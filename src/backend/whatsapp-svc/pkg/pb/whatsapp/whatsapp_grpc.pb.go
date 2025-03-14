// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.25.3
// source: proto/whatsapp.proto

package whatsapp

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
	WhatsappService_SendOrderCustomerMessage_FullMethodName    = "/whatsapp.WhatsappService/SendOrderCustomerMessage"
	WhatsappService_SendOrderBusinessMessage_FullMethodName    = "/whatsapp.WhatsappService/SendOrderBusinessMessage"
	WhatsappService_CreateBusinessPhone_FullMethodName         = "/whatsapp.WhatsappService/CreateBusinessPhone"
	WhatsappService_UpdateBusinessPhone_FullMethodName         = "/whatsapp.WhatsappService/UpdateBusinessPhone"
	WhatsappService_GetBusinessPhoneByCompanyId_FullMethodName = "/whatsapp.WhatsappService/GetBusinessPhoneByCompanyId"
)

// WhatsappServiceClient is the client API for WhatsappService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WhatsappServiceClient interface {
	SendOrderCustomerMessage(ctx context.Context, in *SendOrderCustomerMessageRequest, opts ...grpc.CallOption) (*SendOrderCustomerMessageResponse, error)
	SendOrderBusinessMessage(ctx context.Context, in *SendOrderBusinessMessageRequest, opts ...grpc.CallOption) (*SendOrderBusinessMessageResponse, error)
	CreateBusinessPhone(ctx context.Context, in *CreateBusinessPhoneRequest, opts ...grpc.CallOption) (*CreateBusinessPhoneResponse, error)
	UpdateBusinessPhone(ctx context.Context, in *UpdateBusinessPhoneRequest, opts ...grpc.CallOption) (*UpdateBusinessPhoneResponse, error)
	GetBusinessPhoneByCompanyId(ctx context.Context, in *GetBusinessPhoneByCompanyIdRequest, opts ...grpc.CallOption) (*GetBusinessPhoneByCompanyIdResponse, error)
}

type whatsappServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWhatsappServiceClient(cc grpc.ClientConnInterface) WhatsappServiceClient {
	return &whatsappServiceClient{cc}
}

