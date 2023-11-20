package project_impl

import (
	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/project"
	"github.com/quarkloop/quarkloop/pkg/store/repository"
)

type projectService struct {
	UserService    interface{}
	ProjectService interface{}
	QuotaService   interface{}

	dataStore *repository.Repository
}

func NewProjectService(ds *repository.Repository) project.Service {
	return &projectService{
		dataStore: ds,
	}
}

func (s *projectService) GetProjectList(p *project.GetProjectListParams) ([]model.Project, error) {
	orgList, err := s.dataStore.ListProjects(p.Context, p.OrgId, p.WorkspaceId)
	return orgList, err
}

func (s *projectService) GetProjectById(p *project.GetProjectByIdParams) (*model.Project, error) {
	org, err := s.dataStore.FindUniqueProject(p.Context, p.ProjectId)
	return org, err
}

func (s *projectService) GetProject(p *project.GetProjectParams) (*model.Project, error) {
	org, err := s.dataStore.FindFirstProject(p.Context, &p.Project)
	return org, err
}

func (s *projectService) CreateProject(p *project.CreateProjectParams) (*model.Project, error) {
	org, err := s.dataStore.CreateProject(p.Context, p.OrgId, p.WorkspaceId, &p.Project)
	return org, err
}

func (s *projectService) UpdateProjectById(p *project.UpdateProjectByIdParams) error {
	err := s.dataStore.UpdateProjectById(p.Context, p.ProjectId, &p.Project)
	return err
}

func (s *projectService) DeleteProjectById(p *project.DeleteProjectByIdParams) error {
	err := s.dataStore.DeleteProjectById(p.Context, p.ProjectId)
	return err
}
