package workspace_impl

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/contextdata"
	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/accesscontrol"
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

func (s *workspaceService) GetWorkspaceList(ctx *gin.Context, params *workspace.GetWorkspaceListQuery) ([]*workspace.Workspace, error) {
	if contextdata.IsUserAnonymous(ctx) {
		// anonymous user => return public workspaces
		return s.getWorkspaceList(ctx, model.PublicVisibility, params)
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
			return s.getWorkspaceList(ctx, model.PublicVisibility, params)
		}
		return nil, err
	}

	// authorized user => return public + private workspaces
	return s.getWorkspaceList(ctx, model.AllVisibility, params)
}

func (s *workspaceService) getWorkspaceList(ctx context.Context, visibility model.ScopeVisibility, params *workspace.GetWorkspaceListQuery) ([]*workspace.Workspace, error) {
	workspaceList, err := s.store.ListWorkspaces(ctx, visibility, params.OrgId)
	if err != nil {
		return nil, err
	}

	for i := range workspaceList {
		workspace := workspaceList[i]
		workspace.GeneratePath()
	}
	return workspaceList, nil
}

func (s *workspaceService) GetWorkspaceById(ctx *gin.Context, params *workspace.GetWorkspaceByIdQuery) (*workspace.Workspace, error) {
	ws, err := s.store.GetWorkspaceById(ctx, params.WorkspaceId)
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
			WorkspaceId: params.WorkspaceId,
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

// func (s *workspaceService) GetWorkspace(ctx context.Context, params *workspace.GetWorkspaceQuery) (*workspace.Workspace, error) {
// 	workspace, err := s.store.GetWorkspace(ctx, params.OrgId, &params.Workspace)
// 	if err != nil {
// 		return nil, err
// 	}

// 	workspace.GeneratePath()
// 	return workspace, nil
// }

func (s *workspaceService) CreateWorkspace(ctx *gin.Context, params *workspace.CreateWorkspaceCommand) (*workspace.Workspace, error) {
	if contextdata.IsUserAnonymous(ctx) {
		return nil, errors.New("not authorized")
	}

	user := contextdata.GetUser(ctx)

	// check permissions
	err := s.aclService.Evaluate(ctx, accesscontrol.ActionWorkspaceCreate, &accesscontrol.EvaluateFilterParams{
		OrgId:  params.OrgId,
		UserId: user.GetId(),
	})
	if err != nil {
		return nil, err
	}

	// check quotas
	if err := s.quotaService.CheckCreateWorkspaceQuotaReached(ctx, params.OrgId); err != nil {
		return nil, err
	}

	ws, err := s.store.CreateWorkspace(ctx, params.OrgId, &params.Workspace)
	if err != nil {
		return nil, err
	}
	ws.GeneratePath()

	return ws, nil
}

func (s *workspaceService) UpdateWorkspaceById(ctx *gin.Context, params *workspace.UpdateWorkspaceByIdCommand) error {
	if contextdata.IsUserAnonymous(ctx) {
		return errors.New("not authorized")
	}

	user := contextdata.GetUser(ctx)
	scope := contextdata.GetScope(ctx)

	// check permissions
	err := s.aclService.Evaluate(ctx, accesscontrol.ActionWorkspaceUpdate, &accesscontrol.EvaluateFilterParams{
		UserId:      user.GetId(),
		OrgId:       scope.OrgId(),
		WorkspaceId: params.WorkspaceId,
	})
	if err != nil {
		return err
	}

	return s.store.UpdateWorkspaceById(ctx, params.WorkspaceId, &params.Workspace)
}

func (s *workspaceService) DeleteWorkspaceById(ctx *gin.Context, params *workspace.DeleteWorkspaceByIdCommand) error {
	if contextdata.IsUserAnonymous(ctx) {
		return errors.New("not authorized")
	}

	user := contextdata.GetUser(ctx)
	scope := contextdata.GetScope(ctx)

	// check permissions
	err := s.aclService.Evaluate(ctx, accesscontrol.ActionWorkspaceDelete, &accesscontrol.EvaluateFilterParams{
		UserId:      user.GetId(),
		OrgId:       scope.OrgId(),
		WorkspaceId: params.WorkspaceId,
	})
	if err != nil {
		return err
	}

	return s.store.DeleteWorkspaceById(ctx, params.WorkspaceId)
}
