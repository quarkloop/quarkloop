package project_service_impl

import (
	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/project_service"
	"github.com/quarkloop/quarkloop/pkg/store/repository"
)

type projectService struct {
	dataStore *repository.Repository
}

func NewProjectServiceService(ds *repository.Repository) project_service.Service {
	return &projectService{
		dataStore: ds,
	}
}

func (s *projectService) GetProjectServiceList(p *project_service.GetProjectServiceListParams) ([]model.ProjectService, error) {
	projectList, err := s.dataStore.ListProjectServices(p.Context, p.ProjectId)
	return projectList, err
}

func (s *projectService) GetProjectServiceById(p *project_service.GetProjectServiceByIdParams) (*model.ProjectService, error) {
	project, err := s.dataStore.FindUniqueProjectService(p.Context, p.ProjectId, p.ProjectServiceId)
	return project, err
}

func (s *projectService) CreateProjectService(p *project_service.CreateProjectServiceParams) (*model.ProjectService, error) {
	project, err := s.dataStore.CreateProjectService(p.Context, p.ProjectId, p.ProjectService)
	return project, err
}

func (s *projectService) CreateBulkProjectService(p *project_service.CreateBulkProjectServiceParams) (int64, error) {
	project, err := s.dataStore.CreateBulkProjectService(p.Context, p.ProjectId, p.ProjectServiceList)
	return project, err
}

func (s *projectService) UpdateProjectServiceById(p *project_service.UpdateProjectServiceByIdParams) error {
	err := s.dataStore.UpdateProjectServiceById(p.Context, p.ProjectServiceId, p.ProjectService)
	return err
}

func (s *projectService) DeleteProjectServiceById(p *project_service.DeleteProjectServiceByIdParams) error {
	err := s.dataStore.DeleteProjectServiceById(p.Context, p.ProjectServiceId)
	return err
}
