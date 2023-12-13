package workspace_impl

import (
	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/workspace"
	"github.com/quarkloop/quarkloop/pkg/store/repository"
)

type workspaceService struct {
	UserService      interface{}
	WorkspaceService interface{}
	QuotaService     interface{}

	dataStore *repository.Repository
}

func NewWorkspaceService(ds *repository.Repository) workspace.Service {
	return &workspaceService{
		dataStore: ds,
	}
}

func (s *workspaceService) GetWorkspaceList(p *workspace.GetWorkspaceListParams) ([]model.Workspace, error) {
	workspaceList, err := s.dataStore.ListWorkspaces(p.Context, p.OrgId)
	return workspaceList, err
}

func (s *workspaceService) GetWorkspaceById(p *workspace.GetWorkspaceByIdParams) (*model.Workspace, error) {
	workspace, err := s.dataStore.GetWorkspaceById(p.Context, p.WorkspaceId)
	return workspace, err
}

func (s *workspaceService) GetWorkspace(p *workspace.GetWorkspaceParams) (*model.Workspace, error) {
	workspace, err := s.dataStore.FindFirstWorkspace(p.Context, p.OrgId, &p.Workspace)
	return workspace, err
}

func (s *workspaceService) CreateWorkspace(p *workspace.CreateWorkspaceParams) (*model.Workspace, error) {
	workspace, err := s.dataStore.CreateWorkspace(p.Context, p.OrgId, &p.Workspace)
	return workspace, err
}

func (s *workspaceService) UpdateWorkspaceById(p *workspace.UpdateWorkspaceByIdParams) error {
	err := s.dataStore.UpdateWorkspaceById(p.Context, p.WorkspaceId, &p.Workspace)
	return err
}

func (s *workspaceService) DeleteWorkspaceById(p *workspace.DeleteWorkspaceByIdParams) error {
	err := s.dataStore.DeleteWorkspaceById(p.Context, p.WorkspaceId)
	return err
}
