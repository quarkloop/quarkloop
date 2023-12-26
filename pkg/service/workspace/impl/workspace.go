package workspace_impl

import (
	"context"

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

func (s *workspaceService) GetWorkspaceList(ctx context.Context, p *workspace.GetWorkspaceListParams) ([]workspace.Workspace, error) {
	workspaceList, err := s.store.ListWorkspaces(ctx, p.OrgId)
	if err != nil {
		return nil, err
	}

	for i := range workspaceList {
		workspace := &workspaceList[i]
		workspace.GeneratePath()
	}
	return workspaceList, nil
}

func (s *workspaceService) GetWorkspaceById(ctx context.Context, p *workspace.GetWorkspaceByIdParams) (*workspace.Workspace, error) {
	workspace, err := s.store.GetWorkspaceById(ctx, p.WorkspaceId)
	if err != nil {
		return nil, err
	}

	workspace.GeneratePath()
	return workspace, nil
}

func (s *workspaceService) GetWorkspace(ctx context.Context, p *workspace.GetWorkspaceParams) (*workspace.Workspace, error) {
	workspace, err := s.store.GetWorkspace(ctx, p.OrgId, &p.Workspace)
	if err != nil {
		return nil, err
	}

	workspace.GeneratePath()
	return workspace, nil
}

func (s *workspaceService) CreateWorkspace(ctx context.Context, p *workspace.CreateWorkspaceParams) (*workspace.Workspace, error) {
	permission, err := s.aclService.Evaluate(ctx, accesscontrol.ActionWorkspaceCreate, &accesscontrol.EvaluateFilterParams{
		OrgId:  p.OrgId,
		UserId: 0,
	})
	if err != nil {
		return nil, err
	}
	if !permission {
		return nil, accesscontrol.ErrPermissionDenied
	}

	if err := s.quotaService.CheckCreateWorkspaceQuotaReached(ctx, p.OrgId); err != nil {
		return nil, err
	}

	workspace, err := s.store.CreateWorkspace(ctx, p.OrgId, &p.Workspace)
	if err != nil {
		return nil, err
	}
	workspace.GeneratePath()

	return workspace, nil
}

func (s *workspaceService) UpdateWorkspaceById(ctx context.Context, p *workspace.UpdateWorkspaceByIdParams) error {
	permission, err := s.aclService.Evaluate(ctx, accesscontrol.ActionWorkspaceUpdate, &accesscontrol.EvaluateFilterParams{
		OrgId:       p.WorkspaceId,
		WorkspaceId: p.WorkspaceId,
		UserId:      0,
	})
	if err != nil {
		return err
	}
	if !permission {
		return accesscontrol.ErrPermissionDenied
	}

	return s.store.UpdateWorkspaceById(ctx, p.WorkspaceId, &p.Workspace)
}

func (s *workspaceService) DeleteWorkspaceById(ctx context.Context, p *workspace.DeleteWorkspaceByIdParams) error {
	permission, err := s.aclService.Evaluate(ctx, accesscontrol.ActionWorkspaceDelete, &accesscontrol.EvaluateFilterParams{
		OrgId:       p.WorkspaceId,
		WorkspaceId: p.WorkspaceId,
		UserId:      0,
	})
	if err != nil {
		return err
	}
	if !permission {
		return accesscontrol.ErrPermissionDenied
	}

	return s.store.DeleteWorkspaceById(ctx, p.WorkspaceId)
}
