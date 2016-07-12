// Code generated by protoc-gen-go.
// source: proto/security_e2ekeys_v2/e2ekeys.proto
// DO NOT EDIT!

/*
Package security_e2ekeys_v2 is a generated protocol buffer package.

It is generated from these files:
	proto/security_e2ekeys_v2/e2ekeys.proto

It has these top-level messages:
*/
package security_e2ekeys_v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import security_e2ekeys "github.com/google/e2e-key-server/proto/security_e2ekeys"
import _ "github.com/gengo/grpc-gateway/third_party/googleapis/google/api"

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
const _ = grpc.SupportPackageIsVersion3

// Client API for E2EKeyService service

type E2EKeyServiceClient interface {
	// GetEntry returns a user's entry in the Merkle Tree. Entries contain
	// signed commitments to a profile, which is also returned.
	GetEntry(ctx context.Context, in *security_e2ekeys.GetEntryRequest, opts ...grpc.CallOption) (*security_e2ekeys.GetEntryResponse, error)
	// ListEntryHistory returns a list of GetEntryRespons covering several epochs.
	ListEntryHistory(ctx context.Context, in *security_e2ekeys.ListEntryHistoryRequest, opts ...grpc.CallOption) (*security_e2ekeys.ListEntryHistoryResponse, error)
	// blocking or polling?
	// UpdateEntry submits a SignedEntryUpdate.  Returns empty until this update
	// has been included in an epoch.  Clients must retry until this function
	// returns a proof.
	UpdateEntry(ctx context.Context, in *security_e2ekeys.UpdateEntryRequest, opts ...grpc.CallOption) (*security_e2ekeys.UpdateEntryResponse, error)
}

type e2EKeyServiceClient struct {
	cc *grpc.ClientConn
}

func NewE2EKeyServiceClient(cc *grpc.ClientConn) E2EKeyServiceClient {
	return &e2EKeyServiceClient{cc}
}

