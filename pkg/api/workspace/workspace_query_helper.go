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
	ws, err := s.workspaceService.GetWorkspaceById(ctx, query)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	return api.Success(http.StatusOK, ws)
}

func (s *WorkspaceApi) getWorkspaceList(ctx *gin.Context) api.Response {
	if contextdata.IsUserAnonymous(ctx) {
		// anonymous user => return public workspaces
		wsList, err := s.workspaceService.GetWorkspaceList(ctx, &workspace.GetWorkspaceListQuery{Visibility: model.PublicVisibility})
		if err != nil {
			return api.Error(http.StatusInternalServerError, err) // TODO: change status code
		}

		return api.Success(http.StatusOK, &wsList)
	}

	user := contextdata.GetUser(ctx)
	// check permissions
	query := &accesscontrol.EvaluateFilterQuery{
		UserId: user.GetId(),
	}
	if err := s.aclService.Evaluate(ctx, accesscontrol.ActionOrgRead, query); err != nil {
		if err == accesscontrol.ErrPermissionDenied {
			// unauthorized user (permission denied) => return public workspaces
			wsList, err := s.workspaceService.GetWorkspaceList(ctx, &workspace.GetWorkspaceListQuery{
				UserId:     user.GetId(),
				Visibility: model.PublicVisibility,
			})
			if err != nil {
				return api.Error(http.StatusInternalServerError, err) // TODO: change status code
			}

			return api.Success(http.StatusOK, &wsList)
		}

		return api.Error(http.StatusInternalServerError, err)
	}

	// authorized user => return public + private workspaces
	wsList, err := s.workspaceService.GetWorkspaceList(ctx, &workspace.GetWorkspaceListQuery{
		UserId:     user.GetId(),
		Visibility: model.AllVisibility,
	})
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	return api.Success(http.StatusOK, &wsList)
}

func (s *WorkspaceApi) getProjectList(ctx *gin.Context, query *workspace.GetProjectListQuery) api.Response {
	if contextdata.IsUserAnonymous(ctx) {
		// anonymous user => return public projects
		projectList, err := s.workspaceService.GetProjectList(ctx, &workspace.GetProjectListQuery{
			OrgId:       query.OrgId,
			WorkspaceId: query.WorkspaceId,
			Visibility:  model.PublicVisibility,
		})
		if err != nil {
			return api.Error(http.StatusInternalServerError, err) // TODO: change status code
		}

		return api.Success(http.StatusOK, &projectList)
	}

	user := contextdata.GetUser(ctx)

	// check permissions
	err := s.aclService.Evaluate(ctx, accesscontrol.ActionProjectRead, &accesscontrol.EvaluateFilterQuery{UserId: user.GetId()})
	if err != nil {
		if err == accesscontrol.ErrPermissionDenied {
			// unauthorized user (permission denied) => return public projects
			projectList, err := s.workspaceService.GetProjectList(ctx, &workspace.GetProjectListQuery{
				OrgId:       query.OrgId,
				WorkspaceId: query.WorkspaceId,
				Visibility:  model.PublicVisibility,
			})
			if err != nil {
				return api.Error(http.StatusInternalServerError, err) // TODO: change status code
			}

			return api.Success(http.StatusOK, &projectList)
		}

		return api.Error(http.StatusInternalServerError, err)
	}

	// authorized user => return public + private projects
	projectList, err := s.workspaceService.GetProjectList(ctx, &workspace.GetProjectListQuery{
		OrgId:       query.OrgId,
		WorkspaceId: query.WorkspaceId,
		Visibility:  model.AllVisibility,
	})
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	return api.Success(http.StatusOK, &projectList)
}

func (s *WorkspaceApi) getMemberList(ctx *gin.Context, query *workspace.GetMemberListQuery) api.Response {
	ws, err := s.workspaceService.GetWorkspaceById(ctx, &workspace.GetWorkspaceByIdQuery{
		OrgId:       query.OrgId,
		WorkspaceId: query.WorkspaceId,
	})
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	isPrivate := ws.Visibility == model.PrivateVisibility

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
