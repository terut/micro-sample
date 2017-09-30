// Code generated by protoc-gen-go. DO NOT EDIT.
// source: word.proto

/*
Package word is a generated protocol buffer package.

It is generated from these files:
	word.proto

It has these top-level messages:
	WordRequest
	WordResponse
*/
package word

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

type WordRequest struct {
	Val string `protobuf:"bytes,1,opt,name=val" json:"val,omitempty"`
}

func (m *WordRequest) Reset()                    { *m = WordRequest{} }
func (m *WordRequest) String() string            { return proto.CompactTextString(m) }
func (*WordRequest) ProtoMessage()               {}
func (*WordRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *WordRequest) GetVal() string {
	if m != nil {
		return m.Val
	}
	return ""
}

type WordResponse struct {
	Val string `protobuf:"bytes,1,opt,name=val" json:"val,omitempty"`
}

func (m *WordResponse) Reset()                    { *m = WordResponse{} }
func (m *WordResponse) String() string            { return proto.CompactTextString(m) }
func (*WordResponse) ProtoMessage()               {}
func (*WordResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *WordResponse) GetVal() string {
	if m != nil {
		return m.Val
	}
	return ""
}

func init() {
	proto.RegisterType((*WordRequest)(nil), "WordRequest")
	proto.RegisterType((*WordResponse)(nil), "WordResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Word service

type WordClient interface {
	Build(ctx context.Context, in *WordRequest, opts ...grpc.CallOption) (*WordResponse, error)
}

type wordClient struct {
	cc *grpc.ClientConn
}

func NewWordClient(cc *grpc.ClientConn) WordClient {
	return &wordClient{cc}
}

func (c *wordClient) Build(ctx context.Context, in *WordRequest, opts ...grpc.CallOption) (*WordResponse, error) {
	out := new(WordResponse)
	err := grpc.Invoke(ctx, "/Word/Build", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Word service

type WordServer interface {
	Build(context.Context, *WordRequest) (*WordResponse, error)
}

func RegisterWordServer(s *grpc.Server, srv WordServer) {
	s.RegisterService(&_Word_serviceDesc, srv)
}

func _Word_Build_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WordServer).Build(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Word/Build",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WordServer).Build(ctx, req.(*WordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Word_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Word",
	HandlerType: (*WordServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Build",
			Handler:    _Word_Build_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "word.proto",
}

func init() { proto.RegisterFile("word.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 113 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0xcf, 0x2f, 0x4a,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x92, 0xe7, 0xe2, 0x0e, 0xcf, 0x2f, 0x4a, 0x09, 0x4a,
	0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe0, 0x62, 0x2e, 0x4b, 0xcc, 0x91, 0x60, 0x54, 0x60,
	0xd4, 0xe0, 0x0c, 0x02, 0x31, 0x95, 0x14, 0xb8, 0x78, 0x20, 0x0a, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a,
	0x53, 0x31, 0x55, 0x18, 0xe9, 0x71, 0xb1, 0x80, 0x54, 0x08, 0xa9, 0x71, 0xb1, 0x3a, 0x95, 0x66,
	0xe6, 0xa4, 0x08, 0xf1, 0xe8, 0x21, 0x19, 0x29, 0xc5, 0xab, 0x87, 0xac, 0x5f, 0x89, 0x21, 0x89,
	0x0d, 0x6c, 0xb3, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xa7, 0xac, 0x21, 0xfb, 0x87, 0x00, 0x00,
	0x00,
}