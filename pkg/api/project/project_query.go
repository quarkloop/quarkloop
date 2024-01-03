package project

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/service/project"
)

// GET /orgs/:orgId/workspaces/:workspaceId/projects/:projectId
//
// Get project by id.
//
// Response status:
// 200: StatusOK
// 400: StatusBadRequest
// 500: StatusInternalServerError

func (s *ProjectApi) GetProjectById(ctx *gin.Context) {
	uriParams := &project.GetProjectByIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	query := &project.GetProjectByIdQuery{
		OrgId:       uriParams.OrgId,
		WorkspaceId: uriParams.WorkspaceId,
		ProjectId:   uriParams.ProjectId,
	}
	res := s.getProjectById(ctx, query)
	ctx.JSON(res.Status(), res.Body())
}

// GET /projects
//
// Get global project list.
//
// Response status:
// 200: StatusOK
// 500: StatusInternalServerError

func (s *ProjectApi) GetProjectList(ctx *gin.Context) {
	res := s.getProjectList(ctx)
	ctx.JSON(res.Status(), res.Body())
}

// GET /orgs/:orgId/workspaces/:workspaceId/projects/:projectId/users
//
// Get project user list.
//
// Response status:
// 200: StatusOK
// 400: StatusBadRequest
// 500: StatusInternalServerError

func (s *ProjectApi) GetMemberList(ctx *gin.Context) {
	uriParams := &project.GetMemberListUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	query := &project.GetMemberListQuery{
		OrgId:       uriParams.OrgId,
		WorkspaceId: uriParams.WorkspaceId,
		ProjectId:   uriParams.ProjectId,
	}
	res := s.getMemberList(ctx, query)
	ctx.JSON(res.Status(), res.Body())
}
