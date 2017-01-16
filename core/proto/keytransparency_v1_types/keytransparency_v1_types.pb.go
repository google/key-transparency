// Code generated by protoc-gen-go.
// source: keytransparency_v1_types.proto
// DO NOT EDIT!

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
	Profile
	EntryUpdate
	Entry
	PublicKey
	KeyValue
	SignedKV
	GetEntryRequest
	GetEntryResponse
	ListEntryHistoryRequest
	ListEntryHistoryResponse
	UpdateEntryRequest
	UpdateEntryResponse
*/
package keytransparency_v1_types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import ctmap "github.com/google/key-transparency/core/proto/ctmap"

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

// Profile contains data hidden behind the cryptographic commitment.
type Profile struct {
	// Keys is a map of application IDs to keys.
	Keys map[string][]byte `protobuf:"bytes,1,rep,name=keys" json:"keys,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *Profile) Reset()                    { *m = Profile{} }
func (m *Profile) String() string            { return proto.CompactTextString(m) }
func (*Profile) ProtoMessage()               {}
func (*Profile) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Profile) GetKeys() map[string][]byte {
	if m != nil {
		return m.Keys
	}
	return nil
}

// EntryUpdate contains the user entry update(s).
type EntryUpdate struct {
	// update authorizes the change to profile.
	Update *SignedKV `protobuf:"bytes,2,opt,name=update" json:"update,omitempty"`
	// commitment contains the serialized Profile protobuf.
	Committed *Committed `protobuf:"bytes,3,opt,name=committed" json:"committed,omitempty"`
}

func (m *EntryUpdate) Reset()                    { *m = EntryUpdate{} }
func (m *EntryUpdate) String() string            { return proto.CompactTextString(m) }
func (*EntryUpdate) ProtoMessage()               {}
func (*EntryUpdate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

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
	AuthorizedKeys []*PublicKey `protobuf:"bytes,2,rep,name=authorized_keys,json=authorizedKeys" json:"authorized_keys,omitempty"`
}

func (m *Entry) Reset()                    { *m = Entry{} }
func (m *Entry) String() string            { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()               {}
func (*Entry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Entry) GetCommitment() []byte {
	if m != nil {
		return m.Commitment
	}
	return nil
}

func (m *Entry) GetAuthorizedKeys() []*PublicKey {
	if m != nil {
		return m.AuthorizedKeys
	}
	return nil
}

// PublicKey defines a key this domain uses to sign MapHeads with.
type PublicKey struct {
	// Key formats from Keyczar.
	//
	// Types that are valid to be assigned to KeyType:
	//	*PublicKey_Ed25519
	//	*PublicKey_RsaVerifyingSha256_3072
	//	*PublicKey_EcdsaVerifyingP256
	KeyType isPublicKey_KeyType `protobuf_oneof:"key_type"`
}

func (m *PublicKey) Reset()                    { *m = PublicKey{} }
func (m *PublicKey) String() string            { return proto.CompactTextString(m) }
func (*PublicKey) ProtoMessage()               {}
func (*PublicKey) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type isPublicKey_KeyType interface {
	isPublicKey_KeyType()
}

type PublicKey_Ed25519 struct {
	Ed25519 []byte `protobuf:"bytes,1,opt,name=ed25519,proto3,oneof"`
}
type PublicKey_RsaVerifyingSha256_3072 struct {
	RsaVerifyingSha256_3072 []byte `protobuf:"bytes,2,opt,name=rsa_verifying_sha256_3072,json=rsaVerifyingSha2563072,proto3,oneof"`
}
type PublicKey_EcdsaVerifyingP256 struct {
	EcdsaVerifyingP256 []byte `protobuf:"bytes,3,opt,name=ecdsa_verifying_p256,json=ecdsaVerifyingP256,proto3,oneof"`
}

func (*PublicKey_Ed25519) isPublicKey_KeyType()                 {}
func (*PublicKey_RsaVerifyingSha256_3072) isPublicKey_KeyType() {}
func (*PublicKey_EcdsaVerifyingP256) isPublicKey_KeyType()      {}

func (m *PublicKey) GetKeyType() isPublicKey_KeyType {
	if m != nil {
		return m.KeyType
	}
	return nil
}

func (m *PublicKey) GetEd25519() []byte {
	if x, ok := m.GetKeyType().(*PublicKey_Ed25519); ok {
		return x.Ed25519
	}
	return nil
}

func (m *PublicKey) GetRsaVerifyingSha256_3072() []byte {
	if x, ok := m.GetKeyType().(*PublicKey_RsaVerifyingSha256_3072); ok {
		return x.RsaVerifyingSha256_3072
	}
	return nil
}

func (m *PublicKey) GetEcdsaVerifyingP256() []byte {
	if x, ok := m.GetKeyType().(*PublicKey_EcdsaVerifyingP256); ok {
		return x.EcdsaVerifyingP256
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*PublicKey) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _PublicKey_OneofMarshaler, _PublicKey_OneofUnmarshaler, _PublicKey_OneofSizer, []interface{}{
		(*PublicKey_Ed25519)(nil),
		(*PublicKey_RsaVerifyingSha256_3072)(nil),
		(*PublicKey_EcdsaVerifyingP256)(nil),
	}
}

func _PublicKey_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*PublicKey)
	// key_type
	switch x := m.KeyType.(type) {
	case *PublicKey_Ed25519:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.Ed25519)
	case *PublicKey_RsaVerifyingSha256_3072:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.RsaVerifyingSha256_3072)
	case *PublicKey_EcdsaVerifyingP256:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.EcdsaVerifyingP256)
	case nil:
	default:
		return fmt.Errorf("PublicKey.KeyType has unexpected type %T", x)
	}
	return nil
}

func _PublicKey_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*PublicKey)
	switch tag {
	case 1: // key_type.ed25519
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.KeyType = &PublicKey_Ed25519{x}
		return true, err
	case 2: // key_type.rsa_verifying_sha256_3072
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.KeyType = &PublicKey_RsaVerifyingSha256_3072{x}
		return true, err
	case 3: // key_type.ecdsa_verifying_p256
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.KeyType = &PublicKey_EcdsaVerifyingP256{x}
		return true, err
	default:
		return false, nil
	}
}

func _PublicKey_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*PublicKey)
	// key_type
	switch x := m.KeyType.(type) {
	case *PublicKey_Ed25519:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Ed25519)))
		n += len(x.Ed25519)
	case *PublicKey_RsaVerifyingSha256_3072:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.RsaVerifyingSha256_3072)))
		n += len(x.RsaVerifyingSha256_3072)
	case *PublicKey_EcdsaVerifyingP256:
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.EcdsaVerifyingP256)))
		n += len(x.EcdsaVerifyingP256)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// KeyValue is a map entry.
type KeyValue struct {
	// key contains the map entry key.
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// value contains the map entry value.
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *KeyValue) Reset()                    { *m = KeyValue{} }
func (m *KeyValue) String() string            { return proto.CompactTextString(m) }
func (*KeyValue) ProtoMessage()               {}
func (*KeyValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

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
	Signatures map[string]*ctmap.DigitallySigned `protobuf:"bytes,2,rep,name=signatures" json:"signatures,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// previous contains the hash of the previous entry that this mutation is
	// modifying creating a hash chain of all mutations. The hash used is
	// CommonJSON in "github.com/benlaurie/objecthash/go/objecthash".
	Previous []byte `protobuf:"bytes,3,opt,name=previous,proto3" json:"previous,omitempty"`
}

