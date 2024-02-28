// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: pkg/client/pb/client.proto

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

// ClientServiceClient is the client API for ClientService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientServiceClient interface {
	CreateClient(ctx context.Context, in *CreateClientRequest, opts ...grpc.CallOption) (*CreateClientResponse, error)
	GetClient(ctx context.Context, in *GetClientRequest, opts ...grpc.CallOption) (*GetClientResponse, error)
	GetClients(ctx context.Context, in *GetClientsRequest, opts ...grpc.CallOption) (*GetClientsResponse, error)
	GetClientsByCompanyId(ctx context.Context, in *GetClientsByCompanyIdRequest, opts ...grpc.CallOption) (*GetClientsResponse, error)
	UpdateClient(ctx context.Context, in *UpdateClientRequest, opts ...grpc.CallOption) (*UpdateClientResponse, error)
	DeleteClient(ctx context.Context, in *DeleteClientRequest, opts ...grpc.CallOption) (*DeleteClientResponse, error)
}

type clientServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClientServiceClient(cc grpc.ClientConnInterface) ClientServiceClient {
	return &clientServiceClient{cc}
}

func (c *clientServiceClient) CreateClient(ctx context.Context, in *CreateClientRequest, opts ...grpc.CallOption) (*CreateClientResponse, error) {
	out := new(CreateClientResponse)
	err := c.cc.Invoke(ctx, "/pb.ClientService/CreateClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) GetClient(ctx context.Context, in *GetClientRequest, opts ...grpc.CallOption) (*GetClientResponse, error) {
	out := new(GetClientResponse)
	err := c.cc.Invoke(ctx, "/pb.ClientService/GetClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) GetClients(ctx context.Context, in *GetClientsRequest, opts ...grpc.CallOption) (*GetClientsResponse, error) {
	out := new(GetClientsResponse)
	err := c.cc.Invoke(ctx, "/pb.ClientService/GetClients", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) GetClientsByCompanyId(ctx context.Context, in *GetClientsByCompanyIdRequest, opts ...grpc.CallOption) (*GetClientsResponse, error) {
	out := new(GetClientsResponse)
	err := c.cc.Invoke(ctx, "/pb.ClientService/GetClientsByCompanyId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) UpdateClient(ctx context.Context, in *UpdateClientRequest, opts ...grpc.CallOption) (*UpdateClientResponse, error) {
	out := new(UpdateClientResponse)
	err := c.cc.Invoke(ctx, "/pb.ClientService/UpdateClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) DeleteClient(ctx context.Context, in *DeleteClientRequest, opts ...grpc.CallOption) (*DeleteClientResponse, error) {
	out := new(DeleteClientResponse)
	err := c.cc.Invoke(ctx, "/pb.ClientService/DeleteClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientServiceServer is the server API for ClientService service.
// All implementations must embed UnimplementedClientServiceServer
// for forward compatibility
type ClientServiceServer interface {
	CreateClient(context.Context, *CreateClientRequest) (*CreateClientResponse, error)
	GetClient(context.Context, *GetClientRequest) (*GetClientResponse, error)
	GetClients(context.Context, *GetClientsRequest) (*GetClientsResponse, error)
	GetClientsByCompanyId(context.Context, *GetClientsByCompanyIdRequest) (*GetClientsResponse, error)
	UpdateClient(context.Context, *UpdateClientRequest) (*UpdateClientResponse, error)
	DeleteClient(context.Context, *DeleteClientRequest) (*DeleteClientResponse, error)
	mustEmbedUnimplementedClientServiceServer()
}

// UnimplementedClientServiceServer must be embedded to have forward compatible implementations.
type UnimplementedClientServiceServer struct {
}

func (UnimplementedClientServiceServer) CreateClient(context.Context, *CreateClientRequest) (*CreateClientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateClient not implemented")
}
func (UnimplementedClientServiceServer) GetClient(context.Context, *GetClientRequest) (*GetClientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClient not implemented")
}
func (UnimplementedClientServiceServer) GetClients(context.Context, *GetClientsRequest) (*GetClientsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClients not implemented")
}
func (UnimplementedClientServiceServer) GetClientsByCompanyId(context.Context, *GetClientsByCompanyIdRequest) (*GetClientsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClientsByCompanyId not implemented")
}
func (UnimplementedClientServiceServer) UpdateClient(context.Context, *UpdateClientRequest) (*UpdateClientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateClient not implemented")
}
func (UnimplementedClientServiceServer) DeleteClient(context.Context, *DeleteClientRequest) (*DeleteClientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteClient not implemented")
}
func (UnimplementedClientServiceServer) mustEmbedUnimplementedClientServiceServer() {}

// UnsafeClientServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientServiceServer will
// result in compilation errors.
type UnsafeClientServiceServer interface {
	mustEmbedUnimplementedClientServiceServer()
}

func RegisterClientServiceServer(s grpc.ServiceRegistrar, srv ClientServiceServer) {
	s.RegisterService(&ClientService_ServiceDesc, srv)
}

func _ClientService_CreateClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).CreateClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ClientService/CreateClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).CreateClient(ctx, req.(*CreateClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_GetClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).GetClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ClientService/GetClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).GetClient(ctx, req.(*GetClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_GetClients_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClientsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).GetClients(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ClientService/GetClients",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).GetClients(ctx, req.(*GetClientsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_GetClientsByCompanyId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClientsByCompanyIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).GetClientsByCompanyId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ClientService/GetClientsByCompanyId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).GetClientsByCompanyId(ctx, req.(*GetClientsByCompanyIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_UpdateClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).UpdateClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ClientService/UpdateClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).UpdateClient(ctx, req.(*UpdateClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_DeleteClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).DeleteClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ClientService/DeleteClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).DeleteClient(ctx, req.(*DeleteClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ClientService_ServiceDesc is the grpc.ServiceDesc for ClientService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClientService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ClientService",
	HandlerType: (*ClientServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateClient",
			Handler:    _ClientService_CreateClient_Handler,
		},
		{
			MethodName: "GetClient",
			Handler:    _ClientService_GetClient_Handler,
		},
		{
			MethodName: "GetClients",
			Handler:    _ClientService_GetClients_Handler,
		},
		{
			MethodName: "GetClientsByCompanyId",
			Handler:    _ClientService_GetClientsByCompanyId_Handler,
		},
		{
			MethodName: "UpdateClient",
			Handler:    _ClientService_UpdateClient_Handler,
		},
		{
			MethodName: "DeleteClient",
			Handler:    _ClientService_DeleteClient_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/client/pb/client.proto",
}