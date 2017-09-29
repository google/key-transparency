// Code generated by protoc-gen-go. DO NOT EDIT.
// source: keytransparency_v1_types.proto

/*
Package keytransparency_v1_types is a generated protocol buffer package.

Key Transparency Service

The Key Transparency Service API consists of a map of user names to public
keys. Each user name also has a history of public keys that have been
associated with it.

It is generated from these files:
	keytransparency_v1_types.proto

It has these top-level messages:
	Committed
	EntryUpdate
	Entry
	KeyValue
	SignedKV
	MutationProof
	MapperMetadata
	GetEntryRequest
	GetEntryResponse
	ListEntryHistoryRequest
	ListEntryHistoryResponse
	UpdateEntryRequest
	UpdateEntryResponse
	GetMutationsRequest
	GetMutationsResponse
	GetDomainInfoRequest
	GetDomainInfoResponse
	UserProfile
	BatchUpdateEntriesRequest
	BatchUpdateEntriesResponse
	GetEpochsRequest
	GetEpochsResponse
*/
package keytransparency_v1_types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import keyspb "github.com/google/trillian/crypto/keyspb"
import sigpb "github.com/google/trillian/crypto/sigpb"
import trillian "github.com/google/trillian"
import trillian1 "github.com/google/trillian"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Committed represents the data committed to in a cryptographic commitment.
// commitment = HMAC_SHA512_256(key, data)
type Committed struct {
	// key is the 16 byte random commitment key.
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// data is the data being committed to.
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Committed) Reset()                    { *m = Committed{} }
func (m *Committed) String() string            { return proto.CompactTextString(m) }
func (*Committed) ProtoMessage()               {}
func (*Committed) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Committed) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Committed) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// EntryUpdate contains the user entry update(s).
type EntryUpdate struct {
	// update authorizes the change to entry.
	Update *SignedKV `protobuf:"bytes,2,opt,name=update" json:"update,omitempty"`
	// commitment contains the data committed to in update.commitment.
	Committed *Committed `protobuf:"bytes,3,opt,name=committed" json:"committed,omitempty"`
}

func (m *EntryUpdate) Reset()                    { *m = EntryUpdate{} }
func (m *EntryUpdate) String() string            { return proto.CompactTextString(m) }
func (*EntryUpdate) ProtoMessage()               {}
func (*EntryUpdate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *EntryUpdate) GetUpdate() *SignedKV {
	if m != nil {
		return m.Update
	}
	return nil
}

func (m *EntryUpdate) GetCommitted() *Committed {
	if m != nil {
		return m.Committed
	}
	return nil
}

// Entry contains a commitment to profile and a set of authorized update keys.
// Entry is placed in the verifiable map as leaf data.
type Entry struct {
	// commitment is a cryptographic commitment to arbitrary data.
	Commitment []byte `protobuf:"bytes,1,opt,name=commitment,proto3" json:"commitment,omitempty"`
	// authorized_keys is the set of keys allowed to sign updates for this entry.
	AuthorizedKeys []*keyspb.PublicKey `protobuf:"bytes,2,rep,name=authorized_keys,json=authorizedKeys" json:"authorized_keys,omitempty"`
	// previous contains the hash of the previous entry that this mutation is
	// modifying creating a hash chain of all mutations. The hash used is
	// CommonJSON in "github.com/benlaurie/objecthash/go/objecthash".
	Previous []byte `protobuf:"bytes,3,opt,name=previous,proto3" json:"previous,omitempty"`
}

func (m *Entry) Reset()                    { *m = Entry{} }
func (m *Entry) String() string            { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()               {}
func (*Entry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Entry) GetCommitment() []byte {
	if m != nil {
		return m.Commitment
	}
	return nil
}

func (m *Entry) GetAuthorizedKeys() []*keyspb.PublicKey {
	if m != nil {
		return m.AuthorizedKeys
	}
	return nil
}

func (m *Entry) GetPrevious() []byte {
	if m != nil {
		return m.Previous
	}
	return nil
}

// KeyValue is a map entry.
type KeyValue struct {
	// key contains the map entry key.
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// TODO(ismail): make proto.Any
	// value contains the map entry value.
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *KeyValue) Reset()                    { *m = KeyValue{} }
func (m *KeyValue) String() string            { return proto.CompactTextString(m) }
func (*KeyValue) ProtoMessage()               {}
func (*KeyValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *KeyValue) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *KeyValue) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

