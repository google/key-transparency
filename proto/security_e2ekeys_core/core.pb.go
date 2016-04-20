// Code generated by protoc-gen-go.
// source: proto/security_e2ekeys_core/core.proto
// DO NOT EDIT!

/*
Package security_e2ekeys_core is a generated protocol buffer package.

It is generated from these files:
	proto/security_e2ekeys_core/core.proto

It has these top-level messages:
	EntryStorage
	EpochInfo
	DomainInfo
	VerifierInfo
*/
package security_e2ekeys_core

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import security_ctmap "github.com/gdbelvin/e2e-key-server/proto/security_ctmap"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

// EntryStorage is what gets serialized to disk / log.
type EntryStorage struct {
	// commitment_timestamp is a sequential update number of the SignedEntryUpdate.
	CommitmentTimestamp int64 `protobuf:"varint,1,opt,name=commitment_timestamp,json=commitmentTimestamp" json:"commitment_timestamp,omitempty"`
	// entry_update contains a SignedEntryUpdate proto with the actual update
	// contents.
	SignedEntryUpdate *security_ctmap.SignedEntryUpdate `protobuf:"bytes,2,opt,name=signed_entry_update,json=signedEntryUpdate" json:"signed_entry_update,omitempty"`
	// profile is the serialized protobuf Profile.
	// profile is private and must not be released to verifiers.
	Profile []byte `protobuf:"bytes,3,opt,name=profile,proto3" json:"profile,omitempty"`
	// commitment_key is at least 16 random bytes.
	CommitmentKey []byte `protobuf:"bytes,4,opt,name=commitment_key,json=commitmentKey,proto3" json:"commitment_key,omitempty"`
	// domain allows servers to keep track of multiple trees at once.
	Domain string `protobuf:"bytes,5,opt,name=domain" json:"domain,omitempty"`
}

func (m *EntryStorage) Reset()                    { *m = EntryStorage{} }
func (m *EntryStorage) String() string            { return proto.CompactTextString(m) }
func (*EntryStorage) ProtoMessage()               {}
func (*EntryStorage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *EntryStorage) GetSignedEntryUpdate() *security_ctmap.SignedEntryUpdate {
	if m != nil {
		return m.SignedEntryUpdate
	}
	return nil
}

// EpochInfo is what gets serialized to disk / log.
type EpochInfo struct {
	// signed_epoch_head is the signed epoch head of the created epoch.
	SignedEpochHead *security_ctmap.SignedEpochHead `protobuf:"bytes,1,opt,name=signed_epoch_head,json=signedEpochHead" json:"signed_epoch_head,omitempty"`
	// last_commitment_timestamp is the timestamp of the last update included in
	// created epoch.
	LastCommitmentTimestamp int64 `protobuf:"varint,2,opt,name=last_commitment_timestamp,json=lastCommitmentTimestamp" json:"last_commitment_timestamp,omitempty"`
}

func (m *EpochInfo) Reset()                    { *m = EpochInfo{} }
func (m *EpochInfo) String() string            { return proto.CompactTextString(m) }
func (*EpochInfo) ProtoMessage()               {}
func (*EpochInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *EpochInfo) GetSignedEpochHead() *security_ctmap.SignedEpochHead {
	if m != nil {
		return m.SignedEpochHead
	}
	return nil
}

// DomainInfo is the information that need to be baked into an application
// in order to verify information for a domain.
type DomainInfo struct {
	PublicKeys []*security_ctmap.PublicKey `protobuf:"bytes,1,rep,name=public_keys,json=publicKeys" json:"public_keys,omitempty"`
	// required_sigs is the number of valid signatures to require out of
	// public_keys before considering signed_tree heads legitimate.
	RequiredSigs uint32 `protobuf:"varint,2,opt,name=required_sigs,json=requiredSigs" json:"required_sigs,omitempty"`
	// domain is the doman suffix to use when resolving a user_id to a domain.
	Domain string `protobuf:"bytes,3,opt,name=domain" json:"domain,omitempty"`
	// api_url is the url prefix to use when querying users on this domain.
	ApiUrl string `protobuf:"bytes,4,opt,name=api_url,json=apiUrl" json:"api_url,omitempty"`
}

