// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: user.proto

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

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	FindOneLocalUser(ctx context.Context, in *FindOneLocalUserRequest, opts ...grpc.CallOption) (*FindOneLocalUserResponse, error)
	FindOneUserStudent(ctx context.Context, in *FindOneUserStudentRequest, opts ...grpc.CallOption) (*FindOneUserStudentResponse, error)
	Update(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error)
	FindUsersFromList(ctx context.Context, in *FindUsersFromListRequest, opts ...grpc.CallOption) (*FindUsersFromListResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) FindOneLocalUser(ctx context.Context, in *FindOneLocalUserRequest, opts ...grpc.CallOption) (*FindOneLocalUserResponse, error) {
	out := new(FindOneLocalUserResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/FindOneLocalUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) FindOneUserStudent(ctx context.Context, in *FindOneUserStudentRequest, opts ...grpc.CallOption) (*FindOneUserStudentResponse, error) {
	out := new(FindOneUserStudentResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/FindOneUserStudent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Update(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error) {
	out := new(UpdateUserResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) FindUsersFromList(ctx context.Context, in *FindUsersFromListRequest, opts ...grpc.CallOption) (*FindUsersFromListResponse, error) {
	out := new(FindUsersFromListResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/FindUsersFromList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations should embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	FindOneLocalUser(context.Context, *FindOneLocalUserRequest) (*FindOneLocalUserResponse, error)
	FindOneUserStudent(context.Context, *FindOneUserStudentRequest) (*FindOneUserStudentResponse, error)
	Update(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
	FindUsersFromList(context.Context, *FindUsersFromListRequest) (*FindUsersFromListResponse, error)
}

// UnimplementedUserServiceServer should be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) FindOneLocalUser(context.Context, *FindOneLocalUserRequest) (*FindOneLocalUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindOneLocalUser not implemented")
}
func (UnimplementedUserServiceServer) FindOneUserStudent(context.Context, *FindOneUserStudentRequest) (*FindOneUserStudentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindOneUserStudent not implemented")
}
func (UnimplementedUserServiceServer) Update(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedUserServiceServer) FindUsersFromList(context.Context, *FindUsersFromListRequest) (*FindUsersFromListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindUsersFromList not implemented")
}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_FindOneLocalUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindOneLocalUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).FindOneLocalUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/FindOneLocalUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).FindOneLocalUser(ctx, req.(*FindOneLocalUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_FindOneUserStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindOneUserStudentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).FindOneUserStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/FindOneUserStudent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).FindOneUserStudent(ctx, req.(*FindOneUserStudentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Update(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_FindUsersFromList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindUsersFromListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).FindUsersFromList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/FindUsersFromList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).FindUsersFromList(ctx, req.(*FindUsersFromListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindOneLocalUser",
			Handler:    _UserService_FindOneLocalUser_Handler,
		},
		{
			MethodName: "FindOneUserStudent",
			Handler:    _UserService_FindOneUserStudent_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _UserService_Update_Handler,
		},
		{
			MethodName: "FindUsersFromList",
			Handler:    _UserService_FindUsersFromList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
