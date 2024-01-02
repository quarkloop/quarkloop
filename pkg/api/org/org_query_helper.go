package org

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/contextdata"
	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/accesscontrol"
	"github.com/quarkloop/quarkloop/pkg/service/org"
)

func (s *OrgApi) getOrgById(ctx *gin.Context, orgId int) api.Reponse {
	o, err := s.orgService.GetOrgById(ctx, &org.GetOrgByIdQuery{
		OrgId: orgId,
	})
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	isPrivate := *o.Visibility == model.PrivateVisibility

	// anonymous user => return org not found error
	if isPrivate && contextdata.IsUserAnonymous(ctx) {
		return api.Error(http.StatusInternalServerError, org.ErrOrgNotFound)
	}
	if isPrivate {
		user := contextdata.GetUser(ctx)
		// check permissions
		query := &accesscontrol.EvaluateFilterQuery{
			UserId: user.GetId(),
			OrgId:  orgId,
		}
		if err := s.aclService.Evaluate(ctx, accesscontrol.ActionOrgRead, query); err != nil {
			if err == accesscontrol.ErrPermissionDenied {
				// unauthorized user (permission denied) => return org not found error
				return api.Error(http.StatusInternalServerError, org.ErrOrgNotFound) // TODO: change status code
			}

			return api.Error(http.StatusInternalServerError, err)
		}
	}

	// anonymous and unauthorized user => return public org
	// authorized user => return public or private org
	return api.Success(http.StatusOK, o)
}

func (s *OrgApi) getOrgList(ctx *gin.Context) api.Reponse {
	if contextdata.IsUserAnonymous(ctx) {
		// anonymous user => return public orgs
		orgList, err := s.orgService.GetOrgList(ctx, &org.GetOrgListQuery{Visibility: model.PublicVisibility})
		if err != nil {
			return api.Error(http.StatusInternalServerError, err)
		}

		return api.Success(http.StatusOK, &orgList)
	}

	user := contextdata.GetUser(ctx)
	// check permissions
	query := &accesscontrol.EvaluateFilterQuery{UserId: user.GetId()}
	if err := s.aclService.Evaluate(ctx, accesscontrol.ActionOrgRead, query); err != nil {
		if err == accesscontrol.ErrPermissionDenied {
			// unauthorized user (permission denied) => return public orgs
			orgList, err := s.orgService.GetOrgList(ctx, &org.GetOrgListQuery{Visibility: model.PublicVisibility})
			if err != nil {
				return api.Error(http.StatusInternalServerError, err) // TODO: change status code
			}

			return api.Success(http.StatusOK, &orgList)
		}

		return api.Error(http.StatusInternalServerError, err)
	}

	// authorized user => return public + private orgs
	orgList, err := s.orgService.GetOrgList(ctx, &org.GetOrgListQuery{
		UserId:     user.GetId(),
		Visibility: model.AllVisibility,
	})
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	return api.Success(http.StatusOK, &orgList)
}

func (s *OrgApi) getWorkspaceList(ctx *gin.Context, orgId int) api.Reponse {
	if contextdata.IsUserAnonymous(ctx) {
		// anonymous user => return public workspaces
		wsList, err := s.orgService.GetWorkspaceList(ctx, &org.GetWorkspaceListQuery{
			OrgId:      orgId,
			Visibility: model.PublicVisibility,
		})
		if err != nil {
			return api.Error(http.StatusInternalServerError, err) // TODO: change status code
		}

		return api.Success(http.StatusOK, &wsList)
	}

	user := contextdata.GetUser(ctx)
	// check permissions
	query := &accesscontrol.EvaluateFilterQuery{
		UserId: user.GetId(),
		OrgId:  orgId,
	}
	if err := s.aclService.Evaluate(ctx, accesscontrol.ActionOrgRead, query); err != nil {
		if err == accesscontrol.ErrPermissionDenied {
			// unauthorized user (permission denied) => return public workspaces
			wsList, err := s.orgService.GetWorkspaceList(ctx, &org.GetWorkspaceListQuery{
				OrgId:      orgId,
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
	wsList, err := s.orgService.GetWorkspaceList(ctx, &org.GetWorkspaceListQuery{
		OrgId:      orgId,
		Visibility: model.AllVisibility,
	})
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	return api.Success(http.StatusOK, &wsList)
}

func (s *OrgApi) getProjectList(ctx *gin.Context, orgId int) api.Reponse {
	if contextdata.IsUserAnonymous(ctx) {
		// anonymous user => return public projects
		projectList, err := s.orgService.GetProjectList(ctx, &org.GetProjectListQuery{
			OrgId:      orgId,
			Visibility: model.PublicVisibility,
		})
		if err != nil {
			return api.Error(http.StatusInternalServerError, err) // TODO: change status code
		}

		return api.Success(http.StatusOK, &projectList)
	}

	user := contextdata.GetUser(ctx)

	// check permissions
	err := s.aclService.Evaluate(ctx, accesscontrol.ActionProjectRead, &accesscontrol.EvaluateFilterQuery{
		UserId: user.GetId(),
		OrgId:  orgId,
	})
	if err != nil {
		if err == accesscontrol.ErrPermissionDenied {
			// unauthorized user (permission denied) => return public projects
			projectList, err := s.orgService.GetProjectList(ctx, &org.GetProjectListQuery{
				OrgId:      orgId,
				Visibility: model.PublicVisibility,
			})
			if err != nil {
				return api.Error(http.StatusInternalServerError, err) // TODO: change status code
			}

			return api.Success(http.StatusOK, &projectList)
		}

		return api.Error(http.StatusInternalServerError, err)
	}

	// authorized user => return public + private projects
	projectList, err := s.orgService.GetProjectList(ctx, &org.GetProjectListQuery{
		OrgId:      orgId,
		Visibility: model.AllVisibility,
	})
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	return api.Success(http.StatusOK, &projectList)
}

func (s *OrgApi) getUserList(ctx *gin.Context, query *org.GetUserListQuery) api.Reponse {
	o, err := s.orgService.GetOrgById(ctx, &org.GetOrgByIdQuery{OrgId: query.OrgId})
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	isPrivate := *o.Visibility == model.PrivateVisibility

	// anonymous user => return org not found error
	if isPrivate && contextdata.IsUserAnonymous(ctx) {
		return api.Error(http.StatusInternalServerError, org.ErrOrgNotFound) // TODO: change sttaus code
	}
	if isPrivate {
		user := contextdata.GetUser(ctx)
		// check permissions
		query := &accesscontrol.EvaluateFilterQuery{
			UserId: user.GetId(),
			OrgId:  query.OrgId,
		}
		if err := s.aclService.Evaluate(ctx, accesscontrol.ActionOrgUserRead, query); err != nil {
			if err == accesscontrol.ErrPermissionDenied {
				// unauthorized user (permission denied) => return org not found error
				return api.Error(http.StatusInternalServerError, org.ErrOrgNotFound) // TODO: change status code
			}

			return api.Error(http.StatusInternalServerError, err)
		}
	}

	userList, err := s.orgService.GetUserList(ctx, query)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	// anonymous and unauthorized user => return users of public org
	// authorized user => return users of public or private org
	return api.Success(http.StatusOK, &userList)
}
