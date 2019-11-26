// Code generated by protoc-gen-go. DO NOT EDIT.
// source: entry/entry.proto

package entry

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Item_UpdateType int32

const (
	Item_DELETE Item_UpdateType = 0
	Item_UPDATE Item_UpdateType = 1
	Item_INSERT Item_UpdateType = 2
)

var Item_UpdateType_name = map[int32]string{
	0: "DELETE",
	1: "UPDATE",
	2: "INSERT",
}

var Item_UpdateType_value = map[string]int32{
	"DELETE": 0,
	"UPDATE": 1,
	"INSERT": 2,
}

func (x Item_UpdateType) String() string {
	return proto.EnumName(Item_UpdateType_name, int32(x))
}

func (Item_UpdateType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_aca970399dcefd35, []int{3, 0}
}

type App struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CreatedTime          int64    `protobuf:"varint,3,opt,name=createdTime,proto3" json:"createdTime,omitempty"`
	UpdatedTime          int64    `protobuf:"varint,4,opt,name=updatedTime,proto3" json:"updatedTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *App) Reset()         { *m = App{} }
func (m *App) String() string { return proto.CompactTextString(m) }
func (*App) ProtoMessage()    {}
func (*App) Descriptor() ([]byte, []int) {
	return fileDescriptor_aca970399dcefd35, []int{0}
}

func (m *App) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_App.Unmarshal(m, b)
}
func (m *App) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_App.Marshal(b, m, deterministic)
}
func (m *App) XXX_Merge(src proto.Message) {
	xxx_messageInfo_App.Merge(m, src)
}
func (m *App) XXX_Size() int {
	return xxx_messageInfo_App.Size(m)
}
func (m *App) XXX_DiscardUnknown() {
	xxx_messageInfo_App.DiscardUnknown(m)
}

var xxx_messageInfo_App proto.InternalMessageInfo

func (m *App) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *App) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *App) GetCreatedTime() int64 {
	if m != nil {
		return m.CreatedTime
	}
	return 0
}

func (m *App) GetUpdatedTime() int64 {
	if m != nil {
		return m.UpdatedTime
	}
	return 0
}

type Cluster struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	AppName              string   `protobuf:"bytes,3,opt,name=appName,proto3" json:"appName,omitempty"`
	CreatedTime          int64    `protobuf:"varint,4,opt,name=createdTime,proto3" json:"createdTime,omitempty"`
	UpdatedTime          int64    `protobuf:"varint,5,opt,name=updatedTime,proto3" json:"updatedTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Cluster) Reset()         { *m = Cluster{} }
func (m *Cluster) String() string { return proto.CompactTextString(m) }
func (*Cluster) ProtoMessage()    {}
func (*Cluster) Descriptor() ([]byte, []int) {
	return fileDescriptor_aca970399dcefd35, []int{1}
}

func (m *Cluster) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cluster.Unmarshal(m, b)
}
func (m *Cluster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cluster.Marshal(b, m, deterministic)
}
func (m *Cluster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cluster.Merge(m, src)
}
func (m *Cluster) XXX_Size() int {
	return xxx_messageInfo_Cluster.Size(m)
}
func (m *Cluster) XXX_DiscardUnknown() {
	xxx_messageInfo_Cluster.DiscardUnknown(m)
}

var xxx_messageInfo_Cluster proto.InternalMessageInfo

func (m *Cluster) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Cluster) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Cluster) GetAppName() string {
	if m != nil {
		return m.AppName
	}
	return ""
}

func (m *Cluster) GetCreatedTime() int64 {
	if m != nil {
		return m.CreatedTime
	}
	return 0
}

func (m *Cluster) GetUpdatedTime() int64 {
	if m != nil {
		return m.UpdatedTime
	}
	return 0
}

