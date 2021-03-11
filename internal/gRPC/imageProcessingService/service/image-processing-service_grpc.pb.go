// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package service

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

// ImageProcessingServiceClient is the client API for ImageProcessingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImageProcessingServiceClient interface {
	PyrMeanShiftFiltering(ctx context.Context, in *PyrRequest, opts ...grpc.CallOption) (*DefaultReply, error)
	DrawDefaultContours(ctx context.Context, in *ContoursRequest, opts ...grpc.CallOption) (*DefaultReply, error)
	DrawCustomContours(ctx context.Context, in *ContoursRequest, opts ...grpc.CallOption) (*DefaultReply, error)
	DrawHoughLinesWithParams(ctx context.Context, in *HoughLinesWithPRequest, opts ...grpc.CallOption) (*DefaultReply, error)
	DrawHoughCircles(ctx context.Context, in *HoughCirclesRequest, opts ...grpc.CallOption) (*DefaultReply, error)
	Threshold(ctx context.Context, in *ThresholdRequest, opts ...grpc.CallOption) (*DefaultReply, error)
	Watershed(ctx context.Context, in *WatershedRequest, opts ...grpc.CallOption) (*DefaultReply, error)
	Open(ctx context.Context, in *OpenRequest, opts ...grpc.CallOption) (*DefaultReply, error)
	Close(ctx context.Context, in *CloseRequest, opts ...grpc.CallOption) (*DefaultReply, error)
	FindBlendStructureAmongFabricColorsLUV(ctx context.Context, in *BlendStructureRequest, opts ...grpc.CallOption) (*BlendStructureReply, error)
}

type imageProcessingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewImageProcessingServiceClient(cc grpc.ClientConnInterface) ImageProcessingServiceClient {
	return &imageProcessingServiceClient{cc}
}