func (m *SignedKV) Reset()                    { *m = SignedKV{} }
func (m *SignedKV) String() string            { return proto.CompactTextString(m) }
func (*SignedKV) ProtoMessage()               {}
func (*SignedKV) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *SignedKV) GetKeyValue() *KeyValue {
	if m != nil {
		return m.KeyValue
	}
	return nil
}

func (m *SignedKV) GetSignatures() map[string]*ctmap.DigitallySigned {
	if m != nil {
		return m.Signatures
	}
	return nil
}

func (m *SignedKV) GetPrevious() []byte {
	if m != nil {
		return m.Previous
	}
	return nil
}

// GetEntryRequest for a user object.
type GetEntryRequest struct {
	// user_id is the user identifier. Most commonly an email address.
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
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

// GetEntryResponse returns a requested user entry.
type GetEntryResponse struct {
	// vrf is the output of VRF on user_id.
	Vrf []byte `protobuf:"bytes,1,opt,name=vrf,proto3" json:"vrf,omitempty"`
	// vrf_proof is the proof for VRF on user_id.
	VrfProof []byte `protobuf:"bytes,2,opt,name=vrf_proof,json=vrfProof,proto3" json:"vrf_proof,omitempty"`
	// committed contains the profile for this account and connects the data
	// in profile to the commitment in leaf_proof.
	Committed *Committed `protobuf:"bytes,3,opt,name=committed" json:"committed,omitempty"`
	// leaf_proof contains an Entry and an inclusion proof in the sparse Merkle
	// Tree.
	LeafProof *ctmap.GetLeafResponse `protobuf:"bytes,5,opt,name=leaf_proof,json=leafProof" json:"leaf_proof,omitempty"`
	// smh contains the signed map head for the sparse Merkle Tree.
	// smh is also stored in the append only log.
	Smh *ctmap.SignedMapHead `protobuf:"bytes,6,opt,name=smh" json:"smh,omitempty"`
	// smh_sct is the signed certificate timestamp (SCT) showing that SMH was
	// submitted to CT logs.
	// TODO: Support storing smh in multiple logs.
	SmhSct []byte `protobuf:"bytes,7,opt,name=smh_sct,json=smhSct,proto3" json:"smh_sct,omitempty"`
}

func (m *GetEntryResponse) Reset()                    { *m = GetEntryResponse{} }
func (m *GetEntryResponse) String() string            { return proto.CompactTextString(m) }
func (*GetEntryResponse) ProtoMessage()               {}
func (*GetEntryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *GetEntryResponse) GetVrf() []byte {
	if m != nil {
		return m.Vrf
	}
	return nil
}

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

func (m *GetEntryResponse) GetLeafProof() *ctmap.GetLeafResponse {
	if m != nil {
		return m.LeafProof
	}
	return nil
}

func (m *GetEntryResponse) GetSmh() *ctmap.SignedMapHead {
	if m != nil {
		return m.Smh
	}
	return nil
}

func (m *GetEntryResponse) GetSmhSct() []byte {
	if m != nil {
		return m.SmhSct
	}
	return nil
}

// ListEntryHistoryRequest gets a list of historical keys for a user.
type ListEntryHistoryRequest struct {
	// user_id is the user identifier.
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	// start is the starting epcoh.
	Start int64 `protobuf:"varint,2,opt,name=start" json:"start,omitempty"`
	// page_size is the maximum number of entries to return.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
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
	// entry_update contains the user submitted update.
	EntryUpdate *EntryUpdate `protobuf:"bytes,2,opt,name=entry_update,json=entryUpdate" json:"entry_update,omitempty"`
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

func (m *UpdateEntryRequest) GetEntryUpdate() *EntryUpdate {
	if m != nil {
		return m.EntryUpdate
	}
	return nil
}

// UpdateEntryResponse contains a proof once the update has been included in
// the Merkel Tree.
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

func init() {
	proto.RegisterType((*Committed)(nil), "keytransparency.v1.types.Committed")
	proto.RegisterType((*Profile)(nil), "keytransparency.v1.types.Profile")
	proto.RegisterType((*EntryUpdate)(nil), "keytransparency.v1.types.EntryUpdate")
	proto.RegisterType((*Entry)(nil), "keytransparency.v1.types.Entry")
	proto.RegisterType((*PublicKey)(nil), "keytransparency.v1.types.PublicKey")
	proto.RegisterType((*KeyValue)(nil), "keytransparency.v1.types.KeyValue")
	proto.RegisterType((*SignedKV)(nil), "keytransparency.v1.types.SignedKV")
	proto.RegisterType((*GetEntryRequest)(nil), "keytransparency.v1.types.GetEntryRequest")
	proto.RegisterType((*GetEntryResponse)(nil), "keytransparency.v1.types.GetEntryResponse")
	proto.RegisterType((*ListEntryHistoryRequest)(nil), "keytransparency.v1.types.ListEntryHistoryRequest")
	proto.RegisterType((*ListEntryHistoryResponse)(nil), "keytransparency.v1.types.ListEntryHistoryResponse")
	proto.RegisterType((*UpdateEntryRequest)(nil), "keytransparency.v1.types.UpdateEntryRequest")
	proto.RegisterType((*UpdateEntryResponse)(nil), "keytransparency.v1.types.UpdateEntryResponse")
}

func init() { proto.RegisterFile("keytransparency_v1_types.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 791 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x55, 0xcf, 0x6f, 0xe3, 0x44,
	0x14, 0x5e, 0x37, 0x9b, 0x34, 0x7e, 0x59, 0xd1, 0xd5, 0x50, 0xed, 0x86, 0x20, 0x56, 0x2b, 0x23,
	0xd0, 0x6a, 0x81, 0x84, 0x7a, 0x95, 0x2e, 0x2c, 0x87, 0xc2, 0x02, 0xda, 0xa0, 0x16, 0xa9, 0x9a,
	0xa8, 0xe5, 0x68, 0x4d, 0xed, 0x17, 0x67, 0x94, 0xf8, 0x07, 0x33, 0x63, 0x83, 0x2b, 0x71, 0xe2,
	0xca, 0x89, 0xbf, 0x81, 0x3f, 0x93, 0x03, 0x9a, 0x19, 0x3b, 0x75, 0x4b, 0x43, 0x41, 0xe2, 0x12,
	0xcd, 0x7b, 0xf3, 0x7d, 0xef, 0xc7, 0xf7, 0x9e, 0x27, 0xf0, 0x64, 0x85, 0x95, 0x12, 0x2c, 0x95,
	0x39, 0x13, 0x98, 0x86, 0x55, 0x50, 0x1e, 0x04, 0xaa, 0xca, 0x51, 0x8e, 0x73, 0x91, 0xa9, 0x8c,
	0x0c, 0x6f, 0xdc, 0x8f, 0xcb, 0x83, 0xb1, 0xb9, 0x1f, 0x1d, 0xc5, 0x5c, 0x2d, 0x8b, 0x8b, 0x71,
	0x98, 0x25, 0x93, 0x38, 0xcb, 0xe2, 0x35, 0x4e, 0x56, 0x58, 0x7d, 0xd2, 0x06, 0x4f, 0xc2, 0x4c,
	0xe0, 0xc4, 0x04, 0x9a, 0x84, 0x2a, 0x61, 0xb9, 0xfd, 0xb5, 0xa1, 0xbd, 0x03, 0x70, 0xbf, 0xce,
	0x92, 0x84, 0x2b, 0x85, 0x11, 0x79, 0x08, 0x9d, 0x15, 0x56, 0x43, 0xe7, 0xa9, 0xf3, 0xec, 0x01,
	0xd5, 0x47, 0x42, 0xe0, 0x7e, 0xc4, 0x14, 0x1b, 0xee, 0x18, 0x97, 0x39, 0x7b, 0xbf, 0x3a, 0xb0,
	0x7b, 0x2a, 0xb2, 0x05, 0x5f, 0x23, 0x39, 0x82, 0xfb, 0x2b, 0xac, 0xe4, 0xd0, 0x79, 0xda, 0x79,
	0x36, 0xf0, 0x3f, 0x1a, 0x6f, 0x2b, 0x74, 0x5c, 0x13, 0xc6, 0xc7, 0x58, 0xc9, 0x6f, 0x53, 0x25,
	0x2a, 0x6a, 0x88, 0xa3, 0x97, 0xe0, 0x6e, 0x5c, 0xed, 0xfc, 0xae, 0xcd, 0xbf, 0x0f, 0xdd, 0x92,
	0xad, 0x0b, 0xac, 0x0b, 0xb0, 0xc6, 0xab, 0x9d, 0xcf, 0x1c, 0xef, 0x37, 0x07, 0x06, 0x86, 0x75,
	0x96, 0x47, 0x4c, 0x21, 0x79, 0x05, 0xbd, 0xc2, 0x9c, 0x0c, 0x74, 0xe0, 0x7b, 0xdb, 0x6b, 0x99,
	0xf3, 0x38, 0xc5, 0xe8, 0xf8, 0x9c, 0xd6, 0x0c, 0xf2, 0x15, 0xb8, 0x61, 0x23, 0xc2, 0xb0, 0x63,
	0xe8, 0xef, 0x6f, 0xa7, 0x6f, 0xf4, 0xa2, 0x57, 0x2c, 0xaf, 0x80, 0xae, 0xed, 0xe1, 0x09, 0x80,
	0xf5, 0x26, 0x98, 0xaa, 0x5a, 0xca, 0x96, 0x87, 0x9c, 0xc0, 0x1e, 0x2b, 0xd4, 0x32, 0x13, 0xfc,
	0x12, 0xa3, 0xc0, 0x88, 0xb7, 0x63, 0xc4, 0xfb, 0x87, 0x8c, 0xa7, 0xc5, 0xc5, 0x9a, 0x87, 0xc7,
	0x58, 0xd1, 0xb7, 0xae, 0xb8, 0x5a, 0x36, 0xef, 0x0f, 0x07, 0xdc, 0xcd, 0x2d, 0x19, 0xc1, 0x2e,
	0x46, 0xfe, 0x74, 0x7a, 0xf0, 0xb9, 0x4d, 0x3c, 0xbb, 0x47, 0x1b, 0x07, 0xf9, 0x02, 0xde, 0x11,
	0x92, 0x05, 0x25, 0x0a, 0xbe, 0xa8, 0x78, 0x1a, 0x07, 0x72, 0xc9, 0xfc, 0xe9, 0x61, 0xf0, 0xe2,
	0xd3, 0x97, 0xbe, 0x55, 0x77, 0x76, 0x8f, 0x3e, 0x12, 0x92, 0x9d, 0x37, 0x88, 0xb9, 0x01, 0xe8,
	0x7b, 0xe2, 0xc3, 0x3e, 0x86, 0xd1, 0x35, 0x7a, 0xee, 0x4f, 0x0f, 0x8d, 0x56, 0x9a, 0x47, 0xcc,
	0xed, 0x86, 0x79, 0xea, 0x4f, 0x0f, 0x5f, 0x03, 0xf4, 0x57, 0x58, 0x99, 0x3d, 0xf6, 0x7c, 0xe8,
	0x1f, 0x63, 0x75, 0xae, 0x87, 0x77, 0xcb, 0x92, 0xdd, 0x3a, 0x64, 0xef, 0xf7, 0x1d, 0xe8, 0x37,
	0x93, 0x22, 0x47, 0xe0, 0xea, 0x60, 0x16, 0xe6, 0xdc, 0x35, 0xe0, 0x26, 0x17, 0xd5, 0x15, 0xd8,
	0xac, 0x14, 0x40, 0xf2, 0x38, 0x65, 0xaa, 0x10, 0xd8, 0x28, 0xee, 0xdf, 0xbd, 0x22, 0xe6, 0x60,
	0x49, 0x76, 0x6b, 0x5b, 0x51, 0xc8, 0x08, 0xfa, 0xb9, 0xc0, 0x92, 0x67, 0x85, 0xb4, 0x4a, 0xd0,
	0x8d, 0x3d, 0x3a, 0x83, 0xbd, 0x1b, 0xd4, 0x5b, 0xb6, 0xfb, 0xe3, 0x76, 0xe3, 0x03, 0xff, 0xd1,
	0xd8, 0x7e, 0x99, 0xdf, 0xf0, 0x98, 0x2b, 0xb6, 0x5e, 0x57, 0xb6, 0x8a, 0xf6, 0xd6, 0x3f, 0x87,
	0xbd, 0x37, 0xa8, 0x6c, 0x29, 0xf8, 0x63, 0x81, 0x52, 0x91, 0xc7, 0xb0, 0x5b, 0x48, 0x14, 0x01,
	0x8f, 0xea, 0xd0, 0x3d, 0x6d, 0x7e, 0x17, 0x79, 0x7f, 0x3a, 0xf0, 0xf0, 0x0a, 0x2c, 0xf3, 0x2c,
	0x95, 0x46, 0xfd, 0x52, 0x2c, 0x1a, 0xf5, 0x4b, 0xb1, 0x20, 0xef, 0x82, 0x5b, 0x8a, 0x45, 0x90,
	0x8b, 0x2c, 0x5b, 0xd4, 0x13, 0xe8, 0x97, 0x62, 0x71, 0xaa, 0xed, 0xff, 0xe1, 0xcb, 0x20, 0x53,
	0x80, 0x35, 0xb2, 0x26, 0x41, 0xf7, 0x5a, 0xa7, 0x6f, 0x50, 0x9d, 0x20, 0x5b, 0x34, 0xd5, 0x51,
	0x57, 0x23, 0x6d, 0xe6, 0x0f, 0xa1, 0x23, 0x93, 0xe5, 0xb0, 0x67, 0xf0, 0xfb, 0x35, 0xde, 0x0a,
	0xf2, 0x3d, 0xcb, 0x67, 0xc8, 0x22, 0xaa, 0x01, 0xba, 0x7d, 0x99, 0x2c, 0x03, 0x19, 0xaa, 0xe1,
	0xae, 0x29, 0xbe, 0x27, 0x93, 0xe5, 0x3c, 0x54, 0x1e, 0xc2, 0xe3, 0x13, 0x2e, 0x6d, 0xfb, 0x33,
	0x2e, 0x55, 0x76, 0xb7, 0x64, 0x7a, 0x13, 0xa5, 0x62, 0x42, 0x19, 0x1d, 0x3a, 0xd4, 0x1a, 0x5a,
	0xa1, 0x9c, 0xc5, 0x18, 0x48, 0x7e, 0x89, 0x46, 0x84, 0x2e, 0xed, 0x6b, 0xc7, 0x9c, 0x5f, 0xa2,
	0xf7, 0x0b, 0x0c, 0xff, 0x9e, 0xa6, 0x16, 0xfb, 0x35, 0xf4, 0xcc, 0xe8, 0x9a, 0xf7, 0xf1, 0xf9,
	0x76, 0xe9, 0x6e, 0x0e, 0x8a, 0xd6, 0x4c, 0xf2, 0x1e, 0x40, 0x8a, 0x3f, 0xab, 0xa0, 0x5d, 0x97,
	0xab, 0x3d, 0x73, 0xed, 0xf0, 0x7e, 0x02, 0x62, 0x1f, 0xc0, 0x7f, 0xb5, 0x13, 0x64, 0x06, 0x0f,
	0x50, 0x03, 0x83, 0x6b, 0x6f, 0xe5, 0x07, 0xdb, 0xeb, 0x6a, 0x3d, 0xb1, 0x74, 0x80, 0x57, 0x86,
	0xf7, 0x03, 0xbc, 0x7d, 0x2d, 0x71, 0xdd, 0xf2, 0x97, 0xd0, 0xb5, 0x83, 0xb6, 0x1f, 0xe9, 0x7f,
	0xe9, 0xd8, 0x12, 0x2f, 0x7a, 0xe6, 0x8f, 0xe9, 0xc5, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xeb,
	0x97, 0x30, 0xd7, 0x15, 0x07, 0x00, 0x00,
}