// SignedKV is a signed change to a map entry.
type SignedKV struct {
	// key_value is a serialized KeyValue.
	KeyValue *KeyValue `protobuf:"bytes,1,opt,name=key_value,json=keyValue" json:"key_value,omitempty"`
	// signatures on key_value. Must be signed by keys from both previous and
	// current epochs. The first proves ownership of new epoch key, and the
	// second proves that the correct owner is making this change.
	Signatures map[string]*sigpb.DigitallySigned `protobuf:"bytes,2,rep,name=signatures" json:"signatures,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *SignedKV) Reset()                    { *m = SignedKV{} }
func (m *SignedKV) String() string            { return proto.CompactTextString(m) }
func (*SignedKV) ProtoMessage()               {}
func (*SignedKV) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *SignedKV) GetKeyValue() *KeyValue {
	if m != nil {
		return m.KeyValue
	}
	return nil
}

func (m *SignedKV) GetSignatures() map[string]*sigpb.DigitallySigned {
	if m != nil {
		return m.Signatures
	}
	return nil
}

// MutationProof contains the information necessary to compute the new leaf value.
// It contains a) the old leaf value with it's inclusion proof and b) the mutation.
// The new leaf value is computed via:
//       Mutate(leaf_value, mutation)
type MutationProof struct {
	// mutation contains the information needed to modify the old leaf.
	// The format of a mutation is specific to the particular Mutate function being used.
	Mutation *SignedKV `protobuf:"bytes,1,opt,name=mutation" json:"mutation,omitempty"`
	// leaf_proof contains the leaf and its inclusion proof for a particular map revision.
	LeafProof *trillian1.MapLeafInclusion `protobuf:"bytes,2,opt,name=leaf_proof,json=leafProof" json:"leaf_proof,omitempty"`
}

func (m *MutationProof) Reset()                    { *m = MutationProof{} }
func (m *MutationProof) String() string            { return proto.CompactTextString(m) }
func (*MutationProof) ProtoMessage()               {}
func (*MutationProof) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *MutationProof) GetMutation() *SignedKV {
	if m != nil {
		return m.Mutation
	}
	return nil
}

func (m *MutationProof) GetLeafProof() *trillian1.MapLeafInclusion {
	if m != nil {
		return m.LeafProof
	}
	return nil
}

// MapperMetadata tracks the mutations that have been mapped so far. It is
// embedded in the Trillian SignedMapHead.
type MapperMetadata struct {
	HighestFullyCompletedSeq int64 `protobuf:"varint,1,opt,name=highest_fully_completed_seq,json=highestFullyCompletedSeq" json:"highest_fully_completed_seq,omitempty"`
}

func (m *MapperMetadata) Reset()                    { *m = MapperMetadata{} }
func (m *MapperMetadata) String() string            { return proto.CompactTextString(m) }
func (*MapperMetadata) ProtoMessage()               {}
func (*MapperMetadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *MapperMetadata) GetHighestFullyCompletedSeq() int64 {
	if m != nil {
		return m.HighestFullyCompletedSeq
	}
	return 0
}

// GetEntryRequest for a user object.
type GetEntryRequest struct {
	// user_id is the user identifier. Most commonly an email address.
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	// app_id is the identifier for the application.
	AppId string `protobuf:"bytes,2,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	// first_tree_size is the tree_size of the currently trusted log root.
	// Omitting this field will omit the log consistency proof from the response.
	FirstTreeSize int64 `protobuf:"varint,3,opt,name=first_tree_size,json=firstTreeSize" json:"first_tree_size,omitempty"`
}

func (m *GetEntryRequest) Reset()                    { *m = GetEntryRequest{} }
func (m *GetEntryRequest) String() string            { return proto.CompactTextString(m) }
func (*GetEntryRequest) ProtoMessage()               {}
func (*GetEntryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *GetEntryRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *GetEntryRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *GetEntryRequest) GetFirstTreeSize() int64 {
	if m != nil {
		return m.FirstTreeSize
	}
	return 0
}

// GetEntryResponse returns a requested user entry.
type GetEntryResponse struct {
	// vrf_proof is the proof for VRF on user_id.
	VrfProof []byte `protobuf:"bytes,1,opt,name=vrf_proof,json=vrfProof,proto3" json:"vrf_proof,omitempty"`
	// committed contains the profile for this account and connects the data
	// in profile to the commitment in leaf_proof.
	Committed *Committed `protobuf:"bytes,2,opt,name=committed" json:"committed,omitempty"`
	// leaf_proof contains an Entry and an inclusion proof in the sparse Merkle
	// Tree.
	LeafProof *trillian1.MapLeafInclusion `protobuf:"bytes,3,opt,name=leaf_proof,json=leafProof" json:"leaf_proof,omitempty"`
	// smr contains the signed map head for the sparse Merkle Tree.
	// smr is also stored in the append only log.
	Smr *trillian.SignedMapRoot `protobuf:"bytes,4,opt,name=smr" json:"smr,omitempty"`
	// log_root is the latest globally consistent log root.
	// TODO: gossip the log root to verify global consistency.
	LogRoot *trillian.SignedLogRoot `protobuf:"bytes,5,opt,name=log_root,json=logRoot" json:"log_root,omitempty"`
	// log_consistency proves that log_root is consistent with previously seen roots.
	LogConsistency [][]byte `protobuf:"bytes,6,rep,name=log_consistency,json=logConsistency,proto3" json:"log_consistency,omitempty"`
	// log_inclusion proves that smr is part of log_root at index=srm.MapRevision.
	LogInclusion [][]byte `protobuf:"bytes,7,rep,name=log_inclusion,json=logInclusion,proto3" json:"log_inclusion,omitempty"`
}

func (m *GetEntryResponse) Reset()                    { *m = GetEntryResponse{} }
func (m *GetEntryResponse) String() string            { return proto.CompactTextString(m) }
func (*GetEntryResponse) ProtoMessage()               {}
func (*GetEntryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *GetEntryResponse) GetVrfProof() []byte {
	if m != nil {
		return m.VrfProof
	}
	return nil
}

func (m *GetEntryResponse) GetCommitted() *Committed {
	if m != nil {
		return m.Committed
	}
	return nil
}

func (m *GetEntryResponse) GetLeafProof() *trillian1.MapLeafInclusion {
	if m != nil {
		return m.LeafProof
	}
	return nil
}

func (m *GetEntryResponse) GetSmr() *trillian.SignedMapRoot {
	if m != nil {
		return m.Smr
	}
	return nil
}

