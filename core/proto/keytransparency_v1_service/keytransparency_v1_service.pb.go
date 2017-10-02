// Code generated by protoc-gen-go. DO NOT EDIT.
// source: keytransparency_v1_service.proto

/*
Package keytransparency_v1_service is a generated protocol buffer package.

Key Transparency Service

The Key Transparency Service API consists of a map of user names to public
keys. Each user name also has a history of public keys that have been
associated with it.

It is generated from these files:
	keytransparency_v1_service.proto

It has these top-level messages:
*/
package keytransparency_v1_service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import keytransparency_v1_types "github.com/google/keytransparency/core/proto/keytransparency_v1"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for KeyTransparencyService service

type KeyTransparencyServiceClient interface {
	// GetEntry returns a user's entry in the Merkle Tree.
	//
	// Entries contain signed commitments to a profile, which is also returned.
	GetEntry(ctx context.Context, in *keytransparency_v1_types.GetEntryRequest, opts ...grpc.CallOption) (*keytransparency_v1_types.GetEntryResponse, error)
	// ListEntryHistory returns a list of historic GetEntry values.
	//
	// Clients verify their account history by observing correct values for their
	// account over time.
	ListEntryHistory(ctx context.Context, in *keytransparency_v1_types.ListEntryHistoryRequest, opts ...grpc.CallOption) (*keytransparency_v1_types.ListEntryHistoryResponse, error)
	// UpdateEntry updates a user's profile.
	//
	// Returns the current user profile.
	// Clients must retry until this function returns a proof containing the desired value.
	UpdateEntry(ctx context.Context, in *keytransparency_v1_types.UpdateEntryRequest, opts ...grpc.CallOption) (*keytransparency_v1_types.UpdateEntryResponse, error)
	// GetDomainInfo returns all info tied to the specified domain.
	//
	// This API to get all necessary data needed to verify a particular
	// key-server. Data contains for instance the tree-info, like for instance the
	// log-/map-id and the corresponding public-keys.
	GetDomainInfo(ctx context.Context, in *keytransparency_v1_types.GetDomainInfoRequest, opts ...grpc.CallOption) (*keytransparency_v1_types.GetDomainInfoResponse, error)
}

type keyTransparencyServiceClient struct {
	cc *grpc.ClientConn
}

func NewKeyTransparencyServiceClient(cc *grpc.ClientConn) KeyTransparencyServiceClient {
	return &keyTransparencyServiceClient{cc}
}

