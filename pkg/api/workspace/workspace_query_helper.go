package workspace

import (
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"

	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/contextdata"
	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/accesscontrol"
	"github.com/quarkloop/quarkloop/pkg/service/user"
	"github.com/quarkloop/quarkloop/pkg/service/workspace"
	"github.com/quarkloop/quarkloop/service/v1/system"
	grpc "github.com/quarkloop/quarkloop/service/v1/system/workspace"
)

func (s *WorkspaceApi) getWorkspaceById(ctx *gin.Context, query *workspace.GetWorkspaceByIdQuery) api.Response {
	visibility, err := s.workspaceService.GetWorkspaceVisibilityById(ctx, &grpc.GetWorkspaceVisibilityByIdQuery{
		OrgId:       query.OrgId,
		WorkspaceId: query.WorkspaceId,
	})
	if err != nil {
		if err == workspace.ErrWorkspaceNotFound {
			return api.Error(http.StatusNotFound, err)
		}
		return api.Error(http.StatusInternalServerError, err)
	}

	isPrivate := model.ScopeVisibility(visibility.GetVisibility()) == model.PrivateVisibility

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

	ws, err := s.workspaceService.GetWorkspaceById(ctx, &grpc.GetWorkspaceByIdQuery{
		OrgId:       query.OrgId,
		WorkspaceId: query.WorkspaceId,
	})
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
	query := &grpc.GetWorkspaceListQuery{Visibility: int32(model.PublicVisibility)}
	if !contextdata.IsUserAnonymous(ctx) {
		// check permissions
		user := contextdata.GetUser(ctx)
		aclQuery := &accesscontrol.GetWorkspaceListQuery{
			Permission: "all",
			UserId:     user.GetId(),
		}
		list, err := s.aclService.GetWorkspaceList(ctx, aclQuery)
		if err != nil {
			return api.Error(http.StatusInternalServerError, err)
		}

		if len(list) > 0 {
			query.WorkspaceIdList = list
			query.Visibility = int32(model.AllVisibility)
		}
	}

	wsList, err := s.workspaceService.GetWorkspaceList(ctx, query)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	// anonymous user => return public workspaces
	// unauthorized user (permission denied) => return public workspaces
	// authorized user => return public + private workspaces
	return api.Success(http.StatusOK, transformGrpcSlice(wsList.WorkspaceList))
}

func (s *WorkspaceApi) getProjectList(ctx *gin.Context, query *workspace.GetProjectListQuery) api.Response {
	var authorizedList []int32 = []int32{}
	queryPayload := &grpc.GetProjectListQuery{
		OrgId:       query.OrgId,
		WorkspaceId: query.WorkspaceId,
		Visibility:  int32(model.PublicVisibility),
	}
	if !contextdata.IsUserAnonymous(ctx) {
		// check permissions
		user := contextdata.GetUser(ctx)
		aclQuery := &accesscontrol.GetProjectListQuery{
			Permission: "all",
			UserId:     user.GetId(),
		}
		list, err := s.aclService.GetProjectList(ctx, aclQuery)
		if err != nil {
			return api.Error(http.StatusInternalServerError, err)
		}

		if len(list) > 0 {
			authorizedList = list
			queryPayload.Visibility = int32(model.AllVisibility)
		}
	}

	var newList []*system.Project = []*system.Project{}
	{
		projectList, err := s.workspaceService.GetProjectList(ctx, queryPayload)
		if err != nil {
			return api.Error(http.StatusInternalServerError, err)
		}

		for _, project := range projectList.GetProjectList() {
			if project.Visibility == int32(model.PublicVisibility) {
				newList = append(newList, project)
			} else if project.Visibility == int32(model.PrivateVisibility) && slices.Contains(authorizedList, project.Id) {
				newList = append(newList, project)
			}
		}
	}

	// anonymous user => return public projects
	// unauthorized user (permission denied) => return public projects
	// authorized user => return public + private projects
	return api.Success(http.StatusOK, newList)
}

func (s *WorkspaceApi) getMemberList(ctx *gin.Context, query *workspace.GetMemberListQuery) api.Response {
	visibility, err := s.workspaceService.GetWorkspaceVisibilityById(ctx, &grpc.GetWorkspaceVisibilityByIdQuery{
		OrgId:       query.OrgId,
		WorkspaceId: query.WorkspaceId,
	})
	if err != nil {
		if err == workspace.ErrWorkspaceNotFound {
			return api.Error(http.StatusNotFound, err)
		}
		return api.Error(http.StatusInternalServerError, err)
	}

	isPrivate := model.ScopeVisibility(visibility.GetVisibility()) == model.PrivateVisibility

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

	uaList, err := s.aclService.GetWorkspaceMemberList(ctx, &accesscontrol.GetWorkspaceMemberListQuery{WorkspaceId: query.WorkspaceId})
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}
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
