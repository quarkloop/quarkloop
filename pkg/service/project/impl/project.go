package project_impl

import (
	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/project"
	"github.com/quarkloop/quarkloop/pkg/service/project/store"
	"github.com/quarkloop/quarkloop/pkg/service/user"
)

type projectService struct {
	store store.ProjectStore
}

func NewProjectService(ds store.ProjectStore) project.Service {
	return &projectService{
		store: ds,
	}
}

func (s *projectService) GetProjectList(ctx *gin.Context, query *project.GetProjectListQuery) ([]*project.Project, error) {
	projectList, err := s.store.GetProjectList(ctx, query)
	if err != nil {
		return nil, err
	}

	for i := range projectList {
		project := projectList[i]
		project.GeneratePath()
	}
	return projectList, nil
}

func (s *projectService) GetProjectById(ctx *gin.Context, query *project.GetProjectByIdQuery) (*project.Project, error) {
	p, err := s.store.GetProjectById(ctx, query)
	if err != nil {
		return nil, err
	}

	p.GeneratePath()
	return p, nil
}

func (s *projectService) GetProjectVisibilityById(ctx *gin.Context, query *project.GetProjectVisibilityByIdQuery) (model.ScopeVisibility, error) {
	return s.store.GetProjectVisibilityById(ctx, query)
}

func (s *projectService) CreateProject(ctx *gin.Context, cmd *project.CreateProjectCommand) (*project.Project, error) {
	p, err := s.store.CreateProject(ctx, cmd)
	if err != nil {
		return nil, err
	}

	p.GeneratePath()
	return p, nil
}

func (s *projectService) UpdateProjectById(ctx *gin.Context, cmd *project.UpdateProjectByIdCommand) error {
	return s.store.UpdateProjectById(ctx, cmd)
}

func (s *projectService) DeleteProjectById(ctx *gin.Context, cmd *project.DeleteProjectByIdCommand) error {
	return s.store.DeleteProjectById(ctx, cmd)
}

func (s *projectService) GetUserAssignmentList(ctx *gin.Context, query *project.GetUserAssignmentListQuery) ([]*user.UserAssignment, error) {
	uaList, err := s.store.GetUserAssignmentList(ctx, query)
	if err != nil {
		return nil, err
	}

	return uaList, nil
}
