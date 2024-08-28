// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.26.1
// source: proto/stats.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Hub_GetDataStream_FullMethodName = "/Hub/GetDataStream"
	Hub_Register_FullMethodName      = "/Hub/Register"
	Hub_PushData_FullMethodName      = "/Hub/PushData"
)

// HubClient is the client API for Hub service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Hub is the stats gathering service (hub)
type HubClient interface {
	// GetDataStream returns a stream of fresh data from the stats hub
	GetDataStream(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (grpc.ServerStreamingClient[DataPoint], error)
	// Register registers the node instance with the stats hub
	Register(ctx context.Context, in *StaticInfo, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// PushData continuously pushes the node data to the stats hub
	PushData(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[DynamicInfo, emptypb.Empty], error)
}

type hubClient struct {
	cc grpc.ClientConnInterface
}

func NewHubClient(cc grpc.ClientConnInterface) HubClient {
	return &hubClient{cc}
}

func (c *hubClient) GetDataStream(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (grpc.ServerStreamingClient[DataPoint], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Hub_ServiceDesc.Streams[0], Hub_GetDataStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[emptypb.Empty, DataPoint]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Hub_GetDataStreamClient = grpc.ServerStreamingClient[DataPoint]

func (c *hubClient) Register(ctx context.Context, in *StaticInfo, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Hub_Register_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubClient) PushData(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[DynamicInfo, emptypb.Empty], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Hub_ServiceDesc.Streams[1], Hub_PushData_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[DynamicInfo, emptypb.Empty]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Hub_PushDataClient = grpc.ClientStreamingClient[DynamicInfo, emptypb.Empty]

// HubServer is the server API for Hub service.
// All implementations must embed UnimplementedHubServer
// for forward compatibility.
//
// Hub is the stats gathering service (hub)
type HubServer interface {
	// GetDataStream returns a stream of fresh data from the stats hub
	GetDataStream(*emptypb.Empty, grpc.ServerStreamingServer[DataPoint]) error
	// Register registers the node instance with the stats hub
	Register(context.Context, *StaticInfo) (*emptypb.Empty, error)
	// PushData continuously pushes the node data to the stats hub
	PushData(grpc.ClientStreamingServer[DynamicInfo, emptypb.Empty]) error
	mustEmbedUnimplementedHubServer()
}

// UnimplementedHubServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedHubServer struct{}

func (UnimplementedHubServer) GetDataStream(*emptypb.Empty, grpc.ServerStreamingServer[DataPoint]) error {
	return status.Errorf(codes.Unimplemented, "method GetDataStream not implemented")
}
func (UnimplementedHubServer) Register(context.Context, *StaticInfo) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedHubServer) PushData(grpc.ClientStreamingServer[DynamicInfo, emptypb.Empty]) error {
	return status.Errorf(codes.Unimplemented, "method PushData not implemented")
}
func (UnimplementedHubServer) mustEmbedUnimplementedHubServer() {}
func (UnimplementedHubServer) testEmbeddedByValue()             {}

// UnsafeHubServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HubServer will
// result in compilation errors.
type UnsafeHubServer interface {
	mustEmbedUnimplementedHubServer()
}

func RegisterHubServer(s grpc.ServiceRegistrar, srv HubServer) {
	// If the following call pancis, it indicates UnimplementedHubServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Hub_ServiceDesc, srv)
}

func _Hub_GetDataStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HubServer).GetDataStream(m, &grpc.GenericServerStream[emptypb.Empty, DataPoint]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Hub_GetDataStreamServer = grpc.ServerStreamingServer[DataPoint]

func _Hub_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StaticInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Hub_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubServer).Register(ctx, req.(*StaticInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hub_PushData_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HubServer).PushData(&grpc.GenericServerStream[DynamicInfo, emptypb.Empty]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Hub_PushDataServer = grpc.ClientStreamingServer[DynamicInfo, emptypb.Empty]

// Hub_ServiceDesc is the grpc.ServiceDesc for Hub service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Hub_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Hub",
	HandlerType: (*HubServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Hub_Register_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetDataStream",
			Handler:       _Hub_GetDataStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "PushData",
			Handler:       _Hub_PushData_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "proto/stats.proto",
}
