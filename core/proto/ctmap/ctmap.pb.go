// Code generated by protoc-gen-go.
// source: ctmap.proto
// DO NOT EDIT!

/*
Package ctmap is a generated protocol buffer package.

It is generated from these files:
	ctmap.proto

It has these top-level messages:
	MapHead
	SignedMapHead
	DigitallySigned
	GetLeafRequest
	GetLeafResponse
	UpdateLeafRequest
	UpdateLeafResponse
	MutationEntry
*/
package ctmap

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// HashAlgorithm defines the approved methods for object hashing.
type DigitallySigned_HashAlgorithm int32

const (
	// No hash algorithm is used.
	DigitallySigned_NONE DigitallySigned_HashAlgorithm = 0
	// SHA256 is used.
	DigitallySigned_SHA256 DigitallySigned_HashAlgorithm = 4
	// SHA512 is used.
	DigitallySigned_SHA512 DigitallySigned_HashAlgorithm = 6
)

var DigitallySigned_HashAlgorithm_name = map[int32]string{
	0: "NONE",
	4: "SHA256",
	6: "SHA512",
}
var DigitallySigned_HashAlgorithm_value = map[string]int32{
	"NONE":   0,
	"SHA256": 4,
	"SHA512": 6,
}

func (x DigitallySigned_HashAlgorithm) String() string {
	return proto.EnumName(DigitallySigned_HashAlgorithm_name, int32(x))
}
func (DigitallySigned_HashAlgorithm) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{2, 0}
}

// SignatureAlgorithm defines the algorithm used to sign the object.
type DigitallySigned_SignatureAlgorithm int32

const (
	// Anonymous signature scheme.
	DigitallySigned_ANONYMOUS DigitallySigned_SignatureAlgorithm = 0
	// ECDSA signature scheme.
	DigitallySigned_ECDSA DigitallySigned_SignatureAlgorithm = 1
)

var DigitallySigned_SignatureAlgorithm_name = map[int32]string{
	0: "ANONYMOUS",
	1: "ECDSA",
}
var DigitallySigned_SignatureAlgorithm_value = map[string]int32{
	"ANONYMOUS": 0,
	"ECDSA":     1,
}

func (x DigitallySigned_SignatureAlgorithm) String() string {
	return proto.EnumName(DigitallySigned_SignatureAlgorithm_name, int32(x))
}
func (DigitallySigned_SignatureAlgorithm) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{2, 1}
}

// MapHead is the head node of the Merkle Tree as well as additional metadata
// for the tree.
type MapHead struct {
	// realm is the domain identifier for the transparent map.
	Realm string `protobuf:"bytes,1,opt,name=realm" json:"realm,omitempty"`
	// epoch is the epoch number of this map head.
	Epoch int64 `protobuf:"varint,2,opt,name=epoch" json:"epoch,omitempty"`
	// root is the value of the root node of the Merkle tree.
	Root []byte `protobuf:"bytes,3,opt,name=root,proto3" json:"root,omitempty"`
	// issue_time is the time when this epoch was created. Monotonically
	// increasing.
	IssueTime *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=issue_time,json=issueTime" json:"issue_time,omitempty"`
}

func (m *MapHead) Reset()                    { *m = MapHead{} }
func (m *MapHead) String() string            { return proto.CompactTextString(m) }
func (*MapHead) ProtoMessage()               {}
func (*MapHead) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *MapHead) GetRealm() string {
	if m != nil {
		return m.Realm
	}
	return ""
}

func (m *MapHead) GetEpoch() int64 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

func (m *MapHead) GetRoot() []byte {
	if m != nil {
		return m.Root
	}
	return nil
}

func (m *MapHead) GetIssueTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.IssueTime
	}
	return nil
}

