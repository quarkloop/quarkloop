package project_impl

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/contextdata"
	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/accesscontrol"
	"github.com/quarkloop/quarkloop/pkg/service/project"
	"github.com/quarkloop/quarkloop/pkg/service/project/store"
	"github.com/quarkloop/quarkloop/pkg/service/quota"
	"github.com/quarkloop/quarkloop/pkg/service/table_branch"
)

type projectService struct {
	store store.ProjectStore

	aclService    accesscontrol.Service
	quotaService  quota.Service
	branchService table_branch.Service
}

func NewProjectService(
	ds store.ProjectStore,
	aclService accesscontrol.Service,
	quotaService quota.Service,
	branchService table_branch.Service,
) project.Service {
	return &projectService{
		store:         ds,
		aclService:    aclService,
		quotaService:  quotaService,
		branchService: branchService,
	}
}

func (s *projectService) GetProjectList(ctx *gin.Context, query *project.GetProjectListQuery) ([]*project.Project, error) {
	if contextdata.IsUserAnonymous(ctx) {
		// anonymous user => return public projects
		return s.getProjectList(ctx, model.PublicVisibility, query)
	}

	user := contextdata.GetUser(ctx)
	scope := contextdata.GetScope(ctx)

	// check permissions
	err := s.aclService.Evaluate(ctx, accesscontrol.ActionProjectRead, &accesscontrol.EvaluateFilterQuery{
		UserId: user.GetId(),
		OrgId:  scope.OrgId(),
	})
	if err != nil {
		if err == accesscontrol.ErrPermissionDenied {
			// unauthorized user (permission denied) => return public projects
			return s.getProjectList(ctx, model.PublicVisibility, query)
		}
		return nil, err
	}

	// authorized user => return public + private projects
	return s.getProjectList(ctx, model.AllVisibility, query)
}

func (s *projectService) getProjectList(ctx context.Context, visibility model.ScopeVisibility, query *project.GetProjectListQuery) ([]*project.Project, error) {
	projectList, err := s.store.GetProjectList(ctx, visibility, query.UserId)
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
	p, err := s.store.GetProjectById(ctx, query.ProjectId)
	if err != nil {
		return nil, err
	}

	isPrivate := *p.Visibility == model.PrivateVisibility

	// anonymous user => return project not found error
	if isPrivate && contextdata.IsUserAnonymous(ctx) {
		return nil, project.ErrProjectNotFound
	}
	if isPrivate {
		user := contextdata.GetUser(ctx)
		scope := contextdata.GetScope(ctx)

		// check permissions
		err := s.aclService.Evaluate(ctx, accesscontrol.ActionProjectRead, &accesscontrol.EvaluateFilterQuery{
			UserId:      user.GetId(),
			OrgId:       scope.OrgId(),
			WorkspaceId: scope.WorkspaceId(),
			ProjectId:   query.ProjectId,
		})
		if err != nil {
			if err == accesscontrol.ErrPermissionDenied {
				// unauthorized user (permission denied) => return project not found error
				return nil, project.ErrProjectNotFound
			}
			return nil, err
		}
	}

	// anonymous and unauthorized user => return public project
	// authorized user => return public or private project
	p.GeneratePath()
	return p, nil
}

// TODO
// func (s *projectService) GetProject(ctx context.Context, cmd *project.GetProjectQuery) (*project.Project, error) {
// 	p, err := s.store.GetProject(ctx, &cmd.Project)
// 	if err != nil {
// 		return nil, err
// 	}

// 	isPrivate := *p.Visibility == project.Private

// 	// anonymous user => return project not found error
// 	if isPrivate && contextdata.IsUserAnonymous(ctx) {
// 		return nil, project.ErrProjectNotFound
// 	}

// 	if isPrivate {
// 		// authorize signed-in user
// 		user := contextdata.GetUser(ctx)
// 		err := s.aclService.Evaluate(ctx, accesscontrol.ActionProjectRead, &accesscontrol.EvaluateFilterQuery{
// 			OrgId:       cmd.OrgId,
// 			WorkspaceId: cmd.Project.WorkspaceId,
// 			ProjectId:   cmd.Project.Id,
// 			UserId:      user.Id, // TODO
// 		})
// 		if err != nil {
// 			return nil, err
// 		}
// 	}

// 	// anonymous and unauthorized user => return public project
// 	// authorized user => return public or private project
// 	p.GeneratePath()
// 	return p, nil
// }

func (s *projectService) CreateProject(ctx *gin.Context, cmd *project.CreateProjectCommand) (*project.Project, error) {
	p, err := s.store.CreateProject(ctx, cmd.OrgId, cmd.WorkspaceId, &cmd.Project)
	if err != nil {
		return nil, err
	}
	p.GeneratePath()

	return p, nil
}

func (s *projectService) UpdateProjectById(ctx *gin.Context, cmd *project.UpdateProjectByIdCommand) error {
	if contextdata.IsUserAnonymous(ctx) {
		return errors.New("not authorized")
	}

	user := contextdata.GetUser(ctx)
	scope := contextdata.GetScope(ctx)

	// check permissions
	err := s.aclService.Evaluate(ctx, accesscontrol.ActionProjectUpdate, &accesscontrol.EvaluateFilterQuery{
		UserId:      user.GetId(),
		OrgId:       scope.OrgId(),
		WorkspaceId: scope.WorkspaceId(),
		ProjectId:   cmd.ProjectId,
	})
	if err != nil {
		return err
	}

	return s.store.UpdateProjectById(ctx, cmd.ProjectId, &cmd.Project)
}

func (s *projectService) DeleteProjectById(ctx *gin.Context, cmd *project.DeleteProjectByIdCommand) error {
	if contextdata.IsUserAnonymous(ctx) {
		return errors.New("not authorized")
	}

	user := contextdata.GetUser(ctx)
	scope := contextdata.GetScope(ctx)

	// check permissions
	err := s.aclService.Evaluate(ctx, accesscontrol.ActionProjectDelete, &accesscontrol.EvaluateFilterQuery{
		UserId:      user.GetId(),
		OrgId:       scope.OrgId(),
		WorkspaceId: scope.WorkspaceId(),
		ProjectId:   cmd.ProjectId,
	})
	if err != nil {
		return err
	}

	return s.store.DeleteProjectById(ctx, cmd.ProjectId)
}
