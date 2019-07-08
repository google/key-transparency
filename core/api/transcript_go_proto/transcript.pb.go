// Code generated by protoc-gen-go. DO NOT EDIT.
// source: transcript/transcript.proto

package transcript_go_proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	keytransparency_go_proto "github.com/google/keytransparency/core/api/v1/keytransparency_go_proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Transcript struct {
	Description          string                              `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	Directory            *keytransparency_go_proto.Directory `protobuf:"bytes,2,opt,name=directory,proto3" json:"directory,omitempty"`
	Rpcs                 []*Unary                            `protobuf:"bytes,3,rep,name=rpcs,proto3" json:"rpcs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *Transcript) Reset()         { *m = Transcript{} }
func (m *Transcript) String() string { return proto.CompactTextString(m) }
func (*Transcript) ProtoMessage()    {}
func (*Transcript) Descriptor() ([]byte, []int) {
	return fileDescriptor_5bd339d1227847d1, []int{0}
}

func (m *Transcript) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transcript.Unmarshal(m, b)
}
func (m *Transcript) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transcript.Marshal(b, m, deterministic)
}
func (m *Transcript) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transcript.Merge(m, src)
}
func (m *Transcript) XXX_Size() int {
	return xxx_messageInfo_Transcript.Size(m)
}
func (m *Transcript) XXX_DiscardUnknown() {
	xxx_messageInfo_Transcript.DiscardUnknown(m)
}

var xxx_messageInfo_Transcript proto.InternalMessageInfo

func (m *Transcript) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Transcript) GetDirectory() *keytransparency_go_proto.Directory {
	if m != nil {
		return m.Directory
	}
	return nil
}

func (m *Transcript) GetRpcs() []*Unary {
	if m != nil {
		return m.Rpcs
	}
	return nil
}