// SignedMapHead represents a signed state of the Merkel Tree.
type SignedMapHead struct {
	// map_head contains the head node of the Merkle Tree along with other
	// metadata.
	MapHead *MapHead `protobuf:"bytes,1,opt,name=map_head,json=mapHead" json:"map_head,omitempty"`
	// signatures is a set of map_head signatures. Each signature is identified by
	// the first 64 bits of the public key that verifies it.
	Signatures map[string]*DigitallySigned `protobuf:"bytes,2,rep,name=signatures" json:"signatures,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *SignedMapHead) Reset()                    { *m = SignedMapHead{} }
func (m *SignedMapHead) String() string            { return proto.CompactTextString(m) }
func (*SignedMapHead) ProtoMessage()               {}
func (*SignedMapHead) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SignedMapHead) GetMapHead() *MapHead {
	if m != nil {
		return m.MapHead
	}
	return nil
}

func (m *SignedMapHead) GetSignatures() map[string]*DigitallySigned {
	if m != nil {
		return m.Signatures
	}
	return nil
}

// DigitallySigned defines a way to digitally sign objects.
type DigitallySigned struct {
	// hash_algorithm contains the hash algorithm used.
	HashAlgorithm DigitallySigned_HashAlgorithm `protobuf:"varint,1,opt,name=hash_algorithm,json=hashAlgorithm,enum=ctmap.DigitallySigned_HashAlgorithm" json:"hash_algorithm,omitempty"`
	// sig_algorithm contains the signing algorithm used.
	SigAlgorithm DigitallySigned_SignatureAlgorithm `protobuf:"varint,2,opt,name=sig_algorithm,json=sigAlgorithm,enum=ctmap.DigitallySigned_SignatureAlgorithm" json:"sig_algorithm,omitempty"`
	// signature contains the object signature.
	Signature []byte `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *DigitallySigned) Reset()                    { *m = DigitallySigned{} }
func (m *DigitallySigned) String() string            { return proto.CompactTextString(m) }
func (*DigitallySigned) ProtoMessage()               {}
func (*DigitallySigned) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DigitallySigned) GetHashAlgorithm() DigitallySigned_HashAlgorithm {
	if m != nil {
		return m.HashAlgorithm
	}
	return DigitallySigned_NONE
}

func (m *DigitallySigned) GetSigAlgorithm() DigitallySigned_SignatureAlgorithm {
	if m != nil {
		return m.SigAlgorithm
	}
	return DigitallySigned_ANONYMOUS
}

func (m *DigitallySigned) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

// GetLeafRequest for a verifiable map leaf.
type GetLeafRequest struct {
	// index includes the requested leaf index.
	Index []byte `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	// epoch contains the epoch at which the leaf is requested.
	Epoch uint64 `protobuf:"varint,2,opt,name=epoch" json:"epoch,omitempty"`
}

func (m *GetLeafRequest) Reset()                    { *m = GetLeafRequest{} }
func (m *GetLeafRequest) String() string            { return proto.CompactTextString(m) }
func (*GetLeafRequest) ProtoMessage()               {}
func (*GetLeafRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetLeafRequest) GetIndex() []byte {
	if m != nil {
		return m.Index
	}
	return nil
}

func (m *GetLeafRequest) GetEpoch() uint64 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

// GetLeafResponse for a verifiable map leaf.
type GetLeafResponse struct {
	// leaf_data contains an entry stored in the leaf node.
	LeafData []byte `protobuf:"bytes,1,opt,name=leaf_data,json=leafData,proto3" json:"leaf_data,omitempty"`
	// neighbors is a list of all the adjacent nodes along the path
	// from the deepest node to the head.
	Neighbors [][]byte `protobuf:"bytes,2,rep,name=neighbors,proto3" json:"neighbors,omitempty"`
}

func (m *GetLeafResponse) Reset()                    { *m = GetLeafResponse{} }
func (m *GetLeafResponse) String() string            { return proto.CompactTextString(m) }
func (*GetLeafResponse) ProtoMessage()               {}
func (*GetLeafResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GetLeafResponse) GetLeafData() []byte {
	if m != nil {
		return m.LeafData
	}
	return nil
}

func (m *GetLeafResponse) GetNeighbors() [][]byte {
	if m != nil {
		return m.Neighbors
	}
	return nil
}

// UpdateLeafRequest submits a change for the value at index.
type UpdateLeafRequest struct {
	// index includes the updated leaf index.
	Index []byte `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	// mutation contains the actual update.
	Mutation []byte `protobuf:"bytes,2,opt,name=mutation,proto3" json:"mutation,omitempty"`
}

func (m *UpdateLeafRequest) Reset()                    { *m = UpdateLeafRequest{} }
func (m *UpdateLeafRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateLeafRequest) ProtoMessage()               {}
func (*UpdateLeafRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UpdateLeafRequest) GetIndex() []byte {
	if m != nil {
		return m.Index
	}
	return nil
}

