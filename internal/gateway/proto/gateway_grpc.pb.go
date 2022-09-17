// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: internal/gateway/proto/gateway.proto

package gateway

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

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	Transport(ctx context.Context, opts ...grpc.CallOption) (Gateway_TransportClient, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) Transport(ctx context.Context, opts ...grpc.CallOption) (Gateway_TransportClient, error) {
	stream, err := c.cc.NewStream(ctx, &Gateway_ServiceDesc.Streams[0], "/gateway.Gateway/Transport", opts...)
	if err != nil {
		return nil, err
	}
	x := &gatewayTransportClient{stream}
	return x, nil
}

type Gateway_TransportClient interface {
	Send(*ClientMsg) error
	Recv() (*ServerMsg, error)
	grpc.ClientStream
}

type gatewayTransportClient struct {
	grpc.ClientStream
}

func (x *gatewayTransportClient) Send(m *ClientMsg) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gatewayTransportClient) Recv() (*ServerMsg, error) {
	m := new(ServerMsg)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	Transport(Gateway_TransportServer) error
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) Transport(Gateway_TransportServer) error {
	return status.Errorf(codes.Unimplemented, "method Transport not implemented")
}
func (UnimplementedGatewayServer) mustEmbedUnimplementedGatewayServer() {}

// UnsafeGatewayServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GatewayServer will
// result in compilation errors.
type UnsafeGatewayServer interface {
	mustEmbedUnimplementedGatewayServer()
}

func RegisterGatewayServer(s grpc.ServiceRegistrar, srv GatewayServer) {
	s.RegisterService(&Gateway_ServiceDesc, srv)
}

func _Gateway_Transport_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GatewayServer).Transport(&gatewayTransportServer{stream})
}

type Gateway_TransportServer interface {
	Send(*ServerMsg) error
	Recv() (*ClientMsg, error)
	grpc.ServerStream
}

type gatewayTransportServer struct {
	grpc.ServerStream
}

func (x *gatewayTransportServer) Send(m *ServerMsg) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gatewayTransportServer) Recv() (*ClientMsg, error) {
	m := new(ClientMsg)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gateway.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Transport",
			Handler:       _Gateway_Transport_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "internal/gateway/proto/gateway.proto",
}
