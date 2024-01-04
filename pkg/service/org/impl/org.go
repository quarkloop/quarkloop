package org_impl

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/org"
	"github.com/quarkloop/quarkloop/pkg/service/org/store"
	"github.com/quarkloop/quarkloop/pkg/service/project"
	"github.com/quarkloop/quarkloop/pkg/service/user"
	"github.com/quarkloop/quarkloop/pkg/service/workspace"
)

type orgService struct {
	store store.OrgStore
}

func NewOrgService(ds store.OrgStore) org.Service {
	return &orgService{
		store: ds,
	}
}

func (s *orgService) GetOrgList(ctx *gin.Context, query *org.GetOrgListQuery) ([]*org.Org, error) {
	orgList, err := s.store.GetOrgList(ctx, query)
	if err != nil {
		return nil, err
	}

	for i := range orgList {
		org := orgList[i]
		org.GeneratePath()
	}

	return orgList, nil
}

func (s *orgService) GetOrgById(ctx *gin.Context, query *org.GetOrgByIdQuery) (*org.Org, error) {
	o, err := s.store.GetOrgById(ctx, query)
	if err != nil {
		return nil, err
	}

	o.GeneratePath()
	return o, nil
}

func (s *orgService) GetOrgVisibilityById(ctx *gin.Context, query *org.GetOrgVisibilityByIdQuery) (model.ScopeVisibility, error) {
	return s.store.GetOrgVisibilityById(ctx, query)
}

func (s *orgService) CreateOrg(ctx *gin.Context, cmd *org.CreateOrgCommand) (*org.Org, error) {
	o, err := s.store.CreateOrg(ctx, cmd)
	if err != nil {
		return nil, err
	}
	o.GeneratePath()

	return o, nil
}

func (s *orgService) UpdateOrgById(ctx *gin.Context, cmd *org.UpdateOrgByIdCommand) error {
	return s.store.UpdateOrgById(ctx, cmd)
}

func (s *orgService) DeleteOrgById(ctx *gin.Context, cmd *org.DeleteOrgByIdCommand) error {
	return s.store.DeleteOrgById(ctx, cmd)
}

func (s *orgService) GetWorkspaceList(ctx *gin.Context, query *org.GetWorkspaceListQuery) ([]*workspace.Workspace, error) {
	return s.getWorkspaceList(ctx, query)
}

func (s *orgService) getWorkspaceList(ctx context.Context, query *org.GetWorkspaceListQuery) ([]*workspace.Workspace, error) {
	workspaceList, err := s.store.GetWorkspaceList(ctx, query)
	if err != nil {
		return nil, err
	}

	for i := range workspaceList {
		p := workspaceList[i]
		p.GeneratePath()
	}

	return workspaceList, nil
}

func (s *orgService) GetProjectList(ctx *gin.Context, query *org.GetProjectListQuery) ([]*project.Project, error) {
	projectList, err := s.store.GetProjectList(ctx, query)
	if err != nil {
		return nil, err
	}

	for i := range projectList {
		p := projectList[i]
		p.GeneratePath()
	}

	return projectList, nil
}

func (s *orgService) GetUserAssignmentList(ctx *gin.Context, query *org.GetUserAssignmentListQuery) ([]*user.UserAssignment, error) {
	uaList, err := s.store.GetUserAssignmentList(ctx, query)
	if err != nil {
		return nil, err
	}

	return uaList, nil
}
