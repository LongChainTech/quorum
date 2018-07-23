// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc.proto

package protofiles

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ClientClient is the client API for Client service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClientClient interface {
	Version(ctx context.Context, in *ApiVersion, opts ...grpc.CallOption) (*ApiVersion, error)
	Upcheck(ctx context.Context, in *UpCheckResponse, opts ...grpc.CallOption) (*UpCheckResponse, error)
	Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error)
	Receive(ctx context.Context, in *ReceiveRequest, opts ...grpc.CallOption) (*ReceiveResponse, error)
	UpdatePartyInfo(ctx context.Context, in *PartyInfo, opts ...grpc.CallOption) (*PartyInfoResponse, error)
	Push(ctx context.Context, in *PushPayload, opts ...grpc.CallOption) (*PartyInfoResponse, error)
}

type clientClient struct {
	cc *grpc.ClientConn
}

func NewClientClient(cc *grpc.ClientConn) ClientClient {
	return &clientClient{cc}
}

func (c *clientClient) Version(ctx context.Context, in *ApiVersion, opts ...grpc.CallOption) (*ApiVersion, error) {
	out := new(ApiVersion)
	err := c.cc.Invoke(ctx, "/protofiles.Client/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientClient) Upcheck(ctx context.Context, in *UpCheckResponse, opts ...grpc.CallOption) (*UpCheckResponse, error) {
	out := new(UpCheckResponse)
	err := c.cc.Invoke(ctx, "/protofiles.Client/Upcheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientClient) Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error) {
	out := new(SendResponse)
	err := c.cc.Invoke(ctx, "/protofiles.Client/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientClient) Receive(ctx context.Context, in *ReceiveRequest, opts ...grpc.CallOption) (*ReceiveResponse, error) {
	out := new(ReceiveResponse)
	err := c.cc.Invoke(ctx, "/protofiles.Client/Receive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientClient) UpdatePartyInfo(ctx context.Context, in *PartyInfo, opts ...grpc.CallOption) (*PartyInfoResponse, error) {
	out := new(PartyInfoResponse)
	err := c.cc.Invoke(ctx, "/protofiles.Client/UpdatePartyInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientClient) Push(ctx context.Context, in *PushPayload, opts ...grpc.CallOption) (*PartyInfoResponse, error) {
	out := new(PartyInfoResponse)
	err := c.cc.Invoke(ctx, "/protofiles.Client/Push", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientServer is the server API for Client service.
type ClientServer interface {
	Version(context.Context, *ApiVersion) (*ApiVersion, error)
	Upcheck(context.Context, *UpCheckResponse) (*UpCheckResponse, error)
	Send(context.Context, *SendRequest) (*SendResponse, error)
	Receive(context.Context, *ReceiveRequest) (*ReceiveResponse, error)
	UpdatePartyInfo(context.Context, *PartyInfo) (*PartyInfoResponse, error)
	Push(context.Context, *PushPayload) (*PartyInfoResponse, error)
}

func RegisterClientServer(s *grpc.Server, srv ClientServer) {
	s.RegisterService(&_Client_serviceDesc, srv)
}

func _Client_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApiVersion)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protofiles.Client/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServer).Version(ctx, req.(*ApiVersion))
	}
	return interceptor(ctx, in, info, handler)
}

func _Client_Upcheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpCheckResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServer).Upcheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protofiles.Client/Upcheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServer).Upcheck(ctx, req.(*UpCheckResponse))
	}
	return interceptor(ctx, in, info, handler)
}

func _Client_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protofiles.Client/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServer).Send(ctx, req.(*SendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Client_Receive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReceiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServer).Receive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protofiles.Client/Receive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServer).Receive(ctx, req.(*ReceiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Client_UpdatePartyInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PartyInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServer).UpdatePartyInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protofiles.Client/UpdatePartyInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServer).UpdatePartyInfo(ctx, req.(*PartyInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Client_Push_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushPayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServer).Push(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protofiles.Client/Push",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServer).Push(ctx, req.(*PushPayload))
	}
	return interceptor(ctx, in, info, handler)
}

var _Client_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protofiles.Client",
	HandlerType: (*ClientServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _Client_Version_Handler,
		},
		{
			MethodName: "Upcheck",
			Handler:    _Client_Upcheck_Handler,
		},
		{
			MethodName: "Send",
			Handler:    _Client_Send_Handler,
		},
		{
			MethodName: "Receive",
			Handler:    _Client_Receive_Handler,
		},
		{
			MethodName: "UpdatePartyInfo",
			Handler:    _Client_UpdatePartyInfo_Handler,
		},
		{
			MethodName: "Push",
			Handler:    _Client_Push_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc.proto",
}