func (m *UpdateLeafRequest) GetMutation() []byte {
	if m != nil {
		return m.Mutation
	}
	return nil
}

// UpdateLeafResponse returns the current value of index.
type UpdateLeafResponse struct {
	// proof contains a proof that the update is accepted and processed.
	Proof *GetLeafResponse `protobuf:"bytes,1,opt,name=proof" json:"proof,omitempty"`
}

func (m *UpdateLeafResponse) Reset()                    { *m = UpdateLeafResponse{} }
func (m *UpdateLeafResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateLeafResponse) ProtoMessage()               {}
func (*UpdateLeafResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *UpdateLeafResponse) GetProof() *GetLeafResponse {
	if m != nil {
		return m.Proof
	}
	return nil
}

// MutationEntry is either a mutation or an epoch advancement.
type MutationEntry struct {
	// type specifies the entry mutation type.
	//
	// Types that are valid to be assigned to Type:
	//	*MutationEntry_Update
	//	*MutationEntry_AdvanceEpoch
	Type isMutationEntry_Type `protobuf_oneof:"type"`
}

func (m *MutationEntry) Reset()                    { *m = MutationEntry{} }
func (m *MutationEntry) String() string            { return proto.CompactTextString(m) }
func (*MutationEntry) ProtoMessage()               {}
func (*MutationEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type isMutationEntry_Type interface {
	isMutationEntry_Type()
}

type MutationEntry_Update struct {
	Update *UpdateLeafRequest `protobuf:"bytes,1,opt,name=update,oneof"`
}
type MutationEntry_AdvanceEpoch struct {
	AdvanceEpoch bool `protobuf:"varint,2,opt,name=advance_epoch,json=advanceEpoch,oneof"`
}

func (*MutationEntry_Update) isMutationEntry_Type()       {}
func (*MutationEntry_AdvanceEpoch) isMutationEntry_Type() {}

func (m *MutationEntry) GetType() isMutationEntry_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *MutationEntry) GetUpdate() *UpdateLeafRequest {
	if x, ok := m.GetType().(*MutationEntry_Update); ok {
		return x.Update
	}
	return nil
}

func (m *MutationEntry) GetAdvanceEpoch() bool {
	if x, ok := m.GetType().(*MutationEntry_AdvanceEpoch); ok {
		return x.AdvanceEpoch
	}
	return false
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*MutationEntry) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _MutationEntry_OneofMarshaler, _MutationEntry_OneofUnmarshaler, _MutationEntry_OneofSizer, []interface{}{
		(*MutationEntry_Update)(nil),
		(*MutationEntry_AdvanceEpoch)(nil),
	}
}

func _MutationEntry_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*MutationEntry)
	// type
	switch x := m.Type.(type) {
	case *MutationEntry_Update:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Update); err != nil {
			return err
		}
	case *MutationEntry_AdvanceEpoch:
		t := uint64(0)
		if x.AdvanceEpoch {
			t = 1
		}
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case nil:
	default:
		return fmt.Errorf("MutationEntry.Type has unexpected type %T", x)
	}
	return nil
}

func _MutationEntry_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*MutationEntry)
	switch tag {
	case 1: // type.update
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(UpdateLeafRequest)
		err := b.DecodeMessage(msg)
		m.Type = &MutationEntry_Update{msg}
		return true, err
	case 2: // type.advance_epoch
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Type = &MutationEntry_AdvanceEpoch{x != 0}
		return true, err
	default:
		return false, nil
	}
}

func _MutationEntry_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*MutationEntry)
	// type
	switch x := m.Type.(type) {
	case *MutationEntry_Update:
		s := proto.Size(x.Update)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *MutationEntry_AdvanceEpoch:
		n += proto.SizeVarint(2<<3 | proto.WireVarint)
		n += 1
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*MapHead)(nil), "ctmap.MapHead")
	proto.RegisterType((*SignedMapHead)(nil), "ctmap.SignedMapHead")
	proto.RegisterType((*DigitallySigned)(nil), "ctmap.DigitallySigned")
	proto.RegisterType((*GetLeafRequest)(nil), "ctmap.GetLeafRequest")
	proto.RegisterType((*GetLeafResponse)(nil), "ctmap.GetLeafResponse")
	proto.RegisterType((*UpdateLeafRequest)(nil), "ctmap.UpdateLeafRequest")
	proto.RegisterType((*UpdateLeafResponse)(nil), "ctmap.UpdateLeafResponse")
	proto.RegisterType((*MutationEntry)(nil), "ctmap.MutationEntry")
	proto.RegisterEnum("ctmap.DigitallySigned_HashAlgorithm", DigitallySigned_HashAlgorithm_name, DigitallySigned_HashAlgorithm_value)
	proto.RegisterEnum("ctmap.DigitallySigned_SignatureAlgorithm", DigitallySigned_SignatureAlgorithm_name, DigitallySigned_SignatureAlgorithm_value)
}