type Unary struct {
	Desc string `protobuf:"bytes,1,opt,name=desc,proto3" json:"desc,omitempty"`
	// Types that are valid to be assigned to ReqRespPair:
	//	*Unary_GetUser
	//	*Unary_BatchGetUser
	//	*Unary_BatchListUserRevisions
	ReqRespPair          isUnary_ReqRespPair `protobuf_oneof:"req_resp_pair"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Unary) Reset()         { *m = Unary{} }
func (m *Unary) String() string { return proto.CompactTextString(m) }
func (*Unary) ProtoMessage()    {}
func (*Unary) Descriptor() ([]byte, []int) {
	return fileDescriptor_5bd339d1227847d1, []int{1}
}

func (m *Unary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Unary.Unmarshal(m, b)
}
func (m *Unary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Unary.Marshal(b, m, deterministic)
}
func (m *Unary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Unary.Merge(m, src)
}
func (m *Unary) XXX_Size() int {
	return xxx_messageInfo_Unary.Size(m)
}
func (m *Unary) XXX_DiscardUnknown() {
	xxx_messageInfo_Unary.DiscardUnknown(m)
}

var xxx_messageInfo_Unary proto.InternalMessageInfo

func (m *Unary) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

type isUnary_ReqRespPair interface {
	isUnary_ReqRespPair()
}

type Unary_GetUser struct {
	GetUser *GetUser `protobuf:"bytes,4,opt,name=get_user,json=getUser,proto3,oneof"`
}

type Unary_BatchGetUser struct {
	BatchGetUser *BatchGetUser `protobuf:"bytes,5,opt,name=batch_get_user,json=batchGetUser,proto3,oneof"`
}

type Unary_BatchListUserRevisions struct {
	BatchListUserRevisions *BatchListUserRevisions `protobuf:"bytes,6,opt,name=batch_list_user_revisions,json=batchListUserRevisions,proto3,oneof"`
}

func (*Unary_GetUser) isUnary_ReqRespPair() {}

func (*Unary_BatchGetUser) isUnary_ReqRespPair() {}

func (*Unary_BatchListUserRevisions) isUnary_ReqRespPair() {}

func (m *Unary) GetReqRespPair() isUnary_ReqRespPair {
	if m != nil {
		return m.ReqRespPair
	}
	return nil
}

func (m *Unary) GetGetUser() *GetUser {
	if x, ok := m.GetReqRespPair().(*Unary_GetUser); ok {
		return x.GetUser
	}
	return nil
}

func (m *Unary) GetBatchGetUser() *BatchGetUser {
	if x, ok := m.GetReqRespPair().(*Unary_BatchGetUser); ok {
		return x.BatchGetUser
	}
	return nil
}

func (m *Unary) GetBatchListUserRevisions() *BatchListUserRevisions {
	if x, ok := m.GetReqRespPair().(*Unary_BatchListUserRevisions); ok {
		return x.BatchListUserRevisions
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Unary) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Unary_GetUser)(nil),
		(*Unary_BatchGetUser)(nil),
		(*Unary_BatchListUserRevisions)(nil),
	}
}

type GetUser struct {
	Request              *keytransparency_go_proto.GetUserRequest  `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	Response             *keytransparency_go_proto.GetUserResponse `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                  `json:"-"`
	XXX_unrecognized     []byte                                    `json:"-"`
	XXX_sizecache        int32                                     `json:"-"`
}

func (m *GetUser) Reset()         { *m = GetUser{} }
func (m *GetUser) String() string { return proto.CompactTextString(m) }
func (*GetUser) ProtoMessage()    {}
func (*GetUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_5bd339d1227847d1, []int{2}
}

func (m *GetUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUser.Unmarshal(m, b)
}
func (m *GetUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUser.Marshal(b, m, deterministic)
}
func (m *GetUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUser.Merge(m, src)
}
func (m *GetUser) XXX_Size() int {
	return xxx_messageInfo_GetUser.Size(m)
}
func (m *GetUser) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUser.DiscardUnknown(m)
}

var xxx_messageInfo_GetUser proto.InternalMessageInfo

func (m *GetUser) GetRequest() *keytransparency_go_proto.GetUserRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *GetUser) GetResponse() *keytransparency_go_proto.GetUserResponse {
	if m != nil {
		return m.Response
	}
	return nil
}

type BatchGetUser struct {
	Request              *keytransparency_go_proto.BatchGetUserRequest  `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	Response             *keytransparency_go_proto.BatchGetUserResponse `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                       `json:"-"`
	XXX_unrecognized     []byte                                         `json:"-"`
	XXX_sizecache        int32                                          `json:"-"`
}

func (m *BatchGetUser) Reset()         { *m = BatchGetUser{} }
func (m *BatchGetUser) String() string { return proto.CompactTextString(m) }
func (*BatchGetUser) ProtoMessage()    {}
func (*BatchGetUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_5bd339d1227847d1, []int{3}
}

func (m *BatchGetUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchGetUser.Unmarshal(m, b)
}
func (m *BatchGetUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchGetUser.Marshal(b, m, deterministic)
}
func (m *BatchGetUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchGetUser.Merge(m, src)
}
func (m *BatchGetUser) XXX_Size() int {
	return xxx_messageInfo_BatchGetUser.Size(m)
}
func (m *BatchGetUser) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchGetUser.DiscardUnknown(m)
}

var xxx_messageInfo_BatchGetUser proto.InternalMessageInfo

func (m *BatchGetUser) GetRequest() *keytransparency_go_proto.BatchGetUserRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *BatchGetUser) GetResponse() *keytransparency_go_proto.BatchGetUserResponse {
	if m != nil {
		return m.Response
	}
	return nil
}

type BatchListUserRevisions struct {
	Request              *keytransparency_go_proto.BatchListUserRevisionsRequest  `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	Response             *keytransparency_go_proto.BatchListUserRevisionsResponse `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                 `json:"-"`
	XXX_unrecognized     []byte                                                   `json:"-"`
	XXX_sizecache        int32                                                    `json:"-"`
}

func (m *BatchListUserRevisions) Reset()         { *m = BatchListUserRevisions{} }
func (m *BatchListUserRevisions) String() string { return proto.CompactTextString(m) }
func (*BatchListUserRevisions) ProtoMessage()    {}
func (*BatchListUserRevisions) Descriptor() ([]byte, []int) {
	return fileDescriptor_5bd339d1227847d1, []int{4}
}

func (m *BatchListUserRevisions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchListUserRevisions.Unmarshal(m, b)
}
func (m *BatchListUserRevisions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchListUserRevisions.Marshal(b, m, deterministic)
}
func (m *BatchListUserRevisions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchListUserRevisions.Merge(m, src)
}
func (m *BatchListUserRevisions) XXX_Size() int {
	return xxx_messageInfo_BatchListUserRevisions.Size(m)
}
func (m *BatchListUserRevisions) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchListUserRevisions.DiscardUnknown(m)
}

var xxx_messageInfo_BatchListUserRevisions proto.InternalMessageInfo

func (m *BatchListUserRevisions) GetRequest() *keytransparency_go_proto.BatchListUserRevisionsRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *BatchListUserRevisions) GetResponse() *keytransparency_go_proto.BatchListUserRevisionsResponse {
	if m != nil {
		return m.Response
	}
	return nil
}

func init() {
	proto.RegisterType((*Transcript)(nil), "google.keytransparency.transcript.Transcript")
	proto.RegisterType((*Unary)(nil), "google.keytransparency.transcript.Unary")
	proto.RegisterType((*GetUser)(nil), "google.keytransparency.transcript.GetUser")
	proto.RegisterType((*BatchGetUser)(nil), "google.keytransparency.transcript.BatchGetUser")
	proto.RegisterType((*BatchListUserRevisions)(nil), "google.keytransparency.transcript.BatchListUserRevisions")
}

func init() { proto.RegisterFile("transcript/transcript.proto", fileDescriptor_5bd339d1227847d1) }

var fileDescriptor_5bd339d1227847d1 = []byte{
	// 454 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x97, 0xad, 0x5b, 0xb7, 0xd7, 0x31, 0x24, 0x1f, 0xa6, 0x30, 0x2e, 0xa5, 0xe2, 0x50,
	0x76, 0x70, 0xd4, 0x72, 0x61, 0x12, 0x42, 0xa8, 0x20, 0x56, 0x09, 0x4e, 0x16, 0x15, 0x12, 0x97,
	0x28, 0x49, 0x9f, 0x32, 0x8b, 0x2d, 0xce, 0x9e, 0xdd, 0x48, 0xfd, 0x22, 0x7c, 0x03, 0xee, 0x7c,
	0x07, 0x24, 0x3e, 0x17, 0xaa, 0xdd, 0x36, 0x51, 0x32, 0xb6, 0x68, 0xa7, 0x3e, 0xd7, 0xfe, 0xff,
	0xfc, 0xb3, 0x1d, 0x3d, 0x78, 0x6e, 0x28, 0xca, 0x74, 0x42, 0x32, 0x37, 0x41, 0x59, 0xf2, 0x9c,
	0x94, 0x51, 0xec, 0x45, 0xaa, 0x54, 0x7a, 0x8d, 0xfc, 0x07, 0x2e, 0xed, 0x5c, 0x1e, 0x11, 0x66,
	0xc9, 0x92, 0x97, 0x0b, 0xcf, 0xfc, 0x62, 0x14, 0xd4, 0xa7, 0x6d, 0xf8, 0xec, 0xa4, 0x18, 0x05,
	0xd1, 0xfc, 0x46, 0x66, 0x6e, 0x3c, 0xf8, 0xed, 0x01, 0x7c, 0xdd, 0x06, 0x59, 0x1f, 0x7a, 0x73,
	0x74, 0xb5, 0x54, 0x99, 0xef, 0xf5, 0xbd, 0xe1, 0x91, 0xa8, 0xfe, 0xc5, 0x26, 0x70, 0x34, 0x97,
	0x84, 0x89, 0x51, 0xb4, 0xf4, 0x77, 0xfb, 0xde, 0xb0, 0x37, 0x7e, 0xc9, 0xff, 0x63, 0x54, 0x8c,
	0xf8, 0xc7, 0xcd, 0x5a, 0x51, 0xc6, 0xd8, 0x5b, 0xe8, 0x50, 0x9e, 0x68, 0x7f, 0xaf, 0xbf, 0x37,
	0xec, 0x8d, 0x87, 0xfc, 0xc1, 0x03, 0xf1, 0x59, 0x16, 0xd1, 0x52, 0xd8, 0xd4, 0xe0, 0xef, 0x2e,
	0xec, 0xdb, 0x31, 0x63, 0xd0, 0x59, 0xa9, 0xad, 0x35, 0x6d, 0xcd, 0x2e, 0xe1, 0x30, 0x45, 0x13,
	0x2e, 0x34, 0x92, 0xdf, 0xb1, 0x7a, 0xe7, 0x2d, 0xf8, 0x97, 0x68, 0x66, 0x1a, 0x69, 0xba, 0x23,
	0xba, 0xa9, 0x2b, 0xd9, 0x37, 0x38, 0x89, 0x23, 0x93, 0x5c, 0x85, 0x5b, 0xdc, 0xbe, 0xc5, 0x05,
	0x2d, 0x70, 0x93, 0x55, 0xb0, 0x64, 0x1e, 0xc7, 0x95, 0x31, 0x2b, 0xe0, 0x99, 0x03, 0x5f, 0x4b,
	0xed, 0xc8, 0x21, 0x61, 0x21, 0xb5, 0x54, 0x99, 0xf6, 0x0f, 0xec, 0x1e, 0x17, 0x6d, 0xf7, 0xf8,
	0x22, 0xb5, 0x85, 0x8a, 0x0d, 0x60, 0xba, 0x23, 0x4e, 0xe3, 0x3b, 0x67, 0x26, 0x4f, 0xe1, 0x09,
	0xe1, 0x6d, 0x48, 0xa8, 0xf3, 0x30, 0x8f, 0x24, 0x0d, 0x7e, 0x7a, 0xd0, 0xdd, 0x48, 0x7d, 0x80,
	0x2e, 0xe1, 0xed, 0x02, 0xb5, 0xb1, 0xb7, 0xd9, 0x1b, 0xbf, 0xba, 0xe7, 0x51, 0xd7, 0x21, 0xe1,
	0x02, 0x62, 0x93, 0x64, 0x9f, 0xe0, 0x70, 0x45, 0x57, 0x99, 0xc6, 0xf5, 0xa7, 0x71, 0xde, 0x86,
	0xe2, 0x12, 0x62, 0x9b, 0x1d, 0xfc, 0xf2, 0xe0, 0xb8, 0x7a, 0x85, 0x6c, 0x5a, 0xb7, 0xe3, 0xf7,
	0x70, 0xab, 0xc9, 0x86, 0xe2, 0xe7, 0x86, 0x62, 0xd0, 0x1a, 0xd5, 0xf0, 0xfc, 0xe3, 0xc1, 0xe9,
	0xdd, 0xcf, 0xc0, 0x44, 0xdd, 0xf8, 0xcd, 0x43, 0xdb, 0x34, 0x18, 0x0d, 0xf7, 0x59, 0xc3, 0xfd,
	0xe2, 0x11, 0xd0, 0xfa, 0x29, 0x26, 0xef, 0xbf, 0xbf, 0x4b, 0xa5, 0xb9, 0x5a, 0xc4, 0x3c, 0x51,
	0x37, 0x81, 0x03, 0xd6, 0xbb, 0x47, 0x90, 0x28, 0xc2, 0x20, 0xca, 0x65, 0xa5, 0x1d, 0x85, 0xa9,
	0x0a, 0x6d, 0x13, 0x89, 0x0f, 0xec, 0xcf, 0xeb, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x29, 0x61,
	0x9f, 0xcf, 0xb7, 0x04, 0x00, 0x00,
}
