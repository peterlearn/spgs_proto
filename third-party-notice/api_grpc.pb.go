// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: third-party-notice/api.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ThirdPartyNotice_SendMessage_FullMethodName = "/third_party_notice.ThirdPartyNotice/SendMessage"
)

// ThirdPartyNoticeClient is the client API for ThirdPartyNotice service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ThirdPartyNoticeClient interface {
	// 发送消息
	SendMessage(ctx context.Context, in *SendMessageReq, opts ...grpc.CallOption) (*Empty, error)
}

type thirdPartyNoticeClient struct {
	cc grpc.ClientConnInterface
}

func NewThirdPartyNoticeClient(cc grpc.ClientConnInterface) ThirdPartyNoticeClient {
	return &thirdPartyNoticeClient{cc}
}

func (c *thirdPartyNoticeClient) SendMessage(ctx context.Context, in *SendMessageReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, ThirdPartyNotice_SendMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ThirdPartyNoticeServer is the server API for ThirdPartyNotice service.
// All implementations must embed UnimplementedThirdPartyNoticeServer
// for forward compatibility
type ThirdPartyNoticeServer interface {
	// 发送消息
	SendMessage(context.Context, *SendMessageReq) (*Empty, error)
	mustEmbedUnimplementedThirdPartyNoticeServer()
}

// UnimplementedThirdPartyNoticeServer must be embedded to have forward compatible implementations.
type UnimplementedThirdPartyNoticeServer struct {
}

func (UnimplementedThirdPartyNoticeServer) SendMessage(context.Context, *SendMessageReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedThirdPartyNoticeServer) mustEmbedUnimplementedThirdPartyNoticeServer() {}

// UnsafeThirdPartyNoticeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ThirdPartyNoticeServer will
// result in compilation errors.
type UnsafeThirdPartyNoticeServer interface {
	mustEmbedUnimplementedThirdPartyNoticeServer()
}

func RegisterThirdPartyNoticeServer(s grpc.ServiceRegistrar, srv ThirdPartyNoticeServer) {
	s.RegisterService(&ThirdPartyNotice_ServiceDesc, srv)
}

func _ThirdPartyNotice_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThirdPartyNoticeServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ThirdPartyNotice_SendMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThirdPartyNoticeServer).SendMessage(ctx, req.(*SendMessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ThirdPartyNotice_ServiceDesc is the grpc.ServiceDesc for ThirdPartyNotice service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ThirdPartyNotice_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "third_party_notice.ThirdPartyNotice",
	HandlerType: (*ThirdPartyNoticeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessage",
			Handler:    _ThirdPartyNotice_SendMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "third-party-notice/api.proto",
}