func init() { proto.RegisterFile("ctmap.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 589 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x53, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0x6d, 0xfa, 0xb5, 0xf6, 0xb6, 0xe9, 0x8a, 0x85, 0x50, 0x55, 0x90, 0xa8, 0x22, 0x90, 0x3a,
	0x69, 0xca, 0xb4, 0xa0, 0x21, 0x40, 0xbc, 0x74, 0xb4, 0xa2, 0x12, 0x5b, 0x27, 0xb9, 0xec, 0x81,
	0xa7, 0xca, 0x5b, 0xdc, 0xc4, 0x22, 0x89, 0x43, 0xec, 0x4c, 0xf4, 0x9d, 0xff, 0xc3, 0xcf, 0xe1,
	0xef, 0xa0, 0xd8, 0xee, 0xe7, 0x84, 0x78, 0xf3, 0xb9, 0x3e, 0xf7, 0xdc, 0xe3, 0x9b, 0x13, 0x68,
	0xdd, 0xcb, 0x98, 0xa4, 0x6e, 0x9a, 0x71, 0xc9, 0x51, 0x4d, 0x81, 0xfe, 0xcb, 0x80, 0xf3, 0x20,
	0xa2, 0x67, 0xaa, 0x78, 0x97, 0x2f, 0xcf, 0x24, 0x8b, 0xa9, 0x90, 0x24, 0x36, 0x3c, 0xe7, 0x97,
	0x05, 0x47, 0xd7, 0x24, 0x9d, 0x52, 0xe2, 0xa3, 0xa7, 0x50, 0xcb, 0x28, 0x89, 0xe2, 0x9e, 0x35,
	0xb0, 0x86, 0x4d, 0xac, 0x41, 0x51, 0xa5, 0x29, 0xbf, 0x0f, 0x7b, 0xe5, 0x81, 0x35, 0xac, 0x60,
	0x0d, 0x10, 0x82, 0x6a, 0xc6, 0xb9, 0xec, 0x55, 0x06, 0xd6, 0xb0, 0x8d, 0xd5, 0x19, 0xbd, 0x07,
	0x60, 0x42, 0xe4, 0x74, 0x51, 0x0c, 0xe9, 0x55, 0x07, 0xd6, 0xb0, 0xe5, 0xf5, 0x5d, 0xed, 0xc0,
	0x5d, 0x3b, 0x70, 0xbf, 0xae, 0x1d, 0xe0, 0xa6, 0x62, 0x17, 0xd8, 0xf9, 0x63, 0x81, 0x3d, 0x67,
	0x41, 0x42, 0xfd, 0xb5, 0x99, 0x13, 0x68, 0xc4, 0x24, 0x5d, 0x84, 0x94, 0xf8, 0xca, 0x4f, 0xcb,
	0xeb, 0xb8, 0xfa, 0x81, 0x86, 0x81, 0x8f, 0x62, 0x43, 0x1d, 0x03, 0x08, 0x16, 0x24, 0x44, 0xe6,
	0x19, 0x15, 0xbd, 0xf2, 0xa0, 0x32, 0x6c, 0x79, 0xaf, 0x0c, 0x79, 0x4f, 0x54, 0x21, 0x4d, 0x9b,
	0x24, 0x32, 0x5b, 0xe1, 0x9d, 0xbe, 0xfe, 0x2d, 0x1c, 0x1f, 0x5c, 0xa3, 0x2e, 0x54, 0xbe, 0xd3,
	0x95, 0x59, 0x47, 0x71, 0x44, 0xa7, 0x50, 0x7b, 0x20, 0x51, 0x4e, 0xd5, 0x32, 0x5a, 0xde, 0x33,
	0x33, 0x65, 0xcc, 0x02, 0x26, 0x49, 0x14, 0xad, 0xf4, 0x38, 0xac, 0x49, 0x1f, 0xca, 0xef, 0x2c,
	0xe7, 0x77, 0x19, 0x8e, 0x0f, 0xae, 0xd1, 0x17, 0xe8, 0x84, 0x44, 0x84, 0x0b, 0x12, 0x05, 0x3c,
	0x63, 0x32, 0xd4, 0x1b, 0xef, 0x6c, 0x4c, 0x1f, 0xf0, 0xdd, 0x29, 0x11, 0xe1, 0x68, 0xcd, 0xc5,
	0x76, 0xb8, 0x0b, 0xd1, 0x0c, 0x6c, 0xc1, 0x82, 0x1d, 0xad, 0xb2, 0xd2, 0x3a, 0xf9, 0x87, 0xd6,
	0xe6, 0x8d, 0x5b, 0xc1, 0xb6, 0x60, 0xc1, 0x56, 0xef, 0x05, 0x34, 0x37, 0x5b, 0x31, 0x9f, 0x77,
	0x5b, 0x70, 0xce, 0xc1, 0xde, 0x73, 0x83, 0x1a, 0x50, 0x9d, 0xdd, 0xcc, 0x26, 0xdd, 0x12, 0x02,
	0xa8, 0xcf, 0xa7, 0x23, 0xef, 0xe2, 0x6d, 0xb7, 0x6a, 0xce, 0x17, 0xe7, 0x5e, 0xb7, 0xee, 0xb8,
	0x80, 0x1e, 0x0f, 0x45, 0x36, 0x34, 0x47, 0xb3, 0x9b, 0xd9, 0xb7, 0xeb, 0x9b, 0xdb, 0x79, 0xb7,
	0x84, 0x9a, 0x50, 0x9b, 0x7c, 0x1a, 0xcf, 0x47, 0x5d, 0xcb, 0xf9, 0x08, 0x9d, 0xcf, 0x54, 0x5e,
	0x51, 0xb2, 0xc4, 0xf4, 0x47, 0x4e, 0x85, 0x2c, 0x22, 0xc8, 0x12, 0x9f, 0xfe, 0x54, 0x6b, 0x6a,
	0x63, 0x0d, 0xf6, 0x83, 0x59, 0x35, 0xc1, 0x74, 0xae, 0xe0, 0x78, 0xd3, 0x2d, 0x52, 0x9e, 0x08,
	0x8a, 0x9e, 0x43, 0x33, 0xa2, 0x64, 0xb9, 0xf0, 0x89, 0x24, 0x46, 0xa2, 0x51, 0x14, 0xc6, 0x44,
	0x92, 0xe2, 0xb9, 0x09, 0x65, 0x41, 0x78, 0xc7, 0x33, 0x9d, 0x9d, 0x36, 0xde, 0x16, 0x9c, 0x09,
	0x3c, 0xb9, 0x4d, 0x7d, 0x22, 0xe9, 0xff, 0xed, 0xf4, 0xa1, 0x11, 0xe7, 0x92, 0x48, 0xc6, 0x13,
	0xe5, 0xa8, 0x8d, 0x37, 0xd8, 0xb9, 0x04, 0xb4, 0x2b, 0x63, 0x7c, 0x9d, 0x42, 0x2d, 0xcd, 0x38,
	0x5f, 0x9a, 0x7c, 0xaf, 0xc3, 0x74, 0x60, 0x1f, 0x6b, 0x92, 0x93, 0x81, 0x7d, 0x6d, 0xf4, 0x74,
	0x3a, 0x3d, 0xa8, 0xe7, 0x4a, 0xd4, 0xf4, 0xf7, 0x4c, 0xff, 0x23, 0xc3, 0xd3, 0x12, 0x36, 0x4c,
	0xf4, 0x1a, 0x6c, 0xe2, 0x3f, 0x90, 0xe4, 0x9e, 0x2e, 0xb6, 0xbb, 0x6b, 0x4c, 0x4b, 0xb8, 0x6d,
	0xca, 0x93, 0xa2, 0x7a, 0x59, 0x87, 0xaa, 0x5c, 0xa5, 0xf4, 0xae, 0xae, 0xfe, 0xda, 0x37, 0x7f,
	0x03, 0x00, 0x00, 0xff, 0xff, 0xe6, 0xd4, 0x16, 0xa3, 0x5b, 0x04, 0x00, 0x00,
}
