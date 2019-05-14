// Code generated by protoc-gen-go. DO NOT EDIT.
// source: heartbeat.proto

package mustang

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

type HeartbeatReq struct {
	Timestamp            int64    `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	AppName              string   `protobuf:"bytes,3,opt,name=appName,proto3" json:"appName,omitempty"`
	ExecutePath          string   `protobuf:"bytes,4,opt,name=executePath,proto3" json:"executePath,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeartbeatReq) Reset()         { *m = HeartbeatReq{} }
func (m *HeartbeatReq) String() string { return proto.CompactTextString(m) }
func (*HeartbeatReq) ProtoMessage()    {}
func (*HeartbeatReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c667767fb9826a9, []int{0}
}

func (m *HeartbeatReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartbeatReq.Unmarshal(m, b)
}
func (m *HeartbeatReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartbeatReq.Marshal(b, m, deterministic)
}
func (m *HeartbeatReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartbeatReq.Merge(m, src)
}
func (m *HeartbeatReq) XXX_Size() int {
	return xxx_messageInfo_HeartbeatReq.Size(m)
}
func (m *HeartbeatReq) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartbeatReq.DiscardUnknown(m)
}

var xxx_messageInfo_HeartbeatReq proto.InternalMessageInfo

func (m *HeartbeatReq) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *HeartbeatReq) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *HeartbeatReq) GetAppName() string {
	if m != nil {
		return m.AppName
	}
	return ""
}

func (m *HeartbeatReq) GetExecutePath() string {
	if m != nil {
		return m.ExecutePath
	}
	return ""
}

type HeartbeatResp struct {
	Timestamp            int64    `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeartbeatResp) Reset()         { *m = HeartbeatResp{} }
func (m *HeartbeatResp) String() string { return proto.CompactTextString(m) }
func (*HeartbeatResp) ProtoMessage()    {}
func (*HeartbeatResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c667767fb9826a9, []int{1}
}

func (m *HeartbeatResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartbeatResp.Unmarshal(m, b)
}
func (m *HeartbeatResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartbeatResp.Marshal(b, m, deterministic)
}
func (m *HeartbeatResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartbeatResp.Merge(m, src)
}
func (m *HeartbeatResp) XXX_Size() int {
	return xxx_messageInfo_HeartbeatResp.Size(m)
}
func (m *HeartbeatResp) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartbeatResp.DiscardUnknown(m)
}

var xxx_messageInfo_HeartbeatResp proto.InternalMessageInfo

func (m *HeartbeatResp) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *HeartbeatResp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type RegistryReq struct {
	AppName                string   `protobuf:"bytes,1,opt,name=appName,proto3" json:"appName,omitempty"`
	EnableSelfPreservation bool     `protobuf:"varint,2,opt,name=enableSelfPreservation,proto3" json:"enableSelfPreservation,omitempty"`
	ExecutePath            string   `protobuf:"bytes,3,opt,name=executePath,proto3" json:"executePath,omitempty"`
	HeartbeatInterval      int64    `protobuf:"varint,4,opt,name=heartbeatInterval,proto3" json:"heartbeatInterval,omitempty"`
	ConnectTimeout         int64    `protobuf:"varint,5,opt,name=connectTimeout,proto3" json:"connectTimeout,omitempty"`
	Description            string   `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *RegistryReq) Reset()         { *m = RegistryReq{} }
func (m *RegistryReq) String() string { return proto.CompactTextString(m) }
func (*RegistryReq) ProtoMessage()    {}
func (*RegistryReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c667767fb9826a9, []int{2}
}

func (m *RegistryReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegistryReq.Unmarshal(m, b)
}
func (m *RegistryReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegistryReq.Marshal(b, m, deterministic)
}
func (m *RegistryReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegistryReq.Merge(m, src)
}
func (m *RegistryReq) XXX_Size() int {
	return xxx_messageInfo_RegistryReq.Size(m)
}
func (m *RegistryReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RegistryReq.DiscardUnknown(m)
}

var xxx_messageInfo_RegistryReq proto.InternalMessageInfo

func (m *RegistryReq) GetAppName() string {
	if m != nil {
		return m.AppName
	}
	return ""
}

