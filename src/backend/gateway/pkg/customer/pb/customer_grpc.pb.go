// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: pkg/customer/pb/customer.proto

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

// CustomerServiceClient is the client API for CustomerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CustomerServiceClient interface {
	CreateCustomer(ctx context.Context, in *CreateCustomerRequest, opts ...grpc.CallOption) (*CreateCustomerResponse, error)
	GetCustomer(ctx context.Context, in *GetCustomerRequest, opts ...grpc.CallOption) (*GetCustomerResponse, error)
	GetCustomers(ctx context.Context, in *GetCustomersRequest, opts ...grpc.CallOption) (*GetCustomersResponse, error)
	GetCustomersByMonths(ctx context.Context, in *GetCustomersByMonthsRequest, opts ...grpc.CallOption) (*GetCustomersByMonthsResponse, error)
	UpdateCustomer(ctx context.Context, in *UpdateCustomerRequest, opts ...grpc.CallOption) (*UpdateCustomerResponse, error)
	DeleteCustomer(ctx context.Context, in *DeleteCustomerRequest, opts ...grpc.CallOption) (*DeleteCustomerResponse, error)
}

type customerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerServiceClient(cc grpc.ClientConnInterface) CustomerServiceClient {
	return &customerServiceClient{cc}
}

func (c *customerServiceClient) CreateCustomer(ctx context.Context, in *CreateCustomerRequest, opts ...grpc.CallOption) (*CreateCustomerResponse, error) {
	out := new(CreateCustomerResponse)
	err := c.cc.Invoke(ctx, "/pb.CustomerService/CreateCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) GetCustomer(ctx context.Context, in *GetCustomerRequest, opts ...grpc.CallOption) (*GetCustomerResponse, error) {
	out := new(GetCustomerResponse)
	err := c.cc.Invoke(ctx, "/pb.CustomerService/GetCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) GetCustomers(ctx context.Context, in *GetCustomersRequest, opts ...grpc.CallOption) (*GetCustomersResponse, error) {
	out := new(GetCustomersResponse)
	err := c.cc.Invoke(ctx, "/pb.CustomerService/GetCustomers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) GetCustomersByMonths(ctx context.Context, in *GetCustomersByMonthsRequest, opts ...grpc.CallOption) (*GetCustomersByMonthsResponse, error) {
	out := new(GetCustomersByMonthsResponse)
	err := c.cc.Invoke(ctx, "/pb.CustomerService/GetCustomersByMonths", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) UpdateCustomer(ctx context.Context, in *UpdateCustomerRequest, opts ...grpc.CallOption) (*UpdateCustomerResponse, error) {
	out := new(UpdateCustomerResponse)
	err := c.cc.Invoke(ctx, "/pb.CustomerService/UpdateCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) DeleteCustomer(ctx context.Context, in *DeleteCustomerRequest, opts ...grpc.CallOption) (*DeleteCustomerResponse, error) {
	out := new(DeleteCustomerResponse)
	err := c.cc.Invoke(ctx, "/pb.CustomerService/DeleteCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerServiceServer is the server API for CustomerService service.
// All implementations must embed UnimplementedCustomerServiceServer
// for forward compatibility
type CustomerServiceServer interface {
	CreateCustomer(context.Context, *CreateCustomerRequest) (*CreateCustomerResponse, error)
	GetCustomer(context.Context, *GetCustomerRequest) (*GetCustomerResponse, error)
	GetCustomers(context.Context, *GetCustomersRequest) (*GetCustomersResponse, error)
	GetCustomersByMonths(context.Context, *GetCustomersByMonthsRequest) (*GetCustomersByMonthsResponse, error)
	UpdateCustomer(context.Context, *UpdateCustomerRequest) (*UpdateCustomerResponse, error)
	DeleteCustomer(context.Context, *DeleteCustomerRequest) (*DeleteCustomerResponse, error)
	mustEmbedUnimplementedCustomerServiceServer()
}

// UnimplementedCustomerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCustomerServiceServer struct {
}

func (UnimplementedCustomerServiceServer) CreateCustomer(context.Context, *CreateCustomerRequest) (*CreateCustomerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCustomer not implemented")
}
func (UnimplementedCustomerServiceServer) GetCustomer(context.Context, *GetCustomerRequest) (*GetCustomerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomer not implemented")
}
func (UnimplementedCustomerServiceServer) GetCustomers(context.Context, *GetCustomersRequest) (*GetCustomersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomers not implemented")
}
func (UnimplementedCustomerServiceServer) GetCustomersByMonths(context.Context, *GetCustomersByMonthsRequest) (*GetCustomersByMonthsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomersByMonths not implemented")
}
func (UnimplementedCustomerServiceServer) UpdateCustomer(context.Context, *UpdateCustomerRequest) (*UpdateCustomerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCustomer not implemented")
}
func (UnimplementedCustomerServiceServer) DeleteCustomer(context.Context, *DeleteCustomerRequest) (*DeleteCustomerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCustomer not implemented")
}
func (UnimplementedCustomerServiceServer) mustEmbedUnimplementedCustomerServiceServer() {}

// UnsafeCustomerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CustomerServiceServer will
// result in compilation errors.
type UnsafeCustomerServiceServer interface {
	mustEmbedUnimplementedCustomerServiceServer()
}

func RegisterCustomerServiceServer(s grpc.ServiceRegistrar, srv CustomerServiceServer) {
	s.RegisterService(&CustomerService_ServiceDesc, srv)
}

func _CustomerService_CreateCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).CreateCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CustomerService/CreateCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).CreateCustomer(ctx, req.(*CreateCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_GetCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).GetCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CustomerService/GetCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).GetCustomer(ctx, req.(*GetCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_GetCustomers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).GetCustomers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CustomerService/GetCustomers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).GetCustomers(ctx, req.(*GetCustomersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_GetCustomersByMonths_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomersByMonthsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).GetCustomersByMonths(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CustomerService/GetCustomersByMonths",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).GetCustomersByMonths(ctx, req.(*GetCustomersByMonthsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_UpdateCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).UpdateCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CustomerService/UpdateCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).UpdateCustomer(ctx, req.(*UpdateCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_DeleteCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).DeleteCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CustomerService/DeleteCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).DeleteCustomer(ctx, req.(*DeleteCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CustomerService_ServiceDesc is the grpc.ServiceDesc for CustomerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CustomerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.CustomerService",
	HandlerType: (*CustomerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCustomer",
			Handler:    _CustomerService_CreateCustomer_Handler,
		},
		{
			MethodName: "GetCustomer",
			Handler:    _CustomerService_GetCustomer_Handler,
		},
		{
			MethodName: "GetCustomers",
			Handler:    _CustomerService_GetCustomers_Handler,
		},
		{
			MethodName: "GetCustomersByMonths",
			Handler:    _CustomerService_GetCustomersByMonths_Handler,
		},
		{
			MethodName: "UpdateCustomer",
			Handler:    _CustomerService_UpdateCustomer_Handler,
		},
		{
			MethodName: "DeleteCustomer",
			Handler:    _CustomerService_DeleteCustomer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/customer/pb/customer.proto",
}
