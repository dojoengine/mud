// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: proto/ecs-stream.proto

package ecs_stream

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

// ECSStreamServiceClient is the client API for ECSStreamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ECSStreamServiceClient interface {
	// Opens a cursor to receive the latest ECS events and additional data specified via request.
	SubscribeToStreamLatest(ctx context.Context, in *ECSStreamBlockBundleRequest, opts ...grpc.CallOption) (ECSStreamService_SubscribeToStreamLatestClient, error)
}

type eCSStreamServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewECSStreamServiceClient(cc grpc.ClientConnInterface) ECSStreamServiceClient {
	return &eCSStreamServiceClient{cc}
}

func (c *eCSStreamServiceClient) SubscribeToStreamLatest(ctx context.Context, in *ECSStreamBlockBundleRequest, opts ...grpc.CallOption) (ECSStreamService_SubscribeToStreamLatestClient, error) {
	stream, err := c.cc.NewStream(ctx, &ECSStreamService_ServiceDesc.Streams[0], "/ecsstream.ECSStreamService/SubscribeToStreamLatest", opts...)
	if err != nil {
		return nil, err
	}
	x := &eCSStreamServiceSubscribeToStreamLatestClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ECSStreamService_SubscribeToStreamLatestClient interface {
	Recv() (*ECSStreamBlockBundleReply, error)
	grpc.ClientStream
}

type eCSStreamServiceSubscribeToStreamLatestClient struct {
	grpc.ClientStream
}

func (x *eCSStreamServiceSubscribeToStreamLatestClient) Recv() (*ECSStreamBlockBundleReply, error) {
	m := new(ECSStreamBlockBundleReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ECSStreamServiceServer is the server API for ECSStreamService service.
// All implementations must embed UnimplementedECSStreamServiceServer
// for forward compatibility
type ECSStreamServiceServer interface {
	// Opens a cursor to receive the latest ECS events and additional data specified via request.
	SubscribeToStreamLatest(*ECSStreamBlockBundleRequest, ECSStreamService_SubscribeToStreamLatestServer) error
	mustEmbedUnimplementedECSStreamServiceServer()
}

// UnimplementedECSStreamServiceServer must be embedded to have forward compatible implementations.
type UnimplementedECSStreamServiceServer struct {
}

func (UnimplementedECSStreamServiceServer) SubscribeToStreamLatest(*ECSStreamBlockBundleRequest, ECSStreamService_SubscribeToStreamLatestServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeToStreamLatest not implemented")
}
func (UnimplementedECSStreamServiceServer) mustEmbedUnimplementedECSStreamServiceServer() {}

// UnsafeECSStreamServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ECSStreamServiceServer will
// result in compilation errors.
type UnsafeECSStreamServiceServer interface {
	mustEmbedUnimplementedECSStreamServiceServer()
}

func RegisterECSStreamServiceServer(s grpc.ServiceRegistrar, srv ECSStreamServiceServer) {
	s.RegisterService(&ECSStreamService_ServiceDesc, srv)
}

func _ECSStreamService_SubscribeToStreamLatest_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ECSStreamBlockBundleRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ECSStreamServiceServer).SubscribeToStreamLatest(m, &eCSStreamServiceSubscribeToStreamLatestServer{stream})
}

type ECSStreamService_SubscribeToStreamLatestServer interface {
	Send(*ECSStreamBlockBundleReply) error
	grpc.ServerStream
}

type eCSStreamServiceSubscribeToStreamLatestServer struct {
	grpc.ServerStream
}

func (x *eCSStreamServiceSubscribeToStreamLatestServer) Send(m *ECSStreamBlockBundleReply) error {
	return x.ServerStream.SendMsg(m)
}

// ECSStreamService_ServiceDesc is the grpc.ServiceDesc for ECSStreamService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ECSStreamService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ecsstream.ECSStreamService",
	HandlerType: (*ECSStreamServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeToStreamLatest",
			Handler:       _ECSStreamService_SubscribeToStreamLatest_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/ecs-stream.proto",
}