func (m *GetEntryResponse) GetLogRoot() *trillian.SignedLogRoot {
	if m != nil {
		return m.LogRoot
	}
	return nil
}

func (m *GetEntryResponse) GetLogConsistency() [][]byte {
	if m != nil {
		return m.LogConsistency
	}
	return nil
}

func (m *GetEntryResponse) GetLogInclusion() [][]byte {
	if m != nil {
		return m.LogInclusion
	}
	return nil
}

// ListEntryHistoryRequest gets a list of historical keys for a user.
type ListEntryHistoryRequest struct {
	// user_id is the user identifier.
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	// start is the starting epoch.
	Start int64 `protobuf:"varint,2,opt,name=start" json:"start,omitempty"`
	// page_size is the maximum number of entries to return.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
	// app_id is the identifier for the application.
	AppId string `protobuf:"bytes,4,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	// first_tree_size is the tree_size of the currently trusted log root.
	// Omitting this field will omit the log consistency proof from the response.
	FirstTreeSize int64 `protobuf:"varint,5,opt,name=first_tree_size,json=firstTreeSize" json:"first_tree_size,omitempty"`
}

func (m *ListEntryHistoryRequest) Reset()                    { *m = ListEntryHistoryRequest{} }
func (m *ListEntryHistoryRequest) String() string            { return proto.CompactTextString(m) }
func (*ListEntryHistoryRequest) ProtoMessage()               {}
func (*ListEntryHistoryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ListEntryHistoryRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *ListEntryHistoryRequest) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *ListEntryHistoryRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListEntryHistoryRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *ListEntryHistoryRequest) GetFirstTreeSize() int64 {
	if m != nil {
		return m.FirstTreeSize
	}
	return 0
}

// ListEntryHistoryResponse requests a paginated history of keys for a user.
type ListEntryHistoryResponse struct {
	// values represents the list of keys this user_id has contained over time.
	Values []*GetEntryResponse `protobuf:"bytes,1,rep,name=values" json:"values,omitempty"`
	// next_start is the next page token to query for pagination.
	// next_start is 0 when there are no more results to fetch.
	NextStart int64 `protobuf:"varint,2,opt,name=next_start,json=nextStart" json:"next_start,omitempty"`
}

func (m *ListEntryHistoryResponse) Reset()                    { *m = ListEntryHistoryResponse{} }
func (m *ListEntryHistoryResponse) String() string            { return proto.CompactTextString(m) }
func (*ListEntryHistoryResponse) ProtoMessage()               {}
func (*ListEntryHistoryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *ListEntryHistoryResponse) GetValues() []*GetEntryResponse {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *ListEntryHistoryResponse) GetNextStart() int64 {
	if m != nil {
		return m.NextStart
	}
	return 0
}

// UpdateEntryRequest updates a user's profile.
type UpdateEntryRequest struct {
	// user_id specifies the id for the user who's profile is being updated.
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	// app_id is the identifier for the application.
	AppId string `protobuf:"bytes,2,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	// first_tree_size is the tree_size of the currently trusted log root.
	// Omitting this field will omit the log consistency proof from the response.
	FirstTreeSize int64 `protobuf:"varint,3,opt,name=first_tree_size,json=firstTreeSize" json:"first_tree_size,omitempty"`
	// entry_update contains the user submitted update.
	EntryUpdate *EntryUpdate `protobuf:"bytes,4,opt,name=entry_update,json=entryUpdate" json:"entry_update,omitempty"`
}

func (m *UpdateEntryRequest) Reset()                    { *m = UpdateEntryRequest{} }
func (m *UpdateEntryRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateEntryRequest) ProtoMessage()               {}
func (*UpdateEntryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *UpdateEntryRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *UpdateEntryRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *UpdateEntryRequest) GetFirstTreeSize() int64 {
	if m != nil {
		return m.FirstTreeSize
	}
	return 0
}

func (m *UpdateEntryRequest) GetEntryUpdate() *EntryUpdate {
	if m != nil {
		return m.EntryUpdate
	}
	return nil
}

// UpdateEntryResponse contains a proof once the update has been included in
// the Merkle Tree.
type UpdateEntryResponse struct {
	// proof contains a proof that the update has been included in the tree.
	Proof *GetEntryResponse `protobuf:"bytes,1,opt,name=proof" json:"proof,omitempty"`
}

func (m *UpdateEntryResponse) Reset()                    { *m = UpdateEntryResponse{} }
func (m *UpdateEntryResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateEntryResponse) ProtoMessage()               {}
func (*UpdateEntryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *UpdateEntryResponse) GetProof() *GetEntryResponse {
	if m != nil {
		return m.Proof
	}
	return nil
}

