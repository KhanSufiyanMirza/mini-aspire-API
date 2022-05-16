// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: service_mini_aspire.proto

package ma

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

// MiniAspireClient is the client API for MiniAspire service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiniAspireClient interface {
	CreateUser(ctx context.Context, in *CreatUserRequest, opts ...grpc.CallOption) (*CreatUserResponse, error)
	LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error)
}

type miniAspireClient struct {
	cc grpc.ClientConnInterface
}

func NewMiniAspireClient(cc grpc.ClientConnInterface) MiniAspireClient {
	return &miniAspireClient{cc}
}

func (c *miniAspireClient) CreateUser(ctx context.Context, in *CreatUserRequest, opts ...grpc.CallOption) (*CreatUserResponse, error) {
	out := new(CreatUserResponse)
	err := c.cc.Invoke(ctx, "/ma.MiniAspire/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *miniAspireClient) LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error) {
	out := new(LoginUserResponse)
	err := c.cc.Invoke(ctx, "/ma.MiniAspire/LoginUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiniAspireServer is the server API for MiniAspire service.
// All implementations must embed UnimplementedMiniAspireServer
// for forward compatibility
type MiniAspireServer interface {
	CreateUser(context.Context, *CreatUserRequest) (*CreatUserResponse, error)
	LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error)
	mustEmbedUnimplementedMiniAspireServer()
}

// UnimplementedMiniAspireServer must be embedded to have forward compatible implementations.
type UnimplementedMiniAspireServer struct {
}

func (UnimplementedMiniAspireServer) CreateUser(context.Context, *CreatUserRequest) (*CreatUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedMiniAspireServer) LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (UnimplementedMiniAspireServer) mustEmbedUnimplementedMiniAspireServer() {}

// UnsafeMiniAspireServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MiniAspireServer will
// result in compilation errors.
type UnsafeMiniAspireServer interface {
	mustEmbedUnimplementedMiniAspireServer()
}

func RegisterMiniAspireServer(s grpc.ServiceRegistrar, srv MiniAspireServer) {
	s.RegisterService(&MiniAspire_ServiceDesc, srv)
}

func _MiniAspire_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiniAspireServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ma.MiniAspire/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiniAspireServer).CreateUser(ctx, req.(*CreatUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MiniAspire_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiniAspireServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ma.MiniAspire/LoginUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiniAspireServer).LoginUser(ctx, req.(*LoginUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MiniAspire_ServiceDesc is the grpc.ServiceDesc for MiniAspire service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MiniAspire_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ma.MiniAspire",
	HandlerType: (*MiniAspireServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _MiniAspire_CreateUser_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _MiniAspire_LoginUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_mini_aspire.proto",
}