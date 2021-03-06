// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/dict.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	libresp "github.com/pku-hit/libresp"
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

type DictType int32

const (
	DictType_OmitType DictType = 0
	DictType_Root     DictType = 1
	DictType_Group    DictType = 2
	DictType_Node     DictType = 3
)

var DictType_name = map[int32]string{
	0: "OmitType",
	1: "Root",
	2: "Group",
	3: "Node",
}

var DictType_value = map[string]int32{
	"OmitType": 0,
	"Root":     1,
	"Group":    2,
	"Node":     3,
}

func (x DictType) String() string {
	return proto.EnumName(DictType_name, int32(x))
}

func (DictType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_22c1e80adc3f1fcd, []int{0}
}

type DictStatus int32

const (
	DictStatus_OmitStatus DictStatus = 0
	DictStatus_Normal     DictStatus = 1
	DictStatus_Deleted    DictStatus = 2
)

var DictStatus_name = map[int32]string{
	0: "OmitStatus",
	1: "Normal",
	2: "Deleted",
}

var DictStatus_value = map[string]int32{
	"OmitStatus": 0,
	"Normal":     1,
	"Deleted":    2,
}

func (x DictStatus) String() string {
	return proto.EnumName(DictStatus_name, int32(x))
}

func (DictStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_22c1e80adc3f1fcd, []int{1}
}

