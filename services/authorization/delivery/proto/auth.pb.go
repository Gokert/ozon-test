// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type FindIdRequest struct {
	Sid                  string   `protobuf:"bytes,1,opt,name=sid,proto3" json:"sid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindIdRequest) Reset()         { *m = FindIdRequest{} }
func (m *FindIdRequest) String() string { return proto.CompactTextString(m) }
func (*FindIdRequest) ProtoMessage()    {}
func (*FindIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{0}
}

func (m *FindIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindIdRequest.Unmarshal(m, b)
}
func (m *FindIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindIdRequest.Marshal(b, m, deterministic)
}
func (m *FindIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindIdRequest.Merge(m, src)
}
func (m *FindIdRequest) XXX_Size() int {
	return xxx_messageInfo_FindIdRequest.Size(m)
}
func (m *FindIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindIdRequest proto.InternalMessageInfo

func (m *FindIdRequest) GetSid() string {
	if m != nil {
		return m.Sid
	}
	return ""
}

type FindIdResponse struct {
	Value                uint64   `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindIdResponse) Reset()         { *m = FindIdResponse{} }
func (m *FindIdResponse) String() string { return proto.CompactTextString(m) }
func (*FindIdResponse) ProtoMessage()    {}
func (*FindIdResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{1}
}

