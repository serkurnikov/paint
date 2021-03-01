// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	"google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CalculationClient is the client API for Calculation service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalculationClient interface {
	GetCalculationParams(ctx context.Context, in *CalculationParams, opts ...grpc.CallOption) (*CalculationResult, error)
}

type calculationClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculationClient(cc grpc.ClientConnInterface) CalculationClient {
	return &calculationClient{cc}
}

func (c *calculationClient) GetCalculationParams(ctx context.Context, in *CalculationParams, opts ...grpc.CallOption) (*CalculationResult, error) {
	out := new(CalculationResult)
	err := c.cc.Invoke(ctx, "/promo.Calculation/GetCalculationParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculationServer is the server API for Calculation service.
// All implementations must embed UnimplementedCalculationServer
// for forward compatibility
type CalculationServer interface {
	GetCalculationParams(context.Context, *CalculationParams) (*CalculationResult, error)
	mustEmbedUnimplementedCalculationServer()
}

// UnimplementedCalculationServer must be embedded to have forward compatible implementations.
type UnimplementedCalculationServer struct {
}

func (UnimplementedCalculationServer) GetCalculationParams(context.Context, *CalculationParams) (*CalculationResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCalculationParams not implemented")
}
func (UnimplementedCalculationServer) mustEmbedUnimplementedCalculationServer() {}

// UnsafeCalculationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalculationServer will
// result in compilation errors.
type UnsafeCalculationServer interface {
	mustEmbedUnimplementedCalculationServer()
}

func RegisterCalculationServer(s grpc.ServiceRegistrar, srv CalculationServer) {
	s.RegisterService(&Calculation_ServiceDesc, srv)
}

func _Calculation_GetCalculationParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculationParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculationServer).GetCalculationParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/promo.Calculation/GetCalculationParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculationServer).GetCalculationParams(ctx, req.(*CalculationParams))
	}
	return interceptor(ctx, in, info, handler)
}

// Calculation_ServiceDesc is the grpc.ServiceDesc for Calculation service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Calculation_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "promo.Calculation",
	HandlerType: (*CalculationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCalculationParams",
			Handler:    _Calculation_GetCalculationParams_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calculation.proto",
}