func (c *whatsappServiceClient) SendOrderCustomerMessage(ctx context.Context, in *SendOrderCustomerMessageRequest, opts ...grpc.CallOption) (*SendOrderCustomerMessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendOrderCustomerMessageResponse)
	err := c.cc.Invoke(ctx, WhatsappService_SendOrderCustomerMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *whatsappServiceClient) SendOrderBusinessMessage(ctx context.Context, in *SendOrderBusinessMessageRequest, opts ...grpc.CallOption) (*SendOrderBusinessMessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendOrderBusinessMessageResponse)
	err := c.cc.Invoke(ctx, WhatsappService_SendOrderBusinessMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *whatsappServiceClient) CreateBusinessPhone(ctx context.Context, in *CreateBusinessPhoneRequest, opts ...grpc.CallOption) (*CreateBusinessPhoneResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateBusinessPhoneResponse)
	err := c.cc.Invoke(ctx, WhatsappService_CreateBusinessPhone_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *whatsappServiceClient) UpdateBusinessPhone(ctx context.Context, in *UpdateBusinessPhoneRequest, opts ...grpc.CallOption) (*UpdateBusinessPhoneResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateBusinessPhoneResponse)
	err := c.cc.Invoke(ctx, WhatsappService_UpdateBusinessPhone_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *whatsappServiceClient) GetBusinessPhoneByCompanyId(ctx context.Context, in *GetBusinessPhoneByCompanyIdRequest, opts ...grpc.CallOption) (*GetBusinessPhoneByCompanyIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBusinessPhoneByCompanyIdResponse)
	err := c.cc.Invoke(ctx, WhatsappService_GetBusinessPhoneByCompanyId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WhatsappServiceServer is the server API for WhatsappService service.
// All implementations must embed UnimplementedWhatsappServiceServer
// for forward compatibility.
type WhatsappServiceServer interface {
	SendOrderCustomerMessage(context.Context, *SendOrderCustomerMessageRequest) (*SendOrderCustomerMessageResponse, error)
	SendOrderBusinessMessage(context.Context, *SendOrderBusinessMessageRequest) (*SendOrderBusinessMessageResponse, error)
	CreateBusinessPhone(context.Context, *CreateBusinessPhoneRequest) (*CreateBusinessPhoneResponse, error)
	UpdateBusinessPhone(context.Context, *UpdateBusinessPhoneRequest) (*UpdateBusinessPhoneResponse, error)
	GetBusinessPhoneByCompanyId(context.Context, *GetBusinessPhoneByCompanyIdRequest) (*GetBusinessPhoneByCompanyIdResponse, error)
	mustEmbedUnimplementedWhatsappServiceServer()
}

// UnimplementedWhatsappServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedWhatsappServiceServer struct{}

func (UnimplementedWhatsappServiceServer) SendOrderCustomerMessage(context.Context, *SendOrderCustomerMessageRequest) (*SendOrderCustomerMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendOrderCustomerMessage not implemented")
}
func (UnimplementedWhatsappServiceServer) SendOrderBusinessMessage(context.Context, *SendOrderBusinessMessageRequest) (*SendOrderBusinessMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendOrderBusinessMessage not implemented")
}
func (UnimplementedWhatsappServiceServer) CreateBusinessPhone(context.Context, *CreateBusinessPhoneRequest) (*CreateBusinessPhoneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBusinessPhone not implemented")
}
func (UnimplementedWhatsappServiceServer) UpdateBusinessPhone(context.Context, *UpdateBusinessPhoneRequest) (*UpdateBusinessPhoneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBusinessPhone not implemented")
}
func (UnimplementedWhatsappServiceServer) GetBusinessPhoneByCompanyId(context.Context, *GetBusinessPhoneByCompanyIdRequest) (*GetBusinessPhoneByCompanyIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBusinessPhoneByCompanyId not implemented")
}
func (UnimplementedWhatsappServiceServer) mustEmbedUnimplementedWhatsappServiceServer() {}
func (UnimplementedWhatsappServiceServer) testEmbeddedByValue()                         {}

// UnsafeWhatsappServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WhatsappServiceServer will
// result in compilation errors.
type UnsafeWhatsappServiceServer interface {
	mustEmbedUnimplementedWhatsappServiceServer()
}

func RegisterWhatsappServiceServer(s grpc.ServiceRegistrar, srv WhatsappServiceServer) {
	// If the following call pancis, it indicates UnimplementedWhatsappServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&WhatsappService_ServiceDesc, srv)
}

func _WhatsappService_SendOrderCustomerMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendOrderCustomerMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WhatsappServiceServer).SendOrderCustomerMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WhatsappService_SendOrderCustomerMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WhatsappServiceServer).SendOrderCustomerMessage(ctx, req.(*SendOrderCustomerMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WhatsappService_SendOrderBusinessMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendOrderBusinessMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WhatsappServiceServer).SendOrderBusinessMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WhatsappService_SendOrderBusinessMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WhatsappServiceServer).SendOrderBusinessMessage(ctx, req.(*SendOrderBusinessMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WhatsappService_CreateBusinessPhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBusinessPhoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WhatsappServiceServer).CreateBusinessPhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WhatsappService_CreateBusinessPhone_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WhatsappServiceServer).CreateBusinessPhone(ctx, req.(*CreateBusinessPhoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WhatsappService_UpdateBusinessPhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBusinessPhoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WhatsappServiceServer).UpdateBusinessPhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WhatsappService_UpdateBusinessPhone_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WhatsappServiceServer).UpdateBusinessPhone(ctx, req.(*UpdateBusinessPhoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WhatsappService_GetBusinessPhoneByCompanyId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBusinessPhoneByCompanyIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WhatsappServiceServer).GetBusinessPhoneByCompanyId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WhatsappService_GetBusinessPhoneByCompanyId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WhatsappServiceServer).GetBusinessPhoneByCompanyId(ctx, req.(*GetBusinessPhoneByCompanyIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WhatsappService_ServiceDesc is the grpc.ServiceDesc for WhatsappService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WhatsappService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "whatsapp.WhatsappService",
	HandlerType: (*WhatsappServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendOrderCustomerMessage",
			Handler:    _WhatsappService_SendOrderCustomerMessage_Handler,
		},
		{
			MethodName: "SendOrderBusinessMessage",
			Handler:    _WhatsappService_SendOrderBusinessMessage_Handler,
		},
		{
			MethodName: "CreateBusinessPhone",
			Handler:    _WhatsappService_CreateBusinessPhone_Handler,
		},
		{
			MethodName: "UpdateBusinessPhone",
			Handler:    _WhatsappService_UpdateBusinessPhone_Handler,
		},
		{
			MethodName: "GetBusinessPhoneByCompanyId",
			Handler:    _WhatsappService_GetBusinessPhoneByCompanyId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/whatsapp.proto",
}
