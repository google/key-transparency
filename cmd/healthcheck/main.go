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

// healthcheck makes a HTTP(S) request to the supplied url and exits with 0
// (success) if the HTTP response code was in the 2xx range or 1 (unhealthy).
//
// TLS certificate errors are ignored in order to support self-signed certs.
package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Expected URL as command-line argument")
	}
	url := os.Args[1]
	//nolint:gas
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	//nolint:gas
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("http.Get(%v): %v", url, err)
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		log.Fatalf("HTTP status %v was not in the 2xx range", resp.StatusCode)
	}
}
