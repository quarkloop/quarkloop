package workspace

import (
	"context"

	"github.com/quarkloop/quarkloop/pkg/grpc/v1/system"
	grpc "github.com/quarkloop/quarkloop/pkg/grpc/v1/system/workspace"
	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/services/workspace/store"
)

func (s *workspaceService) GetOrgId(ctx context.Context, query *grpc.GetWorkspaceIdQuery) (*grpc.GetWorkspaceIdReply, error) {
	orgId, workspaceId, err := s.store.GetWorkspaceId(ctx, &store.GetWorkspaceIdQuery{
		OrgSid:       query.OrgSid,
		WorkspaceSid: query.WorkspaceSid,
	})
	if err != nil {
		return nil, err
	}

	reply := &grpc.GetWorkspaceIdReply{
		OrgId:       orgId,
		WorkspaceId: workspaceId,
	}
	return reply, nil
}

func (s *workspaceService) GetWorkspaceById(ctx context.Context, query *grpc.GetWorkspaceByIdQuery) (*grpc.GetWorkspaceByIdReply, error) {
	ws, err := s.store.GetWorkspaceById(ctx, &store.GetWorkspaceByIdQuery{
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
	visibility, err := s.store.GetWorkspaceVisibilityById(ctx, &store.GetWorkspaceVisibilityByIdQuery{
		OrgId:       query.OrgId,
		WorkspaceId: query.WorkspaceId,
	})
	if err != nil {
		return nil, err
	}

	reply := &grpc.GetWorkspaceVisibilityByIdReply{Visibility: int32(visibility)}
	return reply, nil
}

func (s *workspaceService) GetWorkspaceList(ctx context.Context, query *grpc.GetWorkspaceListQuery) (*grpc.GetWorkspaceListReply, error) {
	wsList, err := s.store.GetWorkspaceList(ctx, &store.GetWorkspaceListQuery{
		WorkspaceIdList: query.WorkspaceIdList,
		Visibility:      model.ScopeVisibility(query.Visibility),
	})
	if err != nil {
		return nil, err
	}

	res := &grpc.GetWorkspaceListReply{WorkspaceList: make([]*system.Workspace, len(wsList))}
	for i, ws := range wsList {
		if ws == nil {
			continue
		}

		ws.GeneratePath()
		proto := ws.Proto()
		res.WorkspaceList[i] = proto
	}

	return res, nil
}

func (s *workspaceService) GetProjectList(ctx context.Context, query *grpc.GetProjectListQuery) (*grpc.GetProjectListReply, error) {
	projectList, err := s.store.GetProjectList(ctx, &store.GetProjectListQuery{
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
