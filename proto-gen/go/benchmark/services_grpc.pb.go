// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: benchmark/services.proto

package benchmark

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
	BenchmarkService_Execute_FullMethodName = "/BenchmarkService/Execute"
)

// BenchmarkServiceClient is the client API for BenchmarkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BenchmarkServiceClient interface {
	Execute(ctx context.Context, in *ExecuteRequest, opts ...grpc.CallOption) (BenchmarkService_ExecuteClient, error)
}

type benchmarkServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBenchmarkServiceClient(cc grpc.ClientConnInterface) BenchmarkServiceClient {
	return &benchmarkServiceClient{cc}
}

func (c *benchmarkServiceClient) Execute(ctx context.Context, in *ExecuteRequest, opts ...grpc.CallOption) (BenchmarkService_ExecuteClient, error) {
	stream, err := c.cc.NewStream(ctx, &BenchmarkService_ServiceDesc.Streams[0], BenchmarkService_Execute_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &benchmarkServiceExecuteClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BenchmarkService_ExecuteClient interface {
	Recv() (*ExecuteResponse, error)
	grpc.ClientStream
}

type benchmarkServiceExecuteClient struct {
	grpc.ClientStream
}

func (x *benchmarkServiceExecuteClient) Recv() (*ExecuteResponse, error) {
	m := new(ExecuteResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BenchmarkServiceServer is the server API for BenchmarkService service.
// All implementations should embed UnimplementedBenchmarkServiceServer
// for forward compatibility
type BenchmarkServiceServer interface {
	Execute(*ExecuteRequest, BenchmarkService_ExecuteServer) error
}

// UnimplementedBenchmarkServiceServer should be embedded to have forward compatible implementations.
type UnimplementedBenchmarkServiceServer struct {
}

func (UnimplementedBenchmarkServiceServer) Execute(*ExecuteRequest, BenchmarkService_ExecuteServer) error {
	return status.Errorf(codes.Unimplemented, "method Execute not implemented")
}

// UnsafeBenchmarkServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BenchmarkServiceServer will
// result in compilation errors.
type UnsafeBenchmarkServiceServer interface {
	mustEmbedUnimplementedBenchmarkServiceServer()
}

func RegisterBenchmarkServiceServer(s grpc.ServiceRegistrar, srv BenchmarkServiceServer) {
	s.RegisterService(&BenchmarkService_ServiceDesc, srv)
}

func _BenchmarkService_Execute_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ExecuteRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BenchmarkServiceServer).Execute(m, &benchmarkServiceExecuteServer{stream})
}

type BenchmarkService_ExecuteServer interface {
	Send(*ExecuteResponse) error
	grpc.ServerStream
}

type benchmarkServiceExecuteServer struct {
	grpc.ServerStream
}

func (x *benchmarkServiceExecuteServer) Send(m *ExecuteResponse) error {
	return x.ServerStream.SendMsg(m)
}

// BenchmarkService_ServiceDesc is the grpc.ServiceDesc for BenchmarkService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BenchmarkService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "BenchmarkService",
	HandlerType: (*BenchmarkServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Execute",
			Handler:       _BenchmarkService_Execute_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "benchmark/services.proto",
}
