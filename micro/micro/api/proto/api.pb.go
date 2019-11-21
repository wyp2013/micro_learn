// Code generated by protoc-gen-go. DO NOT EDIT.
// source: micro_learn/micro/micro/api/proto/api.proto

package go_micro_api

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

type Pair struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Values               []string `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pair) Reset()         { *m = Pair{} }
func (m *Pair) String() string { return proto.CompactTextString(m) }
func (*Pair) ProtoMessage()    {}
func (*Pair) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_f8ef31c2e6dafc0b, []int{0}
}
func (m *Pair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pair.Unmarshal(m, b)
}
func (m *Pair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pair.Marshal(b, m, deterministic)
}
func (dst *Pair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pair.Merge(dst, src)
}
func (m *Pair) XXX_Size() int {
	return xxx_messageInfo_Pair.Size(m)
}
func (m *Pair) XXX_DiscardUnknown() {
	xxx_messageInfo_Pair.DiscardUnknown(m)
}

var xxx_messageInfo_Pair proto.InternalMessageInfo

func (m *Pair) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Pair) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

type Request struct {
	Method               string           `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Path                 string           `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Header               map[string]*Pair `protobuf:"bytes,3,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Get                  map[string]*Pair `protobuf:"bytes,4,rep,name=get,proto3" json:"get,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Post                 map[string]*Pair `protobuf:"bytes,5,rep,name=post,proto3" json:"post,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Body                 string           `protobuf:"bytes,6,opt,name=body,proto3" json:"body,omitempty"`
	Url                  string           `protobuf:"bytes,7,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_f8ef31c2e6dafc0b, []int{1}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *Request) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Request) GetHeader() map[string]*Pair {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Request) GetGet() map[string]*Pair {
	if m != nil {
		return m.Get
	}
	return nil
}

func (m *Request) GetPost() map[string]*Pair {
	if m != nil {
		return m.Post
	}
	return nil
}

func (m *Request) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *Request) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type Response struct {
	StatusCode           int32            `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	Header               map[string]*Pair `protobuf:"bytes,2,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Body                 string           `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_f8ef31c2e6dafc0b, []int{2}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *Response) GetHeader() map[string]*Pair {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Response) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func init() {
	proto.RegisterType((*Pair)(nil), "go.micro.api.Pair")
	proto.RegisterType((*Request)(nil), "go.micro.api.Request")
	proto.RegisterMapType((map[string]*Pair)(nil), "go.micro.api.Request.GetEntry")
	proto.RegisterMapType((map[string]*Pair)(nil), "go.micro.api.Request.HeaderEntry")
	proto.RegisterMapType((map[string]*Pair)(nil), "go.micro.api.Request.PostEntry")
	proto.RegisterType((*Response)(nil), "go.micro.api.Response")
	proto.RegisterMapType((map[string]*Pair)(nil), "go.micro.api.Response.HeaderEntry")
}

func init() {
	proto.RegisterFile("micro_learn/micro/micro/api/proto/api.proto", fileDescriptor_api_f8ef31c2e6dafc0b)
}

