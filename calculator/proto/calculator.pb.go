// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calculator.proto

package scenthound

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Operations int32

const (
	Operations_ADD      Operations = 0
	Operations_SUBTRACT Operations = 1
	Operations_MULTIPLY Operations = 2
	Operations_DIVIDE   Operations = 3
)

var Operations_name = map[int32]string{
	0: "ADD",
	1: "SUBTRACT",
	2: "MULTIPLY",
	3: "DIVIDE",
}

var Operations_value = map[string]int32{
	"ADD":      0,
	"SUBTRACT": 1,
	"MULTIPLY": 2,
	"DIVIDE":   3,
}

func (x Operations) String() string {
	return proto.EnumName(Operations_name, int32(x))
}

func (Operations) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c686ea360062a8cf, []int{0}
}

type Operands struct {
	Number1              float64  `protobuf:"fixed64,1,opt,name=number1,proto3" json:"number1,omitempty"`
	Number2              float64  `protobuf:"fixed64,2,opt,name=number2,proto3" json:"number2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Operands) Reset()         { *m = Operands{} }
func (m *Operands) String() string { return proto.CompactTextString(m) }
func (*Operands) ProtoMessage()    {}
func (*Operands) Descriptor() ([]byte, []int) {
	return fileDescriptor_c686ea360062a8cf, []int{0}
}

func (m *Operands) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Operands.Unmarshal(m, b)
}
func (m *Operands) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Operands.Marshal(b, m, deterministic)
}
func (m *Operands) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Operands.Merge(m, src)
}
func (m *Operands) XXX_Size() int {
	return xxx_messageInfo_Operands.Size(m)
}
func (m *Operands) XXX_DiscardUnknown() {
	xxx_messageInfo_Operands.DiscardUnknown(m)
}

var xxx_messageInfo_Operands proto.InternalMessageInfo

func (m *Operands) GetNumber1() float64 {
	if m != nil {
		return m.Number1
	}
	return 0
}

func (m *Operands) GetNumber2() float64 {
	if m != nil {
		return m.Number2
	}
	return 0
}

type Expression struct {
	Operation            Operations `protobuf:"varint,1,opt,name=operation,proto3,enum=scenthound.Operations" json:"operation,omitempty"`
	Operands             *Operands  `protobuf:"bytes,2,opt,name=operands,proto3" json:"operands,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Expression) Reset()         { *m = Expression{} }
func (m *Expression) String() string { return proto.CompactTextString(m) }
func (*Expression) ProtoMessage()    {}
func (*Expression) Descriptor() ([]byte, []int) {
	return fileDescriptor_c686ea360062a8cf, []int{1}
}

func (m *Expression) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Expression.Unmarshal(m, b)
}
func (m *Expression) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Expression.Marshal(b, m, deterministic)
}
func (m *Expression) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Expression.Merge(m, src)
}
func (m *Expression) XXX_Size() int {
	return xxx_messageInfo_Expression.Size(m)
}
func (m *Expression) XXX_DiscardUnknown() {
	xxx_messageInfo_Expression.DiscardUnknown(m)
}

var xxx_messageInfo_Expression proto.InternalMessageInfo

func (m *Expression) GetOperation() Operations {
	if m != nil {
		return m.Operation
	}
	return Operations_ADD
}

func (m *Expression) GetOperands() *Operands {
	if m != nil {
		return m.Operands
	}
	return nil
}

