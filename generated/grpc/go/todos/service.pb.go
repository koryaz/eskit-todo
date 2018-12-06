// Code generated by protoc-gen-go. DO NOT EDIT.
// source: todos/service.proto

package todos

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	common "github.com/makkalot/eskit/generated/grpc/go/common"
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

type CreateRequest struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Text                 string   `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7834060e39fdfb48, []int{0}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *CreateRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *CreateRequest) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type CreateResponse struct {
	Todo                 *Todo    `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7834060e39fdfb48, []int{1}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetTodo() *Todo {
	if m != nil {
		return m.Todo
	}
	return nil
}

type FindRequest struct {
	Originator           *common.Originator `protobuf:"bytes,1,opt,name=originator,proto3" json:"originator,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *FindRequest) Reset()         { *m = FindRequest{} }
func (m *FindRequest) String() string { return proto.CompactTextString(m) }
func (*FindRequest) ProtoMessage()    {}
func (*FindRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7834060e39fdfb48, []int{2}
}

func (m *FindRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindRequest.Unmarshal(m, b)
}
func (m *FindRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindRequest.Marshal(b, m, deterministic)
}
func (m *FindRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindRequest.Merge(m, src)
}
func (m *FindRequest) XXX_Size() int {
	return xxx_messageInfo_FindRequest.Size(m)
}
func (m *FindRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindRequest proto.InternalMessageInfo

func (m *FindRequest) GetOriginator() *common.Originator {
	if m != nil {
		return m.Originator
	}
	return nil
}

type FindResponse struct {
	Todo                 *Todo    `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindResponse) Reset()         { *m = FindResponse{} }
func (m *FindResponse) String() string { return proto.CompactTextString(m) }
func (*FindResponse) ProtoMessage()    {}
func (*FindResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7834060e39fdfb48, []int{3}
}

