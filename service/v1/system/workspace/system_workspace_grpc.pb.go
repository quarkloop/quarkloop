// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: system_workspace.proto

package workspace

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

// WorkspaceServiceClient is the client API for WorkspaceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkspaceServiceClient interface {
	// query
	GetWorkspaceById(ctx context.Context, in *GetWorkspaceByIdQuery, opts ...grpc.CallOption) (*GetWorkspaceByIdReply, error)
	GetWorkspaceVisibilityById(ctx context.Context, in *GetWorkspaceVisibilityByIdQuery, opts ...grpc.CallOption) (*GetWorkspaceVisibilityByIdReply, error)
	GetWorkspaceList(ctx context.Context, in *GetWorkspaceListQuery, opts ...grpc.CallOption) (*GetWorkspaceListReply, error)
	GetProjectList(ctx context.Context, in *GetProjectListQuery, opts ...grpc.CallOption) (*GetProjectListReply, error)
	// mutation
	CreateWorkspace(ctx context.Context, in *CreateWorkspaceCommand, opts ...grpc.CallOption) (*CreateWorkspaceReply, error)
	UpdateWorkspaceById(ctx context.Context, in *UpdateWorkspaceByIdCommand, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteWorkspaceById(ctx context.Context, in *DeleteWorkspaceByIdCommand, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type workspaceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkspaceServiceClient(cc grpc.ClientConnInterface) WorkspaceServiceClient {
	return &workspaceServiceClient{cc}
}

func (c *workspaceServiceClient) GetWorkspaceById(ctx context.Context, in *GetWorkspaceByIdQuery, opts ...grpc.CallOption) (*GetWorkspaceByIdReply, error) {
	out := new(GetWorkspaceByIdReply)
	err := c.cc.Invoke(ctx, "/system.workspace.WorkspaceService/GetWorkspaceById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) GetWorkspaceVisibilityById(ctx context.Context, in *GetWorkspaceVisibilityByIdQuery, opts ...grpc.CallOption) (*GetWorkspaceVisibilityByIdReply, error) {
	out := new(GetWorkspaceVisibilityByIdReply)
	err := c.cc.Invoke(ctx, "/system.workspace.WorkspaceService/GetWorkspaceVisibilityById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) GetWorkspaceList(ctx context.Context, in *GetWorkspaceListQuery, opts ...grpc.CallOption) (*GetWorkspaceListReply, error) {
	out := new(GetWorkspaceListReply)
	err := c.cc.Invoke(ctx, "/system.workspace.WorkspaceService/GetWorkspaceList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) GetProjectList(ctx context.Context, in *GetProjectListQuery, opts ...grpc.CallOption) (*GetProjectListReply, error) {
	out := new(GetProjectListReply)
	err := c.cc.Invoke(ctx, "/system.workspace.WorkspaceService/GetProjectList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) CreateWorkspace(ctx context.Context, in *CreateWorkspaceCommand, opts ...grpc.CallOption) (*CreateWorkspaceReply, error) {
	out := new(CreateWorkspaceReply)
	err := c.cc.Invoke(ctx, "/system.workspace.WorkspaceService/CreateWorkspace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) UpdateWorkspaceById(ctx context.Context, in *UpdateWorkspaceByIdCommand, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/system.workspace.WorkspaceService/UpdateWorkspaceById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) DeleteWorkspaceById(ctx context.Context, in *DeleteWorkspaceByIdCommand, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/system.workspace.WorkspaceService/DeleteWorkspaceById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkspaceServiceServer is the server API for WorkspaceService service.
// All implementations must embed UnimplementedWorkspaceServiceServer
// for forward compatibility
type WorkspaceServiceServer interface {
	// query
	GetWorkspaceById(context.Context, *GetWorkspaceByIdQuery) (*GetWorkspaceByIdReply, error)
	GetWorkspaceVisibilityById(context.Context, *GetWorkspaceVisibilityByIdQuery) (*GetWorkspaceVisibilityByIdReply, error)
	GetWorkspaceList(context.Context, *GetWorkspaceListQuery) (*GetWorkspaceListReply, error)
	GetProjectList(context.Context, *GetProjectListQuery) (*GetProjectListReply, error)
	// mutation
	CreateWorkspace(context.Context, *CreateWorkspaceCommand) (*CreateWorkspaceReply, error)
	UpdateWorkspaceById(context.Context, *UpdateWorkspaceByIdCommand) (*emptypb.Empty, error)
	DeleteWorkspaceById(context.Context, *DeleteWorkspaceByIdCommand) (*emptypb.Empty, error)
	mustEmbedUnimplementedWorkspaceServiceServer()
}

// UnimplementedWorkspaceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWorkspaceServiceServer struct {
}

func (UnimplementedWorkspaceServiceServer) GetWorkspaceById(context.Context, *GetWorkspaceByIdQuery) (*GetWorkspaceByIdReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkspaceById not implemented")
}
func (UnimplementedWorkspaceServiceServer) GetWorkspaceVisibilityById(context.Context, *GetWorkspaceVisibilityByIdQuery) (*GetWorkspaceVisibilityByIdReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkspaceVisibilityById not implemented")
}
func (UnimplementedWorkspaceServiceServer) GetWorkspaceList(context.Context, *GetWorkspaceListQuery) (*GetWorkspaceListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkspaceList not implemented")
}
func (UnimplementedWorkspaceServiceServer) GetProjectList(context.Context, *GetProjectListQuery) (*GetProjectListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProjectList not implemented")
}
func (UnimplementedWorkspaceServiceServer) CreateWorkspace(context.Context, *CreateWorkspaceCommand) (*CreateWorkspaceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWorkspace not implemented")
}
func (UnimplementedWorkspaceServiceServer) UpdateWorkspaceById(context.Context, *UpdateWorkspaceByIdCommand) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWorkspaceById not implemented")
}
func (UnimplementedWorkspaceServiceServer) DeleteWorkspaceById(context.Context, *DeleteWorkspaceByIdCommand) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWorkspaceById not implemented")
}
func (UnimplementedWorkspaceServiceServer) mustEmbedUnimplementedWorkspaceServiceServer() {}

// UnsafeWorkspaceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkspaceServiceServer will
// result in compilation errors.
type UnsafeWorkspaceServiceServer interface {
	mustEmbedUnimplementedWorkspaceServiceServer()
}

func RegisterWorkspaceServiceServer(s grpc.ServiceRegistrar, srv WorkspaceServiceServer) {
	s.RegisterService(&WorkspaceService_ServiceDesc, srv)
}

func _WorkspaceService_GetWorkspaceById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkspaceByIdQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).GetWorkspaceById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/system.workspace.WorkspaceService/GetWorkspaceById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).GetWorkspaceById(ctx, req.(*GetWorkspaceByIdQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_GetWorkspaceVisibilityById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkspaceVisibilityByIdQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).GetWorkspaceVisibilityById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/system.workspace.WorkspaceService/GetWorkspaceVisibilityById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).GetWorkspaceVisibilityById(ctx, req.(*GetWorkspaceVisibilityByIdQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_GetWorkspaceList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkspaceListQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).GetWorkspaceList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/system.workspace.WorkspaceService/GetWorkspaceList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).GetWorkspaceList(ctx, req.(*GetWorkspaceListQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_GetProjectList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProjectListQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).GetProjectList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/system.workspace.WorkspaceService/GetProjectList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).GetProjectList(ctx, req.(*GetProjectListQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_CreateWorkspace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWorkspaceCommand)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).CreateWorkspace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/system.workspace.WorkspaceService/CreateWorkspace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).CreateWorkspace(ctx, req.(*CreateWorkspaceCommand))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_UpdateWorkspaceById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWorkspaceByIdCommand)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).UpdateWorkspaceById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/system.workspace.WorkspaceService/UpdateWorkspaceById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).UpdateWorkspaceById(ctx, req.(*UpdateWorkspaceByIdCommand))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_DeleteWorkspaceById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteWorkspaceByIdCommand)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).DeleteWorkspaceById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/system.workspace.WorkspaceService/DeleteWorkspaceById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).DeleteWorkspaceById(ctx, req.(*DeleteWorkspaceByIdCommand))
	}
	return interceptor(ctx, in, info, handler)
}

// WorkspaceService_ServiceDesc is the grpc.ServiceDesc for WorkspaceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WorkspaceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "system.workspace.WorkspaceService",
	HandlerType: (*WorkspaceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetWorkspaceById",
			Handler:    _WorkspaceService_GetWorkspaceById_Handler,
		},
		{
			MethodName: "GetWorkspaceVisibilityById",
			Handler:    _WorkspaceService_GetWorkspaceVisibilityById_Handler,
		},
		{
			MethodName: "GetWorkspaceList",
			Handler:    _WorkspaceService_GetWorkspaceList_Handler,
		},
		{
			MethodName: "GetProjectList",
			Handler:    _WorkspaceService_GetProjectList_Handler,
		},
		{
			MethodName: "CreateWorkspace",
			Handler:    _WorkspaceService_CreateWorkspace_Handler,
		},
		{
			MethodName: "UpdateWorkspaceById",
			Handler:    _WorkspaceService_UpdateWorkspaceById_Handler,
		},
		{
			MethodName: "DeleteWorkspaceById",
			Handler:    _WorkspaceService_DeleteWorkspaceById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "system_workspace.proto",
}
