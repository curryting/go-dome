// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: user.proto

package __

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

// UserCenterClient is the client API for UserCenter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserCenterClient interface {
	FindUser(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*UserRes, error)
	AddUser(ctx context.Context, in *AddUserReq, opts ...grpc.CallOption) (*CommonRes, error)
	UpdateUser(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*CommonRes, error)
}

type userCenterClient struct {
	cc grpc.ClientConnInterface
}

func NewUserCenterClient(cc grpc.ClientConnInterface) UserCenterClient {
	return &userCenterClient{cc}
}

func (c *userCenterClient) FindUser(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*UserRes, error) {
	out := new(UserRes)
	err := c.cc.Invoke(ctx, "/pb.UserCenter/FindUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCenterClient) AddUser(ctx context.Context, in *AddUserReq, opts ...grpc.CallOption) (*CommonRes, error) {
	out := new(CommonRes)
	err := c.cc.Invoke(ctx, "/pb.UserCenter/AddUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCenterClient) UpdateUser(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*CommonRes, error) {
	out := new(CommonRes)
	err := c.cc.Invoke(ctx, "/pb.UserCenter/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserCenterServer is the server API for UserCenter service.
// All implementations must embed UnimplementedUserCenterServer
// for forward compatibility
type UserCenterServer interface {
	FindUser(context.Context, *UserReq) (*UserRes, error)
	AddUser(context.Context, *AddUserReq) (*CommonRes, error)
	UpdateUser(context.Context, *UserReq) (*CommonRes, error)
	mustEmbedUnimplementedUserCenterServer()
}

// UnimplementedUserCenterServer must be embedded to have forward compatible implementations.
type UnimplementedUserCenterServer struct {
}

func (UnimplementedUserCenterServer) FindUser(context.Context, *UserReq) (*UserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindUser not implemented")
}
func (UnimplementedUserCenterServer) AddUser(context.Context, *AddUserReq) (*CommonRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUser not implemented")
}
func (UnimplementedUserCenterServer) UpdateUser(context.Context, *UserReq) (*CommonRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedUserCenterServer) mustEmbedUnimplementedUserCenterServer() {}

// UnsafeUserCenterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserCenterServer will
// result in compilation errors.
type UnsafeUserCenterServer interface {
	mustEmbedUnimplementedUserCenterServer()
}

func RegisterUserCenterServer(s grpc.ServiceRegistrar, srv UserCenterServer) {
	s.RegisterService(&UserCenter_ServiceDesc, srv)
}

func _UserCenter_FindUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCenterServer).FindUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserCenter/FindUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCenterServer).FindUser(ctx, req.(*UserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCenter_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCenterServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserCenter/AddUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCenterServer).AddUser(ctx, req.(*AddUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCenter_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCenterServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserCenter/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCenterServer).UpdateUser(ctx, req.(*UserReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserCenter_ServiceDesc is the grpc.ServiceDesc for UserCenter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserCenter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UserCenter",
	HandlerType: (*UserCenterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindUser",
			Handler:    _UserCenter_FindUser_Handler,
		},
		{
			MethodName: "AddUser",
			Handler:    _UserCenter_AddUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserCenter_UpdateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