func (c *imageProcessingServiceClient) PyrMeanShiftFiltering(ctx context.Context, in *PyrRequest, opts ...grpc.CallOption) (*DefaultReply, error) {
	out := new(DefaultReply)
	err := c.cc.Invoke(ctx, "/service.ImageProcessingService/PyrMeanShiftFiltering", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageProcessingServiceClient) DrawDefaultContours(ctx context.Context, in *ContoursRequest, opts ...grpc.CallOption) (*DefaultReply, error) {
	out := new(DefaultReply)
	err := c.cc.Invoke(ctx, "/service.ImageProcessingService/DrawDefaultContours", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageProcessingServiceClient) DrawCustomContours(ctx context.Context, in *ContoursRequest, opts ...grpc.CallOption) (*DefaultReply, error) {
	out := new(DefaultReply)
	err := c.cc.Invoke(ctx, "/service.ImageProcessingService/DrawCustomContours", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageProcessingServiceClient) DrawHoughLinesWithParams(ctx context.Context, in *HoughLinesWithPRequest, opts ...grpc.CallOption) (*DefaultReply, error) {
	out := new(DefaultReply)
	err := c.cc.Invoke(ctx, "/service.ImageProcessingService/DrawHoughLinesWithParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageProcessingServiceClient) DrawHoughCircles(ctx context.Context, in *HoughCirclesRequest, opts ...grpc.CallOption) (*DefaultReply, error) {
	out := new(DefaultReply)
	err := c.cc.Invoke(ctx, "/service.ImageProcessingService/DrawHoughCircles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageProcessingServiceClient) Threshold(ctx context.Context, in *ThresholdRequest, opts ...grpc.CallOption) (*DefaultReply, error) {
	out := new(DefaultReply)
	err := c.cc.Invoke(ctx, "/service.ImageProcessingService/Threshold", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageProcessingServiceClient) Watershed(ctx context.Context, in *WatershedRequest, opts ...grpc.CallOption) (*DefaultReply, error) {
	out := new(DefaultReply)
	err := c.cc.Invoke(ctx, "/service.ImageProcessingService/Watershed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageProcessingServiceClient) Open(ctx context.Context, in *OpenRequest, opts ...grpc.CallOption) (*DefaultReply, error) {
	out := new(DefaultReply)
	err := c.cc.Invoke(ctx, "/service.ImageProcessingService/Open", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageProcessingServiceClient) Close(ctx context.Context, in *CloseRequest, opts ...grpc.CallOption) (*DefaultReply, error) {
	out := new(DefaultReply)
	err := c.cc.Invoke(ctx, "/service.ImageProcessingService/Close", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageProcessingServiceClient) FindBlendStructureAmongFabricColorsLUV(ctx context.Context, in *BlendStructureRequest, opts ...grpc.CallOption) (*BlendStructureReply, error) {
	out := new(BlendStructureReply)
	err := c.cc.Invoke(ctx, "/service.ImageProcessingService/FindBlendStructureAmongFabricColorsLUV", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImageProcessingServiceServer is the server API for ImageProcessingService service.
// All implementations must embed UnimplementedImageProcessingServiceServer
// for forward compatibility
type ImageProcessingServiceServer interface {
	PyrMeanShiftFiltering(context.Context, *PyrRequest) (*DefaultReply, error)
	DrawDefaultContours(context.Context, *ContoursRequest) (*DefaultReply, error)
	DrawCustomContours(context.Context, *ContoursRequest) (*DefaultReply, error)
	DrawHoughLinesWithParams(context.Context, *HoughLinesWithPRequest) (*DefaultReply, error)
	DrawHoughCircles(context.Context, *HoughCirclesRequest) (*DefaultReply, error)
	Threshold(context.Context, *ThresholdRequest) (*DefaultReply, error)
	Watershed(context.Context, *WatershedRequest) (*DefaultReply, error)
	Open(context.Context, *OpenRequest) (*DefaultReply, error)
	Close(context.Context, *CloseRequest) (*DefaultReply, error)
	FindBlendStructureAmongFabricColorsLUV(context.Context, *BlendStructureRequest) (*BlendStructureReply, error)
	mustEmbedUnimplementedImageProcessingServiceServer()
}

// UnimplementedImageProcessingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedImageProcessingServiceServer struct {
}

func (UnimplementedImageProcessingServiceServer) PyrMeanShiftFiltering(context.Context, *PyrRequest) (*DefaultReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PyrMeanShiftFiltering not implemented")
}
func (UnimplementedImageProcessingServiceServer) DrawDefaultContours(context.Context, *ContoursRequest) (*DefaultReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DrawDefaultContours not implemented")
}
func (UnimplementedImageProcessingServiceServer) DrawCustomContours(context.Context, *ContoursRequest) (*DefaultReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DrawCustomContours not implemented")
}
func (UnimplementedImageProcessingServiceServer) DrawHoughLinesWithParams(context.Context, *HoughLinesWithPRequest) (*DefaultReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DrawHoughLinesWithParams not implemented")
}
func (UnimplementedImageProcessingServiceServer) DrawHoughCircles(context.Context, *HoughCirclesRequest) (*DefaultReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DrawHoughCircles not implemented")
}
func (UnimplementedImageProcessingServiceServer) Threshold(context.Context, *ThresholdRequest) (*DefaultReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Threshold not implemented")
}
func (UnimplementedImageProcessingServiceServer) Watershed(context.Context, *WatershedRequest) (*DefaultReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Watershed not implemented")
}
func (UnimplementedImageProcessingServiceServer) Open(context.Context, *OpenRequest) (*DefaultReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Open not implemented")
}
func (UnimplementedImageProcessingServiceServer) Close(context.Context, *CloseRequest) (*DefaultReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Close not implemented")
}
func (UnimplementedImageProcessingServiceServer) FindBlendStructureAmongFabricColorsLUV(context.Context, *BlendStructureRequest) (*BlendStructureReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindBlendStructureAmongFabricColorsLUV not implemented")
}
func (UnimplementedImageProcessingServiceServer) mustEmbedUnimplementedImageProcessingServiceServer() {
}

// UnsafeImageProcessingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImageProcessingServiceServer will
// result in compilation errors.
type UnsafeImageProcessingServiceServer interface {
	mustEmbedUnimplementedImageProcessingServiceServer()
}

func RegisterImageProcessingServiceServer(s grpc.ServiceRegistrar, srv ImageProcessingServiceServer) {
	s.RegisterService(&ImageProcessingService_ServiceDesc, srv)
}

func _ImageProcessingService_PyrMeanShiftFiltering_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PyrRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageProcessingServiceServer).PyrMeanShiftFiltering(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.ImageProcessingService/PyrMeanShiftFiltering",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageProcessingServiceServer).PyrMeanShiftFiltering(ctx, req.(*PyrRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageProcessingService_DrawDefaultContours_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContoursRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageProcessingServiceServer).DrawDefaultContours(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.ImageProcessingService/DrawDefaultContours",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageProcessingServiceServer).DrawDefaultContours(ctx, req.(*ContoursRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageProcessingService_DrawCustomContours_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContoursRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageProcessingServiceServer).DrawCustomContours(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.ImageProcessingService/DrawCustomContours",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageProcessingServiceServer).DrawCustomContours(ctx, req.(*ContoursRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageProcessingService_DrawHoughLinesWithParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HoughLinesWithPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageProcessingServiceServer).DrawHoughLinesWithParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.ImageProcessingService/DrawHoughLinesWithParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageProcessingServiceServer).DrawHoughLinesWithParams(ctx, req.(*HoughLinesWithPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageProcessingService_DrawHoughCircles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HoughCirclesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageProcessingServiceServer).DrawHoughCircles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.ImageProcessingService/DrawHoughCircles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageProcessingServiceServer).DrawHoughCircles(ctx, req.(*HoughCirclesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageProcessingService_Threshold_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ThresholdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageProcessingServiceServer).Threshold(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.ImageProcessingService/Threshold",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageProcessingServiceServer).Threshold(ctx, req.(*ThresholdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageProcessingService_Watershed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WatershedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageProcessingServiceServer).Watershed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.ImageProcessingService/Watershed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageProcessingServiceServer).Watershed(ctx, req.(*WatershedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageProcessingService_Open_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageProcessingServiceServer).Open(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.ImageProcessingService/Open",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageProcessingServiceServer).Open(ctx, req.(*OpenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageProcessingService_Close_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageProcessingServiceServer).Close(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.ImageProcessingService/Close",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageProcessingServiceServer).Close(ctx, req.(*CloseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageProcessingService_FindBlendStructureAmongFabricColorsLUV_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlendStructureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageProcessingServiceServer).FindBlendStructureAmongFabricColorsLUV(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.ImageProcessingService/FindBlendStructureAmongFabricColorsLUV",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageProcessingServiceServer).FindBlendStructureAmongFabricColorsLUV(ctx, req.(*BlendStructureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ImageProcessingService_ServiceDesc is the grpc.ServiceDesc for ImageProcessingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ImageProcessingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.ImageProcessingService",
	HandlerType: (*ImageProcessingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PyrMeanShiftFiltering",
			Handler:    _ImageProcessingService_PyrMeanShiftFiltering_Handler,
		},
		{
			MethodName: "DrawDefaultContours",
			Handler:    _ImageProcessingService_DrawDefaultContours_Handler,
		},
		{
			MethodName: "DrawCustomContours",
			Handler:    _ImageProcessingService_DrawCustomContours_Handler,
		},
		{
			MethodName: "DrawHoughLinesWithParams",
			Handler:    _ImageProcessingService_DrawHoughLinesWithParams_Handler,
		},
		{
			MethodName: "DrawHoughCircles",
			Handler:    _ImageProcessingService_DrawHoughCircles_Handler,
		},
		{
			MethodName: "Threshold",
			Handler:    _ImageProcessingService_Threshold_Handler,
		},
		{
			MethodName: "Watershed",
			Handler:    _ImageProcessingService_Watershed_Handler,
		},
		{
			MethodName: "Open",
			Handler:    _ImageProcessingService_Open_Handler,
		},
		{
			MethodName: "Close",
			Handler:    _ImageProcessingService_Close_Handler,
		},
		{
			MethodName: "FindBlendStructureAmongFabricColorsLUV",
			Handler:    _ImageProcessingService_FindBlendStructureAmongFabricColorsLUV_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "image-processing-service.proto",
}
