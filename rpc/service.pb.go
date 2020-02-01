// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/service.proto

package rpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ProverbRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProverbRequest) Reset()         { *m = ProverbRequest{} }
func (m *ProverbRequest) String() string { return proto.CompactTextString(m) }
func (*ProverbRequest) ProtoMessage()    {}
func (*ProverbRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c33392ef2c1961ba, []int{0}
}

func (m *ProverbRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProverbRequest.Unmarshal(m, b)
}
func (m *ProverbRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProverbRequest.Marshal(b, m, deterministic)
}
func (m *ProverbRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProverbRequest.Merge(m, src)
}
func (m *ProverbRequest) XXX_Size() int {
	return xxx_messageInfo_ProverbRequest.Size(m)
}
func (m *ProverbRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProverbRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProverbRequest proto.InternalMessageInfo

type ProverbResponse struct {
	Proverb              string   `protobuf:"bytes,1,opt,name=proverb,proto3" json:"proverb,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProverbResponse) Reset()         { *m = ProverbResponse{} }
func (m *ProverbResponse) String() string { return proto.CompactTextString(m) }
func (*ProverbResponse) ProtoMessage()    {}
func (*ProverbResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c33392ef2c1961ba, []int{1}
}

func (m *ProverbResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProverbResponse.Unmarshal(m, b)
}
func (m *ProverbResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProverbResponse.Marshal(b, m, deterministic)
}
func (m *ProverbResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProverbResponse.Merge(m, src)
}
func (m *ProverbResponse) XXX_Size() int {
	return xxx_messageInfo_ProverbResponse.Size(m)
}
func (m *ProverbResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProverbResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProverbResponse proto.InternalMessageInfo

func (m *ProverbResponse) GetProverb() string {
	if m != nil {
		return m.Proverb
	}
	return ""
}

func init() {
	proto.RegisterType((*ProverbRequest)(nil), "rpc.ProverbRequest")
	proto.RegisterType((*ProverbResponse)(nil), "rpc.ProverbResponse")
}

func init() { proto.RegisterFile("proto/service.proto", fileDescriptor_c33392ef2c1961ba) }

var fileDescriptor_c33392ef2c1961ba = []byte{
	// 169 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x03, 0xf3, 0x84, 0x98, 0x8b, 0x0a,
	0x92, 0xa5, 0x64, 0xd2, 0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5, 0x13, 0x0b, 0x32, 0xf5, 0x13, 0xf3,
	0xf2, 0xf2, 0x4b, 0x12, 0x4b, 0x32, 0xf3, 0xf3, 0x8a, 0x21, 0x4a, 0x94, 0x04, 0xb8, 0xf8, 0x02,
	0x8a, 0xf2, 0xcb, 0x52, 0x8b, 0x92, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x94, 0xb4, 0xb9,
	0xf8, 0xe1, 0x22, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x12, 0x5c, 0xec, 0x05, 0x10, 0x21,
	0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x18, 0xd7, 0x28, 0x8c, 0x8b, 0x13, 0xaa, 0x38, 0xb5,
	0x48, 0xc8, 0x93, 0x8b, 0x1d, 0xca, 0x11, 0x12, 0xd6, 0x2b, 0x2a, 0x48, 0xd6, 0x43, 0x35, 0x59,
	0x4a, 0x04, 0x55, 0x10, 0x62, 0xb8, 0x92, 0x70, 0xd3, 0xe5, 0x27, 0x93, 0x99, 0x78, 0x95, 0x38,
	0xf4, 0xa1, 0x86, 0x5a, 0x31, 0x6a, 0x25, 0xb1, 0x81, 0x5d, 0x67, 0x0c, 0x08, 0x00, 0x00, 0xff,
	0xff, 0x3c, 0x5f, 0x91, 0x9b, 0xd7, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProverberClient is the client API for Proverber service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProverberClient interface {
	// Proverb returns a Go proverb.
	Proverb(ctx context.Context, in *ProverbRequest, opts ...grpc.CallOption) (*ProverbResponse, error)
}

type proverberClient struct {
	cc grpc.ClientConnInterface
}

func NewProverberClient(cc grpc.ClientConnInterface) ProverberClient {
	return &proverberClient{cc}
}

func (c *proverberClient) Proverb(ctx context.Context, in *ProverbRequest, opts ...grpc.CallOption) (*ProverbResponse, error) {
	out := new(ProverbResponse)
	err := c.cc.Invoke(ctx, "/rpc.Proverber/Proverb", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProverberServer is the server API for Proverber service.
type ProverberServer interface {
	// Proverb returns a Go proverb.
	Proverb(context.Context, *ProverbRequest) (*ProverbResponse, error)
}

// UnimplementedProverberServer can be embedded to have forward compatible implementations.
type UnimplementedProverberServer struct {
}

func (*UnimplementedProverberServer) Proverb(ctx context.Context, req *ProverbRequest) (*ProverbResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Proverb not implemented")
}

func RegisterProverberServer(s *grpc.Server, srv ProverberServer) {
	s.RegisterService(&_Proverber_serviceDesc, srv)
}

func _Proverber_Proverb_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProverbRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProverberServer).Proverb(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Proverber/Proverb",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProverberServer).Proverb(ctx, req.(*ProverbRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Proverber_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.Proverber",
	HandlerType: (*ProverberServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Proverb",
			Handler:    _Proverber_Proverb_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/service.proto",
}
