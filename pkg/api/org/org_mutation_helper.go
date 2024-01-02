package org

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/contextdata"
	"github.com/quarkloop/quarkloop/pkg/service/accesscontrol"
	"github.com/quarkloop/quarkloop/pkg/service/org"
	"github.com/quarkloop/quarkloop/pkg/service/quota"
)

func (s *OrgApi) createOrg(ctx *gin.Context, cmd *org.CreateOrgCommand) api.Reponse {
	// check permissions
	if err := s.evaluateCreatePermission(ctx, accesscontrol.ActionOrgCreate); err != nil {
		return api.Error(http.StatusInternalServerError, err) // TODO: change status
	}

	// check quotas
	user := contextdata.GetUser(ctx)
	quotaQuery := &quota.CheckCreateOrgQuotaQuery{UserId: user.GetId()}
	if err := s.quotaService.CheckCreateOrgQuota(ctx, quotaQuery); err != nil {
		return api.Error(http.StatusInternalServerError, err) // TODO: change status
	}

	org, err := s.orgService.CreateOrg(ctx, cmd)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	return api.Success(http.StatusCreated, org)
}

func (s *OrgApi) updateOrgById(ctx *gin.Context, cmd *org.UpdateOrgByIdCommand) api.Reponse {
	// check permissions
	if err := s.evaluatePermission(ctx, accesscontrol.ActionProjectUpdate, cmd.OrgId); err != nil {
		return api.Error(http.StatusInternalServerError, err) // TODO: change status
	}

	err := s.orgService.UpdateOrgById(ctx, cmd)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	return api.Success(http.StatusOK, nil)
}

func (s *OrgApi) deleteOrgById(ctx *gin.Context, orgId int) api.Reponse {
	// check permissions
	if err := s.evaluatePermission(ctx, accesscontrol.ActionProjectDelete, orgId); err != nil {
		return api.Error(http.StatusInternalServerError, err) // TODO: change status
	}

	err := s.orgService.DeleteOrgById(ctx, &org.DeleteOrgByIdCommand{OrgId: orgId})
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	return api.Success(http.StatusNoContent, nil)
}

func (s *OrgApi) evaluateCreatePermission(ctx *gin.Context, permission string) error {
	user := contextdata.GetUser(ctx)
	query := &accesscontrol.EvaluateFilterQuery{UserId: user.GetId()}

	return s.aclService.Evaluate(ctx, permission, query)
}

func (s *OrgApi) evaluatePermission(ctx *gin.Context, permission string, orgId int) error {
	user := contextdata.GetUser(ctx)
	query := &accesscontrol.EvaluateFilterQuery{
		UserId: user.GetId(),
		OrgId:  orgId,
	}

	return s.aclService.Evaluate(ctx, permission, query)
}
