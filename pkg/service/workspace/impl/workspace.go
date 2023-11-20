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
	orgList, err := s.dataStore.ListWorkspaces(p.Context, p.OrgId)
	return orgList, err
}

func (s *workspaceService) GetWorkspaceById(p *workspace.GetWorkspaceByIdParams) (*model.Workspace, error) {
	org, err := s.dataStore.FindUniqueWorkspace(p.Context, p.WorkspaceId)
	return org, err
}

func (s *workspaceService) GetWorkspace(p *workspace.GetWorkspaceParams) (*model.Workspace, error) {
	org, err := s.dataStore.FindFirstWorkspace(p.Context, p.OrgId, &p.Workspace)
	return org, err
}

func (s *workspaceService) CreateWorkspace(p *workspace.CreateWorkspaceParams) (*model.Workspace, error) {
	org, err := s.dataStore.CreateWorkspace(p.Context, p.OrgId, &p.Workspace)
	return org, err
}

func (s *workspaceService) UpdateWorkspaceById(p *workspace.UpdateWorkspaceByIdParams) error {
	err := s.dataStore.UpdateWorkspaceById(p.Context, p.WorkspaceId, &p.Workspace)
	return err
}

func (s *workspaceService) DeleteWorkspaceById(p *workspace.DeleteWorkspaceByIdParams) error {
	err := s.dataStore.DeleteWorkspaceById(p.Context, p.WorkspaceId)
	return err
}
