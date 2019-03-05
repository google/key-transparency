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

// Package runner executes the mapper pipeline.
package runner

import (
	"context"
	"fmt"

	"github.com/golang/glog"

	"github.com/google/keytransparency/core/mutator"
	"github.com/google/keytransparency/core/mutator/entry"
	"github.com/google/keytransparency/core/sequencer/mapper"

	pb "github.com/google/keytransparency/core/api/v1/keytransparency_go_proto"
	spb "github.com/google/keytransparency/core/sequencer/sequencer_go_proto"
	tpb "github.com/google/trillian"
)

// ApplyMutations takes the set of mutations and applies them to given leaves.
// Returns a list of map leaves that should be updated.
func ApplyMutations(reduceFn mutator.ReduceMutationFn,
	indexedUpdates []*entry.IndexedValue, leaves []*tpb.MapLeaf, emitErr func(error)) ([]*tpb.MapLeaf, error) {

	indexedLeaves, err := DoMapMapLeafFn(mapper.MapMapLeafFn, leaves)
	if err != nil {
		return nil, err
	}

	joined := Join(indexedLeaves, indexedUpdates)

	ret := make([]*tpb.MapLeaf, 0, len(joined))
	for _, j := range joined {
		reduceFn(j.Values1, j.Values2,
			func(e *pb.EntryUpdate) {
				mapLeaf, err := (&entry.IndexedValue{Index: j.Index, Value: e}).Marshal()
				if err != nil {
					emitErr(err)
				}
				ret = append(ret, mapLeaf)
			},
			func(err error) { emitErr(fmt.Errorf("reduceFn on index %x: %v", j.Index, err)) },
		)
	}
	glog.V(2).Infof("ApplyMutations applied %v mutations to %v leaves", len(indexedUpdates), len(leaves))
	return ret, nil
}

// Joined is the result of a CoGroupByKey on []*MapLeaf and []*IndexedValue.
type Joined struct {
	Index   []byte
	Values1 []*pb.EntryUpdate
	Values2 []*pb.EntryUpdate
}

// Join pairs up MapLeaves and IndexedValue by index.
func Join(leaves []*entry.IndexedValue, msgs []*entry.IndexedValue) []*Joined {
	joinMap := make(map[string]*Joined)
	for _, l := range leaves {
		row, ok := joinMap[string(l.Index)]
		if !ok {
			row = &Joined{Index: l.Index}
		}
		row.Values1 = append(row.Values1, l.Value)
		joinMap[string(l.Index)] = row
	}
	for _, m := range msgs {
		row, ok := joinMap[string(m.Index)]
		if !ok {
			row = &Joined{Index: m.Index}
		}
		row.Values2 = append(row.Values2, m.Value)
		joinMap[string(m.Index)] = row
	}
	ret := make([]*Joined, 0, len(joinMap))
	for _, r := range joinMap {
		ret = append(ret, r)
	}
	return ret
}

// ReadSliceFn emits the log messages referenced by slice.
type ReadSliceFn func(ctx context.Context, slice *spb.MapMetadata_SourceSlice,
	directoryID string, chunkSize int32,
	emit func(*mutator.LogMessage)) error

// DoReadFn runs ReadSliceFn on every source slice and collects the outputs.
func DoReadFn(ctx context.Context, fn ReadSliceFn, slices []*spb.MapMetadata_SourceSlice,
	directoryID string, chunkSize int32) ([]*mutator.LogMessage, error) {
	outs := make([]*mutator.LogMessage, 0, len(slices))
	for _, s := range slices {
		if err := fn(ctx, s, directoryID, chunkSize,
			func(msg *mutator.LogMessage) { outs = append(outs, msg) },
		); err != nil {
			return nil, err
		}
	}
	return outs, nil
}

// MapMetaFn emits a source slice for every map slice.
type MapMetaFn func(meta *spb.MapMetadata, emit func(*spb.MapMetadata_SourceSlice))

// DoMapMetaFn runs MapMetaFn on meta and collects the outputs.
func DoMapMetaFn(fn MapMetaFn, meta *spb.MapMetadata) []*spb.MapMetadata_SourceSlice {
	outs := make([]*spb.MapMetadata_SourceSlice, 0, 1)
	fn(meta, func(slice *spb.MapMetadata_SourceSlice) { outs = append(outs, slice) })
	return outs
}

// DoMapLogItemsFn runs the MapLogItemsFn on each element of msgs.
func DoMapLogItemsFn(fn mutator.MapLogItemFn, msgs []*mutator.LogMessage, emitErr func(error)) []*entry.IndexedValue {
	outs := make([]*entry.IndexedValue, 0, len(msgs))
	for _, m := range msgs {
		fn(m,
			func(index []byte, value *pb.EntryUpdate) {
				outs = append(outs, &entry.IndexedValue{Index: index, Value: value})
			},
			func(err error) { emitErr(fmt.Errorf("mapLogItemFn: %v", err)) },
		)
	}
	return outs
}

// MapMapLeafFn converts an update into an IndexedValue.
type MapMapLeafFn func(*tpb.MapLeaf) (*entry.IndexedValue, error)

func DoMapMapLeafFn(fn MapMapLeafFn, leaves []*tpb.MapLeaf) ([]*entry.IndexedValue, error) {
	outs := make([]*entry.IndexedValue, 0, len(leaves))
	for _, m := range leaves {
		out, err := fn(m)
		if err != nil {
			return nil, err
		}
		outs = append(outs, out)
	}
	return outs, nil
}
