// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: dataextractor.proto

package dataextractorpb

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

// DataExtractorClient is the client API for DataExtractor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataExtractorClient interface {
	ExtractText(ctx context.Context, in *ExtractTextRequest, opts ...grpc.CallOption) (*ExtractTextResponse, error)
	ExtractDigits(ctx context.Context, in *ExtractDigitsRequest, opts ...grpc.CallOption) (*ExtractDigitsResponse, error)
	ExtractLabeledValues(ctx context.Context, in *ExtractLabeledValuesRequest, opts ...grpc.CallOption) (*ExtractLabeledValuesResponse, error)
}

type dataExtractorClient struct {
	cc grpc.ClientConnInterface
}

func NewDataExtractorClient(cc grpc.ClientConnInterface) DataExtractorClient {
	return &dataExtractorClient{cc}
}

func (c *dataExtractorClient) ExtractText(ctx context.Context, in *ExtractTextRequest, opts ...grpc.CallOption) (*ExtractTextResponse, error) {
	out := new(ExtractTextResponse)
	err := c.cc.Invoke(ctx, "/DataExtractor/ExtractText", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataExtractorClient) ExtractDigits(ctx context.Context, in *ExtractDigitsRequest, opts ...grpc.CallOption) (*ExtractDigitsResponse, error) {
	out := new(ExtractDigitsResponse)
	err := c.cc.Invoke(ctx, "/DataExtractor/ExtractDigits", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataExtractorClient) ExtractLabeledValues(ctx context.Context, in *ExtractLabeledValuesRequest, opts ...grpc.CallOption) (*ExtractLabeledValuesResponse, error) {
	out := new(ExtractLabeledValuesResponse)
	err := c.cc.Invoke(ctx, "/DataExtractor/ExtractLabeledValues", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataExtractorServer is the server API for DataExtractor service.
// All implementations must embed UnimplementedDataExtractorServer
// for forward compatibility
type DataExtractorServer interface {
	ExtractText(context.Context, *ExtractTextRequest) (*ExtractTextResponse, error)
	ExtractDigits(context.Context, *ExtractDigitsRequest) (*ExtractDigitsResponse, error)
	ExtractLabeledValues(context.Context, *ExtractLabeledValuesRequest) (*ExtractLabeledValuesResponse, error)
	mustEmbedUnimplementedDataExtractorServer()
}

// UnimplementedDataExtractorServer must be embedded to have forward compatible implementations.
type UnimplementedDataExtractorServer struct {
}

func (UnimplementedDataExtractorServer) ExtractText(context.Context, *ExtractTextRequest) (*ExtractTextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExtractText not implemented")
}
func (UnimplementedDataExtractorServer) ExtractDigits(context.Context, *ExtractDigitsRequest) (*ExtractDigitsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExtractDigits not implemented")
}
func (UnimplementedDataExtractorServer) ExtractLabeledValues(context.Context, *ExtractLabeledValuesRequest) (*ExtractLabeledValuesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExtractLabeledValues not implemented")
}
func (UnimplementedDataExtractorServer) mustEmbedUnimplementedDataExtractorServer() {}

// UnsafeDataExtractorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataExtractorServer will
// result in compilation errors.
type UnsafeDataExtractorServer interface {
	mustEmbedUnimplementedDataExtractorServer()
}

func RegisterDataExtractorServer(s grpc.ServiceRegistrar, srv DataExtractorServer) {
	s.RegisterService(&DataExtractor_ServiceDesc, srv)
}

func _DataExtractor_ExtractText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExtractTextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataExtractorServer).ExtractText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DataExtractor/ExtractText",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataExtractorServer).ExtractText(ctx, req.(*ExtractTextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataExtractor_ExtractDigits_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExtractDigitsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataExtractorServer).ExtractDigits(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DataExtractor/ExtractDigits",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataExtractorServer).ExtractDigits(ctx, req.(*ExtractDigitsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataExtractor_ExtractLabeledValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExtractLabeledValuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataExtractorServer).ExtractLabeledValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DataExtractor/ExtractLabeledValues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataExtractorServer).ExtractLabeledValues(ctx, req.(*ExtractLabeledValuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DataExtractor_ServiceDesc is the grpc.ServiceDesc for DataExtractor service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataExtractor_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "DataExtractor",
	HandlerType: (*DataExtractorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExtractText",
			Handler:    _DataExtractor_ExtractText_Handler,
		},
		{
			MethodName: "ExtractDigits",
			Handler:    _DataExtractor_ExtractDigits_Handler,
		},
		{
			MethodName: "ExtractLabeledValues",
			Handler:    _DataExtractor_ExtractLabeledValues_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dataextractor.proto",
}