type DictItem struct {
	DictUniqueId         string     `protobuf:"bytes,1,opt,name=dictUniqueId,proto3" json:"dictUniqueId,omitempty"`
	Code                 string     `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Name                 string     `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Type                 DictType   `protobuf:"varint,4,opt,name=type,proto3,enum=proto.DictType" json:"type,omitempty"`
	Status               DictStatus `protobuf:"varint,5,opt,name=status,proto3,enum=proto.DictStatus" json:"status,omitempty"`
	Value                string     `protobuf:"bytes,6,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *DictItem) Reset()         { *m = DictItem{} }
func (m *DictItem) String() string { return proto.CompactTextString(m) }
func (*DictItem) ProtoMessage()    {}
func (*DictItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_22c1e80adc3f1fcd, []int{0}
}

func (m *DictItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DictItem.Unmarshal(m, b)
}
func (m *DictItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DictItem.Marshal(b, m, deterministic)
}
func (m *DictItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DictItem.Merge(m, src)
}
func (m *DictItem) XXX_Size() int {
	return xxx_messageInfo_DictItem.Size(m)
}
func (m *DictItem) XXX_DiscardUnknown() {
	xxx_messageInfo_DictItem.DiscardUnknown(m)
}

var xxx_messageInfo_DictItem proto.InternalMessageInfo

func (m *DictItem) GetDictUniqueId() string {
	if m != nil {
		return m.DictUniqueId
	}
	return ""
}

func (m *DictItem) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *DictItem) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DictItem) GetType() DictType {
	if m != nil {
		return m.Type
	}
	return DictType_OmitType
}

func (m *DictItem) GetStatus() DictStatus {
	if m != nil {
		return m.Status
	}
	return DictStatus_OmitStatus
}

func (m *DictItem) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type ListChildrenRequest struct {
	Type                 DictType `protobuf:"varint,1,opt,name=type,proto3,enum=proto.DictType" json:"type,omitempty"`
	ParentId             string   `protobuf:"bytes,2,opt,name=parentId,proto3" json:"parentId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListChildrenRequest) Reset()         { *m = ListChildrenRequest{} }
func (m *ListChildrenRequest) String() string { return proto.CompactTextString(m) }
func (*ListChildrenRequest) ProtoMessage()    {}
func (*ListChildrenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_22c1e80adc3f1fcd, []int{1}
}

func (m *ListChildrenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListChildrenRequest.Unmarshal(m, b)
}
func (m *ListChildrenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListChildrenRequest.Marshal(b, m, deterministic)
}
func (m *ListChildrenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListChildrenRequest.Merge(m, src)
}
func (m *ListChildrenRequest) XXX_Size() int {
	return xxx_messageInfo_ListChildrenRequest.Size(m)
}
func (m *ListChildrenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListChildrenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListChildrenRequest proto.InternalMessageInfo

func (m *ListChildrenRequest) GetType() DictType {
	if m != nil {
		return m.Type
	}
	return DictType_OmitType
}

func (m *ListChildrenRequest) GetParentId() string {
	if m != nil {
		return m.ParentId
	}
	return ""
}

type AddDictRequest struct {
	Category             string   `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
	ParentId             string   `protobuf:"bytes,2,opt,name=parentId,proto3" json:"parentId,omitempty"`
	Type                 DictType `protobuf:"varint,3,opt,name=type,proto3,enum=proto.DictType" json:"type,omitempty"`
	Code                 string   `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
	PyCode               string   `protobuf:"bytes,5,opt,name=pyCode,proto3" json:"pyCode,omitempty"`
	Name                 string   `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	Value                string   `protobuf:"bytes,7,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddDictRequest) Reset()         { *m = AddDictRequest{} }
func (m *AddDictRequest) String() string { return proto.CompactTextString(m) }
func (*AddDictRequest) ProtoMessage()    {}
func (*AddDictRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_22c1e80adc3f1fcd, []int{2}
}

func (m *AddDictRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddDictRequest.Unmarshal(m, b)
}
func (m *AddDictRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddDictRequest.Marshal(b, m, deterministic)
}
func (m *AddDictRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddDictRequest.Merge(m, src)
}
func (m *AddDictRequest) XXX_Size() int {
	return xxx_messageInfo_AddDictRequest.Size(m)
}
func (m *AddDictRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddDictRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddDictRequest proto.InternalMessageInfo

func (m *AddDictRequest) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *AddDictRequest) GetParentId() string {
	if m != nil {
		return m.ParentId
	}
	return ""
}

func (m *AddDictRequest) GetType() DictType {
	if m != nil {
		return m.Type
	}
	return DictType_OmitType
}

func (m *AddDictRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *AddDictRequest) GetPyCode() string {
	if m != nil {
		return m.PyCode
	}
	return ""
}

func (m *AddDictRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AddDictRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterEnum("proto.DictType", DictType_name, DictType_value)
	proto.RegisterEnum("proto.DictStatus", DictStatus_name, DictStatus_value)
	proto.RegisterType((*DictItem)(nil), "proto.DictItem")
	proto.RegisterType((*ListChildrenRequest)(nil), "proto.ListChildrenRequest")
	proto.RegisterType((*AddDictRequest)(nil), "proto.AddDictRequest")
}

func init() { proto.RegisterFile("proto/dict.proto", fileDescriptor_22c1e80adc3f1fcd) }

var fileDescriptor_22c1e80adc3f1fcd = []byte{
	// 526 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x8e, 0x13, 0xc7, 0x71, 0x87, 0x28, 0xb8, 0x0b, 0xad, 0x2c, 0x83, 0x50, 0x15, 0x2e, 0xa5,
	0x52, 0x6c, 0xa9, 0x08, 0x71, 0x80, 0x0b, 0x4a, 0xaa, 0x28, 0x12, 0x2a, 0x52, 0x0a, 0xbd, 0x3b,
	0xf6, 0xe0, 0xac, 0xb0, 0xbd, 0xdb, 0xf5, 0x1a, 0x94, 0xc7, 0xe0, 0x7d, 0xb8, 0xf1, 0x62, 0x68,
	0xd7, 0x1b, 0x93, 0x00, 0x0d, 0xa7, 0x9d, 0xbf, 0x6f, 0x66, 0xf6, 0x9b, 0x0f, 0x3c, 0x2e, 0x98,
	0x64, 0x51, 0x4a, 0x13, 0x19, 0x6a, 0x93, 0xf4, 0xf5, 0x13, 0x4c, 0x32, 0x2a, 0xd7, 0xf5, 0x2a,
	0x4c, 0x58, 0x11, 0xf1, 0x2f, 0xf5, 0x64, 0x4d, 0x65, 0x94, 0xd3, 0x95, 0xc0, 0x8a, 0x47, 0x45,
	0x95, 0x4d, 0x94, 0xc1, 0xca, 0x0a, 0x1b, 0x54, 0xf0, 0x24, 0x63, 0x2c, 0xcb, 0x31, 0xd2, 0xde,
	0xaa, 0xfe, 0x1c, 0x61, 0xc1, 0xe5, 0xc6, 0x24, 0x9f, 0xfd, 0x99, 0xfc, 0x26, 0x62, 0xce, 0x51,
	0x54, 0x4d, 0x7e, 0xfc, 0xc3, 0x02, 0x77, 0x46, 0x13, 0xb9, 0x90, 0x58, 0x90, 0x31, 0x0c, 0xd5,
	0x36, 0x9f, 0x4a, 0x7a, 0x57, 0xe3, 0x22, 0xf5, 0xad, 0x33, 0xeb, 0xfc, 0x68, 0xb9, 0x17, 0x23,
	0x04, 0xec, 0x84, 0xa5, 0xe8, 0x77, 0x75, 0x4e, 0xdb, 0x2a, 0x56, 0xc6, 0x05, 0xfa, 0xbd, 0x26,
	0xa6, 0x6c, 0xf2, 0x1c, 0x6c, 0xb9, 0xe1, 0xe8, 0xdb, 0x67, 0xd6, 0xf9, 0xe8, 0xf2, 0x61, 0x33,
	0x2e, 0x54, 0xa3, 0x3e, 0x6e, 0x38, 0x2e, 0x75, 0x92, 0xbc, 0x00, 0xa7, 0x92, 0xb1, 0xac, 0x2b,
	0xbf, 0xaf, 0xcb, 0x8e, 0x77, 0xca, 0x6e, 0x74, 0x62, 0x69, 0x0a, 0xc8, 0x63, 0xe8, 0x7f, 0x8d,
	0xf3, 0x1a, 0x7d, 0x47, 0x0f, 0x69, 0x9c, 0xf1, 0x2d, 0x3c, 0x7a, 0x4f, 0x2b, 0x39, 0x5d, 0xd3,
	0x3c, 0x15, 0x58, 0x2e, 0xf1, 0xae, 0xc6, 0x4a, 0xb6, 0xc3, 0xad, 0x43, 0xc3, 0x03, 0x70, 0x79,
	0x2c, 0xb0, 0x94, 0x8b, 0xd4, 0xfc, 0xa6, 0xf5, 0xc7, 0x3f, 0x2d, 0x18, 0xbd, 0x4b, 0x53, 0x85,
	0xd8, 0xf6, 0x0c, 0xc0, 0x4d, 0x62, 0x89, 0x19, 0x13, 0x1b, 0x43, 0x4c, 0xeb, 0x1f, 0x6a, 0xd5,
	0xee, 0xd2, 0x3b, 0xb4, 0xcb, 0x96, 0x55, 0x7b, 0x87, 0xd5, 0x53, 0x70, 0xf8, 0x66, 0xaa, 0xa2,
	0x7d, 0x1d, 0x35, 0x5e, 0xcb, 0xb6, 0xb3, 0xc3, 0x76, 0xcb, 0xce, 0x60, 0x87, 0x9d, 0x8b, 0xd7,
	0xcd, 0x6d, 0xd5, 0x1c, 0x32, 0x04, 0xf7, 0x43, 0x41, 0xb5, 0xed, 0x75, 0x88, 0x0b, 0xf6, 0x92,
	0x31, 0xe9, 0x59, 0xe4, 0x08, 0xfa, 0x73, 0xc1, 0x6a, 0xee, 0x75, 0x55, 0xf0, 0x9a, 0xa5, 0xe8,
	0xf5, 0x2e, 0x5e, 0x01, 0xfc, 0x3e, 0x01, 0x19, 0x01, 0x28, 0x68, 0xe3, 0x79, 0x1d, 0x02, 0xe0,
	0x5c, 0x33, 0x51, 0xc4, 0xb9, 0x67, 0x91, 0x07, 0x30, 0x98, 0x61, 0x8e, 0x12, 0x53, 0xaf, 0x7b,
	0xf9, 0xbd, 0x07, 0xb6, 0xc2, 0x91, 0x37, 0xe0, 0xaa, 0xb3, 0xa8, 0x11, 0xe4, 0x34, 0x6c, 0x24,
	0x18, 0x6e, 0x25, 0x18, 0x5e, 0x29, 0x7d, 0x06, 0x27, 0xa1, 0xd1, 0x74, 0xa8, 0x4b, 0x8d, 0xa6,
	0xc7, 0x1d, 0x72, 0x05, 0x43, 0x7d, 0xd3, 0x2d, 0xb9, 0x4f, 0xff, 0x6a, 0x70, 0x23, 0x05, 0x2d,
	0xb3, 0x5b, 0xf5, 0xc7, 0xfb, 0xdb, 0x4c, 0x4d, 0x1b, 0x23, 0x0d, 0x12, 0x18, 0xe6, 0xff, 0xa1,
	0x97, 0xfb, 0x9b, 0xbc, 0x85, 0x81, 0x91, 0x01, 0x39, 0x31, 0xf8, 0x7d, 0x59, 0x04, 0x7e, 0x0b,
	0x9d, 0x63, 0x89, 0x82, 0x26, 0xfb, 0xe8, 0x19, 0xe6, 0x1a, 0x7d, 0xf8, 0x13, 0xc7, 0x6d, 0x93,
	0x1d, 0xf4, 0x0c, 0xdc, 0x39, 0x4a, 0x5d, 0xf0, 0x1f, 0xf8, 0x81, 0x1d, 0x56, 0x8e, 0x46, 0xbc,
	0xfc, 0x15, 0x00, 0x00, 0xff, 0xff, 0x27, 0xc2, 0xf3, 0xdd, 0x6e, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DictClient is the client API for Dict service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DictClient interface {
	// 获取全部Root节点列表
	ListRoot(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*libresp.ListResponse, error)
	// 获取某一分类下的Root节点列表
	ListCategory(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (*libresp.ListResponse, error)
	// 获取某一节点组下的子节点列表
	ListChildren(ctx context.Context, in *ListChildrenRequest, opts ...grpc.CallOption) (*libresp.ListResponse, error)
	// 添加字典
	AddDict(ctx context.Context, in *AddDictRequest, opts ...grpc.CallOption) (*libresp.GenericResponse, error)
	// 删除字典项
	DelDict(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (*libresp.Response, error)
	// 通过字典ID获取指定节点的值
	GetValue(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (*libresp.GenericResponse, error)
}

type dictClient struct {
	cc *grpc.ClientConn
}

func NewDictClient(cc *grpc.ClientConn) DictClient {
	return &dictClient{cc}
}

func (c *dictClient) ListRoot(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*libresp.ListResponse, error) {
	out := new(libresp.ListResponse)
	err := c.cc.Invoke(ctx, "/proto.Dict/ListRoot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictClient) ListCategory(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (*libresp.ListResponse, error) {
	out := new(libresp.ListResponse)
	err := c.cc.Invoke(ctx, "/proto.Dict/ListCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictClient) ListChildren(ctx context.Context, in *ListChildrenRequest, opts ...grpc.CallOption) (*libresp.ListResponse, error) {
	out := new(libresp.ListResponse)
	err := c.cc.Invoke(ctx, "/proto.Dict/ListChildren", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictClient) AddDict(ctx context.Context, in *AddDictRequest, opts ...grpc.CallOption) (*libresp.GenericResponse, error) {
	out := new(libresp.GenericResponse)
	err := c.cc.Invoke(ctx, "/proto.Dict/AddDict", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictClient) DelDict(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (*libresp.Response, error) {
	out := new(libresp.Response)
	err := c.cc.Invoke(ctx, "/proto.Dict/DelDict", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictClient) GetValue(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (*libresp.GenericResponse, error) {
	out := new(libresp.GenericResponse)
	err := c.cc.Invoke(ctx, "/proto.Dict/GetValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DictServer is the server API for Dict service.
type DictServer interface {
	// 获取全部Root节点列表
	ListRoot(context.Context, *empty.Empty) (*libresp.ListResponse, error)
	// 获取某一分类下的Root节点列表
	ListCategory(context.Context, *wrappers.StringValue) (*libresp.ListResponse, error)
	// 获取某一节点组下的子节点列表
	ListChildren(context.Context, *ListChildrenRequest) (*libresp.ListResponse, error)
	// 添加字典
	AddDict(context.Context, *AddDictRequest) (*libresp.GenericResponse, error)
	// 删除字典项
	DelDict(context.Context, *wrappers.StringValue) (*libresp.Response, error)
	// 通过字典ID获取指定节点的值
	GetValue(context.Context, *wrappers.StringValue) (*libresp.GenericResponse, error)
}

// UnimplementedDictServer can be embedded to have forward compatible implementations.
type UnimplementedDictServer struct {
}

func (*UnimplementedDictServer) ListRoot(ctx context.Context, req *empty.Empty) (*libresp.ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRoot not implemented")
}
func (*UnimplementedDictServer) ListCategory(ctx context.Context, req *wrappers.StringValue) (*libresp.ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCategory not implemented")
}
func (*UnimplementedDictServer) ListChildren(ctx context.Context, req *ListChildrenRequest) (*libresp.ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListChildren not implemented")
}
func (*UnimplementedDictServer) AddDict(ctx context.Context, req *AddDictRequest) (*libresp.GenericResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDict not implemented")
}
func (*UnimplementedDictServer) DelDict(ctx context.Context, req *wrappers.StringValue) (*libresp.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelDict not implemented")
}
func (*UnimplementedDictServer) GetValue(ctx context.Context, req *wrappers.StringValue) (*libresp.GenericResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetValue not implemented")
}

func RegisterDictServer(s *grpc.Server, srv DictServer) {
	s.RegisterService(&_Dict_serviceDesc, srv)
}

func _Dict_ListRoot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictServer).ListRoot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Dict/ListRoot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictServer).ListRoot(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dict_ListCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrappers.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictServer).ListCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Dict/ListCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictServer).ListCategory(ctx, req.(*wrappers.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dict_ListChildren_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListChildrenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictServer).ListChildren(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Dict/ListChildren",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictServer).ListChildren(ctx, req.(*ListChildrenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dict_AddDict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDictRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictServer).AddDict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Dict/AddDict",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictServer).AddDict(ctx, req.(*AddDictRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dict_DelDict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrappers.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictServer).DelDict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Dict/DelDict",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictServer).DelDict(ctx, req.(*wrappers.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dict_GetValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrappers.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictServer).GetValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Dict/GetValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictServer).GetValue(ctx, req.(*wrappers.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

var _Dict_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Dict",
	HandlerType: (*DictServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListRoot",
			Handler:    _Dict_ListRoot_Handler,
		},
		{
			MethodName: "ListCategory",
			Handler:    _Dict_ListCategory_Handler,
		},
		{
			MethodName: "ListChildren",
			Handler:    _Dict_ListChildren_Handler,
		},
		{
			MethodName: "AddDict",
			Handler:    _Dict_AddDict_Handler,
		},
		{
			MethodName: "DelDict",
			Handler:    _Dict_DelDict_Handler,
		},
		{
			MethodName: "GetValue",
			Handler:    _Dict_GetValue_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/dict.proto",
}