type Instance struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AppName              string   `protobuf:"bytes,2,opt,name=appName,proto3" json:"appName,omitempty"`
	ClusterName          string   `protobuf:"bytes,3,opt,name=clusterName,proto3" json:"clusterName,omitempty"`
	Ip                   string   `protobuf:"bytes,4,opt,name=ip,proto3" json:"ip,omitempty"`
	CreatedTime          int64    `protobuf:"varint,5,opt,name=createdTime,proto3" json:"createdTime,omitempty"`
	UpdatedTime          int64    `protobuf:"varint,6,opt,name=updatedTime,proto3" json:"updatedTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Instance) Reset()         { *m = Instance{} }
func (m *Instance) String() string { return proto.CompactTextString(m) }
func (*Instance) ProtoMessage()    {}
func (*Instance) Descriptor() ([]byte, []int) {
	return fileDescriptor_aca970399dcefd35, []int{2}
}

func (m *Instance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Instance.Unmarshal(m, b)
}
func (m *Instance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Instance.Marshal(b, m, deterministic)
}
func (m *Instance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Instance.Merge(m, src)
}
func (m *Instance) XXX_Size() int {
	return xxx_messageInfo_Instance.Size(m)
}
func (m *Instance) XXX_DiscardUnknown() {
	xxx_messageInfo_Instance.DiscardUnknown(m)
}

var xxx_messageInfo_Instance proto.InternalMessageInfo

func (m *Instance) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Instance) GetAppName() string {
	if m != nil {
		return m.AppName
	}
	return ""
}

func (m *Instance) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *Instance) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *Instance) GetCreatedTime() int64 {
	if m != nil {
		return m.CreatedTime
	}
	return 0
}

func (m *Instance) GetUpdatedTime() int64 {
	if m != nil {
		return m.UpdatedTime
	}
	return 0
}

type Item struct {
	Id                   int64           `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	NamespaceId          int64           `protobuf:"varint,2,opt,name=namespaceId,proto3" json:"namespaceId,omitempty"`
	Key                  string          `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Value                string          `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	Comment              string          `protobuf:"bytes,5,opt,name=comment,proto3" json:"comment,omitempty"`
	LineNo               int32           `protobuf:"varint,6,opt,name=lineNo,proto3" json:"lineNo,omitempty"`
	UpdateType           Item_UpdateType `protobuf:"varint,7,opt,name=updateType,proto3,enum=Item_UpdateType" json:"updateType,omitempty"`
	CreatedTime          int64           `protobuf:"varint,8,opt,name=createdTime,proto3" json:"createdTime,omitempty"`
	UpdatedTime          int64           `protobuf:"varint,9,opt,name=updatedTime,proto3" json:"updatedTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}
func (*Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_aca970399dcefd35, []int{3}
}

func (m *Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Item.Unmarshal(m, b)
}
func (m *Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Item.Marshal(b, m, deterministic)
}
func (m *Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Item.Merge(m, src)
}
func (m *Item) XXX_Size() int {
	return xxx_messageInfo_Item.Size(m)
}
func (m *Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Item proto.InternalMessageInfo

func (m *Item) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Item) GetNamespaceId() int64 {
	if m != nil {
		return m.NamespaceId
	}
	return 0
}

func (m *Item) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Item) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Item) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *Item) GetLineNo() int32 {
	if m != nil {
		return m.LineNo
	}
	return 0
}

func (m *Item) GetUpdateType() Item_UpdateType {
	if m != nil {
		return m.UpdateType
	}
	return Item_DELETE
}

func (m *Item) GetCreatedTime() int64 {
	if m != nil {
		return m.CreatedTime
	}
	return 0
}

func (m *Item) GetUpdatedTime() int64 {
	if m != nil {
		return m.UpdatedTime
	}
	return 0
}

type Namespace struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AppName              string   `protobuf:"bytes,2,opt,name=appName,proto3" json:"appName,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	ClusterName          string   `protobuf:"bytes,4,opt,name=clusterName,proto3" json:"clusterName,omitempty"`
	CreatedTime          int64    `protobuf:"varint,5,opt,name=createdTime,proto3" json:"createdTime,omitempty"`
	UpdatedTime          int64    `protobuf:"varint,6,opt,name=updatedTime,proto3" json:"updatedTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Namespace) Reset()         { *m = Namespace{} }
func (m *Namespace) String() string { return proto.CompactTextString(m) }
func (*Namespace) ProtoMessage()    {}
func (*Namespace) Descriptor() ([]byte, []int) {
	return fileDescriptor_aca970399dcefd35, []int{4}
}

func (m *Namespace) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Namespace.Unmarshal(m, b)
}
func (m *Namespace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Namespace.Marshal(b, m, deterministic)
}
func (m *Namespace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Namespace.Merge(m, src)
}
func (m *Namespace) XXX_Size() int {
	return xxx_messageInfo_Namespace.Size(m)
}
func (m *Namespace) XXX_DiscardUnknown() {
	xxx_messageInfo_Namespace.DiscardUnknown(m)
}

var xxx_messageInfo_Namespace proto.InternalMessageInfo

func (m *Namespace) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Namespace) GetAppName() string {
	if m != nil {
		return m.AppName
	}
	return ""
}

func (m *Namespace) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Namespace) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *Namespace) GetCreatedTime() int64 {
	if m != nil {
		return m.CreatedTime
	}
	return 0
}

func (m *Namespace) GetUpdatedTime() int64 {
	if m != nil {
		return m.UpdatedTime
	}
	return 0
}

type EntryRequest struct {
	App                  *App     `protobuf:"bytes,1,opt,name=app,proto3" json:"app,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EntryRequest) Reset()         { *m = EntryRequest{} }
