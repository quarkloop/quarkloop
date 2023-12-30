package workspace_impl

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/contextdata"
	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/accesscontrol"
	"github.com/quarkloop/quarkloop/pkg/service/project"
	"github.com/quarkloop/quarkloop/pkg/service/quota"
	"github.com/quarkloop/quarkloop/pkg/service/workspace"
	"github.com/quarkloop/quarkloop/pkg/service/workspace/store"
)

type workspaceService struct {
	store        store.WorkspaceStore
	aclService   accesscontrol.Service
	quotaService quota.Service
}

func NewWorkspaceService(ds store.WorkspaceStore, aclService accesscontrol.Service, quotaService quota.Service) workspace.Service {
	return &workspaceService{
		store:        ds,
		aclService:   aclService,
		quotaService: quotaService,
	}
}

func (s *workspaceService) GetWorkspaceList(ctx *gin.Context, query *workspace.GetWorkspaceListQuery) ([]*workspace.Workspace, error) {
	if contextdata.IsUserAnonymous(ctx) {
		// anonymous user => return public workspaces
		return s.getWorkspaceList(ctx, model.PublicVisibility, query)
	}

	user := contextdata.GetUser(ctx)
	scope := contextdata.GetScope(ctx)

	// check permissions
	err := s.aclService.Evaluate(ctx, accesscontrol.ActionProjectRead, &accesscontrol.EvaluateFilterParams{
		UserId: user.GetId(),
		OrgId:  scope.OrgId(),
	})
	if err != nil {
		if err == accesscontrol.ErrPermissionDenied {
			// unauthorized user (permission denied) => return public workspaces
			return s.getWorkspaceList(ctx, model.PublicVisibility, query)
		}
		return nil, err
	}

	// authorized user => return public + private workspaces
	return s.getWorkspaceList(ctx, model.AllVisibility, query)
}

func (s *workspaceService) getWorkspaceList(ctx context.Context, visibility model.ScopeVisibility, query *workspace.GetWorkspaceListQuery) ([]*workspace.Workspace, error) {
	workspaceList, err := s.store.GetWorkspaceList(ctx, visibility, query.OrgId)
	if err != nil {
		return nil, err
	}

	for i := range workspaceList {
		workspace := workspaceList[i]
		workspace.GeneratePath()
	}
	return workspaceList, nil
}

func (s *workspaceService) GetWorkspaceById(ctx *gin.Context, query *workspace.GetWorkspaceByIdQuery) (*workspace.Workspace, error) {
	ws, err := s.store.GetWorkspaceById(ctx, query.WorkspaceId)
	if err != nil {
		return nil, err
	}

	isPrivate := *ws.Visibility == model.PrivateVisibility

	// anonymous user => return workspace not found error
	if isPrivate && contextdata.IsUserAnonymous(ctx) {
		return nil, workspace.ErrWorkspaceNotFound
	}
	if isPrivate {
		user := contextdata.GetUser(ctx)
		scope := contextdata.GetScope(ctx)

		// check permissions
		err := s.aclService.Evaluate(ctx, accesscontrol.ActionProjectRead, &accesscontrol.EvaluateFilterParams{
			UserId:      user.GetId(),
			OrgId:       scope.OrgId(),
			WorkspaceId: query.WorkspaceId,
		})
		if err != nil {
			if err == accesscontrol.ErrPermissionDenied {
				// unauthorized user (permission denied) => return workspace not found error
				return nil, workspace.ErrWorkspaceNotFound
			}
			return nil, err
		}
	}

	// anonymous and unauthorized user => return public workspace
	// authorized user => return public or private workspace
	ws.GeneratePath()
	return ws, nil
}

// func (s *workspaceService) GetWorkspace(ctx context.Context, query *workspace.GetWorkspaceQuery) (*workspace.Workspace, error) {
// 	workspace, err := s.store.GetWorkspace(ctx, query.OrgId, &query.Workspace)
// 	if err != nil {
// 		return nil, err
// 	}

// 	workspace.GeneratePath()
// 	return workspace, nil
// }

func (s *workspaceService) CreateWorkspace(ctx *gin.Context, cmd *workspace.CreateWorkspaceCommand) (*workspace.Workspace, error) {
	if contextdata.IsUserAnonymous(ctx) {
		return nil, errors.New("not authorized")
	}

	user := contextdata.GetUser(ctx)

	// check permissions
	err := s.aclService.Evaluate(ctx, accesscontrol.ActionWorkspaceCreate, &accesscontrol.EvaluateFilterParams{
		OrgId:  cmd.OrgId,
		UserId: user.GetId(),
	})
	if err != nil {
		return nil, err
	}

	// check quotas
	if err := s.quotaService.CheckCreateWorkspaceQuotaReached(ctx, &quota.CheckCreateWorkspaceQuotaReachedQuery{OrgId: cmd.OrgId}); err != nil {
		return nil, err
	}

	ws, err := s.store.CreateWorkspace(ctx, cmd.OrgId, &cmd.Workspace)
	if err != nil {
		return nil, err
	}
	ws.GeneratePath()

	return ws, nil
}

func (s *workspaceService) UpdateWorkspaceById(ctx *gin.Context, cmd *workspace.UpdateWorkspaceByIdCommand) error {
	if contextdata.IsUserAnonymous(ctx) {
		return errors.New("not authorized")
	}

	user := contextdata.GetUser(ctx)
	scope := contextdata.GetScope(ctx)

	// check permissions
	err := s.aclService.Evaluate(ctx, accesscontrol.ActionWorkspaceUpdate, &accesscontrol.EvaluateFilterParams{
		UserId:      user.GetId(),
		OrgId:       scope.OrgId(),
		WorkspaceId: cmd.WorkspaceId,
	})
	if err != nil {
		return err
	}

	return s.store.UpdateWorkspaceById(ctx, cmd.WorkspaceId, &cmd.Workspace)
}

func (s *workspaceService) DeleteWorkspaceById(ctx *gin.Context, cmd *workspace.DeleteWorkspaceByIdCommand) error {
	if contextdata.IsUserAnonymous(ctx) {
		return errors.New("not authorized")
	}

	user := contextdata.GetUser(ctx)
	scope := contextdata.GetScope(ctx)

	// check permissions
	err := s.aclService.Evaluate(ctx, accesscontrol.ActionWorkspaceDelete, &accesscontrol.EvaluateFilterParams{
		UserId:      user.GetId(),
		OrgId:       scope.OrgId(),
		WorkspaceId: cmd.WorkspaceId,
	})
	if err != nil {
		return err
	}

	return s.store.DeleteWorkspaceById(ctx, cmd.WorkspaceId)
}

func (s *workspaceService) GetProjectList(ctx *gin.Context, query *workspace.GetProjectListQuery) ([]*project.Project, error) {
	if contextdata.IsUserAnonymous(ctx) {
		// anonymous user => return public projects
		return s.getProjectList(ctx, model.PublicVisibility, query)
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
			return s.getProjectList(ctx, model.PublicVisibility, query)
		}
		return nil, err
	}

	// authorized user => return public + private projects
	return s.getProjectList(ctx, model.AllVisibility, query)
}

func (s *workspaceService) getProjectList(ctx context.Context, visibility model.ScopeVisibility, query *workspace.GetProjectListQuery) ([]*project.Project, error) {
	projectList, err := s.store.GetProjectList(ctx, visibility, query.OrgId, query.WorkspaceId)
	if err != nil {
		return nil, err
	}
	for i := range projectList {
		p := projectList[i]
		p.GeneratePath()
	}
	return projectList, nil
}