func (c *e2EKeyServiceClient) GetEntry(ctx context.Context, in *security_e2ekeys.GetEntryRequest, opts ...grpc.CallOption) (*security_e2ekeys.GetEntryResponse, error) {
	out := new(security_e2ekeys.GetEntryResponse)
	err := grpc.Invoke(ctx, "/security.e2ekeys.v2.E2EKeyService/GetEntry", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *e2EKeyServiceClient) ListEntryHistory(ctx context.Context, in *security_e2ekeys.ListEntryHistoryRequest, opts ...grpc.CallOption) (*security_e2ekeys.ListEntryHistoryResponse, error) {
	out := new(security_e2ekeys.ListEntryHistoryResponse)
	err := grpc.Invoke(ctx, "/security.e2ekeys.v2.E2EKeyService/ListEntryHistory", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *e2EKeyServiceClient) UpdateEntry(ctx context.Context, in *security_e2ekeys.UpdateEntryRequest, opts ...grpc.CallOption) (*security_e2ekeys.UpdateEntryResponse, error) {
	out := new(security_e2ekeys.UpdateEntryResponse)
	err := grpc.Invoke(ctx, "/security.e2ekeys.v2.E2EKeyService/UpdateEntry", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for E2EKeyService service

type E2EKeyServiceServer interface {
	// GetEntry returns a user's entry in the Merkle Tree. Entries contain
	// signed commitments to a profile, which is also returned.
	GetEntry(context.Context, *security_e2ekeys.GetEntryRequest) (*security_e2ekeys.GetEntryResponse, error)
	// ListEntryHistory returns a list of GetEntryRespons covering several epochs.
	ListEntryHistory(context.Context, *security_e2ekeys.ListEntryHistoryRequest) (*security_e2ekeys.ListEntryHistoryResponse, error)
	// blocking or polling?
	// UpdateEntry submits a SignedEntryUpdate.  Returns empty until this update
	// has been included in an epoch.  Clients must retry until this function
	// returns a proof.
	UpdateEntry(context.Context, *security_e2ekeys.UpdateEntryRequest) (*security_e2ekeys.UpdateEntryResponse, error)
}

func RegisterE2EKeyServiceServer(s *grpc.Server, srv E2EKeyServiceServer) {
	s.RegisterService(&_E2EKeyService_serviceDesc, srv)
}

func _E2EKeyService_GetEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(security_e2ekeys.GetEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(E2EKeyServiceServer).GetEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/security.e2ekeys.v2.E2EKeyService/GetEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(E2EKeyServiceServer).GetEntry(ctx, req.(*security_e2ekeys.GetEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _E2EKeyService_ListEntryHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(security_e2ekeys.ListEntryHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(E2EKeyServiceServer).ListEntryHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/security.e2ekeys.v2.E2EKeyService/ListEntryHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(E2EKeyServiceServer).ListEntryHistory(ctx, req.(*security_e2ekeys.ListEntryHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _E2EKeyService_UpdateEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(security_e2ekeys.UpdateEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(E2EKeyServiceServer).UpdateEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/security.e2ekeys.v2.E2EKeyService/UpdateEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(E2EKeyServiceServer).UpdateEntry(ctx, req.(*security_e2ekeys.UpdateEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _E2EKeyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "security.e2ekeys.v2.E2EKeyService",
	HandlerType: (*E2EKeyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEntry",
			Handler:    _E2EKeyService_GetEntry_Handler,
		},
		{
			MethodName: "ListEntryHistory",
			Handler:    _E2EKeyService_ListEntryHistory_Handler,
		},
		{
			MethodName: "UpdateEntry",
			Handler:    _E2EKeyService_UpdateEntry_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("proto/security_e2ekeys_v2/e2ekeys.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x91, 0xcf, 0x4a, 0xc4, 0x30,
	0x10, 0x87, 0x59, 0x05, 0x91, 0xaa, 0x20, 0x59, 0xbc, 0x74, 0xf5, 0xe0, 0xae, 0x22, 0x2b, 0x6c,
	0x03, 0xf5, 0xe6, 0xbd, 0x28, 0xac, 0x27, 0xc5, 0x73, 0xe9, 0x9f, 0xa1, 0x1b, 0x5c, 0x93, 0xda,
	0x4c, 0x0b, 0x41, 0xbc, 0x08, 0x5e, 0xbd, 0x78, 0xf1, 0xbd, 0x7c, 0x05, 0x1f, 0xc4, 0x34, 0x9b,
	0x80, 0xd8, 0xa2, 0x9e, 0x1a, 0x3a, 0xdf, 0xcc, 0xef, 0xcb, 0xc4, 0x3b, 0x29, 0x2b, 0x81, 0x82,
	0x4a, 0xc8, 0xea, 0x8a, 0xa1, 0x8a, 0x21, 0x84, 0x3b, 0x50, 0x32, 0x6e, 0x42, 0x6a, 0x8f, 0x81,
	0x21, 0xc8, 0xd0, 0x21, 0x81, 0xfb, 0xdf, 0x84, 0xfe, 0xbc, 0x60, 0xb8, 0xa8, 0xd3, 0x20, 0x13,
	0xf7, 0xb4, 0xc8, 0x53, 0x58, 0x36, 0x8c, 0xb7, 0x7d, 0x33, 0x0d, 0xcc, 0x24, 0x54, 0x0d, 0x54,
	0xb4, 0x3f, 0xc0, 0x4e, 0x47, 0x55, 0x82, 0x4d, 0xf0, 0xf7, 0x0b, 0x21, 0x8a, 0x25, 0xd0, 0xa4,
	0x64, 0x34, 0xe1, 0x5c, 0x60, 0x82, 0x4c, 0x70, 0x5b, 0x0d, 0xdf, 0xd7, 0xbd, 0x9d, 0x28, 0x8c,
	0xe6, 0xa0, 0x6e, 0xf4, 0x5c, 0x96, 0x01, 0xe1, 0xde, 0xe6, 0x05, 0x60, 0xc4, 0xb1, 0x52, 0xe4,
	0x30, 0xe8, 0xe8, 0xb9, 0xda, 0x35, 0x3c, 0xd4, 0x20, 0xd1, 0x1f, 0xff, 0x86, 0xc8, 0x52, 0x27,
	0xc1, 0x78, 0xf4, 0xfc, 0xf1, 0xf9, 0xb6, 0xb6, 0x47, 0x86, 0x54, 0xdf, 0xbf, 0xd6, 0x17, 0x90,
	0xf4, 0xb1, 0xfd, 0xc4, 0x2c, 0x7f, 0x22, 0xaf, 0x03, 0x6f, 0xf7, 0x8a, 0xc9, 0x55, 0xcb, 0xa5,
	0x3e, 0x08, 0x1d, 0x3c, 0xed, 0x4e, 0xfd, 0xc9, 0x38, 0x81, 0xd3, 0xff, 0xa0, 0x56, 0x64, 0x62,
	0x44, 0x0e, 0xc8, 0xa8, 0x47, 0x84, 0x2e, 0x6c, 0xf6, 0xcb, 0xc0, 0xdb, 0xba, 0x2d, 0xf3, 0x04,
	0x61, 0xb5, 0x84, 0xa3, 0x6e, 0xc0, 0xb7, 0xb2, 0xd3, 0x38, 0xfe, 0x83, 0xb2, 0x06, 0x53, 0x63,
	0x30, 0xf1, 0xfb, 0x56, 0x71, 0xbe, 0x0d, 0x2d, 0x1b, 0xd7, 0xa6, 0x2f, 0xdd, 0x30, 0x2f, 0x74,
	0xf6, 0x15, 0x00, 0x00, 0xff, 0xff, 0x0a, 0x0b, 0x64, 0x2d, 0x4c, 0x02, 0x00, 0x00,
}
