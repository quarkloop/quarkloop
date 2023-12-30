package org_impl

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/contextdata"
	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/accesscontrol"
	"github.com/quarkloop/quarkloop/pkg/service/org"
	"github.com/quarkloop/quarkloop/pkg/service/org/store"
	"github.com/quarkloop/quarkloop/pkg/service/project"
	"github.com/quarkloop/quarkloop/pkg/service/quota"
	"github.com/quarkloop/quarkloop/pkg/service/workspace"
)

type orgService struct {
	store        store.OrgStore
	aclService   accesscontrol.Service
	quotaService quota.Service
}

func NewOrgService(ds store.OrgStore, aclService accesscontrol.Service, quotaService quota.Service) org.Service {
	return &orgService{
		store:        ds,
		aclService:   aclService,
		quotaService: quotaService,
	}
}

func (s *orgService) GetOrgList(ctx *gin.Context, query *org.GetOrgListQuery) ([]*org.Org, error) {
	if contextdata.IsUserAnonymous(ctx) {
		// anonymous user => return public orgs
		return s.getOrgList(ctx, model.PublicVisibility, query)
	}

	user := contextdata.GetUser(ctx)

	// check permissions
	err := s.aclService.Evaluate(ctx, accesscontrol.ActionOrgRead, &accesscontrol.EvaluateFilterParams{
		UserId: user.GetId(),
	})
	if err != nil {
		if err == accesscontrol.ErrPermissionDenied {
			// unauthorized user (permission denied) => return public orgs
			return s.getOrgList(ctx, model.PublicVisibility, query)
		}
		return nil, err
	}

	// authorized user => return public + private orgs
	return s.getOrgList(ctx, model.AllVisibility, query)
}

func (s *orgService) getOrgList(ctx *gin.Context, visibility model.ScopeVisibility, query *org.GetOrgListQuery) ([]*org.Org, error) {
	orgList, err := s.store.GetOrgList(ctx, visibility, query.UserId)
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
	o, err := s.store.GetOrgById(ctx, query.OrgId)
	if err != nil {
		return nil, err
	}

	isPrivate := *o.Visibility == model.PrivateVisibility

	// anonymous user => return org not found error
	if isPrivate && contextdata.IsUserAnonymous(ctx) {
		return nil, org.ErrOrgNotFound
	}
	if isPrivate {
		user := contextdata.GetUser(ctx)

		// check permissions
		err := s.aclService.Evaluate(ctx, accesscontrol.ActionOrgRead, &accesscontrol.EvaluateFilterParams{
			UserId: user.GetId(),
			OrgId:  query.OrgId,
		})
		if err != nil {
			if err == accesscontrol.ErrPermissionDenied {
				// unauthorized user (permission denied) => return org not found error
				return nil, org.ErrOrgNotFound
			}
			return nil, err
		}
	}

	// anonymous and unauthorized user => return public org
	// authorized user => return public or private org
	o.GeneratePath()
	return o, nil
}

// func (s *orgService) GetOrg(ctx *gin.Context, query *org.GetOrgQuery) (*org.Org, error) {
// 	org, err := s.store.GetOrg(ctx, &query.Org)
// 	if err != nil {
// 		return nil, err
// 	}

// 	org.GeneratePath()
// 	return org, nil
// }

func (s *orgService) CreateOrg(ctx *gin.Context, cmd *org.CreateOrgCommand) (*org.Org, error) {
	if contextdata.IsUserAnonymous(ctx) {
		return nil, errors.New("not authorized")
	}

	user := contextdata.GetUser(ctx)
	scope := contextdata.GetScope(ctx)

	// check permissions
	err := s.aclService.Evaluate(ctx, accesscontrol.ActionOrgCreate, &accesscontrol.EvaluateFilterParams{
		OrgId:  scope.OrgId(),
		UserId: user.GetId(),
	})
	if err != nil {
		return nil, err
	}

	// check quotas
	if err := s.quotaService.CheckCreateOrgQuotaReached(ctx, &quota.CheckCreateOrgQuotaReachedQuery{UserId: user.GetId()}); err != nil {
		return nil, err
	}

	o, err := s.store.CreateOrg(ctx, &cmd.Org)
	if err != nil {
		return nil, err
	}
	o.GeneratePath()

	return o, nil
}

