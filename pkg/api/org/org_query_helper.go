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
)

func (s *orgApi) getOrgById(ctx *gin.Context, query *org.GetOrgByIdQuery) api.Response {
	visibility, err := s.orgService.GetOrgVisibilityById(ctx, &org.GetOrgVisibilityByIdQuery{OrgId: query.OrgId})
	if err != nil {
		if err == org.ErrOrgNotFound {
			return api.Error(http.StatusNotFound, err)
		}
		return api.Error(http.StatusInternalServerError, err)
	}

	isPrivate := visibility == model.PrivateVisibility

	// anonymous user => return org not found error
	if isPrivate && contextdata.IsUserAnonymous(ctx) {
		return api.Error(http.StatusNotFound, org.ErrOrgNotFound)
	}
	if isPrivate {
		// check permissions
		user := contextdata.GetUser(ctx)
		query := &accesscontrol.EvaluateFilterQuery{
			UserId: user.GetId(),
			OrgId:  query.OrgId,
		}
		if err := s.aclService.Evaluate(ctx, accesscontrol.ActionOrgRead, query); err != nil {
			if err == accesscontrol.ErrPermissionDenied {
				// unauthorized user (permission denied) => return org not found error
				return api.Error(http.StatusNotFound, org.ErrOrgNotFound) // TODO: change status code
			}
			return api.Error(http.StatusInternalServerError, err)
		}
	}

	o, err := s.orgService.GetOrgById(ctx, &org.GetOrgByIdQuery{OrgId: query.OrgId})
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
	query := &org.GetOrgListQuery{Visibility: model.PublicVisibility}
	if !contextdata.IsUserAnonymous(ctx) {
		// check permissions
		user := contextdata.GetUser(ctx)
		evalQuery := &accesscontrol.EvaluateFilterQuery{UserId: user.GetId()}
		if err := s.aclService.Evaluate(ctx, accesscontrol.ActionOrgList, evalQuery); err != nil {
			if err != accesscontrol.ErrPermissionDenied {
				return api.Error(http.StatusInternalServerError, err)
			}
		} else {
			query.UserId = user.GetId()
			query.Visibility = model.AllVisibility
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

func (s *orgApi) getWorkspaceList(ctx *gin.Context, orgId int) api.Response {
	visibility := model.PublicVisibility
	if !contextdata.IsUserAnonymous(ctx) {
		// check permissions
		user := contextdata.GetUser(ctx)
		query := &accesscontrol.EvaluateFilterQuery{
			UserId: user.GetId(),
			OrgId:  orgId,
		}
		if err := s.aclService.Evaluate(ctx, accesscontrol.ActionWorkspaceList, query); err != nil {
			if err != accesscontrol.ErrPermissionDenied {
				return api.Error(http.StatusInternalServerError, err)
			}
		} else {
			visibility = model.AllVisibility
		}
	}

	wsList, err := s.orgService.GetWorkspaceList(ctx, &org.GetWorkspaceListQuery{
		OrgId:      orgId,
		Visibility: visibility,
	})
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	// anonymous user => return public workspaces
	// unauthorized user (permission denied) => return public workspaces
	// authorized user => return public + private workspaces
	return api.Success(http.StatusOK, &wsList)
}

func (s *orgApi) getProjectList(ctx *gin.Context, orgId int) api.Response {
	visibility := model.PublicVisibility
	if !contextdata.IsUserAnonymous(ctx) {
		// check permissions
		user := contextdata.GetUser(ctx)
		err := s.aclService.Evaluate(ctx, accesscontrol.ActionProjectList, &accesscontrol.EvaluateFilterQuery{
			UserId: user.GetId(),
			OrgId:  orgId,
		})
		if err != nil {
			if err != accesscontrol.ErrPermissionDenied {
				return api.Error(http.StatusInternalServerError, err)
			}
		} else {
			visibility = model.AllVisibility
		}
	}

	projectList, err := s.orgService.GetProjectList(ctx, &org.GetProjectListQuery{
		OrgId:      orgId,
		Visibility: visibility,
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
	visibility, err := s.orgService.GetOrgVisibilityById(ctx, &org.GetOrgVisibilityByIdQuery{OrgId: query.OrgId})
	if err != nil {
		if err == org.ErrOrgNotFound {
			return api.Error(http.StatusNotFound, err)
		}
		return api.Error(http.StatusInternalServerError, err)
	}

	isPrivate := visibility == model.PrivateVisibility

	// anonymous user => return org not found error
	if isPrivate && contextdata.IsUserAnonymous(ctx) {
		return api.Error(http.StatusNotFound, org.ErrOrgNotFound) // TODO: change sttaus code
	}
	if isPrivate {
		// check permissions
		user := contextdata.GetUser(ctx)
		query := &accesscontrol.EvaluateFilterQuery{
			UserId: user.GetId(),
			OrgId:  query.OrgId,
		}
		if err := s.aclService.Evaluate(ctx, accesscontrol.ActionOrgUserRead, query); err != nil {
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
