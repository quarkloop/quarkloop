package project_impl

import (
	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/project"
	"github.com/quarkloop/quarkloop/pkg/store/repository"
)

type projectService struct {
	dataStore *repository.Repository
}

func NewProjectService(ds *repository.Repository) project.Service {
	return &projectService{
		dataStore: ds,
	}
}

func (s *projectService) GetProjectList(p *project.GetProjectListParams) ([]model.Project, error) {
	projectList, err := s.dataStore.ListProjects(p.Context, p.OrgId, p.WorkspaceId)
	return projectList, err
}

func (s *projectService) GetProjectById(p *project.GetProjectByIdParams) (*model.Project, error) {
	project, err := s.dataStore.GetProjectById(p.Context, p.ProjectId)
	return project, err
}

func (s *projectService) GetProject(p *project.GetProjectParams) (*model.Project, error) {
	project, err := s.dataStore.GetProject(p.Context, &p.Project)
	return project, err
}

func (s *projectService) CreateProject(p *project.CreateProjectParams) (*model.Project, error) {
	project, err := s.dataStore.CreateProject(p.Context, p.OrgId, p.WorkspaceId, &p.Project)
	if err != nil {
		return nil, err
	}

	// serviceData := model.ProjectServiceData{
	// 	DiscussionsEnabled: false,
	// 	DiscussionsSettings: model.ProjectDiscussion{
	// 		MaxMessageLimit: 2048,
	// 	},
	// }

	// serviceList := []model.ProjectService{
	// 	{
	// 		Name:        "Discussions",
	// 		Type:        model.DiscussionsService,
	// 		Description: "Used for discussions",
	// 		Metadata:    json.RawMessage{},
	// 		Data:        serviceData,
	// 	},
	// }

	// create project services
	// var _, pErr = s.projectSvc.CreateBulkProjectService(
	// 	&project_table.CreateBulkProjectServiceParams{
	// 		Context:   p.Context,
	// 		ProjectId: project.Id,
	// 		//ProjectServiceList: serviceList,
	// 	},
	// )
	// if pErr != nil {
	// 	return nil, pErr
	// }

	return project, nil
}

func (s *projectService) UpdateProjectById(p *project.UpdateProjectByIdParams) error {
	err := s.dataStore.UpdateProjectById(p.Context, p.ProjectId, &p.Project)
	return err
}

func (s *projectService) DeleteProjectById(p *project.DeleteProjectByIdParams) error {
	err := s.dataStore.DeleteProjectById(p.Context, p.ProjectId)
	return err
}