func (s *orgService) UpdateOrgById(ctx *gin.Context, cmd *org.UpdateOrgByIdCommand) error {
	if contextdata.IsUserAnonymous(ctx) {
		return errors.New("not authorized")
	}

	user := contextdata.GetUser(ctx)

	// check permissions
	err := s.aclService.Evaluate(ctx, accesscontrol.ActionOrgUpdate, &accesscontrol.EvaluateFilterParams{
		OrgId:  cmd.OrgId,
		UserId: user.GetId(),
	})
	if err != nil {
		return err
	}

	return s.store.UpdateOrgById(ctx, cmd.OrgId, &cmd.Org)
}

func (s *orgService) DeleteOrgById(ctx *gin.Context, cmd *org.DeleteOrgByIdCommand) error {
	if contextdata.IsUserAnonymous(ctx) {
		return errors.New("not authorized")
	}

	user := contextdata.GetUser(ctx)

	// check permissions
	err := s.aclService.Evaluate(ctx, accesscontrol.ActionOrgDelete, &accesscontrol.EvaluateFilterParams{
		OrgId:  cmd.OrgId,
		UserId: user.GetId(),
	})
	if err != nil {
		return err
	}

	return s.store.DeleteOrgById(ctx, cmd.OrgId)
}

func (s *orgService) GetWorkspaceList(ctx *gin.Context, query *org.GetWorkspaceListQuery) ([]*workspace.Workspace, error) {
	if contextdata.IsUserAnonymous(ctx) {
		// anonymous user => return public workspaces
		return s.getWorkspaceList(ctx, model.PublicVisibility, query)
	}

	user := contextdata.GetUser(ctx)
	scope := contextdata.GetScope(ctx)

	// check permissions
	err := s.aclService.Evaluate(ctx, accesscontrol.ActionWorkspaceRead, &accesscontrol.EvaluateFilterParams{
		UserId: user.GetId(),
		OrgId:  scope.OrgId(),
	})
	if err != nil {
		if err == accesscontrol.ErrPermissionDenied {
			// unauthorized user (permission denied) => return public workspaces
			return s.getWorkspaceList(ctx, model.PublicVisibility, query)
		}
		return nil, err
	}

	// authorized user => return public + private workspaces
	return s.getWorkspaceList(ctx, model.AllVisibility, query)
}

func (s *orgService) getWorkspaceList(ctx context.Context, visibility model.ScopeVisibility, query *org.GetWorkspaceListQuery) ([]*workspace.Workspace, error) {
	workspaceList, err := s.store.GetWorkspaceList(ctx, visibility, query.OrgId)
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
	if contextdata.IsUserAnonymous(ctx) {
		// anonymous user => return public projects
		return s.getProjectList(ctx, model.PublicVisibility, query)
	}

	user := contextdata.GetUser(ctx)
	scope := contextdata.GetScope(ctx)

	// check permissions
	err := s.aclService.Evaluate(ctx, accesscontrol.ActionProjectRead, &accesscontrol.EvaluateFilterParams{
		UserId:      user.GetId(),
		OrgId:       scope.OrgId(),
		WorkspaceId: scope.WorkspaceId(),
	})
	if err != nil {
		if err == accesscontrol.ErrPermissionDenied {
			// unauthorized user (permission denied) => return public projects
			return s.getProjectList(ctx, model.PublicVisibility, query)
		}
		return nil, err
	}

	// authorized user => return public + private projects
	return s.getProjectList(ctx, model.AllVisibility, query)
}

func (s *orgService) getProjectList(ctx context.Context, visibility model.ScopeVisibility, query *org.GetProjectListQuery) ([]*project.Project, error) {
	projectList, err := s.store.GetProjectList(ctx, visibility, query.OrgId)
	if err != nil {
		return nil, err
	}
	for i := range projectList {
		p := projectList[i]
		p.GeneratePath()
	}
	return projectList, nil
}
