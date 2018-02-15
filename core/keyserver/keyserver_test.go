// Copyright 2018 Google Inc. All Rights Reserved.
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

package keyserver

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/keytransparency/core/domain"
	"github.com/google/keytransparency/core/fake"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/google/keytransparency/core/api/v1/keytransparency_proto"
	tpb "github.com/google/trillian"
	tfake "github.com/google/trillian/testonly/fake"
)

func TestLatestRevision(t *testing.T) {
	ctx := context.Background()
	mapID := int64(2)
	fakeAdmin := fake.NewDomainStorage()
	if err := fakeAdmin.Write(ctx, &domain.Domain{
		DomainID:    domainID,
		MapID:       mapID,
		MinInterval: 1 * time.Second,
		MaxInterval: 5 * time.Second,
	}); err != nil {
		t.Fatalf("admin.Write(): %v", err)
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	s := &tfake.Server{
		MockTrillianMapServer:   tpb.NewMockTrillianMapServer(ctrl),
		MockTrillianLogServer:   tpb.NewMockTrillianLogServer(ctrl),
		MockTrillianAdminServer: tpb.NewMockTrillianAdminServer(ctrl),
	}
	lis, stopFakeServer, err := tfake.StartServer(s)
	if err != nil {
		t.Fatalf("Error starting fake server: %v", err)
	}
	defer stopFakeServer()
	cc, err := grpc.Dial(lis.Addr().String(), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Dial(%v): %v", lis.Addr().String(), err)
	}

	srv := &Server{
		domains: fakeAdmin,
		tlog:    tpb.NewTrillianLogClient(cc),
		tmap:    tpb.NewTrillianMapClient(cc),
		indexFunc: func(context.Context, *domain.Domain, string, string) ([32]byte, []byte, error) {
			return [32]byte{}, []byte(""), nil
		},
	}

	for _, tc := range []struct {
		desc     string
		treeSize int64
		wantErr  codes.Code
		wantRev  int64
	}{
		{desc: "not initialized", treeSize: 0, wantErr: codes.Internal},
		{desc: "log controls revision", treeSize: 2, wantErr: codes.OK, wantRev: 1},
	} {
		t.Run(tc.desc+" GetEntry", func(t *testing.T) {
			resp, err := srv.GetEntry(ctx, &pb.GetEntryRequest{
				DomainId: domainID,
			})
			if got, want := status.Code(err), tc.wantErr; got != want {
				t.Errorf("GetEntry(): %v, want %v", err, want)
			}
			if err != nil {
				return
			}
			if got, want := resp.Smr.MapRevision, tc.wantRev; got != want {
				t.Errorf("GetEntry().Rev: %v, want %v", got, want)
			}
		})
		t.Run(tc.desc+" GetEntryHistory", func(t *testing.T) {
			resp2, err := srv.ListEntryHistory(ctx, &pb.ListEntryHistoryRequest{
				DomainId: domainID,
			})
			if got, want := status.Code(err), tc.wantErr; got != want {
				t.Errorf("ListEntryHistory(): %v, want %v", err, tc.wantErr)
			}
			if err != nil {
				return
			}
			i := len(resp2.Values) - 1 // Get last value.
			if got, want := resp2.Values[i].Smr.MapRevision, tc.wantRev; got != want {
				t.Errorf("ListEntryHistory().Rev: %v, want %v", got, want)
			}
		})
	}

}
