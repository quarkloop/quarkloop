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

func (s *projectService) GetProjectList(ctx context.Context, p *project.GetProjectListParams) ([]project.Project, error) {
	projectList, err := s.store.ListProjects(ctx, p.OrgId, p.WorkspaceId)
	if err != nil {
		return nil, err
	}

	for i := range projectList {
		project := &projectList[i]
		project.GeneratePath()
	}
	return projectList, nil
}

func (s *projectService) GetProjectById(ctx context.Context, p *project.GetProjectByIdParams) (*project.Project, error) {
	project, err := s.store.GetProjectById(ctx, p.ProjectId)
	if err != nil {
		return nil, err
	}

	project.GeneratePath()
	return project, nil
}

func (s *projectService) GetProject(ctx context.Context, p *project.GetProjectParams) (*project.Project, error) {
	project, err := s.store.GetProject(ctx, &p.Project)
	if err != nil {
		return nil, err
	}

	project.GeneratePath()
	return project, nil
}

func (s *projectService) CreateProject(ctx context.Context, p *project.CreateProjectParams) (*project.Project, error) {
	project, err := s.store.CreateProject(ctx, p.OrgId, p.WorkspaceId, &p.Project)
	if err != nil {
		return nil, err
	}

	project.GeneratePath()

	mainBranch, err := s.branchService.CreateTableBranch(ctx, &table_branch.CreateTableBranchParams{
		ProjectId: project.Id,
		Branch: &table_branch.TableBranch{
			Name:        "main",
			Type:        "main",
			Default:     true,
			Description: "main branch",
			CreatedBy:   "user",
		},
	})
	if err != nil {
		return nil, err
	}
	project.Branches = append(project.Branches, mainBranch)

	submissionBranch, err := s.branchService.CreateTableBranch(ctx, &table_branch.CreateTableBranchParams{
		ProjectId: project.Id,
		Branch: &table_branch.TableBranch{
			Name:        "submission",
			Type:        "submission",
			Default:     false,
			Description: "submission branch",
			CreatedBy:   "user",
		},
	})
	if err != nil {
		return nil, err
	}
	project.Branches = append(project.Branches, submissionBranch)

	return project, nil
}

func (s *projectService) UpdateProjectById(ctx context.Context, p *project.UpdateProjectByIdParams) error {
	err := s.store.UpdateProjectById(ctx, p.ProjectId, &p.Project)
	return err
}

func (s *projectService) DeleteProjectById(ctx context.Context, p *project.DeleteProjectByIdParams) error {
	err := s.store.DeleteProjectById(ctx, p.ProjectId)
	return err
}
