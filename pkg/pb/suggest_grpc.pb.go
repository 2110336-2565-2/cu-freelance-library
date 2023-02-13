// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: suggest.proto

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

// SuggestServiceClient is the client API for SuggestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SuggestServiceClient interface {
	Suggest(ctx context.Context, in *SuggestRequest, opts ...grpc.CallOption) (*SuggestResponse, error)
}

type suggestServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSuggestServiceClient(cc grpc.ClientConnInterface) SuggestServiceClient {
	return &suggestServiceClient{cc}
}

func (c *suggestServiceClient) Suggest(ctx context.Context, in *SuggestRequest, opts ...grpc.CallOption) (*SuggestResponse, error) {
	out := new(SuggestResponse)
	err := c.cc.Invoke(ctx, "/course_suggest.SuggestService/Suggest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SuggestServiceServer is the server API for SuggestService service.
// All implementations should embed UnimplementedSuggestServiceServer
// for forward compatibility
type SuggestServiceServer interface {
	Suggest(context.Context, *SuggestRequest) (*SuggestResponse, error)
}

// UnimplementedSuggestServiceServer should be embedded to have forward compatible implementations.
type UnimplementedSuggestServiceServer struct {
}

func (UnimplementedSuggestServiceServer) Suggest(context.Context, *SuggestRequest) (*SuggestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Suggest not implemented")
}

// UnsafeSuggestServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SuggestServiceServer will
// result in compilation errors.
type UnsafeSuggestServiceServer interface {
	mustEmbedUnimplementedSuggestServiceServer()
}

func RegisterSuggestServiceServer(s grpc.ServiceRegistrar, srv SuggestServiceServer) {
	s.RegisterService(&SuggestService_ServiceDesc, srv)
}

func _SuggestService_Suggest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SuggestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SuggestServiceServer).Suggest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/course_suggest.SuggestService/Suggest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SuggestServiceServer).Suggest(ctx, req.(*SuggestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SuggestService_ServiceDesc is the grpc.ServiceDesc for SuggestService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SuggestService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "course_suggest.SuggestService",
	HandlerType: (*SuggestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Suggest",
			Handler:    _SuggestService_Suggest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "suggest.proto",
}
