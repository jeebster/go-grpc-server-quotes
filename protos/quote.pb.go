// Code generated by protoc-gen-go. DO NOT EDIT.
// source: quote.proto

package quote

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

type QuotesRequest struct {
	Limit                int32    `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QuotesRequest) Reset()         { *m = QuotesRequest{} }
func (m *QuotesRequest) String() string { return proto.CompactTextString(m) }
func (*QuotesRequest) ProtoMessage()    {}
func (*QuotesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_12e4b46109c2c385, []int{0}
}

func (m *QuotesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuotesRequest.Unmarshal(m, b)
}
func (m *QuotesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuotesRequest.Marshal(b, m, deterministic)
}
func (m *QuotesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuotesRequest.Merge(m, src)
}
func (m *QuotesRequest) XXX_Size() int {
	return xxx_messageInfo_QuotesRequest.Size(m)
}
func (m *QuotesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QuotesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QuotesRequest proto.InternalMessageInfo

func (m *QuotesRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type QuotesResponse struct {
	Quotes               []*Quote `protobuf:"bytes,1,rep,name=quotes,proto3" json:"quotes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QuotesResponse) Reset()         { *m = QuotesResponse{} }
func (m *QuotesResponse) String() string { return proto.CompactTextString(m) }
func (*QuotesResponse) ProtoMessage()    {}
func (*QuotesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_12e4b46109c2c385, []int{1}
}

func (m *QuotesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuotesResponse.Unmarshal(m, b)
}
func (m *QuotesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuotesResponse.Marshal(b, m, deterministic)
}
func (m *QuotesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuotesResponse.Merge(m, src)
}
func (m *QuotesResponse) XXX_Size() int {
	return xxx_messageInfo_QuotesResponse.Size(m)
}
func (m *QuotesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QuotesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QuotesResponse proto.InternalMessageInfo

func (m *QuotesResponse) GetQuotes() []*Quote {
	if m != nil {
		return m.Quotes
	}
	return nil
}

type Quote struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Body                 string   `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	CreatedAt            string   `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Quote) Reset()         { *m = Quote{} }
func (m *Quote) String() string { return proto.CompactTextString(m) }
func (*Quote) ProtoMessage()    {}
func (*Quote) Descriptor() ([]byte, []int) {
	return fileDescriptor_12e4b46109c2c385, []int{2}
}

func (m *Quote) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Quote.Unmarshal(m, b)
}
func (m *Quote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Quote.Marshal(b, m, deterministic)
}
func (m *Quote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Quote.Merge(m, src)
}
func (m *Quote) XXX_Size() int {
	return xxx_messageInfo_Quote.Size(m)
}
func (m *Quote) XXX_DiscardUnknown() {
	xxx_messageInfo_Quote.DiscardUnknown(m)
}

var xxx_messageInfo_Quote proto.InternalMessageInfo

func (m *Quote) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Quote) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *Quote) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Quote) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func init() {
	proto.RegisterType((*QuotesRequest)(nil), "QuotesRequest")
	proto.RegisterType((*QuotesResponse)(nil), "QuotesResponse")
	proto.RegisterType((*Quote)(nil), "Quote")
}

func init() {
	proto.RegisterFile("quote.proto", fileDescriptor_12e4b46109c2c385)
}

var fileDescriptor_12e4b46109c2c385 = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x31, 0x4b, 0xc6, 0x30,
	0x10, 0x86, 0x49, 0xbf, 0xaf, 0x85, 0xde, 0xa7, 0x15, 0x0e, 0x87, 0x20, 0x28, 0xa5, 0x20, 0x74,
	0x90, 0x20, 0x75, 0x75, 0xe9, 0xe4, 0x6c, 0xfc, 0x01, 0xd2, 0x36, 0x37, 0x04, 0xd4, 0xb4, 0xcd,
	0x55, 0xf0, 0xdf, 0x4b, 0xd3, 0x74, 0xe8, 0x96, 0x7b, 0x9e, 0xe3, 0x4d, 0xde, 0xc0, 0x65, 0x5a,
	0x1c, 0x93, 0x1a, 0x67, 0xc7, 0xae, 0x7a, 0x84, 0xeb, 0xf7, 0x75, 0xf4, 0x9a, 0xa6, 0x85, 0x3c,
	0xe3, 0x2d, 0xa4, 0x5f, 0xf6, 0xdb, 0xb2, 0x14, 0xa5, 0xa8, 0x53, 0xbd, 0x0d, 0xd5, 0x33, 0x14,
	0xfb, 0x9a, 0x1f, 0xdd, 0x8f, 0x27, 0x7c, 0x80, 0x2c, 0xe4, 0x78, 0x29, 0xca, 0x53, 0x7d, 0x69,
	0x32, 0x15, 0x16, 0x74, 0xa4, 0x95, 0x85, 0x34, 0x00, 0x2c, 0x20, 0xb1, 0x26, 0xa6, 0x25, 0xd6,
	0x20, 0xc2, 0xb9, 0x77, 0xe6, 0x4f, 0x26, 0xa5, 0xa8, 0x73, 0x1d, 0xce, 0x78, 0x0f, 0x30, 0xcc,
	0xd4, 0x31, 0x99, 0xcf, 0x8e, 0xe5, 0x29, 0x98, 0x3c, 0x92, 0x96, 0x57, 0xbd, 0x8c, 0x66, 0xd7,
	0xe7, 0x4d, 0x47, 0xd2, 0x72, 0xf3, 0x0a, 0x57, 0xe1, 0xaa, 0x0f, 0x9a, 0x7f, 0xed, 0x40, 0xf8,
	0x04, 0xf9, 0x1b, 0xf1, 0xf6, 0x5e, 0x2c, 0xd4, 0xa1, 0xdf, 0xdd, 0x8d, 0x3a, 0x16, 0xe9, 0xb3,
	0xf0, 0x11, 0x2f, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x25, 0x9c, 0x76, 0x19, 0x17, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// QuoteServiceClient is the client API for QuoteService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QuoteServiceClient interface {
	GetQuotes(ctx context.Context, in *QuotesRequest, opts ...grpc.CallOption) (*QuotesResponse, error)
}

type quoteServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewQuoteServiceClient(cc grpc.ClientConnInterface) QuoteServiceClient {
	return &quoteServiceClient{cc}
}

func (c *quoteServiceClient) GetQuotes(ctx context.Context, in *QuotesRequest, opts ...grpc.CallOption) (*QuotesResponse, error) {
	out := new(QuotesResponse)
	err := c.cc.Invoke(ctx, "/QuoteService/GetQuotes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuoteServiceServer is the server API for QuoteService service.
type QuoteServiceServer interface {
	GetQuotes(context.Context, *QuotesRequest) (*QuotesResponse, error)
}

// UnimplementedQuoteServiceServer can be embedded to have forward compatible implementations.
type UnimplementedQuoteServiceServer struct {
}

func (*UnimplementedQuoteServiceServer) GetQuotes(ctx context.Context, req *QuotesRequest) (*QuotesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuotes not implemented")
}

func RegisterQuoteServiceServer(s *grpc.Server, srv QuoteServiceServer) {
	s.RegisterService(&_QuoteService_serviceDesc, srv)
}

func _QuoteService_GetQuotes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuotesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuoteServiceServer).GetQuotes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/QuoteService/GetQuotes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuoteServiceServer).GetQuotes(ctx, req.(*QuotesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _QuoteService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "QuoteService",
	HandlerType: (*QuoteServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetQuotes",
			Handler:    _QuoteService_GetQuotes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "quote.proto",
}