var fileDescriptor_api_f8ef31c2e6dafc0b = []byte{
	// 347 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xcd, 0x6e, 0xe2, 0x30,
	0x14, 0x85, 0x95, 0x1f, 0x02, 0x5c, 0x66, 0x31, 0xf2, 0x62, 0x64, 0xb1, 0x60, 0x98, 0xac, 0xa2,
	0x59, 0x18, 0x04, 0x9b, 0x96, 0x6d, 0x55, 0xb5, 0x6a, 0x55, 0x09, 0xf9, 0x0d, 0x0c, 0xb1, 0x48,
	0x54, 0xc0, 0xa9, 0xed, 0x54, 0xe2, 0x11, 0xfb, 0x20, 0x7d, 0x8f, 0xca, 0x37, 0x86, 0xd2, 0x2a,
	0x3b, 0xba, 0x89, 0xae, 0x9d, 0x73, 0x4e, 0x8e, 0x3f, 0x07, 0xfe, 0x6f, 0x4a, 0x5b, 0xd4, 0x2b,
	0xb6, 0x56, 0xbb, 0xc9, 0xae, 0x5c, 0x6b, 0xe5, 0x9f, 0xa2, 0x2a, 0x27, 0x95, 0x56, 0x16, 0x27,
	0x86, 0x13, 0xf9, 0xb5, 0x51, 0x0c, 0xdf, 0x32, 0x51, 0x95, 0xe9, 0x14, 0xe2, 0xa5, 0x28, 0x35,
	0xf9, 0x0d, 0xd1, 0xb3, 0x3c, 0xd0, 0x60, 0x1c, 0x64, 0x7d, 0xee, 0x46, 0xf2, 0x07, 0x92, 0x57,
	0xb1, 0xad, 0xa5, 0xa1, 0xe1, 0x38, 0xca, 0xfa, 0xdc, 0xaf, 0xd2, 0xf7, 0x08, 0xba, 0x5c, 0xbe,
	0xd4, 0xd2, 0x58, 0xa7, 0xd9, 0x49, 0x5b, 0xa8, 0xdc, 0x1b, 0xfd, 0x8a, 0x10, 0x88, 0x2b, 0x61,
	0x0b, 0x1a, 0xe2, 0x2e, 0xce, 0xe4, 0x1a, 0x92, 0x42, 0x8a, 0x5c, 0x6a, 0x1a, 0x8d, 0xa3, 0x6c,
	0x30, 0xfb, 0xc7, 0xce, 0x8b, 0x30, 0x1f, 0xc9, 0xee, 0x51, 0x73, 0xbb, 0xb7, 0xfa, 0xc0, 0xbd,
	0x81, 0x4c, 0x21, 0xda, 0x48, 0x4b, 0x63, 0xf4, 0x8d, 0xda, 0x7d, 0x77, 0xd2, 0x36, 0x26, 0x27,
	0x25, 0x73, 0x88, 0x2b, 0x65, 0x2c, 0xed, 0xa0, 0xe5, 0x6f, 0xbb, 0x65, 0xa9, 0x8c, 0xf7, 0xa0,
	0xd8, 0xb5, 0x5e, 0xa9, 0xfc, 0x40, 0x93, 0xa6, 0xb5, 0x9b, 0x1d, 0x97, 0x5a, 0x6f, 0x69, 0xb7,
	0xe1, 0x52, 0xeb, 0xed, 0xf0, 0x09, 0x06, 0x67, 0x1d, 0x5b, 0xc0, 0x65, 0xd0, 0x41, 0x54, 0x78,
	0xfa, 0xc1, 0x8c, 0x7c, 0xfd, 0xb8, 0xa3, 0xcd, 0x1b, 0xc1, 0x22, 0xbc, 0x0a, 0x86, 0x0f, 0xd0,
	0x3b, 0x56, 0xbf, 0x38, 0xeb, 0x11, 0xfa, 0xa7, 0x33, 0x5d, 0x1a, 0x96, 0xbe, 0x05, 0xd0, 0xe3,
	0xd2, 0x54, 0x6a, 0x6f, 0x24, 0x19, 0x01, 0x18, 0x2b, 0x6c, 0x6d, 0x6e, 0x54, 0x2e, 0x31, 0xb3,
	0xc3, 0xcf, 0x76, 0xc8, 0xe2, 0x74, 0xb9, 0x21, 0x12, 0x4f, 0xbf, 0x13, 0x6f, 0x72, 0x5a, 0x6f,
	0xf7, 0x88, 0x3d, 0xfa, 0xc4, 0xfe, 0xc3, 0x90, 0x57, 0x09, 0xfe, 0xfa, 0xf3, 0x8f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x7b, 0x71, 0xca, 0x17, 0x28, 0x03, 0x00, 0x00,
}
