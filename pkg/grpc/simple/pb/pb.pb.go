// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb.proto

package pb

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

type StatusRequest struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusRequest) Reset()         { *m = StatusRequest{} }
func (m *StatusRequest) String() string { return proto.CompactTextString(m) }
func (*StatusRequest) ProtoMessage()    {}
func (*StatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_pb_1a57c384b9bfd9e2, []int{0}
}
func (m *StatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusRequest.Unmarshal(m, b)
}
func (m *StatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusRequest.Marshal(b, m, deterministic)
}
func (dst *StatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusRequest.Merge(dst, src)
}
func (m *StatusRequest) XXX_Size() int {
	return xxx_messageInfo_StatusRequest.Size(m)
}
func (m *StatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StatusRequest proto.InternalMessageInfo

func (m *StatusRequest) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type StatusReply struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusReply) Reset()         { *m = StatusReply{} }
func (m *StatusReply) String() string { return proto.CompactTextString(m) }
func (*StatusReply) ProtoMessage()    {}
func (*StatusReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_pb_1a57c384b9bfd9e2, []int{1}
}
func (m *StatusReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusReply.Unmarshal(m, b)
}
func (m *StatusReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusReply.Marshal(b, m, deterministic)
}
func (dst *StatusReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusReply.Merge(dst, src)
}
func (m *StatusReply) XXX_Size() int {
	return xxx_messageInfo_StatusReply.Size(m)
}
func (m *StatusReply) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusReply.DiscardUnknown(m)
}

var xxx_messageInfo_StatusReply proto.InternalMessageInfo

func (m *StatusReply) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*StatusRequest)(nil), "pb.StatusRequest")
	proto.RegisterType((*StatusReply)(nil), "pb.StatusReply")
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
	Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusReply, error)
}

type simpleClient struct {
	cc *grpc.ClientConn
}

func NewSimpleClient(cc *grpc.ClientConn) SimpleClient {
	return &simpleClient{cc}
}

func (c *simpleClient) Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusReply, error) {
	out := new(StatusReply)
	err := c.cc.Invoke(ctx, "/pb.Simple/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SimpleServer is the server API for Simple service.
type SimpleServer interface {
	Status(context.Context, *StatusRequest) (*StatusReply, error)
}

func RegisterSimpleServer(s *grpc.Server, srv SimpleServer) {
	s.RegisterService(&_Simple_serviceDesc, srv)
}

func _Simple_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimpleServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Simple/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimpleServer).Status(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Simple_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Simple",
	HandlerType: (*SimpleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _Simple_Status_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb.proto",
}

func init() { proto.RegisterFile("pb.proto", fileDescriptor_pb_1a57c384b9bfd9e2) }

var fileDescriptor_pb_1a57c384b9bfd9e2 = []byte{
	// 163 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x28, 0x48, 0xd2, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x52, 0xe6, 0xe2, 0x0d, 0x2e, 0x49, 0x2c,
	0x29, 0x2d, 0x0e, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe2, 0x62, 0x49, 0x49, 0x2c,
	0x49, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x95, 0x14, 0xb9, 0xb8, 0x61, 0x8a,
	0x0a, 0x72, 0x2a, 0xb1, 0x29, 0x31, 0xb2, 0xe0, 0x62, 0x0b, 0xce, 0xcc, 0x2d, 0xc8, 0x49, 0x15,
	0xd2, 0xe3, 0x62, 0x83, 0x28, 0x16, 0x12, 0xd4, 0x2b, 0x48, 0xd2, 0x43, 0x31, 0x5d, 0x8a, 0x1f,
	0x59, 0xa8, 0x20, 0xa7, 0x52, 0x89, 0xc1, 0xc9, 0x80, 0x4b, 0x3a, 0x33, 0x5f, 0x2f, 0xbd, 0xa8,
	0x20, 0x59, 0x2f, 0xb5, 0x22, 0x11, 0x64, 0x44, 0xb1, 0x5e, 0x51, 0x7e, 0x69, 0x49, 0x6a, 0x7a,
	0x69, 0x66, 0x4a, 0xaa, 0x13, 0x7f, 0x10, 0x88, 0xed, 0x0e, 0x62, 0x07, 0x80, 0x5c, 0x1d, 0xc0,
	0x98, 0xc4, 0x06, 0x76, 0xbe, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x70, 0x09, 0x97, 0x54, 0xca,
	0x00, 0x00, 0x00,
}