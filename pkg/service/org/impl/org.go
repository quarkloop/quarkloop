package org_impl

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/org"
	"github.com/quarkloop/quarkloop/pkg/service/org/store"
	"github.com/quarkloop/quarkloop/service/system"
)

type orgService struct {
	store store.OrgStore

	system.UnimplementedOrgServiceServer
}

func NewOrgService(ds store.OrgStore) org.Service {
	return &orgService{store: ds}
}

func (s *orgService) GetOrgList(ctx context.Context, query *system.GetOrgListQuery) (*system.GetOrgListReply, error) {
	orgList, err := s.store.GetOrgList(ctx, &org.GetOrgListQuery{
		UserId:     query.UserId,
		Visibility: model.ScopeVisibility(query.Visibility),
	})
	if err != nil {
		return nil, err
	}

	res := &system.GetOrgListReply{OrgList: make([]*system.Org, len(orgList))}
	for i, org := range orgList {
		if org == nil {
			continue
		}

		org.GeneratePath()
		proto := org.Proto()
		res.OrgList[i] = proto
	}

	return res, nil
}

func (s *orgService) GetOrgById(ctx context.Context, query *system.GetOrgByIdQuery) (*system.GetOrgByIdReply, error) {
	org, err := s.store.GetOrgById(ctx, &org.GetOrgByIdQuery{OrgId: query.OrgId})
	if err != nil {
		return nil, err
	}

	org.GeneratePath()
	proto := org.Proto()

	reply := &system.GetOrgByIdReply{Org: proto}
	return reply, nil
}

func (s *orgService) GetOrgVisibilityById(ctx context.Context, query *system.GetOrgVisibilityByIdQuery) (*system.GetOrgVisibilityByIdReply, error) {
	visibility, err := s.store.GetOrgVisibilityById(ctx, &org.GetOrgVisibilityByIdQuery{OrgId: query.OrgId})
	if err != nil {
		return nil, err
	}

	reply := &system.GetOrgVisibilityByIdReply{Visibility: int32(visibility)}
	return reply, nil
}

func (s *orgService) CreateOrg(ctx context.Context, cmd *system.CreateOrgCommand) (*system.CreateOrgReply, error) {
	org, err := s.store.CreateOrg(ctx, &org.CreateOrgCommand{
		CreatedBy:   cmd.CreatedBy,
		ScopeId:     cmd.ScopeId,
		Name:        cmd.Name,
		Description: cmd.Description,
		Visibility:  model.ScopeVisibility(cmd.Visibility),
	})
	if err != nil {
		return nil, err
	}

	org.GeneratePath()
	proto := org.Proto()

	reply := &system.CreateOrgReply{Org: proto}
	return reply, nil
}

func (s *orgService) UpdateOrgById(ctx context.Context, cmd *system.UpdateOrgByIdCommand) (*emptypb.Empty, error) {
	err := s.store.UpdateOrgById(ctx, &org.UpdateOrgByIdCommand{
		UpdatedBy:   cmd.UpdatedBy,
		ScopeId:     cmd.ScopeId,
		Name:        cmd.Name,
		Description: cmd.Description,
		Visibility:  model.ScopeVisibility(cmd.Visibility),
	})
	return &emptypb.Empty{}, err
}

func (s *orgService) DeleteOrgById(ctx context.Context, cmd *system.DeleteOrgByIdCommand) (*emptypb.Empty, error) {
	err := s.store.DeleteOrgById(ctx, &org.DeleteOrgByIdCommand{OrgId: cmd.OrgId})
	return &emptypb.Empty{}, err
}

func (s *orgService) GetWorkspaceList(ctx context.Context, query *system.GetWorkspaceListQuery) (*system.GetWorkspaceListReply, error) {
	workspaceList, err := s.store.GetWorkspaceList(ctx, &org.GetWorkspaceListQuery{
		OrgId:      query.OrgId,
		Visibility: model.ScopeVisibility(query.Visibility),
	})
	if err != nil {
		return nil, err
	}

	reply := &system.GetWorkspaceListReply{WorkspaceList: make([]*system.Workspace, len(workspaceList))}
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

func (s *orgService) GetProjectList(ctx context.Context, query *system.GetProjectListQuery) (*system.GetProjectListReply, error) {
	projectList, err := s.store.GetProjectList(ctx, &org.GetProjectListQuery{
		OrgId:      query.OrgId,
		Visibility: model.ScopeVisibility(query.Visibility),
	})
	if err != nil {
		return nil, err
	}

	reply := &system.GetProjectListReply{ProjectList: make([]*system.Project, len(projectList))}
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

// func (s *orgService) GetUserAssignmentList(ctx context.Context, query *system.GetUserAssignmentListQuery) (*system.GetUserAssignmentListReply, error) {
// 	uaList, err := s.store.GetUserAssignmentList(ctx, query)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return uaList, nil
// }
