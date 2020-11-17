// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	Get(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*UserResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*MessageResponse, error)
	List(ctx context.Context, in *UserNilRequest, opts ...grpc.CallOption) (*UsersResponse, error)
	Delete(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*MessageResponse, error)
	Deletes(ctx context.Context, in *IdsRequest, opts ...grpc.CallOption) (*MessageResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*MessageResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Ban(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*MessageResponse, error)
	Unban(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*MessageResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Get(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*MessageResponse, error) {
	out := new(MessageResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) List(ctx context.Context, in *UserNilRequest, opts ...grpc.CallOption) (*UsersResponse, error) {
	out := new(UsersResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Delete(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*MessageResponse, error) {
	out := new(MessageResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Deletes(ctx context.Context, in *IdsRequest, opts ...grpc.CallOption) (*MessageResponse, error) {
	out := new(MessageResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/Deletes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*MessageResponse, error) {
	out := new(MessageResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Ban(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*MessageResponse, error) {
	out := new(MessageResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/Ban", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Unban(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*MessageResponse, error) {
	out := new(MessageResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/Unban", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	Get(context.Context, *IdRequest) (*UserResponse, error)
	Create(context.Context, *CreateRequest) (*MessageResponse, error)
	List(context.Context, *UserNilRequest) (*UsersResponse, error)
	Delete(context.Context, *IdRequest) (*MessageResponse, error)
	Deletes(context.Context, *IdsRequest) (*MessageResponse, error)
	Update(context.Context, *UpdateRequest) (*MessageResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Ban(context.Context, *IdRequest) (*MessageResponse, error)
	Unban(context.Context, *IdRequest) (*MessageResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) Get(context.Context, *IdRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedUserServiceServer) Create(context.Context, *CreateRequest) (*MessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedUserServiceServer) List(context.Context, *UserNilRequest) (*UsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedUserServiceServer) Delete(context.Context, *IdRequest) (*MessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedUserServiceServer) Deletes(context.Context, *IdsRequest) (*MessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deletes not implemented")
}
func (UnimplementedUserServiceServer) Update(context.Context, *UpdateRequest) (*MessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedUserServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserServiceServer) Ban(context.Context, *IdRequest) (*MessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ban not implemented")
}
func (UnimplementedUserServiceServer) Unban(context.Context, *IdRequest) (*MessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unban not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Get(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserNilRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).List(ctx, req.(*UserNilRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Delete(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Deletes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Deletes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/Deletes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Deletes(ctx, req.(*IdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
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
		return srv.(UserServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Ban_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Ban(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/Ban",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Ban(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Unban_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Unban(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/Unban",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Unban(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _UserService_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _UserService_Create_Handler,
		},
		{
			MethodName: "List",
			Handler:    _UserService_List_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _UserService_Delete_Handler,
		},
		{
			MethodName: "Deletes",
			Handler:    _UserService_Deletes_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _UserService_Update_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _UserService_Login_Handler,
		},
		{
			MethodName: "Ban",
			Handler:    _UserService_Ban_Handler,
		},
		{
			MethodName: "Unban",
			Handler:    _UserService_Unban_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/infrastructure/delivery/grpc/proto/user.proto",
}