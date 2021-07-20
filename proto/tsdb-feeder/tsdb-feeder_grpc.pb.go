// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package tsdbfeeder

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

// TsdbFeederClient is the client API for TsdbFeeder service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TsdbFeederClient interface {
	SubscribeToDataRates(ctx context.Context, in *IPv4Addresses, opts ...grpc.CallOption) (TsdbFeeder_SubscribeToDataRatesClient, error)
}

type tsdbFeederClient struct {
	cc grpc.ClientConnInterface
}

func NewTsdbFeederClient(cc grpc.ClientConnInterface) TsdbFeederClient {
	return &tsdbFeederClient{cc}
}

func (c *tsdbFeederClient) SubscribeToDataRates(ctx context.Context, in *IPv4Addresses, opts ...grpc.CallOption) (TsdbFeeder_SubscribeToDataRatesClient, error) {
	stream, err := c.cc.NewStream(ctx, &TsdbFeeder_ServiceDesc.Streams[0], "/tsdbfeeder.TsdbFeeder/SubscribeToDataRates", opts...)
	if err != nil {
		return nil, err
	}
	x := &tsdbFeederSubscribeToDataRatesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TsdbFeeder_SubscribeToDataRatesClient interface {
	Recv() (*DataRate, error)
	grpc.ClientStream
}

type tsdbFeederSubscribeToDataRatesClient struct {
	grpc.ClientStream
}

func (x *tsdbFeederSubscribeToDataRatesClient) Recv() (*DataRate, error) {
	m := new(DataRate)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TsdbFeederServer is the server API for TsdbFeeder service.
// All implementations must embed UnimplementedTsdbFeederServer
// for forward compatibility
type TsdbFeederServer interface {
	SubscribeToDataRates(*IPv4Addresses, TsdbFeeder_SubscribeToDataRatesServer) error
	mustEmbedUnimplementedTsdbFeederServer()
}

// UnimplementedTsdbFeederServer must be embedded to have forward compatible implementations.
type UnimplementedTsdbFeederServer struct {
}

func (UnimplementedTsdbFeederServer) SubscribeToDataRates(*IPv4Addresses, TsdbFeeder_SubscribeToDataRatesServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeToDataRates not implemented")
}
func (UnimplementedTsdbFeederServer) mustEmbedUnimplementedTsdbFeederServer() {}

// UnsafeTsdbFeederServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TsdbFeederServer will
// result in compilation errors.
type UnsafeTsdbFeederServer interface {
	mustEmbedUnimplementedTsdbFeederServer()
}

func RegisterTsdbFeederServer(s grpc.ServiceRegistrar, srv TsdbFeederServer) {
	s.RegisterService(&TsdbFeeder_ServiceDesc, srv)
}

func _TsdbFeeder_SubscribeToDataRates_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(IPv4Addresses)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TsdbFeederServer).SubscribeToDataRates(m, &tsdbFeederSubscribeToDataRatesServer{stream})
}

type TsdbFeeder_SubscribeToDataRatesServer interface {
	Send(*DataRate) error
	grpc.ServerStream
}

type tsdbFeederSubscribeToDataRatesServer struct {
	grpc.ServerStream
}

func (x *tsdbFeederSubscribeToDataRatesServer) Send(m *DataRate) error {
	return x.ServerStream.SendMsg(m)
}

// TsdbFeeder_ServiceDesc is the grpc.ServiceDesc for TsdbFeeder service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TsdbFeeder_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tsdbfeeder.TsdbFeeder",
	HandlerType: (*TsdbFeederServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeToDataRates",
			Handler:       _TsdbFeeder_SubscribeToDataRates_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "tsdb-feeder.proto",
}