func (m *FindResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindResponse.Unmarshal(m, b)
}
func (m *FindResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindResponse.Marshal(b, m, deterministic)
}
func (m *FindResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindResponse.Merge(m, src)
}
func (m *FindResponse) XXX_Size() int {
	return xxx_messageInfo_FindResponse.Size(m)
}
func (m *FindResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindResponse proto.InternalMessageInfo

func (m *FindResponse) GetTodo() *Todo {
	if m != nil {
		return m.Todo
	}
	return nil
}

type UpdateRequest struct {
	// version bit for the update requests are required
	Originator           *common.Originator `protobuf:"bytes,1,opt,name=originator,proto3" json:"originator,omitempty"`
	UpdatedTodo          *Todo              `protobuf:"bytes,2,opt,name=updatedTodo,proto3" json:"updatedTodo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7834060e39fdfb48, []int{4}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetOriginator() *common.Originator {
	if m != nil {
		return m.Originator
	}
	return nil
}

func (m *UpdateRequest) GetUpdatedTodo() *Todo {
	if m != nil {
		return m.UpdatedTodo
	}
	return nil
}

type UpdateResponse struct {
	Todo                 *Todo    `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7834060e39fdfb48, []int{5}
}

func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (m *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(m, src)
}
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

func (m *UpdateResponse) GetTodo() *Todo {
	if m != nil {
		return m.Todo
	}
	return nil
}

type DeleteRequest struct {
	Originator           *common.Originator `protobuf:"bytes,1,opt,name=originator,proto3" json:"originator,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7834060e39fdfb48, []int{6}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetOriginator() *common.Originator {
	if m != nil {
		return m.Originator
	}
	return nil
}

type DeleteResponse struct {
	Originator           *common.Originator `protobuf:"bytes,1,opt,name=originator,proto3" json:"originator,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7834060e39fdfb48, []int{7}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

func (m *DeleteResponse) GetOriginator() *common.Originator {
	if m != nil {
		return m.Originator
	}
	return nil
}

type DeleteAllRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteAllRequest) Reset()         { *m = DeleteAllRequest{} }
func (m *DeleteAllRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteAllRequest) ProtoMessage()    {}
func (*DeleteAllRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7834060e39fdfb48, []int{8}
}

func (m *DeleteAllRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteAllRequest.Unmarshal(m, b)
}
func (m *DeleteAllRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteAllRequest.Marshal(b, m, deterministic)
}
func (m *DeleteAllRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteAllRequest.Merge(m, src)
}
func (m *DeleteAllRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteAllRequest.Size(m)
}
func (m *DeleteAllRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteAllRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteAllRequest proto.InternalMessageInfo

type DeleteAllResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteAllResponse) Reset()         { *m = DeleteAllResponse{} }
func (m *DeleteAllResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteAllResponse) ProtoMessage()    {}
func (*DeleteAllResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7834060e39fdfb48, []int{9}
}

func (m *DeleteAllResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteAllResponse.Unmarshal(m, b)
}
func (m *DeleteAllResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteAllResponse.Marshal(b, m, deterministic)
}
func (m *DeleteAllResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteAllResponse.Merge(m, src)
}
func (m *DeleteAllResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteAllResponse.Size(m)
}
func (m *DeleteAllResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteAllResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteAllResponse proto.InternalMessageInfo

func (m *DeleteAllResponse) GetMessage() string {
	if m != nil {
		return m.Message
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
	return fileDescriptor_7834060e39fdfb48, []int{10}
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
	return fileDescriptor_7834060e39fdfb48, []int{11}
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
	proto.RegisterType((*CreateRequest)(nil), "contracts.todos.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "contracts.todos.CreateResponse")
	proto.RegisterType((*FindRequest)(nil), "contracts.todos.FindRequest")
	proto.RegisterType((*FindResponse)(nil), "contracts.todos.FindResponse")
	proto.RegisterType((*UpdateRequest)(nil), "contracts.todos.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "contracts.todos.UpdateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "contracts.todos.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "contracts.todos.DeleteResponse")
	proto.RegisterType((*DeleteAllRequest)(nil), "contracts.todos.DeleteAllRequest")
	proto.RegisterType((*DeleteAllResponse)(nil), "contracts.todos.DeleteAllResponse")
	proto.RegisterType((*HealthRequest)(nil), "contracts.todos.HealthRequest")
	proto.RegisterType((*HealthResponse)(nil), "contracts.todos.HealthResponse")
}

func init() { proto.RegisterFile("todos/service.proto", fileDescriptor_7834060e39fdfb48) }

var fileDescriptor_7834060e39fdfb48 = []byte{
	// 520 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x6e, 0x9a, 0x10, 0xd4, 0x49, 0x93, 0x86, 0x05, 0x44, 0x64, 0x85, 0xbf, 0x3d, 0x01, 0x52,
	0xbd, 0xa2, 0x20, 0x10, 0x82, 0x0b, 0xbf, 0x42, 0x8a, 0x00, 0x29, 0x94, 0x0b, 0x12, 0x07, 0xd7,
	0x1e, 0x39, 0x56, 0x1d, 0x8f, 0xd9, 0xdd, 0x54, 0xd0, 0x23, 0x07, 0x5e, 0x80, 0x47, 0xe3, 0x15,
	0x78, 0x00, 0x1e, 0x01, 0xed, 0xae, 0x9d, 0xda, 0xd4, 0x01, 0x55, 0xed, 0x25, 0xb2, 0x67, 0xbe,
	0xf9, 0xe6, 0x9b, 0xc9, 0x37, 0x86, 0x8b, 0x9a, 0x22, 0x52, 0x42, 0xa1, 0x3c, 0x48, 0x42, 0xf4,
	0x73, 0x49, 0x9a, 0xd8, 0x56, 0x48, 0x99, 0x96, 0x41, 0xa8, 0x95, 0x6f, 0xd3, 0xde, 0xd0, 0xa1,
	0xcc, 0xaf, 0x83, 0x78, 0x57, 0x42, 0x9a, 0xcf, 0x29, 0x13, 0x24, 0x93, 0x38, 0xc9, 0x02, 0x4d,
	0xb2, 0x48, 0x8c, 0x63, 0xa2, 0x38, 0x45, 0x11, 0xe4, 0x89, 0x08, 0xb2, 0x8c, 0x74, 0xa0, 0x13,
	0xca, 0x94, 0xcb, 0xf2, 0x09, 0xf4, 0x9f, 0x4b, 0x0c, 0x34, 0x4e, 0xf1, 0xf3, 0x02, 0x95, 0x66,
	0x43, 0x68, 0x2f, 0x64, 0x3a, 0x6a, 0xdd, 0x68, 0xdd, 0xda, 0x98, 0x9a, 0x47, 0x76, 0x09, 0xce,
	0xe9, 0x44, 0xa7, 0x38, 0x5a, 0xb7, 0x31, 0xf7, 0xc2, 0x18, 0x74, 0x34, 0x7e, 0xd1, 0xa3, 0xb6,
	0x0d, 0xda, 0x67, 0xfe, 0x18, 0x06, 0x25, 0x99, 0xca, 0x29, 0x53, 0xc8, 0x6e, 0x43, 0xc7, 0x68,
	0xb4, 0x74, 0xbd, 0x9d, 0xcb, 0xfe, 0x5f, 0x73, 0xf8, 0xbb, 0x14, 0xd1, 0xd4, 0x42, 0xf8, 0x04,
	0x7a, 0xaf, 0x92, 0x2c, 0x2a, 0x75, 0x3c, 0x01, 0x38, 0x1a, 0xa5, 0xa8, 0x1f, 0x57, 0xea, 0xdd,
	0xb8, 0xfe, 0xbb, 0x25, 0x66, 0x5a, 0xc1, 0xf3, 0x47, 0xb0, 0xe9, 0xc8, 0x4e, 0xae, 0xe3, 0x7b,
	0x0b, 0xfa, 0x1f, 0xf2, 0xa8, 0xb2, 0x92, 0x53, 0x49, 0x61, 0x0f, 0xa1, 0xb7, 0xb0, 0x74, 0x91,
	0x69, 0x62, 0x97, 0xb8, 0x52, 0x41, 0x15, 0x69, 0xb6, 0x59, 0xea, 0x38, 0xf9, 0x14, 0x6f, 0xa0,
	0xff, 0x02, 0x53, 0x3c, 0xa3, 0x21, 0xf8, 0x5b, 0x18, 0x94, 0x74, 0x85, 0x96, 0xd3, 0xf1, 0x31,
	0x18, 0x3a, 0xbe, 0xa7, 0x69, 0x5a, 0x28, 0xe4, 0xdb, 0x70, 0xa1, 0x12, 0x2b, 0xda, 0x8c, 0xe0,
	0xfc, 0x1c, 0x95, 0x0a, 0x62, 0x2c, 0x2c, 0x59, 0xbe, 0xf2, 0x2d, 0xe8, 0xbf, 0xc6, 0x20, 0xd5,
	0xb3, 0xb2, 0xfe, 0x0e, 0x0c, 0xca, 0xc0, 0xff, 0x8a, 0x77, 0x7e, 0xb7, 0xa1, 0x67, 0xb6, 0xf5,
	0xde, 0x9d, 0x19, 0xfb, 0x04, 0x5d, 0x5b, 0x7b, 0xc8, 0xae, 0x1d, 0xdb, 0x6a, 0xad, 0x8b, 0x77,
	0x7d, 0x65, 0xde, 0x35, 0xe5, 0xec, 0xdb, 0xcf, 0x5f, 0x3f, 0xd6, 0x37, 0x19, 0x88, 0x83, 0xbb,
	0x62, 0xe6, 0x48, 0x27, 0xd0, 0x75, 0x87, 0xd1, 0x40, 0x5f, 0x3b, 0xbf, 0x06, 0xfa, 0xfa, 0x45,
	0xf1, 0x35, 0xf6, 0x12, 0x3a, 0xc6, 0xdb, 0x6c, 0x7c, 0x0c, 0x5a, 0xb9, 0x1f, 0xef, 0xea, 0x8a,
	0xec, 0x92, 0x66, 0x02, 0x5d, 0x67, 0xaf, 0x06, 0x4d, 0x35, 0xff, 0x37, 0x68, 0xaa, 0xfb, 0xd2,
	0x91, 0xb9, 0xff, 0xae, 0x81, 0xac, 0xe6, 0xc3, 0x06, 0xb2, 0xba, 0xb1, 0xf8, 0x1a, 0xdb, 0x85,
	0x8d, 0xa5, 0x11, 0xd8, 0xcd, 0x15, 0xf8, 0x23, 0xe3, 0x78, 0xfc, 0x5f, 0x90, 0x92, 0xf5, 0xd9,
	0x83, 0x8f, 0xf7, 0xe3, 0x44, 0xcf, 0x16, 0x7b, 0xc6, 0x9a, 0x62, 0x9f, 0xe4, 0xd7, 0xe0, 0x50,
	0xa0, 0xda, 0x4f, 0xf4, 0xb6, 0x29, 0x12, 0x31, 0x66, 0x28, 0xcd, 0xf1, 0x89, 0x58, 0xe6, 0xa1,
	0x88, 0xc9, 0x7e, 0x5d, 0xd5, 0x5e, 0xd7, 0x7e, 0x28, 0xef, 0xfd, 0x09, 0x00, 0x00, 0xff, 0xff,
	0xfc, 0xd3, 0x56, 0x55, 0x99, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TodoServiceClient is the client API for TodoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TodoServiceClient interface {
	Healtz(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Find(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*FindResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	DeleteAll(ctx context.Context, in *DeleteAllRequest, opts ...grpc.CallOption) (*DeleteAllResponse, error)
}

type todoServiceClient struct {
	cc *grpc.ClientConn
}

func NewTodoServiceClient(cc *grpc.ClientConn) TodoServiceClient {
	return &todoServiceClient{cc}
}

func (c *todoServiceClient) Healtz(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error) {
	out := new(HealthResponse)
	err := c.cc.Invoke(ctx, "/contracts.todos.TodoService/Healtz", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/contracts.todos.TodoService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) Find(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*FindResponse, error) {
	out := new(FindResponse)
	err := c.cc.Invoke(ctx, "/contracts.todos.TodoService/Find", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/contracts.todos.TodoService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/contracts.todos.TodoService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) DeleteAll(ctx context.Context, in *DeleteAllRequest, opts ...grpc.CallOption) (*DeleteAllResponse, error) {
	out := new(DeleteAllResponse)
	err := c.cc.Invoke(ctx, "/contracts.todos.TodoService/DeleteAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoServiceServer is the server API for TodoService service.
type TodoServiceServer interface {
	Healtz(context.Context, *HealthRequest) (*HealthResponse, error)
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Find(context.Context, *FindRequest) (*FindResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	DeleteAll(context.Context, *DeleteAllRequest) (*DeleteAllResponse, error)
}

func RegisterTodoServiceServer(s *grpc.Server, srv TodoServiceServer) {
	s.RegisterService(&_TodoService_serviceDesc, srv)
}

func _TodoService_Healtz_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).Healtz(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contracts.todos.TodoService/Healtz",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).Healtz(ctx, req.(*HealthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contracts.todos.TodoService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_Find_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).Find(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contracts.todos.TodoService/Find",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).Find(ctx, req.(*FindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contracts.todos.TodoService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contracts.todos.TodoService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_DeleteAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).DeleteAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contracts.todos.TodoService/DeleteAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).DeleteAll(ctx, req.(*DeleteAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TodoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "contracts.todos.TodoService",
	HandlerType: (*TodoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Healtz",
			Handler:    _TodoService_Healtz_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _TodoService_Create_Handler,
		},
		{
			MethodName: "Find",
			Handler:    _TodoService_Find_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _TodoService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _TodoService_Delete_Handler,
		},
		{
			MethodName: "DeleteAll",
			Handler:    _TodoService_DeleteAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "todos/service.proto",
}
