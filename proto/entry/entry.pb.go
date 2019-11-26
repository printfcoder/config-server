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

type UpdateType int32

const (
	UpdateType_DELETE UpdateType = 0
	UpdateType_UPDATE UpdateType = 1
	UpdateType_INSERT UpdateType = 2
)

var UpdateType_name = map[int32]string{
	0: "DELETE",
	1: "UPDATE",
	2: "INSERT",
}

var UpdateType_value = map[string]int32{
	"DELETE": 0,
	"UPDATE": 1,
	"INSERT": 2,
}

func (x UpdateType) String() string {
	return proto.EnumName(UpdateType_name, int32(x))
}

func (UpdateType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_aca970399dcefd35, []int{0}
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
	Id                   int64      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	NamespaceId          int64      `protobuf:"varint,2,opt,name=namespaceId,proto3" json:"namespaceId,omitempty"`
	Key                  string     `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Value                string     `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	Comment              string     `protobuf:"bytes,5,opt,name=comment,proto3" json:"comment,omitempty"`
	LineNo               int32      `protobuf:"varint,6,opt,name=lineNo,proto3" json:"lineNo,omitempty"`
	UpdateType           UpdateType `protobuf:"varint,7,opt,name=updateType,proto3,enum=UpdateType" json:"updateType,omitempty"`
	CreatedTime          int64      `protobuf:"varint,8,opt,name=createdTime,proto3" json:"createdTime,omitempty"`
	UpdatedTime          int64      `protobuf:"varint,9,opt,name=updatedTime,proto3" json:"updatedTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
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

func (m *Item) GetUpdateType() UpdateType {
	if m != nil {
		return m.UpdateType
	}
	return UpdateType_DELETE
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
	proto.RegisterEnum("UpdateType", UpdateType_name, UpdateType_value)
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
	// 586 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x95, 0xcf, 0x6a, 0xdb, 0x4e,
	0x10, 0xc7, 0xa3, 0xbf, 0xb6, 0x46, 0xb1, 0x7f, 0xfe, 0x2d, 0xa5, 0x88, 0x52, 0x8a, 0xd0, 0xa1,
	0xb8, 0x49, 0xa3, 0x04, 0xf5, 0x09, 0x4c, 0x22, 0x82, 0x21, 0x98, 0xb0, 0x75, 0x1e, 0x40, 0xb5,
	0x86, 0x22, 0x6a, 0x49, 0x5b, 0x49, 0x2e, 0xf8, 0xdc, 0x63, 0xdf, 0xa4, 0x97, 0x9e, 0xf2, 0x7e,
	0x65, 0x57, 0x92, 0xbd, 0xb5, 0x0e, 0x52, 0xa1, 0x17, 0x33, 0x3b, 0x3b, 0x9a, 0xf9, 0xcc, 0x77,
	0x76, 0xd7, 0xf0, 0x3f, 0x66, 0x55, 0xb1, 0xbf, 0x16, 0xbf, 0x3e, 0x2b, 0xf2, 0x2a, 0xf7, 0x52,
	0xd0, 0x16, 0x8c, 0x91, 0x29, 0xa8, 0x49, 0xec, 0x28, 0xae, 0x32, 0xd7, 0xa8, 0x9a, 0xc4, 0x84,
	0x80, 0x9e, 0x45, 0x29, 0x3a, 0xaa, 0xab, 0xcc, 0x2d, 0x2a, 0x6c, 0xe2, 0x82, 0xbd, 0x29, 0x30,
	0xaa, 0x30, 0x5e, 0x27, 0x29, 0x3a, 0x9a, 0x08, 0x96, 0x5d, 0x3c, 0x62, 0xc7, 0xe2, 0x43, 0x84,
	0x5e, 0x47, 0x48, 0x2e, 0xef, 0x87, 0x02, 0xa3, 0xdb, 0xed, 0xae, 0xac, 0xb0, 0x18, 0x54, 0xd3,
	0x81, 0x51, 0xc4, 0xd8, 0x2a, 0x6a, 0xea, 0x59, 0xb4, 0x5d, 0x9e, 0xd2, 0xe8, 0xbd, 0x34, 0x46,
	0x97, 0xe6, 0xa7, 0x02, 0xe3, 0x65, 0x56, 0x56, 0x51, 0xb6, 0xc1, 0x0e, 0x8e, 0x54, 0x5a, 0xed,
	0x96, 0xae, 0x7b, 0x90, 0xc0, 0x64, 0x97, 0xc8, 0xc5, 0x04, 0x93, 0x45, 0xd5, 0x84, 0x9d, 0xc2,
	0x1a, 0xbd, 0xb0, 0x66, 0x17, 0xf6, 0xbb, 0x0a, 0xfa, 0xb2, 0xc2, 0xb4, 0x03, 0xea, 0x82, 0xcd,
	0xb5, 0x2a, 0x59, 0xb4, 0xc1, 0x65, 0x2c, 0x60, 0x35, 0x2a, 0xbb, 0xc8, 0x0c, 0xb4, 0x2f, 0xb8,
	0x6f, 0x40, 0xb9, 0x49, 0x5e, 0x80, 0xf1, 0x2d, 0xda, 0xee, 0xb0, 0x61, 0xac, 0x17, 0xbc, 0xe5,
	0x4d, 0x9e, 0xa6, 0x98, 0x55, 0x02, 0xd1, 0xa2, 0xed, 0x92, 0xbc, 0x04, 0x73, 0x9b, 0x64, 0xb8,
	0xca, 0x05, 0x99, 0x41, 0x9b, 0x15, 0xb9, 0x04, 0xa8, 0x19, 0xd7, 0x7b, 0x86, 0xce, 0xc8, 0x55,
	0xe6, 0xd3, 0xc0, 0xf6, 0x9f, 0x0e, 0x2e, 0x2a, 0x6d, 0x9f, 0xaa, 0x30, 0xee, 0x55, 0xc1, 0xea,
	0xaa, 0xf0, 0x4b, 0x01, 0x6b, 0xd5, 0xb6, 0xf6, 0x17, 0x33, 0x6b, 0x0f, 0x97, 0x76, 0x72, 0xa0,
	0xa5, 0x39, 0xea, 0xdd, 0x39, 0xfe, 0x8b, 0xb9, 0xbd, 0x85, 0xf3, 0x90, 0x5f, 0x38, 0x8a, 0x5f,
	0x77, 0x58, 0x72, 0x29, 0xb5, 0x88, 0x31, 0x01, 0x6d, 0x07, 0xba, 0xbf, 0x60, 0x8c, 0x72, 0x87,
	0x77, 0x05, 0x46, 0x58, 0x14, 0x79, 0xc1, 0x51, 0x37, 0x79, 0x8c, 0x22, 0xc2, 0xa0, 0xc2, 0xe6,
	0x13, 0x4c, 0xcb, 0xcf, 0x4d, 0x53, 0xdc, 0xf4, 0x9e, 0x15, 0x98, 0x34, 0x79, 0x4b, 0x96, 0x67,
	0x25, 0x92, 0xd7, 0x60, 0x20, 0x4f, 0x20, 0xa2, 0xec, 0xc0, 0xf4, 0x45, 0x3a, 0x5a, 0x3b, 0xdb,
	0xb2, 0xda, 0x49, 0x59, 0xf2, 0x46, 0x48, 0xf6, 0x90, 0x94, 0x95, 0xa3, 0xbb, 0xda, 0x61, 0xaf,
	0x75, 0x12, 0x0f, 0x46, 0x8d, 0x22, 0xa2, 0x7d, 0x3b, 0x18, 0xfb, 0xcd, 0x05, 0xa6, 0xed, 0x06,
	0xb9, 0x38, 0x08, 0x29, 0xf2, 0x98, 0x22, 0xcf, 0x31, 0x4e, 0xde, 0xbc, 0xb8, 0x01, 0x38, 0x1e,
	0x0f, 0x02, 0x60, 0xde, 0x85, 0x0f, 0xe1, 0x3a, 0x9c, 0x9d, 0x71, 0xfb, 0xe9, 0xf1, 0x6e, 0xb1,
	0x0e, 0x67, 0x0a, 0xb7, 0x97, 0xab, 0x8f, 0x21, 0x5d, 0xcf, 0xd4, 0xe0, 0x59, 0x03, 0x43, 0x74,
	0x4a, 0xde, 0x83, 0x75, 0x2b, 0xb4, 0xe7, 0x4f, 0xd6, 0xc4, 0x97, 0x65, 0x7d, 0x35, 0xf5, 0xff,
	0x50, 0xc3, 0x3b, 0x23, 0x97, 0x30, 0xe6, 0x15, 0x17, 0x8c, 0x95, 0xfd, 0xc1, 0xef, 0xc0, 0xbc,
	0xc7, 0x6a, 0x50, 0xde, 0x1b, 0x98, 0xd4, 0x14, 0xed, 0x43, 0xd6, 0xfb, 0xc5, 0x15, 0xc0, 0x3d,
	0x56, 0x83, 0xc3, 0xaf, 0xe1, 0x9c, 0x83, 0x37, 0xf1, 0xe5, 0x20, 0xa2, 0xc7, 0xdd, 0x76, 0xdb,
	0x3e, 0x65, 0x03, 0xbe, 0x08, 0xe0, 0xbf, 0xba, 0x87, 0xe3, 0x5d, 0xea, 0xfd, 0xc6, 0x07, 0xbb,
	0x9e, 0x1c, 0x7f, 0x85, 0xfa, 0x6b, 0x7c, 0x32, 0xc5, 0x3f, 0xcc, 0x87, 0xdf, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xf6, 0x77, 0x48, 0x40, 0x76, 0x06, 0x00, 0x00,
}