func (m *RegistryReq) GetEnableSelfPreservation() bool {
	if m != nil {
		return m.EnableSelfPreservation
	}
	return false
}

func (m *RegistryReq) GetExecutePath() string {
	if m != nil {
		return m.ExecutePath
	}
	return ""
}

func (m *RegistryReq) GetHeartbeatInterval() int64 {
	if m != nil {
		return m.HeartbeatInterval
	}
	return 0
}

func (m *RegistryReq) GetConnectTimeout() int64 {
	if m != nil {
		return m.ConnectTimeout
	}
	return 0
}

func (m *RegistryReq) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type RegistryResp struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegistryResp) Reset()         { *m = RegistryResp{} }
func (m *RegistryResp) String() string { return proto.CompactTextString(m) }
func (*RegistryResp) ProtoMessage()    {}
func (*RegistryResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c667767fb9826a9, []int{3}
}

func (m *RegistryResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegistryResp.Unmarshal(m, b)
}
func (m *RegistryResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegistryResp.Marshal(b, m, deterministic)
}
func (m *RegistryResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegistryResp.Merge(m, src)
}
func (m *RegistryResp) XXX_Size() int {
	return xxx_messageInfo_RegistryResp.Size(m)
}
func (m *RegistryResp) XXX_DiscardUnknown() {
	xxx_messageInfo_RegistryResp.DiscardUnknown(m)
}

var xxx_messageInfo_RegistryResp proto.InternalMessageInfo

func (m *RegistryResp) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*HeartbeatReq)(nil), "mustang.HeartbeatReq")
	proto.RegisterType((*HeartbeatResp)(nil), "mustang.HeartbeatResp")
	proto.RegisterType((*RegistryReq)(nil), "mustang.RegistryReq")
	proto.RegisterType((*RegistryResp)(nil), "mustang.RegistryResp")
}

func init() { proto.RegisterFile("heartbeat.proto", fileDescriptor_3c667767fb9826a9) }

