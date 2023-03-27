// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: file.proto

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

// FileServiceClient is the client API for FileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileServiceClient interface {
	UploadAvatarImage(ctx context.Context, in *UploadAvatarImageRequest, opts ...grpc.CallOption) (*UploadAvatarImageResponse, error)
	GetUserAvatar(ctx context.Context, in *GetUserAvatarRequest, opts ...grpc.CallOption) (*UserAvatarResponse, error)
	DeleteAvater(ctx context.Context, in *DeleteAvaterRequest, opts ...grpc.CallOption) (*DeleteAvaterResponse, error)
	UploadPortfolioImage(ctx context.Context, in *UploadPortfolioImageRequest, opts ...grpc.CallOption) (*UploadPortfolioImageResponse, error)
	GetPortThumbnail(ctx context.Context, in *GetPortThumbnailRequest, opts ...grpc.CallOption) (*PortThumbnailResponse, error)
	GetAllPortImage(ctx context.Context, in *GetAllPortImageRequest, opts ...grpc.CallOption) (*AllPortImageResponse, error)
	DeletePortfolioImage(ctx context.Context, in *DeletePortfolioImageRequest, opts ...grpc.CallOption) (*DeletePortfolioImageResponse, error)
	SelectThumbnail(ctx context.Context, in *SelectThumbnailRequest, opts ...grpc.CallOption) (*SelectThumbnailResponse, error)
}

type fileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFileServiceClient(cc grpc.ClientConnInterface) FileServiceClient {
	return &fileServiceClient{cc}
}

