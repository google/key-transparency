// Code generated by protoc-gen-go.
// source: authorization.proto
// DO NOT EDIT!

/*
Package authorization is a generated protocol buffer package.

It is generated from these files:
	authorization.proto

It has these top-level messages:
	AuthorizationPolicy
*/
package authorization

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Permission specifies the permission a role can take.
type Permission int32

const (
	// LOG indicates whether access to resources will be logged.
	Permission_LOG   Permission = 0
	Permission_READ  Permission = 1
	Permission_WRITE Permission = 2
)

var Permission_name = map[int32]string{
	0: "LOG",
	1: "READ",
	2: "WRITE",
}
var Permission_value = map[string]int32{
	"LOG":   0,
	"READ":  1,
	"WRITE": 2,
}

func (x Permission) String() string {
	return proto.EnumName(Permission_name, int32(x))
}
func (Permission) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// AuthorizationPolicy contains an authorization policy.
type AuthorizationPolicy struct {
	// roles is a map of roles keyed by labels used in RoleLabels.
	Roles map[string]*AuthorizationPolicy_Role `protobuf:"bytes,2,rep,name=roles" json:"roles,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// resource_to_role_labels specifies the authorization policy keyed by resource
	// map_id|app_id concatenation as a string.
	ResourceToRoleLabels map[string]*AuthorizationPolicy_RoleLabels `protobuf:"bytes,3,rep,name=resource_to_role_labels,json=resourceToRoleLabels" json:"resource_to_role_labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *AuthorizationPolicy) Reset()                    { *m = AuthorizationPolicy{} }
func (m *AuthorizationPolicy) String() string            { return proto.CompactTextString(m) }
func (*AuthorizationPolicy) ProtoMessage()               {}
func (*AuthorizationPolicy) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AuthorizationPolicy) GetRoles() map[string]*AuthorizationPolicy_Role {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *AuthorizationPolicy) GetResourceToRoleLabels() map[string]*AuthorizationPolicy_RoleLabels {
	if m != nil {
		return m.ResourceToRoleLabels
	}
	return nil
}

// Resource contains the resource being accessed.
type AuthorizationPolicy_Resource struct {
	// map_id contains the Key Transparency map identity for this entry.
	MapId int64 `protobuf:"varint,1,opt,name=map_id,json=mapId" json:"map_id,omitempty"`
	// app_id contains the application identity for this entry.
	AppId string `protobuf:"bytes,2,opt,name=app_id,json=appId" json:"app_id,omitempty"`
}

func (m *AuthorizationPolicy_Resource) Reset()                    { *m = AuthorizationPolicy_Resource{} }
func (m *AuthorizationPolicy_Resource) String() string            { return proto.CompactTextString(m) }
func (*AuthorizationPolicy_Resource) ProtoMessage()               {}
func (*AuthorizationPolicy_Resource) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *AuthorizationPolicy_Resource) GetMapId() int64 {
	if m != nil {
		return m.MapId
	}
	return 0
}

func (m *AuthorizationPolicy_Resource) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

// Role contains a specific identity of an authorization entry.
type AuthorizationPolicy_Role struct {
	// principals contains an application specific identifier for this entry.
	Principals []string `protobuf:"bytes,1,rep,name=principals" json:"principals,omitempty"`
	// permissions contains the set of permissions this role is granted.
	Permissions []Permission `protobuf:"varint,2,rep,packed,name=permissions,enum=authorization.Permission" json:"permissions,omitempty"`
}

func (m *AuthorizationPolicy_Role) Reset()                    { *m = AuthorizationPolicy_Role{} }
func (m *AuthorizationPolicy_Role) String() string            { return proto.CompactTextString(m) }
func (*AuthorizationPolicy_Role) ProtoMessage()               {}
func (*AuthorizationPolicy_Role) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

func (m *AuthorizationPolicy_Role) GetPrincipals() []string {
	if m != nil {
		return m.Principals
	}
	return nil
}

func (m *AuthorizationPolicy_Role) GetPermissions() []Permission {
	if m != nil {
		return m.Permissions
	}
	return nil
}

// RoleLabels contains a lot of role labels identifying each role.
type AuthorizationPolicy_RoleLabels struct {
	Labels []string `protobuf:"bytes,1,rep,name=labels" json:"labels,omitempty"`
}

func (m *AuthorizationPolicy_RoleLabels) Reset()         { *m = AuthorizationPolicy_RoleLabels{} }
func (m *AuthorizationPolicy_RoleLabels) String() string { return proto.CompactTextString(m) }
func (*AuthorizationPolicy_RoleLabels) ProtoMessage()    {}
func (*AuthorizationPolicy_RoleLabels) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 2}
}

func (m *AuthorizationPolicy_RoleLabels) GetLabels() []string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func init() {
	proto.RegisterType((*AuthorizationPolicy)(nil), "authorization.AuthorizationPolicy")
	proto.RegisterType((*AuthorizationPolicy_Resource)(nil), "authorization.AuthorizationPolicy.Resource")
	proto.RegisterType((*AuthorizationPolicy_Role)(nil), "authorization.AuthorizationPolicy.Role")
	proto.RegisterType((*AuthorizationPolicy_RoleLabels)(nil), "authorization.AuthorizationPolicy.RoleLabels")
	proto.RegisterEnum("authorization.Permission", Permission_name, Permission_value)
}

