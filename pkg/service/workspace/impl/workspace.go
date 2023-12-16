package workspace_impl

import (
	"github.com/quarkloop/quarkloop/pkg/service/workspace"
	"github.com/quarkloop/quarkloop/pkg/service/workspace/store"
)

type workspaceService struct {
	UserService      interface{}
	WorkspaceService interface{}
	QuotaService     interface{}

	dataStore store.WorkspaceStore
}

func NewWorkspaceService(ds store.WorkspaceStore) workspace.Service {
	return &workspaceService{
		dataStore: ds,
	}
}

func (s *workspaceService) GetWorkspaceList(p *workspace.GetWorkspaceListParams) ([]workspace.Workspace, error) {
	workspaceList, err := s.dataStore.ListWorkspaces(p.Context, p.OrgId)
	if err != nil {
		return nil, err
	}

	for i := range workspaceList {
		workspace := &workspaceList[i]
		workspace.GeneratePath()
	}
	return workspaceList, nil
}

func (s *workspaceService) GetWorkspaceById(p *workspace.GetWorkspaceByIdParams) (*workspace.Workspace, error) {
	workspace, err := s.dataStore.GetWorkspaceById(p.Context, p.WorkspaceId)
	if err != nil {
		return nil, err
	}

	workspace.GeneratePath()
	return workspace, nil
}

func (s *workspaceService) GetWorkspace(p *workspace.GetWorkspaceParams) (*workspace.Workspace, error) {
	workspace, err := s.dataStore.GetWorkspace(p.Context, p.OrgId, &p.Workspace)
	if err != nil {
		return nil, err
	}

	workspace.GeneratePath()
	return workspace, nil
}

func (s *workspaceService) CreateWorkspace(p *workspace.CreateWorkspaceParams) (*workspace.Workspace, error) {
	workspace, err := s.dataStore.CreateWorkspace(p.Context, p.OrgId, &p.Workspace)
	if err != nil {
		return nil, err
	}

	workspace.GeneratePath()
	return workspace, nil
}

func (s *workspaceService) UpdateWorkspaceById(p *workspace.UpdateWorkspaceByIdParams) error {
	err := s.dataStore.UpdateWorkspaceById(p.Context, p.WorkspaceId, &p.Workspace)
	return err
}

func (s *workspaceService) DeleteWorkspaceById(p *workspace.DeleteWorkspaceByIdParams) error {
	err := s.dataStore.DeleteWorkspaceById(p.Context, p.WorkspaceId)
	return err
}
