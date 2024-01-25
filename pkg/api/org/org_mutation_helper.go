package org

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/contextdata"
	"github.com/quarkloop/quarkloop/pkg/service/accesscontrol"
	"github.com/quarkloop/quarkloop/pkg/service/org"
	"github.com/quarkloop/quarkloop/pkg/service/quota"
	grpc "github.com/quarkloop/quarkloop/service/v1/system/org"
)

func (s *orgApi) createOrg(ctx *gin.Context, cmd *org.CreateOrgCommand) api.Response {
	// check quotas
	user := contextdata.GetUser(ctx)
	quotaQuery := &quota.CheckCreateOrgQuotaQuery{UserId: user.GetId()}
	if err := s.quotaService.CheckCreateOrgQuota(ctx, quotaQuery); err != nil {
		fmt.Printf("\nOrgQuota => %+v\n\n", err)
		return api.Error(http.StatusTooManyRequests, err) // TODO: change status
	}

	org, err := s.orgService.CreateOrg(ctx, &grpc.CreateOrgCommand{
		CreatedBy:   cmd.CreatedBy,
		ScopeId:     cmd.ScopeId,
		Name:        cmd.Name,
		Description: cmd.Description,
		Visibility:  int32(cmd.Visibility),
	})
	fmt.Printf("\nCreateOrg => %+v => %+v\n\n", err, org)
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.AlreadyExists:
				return api.Error(http.StatusConflict, err)
			case codes.Internal:
				return api.Error(http.StatusInternalServerError, err)
			case codes.InvalidArgument:
				return api.Error(http.StatusBadRequest, err)
			}
		}
		return api.Error(http.StatusInternalServerError, err)
	}

	aclCommand := &accesscontrol.GrantUserAccessCommand{
		UserId: user.Id,
		OrgId:  org.Org.Id,
		Role:   accesscontrol.RoleOwner,
	}
	err = s.aclService.GrantUserAccess(ctx, aclCommand)
	fmt.Printf("\nGrantUserAccess => %+v \n\n", err)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	return api.Success(http.StatusCreated, org.GetOrg())
}

func (s *orgApi) updateOrgById(ctx *gin.Context, cmd *org.UpdateOrgByIdCommand) api.Response {
	// check permissions
	access, err := s.evaluatePermission(ctx, accesscontrol.ActionOrgUpdate, cmd.OrgId)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err) // TODO: change status
	}
	if !access {
		// unauthorized user (permission denied) => return org not found error
		return api.Error(http.StatusNotFound, org.ErrOrgNotFound) // TODO: change status code
	}

	_, err = s.orgService.UpdateOrgById(ctx, &grpc.UpdateOrgByIdCommand{
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
	user := contextdata.GetUser(ctx)

	// check permissions
	access, err := s.evaluatePermission(ctx, accesscontrol.ActionOrgDelete, cmd.OrgId)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err) // TODO: change status
	}
	if !access {
		// unauthorized user (permission denied) => return org not found error
		return api.Error(http.StatusNotFound, org.ErrOrgNotFound) // TODO: change status code
	}
	fmt.Printf("\nRevokeUserAccess (evaluatePermission) => %+v => %+v \n\n", access, err)

	_, err = s.orgService.DeleteOrgById(ctx, &grpc.DeleteOrgByIdCommand{OrgId: cmd.OrgId})
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	aclCommand := &accesscontrol.RevokeUserAccessCommand{
		UserId: user.Id,
		OrgId:  cmd.OrgId,
		Role:   accesscontrol.RoleOwner,
	}
	err = s.aclService.RevokeUserAccess(ctx, aclCommand)
	fmt.Printf("\nRevokeUserAccess => %+v \n\n", err)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	return api.Success(http.StatusNoContent, nil)
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
