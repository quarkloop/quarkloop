package workspace_impl

import (
	"context"
	"errors"

	"github.com/quarkloop/quarkloop/pkg/contextdata"
	"github.com/quarkloop/quarkloop/pkg/service/accesscontrol"
	"github.com/quarkloop/quarkloop/pkg/service/quota"
	"github.com/quarkloop/quarkloop/pkg/service/workspace"
	"github.com/quarkloop/quarkloop/pkg/service/workspace/store"
)

type workspaceService struct {
	store        store.WorkspaceStore
	aclService   accesscontrol.Service
	quotaService quota.Service
}

func NewWorkspaceService(ds store.WorkspaceStore, aclService accesscontrol.Service, quotaService quota.Service) workspace.Service {
	return &workspaceService{
		store:        ds,
		aclService:   aclService,
		quotaService: quotaService,
	}
}

func (s *workspaceService) GetWorkspaceList(ctx context.Context, params *workspace.GetWorkspaceListParams) ([]workspace.Workspace, error) {
	workspaceList, err := s.store.ListWorkspaces(ctx, params.OrgId)
	if err != nil {
		return nil, err
	}

	for i := range workspaceList {
		workspace := &workspaceList[i]
		workspace.GeneratePath()
	}
	return workspaceList, nil
}

func (s *workspaceService) GetWorkspaceById(ctx context.Context, params *workspace.GetWorkspaceByIdParams) (*workspace.Workspace, error) {
	workspace, err := s.store.GetWorkspaceById(ctx, params.WorkspaceId)
	if err != nil {
		return nil, err
	}

	workspace.GeneratePath()
	return workspace, nil
}

// func (s *workspaceService) GetWorkspace(ctx context.Context, params *workspace.GetWorkspaceParams) (*workspace.Workspace, error) {
// 	workspace, err := s.store.GetWorkspace(ctx, params.OrgId, &params.Workspace)
// 	if err != nil {
// 		return nil, err
// 	}

// 	workspace.GeneratePath()
// 	return workspace, nil
// }

func (s *workspaceService) CreateWorkspace(ctx context.Context, params *workspace.CreateWorkspaceParams) (*workspace.Workspace, error) {
	if contextdata.IsUserAnonymous(ctx) {
		return nil, errors.New("not authorized")
	}

	user := contextdata.GetUser(ctx)

	// check permissions
	err := s.aclService.Evaluate(ctx, accesscontrol.ActionWorkspaceCreate, &accesscontrol.EvaluateFilterParams{
		OrgId:  params.OrgId,
		UserId: user.GetId(),
	})
	if err != nil {
		return nil, err
	}

	// check quotas
	if err := s.quotaService.CheckCreateWorkspaceQuotaReached(ctx, params.OrgId); err != nil {
		return nil, err
	}

	workspace, err := s.store.CreateWorkspace(ctx, params.OrgId, &params.Workspace)
	if err != nil {
		return nil, err
	}
	workspace.GeneratePath()

	return workspace, nil
}

func (s *workspaceService) UpdateWorkspaceById(ctx context.Context, params *workspace.UpdateWorkspaceByIdParams) error {
	if contextdata.IsUserAnonymous(ctx) {
		return errors.New("not authorized")
	}

	user := contextdata.GetUser(ctx)
	scope := contextdata.GetScope(ctx)

	// check permissions
	err := s.aclService.Evaluate(ctx, accesscontrol.ActionWorkspaceUpdate, &accesscontrol.EvaluateFilterParams{
		UserId:      user.GetId(),
		OrgId:       scope.OrgId(),
		WorkspaceId: params.WorkspaceId,
	})
	if err != nil {
		return err
	}

	return s.store.UpdateWorkspaceById(ctx, params.WorkspaceId, &params.Workspace)
}

func (s *workspaceService) DeleteWorkspaceById(ctx context.Context, params *workspace.DeleteWorkspaceByIdParams) error {
	if contextdata.IsUserAnonymous(ctx) {
		return errors.New("not authorized")
	}

	user := contextdata.GetUser(ctx)
	scope := contextdata.GetScope(ctx)

	// check permissions
	err := s.aclService.Evaluate(ctx, accesscontrol.ActionWorkspaceDelete, &accesscontrol.EvaluateFilterParams{
		UserId:      user.GetId(),
		OrgId:       scope.OrgId(),
		WorkspaceId: params.WorkspaceId,
	})
	if err != nil {
		return err
	}

	return s.store.DeleteWorkspaceById(ctx, params.WorkspaceId)
}
