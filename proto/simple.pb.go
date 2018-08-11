// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/simple.proto

package simple

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_simple_53ea4365f90f8929, []int{0}
}
func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (dst *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(dst, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HelloReplay struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloReplay) Reset()         { *m = HelloReplay{} }
func (m *HelloReplay) String() string { return proto.CompactTextString(m) }
func (*HelloReplay) ProtoMessage()    {}
func (*HelloReplay) Descriptor() ([]byte, []int) {
	return fileDescriptor_simple_53ea4365f90f8929, []int{1}
}
func (m *HelloReplay) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReplay.Unmarshal(m, b)
}
func (m *HelloReplay) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReplay.Marshal(b, m, deterministic)
}
func (dst *HelloReplay) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReplay.Merge(dst, src)
}
func (m *HelloReplay) XXX_Size() int {
	return xxx_messageInfo_HelloReplay.Size(m)
}
func (m *HelloReplay) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReplay.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReplay proto.InternalMessageInfo

func (m *HelloReplay) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "HelloRequest")
	proto.RegisterType((*HelloReplay)(nil), "HelloReplay")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SimpleClient is the client API for Simple service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SimpleClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReplay, error)
}

type simpleClient struct {
	cc *grpc.ClientConn
}

func NewSimpleClient(cc *grpc.ClientConn) SimpleClient {
	return &simpleClient{cc}
}

func (c *simpleClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReplay, error) {
	out := new(HelloReplay)
	err := c.cc.Invoke(ctx, "/Simple/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SimpleServer is the server API for Simple service.
type SimpleServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReplay, error)
}

func RegisterSimpleServer(s *grpc.Server, srv SimpleServer) {
	s.RegisterService(&_Simple_serviceDesc, srv)
}

func _Simple_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimpleServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Simple/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimpleServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Simple_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Simple",
	HandlerType: (*SimpleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Simple_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/simple.proto",
}

func init() { proto.RegisterFile("proto/simple.proto", fileDescriptor_simple_53ea4365f90f8929) }

var fileDescriptor_simple_53ea4365f90f8929 = []byte{
	// 131 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0xce, 0xcc, 0x2d, 0xc8, 0x49, 0xd5, 0x03, 0x73, 0x94, 0x94, 0xb8, 0x78, 0x3c,
	0x52, 0x73, 0x72, 0xf2, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x84, 0xb8, 0x58, 0xf2,
	0x12, 0x73, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x25, 0x75, 0x2e, 0x6e,
	0xa8, 0x9a, 0x82, 0x9c, 0xc4, 0x4a, 0x21, 0x09, 0x2e, 0xf6, 0xdc, 0xd4, 0xe2, 0xe2, 0xc4, 0x74,
	0x98, 0x2a, 0x18, 0xd7, 0xc8, 0x98, 0x8b, 0x2d, 0x18, 0x6c, 0xb8, 0x90, 0x26, 0x17, 0x47, 0x70,
	0x62, 0x25, 0x58, 0x97, 0x10, 0xaf, 0x1e, 0xb2, 0x0d, 0x52, 0x3c, 0x7a, 0x48, 0x86, 0x29, 0x31,
	0x24, 0xb1, 0x81, 0x1d, 0x62, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xec, 0x52, 0xf7, 0x5b, 0x9e,
	0x00, 0x00, 0x00,
}
