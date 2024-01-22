package org

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/contextdata"
	"github.com/quarkloop/quarkloop/pkg/service/accesscontrol"
	"github.com/quarkloop/quarkloop/pkg/service/org"
	"github.com/quarkloop/quarkloop/pkg/service/quota"
	"github.com/quarkloop/quarkloop/service/system"
)

func (s *orgApi) createOrg(ctx *gin.Context, cmd *org.CreateOrgCommand) api.Response {
	// check permissions
	access, err := s.evaluateCreatePermission(ctx, accesscontrol.ActionOrgCreate)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err) // TODO: change status
	}
	if !access {
		// unauthorized user (permission denied) => return org not found error
		return api.Error(http.StatusNotFound, org.ErrOrgNotFound) // TODO: change status code and error
	}

	// check quotas
	user := contextdata.GetUser(ctx)
	quotaQuery := &quota.CheckCreateOrgQuotaQuery{UserId: user.GetId()}
	if err := s.quotaService.CheckCreateOrgQuota(ctx, quotaQuery); err != nil {
		return api.Error(http.StatusInternalServerError, err) // TODO: change status
	}

	org, err := s.orgService.CreateOrg(ctx, &system.CreateOrgCommand{
		CreatedBy:   cmd.CreatedBy,
		ScopeId:     cmd.ScopeId,
		Name:        cmd.Name,
		Description: cmd.Description,
		Visibility:  int32(cmd.Visibility),
	})
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	return api.Success(http.StatusCreated, org)
}

func (s *orgApi) updateOrgById(ctx *gin.Context, cmd *org.UpdateOrgByIdCommand) api.Response {
	// check permissions
	access, err := s.evaluatePermission(ctx, accesscontrol.ActionProjectUpdate, cmd.OrgId)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err) // TODO: change status
	}
	if !access {
		// unauthorized user (permission denied) => return org not found error
		return api.Error(http.StatusNotFound, org.ErrOrgNotFound) // TODO: change status code
	}

	_, err = s.orgService.UpdateOrgById(ctx, &system.UpdateOrgByIdCommand{
		OrgId:       cmd.OrgId,
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

func (s *orgApi) deleteOrgById(ctx *gin.Context, cmd *org.DeleteOrgByIdCommand) api.Response {
	// check permissions
	access, err := s.evaluatePermission(ctx, accesscontrol.ActionProjectDelete, cmd.OrgId)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err) // TODO: change status
	}
	if !access {
		// unauthorized user (permission denied) => return org not found error
		return api.Error(http.StatusNotFound, org.ErrOrgNotFound) // TODO: change status code
	}

	_, err = s.orgService.DeleteOrgById(ctx, &system.DeleteOrgByIdCommand{OrgId: cmd.OrgId})
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	return api.Success(http.StatusNoContent, nil)
}

func (s *orgApi) evaluateCreatePermission(ctx *gin.Context, permission string) (bool, error) {
	user := contextdata.GetUser(ctx)
	query := &accesscontrol.EvaluateQuery{
		Permission: permission,
		UserId:     user.GetId(),
	}

	return s.aclService.EvaluateUserAccess(ctx, query)
}

func (s *orgApi) evaluatePermission(ctx *gin.Context, permission string, orgId int32) (bool, error) {
	user := contextdata.GetUser(ctx)
	query := &accesscontrol.EvaluateQuery{
		Permission: permission,
		UserId:     user.GetId(),
		OrgId:      orgId,
	}

	return s.aclService.EvaluateUserAccess(ctx, query)
}
