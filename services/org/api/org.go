package org

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/quarkloop/quarkloop/pkg/grpc/v1/system/org"

	userGrpc "github.com/quarkloop/quarkloop/pkg/grpc/v1/auth/user"
	grpc "github.com/quarkloop/quarkloop/pkg/grpc/v1/system/org"
	accesscontrol "github.com/quarkloop/quarkloop/services/accesscontrol/service"
	orgErrors "github.com/quarkloop/quarkloop/services/org/errors"
	quota "github.com/quarkloop/quarkloop/services/quota/service"
)

type Api interface {
	// query
	GetOrgById(*gin.Context)
	GetOrgList(*gin.Context)
	GetWorkspaceList(*gin.Context)
	GetMemberList(*gin.Context)

	// mutation
	CreateOrg(*gin.Context)
	UpdateOrgById(*gin.Context)
	DeleteOrgById(*gin.Context)
	ChangeOrgVisibility(*gin.Context)
}

type orgApi struct {
	orgService   grpc.OrgServiceClient
	userService  userGrpc.UserServiceClient
	aclService   accesscontrol.Service
	quotaService quota.Service
}

func NewOrgApi(orgService grpc.OrgServiceClient, userService userGrpc.UserServiceClient, aclService accesscontrol.Service, quotaService quota.Service) *orgApi {
	return &orgApi{
		orgService:   orgService,
		userService:  userService,
		aclService:   aclService,
		quotaService: quotaService,
	}
}

func GetOrgId(ctx context.Context, service org.OrgServiceClient, orgSid string) (orgId int64, err error, errStatus int) {
	reply, err := service.GetOrgId(ctx, &org.GetOrgIdQuery{OrgSid: orgSid})
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				err = orgErrors.ErrOrgNotFound
				errStatus = http.StatusNotFound
				return
			case codes.AlreadyExists:
				err = orgErrors.ErrOrgAlreadyExists
				errStatus = http.StatusConflict
				return
			case codes.Internal:
				err = errors.New(e.Message())
				errStatus = http.StatusInternalServerError
				return
			case codes.InvalidArgument:
				err = errors.New(e.Message())
				errStatus = http.StatusBadRequest
				return
			}
		}
		errStatus = http.StatusInternalServerError
		return
	}

	orgId = reply.OrgId
	return
}

// func (api *orgApi) GetService() org.Service {
// 	return api.orgService
// }

// func transformGrpcSlice(slice []*system.Org) []*system.Org {
// 	if slice == nil {
// 		return []*system.Org{}
// 	}
// 	return slice
// }

// func (s *orgApi) getOrgId(ctx *gin.Context, orgSid string) int64 {
// 	return 0
// }
