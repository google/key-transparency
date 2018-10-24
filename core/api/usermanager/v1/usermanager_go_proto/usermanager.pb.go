// Code generated by protoc-gen-go. DO NOT EDIT.
// source: usermanager/v1/usermanager.proto

package usermanager_go_proto // import "github.com/google/keytransparency/core/api/usermanager/v1/usermanager_go_proto"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import type_go_proto "github.com/google/keytransparency/core/api/type/type_go_proto"
import tink_go_proto "github.com/google/tink/proto/tink_go_proto"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import field_mask "google.golang.org/genproto/protobuf/field_mask"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// GetKeySetsRequest requests the keyset of a directory_id/app_id
type GetKeySetRequest struct {
	// directory_id identifies the directory.
	DirectoryId string `protobuf:"bytes,1,opt,name=directory_id,json=directoryId,proto3" json:"directory_id,omitempty"`
	// app_id identifies the application.
	AppId                string   `protobuf:"bytes,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetKeySetRequest) Reset()         { *m = GetKeySetRequest{} }
func (m *GetKeySetRequest) String() string { return proto.CompactTextString(m) }
func (*GetKeySetRequest) ProtoMessage()    {}
func (*GetKeySetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_usermanager_4313299c84349f00, []int{0}
}
func (m *GetKeySetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetKeySetRequest.Unmarshal(m, b)
}
func (m *GetKeySetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetKeySetRequest.Marshal(b, m, deterministic)
}
func (dst *GetKeySetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetKeySetRequest.Merge(dst, src)
}
func (m *GetKeySetRequest) XXX_Size() int {
	return xxx_messageInfo_GetKeySetRequest.Size(m)
}
func (m *GetKeySetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetKeySetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetKeySetRequest proto.InternalMessageInfo

func (m *GetKeySetRequest) GetDirectoryId() string {
	if m != nil {
		return m.DirectoryId
	}
	return ""
}

func (m *GetKeySetRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

// CreateUserRequest specifies the information with which a new user should be initialized.
// New users will be signed with the current active key.
// It is the responsibility of authorized callers to verify that directory_id/app_id/user_id is correct.
type CreateUserRequest struct {
	// user is the user to create.
	User *type_go_proto.User `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	// add_signing_keys specifies whether to add this service's signing keys to the set of authorized_keys.
	// This must be set to true if any further operations from this API are meant to succeed.
	// If set to false, there must be at least one key in authorized_keys.
	AddSigningKeys       bool     `protobuf:"varint,7,opt,name=add_signing_keys,json=addSigningKeys,proto3" json:"add_signing_keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_usermanager_4313299c84349f00, []int{1}
}
func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (dst *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(dst, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetUser() *type_go_proto.User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *CreateUserRequest) GetAddSigningKeys() bool {
	if m != nil {
		return m.AddSigningKeys
	}
	return false
}

// UpdateUserRequest sets the data field for the user.
// The user must have the service's current signing key in its list of
// authorized_keys in order to succeed.
type UpdateUserRequest struct {
	// user contains data which will be applied to the user.
	User *type_go_proto.User `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	// update_mask specifies which fields of user to update.
	// For example: "data" or "authorized_keys"
	UpdateMask           *field_mask.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdateUserRequest) Reset()         { *m = UpdateUserRequest{} }
func (m *UpdateUserRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateUserRequest) ProtoMessage()    {}
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_usermanager_4313299c84349f00, []int{2}
}
func (m *UpdateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserRequest.Unmarshal(m, b)
}
func (m *UpdateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserRequest.Merge(dst, src)
}
func (m *UpdateUserRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateUserRequest.Size(m)
}
func (m *UpdateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserRequest proto.InternalMessageInfo

func (m *UpdateUserRequest) GetUser() *type_go_proto.User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *UpdateUserRequest) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

