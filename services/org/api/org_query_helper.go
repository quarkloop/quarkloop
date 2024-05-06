package org

import (
	"fmt"
	"log"
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"

	userGrpc "github.com/quarkloop/quarkloop/pkg/grpc/v1/auth/user"
	grpc "github.com/quarkloop/quarkloop/pkg/grpc/v1/system/org"
	permission "github.com/quarkloop/quarkloop/services/accesscontrol/permission"
	accesscontrol "github.com/quarkloop/quarkloop/services/accesscontrol/service"
	orgErrors "github.com/quarkloop/quarkloop/services/org/errors"

	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/contextdata"
	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/services/user"
)

func (s *orgApi) getOrgById(ctx *gin.Context, query *GetOrgByIdQuery) api.Response {
	visibility, err := s.orgService.GetOrgVisibilityById(ctx, &grpc.GetOrgVisibilityByIdQuery{OrgId: query.OrgId})
	if err != nil {
		if err == orgErrors.ErrOrgNotFound {
			return api.Error(http.StatusNotFound, err)
		}
		return api.Error(http.StatusInternalServerError, err)
	}

	isPrivate := model.ScopeVisibility(visibility.GetVisibility()) == model.PrivateVisibility

	// anonymous user => return org not found error
	if isPrivate && contextdata.IsUserAnonymous(ctx) {
		return api.Error(http.StatusNotFound, orgErrors.ErrOrgNotFound)
	}
	if isPrivate {
		// check permissions
		user := contextdata.GetUser(ctx)
		query := &accesscontrol.CheckPermissionQuery{
			Permission: permission.ActionOrgRead,
			UserId:     user.GetId(),
			OrgId:      query.OrgId,
		}
		access, err := s.aclService.CheckPermission(ctx, query)
		if err != nil {
			return api.Error(http.StatusInternalServerError, err)
		}
		if !access {
			// unauthorized user (permission denied) => return org not found error
			return api.Error(http.StatusNotFound, orgErrors.ErrOrgNotFound) // TODO: change status code
		}
	}

	reply, err := s.orgService.GetOrgById(ctx, &grpc.GetOrgByIdQuery{OrgId: query.OrgId})
	if err != nil {
		if err == orgErrors.ErrOrgNotFound {
			return api.Error(http.StatusNotFound, err)
		}
		return api.Error(http.StatusInternalServerError, err)
	}

	// anonymous and unauthorized user => return public org
	// authorized user => return public or private org
	return api.Success(http.StatusOK, &GetOrgByIdDTO{Data: model.FromOrgProto(reply.Data)})
}

func (s *orgApi) getOrgList(ctx *gin.Context) api.Response {
	query := &grpc.GetOrgListQuery{Visibility: model.PublicVisibility.ToString()}
	if !contextdata.IsUserAnonymous(ctx) {
		// check permissions
		user := contextdata.GetUser(ctx)
		aclQuery := &accesscontrol.GetOrgListQuery{
			Permission: "all",
			UserId:     user.GetId(),
		}
		list, err := s.aclService.GetOrgList(ctx, aclQuery)
		if err != nil {
			return api.Error(http.StatusInternalServerError, err)
		}

		if len(list) > 0 {
			query.OrgIdList = list
			query.Visibility = model.AllVisibility.ToString()
		}
	}

	reply, err := s.orgService.GetOrgList(ctx, query)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	orgList := []*model.Org{}
	for _, data := range reply.Data {
		orgList = append(orgList, model.FromOrgProto(data))
	}
	// anonymous user => return public orgs
	// unauthorized user (permission denied) => return public orgs
	// authorized user => return public + private orgs
	return api.Success(http.StatusOK, &GetOrgListDTO{Data: orgList})
}

func (s *orgApi) getWorkspaceList(ctx *gin.Context, query *GetWorkspaceListQuery) api.Response {
	var authorizedList []int64 = []int64{}
	queryPayload := &grpc.GetWorkspaceListQuery{
		OrgId:      query.OrgId,
		Visibility: model.PublicVisibility.ToString(),
	}

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
		fmt.Printf("\n\n===> %v\n\n", list)

		if len(list) > 0 {
			authorizedList = list
			queryPayload.Visibility = model.AllVisibility.ToString()
		}
	}

	var workspaceList []*model.Workspace = []*model.Workspace{}
	{
		reply, err := s.orgService.GetWorkspaceList(ctx, queryPayload)
		if err != nil {
			return api.Error(http.StatusInternalServerError, err)
		}

		for _, ws := range reply.GetData() {
			if ws.Visibility == model.PublicVisibility.ToString() {
				workspaceList = append(workspaceList, model.FromWorkspaceProto(ws))
			} else if ws.Visibility == model.PrivateVisibility.ToString() && slices.Contains(authorizedList, ws.Id) {
				workspaceList = append(workspaceList, model.FromWorkspaceProto(ws))
			}
		}
	}

	// anonymous user => return public workspaces
	// unauthorized user (permission denied) => return public workspaces
	// authorized user => return public + private workspaces
	return api.Success(http.StatusOK, &GetWorkspaceListDTO{Data: workspaceList})
}

func (s *orgApi) getMemberList(ctx *gin.Context, query *GetMemberListQuery) api.Response {
	visibility, err := s.orgService.GetOrgVisibilityById(ctx, &grpc.GetOrgVisibilityByIdQuery{OrgId: query.OrgId})
	if err != nil {
		if err == orgErrors.ErrOrgNotFound {
			return api.Error(http.StatusNotFound, err)
		}
		return api.Error(http.StatusInternalServerError, err)
	}

	isPrivate := model.ScopeVisibility(visibility.GetVisibility()) == model.PrivateVisibility
	// anonymous user => return org not found error
	if isPrivate && contextdata.IsUserAnonymous(ctx) {
		return api.Error(http.StatusNotFound, orgErrors.ErrOrgNotFound) // TODO: change sttaus code
	}
	if isPrivate {
		// check permissions
		user := contextdata.GetUser(ctx)
		evalQuery := &accesscontrol.CheckPermissionQuery{
			Permission: permission.ActionOrgMemberRead,
			UserId:     user.GetId(),
			OrgId:      query.OrgId,
		}
		access, err := s.aclService.CheckPermission(ctx, evalQuery)
		if err != nil {
			return api.Error(http.StatusInternalServerError, err)
		}
		if !access {
			// unauthorized user (permission denied) => return org not found error
			return api.Error(http.StatusNotFound, orgErrors.ErrOrgNotFound) // TODO: change status code
		}
	}

	uaList, err := s.aclService.GetOrgMemberList(ctx, &accesscontrol.GetOrgMemberListQuery{OrgId: query.OrgId})
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	var memberList []*user.MemberDTO = []*user.MemberDTO{}
	for i := range uaList {
		ua := uaList[i]
		u, err := s.userService.GetUserById(ctx, &userGrpc.GetUserByIdQuery{UserId: ua.UserId})
		if err != nil {
			log.Printf("[GetOrgMemberList] error: %v\n", err)
			continue
		}

		user := user.MemberDTO{
			User: &model.User{
				Id:       u.User.Id,
				Username: &u.User.Username,
				Name:     u.User.Name,
				Email:    u.User.Email,
				Status:   u.User.Status,
			},
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
