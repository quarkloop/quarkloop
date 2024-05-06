package org

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/contextdata"
	"github.com/quarkloop/quarkloop/pkg/model"

	grpc "github.com/quarkloop/quarkloop/pkg/grpc/v1/system/org"
	permission "github.com/quarkloop/quarkloop/services/accesscontrol/permission"
	accesscontrol "github.com/quarkloop/quarkloop/services/accesscontrol/service"
	orgErrors "github.com/quarkloop/quarkloop/services/org/errors"
	quota "github.com/quarkloop/quarkloop/services/quota/service"
)

func (s *orgApi) createOrg(ctx *gin.Context, cmd *CreateOrgCommand) api.Response {
	user := contextdata.GetUser(ctx)

	// check quotas
	quotaQuery := &quota.CheckCreateOrgQuotaQuery{UserId: user.GetId()}
	if err := s.quotaService.CheckCreateOrgQuota(ctx, quotaQuery); err != nil {
		return api.Error(http.StatusTooManyRequests, err)
	}

	// create org
	reply, err := s.orgService.CreateOrg(ctx, &grpc.CreateOrgCommand{
		CreatedBy: "admin",
		Payload:   cmd.Payload.ToProto(),
	})
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

	// grant org access to user
	err = s.aclService.GrantUserAccess(ctx, &accesscontrol.GrantUserAccessCommand{
		UserId: int64(user.Id),
		OrgId:  reply.Data.Id,
		Role:   string(permission.RoleOwner),
	})
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	return api.Success(http.StatusCreated, &CreateOrgDTO{Data: model.FromOrgProto(reply.GetData())})
}

func (s *orgApi) updateOrgById(ctx *gin.Context, cmd *UpdateOrgByIdCommand) api.Response {
	// check permissions
	access, err := s.evaluatePermission(ctx, permission.ActionOrgUpdate, cmd.OrgId)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err) // TODO: change status
	}
	if !access {
		// unauthorized user (permission denied) => return org not found error
		return api.Error(http.StatusNotFound, orgErrors.ErrOrgNotFound) // TODO: change status code
	}

	_, err = s.orgService.UpdateOrgById(ctx, &grpc.UpdateOrgByIdCommand{
		UpdatedBy: "admin",
		OrgId:     cmd.OrgId,
		Payload:   cmd.Payload.ToProto(),
	})
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	return api.Success(http.StatusOK, nil)
}

func (s *orgApi) deleteOrgById(ctx *gin.Context, cmd *DeleteOrgByIdCommand) api.Response {
	// check permissions
	access, err := s.evaluatePermission(ctx, permission.ActionOrgDelete, cmd.OrgId)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err) // TODO: change status
	}
	if !access {
		// unauthorized user (permission denied) => return org not found error
		return api.Error(http.StatusNotFound, orgErrors.ErrOrgNotFound) // TODO: change status code
	}

	_, err = s.orgService.DeleteOrgById(ctx, &grpc.DeleteOrgByIdCommand{OrgId: cmd.OrgId})
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	err = s.aclService.RevokeUserAccess(ctx, &accesscontrol.RevokeUserAccessCommand{OrgId: cmd.OrgId})
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	return api.Success(http.StatusNoContent, nil)
}

func (s *orgApi) evaluatePermission(ctx *gin.Context, permission string, orgId int64) (bool, error) {
	user := contextdata.GetUser(ctx)
	query := &accesscontrol.CheckPermissionQuery{
		Permission: permission,
		UserId:     user.GetId(),
		OrgId:      orgId,
	}

	return s.aclService.CheckPermission(ctx, query)
}

func (s *orgApi) changeOrgVisibility(ctx *gin.Context, cmd *ChangeOrgVisibilityCommand) api.Response {
	// check permissions
	access, err := s.evaluatePermission(ctx, permission.ActionOrgUpdate, cmd.OrgId)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err) // TODO: change status
	}
	if !access {
		// unauthorized user (permission denied) => return org not found error
		return api.Error(http.StatusNotFound, orgErrors.ErrOrgNotFound) // TODO: change status code
	}

	_, err = s.orgService.ChangeOrgVisibility(ctx, &grpc.ChangeOrgVisibilityCommand{
		UpdatedBy:  "admin",
		OrgId:      cmd.OrgId,
		Visibility: string(cmd.Visibility),
	})
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	return api.Success(http.StatusOK, nil)
}
