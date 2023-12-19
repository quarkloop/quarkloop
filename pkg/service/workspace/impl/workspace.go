package workspace_impl

import (
	"context"

	"github.com/quarkloop/quarkloop/pkg/service/quota"
	"github.com/quarkloop/quarkloop/pkg/service/workspace"
	"github.com/quarkloop/quarkloop/pkg/service/workspace/store"
)

type workspaceService struct {
	store        store.WorkspaceStore
	quotaService quota.Service
}

func NewWorkspaceService(ds store.WorkspaceStore, quotaService quota.Service) workspace.Service {
	return &workspaceService{
		store:        ds,
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
	_, err := s.quotaService.CheckWorkspaceQuotaReached(ctx, p.OrgId)
	if err != nil {
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
	err := s.store.UpdateWorkspaceById(ctx, p.WorkspaceId, &p.Workspace)
	return err
}

func (s *workspaceService) DeleteWorkspaceById(ctx context.Context, p *workspace.DeleteWorkspaceByIdParams) error {
	err := s.store.DeleteWorkspaceById(ctx, p.WorkspaceId)
	return err
}
