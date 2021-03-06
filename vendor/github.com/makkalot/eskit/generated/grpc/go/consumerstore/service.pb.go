// Code generated by protoc-gen-go. DO NOT EDIT.
// source: consumerstore/service.proto

package consumerstore

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/makkalot/eskit/generated/grpc/go/common"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ListConsumersRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListConsumersRequest) Reset()         { *m = ListConsumersRequest{} }
func (m *ListConsumersRequest) String() string { return proto.CompactTextString(m) }
func (*ListConsumersRequest) ProtoMessage()    {}
func (*ListConsumersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0042378e09c1851b, []int{0}
}

func (m *ListConsumersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListConsumersRequest.Unmarshal(m, b)
}
func (m *ListConsumersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListConsumersRequest.Marshal(b, m, deterministic)
}
func (m *ListConsumersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListConsumersRequest.Merge(m, src)
}
func (m *ListConsumersRequest) XXX_Size() int {
	return xxx_messageInfo_ListConsumersRequest.Size(m)
}
func (m *ListConsumersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListConsumersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListConsumersRequest proto.InternalMessageInfo

type ListConsumersResponse struct {
	Consumers            []*GetAppLogConsumeResponse `protobuf:"bytes,1,rep,name=consumers,proto3" json:"consumers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *ListConsumersResponse) Reset()         { *m = ListConsumersResponse{} }
func (m *ListConsumersResponse) String() string { return proto.CompactTextString(m) }
func (*ListConsumersResponse) ProtoMessage()    {}
func (*ListConsumersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0042378e09c1851b, []int{1}
}

func (m *ListConsumersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListConsumersResponse.Unmarshal(m, b)
}
func (m *ListConsumersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListConsumersResponse.Marshal(b, m, deterministic)
}
func (m *ListConsumersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListConsumersResponse.Merge(m, src)
}
func (m *ListConsumersResponse) XXX_Size() int {
	return xxx_messageInfo_ListConsumersResponse.Size(m)
}
func (m *ListConsumersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListConsumersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListConsumersResponse proto.InternalMessageInfo

func (m *ListConsumersResponse) GetConsumers() []*GetAppLogConsumeResponse {
	if m != nil {
		return m.Consumers
	}
	return nil
}

type AppLogConsumeRequest struct {
	// the consumer id
	ConsumerId           string   `protobuf:"bytes,1,opt,name=consumer_id,json=consumerId,proto3" json:"consumer_id,omitempty"`
	Offset               string   `protobuf:"bytes,2,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppLogConsumeRequest) Reset()         { *m = AppLogConsumeRequest{} }
func (m *AppLogConsumeRequest) String() string { return proto.CompactTextString(m) }
func (*AppLogConsumeRequest) ProtoMessage()    {}
func (*AppLogConsumeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0042378e09c1851b, []int{2}
}

func (m *AppLogConsumeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppLogConsumeRequest.Unmarshal(m, b)
}
func (m *AppLogConsumeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppLogConsumeRequest.Marshal(b, m, deterministic)
}
func (m *AppLogConsumeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppLogConsumeRequest.Merge(m, src)
}
func (m *AppLogConsumeRequest) XXX_Size() int {
	return xxx_messageInfo_AppLogConsumeRequest.Size(m)
}
func (m *AppLogConsumeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AppLogConsumeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AppLogConsumeRequest proto.InternalMessageInfo

func (m *AppLogConsumeRequest) GetConsumerId() string {
	if m != nil {
		return m.ConsumerId
	}
	return ""
}

func (m *AppLogConsumeRequest) GetOffset() string {
	if m != nil {
		return m.Offset
	}
	return ""
}

type AppLogConsumeResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppLogConsumeResponse) Reset()         { *m = AppLogConsumeResponse{} }
func (m *AppLogConsumeResponse) String() string { return proto.CompactTextString(m) }
func (*AppLogConsumeResponse) ProtoMessage()    {}
func (*AppLogConsumeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0042378e09c1851b, []int{3}
}

func (m *AppLogConsumeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppLogConsumeResponse.Unmarshal(m, b)
}
func (m *AppLogConsumeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppLogConsumeResponse.Marshal(b, m, deterministic)
}
func (m *AppLogConsumeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppLogConsumeResponse.Merge(m, src)
}
func (m *AppLogConsumeResponse) XXX_Size() int {
	return xxx_messageInfo_AppLogConsumeResponse.Size(m)
}
func (m *AppLogConsumeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AppLogConsumeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AppLogConsumeResponse proto.InternalMessageInfo

type GetAppLogConsumeRequest struct {
	ConsumerId           string   `protobuf:"bytes,1,opt,name=consumer_id,json=consumerId,proto3" json:"consumer_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAppLogConsumeRequest) Reset()         { *m = GetAppLogConsumeRequest{} }
func (m *GetAppLogConsumeRequest) String() string { return proto.CompactTextString(m) }
func (*GetAppLogConsumeRequest) ProtoMessage()    {}
func (*GetAppLogConsumeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0042378e09c1851b, []int{4}
}

func (m *GetAppLogConsumeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAppLogConsumeRequest.Unmarshal(m, b)
}
func (m *GetAppLogConsumeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAppLogConsumeRequest.Marshal(b, m, deterministic)
}
func (m *GetAppLogConsumeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAppLogConsumeRequest.Merge(m, src)
}
func (m *GetAppLogConsumeRequest) XXX_Size() int {
	return xxx_messageInfo_GetAppLogConsumeRequest.Size(m)
}
func (m *GetAppLogConsumeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAppLogConsumeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAppLogConsumeRequest proto.InternalMessageInfo

func (m *GetAppLogConsumeRequest) GetConsumerId() string {
	if m != nil {
		return m.ConsumerId
	}
	return ""
}

type GetAppLogConsumeResponse struct {
	ConsumerId           string   `protobuf:"bytes,1,opt,name=consumer_id,json=consumerId,proto3" json:"consumer_id,omitempty"`
	Offset               string   `protobuf:"bytes,2,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAppLogConsumeResponse) Reset()         { *m = GetAppLogConsumeResponse{} }
func (m *GetAppLogConsumeResponse) String() string { return proto.CompactTextString(m) }
func (*GetAppLogConsumeResponse) ProtoMessage()    {}
func (*GetAppLogConsumeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0042378e09c1851b, []int{5}
}

func (m *GetAppLogConsumeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAppLogConsumeResponse.Unmarshal(m, b)
}
func (m *GetAppLogConsumeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAppLogConsumeResponse.Marshal(b, m, deterministic)
}
func (m *GetAppLogConsumeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAppLogConsumeResponse.Merge(m, src)
}
func (m *GetAppLogConsumeResponse) XXX_Size() int {
	return xxx_messageInfo_GetAppLogConsumeResponse.Size(m)
}
func (m *GetAppLogConsumeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAppLogConsumeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAppLogConsumeResponse proto.InternalMessageInfo

func (m *GetAppLogConsumeResponse) GetConsumerId() string {
	if m != nil {
		return m.ConsumerId
	}
	return ""
}

func (m *GetAppLogConsumeResponse) GetOffset() string {
	if m != nil {
		return m.Offset
	}
	return ""
}

type HealthRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthRequest) Reset()         { *m = HealthRequest{} }
func (m *HealthRequest) String() string { return proto.CompactTextString(m) }
func (*HealthRequest) ProtoMessage()    {}
func (*HealthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0042378e09c1851b, []int{6}
}

func (m *HealthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthRequest.Unmarshal(m, b)
}
func (m *HealthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthRequest.Marshal(b, m, deterministic)
}
func (m *HealthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthRequest.Merge(m, src)
}
func (m *HealthRequest) XXX_Size() int {
	return xxx_messageInfo_HealthRequest.Size(m)
}
func (m *HealthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HealthRequest proto.InternalMessageInfo

type HealthResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthResponse) Reset()         { *m = HealthResponse{} }
func (m *HealthResponse) String() string { return proto.CompactTextString(m) }
func (*HealthResponse) ProtoMessage()    {}
func (*HealthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0042378e09c1851b, []int{7}
}

func (m *HealthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthResponse.Unmarshal(m, b)
}
func (m *HealthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthResponse.Marshal(b, m, deterministic)
}
func (m *HealthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthResponse.Merge(m, src)
}
func (m *HealthResponse) XXX_Size() int {
	return xxx_messageInfo_HealthResponse.Size(m)
}
func (m *HealthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HealthResponse proto.InternalMessageInfo

func (m *HealthResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*ListConsumersRequest)(nil), "contracts.consumerstore.ListConsumersRequest")
	proto.RegisterType((*ListConsumersResponse)(nil), "contracts.consumerstore.ListConsumersResponse")
	proto.RegisterType((*AppLogConsumeRequest)(nil), "contracts.consumerstore.AppLogConsumeRequest")
	proto.RegisterType((*AppLogConsumeResponse)(nil), "contracts.consumerstore.AppLogConsumeResponse")
	proto.RegisterType((*GetAppLogConsumeRequest)(nil), "contracts.consumerstore.GetAppLogConsumeRequest")
	proto.RegisterType((*GetAppLogConsumeResponse)(nil), "contracts.consumerstore.GetAppLogConsumeResponse")
	proto.RegisterType((*HealthRequest)(nil), "contracts.consumerstore.HealthRequest")
	proto.RegisterType((*HealthResponse)(nil), "contracts.consumerstore.HealthResponse")
}

func init() { proto.RegisterFile("consumerstore/service.proto", fileDescriptor_0042378e09c1851b) }

var fileDescriptor_0042378e09c1851b = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xc1, 0x6e, 0x13, 0x31,
	0x10, 0x4d, 0x5a, 0x14, 0xd4, 0x29, 0xa5, 0x92, 0xd5, 0x36, 0xab, 0x80, 0x44, 0xe5, 0x03, 0x54,
	0x48, 0xac, 0x69, 0x39, 0x01, 0x27, 0xe0, 0x50, 0x90, 0x2a, 0x55, 0x4a, 0x6f, 0x5c, 0x90, 0xbb,
	0x99, 0x3a, 0x56, 0x62, 0xcf, 0x62, 0x4f, 0x72, 0xe0, 0xc8, 0x2f, 0xf0, 0x63, 0x48, 0xfc, 0x02,
	0x1f, 0x82, 0xc8, 0xae, 0x29, 0x29, 0xbb, 0xa5, 0xe9, 0xd1, 0xe3, 0xe7, 0xf7, 0xde, 0x8c, 0xe7,
	0xc1, 0x83, 0x82, 0x7c, 0x9c, 0x39, 0x0c, 0x91, 0x29, 0xa0, 0x8a, 0x18, 0xe6, 0xb6, 0xc0, 0xbc,
	0x0c, 0xc4, 0x24, 0xfa, 0x05, 0x79, 0x0e, 0xba, 0xe0, 0x98, 0x2f, 0xc1, 0x06, 0xfd, 0x82, 0x9c,
	0x23, 0xaf, 0x28, 0x58, 0x63, 0xbd, 0x66, 0x0a, 0xd5, 0x8b, 0xc1, 0x43, 0x43, 0x64, 0xa6, 0xa8,
	0x74, 0x69, 0x95, 0xf6, 0x9e, 0x58, 0xb3, 0x25, 0x1f, 0xab, 0x5b, 0xb9, 0x07, 0x3b, 0x27, 0x36,
	0xf2, 0xbb, 0xc4, 0x35, 0xc4, 0xcf, 0x33, 0x8c, 0x2c, 0xc7, 0xb0, 0x7b, 0xa5, 0x1e, 0x4b, 0xf2,
	0x11, 0xc5, 0x29, 0x6c, 0xfc, 0x11, 0xce, 0xba, 0xfb, 0xeb, 0x07, 0x9b, 0x47, 0x87, 0x79, 0x8b,
	0xa9, 0xfc, 0x18, 0xf9, 0x4d, 0x59, 0x9e, 0x90, 0xa9, 0x79, 0x12, 0xcb, 0xf0, 0x92, 0x43, 0x9e,
	0xc2, 0xce, 0x15, 0xcc, 0xc2, 0x81, 0x78, 0x04, 0x9b, 0x09, 0xf4, 0xc9, 0x8e, 0xb2, 0xee, 0x7e,
	0xf7, 0x60, 0x63, 0x08, 0xa9, 0xf4, 0x61, 0x24, 0xf6, 0xa0, 0x47, 0x17, 0x17, 0x11, 0x39, 0x5b,
	0x5b, 0xdc, 0xd5, 0x27, 0xd9, 0x87, 0xdd, 0x46, 0x51, 0xf9, 0x0a, 0xfa, 0xff, 0x1a, 0xba, 0x99,
	0x98, 0x3c, 0x83, 0xac, 0xad, 0x99, 0xdb, 0x3b, 0xdd, 0x86, 0xad, 0xf7, 0xa8, 0xa7, 0x3c, 0x4e,
	0x53, 0x7f, 0x0a, 0xf7, 0x53, 0xa1, 0xe6, 0xce, 0xe0, 0xae, 0xc3, 0x18, 0xb5, 0xc1, 0x9a, 0x37,
	0x1d, 0x8f, 0xbe, 0xaf, 0xc3, 0x76, 0xfa, 0x9e, 0xb3, 0x6a, 0x47, 0x84, 0x83, 0xde, 0xe2, 0xfd,
	0x17, 0xf1, 0xb8, 0xf5, 0x4f, 0x96, 0x14, 0x07, 0x4f, 0xfe, 0x8b, 0xab, 0x87, 0x27, 0xbe, 0xfe,
	0xf8, 0xf9, 0x6d, 0xed, 0x9e, 0x00, 0x35, 0x3f, 0x54, 0xe3, 0x4a, 0xc4, 0x01, 0x5c, 0x8e, 0x43,
	0x3c, 0x6b, 0xa5, 0x6a, 0x1a, 0xf9, 0x20, 0xbf, 0x29, 0xbc, 0x36, 0xd0, 0x11, 0x73, 0xd8, 0x3a,
	0x46, 0xfe, 0x4b, 0xf1, 0xf9, 0x0a, 0x8b, 0x57, 0x89, 0xae, 0xbe, 0xaa, 0xb2, 0x23, 0x0c, 0xdc,
	0xf9, 0x9d, 0x85, 0x6b, 0x1a, 0x6c, 0x8a, 0xd0, 0x35, 0x0d, 0x36, 0x26, 0x4b, 0x76, 0xde, 0xbe,
	0xfe, 0xf8, 0xd2, 0x58, 0x1e, 0xcf, 0xce, 0xf3, 0x82, 0x9c, 0x72, 0x7a, 0x32, 0xd1, 0x53, 0x62,
	0x85, 0x71, 0x62, 0x59, 0x19, 0xf4, 0x18, 0x34, 0xe3, 0x48, 0x99, 0x50, 0x16, 0xca, 0x90, 0x5a,
	0x22, 0x3d, 0xef, 0x2d, 0x02, 0xfd, 0xe2, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa6, 0x67, 0x92,
	0x14, 0x3f, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ConsumerServiceClient is the client API for ConsumerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConsumerServiceClient interface {
	Healtz(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error)
	LogConsume(ctx context.Context, in *AppLogConsumeRequest, opts ...grpc.CallOption) (*AppLogConsumeResponse, error)
	GetLogConsume(ctx context.Context, in *GetAppLogConsumeRequest, opts ...grpc.CallOption) (*GetAppLogConsumeResponse, error)
	List(ctx context.Context, in *ListConsumersRequest, opts ...grpc.CallOption) (*ListConsumersResponse, error)
}

type consumerServiceClient struct {
	cc *grpc.ClientConn
}

func NewConsumerServiceClient(cc *grpc.ClientConn) ConsumerServiceClient {
	return &consumerServiceClient{cc}
}

func (c *consumerServiceClient) Healtz(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error) {
	out := new(HealthResponse)
	err := c.cc.Invoke(ctx, "/contracts.consumerstore.ConsumerService/Healtz", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consumerServiceClient) LogConsume(ctx context.Context, in *AppLogConsumeRequest, opts ...grpc.CallOption) (*AppLogConsumeResponse, error) {
	out := new(AppLogConsumeResponse)
	err := c.cc.Invoke(ctx, "/contracts.consumerstore.ConsumerService/LogConsume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consumerServiceClient) GetLogConsume(ctx context.Context, in *GetAppLogConsumeRequest, opts ...grpc.CallOption) (*GetAppLogConsumeResponse, error) {
	out := new(GetAppLogConsumeResponse)
	err := c.cc.Invoke(ctx, "/contracts.consumerstore.ConsumerService/GetLogConsume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consumerServiceClient) List(ctx context.Context, in *ListConsumersRequest, opts ...grpc.CallOption) (*ListConsumersResponse, error) {
	out := new(ListConsumersResponse)
	err := c.cc.Invoke(ctx, "/contracts.consumerstore.ConsumerService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConsumerServiceServer is the server API for ConsumerService service.
type ConsumerServiceServer interface {
	Healtz(context.Context, *HealthRequest) (*HealthResponse, error)
	LogConsume(context.Context, *AppLogConsumeRequest) (*AppLogConsumeResponse, error)
	GetLogConsume(context.Context, *GetAppLogConsumeRequest) (*GetAppLogConsumeResponse, error)
	List(context.Context, *ListConsumersRequest) (*ListConsumersResponse, error)
}

func RegisterConsumerServiceServer(s *grpc.Server, srv ConsumerServiceServer) {
	s.RegisterService(&_ConsumerService_serviceDesc, srv)
}

func _ConsumerService_Healtz_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsumerServiceServer).Healtz(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contracts.consumerstore.ConsumerService/Healtz",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsumerServiceServer).Healtz(ctx, req.(*HealthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsumerService_LogConsume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppLogConsumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsumerServiceServer).LogConsume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contracts.consumerstore.ConsumerService/LogConsume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsumerServiceServer).LogConsume(ctx, req.(*AppLogConsumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsumerService_GetLogConsume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppLogConsumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsumerServiceServer).GetLogConsume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contracts.consumerstore.ConsumerService/GetLogConsume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsumerServiceServer).GetLogConsume(ctx, req.(*GetAppLogConsumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsumerService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListConsumersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsumerServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contracts.consumerstore.ConsumerService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsumerServiceServer).List(ctx, req.(*ListConsumersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ConsumerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "contracts.consumerstore.ConsumerService",
	HandlerType: (*ConsumerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Healtz",
			Handler:    _ConsumerService_Healtz_Handler,
		},
		{
			MethodName: "LogConsume",
			Handler:    _ConsumerService_LogConsume_Handler,
		},
		{
			MethodName: "GetLogConsume",
			Handler:    _ConsumerService_GetLogConsume_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ConsumerService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "consumerstore/service.proto",
}