// BatchCreateUserRequest creates multiple users all at once.
type BatchCreateUserRequest struct {
	// directory_id identifies the directory.
	DirectoryId string `protobuf:"bytes,1,opt,name=directory_id,json=directoryId,proto3" json:"directory_id,omitempty"`
	// app_id identifies the application.
	AppId string `protobuf:"bytes,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	// users is the set of users to create.
	Users []*type_go_proto.User `protobuf:"bytes,3,rep,name=users,proto3" json:"users,omitempty"`
	// add_signing_keys specifies whether to add this service's signing_keys to the set of authorized_keys.
	// This must be set to true if any further operations from this API are meant to succeed.
	// If set to false, there must be at least one key in authorized_keys.
	AddSigningKeys       bool     `protobuf:"varint,4,opt,name=add_signing_keys,json=addSigningKeys,proto3" json:"add_signing_keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BatchCreateUserRequest) Reset()         { *m = BatchCreateUserRequest{} }
func (m *BatchCreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*BatchCreateUserRequest) ProtoMessage()    {}
func (*BatchCreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_usermanager_4313299c84349f00, []int{3}
}
func (m *BatchCreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchCreateUserRequest.Unmarshal(m, b)
}
func (m *BatchCreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchCreateUserRequest.Marshal(b, m, deterministic)
}
func (dst *BatchCreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchCreateUserRequest.Merge(dst, src)
}
func (m *BatchCreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_BatchCreateUserRequest.Size(m)
}
func (m *BatchCreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchCreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BatchCreateUserRequest proto.InternalMessageInfo

func (m *BatchCreateUserRequest) GetDirectoryId() string {
	if m != nil {
		return m.DirectoryId
	}
	return ""
}

func (m *BatchCreateUserRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *BatchCreateUserRequest) GetUsers() []*type_go_proto.User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *BatchCreateUserRequest) GetAddSigningKeys() bool {
	if m != nil {
		return m.AddSigningKeys
	}
	return false
}

// BatchCreateUserResponse creates multiple users at once.
type BatchCreateUserResponse struct {
	// users returns the list of created users.
	Users                []*type_go_proto.User `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *BatchCreateUserResponse) Reset()         { *m = BatchCreateUserResponse{} }
func (m *BatchCreateUserResponse) String() string { return proto.CompactTextString(m) }
func (*BatchCreateUserResponse) ProtoMessage()    {}
func (*BatchCreateUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_usermanager_4313299c84349f00, []int{4}
}
func (m *BatchCreateUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchCreateUserResponse.Unmarshal(m, b)
}
func (m *BatchCreateUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchCreateUserResponse.Marshal(b, m, deterministic)
}
func (dst *BatchCreateUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchCreateUserResponse.Merge(dst, src)
}
func (m *BatchCreateUserResponse) XXX_Size() int {
	return xxx_messageInfo_BatchCreateUserResponse.Size(m)
}
func (m *BatchCreateUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchCreateUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BatchCreateUserResponse proto.InternalMessageInfo

func (m *BatchCreateUserResponse) GetUsers() []*type_go_proto.User {
	if m != nil {
		return m.Users
	}
	return nil
}

func init() {
	proto.RegisterType((*GetKeySetRequest)(nil), "google.keytransparency.usermanager.v1.GetKeySetRequest")
	proto.RegisterType((*CreateUserRequest)(nil), "google.keytransparency.usermanager.v1.CreateUserRequest")
	proto.RegisterType((*UpdateUserRequest)(nil), "google.keytransparency.usermanager.v1.UpdateUserRequest")
	proto.RegisterType((*BatchCreateUserRequest)(nil), "google.keytransparency.usermanager.v1.BatchCreateUserRequest")
	proto.RegisterType((*BatchCreateUserResponse)(nil), "google.keytransparency.usermanager.v1.BatchCreateUserResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserManagerClient is the client API for UserManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserManagerClient interface {
	// GetKeySet returns a list of public keys (a keyset) that corresponds to the signing keys
	// this service has for a given directory and app.
	GetKeySet(ctx context.Context, in *GetKeySetRequest, opts ...grpc.CallOption) (*tink_go_proto.Keyset, error)
	// CreateUser creates a new user and initializes it.
	// If the user already exists, this operation will fail.
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*type_go_proto.User, error)
	// UpdateUserData sets the public key for an user.
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*type_go_proto.User, error)
	// BatchCreateUser creates a set of new users.
	BatchCreateUser(ctx context.Context, in *BatchCreateUserRequest, opts ...grpc.CallOption) (*type_go_proto.User, error)
}

type userManagerClient struct {
	cc *grpc.ClientConn
}

func NewUserManagerClient(cc *grpc.ClientConn) UserManagerClient {
	return &userManagerClient{cc}
}

func (c *userManagerClient) GetKeySet(ctx context.Context, in *GetKeySetRequest, opts ...grpc.CallOption) (*tink_go_proto.Keyset, error) {
	out := new(tink_go_proto.Keyset)
	err := c.cc.Invoke(ctx, "/google.keytransparency.usermanager.v1.UserManager/GetKeySet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagerClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*type_go_proto.User, error) {
	out := new(type_go_proto.User)
	err := c.cc.Invoke(ctx, "/google.keytransparency.usermanager.v1.UserManager/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagerClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*type_go_proto.User, error) {
	out := new(type_go_proto.User)
	err := c.cc.Invoke(ctx, "/google.keytransparency.usermanager.v1.UserManager/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagerClient) BatchCreateUser(ctx context.Context, in *BatchCreateUserRequest, opts ...grpc.CallOption) (*type_go_proto.User, error) {
	out := new(type_go_proto.User)
	err := c.cc.Invoke(ctx, "/google.keytransparency.usermanager.v1.UserManager/BatchCreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserManagerServer is the server API for UserManager service.
type UserManagerServer interface {
	// GetKeySet returns a list of public keys (a keyset) that corresponds to the signing keys
	// this service has for a given directory and app.
	GetKeySet(context.Context, *GetKeySetRequest) (*tink_go_proto.Keyset, error)
	// CreateUser creates a new user and initializes it.
	// If the user already exists, this operation will fail.
	CreateUser(context.Context, *CreateUserRequest) (*type_go_proto.User, error)
	// UpdateUserData sets the public key for an user.
	UpdateUser(context.Context, *UpdateUserRequest) (*type_go_proto.User, error)
	// BatchCreateUser creates a set of new users.
	BatchCreateUser(context.Context, *BatchCreateUserRequest) (*type_go_proto.User, error)
}

func RegisterUserManagerServer(s *grpc.Server, srv UserManagerServer) {
	s.RegisterService(&_UserManager_serviceDesc, srv)
}

func _UserManager_GetKeySet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKeySetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagerServer).GetKeySet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.keytransparency.usermanager.v1.UserManager/GetKeySet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagerServer).GetKeySet(ctx, req.(*GetKeySetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManager_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagerServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.keytransparency.usermanager.v1.UserManager/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagerServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManager_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagerServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.keytransparency.usermanager.v1.UserManager/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagerServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManager_BatchCreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchCreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagerServer).BatchCreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.keytransparency.usermanager.v1.UserManager/BatchCreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagerServer).BatchCreateUser(ctx, req.(*BatchCreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.keytransparency.usermanager.v1.UserManager",
	HandlerType: (*UserManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetKeySet",
			Handler:    _UserManager_GetKeySet_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _UserManager_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserManager_UpdateUser_Handler,
		},
		{
			MethodName: "BatchCreateUser",
			Handler:    _UserManager_BatchCreateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "usermanager/v1/usermanager.proto",
}

func init() {
	proto.RegisterFile("usermanager/v1/usermanager.proto", fileDescriptor_usermanager_4313299c84349f00)
}

var fileDescriptor_usermanager_4313299c84349f00 = []byte{
	// 573 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xd1, 0x6a, 0x13, 0x41,
	0x14, 0x65, 0x93, 0x36, 0xda, 0x89, 0xd8, 0x76, 0x40, 0x0d, 0xc1, 0x87, 0x74, 0x41, 0x08, 0x3e,
	0xec, 0xd0, 0x88, 0x54, 0x22, 0x22, 0x56, 0xb1, 0x84, 0x5a, 0x95, 0x94, 0x22, 0xf8, 0xb2, 0x4c,
	0x77, 0x6f, 0xb7, 0x4b, 0x9a, 0x99, 0x71, 0x66, 0x52, 0x58, 0x4a, 0x5f, 0x7c, 0xf1, 0x03, 0xfc,
	0x07, 0xc1, 0x0f, 0xf0, 0x0f, 0xfc, 0x00, 0xc1, 0x5f, 0xf0, 0x43, 0x64, 0x66, 0x76, 0x93, 0x90,
	0xa4, 0xba, 0x2d, 0xfa, 0xb2, 0xc9, 0xdc, 0x7b, 0x67, 0xee, 0x39, 0x67, 0xce, 0xdd, 0x45, 0xad,
	0x91, 0x02, 0x39, 0xa4, 0x8c, 0x26, 0x20, 0xc9, 0xe9, 0x26, 0x99, 0x5a, 0x06, 0x42, 0x72, 0xcd,
	0xf1, 0xbd, 0x84, 0xf3, 0xe4, 0x04, 0x82, 0x01, 0x64, 0x5a, 0x52, 0xa6, 0x04, 0x95, 0xc0, 0xa2,
	0x2c, 0x98, 0xae, 0x3c, 0xdd, 0x6c, 0xde, 0x75, 0x65, 0x84, 0x8a, 0x94, 0x50, 0xc6, 0xb8, 0xa6,
	0x3a, 0xe5, 0x4c, 0xb9, 0x43, 0x9a, 0xad, 0x3c, 0x6b, 0x57, 0x87, 0xa3, 0x23, 0x72, 0x94, 0xc2,
	0x49, 0x1c, 0x0e, 0xa9, 0x1a, 0xe4, 0x15, 0xab, 0x3a, 0x13, 0x40, 0xcc, 0x23, 0x0f, 0x20, 0x9d,
	0xb2, 0x3c, 0xe9, 0xbf, 0x42, 0x6b, 0x3b, 0xa0, 0x77, 0x21, 0xdb, 0x07, 0xdd, 0x87, 0x0f, 0x23,
	0x50, 0x1a, 0x6f, 0xa0, 0x1b, 0x71, 0x2a, 0x21, 0xd2, 0x5c, 0x66, 0x61, 0x1a, 0x37, 0xbc, 0x96,
	0xd7, 0x5e, 0xe9, 0xd7, 0xc7, 0xb1, 0x5e, 0x8c, 0x6f, 0xa1, 0x1a, 0x15, 0xc2, 0x24, 0x2b, 0x36,
	0xb9, 0x4c, 0x85, 0xe8, 0xc5, 0xbe, 0x46, 0xeb, 0xcf, 0x25, 0x50, 0x0d, 0x07, 0x0a, 0x64, 0x71,
	0xdc, 0x43, 0xb4, 0x64, 0x18, 0xd9, 0xca, 0x7a, 0x67, 0x23, 0xb8, 0x80, 0xb5, 0x05, 0x68, 0xf7,
	0xd9, 0x72, 0xdc, 0x46, 0x6b, 0x34, 0x8e, 0x43, 0x95, 0x26, 0x2c, 0x65, 0x49, 0x38, 0x80, 0x4c,
	0x35, 0xae, 0xb5, 0xbc, 0xf6, 0xf5, 0xfe, 0x4d, 0x1a, 0xc7, 0xfb, 0x2e, 0xbc, 0x0b, 0x99, 0xf2,
	0x3f, 0x79, 0x68, 0xfd, 0x40, 0xc4, 0xff, 0xa6, 0xed, 0x63, 0x54, 0x1f, 0xd9, 0xb3, 0xac, 0x84,
	0x8d, 0xaa, 0xdd, 0xdd, 0x2c, 0x76, 0x17, 0x2a, 0x07, 0x2f, 0x8d, 0xca, 0x7b, 0x54, 0x0d, 0xfa,
	0xc8, 0x95, 0x9b, 0xff, 0xfe, 0x37, 0x0f, 0xdd, 0xde, 0xa6, 0x3a, 0x3a, 0x9e, 0x57, 0xe1, 0xca,
	0xa2, 0xe2, 0x2d, 0xb4, 0x6c, 0x90, 0xa9, 0x46, 0xb5, 0x55, 0x2d, 0xc7, 0xc4, 0xd5, 0x2f, 0x54,
	0x70, 0x69, 0xa1, 0x82, 0x7d, 0x74, 0x67, 0x0e, 0xb6, 0x12, 0x9c, 0x29, 0x98, 0x74, 0xaf, 0x5c,
	0xae, 0x7b, 0xe7, 0x4b, 0x0d, 0xd5, 0xcd, 0x7a, 0xcf, 0x39, 0x19, 0x7f, 0xf5, 0xd0, 0xca, 0xd8,
	0x6a, 0x78, 0x2b, 0x28, 0x65, 0xfe, 0x60, 0xd6, 0x9c, 0xcd, 0xf1, 0x55, 0x44, 0x32, 0x13, 0x9a,
	0x07, 0xd6, 0xcb, 0x86, 0x04, 0x68, 0x7f, 0xe7, 0xe3, 0xcf, 0x5f, 0x9f, 0x2b, 0xcf, 0xf0, 0x53,
	0x32, 0x33, 0x7b, 0x85, 0xca, 0x29, 0x28, 0x72, 0x36, 0x7d, 0x0d, 0xe7, 0x84, 0x0a, 0xa1, 0xc8,
	0x99, 0xd3, 0xfd, 0x9c, 0x0c, 0xec, 0x41, 0xf8, 0x87, 0x87, 0xd0, 0x44, 0x0b, 0xfc, 0xa8, 0x24,
	0xd8, 0xb9, 0x5b, 0x6f, 0xfe, 0x5d, 0x2e, 0x1f, 0x2c, 0xe8, 0xd0, 0x7f, 0xf7, 0x47, 0xd0, 0x26,
	0x17, 0x2c, 0x42, 0x6e, 0x13, 0x05, 0x7c, 0xab, 0x7c, 0x1e, 0x33, 0x0f, 0x13, 0xec, 0x3a, 0x5f,
	0x1b, 0x4a, 0x93, 0x21, 0x29, 0x4d, 0x69, 0x6e, 0xae, 0x2e, 0x41, 0xa9, 0xf3, 0x9f, 0x29, 0x7d,
	0xf7, 0xd0, 0xea, 0x8c, 0x6d, 0xf1, 0x93, 0x92, 0xbc, 0x16, 0x4f, 0x69, 0x19, 0x72, 0x6f, 0x2c,
	0xb9, 0x9e, 0xff, 0xe2, 0xaa, 0x26, 0xeb, 0x4e, 0xf5, 0xee, 0x7a, 0xf7, 0xb7, 0xdf, 0xbe, 0x7f,
	0x9d, 0xa4, 0xfa, 0x78, 0x74, 0x18, 0x44, 0x7c, 0x48, 0xf2, 0xd7, 0xf9, 0x4c, 0x7f, 0x12, 0x71,
	0xe9, 0xbe, 0x00, 0x17, 0x7f, 0x55, 0xc2, 0x84, 0x87, 0xee, 0xfd, 0x54, 0xb3, 0x3f, 0x0f, 0x7e,
	0x07, 0x00, 0x00, 0xff, 0xff, 0x97, 0xd5, 0x4b, 0xab, 0x83, 0x06, 0x00, 0x00,
}
