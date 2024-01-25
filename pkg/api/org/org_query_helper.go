package org

import (
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"

	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/contextdata"
	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/accesscontrol"
	"github.com/quarkloop/quarkloop/pkg/service/org"
	"github.com/quarkloop/quarkloop/pkg/service/user"
	"github.com/quarkloop/quarkloop/service/v1/system"
	grpc "github.com/quarkloop/quarkloop/service/v1/system/org"
)

func (s *orgApi) getOrgById(ctx *gin.Context, query *org.GetOrgByIdQuery) api.Response {
	visibility, err := s.orgService.GetOrgVisibilityById(ctx, &grpc.GetOrgVisibilityByIdQuery{OrgId: query.OrgId})
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
		access, err := s.aclService.EvaluateUserAccess(ctx, query)
		if err != nil {
			return api.Error(http.StatusInternalServerError, err)
		}
		if !access {
			// unauthorized user (permission denied) => return org not found error
			return api.Error(http.StatusNotFound, org.ErrOrgNotFound) // TODO: change status code
		}
	}

	o, err := s.orgService.GetOrgById(ctx, &grpc.GetOrgByIdQuery{OrgId: query.OrgId})
	if err != nil {
		if err == org.ErrOrgNotFound {
			return api.Error(http.StatusNotFound, err)
		}
		return api.Error(http.StatusInternalServerError, err)
	}

	// anonymous and unauthorized user => return public org
	// authorized user => return public or private org
	return api.Success(http.StatusOK, o.GetOrg())
}

func (s *orgApi) getOrgList(ctx *gin.Context) api.Response {
	query := &grpc.GetOrgListQuery{Visibility: int32(model.PublicVisibility)}
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
	return api.Success(http.StatusOK, transformGrpcSlice(orgList.OrgList))
}

func (s *orgApi) getWorkspaceList(ctx *gin.Context, query *org.GetWorkspaceListQuery) api.Response {
	var authorizedList []int32 = []int32{}
	queryPayload := &grpc.GetWorkspaceListQuery{
		OrgId:      query.OrgId,
		Visibility: int32(model.PublicVisibility),
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

		if len(list) > 0 {
			authorizedList = list
			queryPayload.Visibility = int32(model.AllVisibility)
		}
	}

	var newList []*system.Workspace = []*system.Workspace{}
	{
		wsList, err := s.orgService.GetWorkspaceList(ctx, queryPayload)
		if err != nil {
			return api.Error(http.StatusInternalServerError, err)
		}

		for _, ws := range wsList.GetWorkspaceList() {
			if ws.Visibility == int32(model.PublicVisibility) {
				newList = append(newList, ws)
			} else if ws.Visibility == int32(model.PrivateVisibility) && slices.Contains(authorizedList, ws.Id) {
				newList = append(newList, ws)
			}
		}
	}

	// anonymous user => return public workspaces
	// unauthorized user (permission denied) => return public workspaces
	// authorized user => return public + private workspaces
	return api.Success(http.StatusOK, newList)
}

func (s *orgApi) getProjectList(ctx *gin.Context, query *org.GetProjectListQuery) api.Response {
	var authorizedList []int32 = []int32{}
	queryPayload := &grpc.GetProjectListQuery{
		OrgId:      query.OrgId,
		Visibility: int32(model.PublicVisibility),
	}
	if !contextdata.IsUserAnonymous(ctx) {
		// check permissions
		user := contextdata.GetUser(ctx)
		aclQuery := &accesscontrol.GetProjectListQuery{
			Permission: "all",
			UserId:     user.GetId(),
		}
		list, err := s.aclService.GetProjectList(ctx, aclQuery)
		if err != nil {
			return api.Error(http.StatusInternalServerError, err)
		}

		if len(list) > 0 {
			authorizedList = list
			queryPayload.Visibility = int32(model.AllVisibility)
		}
	}

	var newList []*system.Project = []*system.Project{}
	{
		projectList, err := s.orgService.GetProjectList(ctx, queryPayload)
		if err != nil {
			return api.Error(http.StatusInternalServerError, err)
		}

		for _, project := range projectList.GetProjectList() {
			if project.Visibility == int32(model.PublicVisibility) {
				newList = append(newList, project)
			} else if project.Visibility == int32(model.PrivateVisibility) && slices.Contains(authorizedList, project.Id) {
				newList = append(newList, project)
			}
		}
	}

	// anonymous user => return public projects
	// unauthorized user (permission denied) => return public projects
	// authorized user => return public + private projects
	return api.Success(http.StatusOK, newList)
}

func (s *orgApi) getMemberList(ctx *gin.Context, query *org.GetMemberListQuery) api.Response {
	visibility, err := s.orgService.GetOrgVisibilityById(ctx, &grpc.GetOrgVisibilityByIdQuery{OrgId: query.OrgId})
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
		access, err := s.aclService.EvaluateUserAccess(ctx, evalQuery)
		if err != nil {
			return api.Error(http.StatusInternalServerError, err)
		}
		if !access {
			// unauthorized user (permission denied) => return org not found error
			return api.Error(http.StatusNotFound, org.ErrOrgNotFound) // TODO: change status code
		}
	}

	uaList, err := s.aclService.GetOrgMemberList(ctx, &accesscontrol.GetOrgMemberListQuery{OrgId: query.OrgId})
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
