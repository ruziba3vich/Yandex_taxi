// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.3
// source: creations.proto

package genprotos

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	TaxiCreationService_CreateTaxi_FullMethodName = "/creations.TaxiCreationService/CreateTaxi"
)

// TaxiCreationServiceClient is the client API for TaxiCreationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaxiCreationServiceClient interface {
	CreateTaxi(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CreateTaxiResponse, error)
}

type taxiCreationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTaxiCreationServiceClient(cc grpc.ClientConnInterface) TaxiCreationServiceClient {
	return &taxiCreationServiceClient{cc}
}

func (c *taxiCreationServiceClient) CreateTaxi(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CreateTaxiResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateTaxiResponse)
	err := c.cc.Invoke(ctx, TaxiCreationService_CreateTaxi_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaxiCreationServiceServer is the server API for TaxiCreationService service.
// All implementations must embed UnimplementedTaxiCreationServiceServer
// for forward compatibility
type TaxiCreationServiceServer interface {
	CreateTaxi(context.Context, *emptypb.Empty) (*CreateTaxiResponse, error)
	mustEmbedUnimplementedTaxiCreationServiceServer()
}

// UnimplementedTaxiCreationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTaxiCreationServiceServer struct {
}

func (UnimplementedTaxiCreationServiceServer) CreateTaxi(context.Context, *emptypb.Empty) (*CreateTaxiResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTaxi not implemented")
}
func (UnimplementedTaxiCreationServiceServer) mustEmbedUnimplementedTaxiCreationServiceServer() {}

// UnsafeTaxiCreationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaxiCreationServiceServer will
// result in compilation errors.
type UnsafeTaxiCreationServiceServer interface {
	mustEmbedUnimplementedTaxiCreationServiceServer()
}

func RegisterTaxiCreationServiceServer(s grpc.ServiceRegistrar, srv TaxiCreationServiceServer) {
	s.RegisterService(&TaxiCreationService_ServiceDesc, srv)
}

func _TaxiCreationService_CreateTaxi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaxiCreationServiceServer).CreateTaxi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaxiCreationService_CreateTaxi_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaxiCreationServiceServer).CreateTaxi(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// TaxiCreationService_ServiceDesc is the grpc.ServiceDesc for TaxiCreationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TaxiCreationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "creations.TaxiCreationService",
	HandlerType: (*TaxiCreationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTaxi",
			Handler:    _TaxiCreationService_CreateTaxi_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "creations.proto",
}
