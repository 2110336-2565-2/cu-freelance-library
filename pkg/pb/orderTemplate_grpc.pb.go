// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: orderTemplate.proto

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

// OrderTemplateServiceClient is the client API for OrderTemplateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderTemplateServiceClient interface {
	GetOrderTemplatePaginate(ctx context.Context, in *GetOrderTemplatePaginateRequest, opts ...grpc.CallOption) (*OrderTemplatePaginateResponse, error)
	GetOrderTemplateByID(ctx context.Context, in *GetOrderTemplateByIDRequest, opts ...grpc.CallOption) (*GetOrderTemplateByIDResponse, error)
	CreateOrderTemplate(ctx context.Context, in *CreateOrderTemplateRequest, opts ...grpc.CallOption) (*CreateOrderTemplateResponse, error)
	UpdateOrderTemplateByID(ctx context.Context, in *UpdateOrderTemplateRequest, opts ...grpc.CallOption) (*UpdateOrderTemplateResponse, error)
	DeleteOrderTemplateByID(ctx context.Context, in *DeleteOrderTemplateRequest, opts ...grpc.CallOption) (*DeleteOrderTemplateResponse, error)
}

type orderTemplateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderTemplateServiceClient(cc grpc.ClientConnInterface) OrderTemplateServiceClient {
	return &orderTemplateServiceClient{cc}
}

