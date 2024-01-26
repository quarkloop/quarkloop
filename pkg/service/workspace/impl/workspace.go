package workspace_impl

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/workspace"
	"github.com/quarkloop/quarkloop/pkg/service/workspace/store"
	"github.com/quarkloop/quarkloop/service/v1/system"
	grpc "github.com/quarkloop/quarkloop/service/v1/system/workspace"
)

type workspaceService struct {
	store store.WorkspaceStore

	grpc.UnimplementedWorkspaceServiceServer
}

func NewWorkspaceService(ds store.WorkspaceStore) workspace.Service {
	return &workspaceService{store: ds}
}

func (s *workspaceService) GetWorkspaceList(ctx context.Context, query *grpc.GetWorkspaceListQuery) (*grpc.GetWorkspaceListReply, error) {
	workspaceList, err := s.store.GetWorkspaceList(ctx, &workspace.GetWorkspaceListQuery{
		WorkspaceIdList: query.WorkspaceIdList,
		Visibility:      model.ScopeVisibility(query.Visibility),
	})
	if err != nil {
		return nil, err
	}

	res := &grpc.GetWorkspaceListReply{WorkspaceList: make([]*system.Workspace, len(workspaceList))}
	for i, ws := range workspaceList {
		if ws == nil {
			continue
		}

		ws.GeneratePath()
		proto := ws.Proto()
		res.WorkspaceList[i] = proto
	}

	return res, nil
}

func (s *workspaceService) GetWorkspaceById(ctx context.Context, query *grpc.GetWorkspaceByIdQuery) (*grpc.GetWorkspaceByIdReply, error) {
	ws, err := s.store.GetWorkspaceById(ctx, &workspace.GetWorkspaceByIdQuery{
		OrgId:       query.OrgId,
		WorkspaceId: query.WorkspaceId,
	})
	if err != nil {
		return nil, err
	}

	ws.GeneratePath()
	proto := ws.Proto()

	reply := &grpc.GetWorkspaceByIdReply{Workspace: proto}
	return reply, nil
}

func (s *workspaceService) GetWorkspaceVisibilityById(ctx context.Context, query *grpc.GetWorkspaceVisibilityByIdQuery) (*grpc.GetWorkspaceVisibilityByIdReply, error) {
	visibility, err := s.store.GetWorkspaceVisibilityById(ctx, &workspace.GetWorkspaceVisibilityByIdQuery{
		OrgId:       query.OrgId,
		WorkspaceId: query.WorkspaceId,
	})
	if err != nil {
		return nil, err
	}

	reply := &grpc.GetWorkspaceVisibilityByIdReply{Visibility: int32(visibility)}
	return reply, nil
}

func (s *workspaceService) CreateWorkspace(ctx context.Context, cmd *grpc.CreateWorkspaceCommand) (*grpc.CreateWorkspaceReply, error) {
	if cmd.Name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "request missing required field: Name")
	} else if cmd.Description == "" {
		return nil, status.Errorf(codes.InvalidArgument, "request missing required field: Description")
	} else if cmd.CreatedBy == "" {
		return nil, status.Errorf(codes.InvalidArgument, "request missing required field: CreatedBy")
	}

	ws, err := s.store.CreateWorkspace(ctx, &workspace.CreateWorkspaceCommand{
		OrgId:       cmd.OrgId,
		CreatedBy:   cmd.CreatedBy,
		ScopeId:     cmd.ScopeId,
		Name:        cmd.Name,
		Description: cmd.Description,
		Visibility:  model.ScopeVisibility(cmd.Visibility),
	})
	if err != nil {
		if err == workspace.ErrWorkspaceAlreadyExists {
			return nil, status.Errorf(codes.AlreadyExists, err.Error())
		}
		return nil, status.Errorf(codes.Internal, "something went wrong in server, error: %s", err.Error())
	}

	ws.GeneratePath()
	proto := ws.Proto()

	reply := &grpc.CreateWorkspaceReply{Workspace: proto}
	return reply, nil
}

func (s *workspaceService) UpdateWorkspaceById(ctx context.Context, cmd *grpc.UpdateWorkspaceByIdCommand) (*emptypb.Empty, error) {
	err := s.store.UpdateWorkspaceById(ctx, &workspace.UpdateWorkspaceByIdCommand{
		OrgId:       cmd.OrgId,
		WorkspaceId: cmd.WorkspaceId,
		UpdatedBy:   cmd.UpdatedBy,
		ScopeId:     cmd.ScopeId,
		Name:        cmd.Name,
		Description: cmd.Description,
		Visibility:  model.ScopeVisibility(cmd.Visibility),
	})
	return &emptypb.Empty{}, err
}

func (s *workspaceService) DeleteWorkspaceById(ctx context.Context, cmd *grpc.DeleteWorkspaceByIdCommand) (*emptypb.Empty, error) {
	err := s.store.DeleteWorkspaceById(ctx, &workspace.DeleteWorkspaceByIdCommand{
		OrgId:       cmd.OrgId,
		WorkspaceId: cmd.WorkspaceId,
	})
	return &emptypb.Empty{}, err
}

func (s *workspaceService) GetProjectList(ctx context.Context, query *grpc.GetProjectListQuery) (*grpc.GetProjectListReply, error) {
	projectList, err := s.store.GetProjectList(ctx, &workspace.GetProjectListQuery{
		OrgId:       query.OrgId,
		WorkspaceId: query.WorkspaceId,
		Visibility:  model.ScopeVisibility(query.Visibility),
	})
	if err != nil {
		return nil, err
	}

	reply := &grpc.GetProjectListReply{ProjectList: make([]*system.Project, len(projectList))}
	for i, project := range projectList {
		if project == nil {
			continue
		}

		project.GeneratePath()
		proto := project.Proto()
		reply.ProjectList[i] = proto
	}

	return reply, nil
}
