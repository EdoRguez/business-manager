// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.27.1
// source: pkg/company/pb/payment_type.proto

package pb

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
	PaymentTypeService_GetPaymentType_FullMethodName  = "/pb.PaymentTypeService/GetPaymentType"
	PaymentTypeService_GetPaymentTypes_FullMethodName = "/pb.PaymentTypeService/GetPaymentTypes"
)

// PaymentTypeServiceClient is the client API for PaymentTypeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PaymentTypeServiceClient interface {
	GetPaymentType(ctx context.Context, in *GetPaymentTypeRequest, opts ...grpc.CallOption) (*GetPaymentTypeResponse, error)
	GetPaymentTypes(ctx context.Context, in *GetPaymentTypesRequest, opts ...grpc.CallOption) (*GetPaymentTypesResponse, error)
}

type paymentTypeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymentTypeServiceClient(cc grpc.ClientConnInterface) PaymentTypeServiceClient {
	return &paymentTypeServiceClient{cc}
}

func (c *paymentTypeServiceClient) GetPaymentType(ctx context.Context, in *GetPaymentTypeRequest, opts ...grpc.CallOption) (*GetPaymentTypeResponse, error) {
	out := new(GetPaymentTypeResponse)
	err := c.cc.Invoke(ctx, PaymentTypeService_GetPaymentType_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentTypeServiceClient) GetPaymentTypes(ctx context.Context, in *GetPaymentTypesRequest, opts ...grpc.CallOption) (*GetPaymentTypesResponse, error) {
	out := new(GetPaymentTypesResponse)
	err := c.cc.Invoke(ctx, PaymentTypeService_GetPaymentTypes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentTypeServiceServer is the server API for PaymentTypeService service.
// All implementations must embed UnimplementedPaymentTypeServiceServer
// for forward compatibility
type PaymentTypeServiceServer interface {
	GetPaymentType(context.Context, *GetPaymentTypeRequest) (*GetPaymentTypeResponse, error)
	GetPaymentTypes(context.Context, *GetPaymentTypesRequest) (*GetPaymentTypesResponse, error)
	mustEmbedUnimplementedPaymentTypeServiceServer()
}

// UnimplementedPaymentTypeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPaymentTypeServiceServer struct {
}

func (UnimplementedPaymentTypeServiceServer) GetPaymentType(context.Context, *GetPaymentTypeRequest) (*GetPaymentTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPaymentType not implemented")
}
func (UnimplementedPaymentTypeServiceServer) GetPaymentTypes(context.Context, *GetPaymentTypesRequest) (*GetPaymentTypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPaymentTypes not implemented")
}
func (UnimplementedPaymentTypeServiceServer) mustEmbedUnimplementedPaymentTypeServiceServer() {}

// UnsafePaymentTypeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PaymentTypeServiceServer will
// result in compilation errors.
type UnsafePaymentTypeServiceServer interface {
	mustEmbedUnimplementedPaymentTypeServiceServer()
}

func RegisterPaymentTypeServiceServer(s grpc.ServiceRegistrar, srv PaymentTypeServiceServer) {
	s.RegisterService(&PaymentTypeService_ServiceDesc, srv)
}

func _PaymentTypeService_GetPaymentType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPaymentTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentTypeServiceServer).GetPaymentType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentTypeService_GetPaymentType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentTypeServiceServer).GetPaymentType(ctx, req.(*GetPaymentTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentTypeService_GetPaymentTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPaymentTypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentTypeServiceServer).GetPaymentTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentTypeService_GetPaymentTypes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentTypeServiceServer).GetPaymentTypes(ctx, req.(*GetPaymentTypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PaymentTypeService_ServiceDesc is the grpc.ServiceDesc for PaymentTypeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PaymentTypeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.PaymentTypeService",
	HandlerType: (*PaymentTypeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPaymentType",
			Handler:    _PaymentTypeService_GetPaymentType_Handler,
		},
		{
			MethodName: "GetPaymentTypes",
			Handler:    _PaymentTypeService_GetPaymentTypes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/company/pb/payment_type.proto",
}