func (m *EntryRequest) String() string { return proto.CompactTextString(m) }
func (*EntryRequest) ProtoMessage()    {}
func (*EntryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aca970399dcefd35, []int{5}
}

func (m *EntryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntryRequest.Unmarshal(m, b)
}
func (m *EntryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntryRequest.Marshal(b, m, deterministic)
}
func (m *EntryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntryRequest.Merge(m, src)
}
func (m *EntryRequest) XXX_Size() int {
	return xxx_messageInfo_EntryRequest.Size(m)
}
func (m *EntryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EntryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EntryRequest proto.InternalMessageInfo

func (m *EntryRequest) GetApp() *App {
	if m != nil {
		return m.App
	}
	return nil
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_aca970399dcefd35, []int{6}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type EntryResponse struct {
	Error                *Error     `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	App                  *App       `protobuf:"bytes,3,opt,name=app,proto3" json:"app,omitempty"`
	AppList              []*App     `protobuf:"bytes,4,rep,name=appList,proto3" json:"appList,omitempty"`
	Cluster              *Cluster   `protobuf:"bytes,5,opt,name=cluster,proto3" json:"cluster,omitempty"`
	ClusterList          []*Cluster `protobuf:"bytes,6,rep,name=clusterList,proto3" json:"clusterList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *EntryResponse) Reset()         { *m = EntryResponse{} }
func (m *EntryResponse) String() string { return proto.CompactTextString(m) }
func (*EntryResponse) ProtoMessage()    {}
func (*EntryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_aca970399dcefd35, []int{7}
}

func (m *EntryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntryResponse.Unmarshal(m, b)
}
func (m *EntryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntryResponse.Marshal(b, m, deterministic)
}
func (m *EntryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntryResponse.Merge(m, src)
}
func (m *EntryResponse) XXX_Size() int {
	return xxx_messageInfo_EntryResponse.Size(m)
}
func (m *EntryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EntryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EntryResponse proto.InternalMessageInfo

func (m *EntryResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *EntryResponse) GetApp() *App {
	if m != nil {
		return m.App
	}
	return nil
}

func (m *EntryResponse) GetAppList() []*App {
	if m != nil {
		return m.AppList
	}
	return nil
}

func (m *EntryResponse) GetCluster() *Cluster {
	if m != nil {
		return m.Cluster
	}
	return nil
}

func (m *EntryResponse) GetClusterList() []*Cluster {
	if m != nil {
		return m.ClusterList
	}
	return nil
}

func init() {
	proto.RegisterEnum("Item_UpdateType", Item_UpdateType_name, Item_UpdateType_value)
	proto.RegisterType((*App)(nil), "App")
	proto.RegisterType((*Cluster)(nil), "Cluster")
	proto.RegisterType((*Instance)(nil), "Instance")
	proto.RegisterType((*Item)(nil), "Item")
	proto.RegisterType((*Namespace)(nil), "Namespace")
	proto.RegisterType((*EntryRequest)(nil), "EntryRequest")
	proto.RegisterType((*Error)(nil), "Error")
	proto.RegisterType((*EntryResponse)(nil), "EntryResponse")
}

func init() { proto.RegisterFile("entry/entry.proto", fileDescriptor_aca970399dcefd35) }

var fileDescriptor_aca970399dcefd35 = []byte{
	// 615 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x95, 0xdd, 0x8a, 0xd3, 0x40,
	0x14, 0xc7, 0x37, 0x9f, 0x6d, 0x4e, 0x76, 0x6b, 0x1d, 0x44, 0x82, 0x88, 0x84, 0x5c, 0x48, 0xfd,
	0xd8, 0x6c, 0x89, 0x4f, 0x50, 0xb6, 0x61, 0x29, 0x2c, 0x65, 0x19, 0xbb, 0x0f, 0x10, 0x9b, 0x83,
	0x04, 0xf3, 0x31, 0x26, 0x69, 0xa1, 0xcf, 0xe0, 0x13, 0xf8, 0x0a, 0xde, 0x78, 0xe5, 0x8d, 0x4f,
	0x27, 0x33, 0x49, 0xda, 0xd8, 0x5c, 0x24, 0x82, 0x37, 0xcb, 0xcc, 0xc9, 0x7f, 0xce, 0xf9, 0x9d,
	0xff, 0xd9, 0x99, 0xc2, 0x53, 0x4c, 0xcb, 0xfc, 0x70, 0x23, 0xfe, 0xba, 0x2c, 0xcf, 0xca, 0xcc,
	0x49, 0x40, 0x59, 0x30, 0x46, 0x26, 0x20, 0x47, 0xa1, 0x25, 0xd9, 0xd2, 0x4c, 0xa1, 0x72, 0x14,
	0x12, 0x02, 0x6a, 0x1a, 0x24, 0x68, 0xc9, 0xb6, 0x34, 0x33, 0xa8, 0x58, 0x13, 0x1b, 0xcc, 0x6d,
	0x8e, 0x41, 0x89, 0xe1, 0x26, 0x4a, 0xd0, 0x52, 0x84, 0xb8, 0x1d, 0xe2, 0x8a, 0x1d, 0x0b, 0x8f,
	0x0a, 0xb5, 0x52, 0xb4, 0x42, 0xce, 0x37, 0x09, 0x46, 0xb7, 0xf1, 0xae, 0x28, 0x31, 0x1f, 0x54,
	0xd3, 0x82, 0x51, 0xc0, 0xd8, 0x3a, 0xa8, 0xeb, 0x19, 0xb4, 0xd9, 0x9e, 0xd3, 0xa8, 0xbd, 0x34,
	0x5a, 0x97, 0xe6, 0x87, 0x04, 0xe3, 0x55, 0x5a, 0x94, 0x41, 0xba, 0xc5, 0x0e, 0x4e, 0xab, 0xb4,
	0xdc, 0x2d, 0x5d, 0xf5, 0xd0, 0x02, 0x6b, 0x87, 0x44, 0x2e, 0x26, 0x98, 0x0c, 0x2a, 0x47, 0xec,
	0x1c, 0x56, 0xeb, 0x85, 0xd5, 0xbb, 0xb0, 0xbf, 0x65, 0x50, 0x57, 0x25, 0x26, 0x1d, 0x50, 0x1b,
	0x4c, 0xee, 0x55, 0xc1, 0x82, 0x2d, 0xae, 0x42, 0x01, 0xab, 0xd0, 0x76, 0x88, 0x4c, 0x41, 0xf9,
	0x82, 0x87, 0x1a, 0x94, 0x2f, 0xc9, 0x33, 0xd0, 0xf6, 0x41, 0xbc, 0xc3, 0x9a, 0xb1, 0xda, 0xf0,
	0x96, 0xb7, 0x59, 0x92, 0x60, 0x5a, 0x0a, 0x44, 0x83, 0x36, 0x5b, 0xf2, 0x1c, 0xf4, 0x38, 0x4a,
	0x71, 0x9d, 0x09, 0x32, 0x8d, 0xd6, 0x3b, 0x32, 0x07, 0xa8, 0x18, 0x37, 0x07, 0x86, 0xd6, 0xc8,
	0x96, 0x66, 0x13, 0x6f, 0xea, 0x72, 0x4c, 0xf7, 0xf1, 0x18, 0xa7, 0x2d, 0xcd, 0xb9, 0x15, 0xe3,
	0x5e, 0x2b, 0x8c, 0xae, 0x15, 0x73, 0x80, 0x53, 0x76, 0x02, 0xa0, 0x2f, 0xfd, 0x7b, 0x7f, 0xe3,
	0x4f, 0x2f, 0xf8, 0xfa, 0xf1, 0x61, 0xb9, 0xd8, 0xf8, 0x53, 0x89, 0xaf, 0x57, 0xeb, 0x8f, 0x3e,
	0xdd, 0x4c, 0x65, 0xe7, 0xa7, 0x04, 0xc6, 0xba, 0x71, 0xe4, 0x1f, 0x46, 0xdd, 0xfc, 0x4f, 0x2a,
	0x67, 0xf7, 0xa0, 0x35, 0x7e, 0xb5, 0x3b, 0xfe, 0xff, 0x31, 0xee, 0xd7, 0x70, 0xe9, 0xf3, 0x7b,
	0x4a, 0xf1, 0xeb, 0x0e, 0x0b, 0x3e, 0x01, 0x25, 0x60, 0x4c, 0x40, 0x9b, 0x9e, 0xea, 0x2e, 0x18,
	0xa3, 0x3c, 0xe0, 0x5c, 0x83, 0xe6, 0xe7, 0x79, 0x96, 0x73, 0xd4, 0x6d, 0x16, 0xa2, 0x50, 0x68,
	0x54, 0xac, 0xf9, 0xe0, 0x93, 0xe2, 0x73, 0xdd, 0x14, 0x5f, 0x3a, 0xbf, 0x24, 0xb8, 0xaa, 0xf3,
	0x16, 0x2c, 0x4b, 0x0b, 0x24, 0x2f, 0x41, 0x43, 0x9e, 0x40, 0xa8, 0x4c, 0x4f, 0x77, 0x45, 0x3a,
	0x5a, 0x05, 0x9b, 0xb2, 0xca, 0x59, 0x59, 0xf2, 0x4a, 0x58, 0x76, 0x1f, 0x15, 0xa5, 0xa5, 0xda,
	0xca, 0xf1, 0x5b, 0x13, 0x24, 0x0e, 0x8c, 0x6a, 0x47, 0x44, 0xfb, 0xa6, 0x37, 0x76, 0xeb, 0x7b,
	0x4f, 0x9b, 0x0f, 0xe4, 0xed, 0xd1, 0x48, 0x91, 0x47, 0x17, 0x79, 0x4e, 0xba, 0xf6, 0x47, 0xef,
	0xbb, 0x06, 0x9a, 0xe0, 0x26, 0xef, 0xc1, 0xb8, 0x15, 0x4e, 0xf2, 0x77, 0xeb, 0xca, 0x6d, 0x9b,
	0xf4, 0x62, 0xe2, 0xfe, 0xd5, 0x9b, 0x73, 0x41, 0xde, 0xc1, 0x98, 0x9f, 0x5f, 0x30, 0x56, 0xf4,
	0x8b, 0xdf, 0x80, 0x7e, 0x87, 0xe5, 0xa0, 0xbc, 0x73, 0xb8, 0xaa, 0x28, 0x9a, 0xd7, 0xac, 0xf7,
	0xc4, 0x35, 0xc0, 0x1d, 0x96, 0x83, 0xe5, 0x37, 0x70, 0xc9, 0xc1, 0x6b, 0xfd, 0x00, 0xf8, 0xa3,
	0x2f, 0x7e, 0xba, 0x1f, 0xda, 0xea, 0x20, 0x69, 0x6d, 0xa1, 0x9f, 0xee, 0x8b, 0x41, 0xbe, 0x3c,
	0xec, 0xe2, 0xb8, 0x79, 0x55, 0x07, 0x9c, 0xf0, 0xe0, 0x49, 0xc5, 0x7d, 0xba, 0x9f, 0x43, 0xbc,
	0xac, 0xce, 0x88, 0x07, 0x71, 0x88, 0xbc, 0x7a, 0x2f, 0x86, 0xc9, 0x5d, 0x30, 0x97, 0x18, 0x63,
	0x25, 0xef, 0xef, 0xe0, 0x93, 0x2e, 0x7e, 0x4a, 0x3f, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0x29,
	0x49, 0xc6, 0x03, 0x5f, 0x07, 0x00, 0x00,
}
