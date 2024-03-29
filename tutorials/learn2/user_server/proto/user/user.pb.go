// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user/user.proto

package smtl_micro_learn_srv_user

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

type User struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Pwd                  string   `protobuf:"bytes,3,opt,name=pwd,proto3" json:"pwd,omitempty"`
	CreatedTime          uint64   `protobuf:"varint,4,opt,name=createdTime,proto3" json:"createdTime,omitempty"`
	UpdatedTime          uint64   `protobuf:"varint,5,opt,name=updatedTime,proto3" json:"updatedTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetPwd() string {
	if m != nil {
		return m.Pwd
	}
	return ""
}

func (m *User) GetCreatedTime() uint64 {
	if m != nil {
		return m.CreatedTime
	}
	return 0
}

func (m *User) GetUpdatedTime() uint64 {
	if m != nil {
		return m.UpdatedTime
	}
	return 0
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Detail               string   `protobuf:"bytes,2,opt,name=detail,proto3" json:"detail,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{1}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetDetail() string {
	if m != nil {
		return m.Detail
	}
	return ""
}

type Request struct {
	UserID               string   `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	UserPwd              string   `protobuf:"bytes,3,opt,name=userPwd,proto3" json:"userPwd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{2}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *Request) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *Request) GetUserPwd() string {
	if m != nil {
		return m.UserPwd
	}
	return ""
}

type Response struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error                *Error   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	User                 *User    `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{3}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *Response) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *Response) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "smtl.micro.learn.srv.user.user")
	proto.RegisterType((*Error)(nil), "smtl.micro.learn.srv.user.Error")
	proto.RegisterType((*Request)(nil), "smtl.micro.learn.srv.user.Request")
	proto.RegisterType((*Response)(nil), "smtl.micro.learn.srv.user.Response")
}

func init() { proto.RegisterFile("proto/user/user.proto", fileDescriptor_9b283a848145d6b7) }

var fileDescriptor_9b283a848145d6b7 = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x3f, 0x4f, 0xf3, 0x30,
	0x10, 0xc6, 0x5f, 0xb7, 0x49, 0xff, 0x5c, 0xa5, 0x17, 0x64, 0x09, 0x14, 0xba, 0x10, 0x99, 0xa5,
	0x93, 0x91, 0x5a, 0x89, 0x0f, 0x80, 0x60, 0x60, 0x41, 0x60, 0x81, 0x58, 0x58, 0x42, 0x7c, 0x43,
	0xa4, 0xa6, 0x0e, 0x76, 0x02, 0xea, 0xc6, 0x57, 0xe0, 0x1b, 0xa3, 0xbb, 0xa4, 0xa5, 0x0b, 0x59,
	0xac, 0xe7, 0x39, 0xfd, 0x7c, 0xbe, 0xe7, 0x12, 0x38, 0xa9, 0xbc, 0xab, 0xdd, 0x65, 0x13, 0xd0,
	0xf3, 0xa1, 0xd9, 0xcb, 0xb3, 0x50, 0xd6, 0x6b, 0x5d, 0x16, 0xb9, 0x77, 0x7a, 0x8d, 0x99, 0xdf,
	0xe8, 0xe0, 0x3f, 0x34, 0x01, 0xea, 0x4b, 0x40, 0x44, 0x42, 0xfe, 0x87, 0x41, 0x61, 0x13, 0x91,
	0x8a, 0xc5, 0xd0, 0x0c, 0x0a, 0x2b, 0x25, 0x44, 0x9b, 0xac, 0xc4, 0x64, 0x90, 0x8a, 0xc5, 0xd4,
	0xb0, 0x96, 0xc7, 0x30, 0xac, 0x3e, 0x6d, 0x32, 0xe4, 0x12, 0x49, 0x99, 0xc2, 0x2c, 0xf7, 0x98,
	0xd5, 0x68, 0x9f, 0x8a, 0x12, 0x93, 0x28, 0x15, 0x8b, 0xc8, 0x1c, 0x96, 0x88, 0x68, 0x2a, 0xbb,
	0x27, 0xe2, 0x96, 0x38, 0x28, 0xa9, 0x15, 0xc4, 0xb7, 0xde, 0x3b, 0x4f, 0x4f, 0xe6, 0xce, 0x22,
	0x0f, 0x11, 0x1b, 0xd6, 0xf2, 0x14, 0x46, 0x16, 0xeb, 0xac, 0x58, 0x77, 0x83, 0x74, 0x4e, 0xbd,
	0xc0, 0xd8, 0xe0, 0x7b, 0x83, 0xa1, 0x26, 0x84, 0x12, 0xdc, 0xdd, 0xf0, 0xc5, 0xa9, 0xe9, 0x9c,
	0x9c, 0xc3, 0x84, 0xd4, 0xfd, 0x6f, 0x8a, 0xbd, 0x97, 0x09, 0x8c, 0x49, 0x3f, 0xec, 0xd3, 0xec,
	0xac, 0xfa, 0x16, 0x30, 0x31, 0x18, 0x2a, 0xb7, 0x09, 0x8c, 0x85, 0x26, 0xcf, 0x31, 0x04, 0xee,
	0x3d, 0x31, 0x3b, 0x2b, 0xaf, 0x20, 0x46, 0x1a, 0x9a, 0x3b, 0xcf, 0x96, 0xa9, 0xfe, 0x73, 0xc5,
	0x9a, 0xc3, 0x99, 0x16, 0x97, 0xab, 0x76, 0xdd, 0xfc, 0xea, 0x6c, 0x79, 0xde, 0x73, 0x8d, 0x0e,
	0xc3, 0xf0, 0xd2, 0x42, 0xf4, 0x4c, 0xdf, 0xe8, 0x15, 0x8e, 0x1e, 0x1b, 0xf4, 0x5b, 0x32, 0xd7,
	0x5b, 0x0e, 0xa2, 0x7a, 0x3a, 0x74, 0x0b, 0x9a, 0x5f, 0xf4, 0x32, 0x6d, 0x54, 0xf5, 0xef, 0x6d,
	0xc4, 0x3f, 0xcb, 0xea, 0x27, 0x00, 0x00, 0xff, 0xff, 0x75, 0xb3, 0xd0, 0xaa, 0x45, 0x02, 0x00,
	0x00,
}
