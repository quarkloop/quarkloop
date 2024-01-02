package workspace

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/contextdata"
	"github.com/quarkloop/quarkloop/pkg/service/accesscontrol"
	"github.com/quarkloop/quarkloop/pkg/service/quota"
	"github.com/quarkloop/quarkloop/pkg/service/workspace"
)

func (s *WorkspaceApi) createWorkspace(ctx *gin.Context, cmd *workspace.CreateWorkspaceCommand) api.Reponse {
	// check permissions
	if err := s.evaluateCreatePermission(ctx, accesscontrol.ActionWorkspaceCreate, cmd.OrgId); err != nil {
		return api.Error(http.StatusInternalServerError, err) // TODO: change status
	}

	// check quotas
	quotaQuery := &quota.CheckCreateWorkspaceQuotaQuery{OrgId: cmd.OrgId}
	if err := s.quotaService.CheckCreateWorkspaceQuota(ctx, quotaQuery); err != nil {
		return api.Error(http.StatusInternalServerError, err) // TODO: change status
	}

	ws, err := s.workspaceService.CreateWorkspace(ctx, cmd)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	return api.Success(http.StatusCreated, ws)
}

func (s *WorkspaceApi) updateWorkspaceById(ctx *gin.Context, cmd *workspace.UpdateWorkspaceByIdCommand) api.Reponse {
	// check permissions
	if err := s.evaluatePermission(ctx, accesscontrol.ActionProjectUpdate, cmd.OrgId, cmd.WorkspaceId); err != nil {
		return api.Error(http.StatusInternalServerError, err) // TODO: change status
	}

	err := s.workspaceService.UpdateWorkspaceById(ctx, cmd)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	return api.Success(http.StatusOK, nil)
}

func (s *WorkspaceApi) deleteWorkspaceById(ctx *gin.Context, cmd *workspace.DeleteWorkspaceByIdCommand) api.Reponse {
	// check permissions
	if err := s.evaluatePermission(ctx, accesscontrol.ActionProjectDelete, cmd.OrgId, cmd.WorkspaceId); err != nil {
		return api.Error(http.StatusInternalServerError, err) // TODO: change status
	}

	err := s.workspaceService.DeleteWorkspaceById(ctx, cmd)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	return api.Success(http.StatusNoContent, nil)
}

func (s *WorkspaceApi) evaluateCreatePermission(ctx *gin.Context, permission string, orgId int) error {
	user := contextdata.GetUser(ctx)
	query := &accesscontrol.EvaluateFilterQuery{
		UserId: user.GetId(),
		OrgId:  orgId,
	}

	return s.aclService.Evaluate(ctx, permission, query)
}

func (s *WorkspaceApi) evaluatePermission(ctx *gin.Context, permission string, orgId, workspaceId int) error {
	user := contextdata.GetUser(ctx)
	query := &accesscontrol.EvaluateFilterQuery{
		UserId:      user.GetId(),
		OrgId:       orgId,
		WorkspaceId: workspaceId,
	}

	return s.aclService.Evaluate(ctx, permission, query)
}