type Response struct {
	Expression           *Expression `protobuf:"bytes,1,opt,name=expression,proto3" json:"expression,omitempty"`
	Result               float64     `protobuf:"fixed64,2,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_c686ea360062a8cf, []int{2}
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

func (m *Response) GetExpression() *Expression {
	if m != nil {
		return m.Expression
	}
	return nil
}

func (m *Response) GetResult() float64 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterEnum("scenthound.Operations", Operations_name, Operations_value)
	proto.RegisterType((*Operands)(nil), "scenthound.Operands")
	proto.RegisterType((*Expression)(nil), "scenthound.Expression")
	proto.RegisterType((*Response)(nil), "scenthound.Response")
}

func init() {
	proto.RegisterFile("calculator.proto", fileDescriptor_c686ea360062a8cf)
}

var fileDescriptor_c686ea360062a8cf = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x51, 0x4b, 0x4b, 0xc3, 0x40,
	0x18, 0x34, 0x2d, 0xb4, 0x71, 0x2a, 0x12, 0x3e, 0xa4, 0x04, 0x4f, 0x92, 0x93, 0x78, 0x08, 0x1a,
	0x45, 0xbc, 0x28, 0xd4, 0x26, 0x42, 0xa0, 0x52, 0x59, 0x53, 0x41, 0x6f, 0x69, 0xb2, 0x60, 0x21,
	0xee, 0x86, 0xdd, 0x8d, 0xf8, 0xf3, 0xa5, 0x21, 0x8f, 0x16, 0xf1, 0x38, 0xcc, 0xec, 0x3c, 0xf6,
	0x83, 0x93, 0xa5, 0x45, 0x56, 0x15, 0xa9, 0x91, 0xca, 0x2f, 0x95, 0x34, 0x92, 0xa0, 0x33, 0x2e,
	0xcc, 0xa7, 0xac, 0x44, 0xee, 0x3d, 0xc0, 0x5e, 0x96, 0x5c, 0xa5, 0x22, 0xd7, 0xe4, 0x62, 0x2c,
	0xaa, 0xaf, 0x35, 0x57, 0x57, 0xae, 0x75, 0x66, 0x9d, 0x5b, 0xac, 0x85, 0x3d, 0x13, 0xb8, 0x83,
	0x5d, 0x26, 0xf0, 0x0c, 0x10, 0xfd, 0x94, 0x8a, 0x6b, 0xbd, 0x91, 0x82, 0x6e, 0x70, 0x28, 0xb7,
	0x6e, 0x66, 0x23, 0x45, 0xed, 0x71, 0x1c, 0x4c, 0xfd, 0x3e, 0xcd, 0x5f, 0xb6, 0xa4, 0x66, 0xbd,
	0x90, 0x2e, 0x61, 0xcb, 0xa6, 0x43, 0x6d, 0x3f, 0x09, 0x4e, 0xfe, 0x3c, 0x12, 0xb9, 0x66, 0x9d,
	0xca, 0xfb, 0x80, 0xcd, 0xb8, 0x2e, 0xa5, 0xd0, 0x9c, 0x6e, 0x01, 0xde, 0x35, 0xa8, 0x43, 0x27,
	0xfb, 0xa1, 0x7d, 0x3f, 0xb6, 0xa3, 0xa4, 0x29, 0x46, 0x8a, 0xeb, 0xaa, 0x30, 0xcd, 0xa4, 0x06,
	0x5d, 0xdc, 0x03, 0x7d, 0x4d, 0x1a, 0x63, 0x38, 0x0b, 0x43, 0xe7, 0x80, 0x8e, 0x60, 0xbf, 0xae,
	0x1e, 0x13, 0x36, 0x9b, 0x27, 0x8e, 0xb5, 0x45, 0xcf, 0xab, 0x45, 0x12, 0xbf, 0x2c, 0xde, 0x9d,
	0x01, 0x01, 0xa3, 0x30, 0x7e, 0x8b, 0xc3, 0xc8, 0x19, 0x06, 0x4f, 0xc0, 0xbc, 0xfb, 0x70, 0xba,
	0x83, 0x1d, 0x7d, 0xa7, 0x45, 0x95, 0x1a, 0x4e, 0xff, 0x94, 0x3a, 0xdd, 0x1b, 0xdb, 0xce, 0x5a,
	0x8f, 0xea, 0x5b, 0x5d, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x83, 0x43, 0xac, 0x3a, 0xbf, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CalculatorClient is the client API for Calculator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculatorClient interface {
	Evaluate(ctx context.Context, in *Expression, opts ...grpc.CallOption) (*Response, error)
}

type calculatorClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorClient(cc grpc.ClientConnInterface) CalculatorClient {
	return &calculatorClient{cc}
}

func (c *calculatorClient) Evaluate(ctx context.Context, in *Expression, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/scenthound.Calculator/Evaluate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculatorServer is the server API for Calculator service.
type CalculatorServer interface {
	Evaluate(context.Context, *Expression) (*Response, error)
}

// UnimplementedCalculatorServer can be embedded to have forward compatible implementations.
type UnimplementedCalculatorServer struct {
}

func (*UnimplementedCalculatorServer) Evaluate(ctx context.Context, req *Expression) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Evaluate not implemented")
}

func RegisterCalculatorServer(s *grpc.Server, srv CalculatorServer) {
	s.RegisterService(&_Calculator_serviceDesc, srv)
}

func _Calculator_Evaluate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Expression)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).Evaluate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scenthound.Calculator/Evaluate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).Evaluate(ctx, req.(*Expression))
	}
	return interceptor(ctx, in, info, handler)
}

var _Calculator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "scenthound.Calculator",
	HandlerType: (*CalculatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Evaluate",
			Handler:    _Calculator_Evaluate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calculator.proto",
}
