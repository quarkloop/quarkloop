package project

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/contextdata"
	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/accesscontrol"
	"github.com/quarkloop/quarkloop/pkg/service/project"
	"github.com/quarkloop/quarkloop/pkg/service/user"
	grpc "github.com/quarkloop/quarkloop/service/v1/system/project"
)

func (s *ProjectApi) getProjectById(ctx *gin.Context, query *project.GetProjectByIdQuery) api.Response {
	visibility, err := s.projectService.GetProjectVisibilityById(ctx, &grpc.GetProjectVisibilityByIdQuery{
		OrgId:       query.OrgId,
		WorkspaceId: query.WorkspaceId,
		ProjectId:   query.ProjectId,
	})
	if err != nil {
		if err == project.ErrProjectNotFound {
			return api.Error(http.StatusNotFound, err)
		}
		return api.Error(http.StatusInternalServerError, err)
	}

	isPrivate := model.ScopeVisibility(visibility.GetVisibility()) == model.PrivateVisibility

	// anonymous user => return project not found error
	if isPrivate && contextdata.IsUserAnonymous(ctx) {
		return api.Error(http.StatusNotFound, project.ErrProjectNotFound)
	}
	if isPrivate {
		user := contextdata.GetUser(ctx)
		// check permissions
		evalQuery := &accesscontrol.EvaluateQuery{
			Permission:  accesscontrol.ActionProjectRead,
			UserId:      user.GetId(),
			OrgId:       query.OrgId,
			WorkspaceId: query.WorkspaceId,
			ProjectId:   query.ProjectId,
		}
		access, err := s.aclService.EvaluateUserAccess(ctx, evalQuery)
		if err != nil {
			return api.Error(http.StatusInternalServerError, err)
		}
		if !access {
			// unauthorized user (permission denied) => return project not found error
			return api.Error(http.StatusNotFound, project.ErrProjectNotFound) // TODO: change status code
		}
	}

	p, err := s.projectService.GetProjectById(ctx, &grpc.GetProjectByIdQuery{
		OrgId:       query.OrgId,
		WorkspaceId: query.WorkspaceId,
		ProjectId:   query.ProjectId,
	})
	if err != nil {
		if err == project.ErrProjectNotFound {
			return api.Error(http.StatusNotFound, err)
		}
		return api.Error(http.StatusInternalServerError, err)
	}

	// anonymous and unauthorized user => return public project
	// authorized user => return public or private project
	return api.Success(http.StatusOK, p.GetProject())
}

func (s *ProjectApi) getProjectList(ctx *gin.Context) api.Response {
	query := &grpc.GetProjectListQuery{Visibility: int32(model.PublicVisibility)}
	if !contextdata.IsUserAnonymous(ctx) {
		// check permissions
		user := contextdata.GetUser(ctx)
		aclQuery := &accesscontrol.GetProjectListQuery{
			Permission: "all",
			UserId:     user.GetId(),
		}
		list, err := s.aclService.GetProjectList(ctx, aclQuery)
		fmt.Printf("\n[aclService] GetProjectList => %+v => %+v\n", list, err)
		if err != nil {
			return api.Error(http.StatusInternalServerError, err)
		}

		if len(list) > 0 {
			query.ProjectIdList = list
			query.Visibility = int32(model.AllVisibility)
		}
	}

	projectList, err := s.projectService.GetProjectList(ctx, query)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	// anonymous user => return public projects
	// unauthorized user (permission denied) => return public projects
	// authorized user => return public + private projects
	return api.Success(http.StatusOK, transformGrpcSlice(projectList.ProjectList))
}

func (s *ProjectApi) getMemberList(ctx *gin.Context, query *project.GetMemberListQuery) api.Response {
	visibility, err := s.projectService.GetProjectVisibilityById(ctx, &grpc.GetProjectVisibilityByIdQuery{
		OrgId:       query.OrgId,
		WorkspaceId: query.WorkspaceId,
		ProjectId:   query.ProjectId,
	})
	if err != nil {
		if err == project.ErrProjectNotFound {
			return api.Error(http.StatusNotFound, err)
		}
		return api.Error(http.StatusInternalServerError, err)
	}

	isPrivate := model.ScopeVisibility(visibility.GetVisibility()) == model.PrivateVisibility

	// anonymous user => return project not found error
	if isPrivate && contextdata.IsUserAnonymous(ctx) {
		return api.Error(http.StatusNotFound, project.ErrProjectNotFound) // TODO: change sttaus code
	}
	if isPrivate {
		user := contextdata.GetUser(ctx)
		// check permissions
		evalQuery := &accesscontrol.EvaluateQuery{
			Permission:  accesscontrol.ActionWorkspaceUserRead,
			UserId:      user.GetId(),
			OrgId:       query.OrgId,
			WorkspaceId: query.WorkspaceId,
			ProjectId:   query.ProjectId,
		}
		access, err := s.aclService.EvaluateUserAccess(ctx, evalQuery)
		if err != nil {
			return api.Error(http.StatusInternalServerError, err)
		}
		if !access {
			// unauthorized user (permission denied) => return project not found error
			return api.Error(http.StatusNotFound, project.ErrProjectNotFound) // TODO: change status code
		}
	}

	uaList, err := s.aclService.GetProjectMemberList(ctx, &accesscontrol.GetProjectMemberListQuery{ProjectId: query.ProjectId})
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}
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

	// anonymous and unauthorized user => return members of public project
	// authorized user => return members of public or private project
	return api.Success(http.StatusOK, &memberList)
}
