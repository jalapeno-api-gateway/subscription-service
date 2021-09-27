// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package subscriptionservice

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

// SubscriptionServiceClient is the client API for SubscriptionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SubscriptionServiceClient interface {
	SubscribeToLsNodes(ctx context.Context, in *TopologySubscription, opts ...grpc.CallOption) (SubscriptionService_SubscribeToLsNodesClient, error)
	SubscribeToLsLinks(ctx context.Context, in *TopologySubscription, opts ...grpc.CallOption) (SubscriptionService_SubscribeToLsLinksClient, error)
	SubscribeToTelemetryData(ctx context.Context, in *TelemetrySubscription, opts ...grpc.CallOption) (SubscriptionService_SubscribeToTelemetryDataClient, error)
}

type subscriptionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSubscriptionServiceClient(cc grpc.ClientConnInterface) SubscriptionServiceClient {
	return &subscriptionServiceClient{cc}
}

func (c *subscriptionServiceClient) SubscribeToLsNodes(ctx context.Context, in *TopologySubscription, opts ...grpc.CallOption) (SubscriptionService_SubscribeToLsNodesClient, error) {
	stream, err := c.cc.NewStream(ctx, &SubscriptionService_ServiceDesc.Streams[0], "/subscriptionservice.SubscriptionService/SubscribeToLsNodes", opts...)
	if err != nil {
		return nil, err
	}
	x := &subscriptionServiceSubscribeToLsNodesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SubscriptionService_SubscribeToLsNodesClient interface {
	Recv() (*LsNodeEvent, error)
	grpc.ClientStream
}

type subscriptionServiceSubscribeToLsNodesClient struct {
	grpc.ClientStream
}

func (x *subscriptionServiceSubscribeToLsNodesClient) Recv() (*LsNodeEvent, error) {
	m := new(LsNodeEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *subscriptionServiceClient) SubscribeToLsLinks(ctx context.Context, in *TopologySubscription, opts ...grpc.CallOption) (SubscriptionService_SubscribeToLsLinksClient, error) {
	stream, err := c.cc.NewStream(ctx, &SubscriptionService_ServiceDesc.Streams[1], "/subscriptionservice.SubscriptionService/SubscribeToLsLinks", opts...)
	if err != nil {
		return nil, err
	}
	x := &subscriptionServiceSubscribeToLsLinksClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SubscriptionService_SubscribeToLsLinksClient interface {
	Recv() (*LsLinkEvent, error)
	grpc.ClientStream
}

type subscriptionServiceSubscribeToLsLinksClient struct {
	grpc.ClientStream
}

func (x *subscriptionServiceSubscribeToLsLinksClient) Recv() (*LsLinkEvent, error) {
	m := new(LsLinkEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *subscriptionServiceClient) SubscribeToTelemetryData(ctx context.Context, in *TelemetrySubscription, opts ...grpc.CallOption) (SubscriptionService_SubscribeToTelemetryDataClient, error) {
	stream, err := c.cc.NewStream(ctx, &SubscriptionService_ServiceDesc.Streams[2], "/subscriptionservice.SubscriptionService/SubscribeToTelemetryData", opts...)
	if err != nil {
		return nil, err
	}
	x := &subscriptionServiceSubscribeToTelemetryDataClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SubscriptionService_SubscribeToTelemetryDataClient interface {
	Recv() (*TelemetryEvent, error)
	grpc.ClientStream
}

type subscriptionServiceSubscribeToTelemetryDataClient struct {
	grpc.ClientStream
}

func (x *subscriptionServiceSubscribeToTelemetryDataClient) Recv() (*TelemetryEvent, error) {
	m := new(TelemetryEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SubscriptionServiceServer is the server API for SubscriptionService service.
// All implementations must embed UnimplementedSubscriptionServiceServer
// for forward compatibility
type SubscriptionServiceServer interface {
	SubscribeToLsNodes(*TopologySubscription, SubscriptionService_SubscribeToLsNodesServer) error
	SubscribeToLsLinks(*TopologySubscription, SubscriptionService_SubscribeToLsLinksServer) error
	SubscribeToTelemetryData(*TelemetrySubscription, SubscriptionService_SubscribeToTelemetryDataServer) error
	mustEmbedUnimplementedSubscriptionServiceServer()
}

// UnimplementedSubscriptionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSubscriptionServiceServer struct {
}

func (UnimplementedSubscriptionServiceServer) SubscribeToLsNodes(*TopologySubscription, SubscriptionService_SubscribeToLsNodesServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeToLsNodes not implemented")
}
func (UnimplementedSubscriptionServiceServer) SubscribeToLsLinks(*TopologySubscription, SubscriptionService_SubscribeToLsLinksServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeToLsLinks not implemented")
}
func (UnimplementedSubscriptionServiceServer) SubscribeToTelemetryData(*TelemetrySubscription, SubscriptionService_SubscribeToTelemetryDataServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeToTelemetryData not implemented")
}
func (UnimplementedSubscriptionServiceServer) mustEmbedUnimplementedSubscriptionServiceServer() {}

// UnsafeSubscriptionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SubscriptionServiceServer will
// result in compilation errors.
type UnsafeSubscriptionServiceServer interface {
	mustEmbedUnimplementedSubscriptionServiceServer()
}

func RegisterSubscriptionServiceServer(s grpc.ServiceRegistrar, srv SubscriptionServiceServer) {
	s.RegisterService(&SubscriptionService_ServiceDesc, srv)
}

func _SubscriptionService_SubscribeToLsNodes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TopologySubscription)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SubscriptionServiceServer).SubscribeToLsNodes(m, &subscriptionServiceSubscribeToLsNodesServer{stream})
}

type SubscriptionService_SubscribeToLsNodesServer interface {
	Send(*LsNodeEvent) error
	grpc.ServerStream
}

type subscriptionServiceSubscribeToLsNodesServer struct {
	grpc.ServerStream
}

func (x *subscriptionServiceSubscribeToLsNodesServer) Send(m *LsNodeEvent) error {
	return x.ServerStream.SendMsg(m)
}

func _SubscriptionService_SubscribeToLsLinks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TopologySubscription)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SubscriptionServiceServer).SubscribeToLsLinks(m, &subscriptionServiceSubscribeToLsLinksServer{stream})
}

