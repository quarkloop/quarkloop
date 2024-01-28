package workspace

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/contextdata"
	"github.com/quarkloop/quarkloop/pkg/service/accesscontrol"
	"github.com/quarkloop/quarkloop/pkg/service/quota"
	"github.com/quarkloop/quarkloop/pkg/service/workspace"
	grpc "github.com/quarkloop/quarkloop/service/v1/system/workspace"
)

func (s *WorkspaceApi) createWorkspace(ctx *gin.Context, cmd *workspace.CreateWorkspaceCommand) api.Response {
	// check permissions
	access, err := s.evaluateCreatePermission(ctx, accesscontrol.ActionWorkspaceCreate, cmd.OrgId)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err) // TODO: change status
	}
	if !access {
		// unauthorized user (permission denied) => return org not found error
		return api.Error(http.StatusForbidden, accesscontrol.ErrPermissionDenied) // TODO: change status code and error
	}

	// check quotas
	quotaQuery := &quota.CheckCreateWorkspaceQuotaQuery{OrgId: cmd.OrgId}
	if err := s.quotaService.CheckCreateWorkspaceQuota(ctx, quotaQuery); err != nil {
		return api.Error(http.StatusTooManyRequests, err)
	}

	ws, err := s.workspaceService.CreateWorkspace(ctx, &grpc.CreateWorkspaceCommand{
		OrgId:       cmd.OrgId,
		CreatedBy:   cmd.CreatedBy,
		ScopeId:     cmd.ScopeId,
		Name:        cmd.Name,
		Description: cmd.Description,
		Visibility:  int32(cmd.Visibility),
	})
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.AlreadyExists:
				return api.Error(http.StatusConflict, err)
			case codes.Internal:
				return api.Error(http.StatusInternalServerError, err)
			case codes.InvalidArgument:
				return api.Error(http.StatusBadRequest, err)
			}
		}
		return api.Error(http.StatusInternalServerError, err)
	}

	relationCmd := &accesscontrol.MakeParentResourceCommand{
		ParentResource:   "org",
		ParentResourceId: cmd.OrgId,
		ChildResource:    "workspace",
		ChildResourceId:  ws.Workspace.Id,
	}
	err = s.aclService.MakeParentResource(ctx, relationCmd)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	return api.Success(http.StatusCreated, ws.GetWorkspace())
}

func (s *WorkspaceApi) updateWorkspaceById(ctx *gin.Context, cmd *workspace.UpdateWorkspaceByIdCommand) api.Response {
	// check permissions
	access, err := s.evaluatePermission(ctx, accesscontrol.ActionWorkspaceUpdate, cmd.OrgId, cmd.WorkspaceId)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err) // TODO: change status
	}
	if !access {
		// unauthorized user (permission denied) => return org not found error
		return api.Error(http.StatusNotFound, workspace.ErrWorkspaceNotFound) // TODO: change status code
	}

	_, err = s.workspaceService.UpdateWorkspaceById(ctx, &grpc.UpdateWorkspaceByIdCommand{
		OrgId:       cmd.OrgId,
		WorkspaceId: cmd.WorkspaceId,
		UpdatedBy:   cmd.UpdatedBy,
		ScopeId:     cmd.ScopeId,
		Name:        cmd.Name,
		Description: cmd.Description,
		Visibility:  int32(cmd.Visibility),
	})
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	return api.Success(http.StatusOK, nil)
}

func (s *WorkspaceApi) deleteWorkspaceById(ctx *gin.Context, cmd *workspace.DeleteWorkspaceByIdCommand) api.Response {
	// check permissions
	access, err := s.evaluatePermission(ctx, accesscontrol.ActionWorkspaceDelete, cmd.OrgId, cmd.WorkspaceId)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err) // TODO: change status
	}
	if !access {
		// unauthorized user (permission denied) => return org not found error
		return api.Error(http.StatusNotFound, workspace.ErrWorkspaceNotFound) // TODO: change status code
	}

	_, err = s.workspaceService.DeleteWorkspaceById(ctx, &grpc.DeleteWorkspaceByIdCommand{
		OrgId:       cmd.OrgId,
		WorkspaceId: cmd.WorkspaceId,
	})
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	aclCommand := &accesscontrol.RevokeUserAccessCommand{
		OrgId:       cmd.OrgId,
		WorkspaceId: cmd.WorkspaceId,
	}
	err = s.aclService.RevokeUserAccess(ctx, aclCommand)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	return api.Success(http.StatusNoContent, nil)
}

func (s *WorkspaceApi) evaluateCreatePermission(ctx *gin.Context, permission string, orgId int32) (bool, error) {
	user := contextdata.GetUser(ctx)
	query := &accesscontrol.EvaluateQuery{
		Permission: permission,
		UserId:     user.GetId(),
		OrgId:      orgId,
	}

	return s.aclService.EvaluateUserAccess(ctx, query)
}

func (s *WorkspaceApi) evaluatePermission(ctx *gin.Context, permission string, orgId, workspaceId int32) (bool, error) {
	user := contextdata.GetUser(ctx)
	query := &accesscontrol.EvaluateQuery{
		Permission:  permission,
		UserId:      user.GetId(),
		OrgId:       orgId,
		WorkspaceId: workspaceId,
	}

	return s.aclService.EvaluateUserAccess(ctx, query)
}