func (m *FindIdResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindIdResponse.Unmarshal(m, b)
}
func (m *FindIdResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindIdResponse.Marshal(b, m, deterministic)
}
func (m *FindIdResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindIdResponse.Merge(m, src)
}
func (m *FindIdResponse) XXX_Size() int {
	return xxx_messageInfo_FindIdResponse.Size(m)
}
func (m *FindIdResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindIdResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindIdResponse proto.InternalMessageInfo

func (m *FindIdResponse) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type AuthorizationCheckRequest struct {
	Sid                  string   `protobuf:"bytes,1,opt,name=sid,proto3" json:"sid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthorizationCheckRequest) Reset()         { *m = AuthorizationCheckRequest{} }
func (m *AuthorizationCheckRequest) String() string { return proto.CompactTextString(m) }
func (*AuthorizationCheckRequest) ProtoMessage()    {}
func (*AuthorizationCheckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{2}
}

func (m *AuthorizationCheckRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthorizationCheckRequest.Unmarshal(m, b)
}
func (m *AuthorizationCheckRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthorizationCheckRequest.Marshal(b, m, deterministic)
}
func (m *AuthorizationCheckRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthorizationCheckRequest.Merge(m, src)
}
func (m *AuthorizationCheckRequest) XXX_Size() int {
	return xxx_messageInfo_AuthorizationCheckRequest.Size(m)
}
func (m *AuthorizationCheckRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthorizationCheckRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthorizationCheckRequest proto.InternalMessageInfo

func (m *AuthorizationCheckRequest) GetSid() string {
	if m != nil {
		return m.Sid
	}
	return ""
}

type AuthorizationCheckResponse struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthorizationCheckResponse) Reset()         { *m = AuthorizationCheckResponse{} }
func (m *AuthorizationCheckResponse) String() string { return proto.CompactTextString(m) }
func (*AuthorizationCheckResponse) ProtoMessage()    {}
func (*AuthorizationCheckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{3}
}

func (m *AuthorizationCheckResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthorizationCheckResponse.Unmarshal(m, b)
}
func (m *AuthorizationCheckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthorizationCheckResponse.Marshal(b, m, deterministic)
}
func (m *AuthorizationCheckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthorizationCheckResponse.Merge(m, src)
}
func (m *AuthorizationCheckResponse) XXX_Size() int {
	return xxx_messageInfo_AuthorizationCheckResponse.Size(m)
}
func (m *AuthorizationCheckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthorizationCheckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthorizationCheckResponse proto.InternalMessageInfo

func (m *AuthorizationCheckResponse) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

type RoleRequest struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoleRequest) Reset()         { *m = RoleRequest{} }
func (m *RoleRequest) String() string { return proto.CompactTextString(m) }
func (*RoleRequest) ProtoMessage()    {}
func (*RoleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{4}
}

func (m *RoleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleRequest.Unmarshal(m, b)
}
func (m *RoleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleRequest.Marshal(b, m, deterministic)
}
func (m *RoleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleRequest.Merge(m, src)
}
func (m *RoleRequest) XXX_Size() int {
	return xxx_messageInfo_RoleRequest.Size(m)
}
func (m *RoleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RoleRequest proto.InternalMessageInfo

func (m *RoleRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type RoleResponse struct {
	Role                 string   `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoleResponse) Reset()         { *m = RoleResponse{} }
func (m *RoleResponse) String() string { return proto.CompactTextString(m) }
func (*RoleResponse) ProtoMessage()    {}
func (*RoleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{5}
}

func (m *RoleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleResponse.Unmarshal(m, b)
}
func (m *RoleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleResponse.Marshal(b, m, deterministic)
}
func (m *RoleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleResponse.Merge(m, src)
}
func (m *RoleResponse) XXX_Size() int {
	return xxx_messageInfo_RoleResponse.Size(m)
}
func (m *RoleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RoleResponse proto.InternalMessageInfo

func (m *RoleResponse) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

type UserItemRequest struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserItemRequest) Reset()         { *m = UserItemRequest{} }
func (m *UserItemRequest) String() string { return proto.CompactTextString(m) }
func (*UserItemRequest) ProtoMessage()    {}
func (*UserItemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{6}
}

func (m *UserItemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserItemRequest.Unmarshal(m, b)
}
func (m *UserItemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserItemRequest.Marshal(b, m, deterministic)
}
func (m *UserItemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserItemRequest.Merge(m, src)
}
func (m *UserItemRequest) XXX_Size() int {
	return xxx_messageInfo_UserItemRequest.Size(m)
}
func (m *UserItemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserItemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserItemRequest proto.InternalMessageInfo

func (m *UserItemRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type UserItemResponse struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserItemResponse) Reset()         { *m = UserItemResponse{} }
func (m *UserItemResponse) String() string { return proto.CompactTextString(m) }
func (*UserItemResponse) ProtoMessage()    {}
func (*UserItemResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{7}
}

func (m *UserItemResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserItemResponse.Unmarshal(m, b)
}
func (m *UserItemResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserItemResponse.Marshal(b, m, deterministic)
}
func (m *UserItemResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserItemResponse.Merge(m, src)
}
func (m *UserItemResponse) XXX_Size() int {
	return xxx_messageInfo_UserItemResponse.Size(m)
}
func (m *UserItemResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserItemResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserItemResponse proto.InternalMessageInfo

func (m *UserItemResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type UserResponse struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Login                string   `protobuf:"bytes,2,opt,name=login,proto3" json:"login,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserResponse) Reset()         { *m = UserResponse{} }
func (m *UserResponse) String() string { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()    {}
func (*UserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{8}
}

func (m *UserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResponse.Unmarshal(m, b)
}
func (m *UserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResponse.Marshal(b, m, deterministic)
}
func (m *UserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResponse.Merge(m, src)
}
func (m *UserResponse) XXX_Size() int {
	return xxx_messageInfo_UserResponse.Size(m)
}
func (m *UserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserResponse proto.InternalMessageInfo

func (m *UserResponse) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserResponse) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

func (m *UserResponse) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func init() {
	proto.RegisterType((*FindIdRequest)(nil), "auth.FindIdRequest")
	proto.RegisterType((*FindIdResponse)(nil), "auth.FindIdResponse")
	proto.RegisterType((*AuthorizationCheckRequest)(nil), "auth.AuthorizationCheckRequest")
	proto.RegisterType((*AuthorizationCheckResponse)(nil), "auth.AuthorizationCheckResponse")
	proto.RegisterType((*RoleRequest)(nil), "auth.RoleRequest")
	proto.RegisterType((*RoleResponse)(nil), "auth.RoleResponse")
	proto.RegisterType((*UserItemRequest)(nil), "auth.UserItemRequest")
	proto.RegisterType((*UserItemResponse)(nil), "auth.UserItemResponse")
	proto.RegisterType((*UserResponse)(nil), "auth.UserResponse")
}

func init() {
	proto.RegisterFile("auth.proto", fileDescriptor_8bbd6f3875b0e874)
}

var fileDescriptor_8bbd6f3875b0e874 = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x4f, 0x4f, 0x83, 0x40,
	0x10, 0xc5, 0x2d, 0xa5, 0x55, 0xa7, 0x7f, 0xac, 0x6b, 0x25, 0x95, 0xc4, 0xd8, 0xee, 0xa1, 0xf1,
	0x62, 0x9b, 0xd4, 0xc6, 0xa3, 0x89, 0x9a, 0x48, 0xea, 0xc1, 0x03, 0xc6, 0x8b, 0x89, 0x87, 0x55,
	0x26, 0x96, 0x08, 0x6c, 0x85, 0xc5, 0x83, 0x9f, 0xd2, 0x8f, 0x64, 0x76, 0x97, 0x42, 0x69, 0xe4,
	0xc4, 0xcc, 0xdb, 0xf7, 0x7e, 0x0b, 0x33, 0x00, 0xb0, 0x54, 0x2c, 0x27, 0xab, 0x98, 0x0b, 0x4e,
	0x4c, 0x59, 0xd3, 0x11, 0x74, 0xee, 0xfd, 0xc8, 0x5b, 0x78, 0x2e, 0x7e, 0xa5, 0x98, 0x08, 0xd2,
	0x83, 0x7a, 0xe2, 0x7b, 0x83, 0xda, 0xb0, 0x76, 0xbe, 0xef, 0xca, 0x92, 0x8e, 0xa1, 0xbb, 0xb6,
	0x24, 0x2b, 0x1e, 0x25, 0x48, 0xfa, 0xd0, 0xf8, 0x66, 0x41, 0x8a, 0xca, 0x65, 0xba, 0xba, 0xa1,
	0x17, 0x70, 0x72, 0x93, 0x8a, 0x25, 0x8f, 0xfd, 0x1f, 0x26, 0x7c, 0x1e, 0xdd, 0x2d, 0xf1, 0xfd,
	0xb3, 0x1a, 0x3b, 0x07, 0xfb, 0x3f, 0x7b, 0x76, 0x85, 0x05, 0xcd, 0x44, 0x30, 0x91, 0x26, 0x2a,
	0xb2, 0xe7, 0x66, 0x1d, 0x3d, 0x85, 0x96, 0xcb, 0x03, 0x5c, 0x63, 0xbb, 0x60, 0x64, 0x54, 0xd3,
	0x35, 0x7c, 0x8f, 0x52, 0x68, 0xeb, 0xe3, 0x0c, 0x43, 0xc0, 0x8c, 0x79, 0x80, 0xd9, 0xbd, 0xaa,
	0xa6, 0x23, 0x38, 0x78, 0x4e, 0x30, 0x5e, 0x08, 0x0c, 0xab, 0x30, 0x63, 0xe8, 0x15, 0x96, 0x02,
	0x15, 0xb1, 0x30, 0x47, 0xc9, 0x9a, 0x3e, 0x40, 0x5b, 0xfa, 0x72, 0xcf, 0x16, 0x47, 0x0e, 0x2a,
	0xe0, 0x1f, 0x7e, 0x34, 0x30, 0x54, 0x48, 0x37, 0x52, 0xc5, 0x90, 0xf9, 0xc1, 0xa0, 0xae, 0x55,
	0xd5, 0xcc, 0x7e, 0x0d, 0xe8, 0x94, 0x06, 0x42, 0xe6, 0xd0, 0x70, 0x50, 0x2c, 0x3c, 0x72, 0x34,
	0x51, 0x7b, 0x2b, 0x2d, 0xca, 0xee, 0x97, 0x45, 0xfd, 0x06, 0x74, 0x87, 0xbc, 0x82, 0xe5, 0xa0,
	0x28, 0x91, 0x9e, 0xd4, 0xec, 0xc8, 0x99, 0x4e, 0x54, 0x2e, 0xc9, 0x1e, 0x56, 0x1b, 0x72, 0xfc,
	0x0c, 0x76, 0x1d, 0x14, 0x72, 0xc8, 0xe4, 0x50, 0xdb, 0x37, 0xf6, 0x61, 0x93, 0x4d, 0x29, 0xcf,
	0x5c, 0x43, 0xcb, 0x41, 0x21, 0x27, 0xf5, 0xc8, 0x42, 0x24, 0xc7, 0xda, 0xb4, 0xb5, 0x04, 0xdb,
	0xda, 0x96, 0xf3, 0xfc, 0x95, 0xba, 0x53, 0x1e, 0x54, 0x65, 0x49, 0x21, 0x17, 0xb9, 0x5b, 0xeb,
	0xa5, 0x3f, 0x65, 0x9b, 0x1f, 0x33, 0x55, 0xbf, 0xfe, 0x5b, 0x53, 0x3d, 0x2e, 0xff, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x43, 0xbf, 0xfa, 0xbb, 0x0f, 0x03, 0x00, 0x00,
}