type SubscriptionService_SubscribeToLsLinksServer interface {
	Send(*LsLinkEvent) error
	grpc.ServerStream
}

type subscriptionServiceSubscribeToLsLinksServer struct {
	grpc.ServerStream
}

func (x *subscriptionServiceSubscribeToLsLinksServer) Send(m *LsLinkEvent) error {
	return x.ServerStream.SendMsg(m)
}

func _SubscriptionService_SubscribeToTelemetryData_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TelemetrySubscription)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SubscriptionServiceServer).SubscribeToTelemetryData(m, &subscriptionServiceSubscribeToTelemetryDataServer{stream})
}

type SubscriptionService_SubscribeToTelemetryDataServer interface {
	Send(*TelemetryEvent) error
	grpc.ServerStream
}

type subscriptionServiceSubscribeToTelemetryDataServer struct {
	grpc.ServerStream
}

func (x *subscriptionServiceSubscribeToTelemetryDataServer) Send(m *TelemetryEvent) error {
	return x.ServerStream.SendMsg(m)
}

// SubscriptionService_ServiceDesc is the grpc.ServiceDesc for SubscriptionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SubscriptionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "subscriptionservice.SubscriptionService",
	HandlerType: (*SubscriptionServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeToLsNodes",
			Handler:       _SubscriptionService_SubscribeToLsNodes_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SubscribeToLsLinks",
			Handler:       _SubscriptionService_SubscribeToLsLinks_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SubscribeToTelemetryData",
			Handler:       _SubscriptionService_SubscribeToTelemetryData_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "subscriptionservice.proto",
}