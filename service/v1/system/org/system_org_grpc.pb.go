// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: system_org.proto

package org

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OrgServiceClient is the client API for OrgService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrgServiceClient interface {
	// query
	GetOrgById(ctx context.Context, in *GetOrgByIdQuery, opts ...grpc.CallOption) (*GetOrgByIdReply, error)
	GetOrgVisibilityById(ctx context.Context, in *GetOrgVisibilityByIdQuery, opts ...grpc.CallOption) (*GetOrgVisibilityByIdReply, error)
	GetOrgList(ctx context.Context, in *GetOrgListQuery, opts ...grpc.CallOption) (*GetOrgListReply, error)
	GetWorkspaceList(ctx context.Context, in *GetWorkspaceListQuery, opts ...grpc.CallOption) (*GetWorkspaceListReply, error)
	GetProjectList(ctx context.Context, in *GetProjectListQuery, opts ...grpc.CallOption) (*GetProjectListReply, error)
	// mutation
	CreateOrg(ctx context.Context, in *CreateOrgCommand, opts ...grpc.CallOption) (*CreateOrgReply, error)
	UpdateOrgById(ctx context.Context, in *UpdateOrgByIdCommand, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteOrgById(ctx context.Context, in *DeleteOrgByIdCommand, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type orgServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrgServiceClient(cc grpc.ClientConnInterface) OrgServiceClient {
	return &orgServiceClient{cc}
}

func (c *orgServiceClient) GetOrgById(ctx context.Context, in *GetOrgByIdQuery, opts ...grpc.CallOption) (*GetOrgByIdReply, error) {
	out := new(GetOrgByIdReply)
	err := c.cc.Invoke(ctx, "/system.org.OrgService/GetOrgById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgServiceClient) GetOrgVisibilityById(ctx context.Context, in *GetOrgVisibilityByIdQuery, opts ...grpc.CallOption) (*GetOrgVisibilityByIdReply, error) {
	out := new(GetOrgVisibilityByIdReply)
	err := c.cc.Invoke(ctx, "/system.org.OrgService/GetOrgVisibilityById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgServiceClient) GetOrgList(ctx context.Context, in *GetOrgListQuery, opts ...grpc.CallOption) (*GetOrgListReply, error) {
	out := new(GetOrgListReply)
	err := c.cc.Invoke(ctx, "/system.org.OrgService/GetOrgList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgServiceClient) GetWorkspaceList(ctx context.Context, in *GetWorkspaceListQuery, opts ...grpc.CallOption) (*GetWorkspaceListReply, error) {
	out := new(GetWorkspaceListReply)
	err := c.cc.Invoke(ctx, "/system.org.OrgService/GetWorkspaceList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgServiceClient) GetProjectList(ctx context.Context, in *GetProjectListQuery, opts ...grpc.CallOption) (*GetProjectListReply, error) {
	out := new(GetProjectListReply)
	err := c.cc.Invoke(ctx, "/system.org.OrgService/GetProjectList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgServiceClient) CreateOrg(ctx context.Context, in *CreateOrgCommand, opts ...grpc.CallOption) (*CreateOrgReply, error) {
	out := new(CreateOrgReply)
	err := c.cc.Invoke(ctx, "/system.org.OrgService/CreateOrg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgServiceClient) UpdateOrgById(ctx context.Context, in *UpdateOrgByIdCommand, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/system.org.OrgService/UpdateOrgById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgServiceClient) DeleteOrgById(ctx context.Context, in *DeleteOrgByIdCommand, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/system.org.OrgService/DeleteOrgById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrgServiceServer is the server API for OrgService service.
// All implementations must embed UnimplementedOrgServiceServer
// for forward compatibility
type OrgServiceServer interface {
	// query
	GetOrgById(context.Context, *GetOrgByIdQuery) (*GetOrgByIdReply, error)
	GetOrgVisibilityById(context.Context, *GetOrgVisibilityByIdQuery) (*GetOrgVisibilityByIdReply, error)
	GetOrgList(context.Context, *GetOrgListQuery) (*GetOrgListReply, error)
	GetWorkspaceList(context.Context, *GetWorkspaceListQuery) (*GetWorkspaceListReply, error)
	GetProjectList(context.Context, *GetProjectListQuery) (*GetProjectListReply, error)
	// mutation
	CreateOrg(context.Context, *CreateOrgCommand) (*CreateOrgReply, error)
	UpdateOrgById(context.Context, *UpdateOrgByIdCommand) (*emptypb.Empty, error)
	DeleteOrgById(context.Context, *DeleteOrgByIdCommand) (*emptypb.Empty, error)
	mustEmbedUnimplementedOrgServiceServer()
}

// UnimplementedOrgServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOrgServiceServer struct {
}

func (UnimplementedOrgServiceServer) GetOrgById(context.Context, *GetOrgByIdQuery) (*GetOrgByIdReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrgById not implemented")
}
func (UnimplementedOrgServiceServer) GetOrgVisibilityById(context.Context, *GetOrgVisibilityByIdQuery) (*GetOrgVisibilityByIdReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrgVisibilityById not implemented")
}
func (UnimplementedOrgServiceServer) GetOrgList(context.Context, *GetOrgListQuery) (*GetOrgListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrgList not implemented")
}
func (UnimplementedOrgServiceServer) GetWorkspaceList(context.Context, *GetWorkspaceListQuery) (*GetWorkspaceListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkspaceList not implemented")
}
func (UnimplementedOrgServiceServer) GetProjectList(context.Context, *GetProjectListQuery) (*GetProjectListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProjectList not implemented")
}
func (UnimplementedOrgServiceServer) CreateOrg(context.Context, *CreateOrgCommand) (*CreateOrgReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrg not implemented")
}
func (UnimplementedOrgServiceServer) UpdateOrgById(context.Context, *UpdateOrgByIdCommand) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrgById not implemented")
}
func (UnimplementedOrgServiceServer) DeleteOrgById(context.Context, *DeleteOrgByIdCommand) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOrgById not implemented")
}
func (UnimplementedOrgServiceServer) mustEmbedUnimplementedOrgServiceServer() {}

// UnsafeOrgServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrgServiceServer will
// result in compilation errors.
type UnsafeOrgServiceServer interface {
	mustEmbedUnimplementedOrgServiceServer()
}

func RegisterOrgServiceServer(s grpc.ServiceRegistrar, srv OrgServiceServer) {
	s.RegisterService(&OrgService_ServiceDesc, srv)
}

func _OrgService_GetOrgById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrgByIdQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgServiceServer).GetOrgById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/system.org.OrgService/GetOrgById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgServiceServer).GetOrgById(ctx, req.(*GetOrgByIdQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrgService_GetOrgVisibilityById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrgVisibilityByIdQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgServiceServer).GetOrgVisibilityById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/system.org.OrgService/GetOrgVisibilityById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgServiceServer).GetOrgVisibilityById(ctx, req.(*GetOrgVisibilityByIdQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrgService_GetOrgList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrgListQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgServiceServer).GetOrgList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/system.org.OrgService/GetOrgList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgServiceServer).GetOrgList(ctx, req.(*GetOrgListQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrgService_GetWorkspaceList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkspaceListQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgServiceServer).GetWorkspaceList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/system.org.OrgService/GetWorkspaceList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgServiceServer).GetWorkspaceList(ctx, req.(*GetWorkspaceListQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrgService_GetProjectList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProjectListQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgServiceServer).GetProjectList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/system.org.OrgService/GetProjectList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgServiceServer).GetProjectList(ctx, req.(*GetProjectListQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrgService_CreateOrg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrgCommand)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgServiceServer).CreateOrg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/system.org.OrgService/CreateOrg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgServiceServer).CreateOrg(ctx, req.(*CreateOrgCommand))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrgService_UpdateOrgById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrgByIdCommand)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgServiceServer).UpdateOrgById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/system.org.OrgService/UpdateOrgById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgServiceServer).UpdateOrgById(ctx, req.(*UpdateOrgByIdCommand))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrgService_DeleteOrgById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOrgByIdCommand)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgServiceServer).DeleteOrgById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/system.org.OrgService/DeleteOrgById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgServiceServer).DeleteOrgById(ctx, req.(*DeleteOrgByIdCommand))
	}
	return interceptor(ctx, in, info, handler)
}

// OrgService_ServiceDesc is the grpc.ServiceDesc for OrgService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrgService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "system.org.OrgService",
	HandlerType: (*OrgServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOrgById",
			Handler:    _OrgService_GetOrgById_Handler,
		},
		{
			MethodName: "GetOrgVisibilityById",
			Handler:    _OrgService_GetOrgVisibilityById_Handler,
		},
		{
			MethodName: "GetOrgList",
			Handler:    _OrgService_GetOrgList_Handler,
		},
		{
			MethodName: "GetWorkspaceList",
			Handler:    _OrgService_GetWorkspaceList_Handler,
		},
		{
			MethodName: "GetProjectList",
			Handler:    _OrgService_GetProjectList_Handler,
		},
		{
			MethodName: "CreateOrg",
			Handler:    _OrgService_CreateOrg_Handler,
		},
		{
			MethodName: "UpdateOrgById",
			Handler:    _OrgService_UpdateOrgById_Handler,
		},
		{
			MethodName: "DeleteOrgById",
			Handler:    _OrgService_DeleteOrgById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "system_org.proto",
}