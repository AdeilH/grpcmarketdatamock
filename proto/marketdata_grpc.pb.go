// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: marketdata.proto

package proto

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

// MarketDataServiceClient is the client API for MarketDataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MarketDataServiceClient interface {
	// Unary RPC method for getting a single snapshot of market data.
	GetMarketData(ctx context.Context, in *MarketDataRequest, opts ...grpc.CallOption) (*MarketDataResponse, error)
	// Server-streaming RPC method for pushing market data updates.
	StreamMarketData(ctx context.Context, in *MarketDataRequest, opts ...grpc.CallOption) (MarketDataService_StreamMarketDataClient, error)
}

type marketDataServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMarketDataServiceClient(cc grpc.ClientConnInterface) MarketDataServiceClient {
	return &marketDataServiceClient{cc}
}

func (c *marketDataServiceClient) GetMarketData(ctx context.Context, in *MarketDataRequest, opts ...grpc.CallOption) (*MarketDataResponse, error) {
	out := new(MarketDataResponse)
	err := c.cc.Invoke(ctx, "/marketdata.MarketDataService/GetMarketData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketDataServiceClient) StreamMarketData(ctx context.Context, in *MarketDataRequest, opts ...grpc.CallOption) (MarketDataService_StreamMarketDataClient, error) {
	stream, err := c.cc.NewStream(ctx, &MarketDataService_ServiceDesc.Streams[0], "/marketdata.MarketDataService/StreamMarketData", opts...)
	if err != nil {
		return nil, err
	}
	x := &marketDataServiceStreamMarketDataClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MarketDataService_StreamMarketDataClient interface {
	Recv() (*MarketDataUpdate, error)
	grpc.ClientStream
}

type marketDataServiceStreamMarketDataClient struct {
	grpc.ClientStream
}

func (x *marketDataServiceStreamMarketDataClient) Recv() (*MarketDataUpdate, error) {
	m := new(MarketDataUpdate)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MarketDataServiceServer is the server API for MarketDataService service.
// All implementations must embed UnimplementedMarketDataServiceServer
// for forward compatibility
type MarketDataServiceServer interface {
	// Unary RPC method for getting a single snapshot of market data.
	GetMarketData(context.Context, *MarketDataRequest) (*MarketDataResponse, error)
	// Server-streaming RPC method for pushing market data updates.
	StreamMarketData(*MarketDataRequest, MarketDataService_StreamMarketDataServer) error
	mustEmbedUnimplementedMarketDataServiceServer()
}

// UnimplementedMarketDataServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMarketDataServiceServer struct {
}

func (UnimplementedMarketDataServiceServer) GetMarketData(context.Context, *MarketDataRequest) (*MarketDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMarketData not implemented")
}
func (UnimplementedMarketDataServiceServer) StreamMarketData(*MarketDataRequest, MarketDataService_StreamMarketDataServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamMarketData not implemented")
}
func (UnimplementedMarketDataServiceServer) mustEmbedUnimplementedMarketDataServiceServer() {}

// UnsafeMarketDataServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MarketDataServiceServer will
// result in compilation errors.
type UnsafeMarketDataServiceServer interface {
	mustEmbedUnimplementedMarketDataServiceServer()
}

func RegisterMarketDataServiceServer(s grpc.ServiceRegistrar, srv MarketDataServiceServer) {
	s.RegisterService(&MarketDataService_ServiceDesc, srv)
}

func _MarketDataService_GetMarketData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarketDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketDataServiceServer).GetMarketData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/marketdata.MarketDataService/GetMarketData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketDataServiceServer).GetMarketData(ctx, req.(*MarketDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MarketDataService_StreamMarketData_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MarketDataRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MarketDataServiceServer).StreamMarketData(m, &marketDataServiceStreamMarketDataServer{stream})
}

type MarketDataService_StreamMarketDataServer interface {
	Send(*MarketDataUpdate) error
	grpc.ServerStream
}

type marketDataServiceStreamMarketDataServer struct {
	grpc.ServerStream
}

func (x *marketDataServiceStreamMarketDataServer) Send(m *MarketDataUpdate) error {
	return x.ServerStream.SendMsg(m)
}

// MarketDataService_ServiceDesc is the grpc.ServiceDesc for MarketDataService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MarketDataService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "marketdata.MarketDataService",
	HandlerType: (*MarketDataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMarketData",
			Handler:    _MarketDataService_GetMarketData_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamMarketData",
			Handler:       _MarketDataService_StreamMarketData_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "marketdata.proto",
}