func init() { proto.RegisterFile("grpc.proto", fileDescriptor_grpc_038875b6e19575ae) }

var fileDescriptor_grpc_038875b6e19575ae = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x8f, 0x41, 0x4b, 0xc3, 0x30,
	0x18, 0x86, 0x99, 0xce, 0x55, 0x3e, 0xd0, 0x49, 0x64, 0x2a, 0x99, 0x5e, 0x7a, 0xdc, 0xa1, 0x05,
	0x3d, 0xb9, 0x9b, 0xee, 0xe4, 0x45, 0x4a, 0xa5, 0xe2, 0x49, 0x88, 0xed, 0xd7, 0x36, 0x58, 0x93,
	0xd8, 0xa4, 0x83, 0x5d, 0xfd, 0x0b, 0xfe, 0x34, 0x7f, 0x82, 0xfe, 0x10, 0x49, 0x5b, 0x67, 0xc5,
	0xe9, 0x4e, 0x21, 0xcf, 0xfb, 0x7e, 0x0f, 0xbc, 0x00, 0x59, 0xa9, 0x62, 0x4f, 0x95, 0xd2, 0x48,
	0x02, 0xf5, 0x93, 0xf2, 0x02, 0x35, 0x3d, 0xce, 0xa4, 0xcc, 0x0a, 0xf4, 0x99, 0xe2, 0x3e, 0x13,
	0x42, 0x1a, 0x66, 0xb8, 0x14, 0xba, 0x69, 0xd2, 0xdd, 0x27, 0xd4, 0x9a, 0x65, 0xd8, 0xfe, 0x4f,
	0xdf, 0x37, 0x61, 0x30, 0x2b, 0x38, 0x0a, 0x43, 0xae, 0xc1, 0xb9, 0xc5, 0x52, 0x73, 0x29, 0xc8,
	0x81, 0xf7, 0x2d, 0xf4, 0x2e, 0x14, 0x6f, 0x39, 0xfd, 0x83, 0xbb, 0xfb, 0x2f, 0x6f, 0x1f, 0xaf,
	0x1b, 0x3b, 0xee, 0xb6, 0x3f, 0x6f, 0xc8, 0xb4, 0x37, 0x21, 0x77, 0xe0, 0x44, 0x2a, 0xce, 0x31,
	0x7e, 0x24, 0xe3, 0xee, 0x5d, 0xa4, 0x66, 0x16, 0x86, 0xa8, 0x95, 0x14, 0x1a, 0xe9, 0x7f, 0x61,
	0xc7, 0x5c, 0x35, 0x2e, 0x6b, 0x3e, 0x87, 0xfe, 0x0d, 0x8a, 0x84, 0x1c, 0x76, 0x2f, 0x2d, 0x09,
	0xf1, 0xb9, 0x42, 0x6d, 0xe8, 0xd1, 0xef, 0xa0, 0xf1, 0x91, 0x4b, 0x70, 0x42, 0x8c, 0x91, 0xcf,
	0x91, 0xd0, 0x6e, 0xa9, 0x85, 0x5f, 0x82, 0xf1, 0xca, 0xac, 0x75, 0xdc, 0xc3, 0x30, 0x52, 0x09,
	0x33, 0x18, 0xb0, 0xd2, 0x2c, 0xae, 0x44, 0x2a, 0xc9, 0xa8, 0xdb, 0x5f, 0x62, 0x7a, 0xb2, 0x12,
	0x2f, 0xc7, 0x8d, 0xea, 0x71, 0x43, 0x17, 0x7c, 0x65, 0x33, 0x2e, 0x52, 0x69, 0xe7, 0x05, 0xd0,
	0x0f, 0x2a, 0x9d, 0xff, 0x9c, 0x67, 0x49, 0xc0, 0x16, 0x85, 0x64, 0xc9, 0x3a, 0xed, 0x5e, 0xad,
	0x05, 0x77, 0xcb, 0x57, 0x95, 0xce, 0xa7, 0xbd, 0xc9, 0xc3, 0xa0, 0xee, 0x9f, 0x7d, 0x06, 0x00,
	0x00, 0xff, 0xff, 0x21, 0x0d, 0xce, 0xd4, 0x34, 0x02, 0x00, 0x00,
}