func (c *orderTemplateServiceClient) GetOrderTemplatePaginate(ctx context.Context, in *GetOrderTemplatePaginateRequest, opts ...grpc.CallOption) (*OrderTemplatePaginateResponse, error) {
	out := new(OrderTemplatePaginateResponse)
	err := c.cc.Invoke(ctx, "/orderTemplate.OrderTemplateService/GetOrderTemplatePaginate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderTemplateServiceClient) GetOrderTemplateByID(ctx context.Context, in *GetOrderTemplateByIDRequest, opts ...grpc.CallOption) (*GetOrderTemplateByIDResponse, error) {
	out := new(GetOrderTemplateByIDResponse)
	err := c.cc.Invoke(ctx, "/orderTemplate.OrderTemplateService/GetOrderTemplateByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderTemplateServiceClient) CreateOrderTemplate(ctx context.Context, in *CreateOrderTemplateRequest, opts ...grpc.CallOption) (*CreateOrderTemplateResponse, error) {
	out := new(CreateOrderTemplateResponse)
	err := c.cc.Invoke(ctx, "/orderTemplate.OrderTemplateService/CreateOrderTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderTemplateServiceClient) UpdateOrderTemplateByID(ctx context.Context, in *UpdateOrderTemplateRequest, opts ...grpc.CallOption) (*UpdateOrderTemplateResponse, error) {
	out := new(UpdateOrderTemplateResponse)
	err := c.cc.Invoke(ctx, "/orderTemplate.OrderTemplateService/UpdateOrderTemplateByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderTemplateServiceClient) DeleteOrderTemplateByID(ctx context.Context, in *DeleteOrderTemplateRequest, opts ...grpc.CallOption) (*DeleteOrderTemplateResponse, error) {
	out := new(DeleteOrderTemplateResponse)
	err := c.cc.Invoke(ctx, "/orderTemplate.OrderTemplateService/DeleteOrderTemplateByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderTemplateServiceServer is the server API for OrderTemplateService service.
// All implementations should embed UnimplementedOrderTemplateServiceServer
// for forward compatibility
type OrderTemplateServiceServer interface {
	GetOrderTemplatePaginate(context.Context, *GetOrderTemplatePaginateRequest) (*OrderTemplatePaginateResponse, error)
	GetOrderTemplateByID(context.Context, *GetOrderTemplateByIDRequest) (*GetOrderTemplateByIDResponse, error)
	CreateOrderTemplate(context.Context, *CreateOrderTemplateRequest) (*CreateOrderTemplateResponse, error)
	UpdateOrderTemplateByID(context.Context, *UpdateOrderTemplateRequest) (*UpdateOrderTemplateResponse, error)
	DeleteOrderTemplateByID(context.Context, *DeleteOrderTemplateRequest) (*DeleteOrderTemplateResponse, error)
}

// UnimplementedOrderTemplateServiceServer should be embedded to have forward compatible implementations.
type UnimplementedOrderTemplateServiceServer struct {
}

func (UnimplementedOrderTemplateServiceServer) GetOrderTemplatePaginate(context.Context, *GetOrderTemplatePaginateRequest) (*OrderTemplatePaginateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderTemplatePaginate not implemented")
}
func (UnimplementedOrderTemplateServiceServer) GetOrderTemplateByID(context.Context, *GetOrderTemplateByIDRequest) (*GetOrderTemplateByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderTemplateByID not implemented")
}
func (UnimplementedOrderTemplateServiceServer) CreateOrderTemplate(context.Context, *CreateOrderTemplateRequest) (*CreateOrderTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrderTemplate not implemented")
}
func (UnimplementedOrderTemplateServiceServer) UpdateOrderTemplateByID(context.Context, *UpdateOrderTemplateRequest) (*UpdateOrderTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrderTemplateByID not implemented")
}
func (UnimplementedOrderTemplateServiceServer) DeleteOrderTemplateByID(context.Context, *DeleteOrderTemplateRequest) (*DeleteOrderTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOrderTemplateByID not implemented")
}

// UnsafeOrderTemplateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderTemplateServiceServer will
// result in compilation errors.
type UnsafeOrderTemplateServiceServer interface {
	mustEmbedUnimplementedOrderTemplateServiceServer()
}

func RegisterOrderTemplateServiceServer(s grpc.ServiceRegistrar, srv OrderTemplateServiceServer) {
	s.RegisterService(&OrderTemplateService_ServiceDesc, srv)
}

func _OrderTemplateService_GetOrderTemplatePaginate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderTemplatePaginateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderTemplateServiceServer).GetOrderTemplatePaginate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orderTemplate.OrderTemplateService/GetOrderTemplatePaginate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderTemplateServiceServer).GetOrderTemplatePaginate(ctx, req.(*GetOrderTemplatePaginateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderTemplateService_GetOrderTemplateByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderTemplateByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderTemplateServiceServer).GetOrderTemplateByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orderTemplate.OrderTemplateService/GetOrderTemplateByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderTemplateServiceServer).GetOrderTemplateByID(ctx, req.(*GetOrderTemplateByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderTemplateService_CreateOrderTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderTemplateServiceServer).CreateOrderTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orderTemplate.OrderTemplateService/CreateOrderTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderTemplateServiceServer).CreateOrderTemplate(ctx, req.(*CreateOrderTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderTemplateService_UpdateOrderTemplateByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrderTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderTemplateServiceServer).UpdateOrderTemplateByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orderTemplate.OrderTemplateService/UpdateOrderTemplateByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderTemplateServiceServer).UpdateOrderTemplateByID(ctx, req.(*UpdateOrderTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderTemplateService_DeleteOrderTemplateByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOrderTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderTemplateServiceServer).DeleteOrderTemplateByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orderTemplate.OrderTemplateService/DeleteOrderTemplateByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderTemplateServiceServer).DeleteOrderTemplateByID(ctx, req.(*DeleteOrderTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderTemplateService_ServiceDesc is the grpc.ServiceDesc for OrderTemplateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderTemplateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "orderTemplate.OrderTemplateService",
	HandlerType: (*OrderTemplateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOrderTemplatePaginate",
			Handler:    _OrderTemplateService_GetOrderTemplatePaginate_Handler,
		},
		{
			MethodName: "GetOrderTemplateByID",
			Handler:    _OrderTemplateService_GetOrderTemplateByID_Handler,
		},
		{
			MethodName: "CreateOrderTemplate",
			Handler:    _OrderTemplateService_CreateOrderTemplate_Handler,
		},
		{
			MethodName: "UpdateOrderTemplateByID",
			Handler:    _OrderTemplateService_UpdateOrderTemplateByID_Handler,
		},
		{
			MethodName: "DeleteOrderTemplateByID",
			Handler:    _OrderTemplateService_DeleteOrderTemplateByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orderTemplate.proto",
}