func (m *DomainInfo) Reset()                    { *m = DomainInfo{} }
func (m *DomainInfo) String() string            { return proto.CompactTextString(m) }
func (*DomainInfo) ProtoMessage()               {}
func (*DomainInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DomainInfo) GetPublicKeys() []*security_ctmap.PublicKey {
	if m != nil {
		return m.PublicKeys
	}
	return nil
}

type VerifierInfo struct {
	PublicKeys []*security_ctmap.PublicKey `protobuf:"bytes,1,rep,name=public_keys,json=publicKeys" json:"public_keys,omitempty"`
	// required_sigs is the number of valid signatures to require out of
	// public_keys before considering signed_tree heads legitimate.
	RequiredSigs uint32 `protobuf:"varint,2,opt,name=required_sigs,json=requiredSigs" json:"required_sigs,omitempty"`
	// api_url is the url prefix to use when querying users on this domain.
	ApiUrl string `protobuf:"bytes,4,opt,name=api_url,json=apiUrl" json:"api_url,omitempty"`
	// domain is the doman suffix that this verifier is responsible for.
	Domain string `protobuf:"bytes,3,opt,name=domain" json:"domain,omitempty"`
}

func (m *VerifierInfo) Reset()                    { *m = VerifierInfo{} }
func (m *VerifierInfo) String() string            { return proto.CompactTextString(m) }
func (*VerifierInfo) ProtoMessage()               {}
func (*VerifierInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *VerifierInfo) GetPublicKeys() []*security_ctmap.PublicKey {
	if m != nil {
		return m.PublicKeys
	}
	return nil
}

func init() {
	proto.RegisterType((*EntryStorage)(nil), "security.e2ekeys.core.EntryStorage")
	proto.RegisterType((*EpochInfo)(nil), "security.e2ekeys.core.EpochInfo")
	proto.RegisterType((*DomainInfo)(nil), "security.e2ekeys.core.DomainInfo")
	proto.RegisterType((*VerifierInfo)(nil), "security.e2ekeys.core.VerifierInfo")
}