// GetMutationsRequest contains the input parameters of the GetMutation APIs.
type GetMutationsRequest struct {
	// epoch specifies the epoch number in which mutations will be returned.
	Epoch int64 `protobuf:"varint,1,opt,name=epoch" json:"epoch,omitempty"`
	// first_tree_size is the tree_size of the currently trusted log root.
	// Omitting this field will omit the log consistency proof from the response.
	FirstTreeSize int64 `protobuf:"varint,2,opt,name=first_tree_size,json=firstTreeSize" json:"first_tree_size,omitempty"`
	// page_token defines the starting point for pagination. An empty
	// value means start from the beginning. A non-empty value requests the next
	// page of values.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
	// page_size is the maximum number of epochs to return.
	PageSize int32 `protobuf:"varint,4,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
}

func (m *GetMutationsRequest) Reset()                    { *m = GetMutationsRequest{} }
func (m *GetMutationsRequest) String() string            { return proto.CompactTextString(m) }
func (*GetMutationsRequest) ProtoMessage()               {}
func (*GetMutationsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *GetMutationsRequest) GetEpoch() int64 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

func (m *GetMutationsRequest) GetFirstTreeSize() int64 {
	if m != nil {
		return m.FirstTreeSize
	}
	return 0
}

func (m *GetMutationsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func (m *GetMutationsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

// GetMutationsResponse contains the results of GetMutation APIs.
type GetMutationsResponse struct {
	// epoch specifies the epoch number of the returned mutations.
	Epoch int64 `protobuf:"varint,1,opt,name=epoch" json:"epoch,omitempty"`
	// smr contains the signed map root for the sparse Merkle Tree.
	Smr *trillian.SignedMapRoot `protobuf:"bytes,2,opt,name=smr" json:"smr,omitempty"`
	// log_root is the latest globally consistent log root.
	LogRoot *trillian.SignedLogRoot `protobuf:"bytes,3,opt,name=log_root,json=logRoot" json:"log_root,omitempty"`
	// log_consistency proves that log_root is consistent with previously seen roots.
	LogConsistency [][]byte `protobuf:"bytes,4,rep,name=log_consistency,json=logConsistency,proto3" json:"log_consistency,omitempty"`
	// log_inclusion proves that smr is part of log_root at index=srm.MapRevision.
	LogInclusion [][]byte `protobuf:"bytes,5,rep,name=log_inclusion,json=logInclusion,proto3" json:"log_inclusion,omitempty"`
	// mutation contains mutation information.
	Mutations []*MutationProof `protobuf:"bytes,6,rep,name=mutations" json:"mutations,omitempty"`
	// next_page_token is the next page token to query for pagination.
	// An empty value means there are no more results to fetch.
	// A non-zero value may be used by the client to fetch the next page of
	// results.
	NextPageToken string `protobuf:"bytes,7,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *GetMutationsResponse) Reset()                    { *m = GetMutationsResponse{} }
func (m *GetMutationsResponse) String() string            { return proto.CompactTextString(m) }
func (*GetMutationsResponse) ProtoMessage()               {}
func (*GetMutationsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *GetMutationsResponse) GetEpoch() int64 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

func (m *GetMutationsResponse) GetSmr() *trillian.SignedMapRoot {
	if m != nil {
		return m.Smr
	}
	return nil
}

func (m *GetMutationsResponse) GetLogRoot() *trillian.SignedLogRoot {
	if m != nil {
		return m.LogRoot
	}
	return nil
}

func (m *GetMutationsResponse) GetLogConsistency() [][]byte {
	if m != nil {
		return m.LogConsistency
	}
	return nil
}

func (m *GetMutationsResponse) GetLogInclusion() [][]byte {
	if m != nil {
		return m.LogInclusion
	}
	return nil
}

func (m *GetMutationsResponse) GetMutations() []*MutationProof {
	if m != nil {
		return m.Mutations
	}
	return nil
}

func (m *GetMutationsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// GetDomainInfoRequest contains an empty request to query the GetDomainInfo
// APIs.
type GetDomainInfoRequest struct {
}

func (m *GetDomainInfoRequest) Reset()                    { *m = GetDomainInfoRequest{} }
func (m *GetDomainInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*GetDomainInfoRequest) ProtoMessage()               {}
func (*GetDomainInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

// GetDomainInfoResponse contains the results of GetDomainInfo APIs.
type GetDomainInfoResponse struct {
	// Log contains the Log-Tree's info.
	Log *trillian.Tree `protobuf:"bytes,1,opt,name=log" json:"log,omitempty"`
	// Map contains the Map-Tree's info.
	Map *trillian.Tree `protobuf:"bytes,2,opt,name=map" json:"map,omitempty"`
	// Vrf contains the VRF public key.
	Vrf *keyspb.PublicKey `protobuf:"bytes,3,opt,name=vrf" json:"vrf,omitempty"`
}

func (m *GetDomainInfoResponse) Reset()                    { *m = GetDomainInfoResponse{} }
func (m *GetDomainInfoResponse) String() string            { return proto.CompactTextString(m) }
func (*GetDomainInfoResponse) ProtoMessage()               {}
func (*GetDomainInfoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *GetDomainInfoResponse) GetLog() *trillian.Tree {
	if m != nil {
		return m.Log
	}
	return nil
}

func (m *GetDomainInfoResponse) GetMap() *trillian.Tree {
	if m != nil {
		return m.Map
	}
	return nil
}

func (m *GetDomainInfoResponse) GetVrf() *keyspb.PublicKey {
	if m != nil {
		return m.Vrf
	}
	return nil
}

// UserProfile is the data that a client would like to store on the server.
type UserProfile struct {
	// data is the public key data for the user.
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *UserProfile) Reset()                    { *m = UserProfile{} }
func (m *UserProfile) String() string            { return proto.CompactTextString(m) }
func (*UserProfile) ProtoMessage()               {}
func (*UserProfile) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *UserProfile) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// SetEntriesRequest will update the accounts of the given user_ids to be set to
// user data if those accounts have key_id in their set of authorized_keys.
type BatchUpdateEntriesRequest struct {
	// users is a map from user_ids to user data.
	Users map[string]*UserProfile `protobuf:"bytes,1,rep,name=users" json:"users,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// app_id is the application to make this change for.
	AppId string `protobuf:"bytes,2,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	// key_id is the id of the authorized_public key to use when updating accounts.
	// This must be a key that this server has the private key for.
	KeyId string `protobuf:"bytes,3,opt,name=key_id,json=keyId" json:"key_id,omitempty"`
}

func (m *BatchUpdateEntriesRequest) Reset()                    { *m = BatchUpdateEntriesRequest{} }
func (m *BatchUpdateEntriesRequest) String() string            { return proto.CompactTextString(m) }
func (*BatchUpdateEntriesRequest) ProtoMessage()               {}
func (*BatchUpdateEntriesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

func (m *BatchUpdateEntriesRequest) GetUsers() map[string]*UserProfile {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *BatchUpdateEntriesRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *BatchUpdateEntriesRequest) GetKeyId() string {
	if m != nil {
		return m.KeyId
	}
	return ""
}

// BatchUpdateEntriesResponse returns a list of users for which the set operation
// was unsuccessful.
type BatchUpdateEntriesResponse struct {
	// errors is a map from user_ids to errors, if there was an error for that user.
	Errors map[string]string `protobuf:"bytes,1,rep,name=errors" json:"errors,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *BatchUpdateEntriesResponse) Reset()                    { *m = BatchUpdateEntriesResponse{} }
func (m *BatchUpdateEntriesResponse) String() string            { return proto.CompactTextString(m) }
func (*BatchUpdateEntriesResponse) ProtoMessage()               {}
func (*BatchUpdateEntriesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{19} }

func (m *BatchUpdateEntriesResponse) GetErrors() map[string]string {
	if m != nil {
		return m.Errors
	}
	return nil
}

// GetEpochsRequest is an empty proto message used as input to GetEpochs API.
type GetEpochsRequest struct {
}

func (m *GetEpochsRequest) Reset()                    { *m = GetEpochsRequest{} }
func (m *GetEpochsRequest) String() string            { return proto.CompactTextString(m) }
func (*GetEpochsRequest) ProtoMessage()               {}
func (*GetEpochsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{20} }

// GetEpochsResponse contains mutations of a newly created epoch.
type GetEpochsResponse struct {
	// mutations contains all mutations information of a newly created epoch.
	Mutations *GetMutationsResponse `protobuf:"bytes,1,opt,name=mutations" json:"mutations,omitempty"`
}

func (m *GetEpochsResponse) Reset()                    { *m = GetEpochsResponse{} }
func (m *GetEpochsResponse) String() string            { return proto.CompactTextString(m) }
func (*GetEpochsResponse) ProtoMessage()               {}
func (*GetEpochsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{21} }

func (m *GetEpochsResponse) GetMutations() *GetMutationsResponse {
	if m != nil {
		return m.Mutations
	}
	return nil
}

func init() {
	proto.RegisterType((*Committed)(nil), "keytransparency.v1.types.Committed")
	proto.RegisterType((*EntryUpdate)(nil), "keytransparency.v1.types.EntryUpdate")
	proto.RegisterType((*Entry)(nil), "keytransparency.v1.types.Entry")
	proto.RegisterType((*KeyValue)(nil), "keytransparency.v1.types.KeyValue")
	proto.RegisterType((*SignedKV)(nil), "keytransparency.v1.types.SignedKV")
	proto.RegisterType((*MutationProof)(nil), "keytransparency.v1.types.MutationProof")
	proto.RegisterType((*MapperMetadata)(nil), "keytransparency.v1.types.MapperMetadata")
	proto.RegisterType((*GetEntryRequest)(nil), "keytransparency.v1.types.GetEntryRequest")
	proto.RegisterType((*GetEntryResponse)(nil), "keytransparency.v1.types.GetEntryResponse")
	proto.RegisterType((*ListEntryHistoryRequest)(nil), "keytransparency.v1.types.ListEntryHistoryRequest")
	proto.RegisterType((*ListEntryHistoryResponse)(nil), "keytransparency.v1.types.ListEntryHistoryResponse")
	proto.RegisterType((*UpdateEntryRequest)(nil), "keytransparency.v1.types.UpdateEntryRequest")
	proto.RegisterType((*UpdateEntryResponse)(nil), "keytransparency.v1.types.UpdateEntryResponse")
	proto.RegisterType((*GetMutationsRequest)(nil), "keytransparency.v1.types.GetMutationsRequest")
	proto.RegisterType((*GetMutationsResponse)(nil), "keytransparency.v1.types.GetMutationsResponse")
	proto.RegisterType((*GetDomainInfoRequest)(nil), "keytransparency.v1.types.GetDomainInfoRequest")
	proto.RegisterType((*GetDomainInfoResponse)(nil), "keytransparency.v1.types.GetDomainInfoResponse")
	proto.RegisterType((*UserProfile)(nil), "keytransparency.v1.types.UserProfile")
	proto.RegisterType((*BatchUpdateEntriesRequest)(nil), "keytransparency.v1.types.BatchUpdateEntriesRequest")
	proto.RegisterType((*BatchUpdateEntriesResponse)(nil), "keytransparency.v1.types.BatchUpdateEntriesResponse")
	proto.RegisterType((*GetEpochsRequest)(nil), "keytransparency.v1.types.GetEpochsRequest")
	proto.RegisterType((*GetEpochsResponse)(nil), "keytransparency.v1.types.GetEpochsResponse")
}

func init() { proto.RegisterFile("keytransparency_v1_types.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x57, 0xdd, 0x6e, 0x1b, 0x45,
	0x14, 0xd6, 0x7a, 0xb3, 0x8e, 0x7d, 0x9c, 0x9f, 0x76, 0x9a, 0xb6, 0x8b, 0xab, 0x56, 0x61, 0x2b,
	0x68, 0x41, 0xc8, 0xa8, 0xe6, 0x86, 0x16, 0x51, 0x4a, 0x7f, 0x68, 0xa3, 0x24, 0x22, 0xda, 0x24,
	0x85, 0xbb, 0xd5, 0xc4, 0x1e, 0x3b, 0x23, 0xaf, 0x77, 0x26, 0x33, 0x63, 0x8b, 0x8d, 0x84, 0xe8,
	0x35, 0x42, 0xe2, 0x25, 0x78, 0x01, 0x6e, 0x78, 0x06, 0x9e, 0x81, 0x07, 0xe1, 0x1a, 0xcd, 0xcf,
	0xda, 0xeb, 0xc4, 0x6e, 0x9a, 0x5e, 0x70, 0x93, 0xec, 0x9c, 0x9f, 0x99, 0x73, 0xbe, 0xf3, 0x9d,
	0x99, 0x63, 0xb8, 0x33, 0x20, 0xb9, 0x12, 0x38, 0x93, 0x1c, 0x0b, 0x92, 0x75, 0xf2, 0x64, 0xfc,
	0x20, 0x51, 0x39, 0x27, 0xb2, 0xc5, 0x05, 0x53, 0x0c, 0x85, 0x67, 0xf4, 0xad, 0xf1, 0x83, 0x96,
	0xd1, 0x37, 0x9b, 0x1d, 0x91, 0x73, 0xc5, 0x3e, 0x1f, 0x90, 0x5c, 0xf2, 0x23, 0xf7, 0xcf, 0x7a,
	0x35, 0x43, 0xa7, 0x93, 0xb4, 0xcf, 0x8f, 0xec, 0x5f, 0xa7, 0x59, 0x53, 0x82, 0xa6, 0x29, 0xc5,
	0x99, 0x5b, 0xdf, 0x28, 0xd6, 0xc9, 0x10, 0xf3, 0x04, 0x73, 0x6a, 0xe5, 0xd1, 0x03, 0xa8, 0x3f,
	0x63, 0xc3, 0x21, 0x55, 0x8a, 0x74, 0xd1, 0x15, 0xf0, 0x07, 0x24, 0x0f, 0xbd, 0x4d, 0xef, 0xfe,
	0x4a, 0xac, 0x3f, 0x11, 0x82, 0xa5, 0x2e, 0x56, 0x38, 0xac, 0x18, 0x91, 0xf9, 0x8e, 0x7e, 0xf3,
	0xa0, 0xf1, 0x22, 0x53, 0x22, 0x3f, 0xe4, 0x5d, 0xac, 0x08, 0x7a, 0x04, 0xd5, 0x91, 0xf9, 0x32,
	0x56, 0x8d, 0x76, 0xd4, 0x5a, 0x94, 0x4b, 0x6b, 0x9f, 0xf6, 0x33, 0xd2, 0xdd, 0x7e, 0x1d, 0x3b,
	0x0f, 0xf4, 0x2d, 0xd4, 0x3b, 0xc5, 0xf1, 0xa1, 0x6f, 0xdc, 0xef, 0x2e, 0x76, 0x9f, 0x44, 0x1a,
	0x4f, 0xbd, 0xa2, 0x5f, 0x20, 0x30, 0xd1, 0xa0, 0x3b, 0x00, 0x56, 0x3a, 0x24, 0x99, 0x72, 0x49,
	0x94, 0x24, 0xe8, 0x11, 0xac, 0xe3, 0x91, 0x3a, 0x66, 0x82, 0x9e, 0x92, 0x6e, 0xa2, 0x71, 0x0c,
	0x2b, 0x9b, 0xfe, 0xfd, 0x46, 0xfb, 0x6a, 0xcb, 0x81, 0xba, 0x37, 0x3a, 0x4a, 0x69, 0x67, 0x9b,
	0xe4, 0xf1, 0xda, 0xd4, 0x72, 0x9b, 0xe4, 0x12, 0x35, 0xa1, 0xc6, 0x05, 0x19, 0x53, 0x36, 0x92,
	0x26, 0xcc, 0x95, 0x78, 0xb2, 0x8e, 0xda, 0x50, 0xdb, 0x26, 0xf9, 0x6b, 0x9c, 0x8e, 0xc8, 0x1c,
	0x04, 0x37, 0x20, 0x18, 0x6b, 0x95, 0x83, 0xd0, 0x2e, 0xa2, 0x7f, 0x3d, 0xa8, 0x15, 0x60, 0xa0,
	0x6f, 0xa0, 0x3e, 0x20, 0x79, 0x62, 0xcd, 0xbc, 0x8b, 0x30, 0x2c, 0xce, 0x8a, 0x6b, 0x83, 0xe2,
	0xd4, 0x18, 0x40, 0xd2, 0x7e, 0x86, 0xd5, 0x48, 0x90, 0x22, 0xa9, 0xf6, 0xc5, 0x55, 0x30, 0x1f,
	0xd6, 0xc9, 0x20, 0x18, 0x97, 0x76, 0x69, 0x1e, 0xc2, 0xfa, 0x19, 0x75, 0x39, 0xb9, 0xba, 0x4d,
	0xee, 0xb3, 0x72, 0x72, 0x8d, 0xf6, 0x8d, 0x96, 0xa5, 0xe0, 0x73, 0xda, 0xa7, 0x0a, 0xa7, 0x69,
	0x6e, 0x4f, 0x72, 0x49, 0x3f, 0xaa, 0x7c, 0xe9, 0x45, 0xbf, 0x7a, 0xb0, 0xba, 0x3b, 0x52, 0x58,
	0x51, 0x96, 0xed, 0x09, 0xc6, 0x7a, 0xe8, 0x31, 0xd4, 0x86, 0x4e, 0x70, 0x71, 0xf2, 0x13, 0x02,
	0x4d, 0x7c, 0xd0, 0x43, 0x80, 0x94, 0xe0, 0x5e, 0xc2, 0xf5, 0x6e, 0x2e, 0x90, 0x66, 0x6b, 0x42,
	0xff, 0x5d, 0xcc, 0x77, 0x08, 0xee, 0x6d, 0x65, 0x9d, 0x74, 0x24, 0x29, 0xcb, 0xe2, 0xba, 0xb6,
	0x36, 0x47, 0x47, 0xdf, 0xc3, 0xda, 0x2e, 0xe6, 0x9c, 0x88, 0x5d, 0xa2, 0xb0, 0xe6, 0x36, 0xfa,
	0x1a, 0x6e, 0x1d, 0xd3, 0xfe, 0x31, 0x91, 0x2a, 0xe9, 0x8d, 0xd2, 0x34, 0x4f, 0x3a, 0x6c, 0xc8,
	0x53, 0xa2, 0x48, 0x37, 0x91, 0xe4, 0xc4, 0xc4, 0xe7, 0xc7, 0xa1, 0x33, 0xf9, 0x4e, 0x5b, 0x3c,
	0x2b, 0x0c, 0xf6, 0xc9, 0x49, 0x44, 0x61, 0xfd, 0x25, 0x51, 0x16, 0x4c, 0x72, 0x32, 0x22, 0x52,
	0xa1, 0x9b, 0xb0, 0x3c, 0x92, 0x44, 0x24, 0xb4, 0xeb, 0x80, 0xab, 0xea, 0xe5, 0x56, 0x17, 0x5d,
	0x87, 0x2a, 0xe6, 0x5c, 0xcb, 0x2b, 0x46, 0x1e, 0x60, 0xce, 0xb7, 0xba, 0xe8, 0x63, 0x58, 0xef,
	0x51, 0x21, 0x55, 0xa2, 0x04, 0x21, 0x89, 0xa4, 0xa7, 0xc4, 0x10, 0xce, 0x8f, 0x57, 0x8d, 0xf8,
	0x40, 0x10, 0xb2, 0x4f, 0x4f, 0x49, 0xf4, 0x4f, 0x05, 0xae, 0x4c, 0xcf, 0x92, 0x9c, 0x65, 0x92,
	0xa0, 0x5b, 0x50, 0x1f, 0x8b, 0x02, 0x0a, 0x4b, 0xc2, 0xda, 0x58, 0xd8, 0x6c, 0x67, 0x7b, 0xad,
	0xf2, 0x3e, 0xbd, 0x76, 0x06, 0x6b, 0xff, 0x12, 0x58, 0xa3, 0x4f, 0xc0, 0x97, 0x43, 0x11, 0x2e,
	0x19, 0x9f, 0x9b, 0x53, 0x1f, 0x5b, 0xd1, 0x5d, 0xcc, 0x63, 0xc6, 0x54, 0xac, 0x6d, 0x50, 0x1b,
	0x6a, 0x29, 0xeb, 0x27, 0x82, 0x31, 0x15, 0x06, 0xf3, 0xed, 0x77, 0x58, 0xdf, 0xd8, 0x2f, 0xa7,
	0xf6, 0x03, 0xdd, 0x83, 0x75, 0xed, 0xd3, 0x61, 0x99, 0xa4, 0x52, 0xe9, 0x54, 0xc2, 0xea, 0xa6,
	0x7f, 0x7f, 0x25, 0x5e, 0x4b, 0x59, 0xff, 0xd9, 0x54, 0x8a, 0xee, 0xc2, 0xaa, 0x36, 0xa4, 0x45,
	0x8c, 0xe1, 0xb2, 0x31, 0x5b, 0x49, 0x59, 0x7f, 0x12, 0x77, 0xf4, 0x87, 0x07, 0x37, 0x77, 0xa8,
	0xb4, 0xe8, 0xbe, 0xa2, 0x52, 0xb1, 0x77, 0x28, 0xe8, 0x06, 0x04, 0x52, 0x61, 0xa1, 0x0c, 0xb6,
	0x7e, 0x6c, 0x17, 0xba, 0x24, 0x1c, 0xf7, 0x4b, 0x95, 0x0c, 0xe2, 0x9a, 0x16, 0xe8, 0x22, 0x96,
	0x38, 0xb0, 0x74, 0x01, 0x07, 0x82, 0x79, 0x1c, 0xf8, 0x19, 0xc2, 0xf3, 0x51, 0x3a, 0x2a, 0x3c,
	0x85, 0xaa, 0xe9, 0x3a, 0x19, 0x7a, 0xe6, 0x3e, 0xf8, 0x74, 0x71, 0xa9, 0xcf, 0xd2, 0x28, 0x76,
	0x9e, 0xe8, 0x36, 0x40, 0x46, 0x7e, 0x52, 0x49, 0x39, 0xad, 0xba, 0x96, 0xec, 0x6b, 0x41, 0xf4,
	0x97, 0x07, 0xc8, 0xbe, 0x01, 0xff, 0x07, 0xe3, 0xd1, 0x2b, 0x58, 0x21, 0xfa, 0x9c, 0xc4, 0xbd,
	0x36, 0x96, 0x4a, 0x1f, 0x2d, 0xce, 0xab, 0xf4, 0x48, 0xc5, 0x0d, 0x32, 0x5d, 0x44, 0x3f, 0xc0,
	0xb5, 0x99, 0xb8, 0x1d, 0x64, 0x4f, 0x20, 0x98, 0x76, 0xce, 0xe5, 0x10, 0xb3, 0x8e, 0xd1, 0xef,
	0x1e, 0x5c, 0x7b, 0x49, 0x54, 0x71, 0xc1, 0xc9, 0x02, 0x92, 0x0d, 0x08, 0x08, 0x67, 0x9d, 0x63,
	0x77, 0x81, 0xd8, 0xc5, 0xbc, 0xc4, 0x2b, 0xf3, 0x12, 0xbf, 0x0d, 0x60, 0x28, 0xa4, 0xd8, 0x80,
	0x64, 0x06, 0x9b, 0x7a, 0x6c, 0x48, 0x75, 0xa0, 0x05, 0xb3, 0x0c, 0x5b, 0x9a, 0x65, 0x58, 0xf4,
	0x77, 0x05, 0x36, 0x66, 0x23, 0x72, 0xc9, 0xce, 0x0f, 0xc9, 0x75, 0x69, 0xe5, 0x92, 0x5d, 0xea,
	0xbf, 0x7f, 0x97, 0x2e, 0xbd, 0x5b, 0x97, 0x06, 0xe7, 0xbb, 0x14, 0xbd, 0x80, 0x7a, 0xf1, 0x0a,
	0x48, 0xd3, 0xed, 0x8d, 0xf6, 0xbd, 0xc5, 0x35, 0x9b, 0x79, 0x75, 0xe2, 0xa9, 0xa7, 0x2e, 0x83,
	0x61, 0x79, 0x09, 0xe3, 0x65, 0x83, 0xf1, 0xaa, 0x16, 0xef, 0x15, 0x38, 0x47, 0x37, 0x0c, 0x92,
	0xcf, 0xd9, 0x10, 0xd3, 0x6c, 0x2b, 0xeb, 0x31, 0x57, 0xdc, 0xe8, 0x8d, 0x07, 0xd7, 0xcf, 0x28,
	0x1c, 0xc6, 0x9b, 0xe0, 0xa7, 0xac, 0xef, 0xe8, 0xb4, 0x36, 0x45, 0x47, 0x57, 0x36, 0xd6, 0x2a,
	0x6d, 0x31, 0xc4, 0xdc, 0xe1, 0x7d, 0xce, 0x62, 0x88, 0x39, 0xba, 0x0b, 0xfe, 0x58, 0x14, 0x77,
	0xed, 0x9c, 0x49, 0x45, 0x6b, 0xa3, 0x0f, 0xa1, 0x71, 0x28, 0x89, 0xd8, 0x13, 0xac, 0x47, 0x53,
	0x32, 0x99, 0xda, 0xbc, 0xd2, 0xd4, 0xf6, 0xa6, 0x02, 0x1f, 0x3c, 0xc5, 0xaa, 0x73, 0x3c, 0x65,
	0x3e, 0x25, 0x13, 0x82, 0x1e, 0x40, 0xa0, 0x9b, 0xb4, 0xb8, 0x2c, 0x1e, 0x2f, 0x86, 0x71, 0xe1,
	0x1e, 0x2d, 0x1d, 0x81, 0x1b, 0x24, 0xec, 0x66, 0x8b, 0x1a, 0xfe, 0x3a, 0x54, 0xf5, 0xbc, 0x43,
	0xbb, 0x8e, 0xcb, 0xc1, 0x80, 0xe4, 0x5b, 0xdd, 0x66, 0x02, 0x30, 0xdd, 0x62, 0xce, 0xb0, 0xf1,
	0xd5, 0xec, 0xb0, 0xf1, 0x96, 0xc6, 0x2f, 0x61, 0x51, 0x9e, 0x3d, 0xfe, 0xf4, 0xa0, 0x39, 0x2f,
	0x7c, 0x57, 0xad, 0x1f, 0xa1, 0x4a, 0x84, 0x60, 0x13, 0x10, 0x9e, 0x5c, 0x0e, 0x04, 0xbb, 0x4b,
	0xeb, 0x85, 0xd9, 0xc2, 0xc2, 0xe0, 0xf6, 0x6b, 0x3e, 0x84, 0x46, 0x49, 0x3c, 0x27, 0xb5, 0x99,
	0x21, 0xb1, 0x5e, 0x8e, 0x19, 0xd9, 0x57, 0x5e, 0x37, 0x67, 0x01, 0x74, 0x84, 0xe1, 0x6a, 0x49,
	0xe6, 0xa2, 0xdf, 0x29, 0x37, 0x83, 0x65, 0x5c, 0xeb, 0xad, 0x17, 0xd8, 0xb9, 0x2b, 0xa1, 0xd4,
	0x13, 0x47, 0x55, 0xf3, 0xeb, 0xe0, 0x8b, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x55, 0x41, 0x7e,
	0x41, 0xb7, 0x0c, 0x00, 0x00,
}
