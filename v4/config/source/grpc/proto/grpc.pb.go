// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc.proto

/*
Package grpc is a generated protocol buffer package.

It is generated from these files:

	grpc.proto

It has these top-level messages:

	ChangeSet
	ReadRequest
	ReadResponse
	WatchRequest
	WatchResponse
*/
package grpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc1 "google.golang.org/grpc"
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

type ChangeSet struct {
	Data      []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Checksum  string `protobuf:"bytes,2,opt,name=checksum" json:"checksum,omitempty"`
	Format    string `protobuf:"bytes,3,opt,name=format" json:"format,omitempty"`
	Source    string `protobuf:"bytes,4,opt,name=source" json:"source,omitempty"`
	Timestamp int64  `protobuf:"varint,5,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *ChangeSet) Reset()                    { *m = ChangeSet{} }
func (m *ChangeSet) String() string            { return proto.CompactTextString(m) }
func (*ChangeSet) ProtoMessage()               {}
func (*ChangeSet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ChangeSet) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *ChangeSet) GetChecksum() string {
	if m != nil {
		return m.Checksum
	}
	return ""
}

func (m *ChangeSet) GetFormat() string {
	if m != nil {
		return m.Format
	}
	return ""
}

func (m *ChangeSet) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *ChangeSet) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type ReadRequest struct {
	Path string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
}

func (m *ReadRequest) Reset()                    { *m = ReadRequest{} }
func (m *ReadRequest) String() string            { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()               {}
func (*ReadRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ReadRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type ReadResponse struct {
	ChangeSet *ChangeSet `protobuf:"bytes,1,opt,name=change_set,json=changeSet" json:"change_set,omitempty"`
}

func (m *ReadResponse) Reset()                    { *m = ReadResponse{} }
func (m *ReadResponse) String() string            { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()               {}
func (*ReadResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ReadResponse) GetChangeSet() *ChangeSet {
	if m != nil {
		return m.ChangeSet
	}
	return nil
}

type WatchRequest struct {
	Path string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
}

func (m *WatchRequest) Reset()                    { *m = WatchRequest{} }
func (m *WatchRequest) String() string            { return proto.CompactTextString(m) }
func (*WatchRequest) ProtoMessage()               {}
func (*WatchRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *WatchRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type WatchResponse struct {
	ChangeSet *ChangeSet `protobuf:"bytes,1,opt,name=change_set,json=changeSet" json:"change_set,omitempty"`
}

func (m *WatchResponse) Reset()                    { *m = WatchResponse{} }
func (m *WatchResponse) String() string            { return proto.CompactTextString(m) }
func (*WatchResponse) ProtoMessage()               {}
func (*WatchResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *WatchResponse) GetChangeSet() *ChangeSet {
	if m != nil {
		return m.ChangeSet
	}
	return nil
}

func init() {
	proto.RegisterType((*ChangeSet)(nil), "ChangeSet")
	proto.RegisterType((*ReadRequest)(nil), "ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "ReadResponse")
	proto.RegisterType((*WatchRequest)(nil), "WatchRequest")
	proto.RegisterType((*WatchResponse)(nil), "WatchResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc1.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc1.SupportPackageIsVersion4

// Client API for Source service

type SourceClient interface {
	Read(ctx context.Context, in *ReadRequest, opts ...grpc1.CallOption) (*ReadResponse, error)
	Watch(ctx context.Context, in *WatchRequest, opts ...grpc1.CallOption) (Source_WatchClient, error)
}

type sourceClient struct {
	cc *grpc1.ClientConn
}

func NewSourceClient(cc *grpc1.ClientConn) SourceClient {
	return &sourceClient{cc}
}

func (c *sourceClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc1.CallOption) (*ReadResponse, error) {
	out := new(ReadResponse)
	err := grpc1.Invoke(ctx, "/Source/Read", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sourceClient) Watch(ctx context.Context, in *WatchRequest, opts ...grpc1.CallOption) (Source_WatchClient, error) {
	stream, err := grpc1.NewClientStream(ctx, &_Source_serviceDesc.Streams[0], c.cc, "/Source/Watch", opts...)
	if err != nil {
		return nil, err
	}
	x := &sourceWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Source_WatchClient interface {
	Recv() (*WatchResponse, error)
	grpc1.ClientStream
}

type sourceWatchClient struct {
	grpc1.ClientStream
}

func (x *sourceWatchClient) Recv() (*WatchResponse, error) {
	m := new(WatchResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Source service

type SourceServer interface {
	Read(context.Context, *ReadRequest) (*ReadResponse, error)
	Watch(*WatchRequest, Source_WatchServer) error
}

func RegisterSourceServer(s *grpc1.Server, srv SourceServer) {
	s.RegisterService(&_Source_serviceDesc, srv)
}

func _Source_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SourceServer).Read(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Source/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SourceServer).Read(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Source_Watch_Handler(srv interface{}, stream grpc1.ServerStream) error {
	m := new(WatchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SourceServer).Watch(m, &sourceWatchServer{stream})
}

type Source_WatchServer interface {
	Send(*WatchResponse) error
	grpc1.ServerStream
}

type sourceWatchServer struct {
	grpc1.ServerStream
}

func (x *sourceWatchServer) Send(m *WatchResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Source_serviceDesc = grpc1.ServiceDesc{
	ServiceName: "Source",
	HandlerType: (*SourceServer)(nil),
	Methods: []grpc1.MethodDesc{
		{
			MethodName: "Read",
			Handler:    _Source_Read_Handler,
		},
	},
	Streams: []grpc1.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _Source_Watch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "grpc.proto",
}

func init() { proto.RegisterFile("grpc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xb1, 0x4e, 0xf3, 0x30,
	0x14, 0x85, 0xeb, 0xbf, 0x69, 0xf4, 0xfb, 0x36, 0x61, 0xf0, 0x80, 0xac, 0x88, 0x21, 0x58, 0x42,
	0x0a, 0x0c, 0x16, 0x2a, 0x13, 0xac, 0xbc, 0x81, 0x3b, 0x30, 0x30, 0x20, 0xe3, 0x5e, 0x1a, 0x84,
	0x52, 0x9b, 0xd8, 0x79, 0x08, 0xde, 0x1a, 0xf5, 0xa6, 0x94, 0xb0, 0x20, 0xb1, 0x9d, 0x73, 0x72,
	0x74, 0xf3, 0x1d, 0x19, 0x60, 0xdb, 0x07, 0xa7, 0x43, 0xef, 0x93, 0x57, 0x1f, 0x0c, 0xf8, 0x7d,
	0x6b, 0x77, 0x5b, 0x5c, 0x63, 0x12, 0x02, 0xb2, 0x8d, 0x4d, 0x56, 0xb2, 0x9a, 0x35, 0x85, 0x21,
	0x2d, 0x2a, 0xf8, 0xef, 0x5a, 0x74, 0x6f, 0x71, 0xe8, 0xe4, 0xbf, 0x9a, 0x35, 0xdc, 0x1c, 0xbd,
	0x38, 0x85, 0xfc, 0xc5, 0xf7, 0x9d, 0x4d, 0x72, 0x4e, 0x5f, 0x0e, 0x6e, 0x9f, 0x47, 0x3f, 0xf4,
	0x0e, 0x65, 0x36, 0xe6, 0xa3, 0x13, 0x67, 0xc0, 0xd3, 0x6b, 0x87, 0x31, 0xd9, 0x2e, 0xc8, 0x45,
	0xcd, 0x9a, 0xb9, 0xf9, 0x0e, 0xd4, 0x39, 0x2c, 0x0d, 0xda, 0x8d, 0xc1, 0xf7, 0x01, 0x23, 0xc1,
	0x04, 0x9b, 0x5a, 0x82, 0xe1, 0x86, 0xb4, 0xba, 0x85, 0x62, 0xac, 0xc4, 0xe0, 0x77, 0x11, 0xc5,
	0x25, 0x80, 0x23, 0xfa, 0xa7, 0x88, 0x89, 0x9a, 0xcb, 0x15, 0xe8, 0xe3, 0x20, 0xc3, 0xdd, 0x97,
	0x54, 0x0a, 0x8a, 0x07, 0x9b, 0x5c, 0xfb, 0xdb, 0xf9, 0x3b, 0x28, 0x0f, 0x9d, 0x3f, 0xdf, 0x5f,
	0x3d, 0x42, 0xbe, 0x1e, 0x57, 0x5e, 0x40, 0xb6, 0x87, 0x14, 0x85, 0x9e, 0xcc, 0xa9, 0x4a, 0x3d,
	0x25, 0x57, 0x33, 0x71, 0x05, 0x0b, 0xfa, 0x99, 0x28, 0xf5, 0x14, 0xac, 0x3a, 0xd1, 0x3f, 0x18,
	0xd4, 0xec, 0x9a, 0x3d, 0xe7, 0xf4, 0x5a, 0x37, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xbb, 0xaf,
	0xaa, 0x9e, 0xbb, 0x01, 0x00, 0x00,
}
