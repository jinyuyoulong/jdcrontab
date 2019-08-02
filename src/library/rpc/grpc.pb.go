// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc.proto

package rpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type GRPCRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GRPCRequest) Reset()         { *m = GRPCRequest{} }
func (m *GRPCRequest) String() string { return proto.CompactTextString(m) }
func (*GRPCRequest) ProtoMessage()    {}
func (*GRPCRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{0}
}

func (m *GRPCRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GRPCRequest.Unmarshal(m, b)
}
func (m *GRPCRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GRPCRequest.Marshal(b, m, deterministic)
}
func (m *GRPCRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GRPCRequest.Merge(m, src)
}
func (m *GRPCRequest) XXX_Size() int {
	return xxx_messageInfo_GRPCRequest.Size(m)
}
func (m *GRPCRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GRPCRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GRPCRequest proto.InternalMessageInfo

func (m *GRPCRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GRPCReply struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data                 string   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GRPCReply) Reset()         { *m = GRPCReply{} }
func (m *GRPCReply) String() string { return proto.CompactTextString(m) }
func (*GRPCReply) ProtoMessage()    {}
func (*GRPCReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{1}
}

func (m *GRPCReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GRPCReply.Unmarshal(m, b)
}
func (m *GRPCReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GRPCReply.Marshal(b, m, deterministic)
}
func (m *GRPCReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GRPCReply.Merge(m, src)
}
func (m *GRPCReply) XXX_Size() int {
	return xxx_messageInfo_GRPCReply.Size(m)
}
func (m *GRPCReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GRPCReply.DiscardUnknown(m)
}

var xxx_messageInfo_GRPCReply proto.InternalMessageInfo

func (m *GRPCReply) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GRPCReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *GRPCReply) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*GRPCRequest)(nil), "rpc.GRPCRequest")
	proto.RegisterType((*GRPCReply)(nil), "rpc.GRPCReply")
}

func init() { proto.RegisterFile("grpc.proto", fileDescriptor_bedfbfc9b54e5600) }

var fileDescriptor_bedfbfc9b54e5600 = []byte{
	// 168 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x2f, 0x2a, 0x48,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2e, 0x2a, 0x48, 0x56, 0x52, 0xe4, 0xe2, 0x76,
	0x0f, 0x0a, 0x70, 0x0e, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe2, 0x62, 0xc9, 0x4b,
	0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x95, 0x7c, 0xb9, 0x38, 0x21,
	0x4a, 0x0a, 0x72, 0x2a, 0x41, 0x0a, 0x92, 0xf3, 0x53, 0x20, 0x0a, 0x58, 0x83, 0xc0, 0x6c, 0x21,
	0x09, 0x2e, 0xf6, 0xdc, 0xd4, 0xe2, 0xe2, 0xc4, 0xf4, 0x54, 0x09, 0x26, 0xb0, 0x3e, 0x18, 0x17,
	0xa4, 0x3a, 0x25, 0xb1, 0x24, 0x51, 0x82, 0x19, 0x62, 0x1c, 0x88, 0x6d, 0xe4, 0xc8, 0xc5, 0x0d,
	0x72, 0x44, 0x70, 0x6a, 0x51, 0x59, 0x66, 0x72, 0xaa, 0x90, 0x11, 0x17, 0x0f, 0xc4, 0xf4, 0xe2,
	0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x01, 0x3d, 0x90, 0x0b, 0x91, 0xdc, 0x24, 0xc5, 0x87, 0x24,
	0x52, 0x90, 0x53, 0xa9, 0xc4, 0x90, 0xc4, 0x06, 0xf6, 0x80, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff,
	0xd5, 0x51, 0x2e, 0xc8, 0xce, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GrpcServiceClient is the client API for GrpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GrpcServiceClient interface {
	GRPCResponse(ctx context.Context, in *GRPCRequest, opts ...grpc.CallOption) (*GRPCReply, error)
}

type grpcServiceClient struct {
	cc *grpc.ClientConn
}

func NewGrpcServiceClient(cc *grpc.ClientConn) GrpcServiceClient {
	return &grpcServiceClient{cc}
}

func (c *grpcServiceClient) GRPCResponse(ctx context.Context, in *GRPCRequest, opts ...grpc.CallOption) (*GRPCReply, error) {
	out := new(GRPCReply)
	err := c.cc.Invoke(ctx, "/rpc.grpcService/GRPCResponse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GrpcServiceServer is the server API for GrpcService service.
type GrpcServiceServer interface {
	GRPCResponse(context.Context, *GRPCRequest) (*GRPCReply, error)
}

func RegisterGrpcServiceServer(s *grpc.Server, srv GrpcServiceServer) {
	s.RegisterService(&_GrpcService_serviceDesc, srv)
}

func _GrpcService_GRPCResponse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GRPCRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcServiceServer).GRPCResponse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.grpcService/GRPCResponse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcServiceServer).GRPCResponse(ctx, req.(*GRPCRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GrpcService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.grpcService",
	HandlerType: (*GrpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GRPCResponse",
			Handler:    _GrpcService_GRPCResponse_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc.proto",
}