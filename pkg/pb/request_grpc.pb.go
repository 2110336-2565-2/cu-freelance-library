// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: request.proto

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

// RequestServiceClient is the client API for RequestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RequestServiceClient interface {
	GetOneRequest(ctx context.Context, in *GetOneRequestInput, opts ...grpc.CallOption) (*GetOneRequestResponse, error)
	GetRequestsByUser(ctx context.Context, in *GetRequestsByUserInput, opts ...grpc.CallOption) (*GetRequestsByUserResponse, error)
	GetRequestsByStatus(ctx context.Context, in *GetRequestsByStatusInput, opts ...grpc.CallOption) (*GetRequestsByStatusResponse, error)
	GetRequestsByOrderTemplate(ctx context.Context, in *GetRequestsByOrderTemplateInput, opts ...grpc.CallOption) (*GetRequestsByOrderTemplateResponse, error)
	CreateRequest(ctx context.Context, in *CreateRequestInput, opts ...grpc.CallOption) (*CreateRequestResponse, error)
	UpdateRequest(ctx context.Context, in *UpdateRequestInput, opts ...grpc.CallOption) (*UpdateRequestResponse, error)
	DeleteRequest(ctx context.Context, in *DeleteRequestInput, opts ...grpc.CallOption) (*DeleteRequestResponse, error)
}

type requestServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRequestServiceClient(cc grpc.ClientConnInterface) RequestServiceClient {
	return &requestServiceClient{cc}
}

