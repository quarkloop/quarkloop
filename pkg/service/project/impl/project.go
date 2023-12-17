package project_impl

import (
	"context"

	"github.com/quarkloop/quarkloop/pkg/service/project"
	"github.com/quarkloop/quarkloop/pkg/service/project/store"
	"github.com/quarkloop/quarkloop/pkg/service/table_branch"
)

type projectService struct {
	store         store.ProjectStore
	branchService table_branch.Service
}

func NewProjectService(ds store.ProjectStore, branchService table_branch.Service) project.Service {
	return &projectService{
		store:         ds,
		branchService: branchService,
	}
}

func (s *projectService) GetProjectList(p *project.GetProjectListParams) ([]project.Project, error) {
	projectList, err := s.store.ListProjects(p.Context, p.OrgId, p.WorkspaceId)
	if err != nil {
		return nil, err
	}

	for i := range projectList {
		project := &projectList[i]
		project.GeneratePath()
	}
	return projectList, nil
}

func (s *projectService) GetProjectById(p *project.GetProjectByIdParams) (*project.Project, error) {
	project, err := s.store.GetProjectById(p.Context, p.ProjectId)
	if err != nil {
		return nil, err
	}

	project.GeneratePath()
	return project, nil
}

func (s *projectService) GetProject(p *project.GetProjectParams) (*project.Project, error) {
	project, err := s.store.GetProject(p.Context, &p.Project)
	if err != nil {
		return nil, err
	}

	project.GeneratePath()
	return project, nil
}

func (s *projectService) CreateProject(p *project.CreateProjectParams) (*project.Project, error) {
	project, err := s.store.CreateProject(p.Context, p.OrgId, p.WorkspaceId, &p.Project)
	if err != nil {
		return nil, err
	}

	project.GeneratePath()

	s.branchService.CreateTableBranch(&table_branch.CreateTableBranchParams{
		Context:   context.Background(),
		ProjectId: project.Id,
		Branch: &table_branch.TableBranch{
			Name:        "main",
			Type:        "main",
			Default:     true,
			Description: "main branch",
			CreatedBy:   "user",
		},
	})

	return project, nil
}

func (s *projectService) UpdateProjectById(p *project.UpdateProjectByIdParams) error {
	err := s.store.UpdateProjectById(p.Context, p.ProjectId, &p.Project)
	return err
}

func (s *projectService) DeleteProjectById(p *project.DeleteProjectByIdParams) error {
	err := s.store.DeleteProjectById(p.Context, p.ProjectId)
	return err
}
