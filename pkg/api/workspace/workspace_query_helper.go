package workspace

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/contextdata"
	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/accesscontrol"
	"github.com/quarkloop/quarkloop/pkg/service/user"
	"github.com/quarkloop/quarkloop/pkg/service/workspace"
)

func (s *WorkspaceApi) getWorkspaceById(ctx *gin.Context, query *workspace.GetWorkspaceByIdQuery) api.Response {
	visibility, err := s.workspaceService.GetWorkspaceVisibilityById(ctx, &workspace.GetWorkspaceVisibilityByIdQuery{
		OrgId:       query.OrgId,
		WorkspaceId: query.WorkspaceId,
	})
	if err != nil {
		if err == workspace.ErrWorkspaceNotFound {
			return api.Error(http.StatusNotFound, err)
		}
		return api.Error(http.StatusInternalServerError, err)
	}

	isPrivate := visibility == model.PrivateVisibility

	// anonymous user => return workspace not found error
	if isPrivate && contextdata.IsUserAnonymous(ctx) {
		return api.Error(http.StatusNotFound, workspace.ErrWorkspaceNotFound)
	}
	if isPrivate {
		user := contextdata.GetUser(ctx)
		// check permissions
		query := &accesscontrol.EvaluateQuery{
			Permission:  accesscontrol.ActionWorkspaceRead,
			UserId:      user.GetId(),
			OrgId:       query.OrgId,
			WorkspaceId: query.WorkspaceId,
		}
		access, err := s.aclService.EvaluateUserAccess(ctx, query)
		if err != nil {
			return api.Error(http.StatusInternalServerError, err)
		}
		if !access {
			// unauthorized user (permission denied) => return workspace not found error
			return api.Error(http.StatusNotFound, workspace.ErrWorkspaceNotFound)
		}
	}

	ws, err := s.workspaceService.GetWorkspaceById(ctx, query)
	if err != nil {
		if err == workspace.ErrWorkspaceNotFound {
			return api.Error(http.StatusNotFound, err)
		}
		return api.Error(http.StatusInternalServerError, err)
	}

	// anonymous and unauthorized user => return public workspace
	// authorized user => return public or private workspace
	return api.Success(http.StatusOK, ws)
}

func (s *WorkspaceApi) getWorkspaceList(ctx *gin.Context) api.Response {
	query := &workspace.GetWorkspaceListQuery{Visibility: model.PublicVisibility}
	if !contextdata.IsUserAnonymous(ctx) {
		// check permissions
		user := contextdata.GetUser(ctx)
		evalQuery := &accesscontrol.EvaluateQuery{
			Permission: accesscontrol.ActionWorkspaceList,
			UserId:     user.GetId(),
		}
		access, err := s.aclService.EvaluateUserAccess(ctx, evalQuery)
		if err != nil {
			return api.Error(http.StatusInternalServerError, err)
		}
		if !access {
			// unauthorized user (permission denied) => return workspace not found error
			return api.Error(http.StatusNotFound, workspace.ErrWorkspaceNotFound)
		}

		query.UserId = user.GetId()
		query.Visibility = model.AllVisibility
	}

	wsList, err := s.workspaceService.GetWorkspaceList(ctx, query)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	// anonymous user => return public workspaces
	// unauthorized user (permission denied) => return public workspaces
	// authorized user => return public + private workspaces
	return api.Success(http.StatusOK, &wsList)
}

func (s *WorkspaceApi) getProjectList(ctx *gin.Context, query *workspace.GetProjectListQuery) api.Response {
	getProjectListQuery := &workspace.GetProjectListQuery{
		OrgId:       query.OrgId,
		WorkspaceId: query.WorkspaceId,
		Visibility:  model.PublicVisibility,
	}
	if !contextdata.IsUserAnonymous(ctx) {
		// check permissions
		user := contextdata.GetUser(ctx)
		evalQuery := &accesscontrol.EvaluateQuery{
			Permission: accesscontrol.ActionProjectList,
			UserId:     user.GetId(),
		}
		access, err := s.aclService.EvaluateUserAccess(ctx, evalQuery)
		if err != nil {
			return api.Error(http.StatusInternalServerError, err)
		}
		if !access {
			// unauthorized user (permission denied) => return workspace not found error
			return api.Error(http.StatusNotFound, workspace.ErrWorkspaceNotFound)
		}

		getProjectListQuery.Visibility = model.AllVisibility
	}

	projectList, err := s.workspaceService.GetProjectList(ctx, getProjectListQuery)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	// anonymous user => return public projects
	// unauthorized user (permission denied) => return public projects
	// authorized user => return public + private projects
	return api.Success(http.StatusOK, &projectList)
}

func (s *WorkspaceApi) getMemberList(ctx *gin.Context, query *workspace.GetMemberListQuery) api.Response {
	visibility, err := s.workspaceService.GetWorkspaceVisibilityById(ctx, &workspace.GetWorkspaceVisibilityByIdQuery{
		OrgId:       query.OrgId,
		WorkspaceId: query.WorkspaceId,
	})
	if err != nil {
		if err == workspace.ErrWorkspaceNotFound {
			return api.Error(http.StatusNotFound, err)
		}
		return api.Error(http.StatusInternalServerError, err)
	}

	isPrivate := visibility == model.PrivateVisibility

	// anonymous user => return workspace not found error
	if isPrivate && contextdata.IsUserAnonymous(ctx) {
		return api.Error(http.StatusNotFound, workspace.ErrWorkspaceNotFound) // TODO: change sttaus code
	}
	if isPrivate {
		user := contextdata.GetUser(ctx)
		// check permissions
		evalQuery := &accesscontrol.EvaluateQuery{
			Permission:  accesscontrol.ActionWorkspaceUserRead,
			UserId:      user.GetId(),
			OrgId:       query.OrgId,
			WorkspaceId: query.WorkspaceId,
		}
		access, err := s.aclService.EvaluateUserAccess(ctx, evalQuery)
		if err != nil {
			return api.Error(http.StatusInternalServerError, err)
		}
		if !access {
			// unauthorized user (permission denied) => return workspace not found error
			return api.Error(http.StatusNotFound, workspace.ErrWorkspaceNotFound)
		}
	}

	uaList, err := s.workspaceService.GetUserAssignmentList(ctx, &workspace.GetUserAssignmentListQuery{
		OrgId:       query.OrgId,
		WorkspaceId: query.WorkspaceId,
	})
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	var memberList []*user.MemberDTO = []*user.MemberDTO{}
	for i := range uaList {
		ua := uaList[i]

		u, err := s.userService.GetUserById(ctx, &user.GetUserByIdQuery{UserId: ua.UserId})
		if err != nil {
			return api.Error(http.StatusInternalServerError, err)
		}

		user := user.MemberDTO{
			User:      u,
			Role:      ua.Role,
			CreatedAt: ua.CreatedAt,
			CreatedBy: ua.CreatedBy,
			UpdatedAt: ua.UpdatedAt,
			UpdatedBy: ua.UpdatedBy,
		}
		memberList = append(memberList, &user)
	}

	// anonymous and unauthorized user => return members of public workspace
	// authorized user => return members of public or private workspace
	return api.Success(http.StatusOK, &memberList)
}
