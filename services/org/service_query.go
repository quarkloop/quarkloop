package org

import (
	"context"

	"github.com/quarkloop/quarkloop/pkg/grpc/v1/system"
	grpc "github.com/quarkloop/quarkloop/pkg/grpc/v1/system/org"
	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/services/org/store"
)

func (s *orgService) GetOrgId(ctx context.Context, query *grpc.GetOrgIdQuery) (*grpc.GetOrgIdReply, error) {
	orgId, err := s.store.GetOrgId(ctx, query.OrgSid)
	if err != nil {
		return nil, err
	}

	return &grpc.GetOrgIdReply{OrgId: orgId}, nil
}

func (s *orgService) GetOrgList(ctx context.Context, query *grpc.GetOrgListQuery) (*grpc.GetOrgListReply, error) {
	orgList, err := s.store.GetOrgList(ctx, &store.GetOrgListQuery{
		OrgIdList:  query.OrgIdList,
		Visibility: model.ScopeVisibility(query.Visibility),
	})
	if err != nil {
		return nil, err
	}

	reply := &grpc.GetOrgListReply{OrgList: make([]*system.Org, len(orgList))}
	for i, org := range orgList {
		if org == nil {
			continue
		}

		org.GeneratePath()
		proto := org.Proto()
		reply.OrgList[i] = proto
	}

	return reply, nil
}

func (s *orgService) GetOrgById(ctx context.Context, query *grpc.GetOrgByIdQuery) (*grpc.GetOrgByIdReply, error) {
	org, err := s.store.GetOrgById(ctx, &store.GetOrgByIdQuery{OrgId: query.OrgId})
	if err != nil {
		return nil, err
	}

	org.GeneratePath()
	proto := org.Proto()

	reply := &grpc.GetOrgByIdReply{Org: proto}
	return reply, nil
}

func (s *orgService) GetOrgVisibilityById(ctx context.Context, query *grpc.GetOrgVisibilityByIdQuery) (*grpc.GetOrgVisibilityByIdReply, error) {
	visibility, err := s.store.GetOrgVisibilityById(ctx, &store.GetOrgVisibilityByIdQuery{OrgId: query.OrgId})
	if err != nil {
		return nil, err
	}

	reply := &grpc.GetOrgVisibilityByIdReply{Visibility: int32(visibility)}
	return reply, nil
}

func (s *orgService) GetWorkspaceList(ctx context.Context, query *grpc.GetWorkspaceListQuery) (*grpc.GetWorkspaceListReply, error) {
	workspaceList, err := s.store.GetWorkspaceList(ctx, &store.GetWorkspaceListQuery{
		OrgId:      query.OrgId,
		Visibility: model.ScopeVisibility(query.Visibility),
	})
	if err != nil {
		return nil, err
	}

	reply := &grpc.GetWorkspaceListReply{WorkspaceList: make([]*system.Workspace, len(workspaceList))}
	for i, ws := range workspaceList {
		if ws == nil {
			continue
		}

		ws.GeneratePath()
		proto := ws.Proto()
		reply.WorkspaceList[i] = proto
	}

	return reply, nil
}

func (s *orgService) GetProjectList(ctx context.Context, query *grpc.GetProjectListQuery) (*grpc.GetProjectListReply, error) {
	projectList, err := s.store.GetProjectList(ctx, &store.GetProjectListQuery{
		OrgId:      query.OrgId,
		Visibility: model.ScopeVisibility(query.Visibility),
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

// func (s *orgService) GetUserAssignmentList(ctx context.Context, query *grpc.GetUserAssignmentListQuery) (*grpc.GetUserAssignmentListReply, error) {
// 	uaList, err := s.store.GetUserAssignmentList(ctx, query)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return uaList, nil
// }
