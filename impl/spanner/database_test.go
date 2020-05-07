// Copyright 2020 Google Inc. All Rights Reserved.
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

// Package spanner provides functions for interacting with cloudspanner.
package spanner

import (
	"context"
	"testing"
)

func TestNew(t *testing.T) {
	ctx := context.Background()
	dbName := "projects/your-project-id/instances/your-instance-id/databases/your-database-id"
	db, err := New(ctx, dbName)
	if err != nil {
		t.Fatal(err)
	}
	client := db.Get()
	defer client.Close()
}