func (c *fileServiceClient) UploadAvatarImage(ctx context.Context, in *UploadAvatarImageRequest, opts ...grpc.CallOption) (*UploadAvatarImageResponse, error) {
	out := new(UploadAvatarImageResponse)
	err := c.cc.Invoke(ctx, "/file.FileService/UploadAvatarImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) GetUserAvatar(ctx context.Context, in *GetUserAvatarRequest, opts ...grpc.CallOption) (*UserAvatarResponse, error) {
	out := new(UserAvatarResponse)
	err := c.cc.Invoke(ctx, "/file.FileService/GetUserAvatar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) DeleteAvater(ctx context.Context, in *DeleteAvaterRequest, opts ...grpc.CallOption) (*DeleteAvaterResponse, error) {
	out := new(DeleteAvaterResponse)
	err := c.cc.Invoke(ctx, "/file.FileService/DeleteAvater", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) UploadPortfolioImage(ctx context.Context, in *UploadPortfolioImageRequest, opts ...grpc.CallOption) (*UploadPortfolioImageResponse, error) {
	out := new(UploadPortfolioImageResponse)
	err := c.cc.Invoke(ctx, "/file.FileService/UploadPortfolioImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) GetPortThumbnail(ctx context.Context, in *GetPortThumbnailRequest, opts ...grpc.CallOption) (*PortThumbnailResponse, error) {
	out := new(PortThumbnailResponse)
	err := c.cc.Invoke(ctx, "/file.FileService/GetPortThumbnail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) GetAllPortImage(ctx context.Context, in *GetAllPortImageRequest, opts ...grpc.CallOption) (*AllPortImageResponse, error) {
	out := new(AllPortImageResponse)
	err := c.cc.Invoke(ctx, "/file.FileService/GetAllPortImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) DeletePortfolioImage(ctx context.Context, in *DeletePortfolioImageRequest, opts ...grpc.CallOption) (*DeletePortfolioImageResponse, error) {
	out := new(DeletePortfolioImageResponse)
	err := c.cc.Invoke(ctx, "/file.FileService/DeletePortfolioImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) SelectThumbnail(ctx context.Context, in *SelectThumbnailRequest, opts ...grpc.CallOption) (*SelectThumbnailResponse, error) {
	out := new(SelectThumbnailResponse)
	err := c.cc.Invoke(ctx, "/file.FileService/SelectThumbnail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileServiceServer is the server API for FileService service.
// All implementations should embed UnimplementedFileServiceServer
// for forward compatibility
type FileServiceServer interface {
	UploadAvatarImage(context.Context, *UploadAvatarImageRequest) (*UploadAvatarImageResponse, error)
	GetUserAvatar(context.Context, *GetUserAvatarRequest) (*UserAvatarResponse, error)
	DeleteAvater(context.Context, *DeleteAvaterRequest) (*DeleteAvaterResponse, error)
	UploadPortfolioImage(context.Context, *UploadPortfolioImageRequest) (*UploadPortfolioImageResponse, error)
	GetPortThumbnail(context.Context, *GetPortThumbnailRequest) (*PortThumbnailResponse, error)
	GetAllPortImage(context.Context, *GetAllPortImageRequest) (*AllPortImageResponse, error)
	DeletePortfolioImage(context.Context, *DeletePortfolioImageRequest) (*DeletePortfolioImageResponse, error)
	SelectThumbnail(context.Context, *SelectThumbnailRequest) (*SelectThumbnailResponse, error)
}

// UnimplementedFileServiceServer should be embedded to have forward compatible implementations.
type UnimplementedFileServiceServer struct {
}

func (UnimplementedFileServiceServer) UploadAvatarImage(context.Context, *UploadAvatarImageRequest) (*UploadAvatarImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadAvatarImage not implemented")
}
func (UnimplementedFileServiceServer) GetUserAvatar(context.Context, *GetUserAvatarRequest) (*UserAvatarResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserAvatar not implemented")
}
func (UnimplementedFileServiceServer) DeleteAvater(context.Context, *DeleteAvaterRequest) (*DeleteAvaterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAvater not implemented")
}
func (UnimplementedFileServiceServer) UploadPortfolioImage(context.Context, *UploadPortfolioImageRequest) (*UploadPortfolioImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadPortfolioImage not implemented")
}
func (UnimplementedFileServiceServer) GetPortThumbnail(context.Context, *GetPortThumbnailRequest) (*PortThumbnailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPortThumbnail not implemented")
}
func (UnimplementedFileServiceServer) GetAllPortImage(context.Context, *GetAllPortImageRequest) (*AllPortImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllPortImage not implemented")
}
func (UnimplementedFileServiceServer) DeletePortfolioImage(context.Context, *DeletePortfolioImageRequest) (*DeletePortfolioImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePortfolioImage not implemented")
}
func (UnimplementedFileServiceServer) SelectThumbnail(context.Context, *SelectThumbnailRequest) (*SelectThumbnailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SelectThumbnail not implemented")
}

// UnsafeFileServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileServiceServer will
// result in compilation errors.
type UnsafeFileServiceServer interface {
	mustEmbedUnimplementedFileServiceServer()
}

func RegisterFileServiceServer(s grpc.ServiceRegistrar, srv FileServiceServer) {
	s.RegisterService(&FileService_ServiceDesc, srv)
}

func _FileService_UploadAvatarImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadAvatarImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).UploadAvatarImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/file.FileService/UploadAvatarImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).UploadAvatarImage(ctx, req.(*UploadAvatarImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_GetUserAvatar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserAvatarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).GetUserAvatar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/file.FileService/GetUserAvatar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).GetUserAvatar(ctx, req.(*GetUserAvatarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_DeleteAvater_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAvaterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).DeleteAvater(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/file.FileService/DeleteAvater",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).DeleteAvater(ctx, req.(*DeleteAvaterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_UploadPortfolioImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadPortfolioImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).UploadPortfolioImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/file.FileService/UploadPortfolioImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).UploadPortfolioImage(ctx, req.(*UploadPortfolioImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_GetPortThumbnail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPortThumbnailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).GetPortThumbnail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/file.FileService/GetPortThumbnail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).GetPortThumbnail(ctx, req.(*GetPortThumbnailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_GetAllPortImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllPortImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).GetAllPortImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/file.FileService/GetAllPortImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).GetAllPortImage(ctx, req.(*GetAllPortImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_DeletePortfolioImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePortfolioImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).DeletePortfolioImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/file.FileService/DeletePortfolioImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).DeletePortfolioImage(ctx, req.(*DeletePortfolioImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_SelectThumbnail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectThumbnailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).SelectThumbnail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/file.FileService/SelectThumbnail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).SelectThumbnail(ctx, req.(*SelectThumbnailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FileService_ServiceDesc is the grpc.ServiceDesc for FileService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "file.FileService",
	HandlerType: (*FileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadAvatarImage",
			Handler:    _FileService_UploadAvatarImage_Handler,
		},
		{
			MethodName: "GetUserAvatar",
			Handler:    _FileService_GetUserAvatar_Handler,
		},
		{
			MethodName: "DeleteAvater",
			Handler:    _FileService_DeleteAvater_Handler,
		},
		{
			MethodName: "UploadPortfolioImage",
			Handler:    _FileService_UploadPortfolioImage_Handler,
		},
		{
			MethodName: "GetPortThumbnail",
			Handler:    _FileService_GetPortThumbnail_Handler,
		},
		{
			MethodName: "GetAllPortImage",
			Handler:    _FileService_GetAllPortImage_Handler,
		},
		{
			MethodName: "DeletePortfolioImage",
			Handler:    _FileService_DeletePortfolioImage_Handler,
		},
		{
			MethodName: "SelectThumbnail",
			Handler:    _FileService_SelectThumbnail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "file.proto",
}
