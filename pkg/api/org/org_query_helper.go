package org

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/contextdata"
	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/accesscontrol"
	"github.com/quarkloop/quarkloop/pkg/service/org"
	"github.com/quarkloop/quarkloop/pkg/service/user"
	"github.com/quarkloop/quarkloop/service/system"
)

func (s *orgApi) getOrgById(ctx *gin.Context, query *org.GetOrgByIdQuery) api.Response {
	visibility, err := s.orgService.GetOrgVisibilityById(ctx, &system.GetOrgVisibilityByIdQuery{OrgId: query.OrgId})
	if err != nil {
		if err == org.ErrOrgNotFound {
			return api.Error(http.StatusNotFound, err)
		}
		return api.Error(http.StatusInternalServerError, err)
	}

	isPrivate := model.ScopeVisibility(visibility.GetVisibility()) == model.PrivateVisibility

	// anonymous user => return org not found error
	if isPrivate && contextdata.IsUserAnonymous(ctx) {
		return api.Error(http.StatusNotFound, org.ErrOrgNotFound)
	}
	if isPrivate {
		// check permissions
		user := contextdata.GetUser(ctx)
		query := &accesscontrol.EvaluateQuery{
			Permission: accesscontrol.ActionOrgRead,
			UserId:     user.GetId(),
			OrgId:      query.OrgId,
		}
		if err := s.aclService.EvaluateUserAccess(ctx, query); err != nil {
			if err == accesscontrol.ErrPermissionDenied {
				// unauthorized user (permission denied) => return org not found error
				return api.Error(http.StatusNotFound, org.ErrOrgNotFound) // TODO: change status code
			}
			return api.Error(http.StatusInternalServerError, err)
		}
	}

	o, err := s.orgService.GetOrgById(ctx, &system.GetOrgByIdQuery{OrgId: query.OrgId})
	if err != nil {
		if err == org.ErrOrgNotFound {
			return api.Error(http.StatusNotFound, err)
		}
		return api.Error(http.StatusInternalServerError, err)
	}

	// anonymous and unauthorized user => return public org
	// authorized user => return public or private org
	return api.Success(http.StatusOK, o)
}

func (s *orgApi) getOrgList(ctx *gin.Context) api.Response {
	query := &system.GetOrgListQuery{Visibility: int32(model.PublicVisibility)}
	if !contextdata.IsUserAnonymous(ctx) {
		// check permissions
		user := contextdata.GetUser(ctx)
		evalQuery := &accesscontrol.EvaluateQuery{
			Permission: accesscontrol.ActionOrgList,
			UserId:     user.GetId(),
		}
		if err := s.aclService.EvaluateUserAccess(ctx, evalQuery); err != nil {
			if err != accesscontrol.ErrPermissionDenied {
				return api.Error(http.StatusInternalServerError, err)
			}
		} else {
			query.UserId = user.GetId()
			query.Visibility = int32(model.AllVisibility)
		}
	}

	orgList, err := s.orgService.GetOrgList(ctx, query)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	// anonymous user => return public orgs
	// unauthorized user (permission denied) => return public orgs
	// authorized user => return public + private orgs
	return api.Success(http.StatusOK, &orgList)
}

func (s *orgApi) getWorkspaceList(ctx *gin.Context, query *org.GetWorkspaceListQuery) api.Response {
	visibility := model.PublicVisibility
	if !contextdata.IsUserAnonymous(ctx) {
		// check permissions
		user := contextdata.GetUser(ctx)
		evalQuery := &accesscontrol.EvaluateQuery{
			Permission: accesscontrol.ActionWorkspaceList,
			UserId:     user.GetId(),
			OrgId:      query.OrgId,
		}
		if err := s.aclService.EvaluateUserAccess(ctx, evalQuery); err != nil {
			if err != accesscontrol.ErrPermissionDenied {
				return api.Error(http.StatusInternalServerError, err)
			}
		} else {
			visibility = model.AllVisibility
		}
	}

	wsList, err := s.orgService.GetWorkspaceList(ctx, &system.GetWorkspaceListQuery{
		OrgId:      query.OrgId,
		Visibility: int32(visibility),
	})
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	// anonymous user => return public workspaces
	// unauthorized user (permission denied) => return public workspaces
	// authorized user => return public + private workspaces
	return api.Success(http.StatusOK, &wsList)
}

func (s *orgApi) getProjectList(ctx *gin.Context, query *org.GetProjectListQuery) api.Response {
	visibility := model.PublicVisibility
	if !contextdata.IsUserAnonymous(ctx) {
		// check permissions
		user := contextdata.GetUser(ctx)
		evalQuery := &accesscontrol.EvaluateQuery{
			Permission: accesscontrol.ActionProjectList,
			UserId:     user.GetId(),
			OrgId:      query.OrgId,
		}
		err := s.aclService.EvaluateUserAccess(ctx, evalQuery)
		if err != nil {
			if err != accesscontrol.ErrPermissionDenied {
				return api.Error(http.StatusInternalServerError, err)
			}
		} else {
			visibility = model.AllVisibility
		}
	}

	projectList, err := s.orgService.GetProjectList(ctx, &system.GetProjectListQuery{
		OrgId:      query.OrgId,
		Visibility: int32(visibility),
	})
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	// anonymous user => return public projects
	// unauthorized user (permission denied) => return public projects
	// authorized user => return public + private projects
	return api.Success(http.StatusOK, &projectList)
}

func (s *orgApi) getMemberList(ctx *gin.Context, query *org.GetMemberListQuery) api.Response {
	visibility, err := s.orgService.GetOrgVisibilityById(ctx, &system.GetOrgVisibilityByIdQuery{OrgId: query.OrgId})
	if err != nil {
		if err == org.ErrOrgNotFound {
			return api.Error(http.StatusNotFound, err)
		}
		return api.Error(http.StatusInternalServerError, err)
	}

	isPrivate := model.ScopeVisibility(visibility.GetVisibility()) == model.PrivateVisibility

	// anonymous user => return org not found error
	if isPrivate && contextdata.IsUserAnonymous(ctx) {
		return api.Error(http.StatusNotFound, org.ErrOrgNotFound) // TODO: change sttaus code
	}
	if isPrivate {
		// check permissions
		user := contextdata.GetUser(ctx)
		evalQuery := &accesscontrol.EvaluateQuery{
			Permission: accesscontrol.ActionOrgUserRead,
			UserId:     user.GetId(),
			OrgId:      query.OrgId,
		}
		if err := s.aclService.EvaluateUserAccess(ctx, evalQuery); err != nil {
			if err == accesscontrol.ErrPermissionDenied {
				// unauthorized user (permission denied) => return org not found error
				return api.Error(http.StatusNotFound, org.ErrOrgNotFound) // TODO: change status code
			}

			return api.Error(http.StatusInternalServerError, err)
		}
	}

	uaList, err := s.orgService.GetUserAssignmentList(ctx, &org.GetUserAssignmentListQuery{OrgId: query.OrgId})
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

	// anonymous and unauthorized user => return members of public org
	// authorized user => return members of public or private org
	return api.Success(http.StatusOK, &memberList)
}
