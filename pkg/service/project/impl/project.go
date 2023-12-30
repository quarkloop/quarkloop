package project_impl

import (
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
		err := s.aclService.Evaluate(ctx, accesscontrol.ActionProjectRead, &accesscontrol.EvaluateFilterParams{
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
// 		err := s.aclService.Evaluate(ctx, accesscontrol.ActionProjectRead, &accesscontrol.EvaluateFilterParams{
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
	if contextdata.IsUserAnonymous(ctx) {
		return nil, errors.New("not authorized")
	}

	user := contextdata.GetUser(ctx)
	scope := contextdata.GetScope(ctx)

	// check permissions
	err := s.aclService.Evaluate(ctx, accesscontrol.ActionProjectCreate, &accesscontrol.EvaluateFilterParams{
		UserId:      user.GetId(),
		OrgId:       scope.OrgId(),
		WorkspaceId: scope.WorkspaceId(),
	})
	if err != nil {
		return nil, err
	}

	// check quotas
	if err := s.quotaService.CheckCreateProjectQuotaReached(ctx, &quota.CheckCreateProjectQuotaReachedQuery{OrgId: cmd.OrgId}); err != nil {
		return nil, err
	}

	p, err := s.store.CreateProject(ctx, cmd.OrgId, cmd.WorkspaceId, &cmd.Project)
	if err != nil {
		return nil, err
	}
	p.GeneratePath()

	mainBranch, err := s.branchService.CreateTableBranch(ctx, &table_branch.CreateTableBranchParams{
		ProjectId: p.Id,
		Branch: &table_branch.TableBranch{
			Name:        "main",
			Type:        "main",
			Default:     true,
			Description: "main branch",
			CreatedBy:   user.Name, // TODO
		},
	})
	if err != nil {
		return nil, err
	}
	p.Branches = append(p.Branches, mainBranch)

	submissionBranch, err := s.branchService.CreateTableBranch(ctx, &table_branch.CreateTableBranchParams{
		ProjectId: p.Id,
		Branch: &table_branch.TableBranch{
			Name:        "submission",
			Type:        "submission",
			Default:     false,
			Description: "submission branch",
			CreatedBy:   "user", // TODO
		},
	})
	if err != nil {
		return nil, err
	}
	p.Branches = append(p.Branches, submissionBranch)

	return p, nil
}

func (s *projectService) UpdateProjectById(ctx *gin.Context, cmd *project.UpdateProjectByIdCommand) error {
	if contextdata.IsUserAnonymous(ctx) {
		return errors.New("not authorized")
	}

	user := contextdata.GetUser(ctx)
	scope := contextdata.GetScope(ctx)

	// check permissions
	err := s.aclService.Evaluate(ctx, accesscontrol.ActionProjectUpdate, &accesscontrol.EvaluateFilterParams{
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
	err := s.aclService.Evaluate(ctx, accesscontrol.ActionProjectDelete, &accesscontrol.EvaluateFilterParams{
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
