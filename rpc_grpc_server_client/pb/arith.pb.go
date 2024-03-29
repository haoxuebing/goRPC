// Code generated by protoc-gen-go. DO NOT EDIT.
// source: arith.proto

package pb

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

// 算术运算请求结构
type ArithRequest struct {
	A                    int32    `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B                    int32    `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArithRequest) Reset()         { *m = ArithRequest{} }
func (m *ArithRequest) String() string { return proto.CompactTextString(m) }
func (*ArithRequest) ProtoMessage()    {}
func (*ArithRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_211289c5d02710c5, []int{0}
}

func (m *ArithRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArithRequest.Unmarshal(m, b)
}
func (m *ArithRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArithRequest.Marshal(b, m, deterministic)
}
func (m *ArithRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArithRequest.Merge(m, src)
}
func (m *ArithRequest) XXX_Size() int {
	return xxx_messageInfo_ArithRequest.Size(m)
}
func (m *ArithRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ArithRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ArithRequest proto.InternalMessageInfo

func (m *ArithRequest) GetA() int32 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *ArithRequest) GetB() int32 {
	if m != nil {
		return m.B
	}
	return 0
}

// 算术运算响应结构
type ArithResponse struct {
	Pro                  int32    `protobuf:"varint,1,opt,name=pro,proto3" json:"pro,omitempty"`
	Quo                  int32    `protobuf:"varint,2,opt,name=quo,proto3" json:"quo,omitempty"`
	Rem                  int32    `protobuf:"varint,3,opt,name=rem,proto3" json:"rem,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArithResponse) Reset()         { *m = ArithResponse{} }
func (m *ArithResponse) String() string { return proto.CompactTextString(m) }
func (*ArithResponse) ProtoMessage()    {}
func (*ArithResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_211289c5d02710c5, []int{1}
}

func (m *ArithResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArithResponse.Unmarshal(m, b)
}
func (m *ArithResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArithResponse.Marshal(b, m, deterministic)
}
func (m *ArithResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArithResponse.Merge(m, src)
}
func (m *ArithResponse) XXX_Size() int {
	return xxx_messageInfo_ArithResponse.Size(m)
}
func (m *ArithResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ArithResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ArithResponse proto.InternalMessageInfo

func (m *ArithResponse) GetPro() int32 {
	if m != nil {
		return m.Pro
	}
	return 0
}

func (m *ArithResponse) GetQuo() int32 {
	if m != nil {
		return m.Quo
	}
	return 0
}

func (m *ArithResponse) GetRem() int32 {
	if m != nil {
		return m.Rem
	}
	return 0
}

func init() {
	proto.RegisterType((*ArithRequest)(nil), "pb.ArithRequest")
	proto.RegisterType((*ArithResponse)(nil), "pb.ArithResponse")
}

func init() { proto.RegisterFile("arith.proto", fileDescriptor_211289c5d02710c5) }

var fileDescriptor_211289c5d02710c5 = []byte{
	// 174 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x2c, 0xca, 0x2c,
	0xc9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0xd2, 0xe2, 0xe2, 0x71,
	0x04, 0x09, 0x05, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0xf1, 0x70, 0x31, 0x26, 0x4a, 0x30,
	0x2a, 0x30, 0x6a, 0xb0, 0x06, 0x31, 0x26, 0x82, 0x78, 0x49, 0x12, 0x4c, 0x10, 0x5e, 0x92, 0x92,
	0x2b, 0x17, 0x2f, 0x54, 0x6d, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x00, 0x17, 0x73, 0x41,
	0x51, 0x3e, 0x54, 0x39, 0x88, 0x09, 0x12, 0x29, 0x2c, 0xcd, 0x87, 0x6a, 0x01, 0x31, 0x41, 0x22,
	0x45, 0xa9, 0xb9, 0x12, 0xcc, 0x10, 0x91, 0xa2, 0xd4, 0x5c, 0xa3, 0x3c, 0xa8, 0x95, 0xc1, 0xa9,
	0x45, 0x65, 0x99, 0xc9, 0xa9, 0x42, 0xfa, 0x5c, 0x1c, 0xb9, 0xa5, 0x39, 0x25, 0x99, 0x05, 0x39,
	0x95, 0x42, 0x02, 0x7a, 0x05, 0x49, 0x7a, 0xc8, 0x0e, 0x92, 0x12, 0x44, 0x12, 0x81, 0x5a, 0xab,
	0xcb, 0xc5, 0x96, 0x92, 0x59, 0x96, 0x99, 0x92, 0x4a, 0x94, 0xf2, 0x24, 0x36, 0xb0, 0x6f, 0x8d,
	0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xcb, 0x03, 0xde, 0xe9, 0xfc, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ArithServiceClient is the client API for ArithService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ArithServiceClient interface {
	Multiply(ctx context.Context, in *ArithRequest, opts ...grpc.CallOption) (*ArithResponse, error)
	Divide(ctx context.Context, in *ArithRequest, opts ...grpc.CallOption) (*ArithResponse, error)
}

type arithServiceClient struct {
	cc *grpc.ClientConn
}

func NewArithServiceClient(cc *grpc.ClientConn) ArithServiceClient {
	return &arithServiceClient{cc}
}

func (c *arithServiceClient) Multiply(ctx context.Context, in *ArithRequest, opts ...grpc.CallOption) (*ArithResponse, error) {
	out := new(ArithResponse)
	err := c.cc.Invoke(ctx, "/pb.ArithService/multiply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *arithServiceClient) Divide(ctx context.Context, in *ArithRequest, opts ...grpc.CallOption) (*ArithResponse, error) {
	out := new(ArithResponse)
	err := c.cc.Invoke(ctx, "/pb.ArithService/divide", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArithServiceServer is the server API for ArithService service.
type ArithServiceServer interface {
	Multiply(context.Context, *ArithRequest) (*ArithResponse, error)
	Divide(context.Context, *ArithRequest) (*ArithResponse, error)
}

func RegisterArithServiceServer(s *grpc.Server, srv ArithServiceServer) {
	s.RegisterService(&_ArithService_serviceDesc, srv)
}

func _ArithService_Multiply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArithRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArithServiceServer).Multiply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ArithService/Multiply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArithServiceServer).Multiply(ctx, req.(*ArithRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArithService_Divide_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArithRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArithServiceServer).Divide(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ArithService/Divide",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArithServiceServer).Divide(ctx, req.(*ArithRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ArithService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ArithService",
	HandlerType: (*ArithServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "multiply",
			Handler:    _ArithService_Multiply_Handler,
		},
		{
			MethodName: "divide",
			Handler:    _ArithService_Divide_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "arith.proto",
}