var fileDescriptor0 = []byte{
	// 424 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x92, 0xcf, 0x6e, 0xd4, 0x30,
	0x10, 0xc6, 0x95, 0x06, 0x5a, 0x75, 0x76, 0x17, 0x54, 0x17, 0x68, 0xca, 0x05, 0x08, 0x02, 0x71,
	0x69, 0x22, 0x96, 0x5b, 0x2f, 0x1c, 0xa0, 0x12, 0xa8, 0x17, 0x48, 0x59, 0xae, 0x96, 0x37, 0x99,
	0xcd, 0x5a, 0x24, 0xb1, 0xb1, 0x1d, 0xa4, 0xbc, 0x07, 0x67, 0xc4, 0x0b, 0xf2, 0x0e, 0x4c, 0xcc,
	0x46, 0xfb, 0x07, 0xf6, 0xca, 0x21, 0x51, 0x66, 0xe6, 0xcb, 0xf8, 0xf7, 0x7d, 0x32, 0x3c, 0xd7,
	0x46, 0x39, 0x95, 0x5a, 0xcc, 0x5b, 0x23, 0x5d, 0xc7, 0x71, 0x8a, 0x5f, 0xb0, 0xb3, 0x3c, 0x57,
	0x06, 0xd3, 0xfe, 0x95, 0x78, 0x01, 0xbb, 0x3f, 0x28, 0x92, 0x95, 0x22, 0xe9, 0x87, 0x0f, 0x5f,
	0x97, 0xd2, 0x2d, 0xdb, 0x39, 0x15, 0x75, 0x5a, 0x2a, 0x55, 0x56, 0x98, 0xd2, 0xfc, 0x82, 0x04,
	0x17, 0x16, 0xcd, 0x37, 0x34, 0xe9, 0xce, 0xfe, 0xdc, 0xd5, 0x42, 0xa7, 0xf4, 0xfc, 0xd9, 0x1b,
	0xff, 0x0a, 0x60, 0x7c, 0xd5, 0x38, 0xd3, 0xdd, 0x38, 0x65, 0x44, 0x89, 0xec, 0x25, 0xdc, 0xa3,
	0x65, 0xb5, 0x74, 0x35, 0x36, 0x8e, 0x3b, 0x59, 0xa3, 0x75, 0xa2, 0xd6, 0x51, 0xf0, 0x38, 0x78,
	0x11, 0x66, 0xa7, 0xeb, 0xd9, 0xa7, 0x61, 0xc4, 0x3e, 0xc2, 0xa9, 0x95, 0x65, 0x83, 0x05, 0xc7,
	0x7e, 0x13, 0x6f, 0x75, 0x21, 0x1c, 0x46, 0x07, 0xf4, 0xc7, 0x68, 0xfa, 0x24, 0xd9, 0x3e, 0x3b,
	0xb9, 0xf1, 0x52, 0x7f, 0xe6, 0xcc, 0x0b, 0xb3, 0x13, 0xbb, 0xdb, 0x62, 0x11, 0x1c, 0x11, 0xdf,
	0x42, 0x56, 0x18, 0x85, 0xb4, 0x66, 0x9c, 0x0d, 0x25, 0x7b, 0x06, 0x77, 0x36, 0xf8, 0xc8, 0x68,
	0x74, 0xcb, 0x0b, 0x26, 0xeb, 0xee, 0x35, 0x76, 0xec, 0x01, 0x1c, 0x16, 0xaa, 0x16, 0xb2, 0x89,
	0x6e, 0xd3, 0xf8, 0x38, 0x5b, 0x55, 0xf1, 0xf7, 0x00, 0x8e, 0xaf, 0xb4, 0xca, 0x97, 0xef, 0x9b,
	0x85, 0x62, 0xd7, 0x70, 0x32, 0x90, 0xf7, 0x3d, 0xbe, 0x44, 0x51, 0x78, 0xa7, 0xa3, 0xe9, 0xa3,
	0x3d, 0xdc, 0xbd, 0xee, 0x1d, 0xc9, 0xb2, 0xbb, 0x76, 0xbb, 0xc1, 0x2e, 0xe1, 0xbc, 0x12, 0xd6,
	0xf1, 0x7f, 0xc6, 0x77, 0xe0, 0xe3, 0x3b, 0xeb, 0x05, 0x6f, 0xfe, 0x8e, 0x30, 0xfe, 0x11, 0x00,
	0xbc, 0xf5, 0x84, 0x9e, 0xeb, 0x12, 0x46, 0xba, 0x9d, 0x57, 0x32, 0xef, 0x0d, 0x5a, 0x22, 0x0a,
	0x89, 0xe8, 0x7c, 0x97, 0xe8, 0x83, 0x97, 0x90, 0xdb, 0x0c, 0xf4, 0xf0, 0x69, 0xd9, 0x53, 0x98,
	0x18, 0xfc, 0xda, 0x4a, 0x43, 0xae, 0x08, 0xd1, 0xfa, 0xa3, 0x27, 0xd9, 0x78, 0x68, 0x92, 0x0f,
	0xbb, 0x11, 0x4f, 0xb8, 0x19, 0x0f, 0x3b, 0x83, 0x23, 0xa1, 0x25, 0x6f, 0x4d, 0xe5, 0x63, 0xa5,
	0x01, 0x95, 0x33, 0x53, 0xc5, 0x3f, 0xe9, 0x9e, 0x7c, 0x46, 0x23, 0x17, 0x12, 0xcd, 0xff, 0x41,
	0xdc, 0x87, 0xb2, 0x8f, 0x7d, 0x7e, 0xe8, 0x6f, 0xf4, 0xab, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x0a, 0xf6, 0x4d, 0x35, 0x53, 0x03, 0x00, 0x00,
}
