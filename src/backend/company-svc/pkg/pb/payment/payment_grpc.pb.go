// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.25.3
// source: proto/payment.proto

package payment

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	PaymentService_CreatePayment_FullMethodName       = "/payment.PaymentService/CreatePayment"
	PaymentService_GetPayment_FullMethodName          = "/payment.PaymentService/GetPayment"
	PaymentService_GetPayments_FullMethodName         = "/payment.PaymentService/GetPayments"
	PaymentService_GetPaymentsTypes_FullMethodName    = "/payment.PaymentService/GetPaymentsTypes"
	PaymentService_UpdatePayment_FullMethodName       = "/payment.PaymentService/UpdatePayment"
	PaymentService_UpdatePaymentStatus_FullMethodName = "/payment.PaymentService/UpdatePaymentStatus"
	PaymentService_DeletePayment_FullMethodName       = "/payment.PaymentService/DeletePayment"
)

// PaymentServiceClient is the client API for PaymentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PaymentServiceClient interface {
	CreatePayment(ctx context.Context, in *CreatePaymentRequest, opts ...grpc.CallOption) (*CreatePaymentResponse, error)
	GetPayment(ctx context.Context, in *GetPaymentRequest, opts ...grpc.CallOption) (*GetPaymentResponse, error)
	GetPayments(ctx context.Context, in *GetPaymentsRequest, opts ...grpc.CallOption) (*GetPaymentsResponse, error)
	GetPaymentsTypes(ctx context.Context, in *GetPaymentsTypesRequest, opts ...grpc.CallOption) (*GetPaymentsTypesResponse, error)
	UpdatePayment(ctx context.Context, in *UpdatePaymentRequest, opts ...grpc.CallOption) (*UpdatePaymentResponse, error)
	UpdatePaymentStatus(ctx context.Context, in *UpdatePaymentStatusRequest, opts ...grpc.CallOption) (*UpdatePaymentStatusResponse, error)
	DeletePayment(ctx context.Context, in *DeletePaymentRequest, opts ...grpc.CallOption) (*DeletePaymentResponse, error)
}

type paymentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymentServiceClient(cc grpc.ClientConnInterface) PaymentServiceClient {
	return &paymentServiceClient{cc}
}