var fileDescriptor_3c667767fb9826a9 = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x4b, 0x4f, 0xf3, 0x30,
	0x10, 0x54, 0xbe, 0xf4, 0xeb, 0x63, 0x5b, 0x5e, 0x16, 0xad, 0xa2, 0x8a, 0x43, 0x95, 0x03, 0xea,
	0x01, 0xf5, 0x00, 0x02, 0x8e, 0x5c, 0xe1, 0x82, 0x2a, 0x97, 0x3f, 0xe0, 0xa6, 0x4b, 0x1b, 0xa9,
	0x71, 0x8c, 0x77, 0x83, 0x78, 0xfc, 0x72, 0x6e, 0x28, 0x86, 0xb4, 0xa6, 0x2d, 0x07, 0x6e, 0xde,
	0x99, 0xb5, 0x67, 0x66, 0xbd, 0x70, 0xb0, 0x40, 0x65, 0x79, 0x8a, 0x8a, 0x47, 0xc6, 0xe6, 0x9c,
	0x8b, 0x46, 0x56, 0x10, 0x2b, 0x3d, 0x8f, 0xdf, 0xa0, 0x73, 0x5b, 0x71, 0x12, 0x9f, 0xc4, 0x09,
	0xb4, 0x38, 0xcd, 0x90, 0x58, 0x65, 0x26, 0x0a, 0x06, 0xc1, 0x30, 0x94, 0x6b, 0x40, 0x1c, 0x42,
	0x98, 0xd1, 0x3c, 0xfa, 0x37, 0x08, 0x86, 0x2d, 0x59, 0x1e, 0x45, 0x04, 0x0d, 0x65, 0xcc, 0xbd,
	0xca, 0x30, 0x0a, 0x1d, 0x5a, 0x95, 0x62, 0x00, 0x6d, 0x7c, 0xc1, 0xa4, 0x60, 0x1c, 0x2b, 0x5e,
	0x44, 0x35, 0xc7, 0xfa, 0x50, 0x7c, 0x03, 0x7b, 0x9e, 0x36, 0x99, 0xbf, 0x8a, 0xc7, 0x1f, 0x01,
	0xb4, 0x25, 0xce, 0x53, 0x62, 0xfb, 0x5a, 0x9a, 0xf7, 0xcc, 0x04, 0x3f, 0xcd, 0x5c, 0x41, 0x0f,
	0xb5, 0x9a, 0x2e, 0x71, 0x82, 0xcb, 0xc7, 0xb1, 0x45, 0x42, 0xfb, 0xac, 0x38, 0xcd, 0xb5, 0x7b,
	0xae, 0x29, 0x7f, 0x61, 0x37, 0x43, 0x84, 0x5b, 0x21, 0xc4, 0x19, 0x1c, 0xad, 0x86, 0x7b, 0xa7,
	0xb9, 0xbc, 0xb9, 0x74, 0x61, 0x43, 0xb9, 0x4d, 0x88, 0x53, 0xd8, 0x4f, 0x72, 0xad, 0x31, 0xe1,
	0x87, 0x34, 0xc3, 0xbc, 0xe0, 0xe8, 0xbf, 0x6b, 0xdd, 0x40, 0x4b, 0xdd, 0x19, 0x52, 0x62, 0x53,
	0xe3, 0x4c, 0xd6, 0xbf, 0x74, 0x3d, 0x28, 0x1e, 0x42, 0x67, 0x1d, 0x9d, 0x4c, 0x99, 0x9d, 0x8a,
	0x24, 0x41, 0x22, 0x97, 0xbd, 0x29, 0xab, 0xf2, 0xfc, 0x1d, 0x5a, 0xab, 0x31, 0x8b, 0x4b, 0xa8,
	0x4d, 0x50, 0xcf, 0x44, 0x77, 0xf4, 0xbd, 0x01, 0x23, 0xff, 0xfb, 0xfb, 0xbd, 0x5d, 0x30, 0x19,
	0x71, 0x0d, 0xcd, 0x4a, 0x4d, 0x1c, 0xaf, 0x7a, 0xbc, 0xd9, 0xf7, 0xbb, 0x3b, 0x50, 0x32, 0xd3,
	0xba, 0xdb, 0xb7, 0x8b, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe8, 0xda, 0x05, 0x9b, 0x82, 0x02,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HeartbeatClient is the client API for Heartbeat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HeartbeatClient interface {
	Send(ctx context.Context, in *HeartbeatReq, opts ...grpc.CallOption) (*HeartbeatResp, error)
	Registry(ctx context.Context, in *RegistryReq, opts ...grpc.CallOption) (*RegistryResp, error)
}

type heartbeatClient struct {
	cc *grpc.ClientConn
}

func NewHeartbeatClient(cc *grpc.ClientConn) HeartbeatClient {
	return &heartbeatClient{cc}
}

func (c *heartbeatClient) Send(ctx context.Context, in *HeartbeatReq, opts ...grpc.CallOption) (*HeartbeatResp, error) {
	out := new(HeartbeatResp)
	err := c.cc.Invoke(ctx, "/mustang.Heartbeat/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *heartbeatClient) Registry(ctx context.Context, in *RegistryReq, opts ...grpc.CallOption) (*RegistryResp, error) {
	out := new(RegistryResp)
	err := c.cc.Invoke(ctx, "/mustang.Heartbeat/Registry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HeartbeatServer is the server API for Heartbeat service.
type HeartbeatServer interface {
	Send(context.Context, *HeartbeatReq) (*HeartbeatResp, error)
	Registry(context.Context, *RegistryReq) (*RegistryResp, error)
}

func RegisterHeartbeatServer(s *grpc.Server, srv HeartbeatServer) {
	s.RegisterService(&_Heartbeat_serviceDesc, srv)
}

func _Heartbeat_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartbeatReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeartbeatServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mustang.Heartbeat/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeartbeatServer).Send(ctx, req.(*HeartbeatReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Heartbeat_Registry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeartbeatServer).Registry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mustang.Heartbeat/Registry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeartbeatServer).Registry(ctx, req.(*RegistryReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Heartbeat_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mustang.Heartbeat",
	HandlerType: (*HeartbeatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _Heartbeat_Send_Handler,
		},
		{
			MethodName: "Registry",
			Handler:    _Heartbeat_Registry_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "heartbeat.proto",
}
