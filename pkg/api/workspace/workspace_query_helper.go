package workspace

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/contextdata"
	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/accesscontrol"
	"github.com/quarkloop/quarkloop/pkg/service/workspace"
)

func (s *WorkspaceApi) getWorkspaceById(ctx *gin.Context, query *workspace.GetWorkspaceByIdQuery) api.Reponse {
	ws, err := s.workspaceService.GetWorkspaceById(ctx, query)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	return api.Success(http.StatusOK, ws)
}

func (s *WorkspaceApi) getWorkspaceList(ctx *gin.Context) api.Reponse {
	user := contextdata.GetUser(ctx)

	wsList, err := s.workspaceService.GetWorkspaceList(ctx, &workspace.GetWorkspaceListQuery{
		UserId: user.GetId(),
	})
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	return api.Success(http.StatusOK, &wsList)
}

func (s *WorkspaceApi) getProjectList(ctx *gin.Context, query *workspace.GetProjectListQuery) api.Reponse {
	projectList, err := s.workspaceService.GetProjectList(ctx, query)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	return api.Success(http.StatusOK, &projectList)
}

func (s *WorkspaceApi) getUserList(ctx *gin.Context, query *workspace.GetUserListQuery) api.Reponse {
	ws, err := s.workspaceService.GetWorkspaceById(ctx, &workspace.GetWorkspaceByIdQuery{
		OrgId:       query.OrgId,
		WorkspaceId: query.WorkspaceId,
	})
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	isPrivate := *ws.Visibility == model.PrivateVisibility

	// anonymous user => return workspace not found error
	if isPrivate && contextdata.IsUserAnonymous(ctx) {
		return api.Error(http.StatusInternalServerError, workspace.ErrWorkspaceNotFound) // TODO: change sttaus code
	}
	if isPrivate {
		user := contextdata.GetUser(ctx)
		// check permissions
		query := &accesscontrol.EvaluateFilterQuery{
			UserId:      user.GetId(),
			OrgId:       query.OrgId,
			WorkspaceId: query.WorkspaceId,
		}
		if err := s.aclService.Evaluate(ctx, accesscontrol.ActionWorkspaceUserRead, query); err != nil {
			if err == accesscontrol.ErrPermissionDenied {
				// unauthorized user (permission denied) => return workspace not found error
				return api.Error(http.StatusInternalServerError, workspace.ErrWorkspaceNotFound) // TODO: change status code
			}

			return api.Error(http.StatusInternalServerError, err)
		}
	}

	userList, err := s.workspaceService.GetUserList(ctx, query)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	// anonymous and unauthorized user => return users of public workspace
	// authorized user => return users of public or private workspace
	return api.Success(http.StatusOK, &userList)
}
