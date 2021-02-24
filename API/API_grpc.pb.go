// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

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

// NetworkServiceClient is the client API for NetworkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NetworkServiceClient interface {
	// getNetworkMembers returns the members in a network
	GetNetworkMembers(ctx context.Context, in *NetworkKey, opts ...grpc.CallOption) (*UserKeys, error)
}

type networkServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNetworkServiceClient(cc grpc.ClientConnInterface) NetworkServiceClient {
	return &networkServiceClient{cc}
}

func (c *networkServiceClient) GetNetworkMembers(ctx context.Context, in *NetworkKey, opts ...grpc.CallOption) (*UserKeys, error) {
	out := new(UserKeys)
	err := c.cc.Invoke(ctx, "/interview.NetworkService/getNetworkMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetworkServiceServer is the server API for NetworkService service.
// All implementations must embed UnimplementedNetworkServiceServer
// for forward compatibility
type NetworkServiceServer interface {
	// getNetworkMembers returns the members in a network
	GetNetworkMembers(context.Context, *NetworkKey) (*UserKeys, error)
	mustEmbedUnimplementedNetworkServiceServer()
}

// UnimplementedNetworkServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNetworkServiceServer struct {
}

func (UnimplementedNetworkServiceServer) GetNetworkMembers(context.Context, *NetworkKey) (*UserKeys, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNetworkMembers not implemented")
}
func (UnimplementedNetworkServiceServer) mustEmbedUnimplementedNetworkServiceServer() {}

// UnsafeNetworkServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NetworkServiceServer will
// result in compilation errors.
type UnsafeNetworkServiceServer interface {
	mustEmbedUnimplementedNetworkServiceServer()
}

func RegisterNetworkServiceServer(s grpc.ServiceRegistrar, srv NetworkServiceServer) {
	s.RegisterService(&NetworkService_ServiceDesc, srv)
}

func _NetworkService_GetNetworkMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).GetNetworkMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interview.NetworkService/getNetworkMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).GetNetworkMembers(ctx, req.(*NetworkKey))
	}
	return interceptor(ctx, in, info, handler)
}

// NetworkService_ServiceDesc is the grpc.ServiceDesc for NetworkService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NetworkService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "interview.NetworkService",
	HandlerType: (*NetworkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getNetworkMembers",
			Handler:    _NetworkService_GetNetworkMembers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "API/API.proto",
}

// ContactServiceClient is the client API for ContactService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContactServiceClient interface {
	// getCommonContacts returns the contacts two users have in common
	GetCommonContacts(ctx context.Context, in *TwoUserKeys, opts ...grpc.CallOption) (*Contacts, error)
}

type contactServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewContactServiceClient(cc grpc.ClientConnInterface) ContactServiceClient {
	return &contactServiceClient{cc}
}

func (c *contactServiceClient) GetCommonContacts(ctx context.Context, in *TwoUserKeys, opts ...grpc.CallOption) (*Contacts, error) {
	out := new(Contacts)
	err := c.cc.Invoke(ctx, "/interview.ContactService/getCommonContacts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContactServiceServer is the server API for ContactService service.
// All implementations must embed UnimplementedContactServiceServer
// for forward compatibility
type ContactServiceServer interface {
	// getCommonContacts returns the contacts two users have in common
	GetCommonContacts(context.Context, *TwoUserKeys) (*Contacts, error)
	mustEmbedUnimplementedContactServiceServer()
}

// UnimplementedContactServiceServer must be embedded to have forward compatible implementations.
type UnimplementedContactServiceServer struct {
}

func (UnimplementedContactServiceServer) GetCommonContacts(context.Context, *TwoUserKeys) (*Contacts, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommonContacts not implemented")
}
func (UnimplementedContactServiceServer) mustEmbedUnimplementedContactServiceServer() {}

// UnsafeContactServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContactServiceServer will
// result in compilation errors.
type UnsafeContactServiceServer interface {
	mustEmbedUnimplementedContactServiceServer()
}

func RegisterContactServiceServer(s grpc.ServiceRegistrar, srv ContactServiceServer) {
	s.RegisterService(&ContactService_ServiceDesc, srv)
}

func _ContactService_GetCommonContacts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TwoUserKeys)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServiceServer).GetCommonContacts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interview.ContactService/getCommonContacts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServiceServer).GetCommonContacts(ctx, req.(*TwoUserKeys))
	}
	return interceptor(ctx, in, info, handler)
}

// ContactService_ServiceDesc is the grpc.ServiceDesc for ContactService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ContactService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "interview.ContactService",
	HandlerType: (*ContactServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getCommonContacts",
			Handler:    _ContactService_GetCommonContacts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "API/API.proto",
}

// InterestsServiceClient is the client API for InterestsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InterestsServiceClient interface {
	// getSharedInterests returns the topics both users share interest in
	GetSharedInterests(ctx context.Context, in *TwoUserKeys, opts ...grpc.CallOption) (*Interests, error)
}

type interestsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInterestsServiceClient(cc grpc.ClientConnInterface) InterestsServiceClient {
	return &interestsServiceClient{cc}
}

