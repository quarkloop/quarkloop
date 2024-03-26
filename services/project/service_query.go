package project

import (
	"context"

	"github.com/quarkloop/quarkloop/pkg/grpc/v1/system"
	grpc "github.com/quarkloop/quarkloop/pkg/grpc/v1/system/project"
	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/services/project/store"
)

func (s *projectService) GetProjectId(ctx context.Context, query *grpc.GetProjectIdQuery) (*grpc.GetProjectIdReply, error) {
	orgId, workspaceId, projectId, err := s.store.GetProjectId(ctx, &store.GetProjectIdQuery{
		OrgSid:       query.OrgSid,
		WorkspaceSid: query.WorkspaceSid,
		ProjectSid:   query.ProjectSid,
	})
	if err != nil {
		return nil, err
	}

	reply := &grpc.GetProjectIdReply{
		OrgId:       orgId,
		WorkspaceId: workspaceId,
		ProjectId:   projectId,
	}
	return reply, nil
}

func (s *projectService) GetProjectList(ctx context.Context, query *grpc.GetProjectListQuery) (*grpc.GetProjectListReply, error) {
	projectList, err := s.store.GetProjectList(ctx, &store.GetProjectListQuery{
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
	prj, err := s.store.GetProjectById(ctx, &store.GetProjectByIdQuery{
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
	visibility, err := s.store.GetProjectVisibilityById(ctx, &store.GetProjectVisibilityByIdQuery{
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
