package project_impl

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/project"
	"github.com/quarkloop/quarkloop/pkg/service/project/store"
	"github.com/quarkloop/quarkloop/service/v1/system"
	grpc "github.com/quarkloop/quarkloop/service/v1/system/project"
)

type projectService struct {
	store store.ProjectStore

	grpc.UnimplementedProjectServiceServer
}

func NewProjectService(ds store.ProjectStore) project.Service {
	return &projectService{store: ds}
}

func (s *projectService) GetProjectList(ctx context.Context, query *grpc.GetProjectListQuery) (*grpc.GetProjectListReply, error) {
	projectList, err := s.store.GetProjectList(ctx, &project.GetProjectListQuery{
		ProjectIdList: query.ProjectIdList,
		Visibility:    model.ScopeVisibility(query.Visibility),
	})
	if err != nil {
		return nil, err
	}

	res := &grpc.GetProjectListReply{ProjectList: make([]*system.Project, len(projectList))}
	for i, prj := range projectList {
		if prj == nil {
			continue
		}

		prj.GeneratePath()
		proto := prj.Proto()
		res.ProjectList[i] = proto
	}

	return res, nil
}

func (s *projectService) GetProjectById(ctx context.Context, query *grpc.GetProjectByIdQuery) (*grpc.GetProjectByIdReply, error) {
	prj, err := s.store.GetProjectById(ctx, &project.GetProjectByIdQuery{
		OrgId:       query.OrgId,
		WorkspaceId: query.WorkspaceId,
		ProjectId:   query.ProjectId,
	})
	if err != nil {
		return nil, err
	}

	prj.GeneratePath()
	proto := prj.Proto()

	reply := &grpc.GetProjectByIdReply{Project: proto}
	return reply, nil
}

func (s *projectService) GetProjectVisibilityById(ctx context.Context, query *grpc.GetProjectVisibilityByIdQuery) (*grpc.GetProjectVisibilityByIdReply, error) {
	visibility, err := s.store.GetProjectVisibilityById(ctx, &project.GetProjectVisibilityByIdQuery{
		OrgId:       query.OrgId,
		WorkspaceId: query.WorkspaceId,
		ProjectId:   query.ProjectId,
	})
	if err != nil {
		return nil, err
	}

	reply := &grpc.GetProjectVisibilityByIdReply{Visibility: int32(visibility)}
	return reply, nil
}

func (s *projectService) CreateProject(ctx context.Context, cmd *grpc.CreateProjectCommand) (*grpc.CreateProjectReply, error) {
	if cmd.Name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "request missing required field: Name")
	} else if cmd.Description == "" {
		return nil, status.Errorf(codes.InvalidArgument, "request missing required field: Description")
	} else if cmd.CreatedBy == "" {
		return nil, status.Errorf(codes.InvalidArgument, "request missing required field: CreatedBy")
	}

	prj, err := s.store.CreateProject(ctx, &project.CreateProjectCommand{
		OrgId:       cmd.OrgId,
		WorkspaceId: cmd.WorkspaceId,
		CreatedBy:   cmd.CreatedBy,
		ScopeId:     cmd.ScopeId,
		Name:        cmd.Name,
		Description: cmd.Description,
		Visibility:  model.ScopeVisibility(cmd.Visibility),
	})
	if err != nil {
		if err == project.ErrProjectAlreadyExists {
			return nil, status.Errorf(codes.AlreadyExists, err.Error())
		}
		return nil, status.Errorf(codes.Internal, "something went wrong in server, error: %s", err.Error())
	}

	prj.GeneratePath()
	proto := prj.Proto()

	reply := &grpc.CreateProjectReply{Project: proto}
	return reply, nil
}

func (s *projectService) UpdateProjectById(ctx context.Context, cmd *grpc.UpdateProjectByIdCommand) (*emptypb.Empty, error) {
	err := s.store.UpdateProjectById(ctx, &project.UpdateProjectByIdCommand{
		OrgId:       cmd.OrgId,
		WorkspaceId: cmd.WorkspaceId,
		ProjectId:   cmd.ProjectId,
		UpdatedBy:   cmd.UpdatedBy,
		ScopeId:     cmd.ScopeId,
		Name:        cmd.Name,
		Description: cmd.Description,
		Visibility:  model.ScopeVisibility(cmd.Visibility),
	})
	return &emptypb.Empty{}, err
}

func (s *projectService) DeleteProjectById(ctx context.Context, cmd *grpc.DeleteProjectByIdCommand) (*emptypb.Empty, error) {
	err := s.store.DeleteProjectById(ctx, &project.DeleteProjectByIdCommand{
		OrgId:       cmd.OrgId,
		WorkspaceId: cmd.WorkspaceId,
		ProjectId:   cmd.ProjectId,
	})
	return &emptypb.Empty{}, err
}