func (c *paymentServiceClient) CreatePayment(ctx context.Context, in *CreatePaymentRequest, opts ...grpc.CallOption) (*CreatePaymentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreatePaymentResponse)
	err := c.cc.Invoke(ctx, PaymentService_CreatePayment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) GetPayment(ctx context.Context, in *GetPaymentRequest, opts ...grpc.CallOption) (*GetPaymentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPaymentResponse)
	err := c.cc.Invoke(ctx, PaymentService_GetPayment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) GetPayments(ctx context.Context, in *GetPaymentsRequest, opts ...grpc.CallOption) (*GetPaymentsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPaymentsResponse)
	err := c.cc.Invoke(ctx, PaymentService_GetPayments_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) GetPaymentsTypes(ctx context.Context, in *GetPaymentsTypesRequest, opts ...grpc.CallOption) (*GetPaymentsTypesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPaymentsTypesResponse)
	err := c.cc.Invoke(ctx, PaymentService_GetPaymentsTypes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) UpdatePayment(ctx context.Context, in *UpdatePaymentRequest, opts ...grpc.CallOption) (*UpdatePaymentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdatePaymentResponse)
	err := c.cc.Invoke(ctx, PaymentService_UpdatePayment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) UpdatePaymentStatus(ctx context.Context, in *UpdatePaymentStatusRequest, opts ...grpc.CallOption) (*UpdatePaymentStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdatePaymentStatusResponse)
	err := c.cc.Invoke(ctx, PaymentService_UpdatePaymentStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) DeletePayment(ctx context.Context, in *DeletePaymentRequest, opts ...grpc.CallOption) (*DeletePaymentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeletePaymentResponse)
	err := c.cc.Invoke(ctx, PaymentService_DeletePayment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentServiceServer is the server API for PaymentService service.
// All implementations must embed UnimplementedPaymentServiceServer
// for forward compatibility.
type PaymentServiceServer interface {
	CreatePayment(context.Context, *CreatePaymentRequest) (*CreatePaymentResponse, error)
	GetPayment(context.Context, *GetPaymentRequest) (*GetPaymentResponse, error)
	GetPayments(context.Context, *GetPaymentsRequest) (*GetPaymentsResponse, error)
	GetPaymentsTypes(context.Context, *GetPaymentsTypesRequest) (*GetPaymentsTypesResponse, error)
	UpdatePayment(context.Context, *UpdatePaymentRequest) (*UpdatePaymentResponse, error)
	UpdatePaymentStatus(context.Context, *UpdatePaymentStatusRequest) (*UpdatePaymentStatusResponse, error)
	DeletePayment(context.Context, *DeletePaymentRequest) (*DeletePaymentResponse, error)
	mustEmbedUnimplementedPaymentServiceServer()
}

// UnimplementedPaymentServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPaymentServiceServer struct{}

func (UnimplementedPaymentServiceServer) CreatePayment(context.Context, *CreatePaymentRequest) (*CreatePaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePayment not implemented")
}
func (UnimplementedPaymentServiceServer) GetPayment(context.Context, *GetPaymentRequest) (*GetPaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPayment not implemented")
}
func (UnimplementedPaymentServiceServer) GetPayments(context.Context, *GetPaymentsRequest) (*GetPaymentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPayments not implemented")
}
func (UnimplementedPaymentServiceServer) GetPaymentsTypes(context.Context, *GetPaymentsTypesRequest) (*GetPaymentsTypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPaymentsTypes not implemented")
}
func (UnimplementedPaymentServiceServer) UpdatePayment(context.Context, *UpdatePaymentRequest) (*UpdatePaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePayment not implemented")
}
func (UnimplementedPaymentServiceServer) UpdatePaymentStatus(context.Context, *UpdatePaymentStatusRequest) (*UpdatePaymentStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePaymentStatus not implemented")
}
func (UnimplementedPaymentServiceServer) DeletePayment(context.Context, *DeletePaymentRequest) (*DeletePaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePayment not implemented")
}
func (UnimplementedPaymentServiceServer) mustEmbedUnimplementedPaymentServiceServer() {}
func (UnimplementedPaymentServiceServer) testEmbeddedByValue()                        {}

// UnsafePaymentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PaymentServiceServer will
// result in compilation errors.
type UnsafePaymentServiceServer interface {
	mustEmbedUnimplementedPaymentServiceServer()
}

func RegisterPaymentServiceServer(s grpc.ServiceRegistrar, srv PaymentServiceServer) {
	// If the following call pancis, it indicates UnimplementedPaymentServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PaymentService_ServiceDesc, srv)
}

func _PaymentService_CreatePayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).CreatePayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_CreatePayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).CreatePayment(ctx, req.(*CreatePaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_GetPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).GetPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_GetPayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).GetPayment(ctx, req.(*GetPaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_GetPayments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPaymentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).GetPayments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_GetPayments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).GetPayments(ctx, req.(*GetPaymentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_GetPaymentsTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPaymentsTypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).GetPaymentsTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_GetPaymentsTypes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).GetPaymentsTypes(ctx, req.(*GetPaymentsTypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_UpdatePayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).UpdatePayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_UpdatePayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).UpdatePayment(ctx, req.(*UpdatePaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_UpdatePaymentStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePaymentStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).UpdatePaymentStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_UpdatePaymentStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).UpdatePaymentStatus(ctx, req.(*UpdatePaymentStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_DeletePayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).DeletePayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_DeletePayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).DeletePayment(ctx, req.(*DeletePaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PaymentService_ServiceDesc is the grpc.ServiceDesc for PaymentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PaymentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "payment.PaymentService",
	HandlerType: (*PaymentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePayment",
			Handler:    _PaymentService_CreatePayment_Handler,
		},
		{
			MethodName: "GetPayment",
			Handler:    _PaymentService_GetPayment_Handler,
		},
		{
			MethodName: "GetPayments",
			Handler:    _PaymentService_GetPayments_Handler,
		},
		{
			MethodName: "GetPaymentsTypes",
			Handler:    _PaymentService_GetPaymentsTypes_Handler,
		},
		{
			MethodName: "UpdatePayment",
			Handler:    _PaymentService_UpdatePayment_Handler,
		},
		{
			MethodName: "UpdatePaymentStatus",
			Handler:    _PaymentService_UpdatePaymentStatus_Handler,
		},
		{
			MethodName: "DeletePayment",
			Handler:    _PaymentService_DeletePayment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/payment.proto",
}