func (c *interestsServiceClient) GetSharedInterests(ctx context.Context, in *TwoUserKeys, opts ...grpc.CallOption) (*Interests, error) {
	out := new(Interests)
	err := c.cc.Invoke(ctx, "/interview.InterestsService/getSharedInterests", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InterestsServiceServer is the server API for InterestsService service.
// All implementations must embed UnimplementedInterestsServiceServer
// for forward compatibility
type InterestsServiceServer interface {
	// getSharedInterests returns the topics both users share interest in
	GetSharedInterests(context.Context, *TwoUserKeys) (*Interests, error)
	mustEmbedUnimplementedInterestsServiceServer()
}

// UnimplementedInterestsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedInterestsServiceServer struct {
}

func (UnimplementedInterestsServiceServer) GetSharedInterests(context.Context, *TwoUserKeys) (*Interests, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSharedInterests not implemented")
}
func (UnimplementedInterestsServiceServer) mustEmbedUnimplementedInterestsServiceServer() {}

// UnsafeInterestsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InterestsServiceServer will
// result in compilation errors.
type UnsafeInterestsServiceServer interface {
	mustEmbedUnimplementedInterestsServiceServer()
}

func RegisterInterestsServiceServer(s grpc.ServiceRegistrar, srv InterestsServiceServer) {
	s.RegisterService(&InterestsService_ServiceDesc, srv)
}

func _InterestsService_GetSharedInterests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TwoUserKeys)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InterestsServiceServer).GetSharedInterests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interview.InterestsService/getSharedInterests",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InterestsServiceServer).GetSharedInterests(ctx, req.(*TwoUserKeys))
	}
	return interceptor(ctx, in, info, handler)
}

// InterestsService_ServiceDesc is the grpc.ServiceDesc for InterestsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InterestsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "interview.InterestsService",
	HandlerType: (*InterestsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getSharedInterests",
			Handler:    _InterestsService_GetSharedInterests_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "API/API.proto",
}

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	// getUser returns the User associated with a key
	GetUser(ctx context.Context, in *UserKey, opts ...grpc.CallOption) (*User, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetUser(ctx context.Context, in *UserKey, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/interview.UserService/getUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	// getUser returns the User associated with a key
	GetUser(context.Context, *UserKey) (*User, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) GetUser(context.Context, *UserKey) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interview.UserService/getUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUser(ctx, req.(*UserKey))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "interview.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getUser",
			Handler:    _UserService_GetUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "API/API.proto",
}

// ViewNetworkServiceClient is the client API for ViewNetworkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ViewNetworkServiceClient interface {
	// viewNetworkMembers presents a list of network members for display, be enriching the network member list with additional information
	ViewNetworkMembers(ctx context.Context, in *UserViewingNetwork, opts ...grpc.CallOption) (*NetworkMembersView, error)
}

type viewNetworkServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewViewNetworkServiceClient(cc grpc.ClientConnInterface) ViewNetworkServiceClient {
	return &viewNetworkServiceClient{cc}
}

func (c *viewNetworkServiceClient) ViewNetworkMembers(ctx context.Context, in *UserViewingNetwork, opts ...grpc.CallOption) (*NetworkMembersView, error) {
	out := new(NetworkMembersView)
	err := c.cc.Invoke(ctx, "/interview.ViewNetworkService/viewNetworkMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ViewNetworkServiceServer is the server API for ViewNetworkService service.
// All implementations must embed UnimplementedViewNetworkServiceServer
// for forward compatibility
type ViewNetworkServiceServer interface {
	// viewNetworkMembers presents a list of network members for display, be enriching the network member list with additional information
	ViewNetworkMembers(context.Context, *UserViewingNetwork) (*NetworkMembersView, error)
	mustEmbedUnimplementedViewNetworkServiceServer()
}

// UnimplementedViewNetworkServiceServer must be embedded to have forward compatible implementations.
type UnimplementedViewNetworkServiceServer struct {
}

func (UnimplementedViewNetworkServiceServer) ViewNetworkMembers(context.Context, *UserViewingNetwork) (*NetworkMembersView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewNetworkMembers not implemented")
}
func (UnimplementedViewNetworkServiceServer) mustEmbedUnimplementedViewNetworkServiceServer() {}

// UnsafeViewNetworkServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ViewNetworkServiceServer will
// result in compilation errors.
type UnsafeViewNetworkServiceServer interface {
	mustEmbedUnimplementedViewNetworkServiceServer()
}

func RegisterViewNetworkServiceServer(s grpc.ServiceRegistrar, srv ViewNetworkServiceServer) {
	s.RegisterService(&ViewNetworkService_ServiceDesc, srv)
}

func _ViewNetworkService_ViewNetworkMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserViewingNetwork)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViewNetworkServiceServer).ViewNetworkMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interview.ViewNetworkService/viewNetworkMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViewNetworkServiceServer).ViewNetworkMembers(ctx, req.(*UserViewingNetwork))
	}
	return interceptor(ctx, in, info, handler)
}

// ViewNetworkService_ServiceDesc is the grpc.ServiceDesc for ViewNetworkService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ViewNetworkService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "interview.ViewNetworkService",
	HandlerType: (*ViewNetworkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "viewNetworkMembers",
			Handler:    _ViewNetworkService_ViewNetworkMembers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "API/API.proto",
}