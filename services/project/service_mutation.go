package project

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	grpc "github.com/quarkloop/quarkloop/pkg/grpc/v1/system/project"
	"github.com/quarkloop/quarkloop/pkg/model"
	projectErrors "github.com/quarkloop/quarkloop/services/project/errors"
	"github.com/quarkloop/quarkloop/services/project/store"
)

func (s *projectService) CreateProject(ctx context.Context, cmd *grpc.CreateProjectCommand) (*grpc.CreateProjectReply, error) {
	if cmd.Name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "request missing required field: Name")
	} else if cmd.Description == "" {
		return nil, status.Errorf(codes.InvalidArgument, "request missing required field: Description")
	} else if cmd.CreatedBy == "" {
		return nil, status.Errorf(codes.InvalidArgument, "request missing required field: CreatedBy")
	}

	prj, err := s.store.CreateProject(ctx, &store.CreateProjectCommand{
		OrgId:       cmd.OrgId,
		WorkspaceId: cmd.WorkspaceId,
		CreatedBy:   cmd.CreatedBy,
		ScopeId:     cmd.ScopeId,
		Name:        cmd.Name,
		Description: cmd.Description,
		Visibility:  model.ScopeVisibility(cmd.Visibility),
	})
	if err != nil {
		if err == projectErrors.ErrProjectAlreadyExists {
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
	err := s.store.UpdateProjectById(ctx, &store.UpdateProjectByIdCommand{
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
	err := s.store.DeleteProjectById(ctx, &store.DeleteProjectByIdCommand{
		OrgId:       cmd.OrgId,
		WorkspaceId: cmd.WorkspaceId,
		ProjectId:   cmd.ProjectId,
	})
	return &emptypb.Empty{}, err
}
