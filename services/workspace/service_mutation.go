package workspace

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	grpc "github.com/quarkloop/quarkloop/pkg/grpc/v1/system/workspace"
	"github.com/quarkloop/quarkloop/pkg/model"
	wsErros "github.com/quarkloop/quarkloop/services/workspace/errors"
	"github.com/quarkloop/quarkloop/services/workspace/store"
)

func (s *workspaceService) CreateWorkspace(ctx context.Context, cmd *grpc.CreateWorkspaceCommand) (*grpc.CreateWorkspaceReply, error) {
	if cmd.Name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "request missing required field: Name")
	} else if cmd.Description == "" {
		return nil, status.Errorf(codes.InvalidArgument, "request missing required field: Description")
	} else if cmd.CreatedBy == "" {
		return nil, status.Errorf(codes.InvalidArgument, "request missing required field: CreatedBy")
	}

	ws, err := s.store.CreateWorkspace(ctx, &store.CreateWorkspaceCommand{
		OrgId:       cmd.OrgId,
		CreatedBy:   cmd.CreatedBy,
		ScopeId:     cmd.ScopeId,
		Name:        cmd.Name,
		Description: cmd.Description,
		Visibility:  model.ScopeVisibility(cmd.Visibility),
	})
	if err != nil {
		if err == wsErros.ErrWorkspaceAlreadyExists {
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
	err := s.store.UpdateWorkspaceById(ctx, &store.UpdateWorkspaceByIdCommand{
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
	err := s.store.DeleteWorkspaceById(ctx, &store.DeleteWorkspaceByIdCommand{
		OrgId:       cmd.OrgId,
		WorkspaceId: cmd.WorkspaceId,
	})
	return &emptypb.Empty{}, err
}
