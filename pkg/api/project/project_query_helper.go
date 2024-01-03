package project

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/contextdata"
	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/accesscontrol"
	"github.com/quarkloop/quarkloop/pkg/service/project"
	"github.com/quarkloop/quarkloop/pkg/service/user"
)

func (s *ProjectApi) getProjectById(ctx *gin.Context, query *project.GetProjectByIdQuery) api.Response {
	p, err := s.projectService.GetProjectById(ctx, query)
	if err != nil {
		if err == project.ErrProjectNotFound {
			return api.Error(http.StatusNoContent, nil)
		}
		return api.Error(http.StatusInternalServerError, err)
	}

	isPrivate := p.Visibility == model.PrivateVisibility

	// anonymous user => return project not found error
	if isPrivate && contextdata.IsUserAnonymous(ctx) {
		return api.Error(http.StatusNotFound, project.ErrProjectNotFound)
	}
	if isPrivate {
		user := contextdata.GetUser(ctx)
		// check permissions
		query := &accesscontrol.EvaluateFilterQuery{
			UserId:      user.GetId(),
			OrgId:       query.OrgId,
			WorkspaceId: query.WorkspaceId,
			ProjectId:   query.ProjectId,
		}
		if err := s.aclService.Evaluate(ctx, accesscontrol.ActionProjectRead, query); err != nil {
			if err == accesscontrol.ErrPermissionDenied {
				// unauthorized user (permission denied) => return project not found error
				return api.Error(http.StatusNotFound, project.ErrProjectNotFound) // TODO: change status code
			}

			return api.Error(http.StatusInternalServerError, err)
		}
	}

	// anonymous and unauthorized user => return public project
	// authorized user => return public or private project
	return api.Success(http.StatusOK, p)
}

func (s *ProjectApi) getProjectList(ctx *gin.Context) api.Response {
	if contextdata.IsUserAnonymous(ctx) {
		// anonymous user => return public projects
		projectList, err := s.projectService.GetProjectList(ctx, &project.GetProjectListQuery{Visibility: model.PublicVisibility})
		if err != nil {
			return api.Error(http.StatusInternalServerError, err) // TODO: change status code
		}

		return api.Success(http.StatusOK, &projectList)
	}

	user := contextdata.GetUser(ctx)

	// check permissions
	err := s.aclService.Evaluate(ctx, accesscontrol.ActionProjectRead, &accesscontrol.EvaluateFilterQuery{UserId: user.GetId()})
	if err != nil {
		if err == accesscontrol.ErrPermissionDenied {
			// unauthorized user (permission denied) => return public projects
			projectList, err := s.projectService.GetProjectList(ctx, &project.GetProjectListQuery{
				UserId:     user.GetId(),
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
	projectList, err := s.projectService.GetProjectList(ctx, &project.GetProjectListQuery{
		UserId:     user.GetId(),
		Visibility: model.AllVisibility,
	})
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	return api.Success(http.StatusOK, &projectList)
}

func (s *ProjectApi) getMemberList(ctx *gin.Context, query *project.GetMemberListQuery) api.Response {
	ws, err := s.projectService.GetProjectById(ctx, &project.GetProjectByIdQuery{
		OrgId:       query.OrgId,
		WorkspaceId: query.WorkspaceId,
		ProjectId:   query.ProjectId,
	})
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	isPrivate := ws.Visibility == model.PrivateVisibility

	// anonymous user => return project not found error
	if isPrivate && contextdata.IsUserAnonymous(ctx) {
		return api.Error(http.StatusInternalServerError, project.ErrProjectNotFound) // TODO: change sttaus code
	}
	if isPrivate {
		user := contextdata.GetUser(ctx)
		// check permissions
		query := &accesscontrol.EvaluateFilterQuery{
			UserId:      user.GetId(),
			OrgId:       query.OrgId,
			WorkspaceId: query.WorkspaceId,
			ProjectId:   query.ProjectId,
		}
		if err := s.aclService.Evaluate(ctx, accesscontrol.ActionWorkspaceUserRead, query); err != nil {
			if err == accesscontrol.ErrPermissionDenied {
				// unauthorized user (permission denied) => return project not found error
				return api.Error(http.StatusInternalServerError, project.ErrProjectNotFound) // TODO: change status code
			}

			return api.Error(http.StatusInternalServerError, err)
		}
	}

	uaList, err := s.projectService.GetUserAssignmentList(ctx, &project.GetUserAssignmentListQuery{
		OrgId:       query.OrgId,
		WorkspaceId: query.WorkspaceId,
		ProjectId:   query.ProjectId,
	})
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