func init() { proto.RegisterFile("authorization.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x5b, 0x6b, 0xfa, 0x30,
	0x1c, 0xfd, 0xb7, 0xb5, 0xfe, 0xed, 0x4f, 0x36, 0x4a, 0xdc, 0xa5, 0xf6, 0x61, 0x94, 0x31, 0x58,
	0x11, 0xe6, 0x83, 0x7b, 0x91, 0x5d, 0x1e, 0xc4, 0xc9, 0x10, 0x84, 0x49, 0x10, 0xf6, 0x58, 0x62,
	0x0d, 0x2c, 0xac, 0x36, 0x21, 0xa9, 0x82, 0xfb, 0xd2, 0xfb, 0x0a, 0xa3, 0x8d, 0x9b, 0x75, 0x38,
	0xe8, 0x5b, 0x7b, 0x7a, 0x2e, 0x3d, 0x27, 0x81, 0x16, 0x59, 0x65, 0x6f, 0x5c, 0xb2, 0x0f, 0x92,
	0x31, 0x9e, 0x76, 0x85, 0xe4, 0x19, 0x47, 0x47, 0x7b, 0xe0, 0xe5, 0x67, 0x0d, 0x5a, 0x83, 0x32,
	0x32, 0xe5, 0x09, 0x8b, 0x37, 0x68, 0x08, 0xb6, 0xe4, 0x09, 0x55, 0x9e, 0x19, 0x58, 0x61, 0xb3,
	0x77, 0xd3, 0xdd, 0xf7, 0x3a, 0x20, 0xe9, 0xe2, 0x9c, 0x3f, 0x4a, 0x33, 0xb9, 0xc1, 0x5a, 0x8b,
	0x14, 0x9c, 0x4b, 0xaa, 0xf8, 0x4a, 0xc6, 0x34, 0xca, 0x78, 0x94, 0x83, 0x51, 0x42, 0xe6, 0x34,
	0x51, 0x9e, 0x55, 0xd8, 0x3e, 0x54, 0xb1, 0xdd, 0x3a, 0xcc, 0x78, 0x1e, 0x30, 0x29, 0xe4, 0x3a,
	0xe5, 0x44, 0x1e, 0xf8, 0xe4, 0xf7, 0xa1, 0xf1, 0x2d, 0x41, 0xa7, 0x50, 0x5f, 0x12, 0x11, 0xb1,
	0x85, 0x67, 0x04, 0x46, 0x68, 0x61, 0x7b, 0x49, 0xc4, 0x78, 0x91, 0xc3, 0x44, 0x14, 0xb0, 0x19,
	0x18, 0xa1, 0x83, 0x6d, 0x22, 0xc4, 0x78, 0xe1, 0xc7, 0x50, 0xcb, 0x7d, 0xd0, 0x05, 0x80, 0x90,
	0x2c, 0x8d, 0x99, 0x20, 0x89, 0xf2, 0x8c, 0xc0, 0x0a, 0x1d, 0x5c, 0x42, 0xd0, 0x3d, 0x34, 0x05,
	0x95, 0x4b, 0xa6, 0x14, 0xe3, 0xa9, 0x5e, 0xe8, 0xb8, 0xd7, 0xfe, 0x55, 0x65, 0xfa, 0xc3, 0xc0,
	0x65, 0xb6, 0x7f, 0x05, 0xb0, 0xfb, 0x59, 0x74, 0x06, 0xf5, 0xed, 0x20, 0x3a, 0x66, 0xfb, 0xe6,
	0x13, 0xcd, 0xd2, 0x45, 0x91, 0x0b, 0xd6, 0x3b, 0xdd, 0x14, 0x1d, 0x1c, 0x9c, 0x3f, 0xa2, 0x47,
	0xb0, 0xd7, 0x24, 0x59, 0xd1, 0xa2, 0x40, 0xb3, 0x77, 0x5d, 0xf1, 0x78, 0xb0, 0x56, 0xdd, 0x99,
	0x7d, 0xc3, 0x5f, 0x43, 0xfb, 0xcf, 0x69, 0x0f, 0x24, 0x0e, 0xf7, 0x13, 0xab, 0x5e, 0x08, 0x6d,
	0x5a, 0xca, 0xed, 0x74, 0x00, 0x76, 0xdb, 0xa0, 0xff, 0x60, 0x4d, 0x5e, 0x9e, 0xdd, 0x7f, 0xa8,
	0x01, 0x35, 0x3c, 0x1a, 0x3c, 0xb9, 0x06, 0x72, 0xc0, 0x7e, 0xc5, 0xe3, 0xd9, 0xc8, 0x35, 0xe7,
	0xf5, 0xe2, 0xce, 0xde, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0xd6, 0x43, 0x2b, 0x46, 0xca, 0x02,
	0x00, 0x00,
}