func (c *requestServiceClient) GetOneRequest(ctx context.Context, in *GetOneRequestInput, opts ...grpc.CallOption) (*GetOneRequestResponse, error) {
	out := new(GetOneRequestResponse)
	err := c.cc.Invoke(ctx, "/request.RequestService/GetOneRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requestServiceClient) GetRequestsByUser(ctx context.Context, in *GetRequestsByUserInput, opts ...grpc.CallOption) (*GetRequestsByUserResponse, error) {
	out := new(GetRequestsByUserResponse)
	err := c.cc.Invoke(ctx, "/request.RequestService/GetRequestsByUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requestServiceClient) GetRequestsByStatus(ctx context.Context, in *GetRequestsByStatusInput, opts ...grpc.CallOption) (*GetRequestsByStatusResponse, error) {
	out := new(GetRequestsByStatusResponse)
	err := c.cc.Invoke(ctx, "/request.RequestService/GetRequestsByStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requestServiceClient) GetRequestsByOrderTemplate(ctx context.Context, in *GetRequestsByOrderTemplateInput, opts ...grpc.CallOption) (*GetRequestsByOrderTemplateResponse, error) {
	out := new(GetRequestsByOrderTemplateResponse)
	err := c.cc.Invoke(ctx, "/request.RequestService/GetRequestsByOrderTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requestServiceClient) CreateRequest(ctx context.Context, in *CreateRequestInput, opts ...grpc.CallOption) (*CreateRequestResponse, error) {
	out := new(CreateRequestResponse)
	err := c.cc.Invoke(ctx, "/request.RequestService/CreateRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requestServiceClient) UpdateRequest(ctx context.Context, in *UpdateRequestInput, opts ...grpc.CallOption) (*UpdateRequestResponse, error) {
	out := new(UpdateRequestResponse)
	err := c.cc.Invoke(ctx, "/request.RequestService/UpdateRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requestServiceClient) DeleteRequest(ctx context.Context, in *DeleteRequestInput, opts ...grpc.CallOption) (*DeleteRequestResponse, error) {
	out := new(DeleteRequestResponse)
	err := c.cc.Invoke(ctx, "/request.RequestService/DeleteRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RequestServiceServer is the server API for RequestService service.
// All implementations should embed UnimplementedRequestServiceServer
// for forward compatibility
type RequestServiceServer interface {
	GetOneRequest(context.Context, *GetOneRequestInput) (*GetOneRequestResponse, error)
	GetRequestsByUser(context.Context, *GetRequestsByUserInput) (*GetRequestsByUserResponse, error)
	GetRequestsByStatus(context.Context, *GetRequestsByStatusInput) (*GetRequestsByStatusResponse, error)
	GetRequestsByOrderTemplate(context.Context, *GetRequestsByOrderTemplateInput) (*GetRequestsByOrderTemplateResponse, error)
	CreateRequest(context.Context, *CreateRequestInput) (*CreateRequestResponse, error)
	UpdateRequest(context.Context, *UpdateRequestInput) (*UpdateRequestResponse, error)
	DeleteRequest(context.Context, *DeleteRequestInput) (*DeleteRequestResponse, error)
}

// UnimplementedRequestServiceServer should be embedded to have forward compatible implementations.
type UnimplementedRequestServiceServer struct {
}

func (UnimplementedRequestServiceServer) GetOneRequest(context.Context, *GetOneRequestInput) (*GetOneRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOneRequest not implemented")
}
func (UnimplementedRequestServiceServer) GetRequestsByUser(context.Context, *GetRequestsByUserInput) (*GetRequestsByUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRequestsByUser not implemented")
}
func (UnimplementedRequestServiceServer) GetRequestsByStatus(context.Context, *GetRequestsByStatusInput) (*GetRequestsByStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRequestsByStatus not implemented")
}
func (UnimplementedRequestServiceServer) GetRequestsByOrderTemplate(context.Context, *GetRequestsByOrderTemplateInput) (*GetRequestsByOrderTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRequestsByOrderTemplate not implemented")
}
func (UnimplementedRequestServiceServer) CreateRequest(context.Context, *CreateRequestInput) (*CreateRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRequest not implemented")
}
func (UnimplementedRequestServiceServer) UpdateRequest(context.Context, *UpdateRequestInput) (*UpdateRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRequest not implemented")
}
func (UnimplementedRequestServiceServer) DeleteRequest(context.Context, *DeleteRequestInput) (*DeleteRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRequest not implemented")
}

// UnsafeRequestServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RequestServiceServer will
// result in compilation errors.
type UnsafeRequestServiceServer interface {
	mustEmbedUnimplementedRequestServiceServer()
}

func RegisterRequestServiceServer(s grpc.ServiceRegistrar, srv RequestServiceServer) {
	s.RegisterService(&RequestService_ServiceDesc, srv)
}

func _RequestService_GetOneRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOneRequestInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestServiceServer).GetOneRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/request.RequestService/GetOneRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestServiceServer).GetOneRequest(ctx, req.(*GetOneRequestInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _RequestService_GetRequestsByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequestsByUserInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestServiceServer).GetRequestsByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/request.RequestService/GetRequestsByUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestServiceServer).GetRequestsByUser(ctx, req.(*GetRequestsByUserInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _RequestService_GetRequestsByStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequestsByStatusInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestServiceServer).GetRequestsByStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/request.RequestService/GetRequestsByStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestServiceServer).GetRequestsByStatus(ctx, req.(*GetRequestsByStatusInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _RequestService_GetRequestsByOrderTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequestsByOrderTemplateInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestServiceServer).GetRequestsByOrderTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/request.RequestService/GetRequestsByOrderTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestServiceServer).GetRequestsByOrderTemplate(ctx, req.(*GetRequestsByOrderTemplateInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _RequestService_CreateRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequestInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestServiceServer).CreateRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/request.RequestService/CreateRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestServiceServer).CreateRequest(ctx, req.(*CreateRequestInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _RequestService_UpdateRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequestInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestServiceServer).UpdateRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/request.RequestService/UpdateRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestServiceServer).UpdateRequest(ctx, req.(*UpdateRequestInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _RequestService_DeleteRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequestInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestServiceServer).DeleteRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/request.RequestService/DeleteRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestServiceServer).DeleteRequest(ctx, req.(*DeleteRequestInput))
	}
	return interceptor(ctx, in, info, handler)
}

// RequestService_ServiceDesc is the grpc.ServiceDesc for RequestService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RequestService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "request.RequestService",
	HandlerType: (*RequestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOneRequest",
			Handler:    _RequestService_GetOneRequest_Handler,
		},
		{
			MethodName: "GetRequestsByUser",
			Handler:    _RequestService_GetRequestsByUser_Handler,
		},
		{
			MethodName: "GetRequestsByStatus",
			Handler:    _RequestService_GetRequestsByStatus_Handler,
		},
		{
			MethodName: "GetRequestsByOrderTemplate",
			Handler:    _RequestService_GetRequestsByOrderTemplate_Handler,
		},
		{
			MethodName: "CreateRequest",
			Handler:    _RequestService_CreateRequest_Handler,
		},
		{
			MethodName: "UpdateRequest",
			Handler:    _RequestService_UpdateRequest_Handler,
		},
		{
			MethodName: "DeleteRequest",
			Handler:    _RequestService_DeleteRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "request.proto",
}