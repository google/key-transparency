// Copyright 2016 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"database/sql"
	"flag"
	"net"
	"time"

	"github.com/google/keytransparency/core/adminserver"
	"github.com/google/keytransparency/core/mutator"
	"github.com/google/keytransparency/core/sequencer"
	"github.com/google/keytransparency/impl/sql/domain"
	"github.com/google/keytransparency/impl/sql/engine"
	"github.com/google/keytransparency/impl/sql/mutationstorage"

	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	"github.com/google/trillian"
	"google.golang.org/grpc"

	"github.com/google/trillian/crypto/keys/der"
	"github.com/google/trillian/crypto/keyspb"
	"github.com/google/trillian/monitoring/prometheus"

	spb "github.com/google/keytransparency/core/sequencer/sequencer_go_proto"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
)

var (
	serverDBPath = flag.String("db", "db", "Database connection string")

	// Info to connect to the trillian map and log.
	mapURL    = flag.String("map-url", "", "URL of Trillian Map Server")
	logURL    = flag.String("log-url", "", "URL of Trillian Log Server for Signed Map Heads")
	refresh   = flag.Duration("domain-refresh", 5*time.Second, "Time to detect new domain")
	batchSize = flag.Int("batch-size", 100, "Maximum number of mutations to process per map revision")
)

func openDB() *sql.DB {
	db, err := sql.Open(engine.DriverName, *serverDBPath)
	if err != nil {
		glog.Exitf("sql.Open(): %v", err)
	}
	if err := db.Ping(); err != nil {
		glog.Exitf("db.Ping(): %v", err)
	}
	return db
}

func main() {
	flag.Parse()

	// Connect to trillian log and map backends.
	mconn, err := grpc.Dial(*mapURL, grpc.WithInsecure())
	if err != nil {
		glog.Exitf("grpc.Dial(%v): %v", *mapURL, err)
	}
	lconn, err := grpc.Dial(*logURL, grpc.WithInsecure())
	if err != nil {
		glog.Exitf("Failed to connect to %v: %v", *logURL, err)
	}
	tlog := trillian.NewTrillianLogClient(lconn)
	tmap := trillian.NewTrillianMapClient(mconn)
	logAdmin := trillian.NewTrillianAdminClient(lconn)
	mapAdmin := trillian.NewTrillianAdminClient(mconn)

	// Database tables
	sqldb := openDB()
	defer sqldb.Close()

	mutations, err := mutationstorage.New(sqldb)
	if err != nil {
		glog.Exitf("Failed to create mutations object: %v", err)
	}
	domainStorage, err := domain.NewStorage(sqldb)
	if err != nil {
		glog.Exitf("Failed to create domain storage object: %v", err)
	}
	queue := mutator.MutationQueue(mutations)

	// Create server
	sequencerServer := sequencer.NewServer(
		domainStorage,
		logAdmin, mapAdmin,
		tlog, tmap,
		mutations,
		prometheus.MetricFactory{},
	)
	grpcServer := grpc.NewServer()
	grpc_prometheus.Register(grpcServer)
	grpc_prometheus.EnableHandlingTimeHistogram()
	spb.RegisterKeyTransparencySequencerServer(grpcServer, sequencerServer)
	lis, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		glog.Errorf("net.Listen('localhost:0'): %v", err)
	}
	go grpcServer.Serve(lis)
	sequencerConn, err := grpc.Dial(lis.Addr().String(), grpc.WithInsecure())
	if err != nil {
		glog.Exitf("grpc.Dial(%v): %v", *mapURL, err)
	}

	//sequencerClient
	signer := sequencer.New(
		spb.NewKeyTransparencySequencerClient(sequencerConn),
		tlog, tmap,
		mapAdmin,
		domainStorage,
		mutations,
		queue,
		*batchSize)

	keygen := func(ctx context.Context, spec *keyspb.Specification) (proto.Message, error) {
		return der.NewProtoFromSpec(spec)
	}
	adminServer := adminserver.New(tlog, tmap, logAdmin, mapAdmin, domainStorage, keygen)
	glog.Infof("Signer starting")

	// Run servers
	httpServer := startHTTPServer(adminServer)

	cctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if err := signer.ListenForNewDomains(cctx, *refresh); err != nil {
		glog.Errorf("StartSequencingAll(): %v", err)
	}
	httpServer.Shutdown(cctx)
	glog.Errorf("Signer exiting")
}
