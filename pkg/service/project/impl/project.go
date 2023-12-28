package project_impl

import (
	"context"
	"errors"

	"github.com/quarkloop/quarkloop/pkg/contextdata"
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

func (s *projectService) GetProjectList(ctx context.Context, params *project.GetProjectListParams) ([]project.Project, error) {
	if contextdata.IsUserAnonymous(ctx) {
		// anonymous user => return public projects
		return s.getProjectList(ctx, project.Public, params)
	}

	user := contextdata.GetUser(ctx)
	scope := contextdata.GetScope(ctx)

	// check permissions
	err := s.aclService.Evaluate(ctx, accesscontrol.ActionProjectRead, &accesscontrol.EvaluateFilterParams{
		UserId:      user.GetId(),
		OrgId:       scope.OrgId(),
		WorkspaceId: scope.WorkspaceId(),
	})
	if err != nil {
		if err == accesscontrol.ErrPermissionDenied {
			// unauthorized user (permission denied) => return public projects
			return s.getProjectList(ctx, project.Public, params)
		}
		return nil, err
	}

	// authorized user => return public + private projects
	return s.getProjectList(ctx, project.All, params)
}

func (s *projectService) getProjectList(ctx context.Context, visibility project.ScopeVisibility, params *project.GetProjectListParams) ([]project.Project, error) {
	projectList, err := s.store.ListProjects(ctx, visibility, params.OrgId, params.WorkspaceId)
	if err != nil {
		return nil, err
	}
	for i := range projectList {
		p := &projectList[i]
		p.GeneratePath()
	}
	return projectList, nil
}

func (s *projectService) GetProjectById(ctx context.Context, params *project.GetProjectByIdParams) (*project.Project, error) {
	p, err := s.store.GetProjectById(ctx, params.ProjectId)
	if err != nil {
		return nil, err
	}

	isPrivate := *p.Visibility == project.Private

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
			ProjectId:   params.ProjectId,
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
// func (s *projectService) GetProject(ctx context.Context, params *project.GetProjectParams) (*project.Project, error) {
// 	p, err := s.store.GetProject(ctx, &params.Project)
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
// 			OrgId:       params.OrgId,
// 			WorkspaceId: params.Project.WorkspaceId,
// 			ProjectId:   params.Project.Id,
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

func (s *projectService) CreateProject(ctx context.Context, params *project.CreateProjectParams) (*project.Project, error) {
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
	if err := s.quotaService.CheckCreateProjectQuotaReached(ctx, params.OrgId); err != nil {
		return nil, err
	}

	p, err := s.store.CreateProject(ctx, params.OrgId, params.WorkspaceId, &params.Project)
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

func (s *projectService) UpdateProjectById(ctx context.Context, params *project.UpdateProjectByIdParams) error {
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
		ProjectId:   params.ProjectId,
	})
	if err != nil {
		return err
	}

	return s.store.UpdateProjectById(ctx, params.ProjectId, &params.Project)
}

func (s *projectService) DeleteProjectById(ctx context.Context, params *project.DeleteProjectByIdParams) error {
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
		ProjectId:   params.ProjectId,
	})
	if err != nil {
		return err
	}

	return s.store.DeleteProjectById(ctx, params.ProjectId)
}