func (c *keyTransparencyServiceClient) GetEntry(ctx context.Context, in *keytransparency_v1_types.GetEntryRequest, opts ...grpc.CallOption) (*keytransparency_v1_types.GetEntryResponse, error) {
	out := new(keytransparency_v1_types.GetEntryResponse)
	err := grpc.Invoke(ctx, "/keytransparency.v1.service.KeyTransparencyService/GetEntry", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyTransparencyServiceClient) ListEntryHistory(ctx context.Context, in *keytransparency_v1_types.ListEntryHistoryRequest, opts ...grpc.CallOption) (*keytransparency_v1_types.ListEntryHistoryResponse, error) {
	out := new(keytransparency_v1_types.ListEntryHistoryResponse)
	err := grpc.Invoke(ctx, "/keytransparency.v1.service.KeyTransparencyService/ListEntryHistory", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyTransparencyServiceClient) UpdateEntry(ctx context.Context, in *keytransparency_v1_types.UpdateEntryRequest, opts ...grpc.CallOption) (*keytransparency_v1_types.UpdateEntryResponse, error) {
	out := new(keytransparency_v1_types.UpdateEntryResponse)
	err := grpc.Invoke(ctx, "/keytransparency.v1.service.KeyTransparencyService/UpdateEntry", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyTransparencyServiceClient) GetDomainInfo(ctx context.Context, in *keytransparency_v1_types.GetDomainInfoRequest, opts ...grpc.CallOption) (*keytransparency_v1_types.GetDomainInfoResponse, error) {
	out := new(keytransparency_v1_types.GetDomainInfoResponse)
	err := grpc.Invoke(ctx, "/keytransparency.v1.service.KeyTransparencyService/GetDomainInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for KeyTransparencyService service

type KeyTransparencyServiceServer interface {
	// GetEntry returns a user's entry in the Merkle Tree.
	//
	// Entries contain signed commitments to a profile, which is also returned.
	GetEntry(context.Context, *keytransparency_v1_types.GetEntryRequest) (*keytransparency_v1_types.GetEntryResponse, error)
	// ListEntryHistory returns a list of historic GetEntry values.
	//
	// Clients verify their account history by observing correct values for their
	// account over time.
	ListEntryHistory(context.Context, *keytransparency_v1_types.ListEntryHistoryRequest) (*keytransparency_v1_types.ListEntryHistoryResponse, error)
	// UpdateEntry updates a user's profile.
	//
	// Returns the current user profile.
	// Clients must retry until this function returns a proof containing the desired value.
	UpdateEntry(context.Context, *keytransparency_v1_types.UpdateEntryRequest) (*keytransparency_v1_types.UpdateEntryResponse, error)
	// GetDomainInfo returns all info tied to the specified domain.
	//
	// This API to get all necessary data needed to verify a particular
	// key-server. Data contains for instance the tree-info, like for instance the
	// log-/map-id and the corresponding public-keys.
	GetDomainInfo(context.Context, *keytransparency_v1_types.GetDomainInfoRequest) (*keytransparency_v1_types.GetDomainInfoResponse, error)
}

func RegisterKeyTransparencyServiceServer(s *grpc.Server, srv KeyTransparencyServiceServer) {
	s.RegisterService(&_KeyTransparencyService_serviceDesc, srv)
}

func _KeyTransparencyService_GetEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(keytransparency_v1_types.GetEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyTransparencyServiceServer).GetEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keytransparency.v1.service.KeyTransparencyService/GetEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyTransparencyServiceServer).GetEntry(ctx, req.(*keytransparency_v1_types.GetEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyTransparencyService_ListEntryHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(keytransparency_v1_types.ListEntryHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyTransparencyServiceServer).ListEntryHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keytransparency.v1.service.KeyTransparencyService/ListEntryHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyTransparencyServiceServer).ListEntryHistory(ctx, req.(*keytransparency_v1_types.ListEntryHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyTransparencyService_UpdateEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(keytransparency_v1_types.UpdateEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyTransparencyServiceServer).UpdateEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keytransparency.v1.service.KeyTransparencyService/UpdateEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyTransparencyServiceServer).UpdateEntry(ctx, req.(*keytransparency_v1_types.UpdateEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyTransparencyService_GetDomainInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(keytransparency_v1_types.GetDomainInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyTransparencyServiceServer).GetDomainInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keytransparency.v1.service.KeyTransparencyService/GetDomainInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyTransparencyServiceServer).GetDomainInfo(ctx, req.(*keytransparency_v1_types.GetDomainInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KeyTransparencyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "keytransparency.v1.service.KeyTransparencyService",
	HandlerType: (*KeyTransparencyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEntry",
			Handler:    _KeyTransparencyService_GetEntry_Handler,
		},
		{
			MethodName: "ListEntryHistory",
			Handler:    _KeyTransparencyService_ListEntryHistory_Handler,
		},
		{
			MethodName: "UpdateEntry",
			Handler:    _KeyTransparencyService_UpdateEntry_Handler,
		},
		{
			MethodName: "GetDomainInfo",
			Handler:    _KeyTransparencyService_GetDomainInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "keytransparency_v1_service.proto",
}

// Client API for KeyTransparencyAdminService service

type KeyTransparencyAdminServiceClient interface {
	// BatchSetEntries uses an authorized_public key to perform a set request on multiple entries at once.
	BatchUpdateEntries(ctx context.Context, in *keytransparency_v1_types.BatchUpdateEntriesRequest, opts ...grpc.CallOption) (*keytransparency_v1_types.BatchUpdateEntriesResponse, error)
}

type keyTransparencyAdminServiceClient struct {
	cc *grpc.ClientConn
}

func NewKeyTransparencyAdminServiceClient(cc *grpc.ClientConn) KeyTransparencyAdminServiceClient {
	return &keyTransparencyAdminServiceClient{cc}
}

func (c *keyTransparencyAdminServiceClient) BatchUpdateEntries(ctx context.Context, in *keytransparency_v1_types.BatchUpdateEntriesRequest, opts ...grpc.CallOption) (*keytransparency_v1_types.BatchUpdateEntriesResponse, error) {
	out := new(keytransparency_v1_types.BatchUpdateEntriesResponse)
	err := grpc.Invoke(ctx, "/keytransparency.v1.service.KeyTransparencyAdminService/BatchUpdateEntries", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for KeyTransparencyAdminService service

type KeyTransparencyAdminServiceServer interface {
	// BatchSetEntries uses an authorized_public key to perform a set request on multiple entries at once.
	BatchUpdateEntries(context.Context, *keytransparency_v1_types.BatchUpdateEntriesRequest) (*keytransparency_v1_types.BatchUpdateEntriesResponse, error)
}

func RegisterKeyTransparencyAdminServiceServer(s *grpc.Server, srv KeyTransparencyAdminServiceServer) {
	s.RegisterService(&_KeyTransparencyAdminService_serviceDesc, srv)
}

func _KeyTransparencyAdminService_BatchUpdateEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(keytransparency_v1_types.BatchUpdateEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyTransparencyAdminServiceServer).BatchUpdateEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keytransparency.v1.service.KeyTransparencyAdminService/BatchUpdateEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyTransparencyAdminServiceServer).BatchUpdateEntries(ctx, req.(*keytransparency_v1_types.BatchUpdateEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KeyTransparencyAdminService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "keytransparency.v1.service.KeyTransparencyAdminService",
	HandlerType: (*KeyTransparencyAdminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BatchUpdateEntries",
			Handler:    _KeyTransparencyAdminService_BatchUpdateEntries_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "keytransparency_v1_service.proto",
}

func init() { proto.RegisterFile("keytransparency_v1_service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xcf, 0x4b, 0xe3, 0x40,
	0x14, 0xc7, 0xc9, 0xb2, 0x2c, 0xcb, 0xec, 0x2e, 0xdd, 0x9d, 0x45, 0x85, 0x54, 0x41, 0xda, 0x53,
	0x45, 0x33, 0xa4, 0xf5, 0xd4, 0x9b, 0xa2, 0x54, 0xd1, 0x93, 0x3f, 0xce, 0x61, 0x9a, 0xbc, 0xb6,
	0x83, 0x76, 0x26, 0xce, 0x4c, 0x02, 0x41, 0xf4, 0xe0, 0xcd, 0xb3, 0x08, 0x9e, 0xfc, 0x3b, 0xfc,
	0x3f, 0x3c, 0x7b, 0xf3, 0x0f, 0x91, 0xfc, 0xb2, 0x35, 0x26, 0xd2, 0x9e, 0x72, 0x78, 0x9f, 0xf7,
	0x7d, 0x9f, 0x79, 0x99, 0x41, 0xab, 0x67, 0x10, 0x69, 0x49, 0xb9, 0xf2, 0xa9, 0x04, 0xee, 0x46,
	0x4e, 0x68, 0x3b, 0x0a, 0x64, 0xc8, 0x5c, 0xb0, 0x7c, 0x29, 0xb4, 0xc0, 0x66, 0x81, 0xb0, 0x42,
	0xdb, 0xca, 0x08, 0xd3, 0x1b, 0x32, 0x3d, 0x0a, 0xfa, 0x96, 0x2b, 0xc6, 0x64, 0x28, 0xc4, 0xf0,
	0x1c, 0x48, 0x81, 0x26, 0xae, 0x90, 0x40, 0x92, 0x24, 0x52, 0x32, 0x4a, 0x47, 0x3e, 0xa8, 0xca,
	0x42, 0x6a, 0x60, 0x2e, 0x67, 0xd1, 0xd4, 0x67, 0x84, 0x72, 0x2e, 0x34, 0xd5, 0x4c, 0xf0, 0xac,
	0xda, 0x7e, 0xf9, 0x8e, 0x16, 0x0f, 0x20, 0x3a, 0x99, 0x0a, 0x38, 0x4e, 0xf5, 0xf0, 0x35, 0xfa,
	0xd9, 0x03, 0xbd, 0xcb, 0xb5, 0x8c, 0x70, 0xcb, 0x2a, 0x39, 0x47, 0x3a, 0x25, 0x67, 0x8e, 0xe0,
	0x22, 0x00, 0xa5, 0xcd, 0xb5, 0x59, 0x50, 0xe5, 0x0b, 0xae, 0xa0, 0x51, 0xbf, 0x79, 0x7e, 0xbd,
	0xfb, 0xb6, 0x80, 0xff, 0x93, 0xd0, 0x26, 0x81, 0x02, 0xa9, 0xc8, 0x65, 0xfc, 0x71, 0x98, 0x77,
	0x85, 0x1f, 0x0d, 0xf4, 0xf7, 0x90, 0xa9, 0xb4, 0x65, 0x8f, 0x29, 0x2d, 0x64, 0x84, 0xed, 0xea,
	0xf4, 0x22, 0x9b, 0x0b, 0xb5, 0xe7, 0x69, 0xc9, 0xc4, 0x9a, 0x89, 0xd8, 0x0a, 0xae, 0x97, 0x88,
	0x91, 0x51, 0xe6, 0x72, 0x6f, 0xa0, 0x5f, 0xa7, 0xbe, 0x47, 0x35, 0xa4, 0x4b, 0x5a, 0xaf, 0x1e,
	0x34, 0x85, 0xe5, 0x5a, 0x1b, 0x33, 0xd2, 0x99, 0x51, 0x2b, 0x31, 0x6a, 0x9a, 0x65, 0xab, 0xea,
	0xfe, 0x86, 0x98, 0x75, 0x82, 0xa4, 0x0f, 0xdf, 0x1a, 0xe8, 0x4f, 0x0f, 0xf4, 0x8e, 0x18, 0x53,
	0xc6, 0xf7, 0xf9, 0x40, 0x60, 0xeb, 0xcb, 0x7f, 0x32, 0x01, 0x73, 0x37, 0x32, 0x33, 0x9f, 0xd9,
	0x2d, 0x25, 0x76, 0xff, 0x70, 0x2d, 0xb6, 0xf3, 0x92, 0x3a, 0x61, 0x7c, 0x20, 0xda, 0x4f, 0x06,
	0xaa, 0x17, 0xee, 0xd7, 0x96, 0x37, 0x66, 0x3c, 0xbf, 0x64, 0x0f, 0x06, 0xc2, 0xdb, 0x54, 0xbb,
	0xa3, 0xc9, 0x99, 0x19, 0x28, 0xdc, 0xa9, 0x16, 0xf8, 0x4c, 0xe7, 0xd6, 0x9b, 0xf3, 0x35, 0x7d,
	0x54, 0x6f, 0xd4, 0xde, 0x17, 0xdb, 0xed, 0xc7, 0x74, 0xff, 0x47, 0xf2, 0x42, 0x3a, 0x6f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x66, 0xde, 0xa2, 0xf5, 0xe5, 0x03, 0x00, 0x00,
}
