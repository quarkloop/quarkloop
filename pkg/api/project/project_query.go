package project

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/service/project"
)

// GET /orgs/:orgId/workspaces/:workspaceId/projects/:projectId
//
// Get project by id.
//
// Response status:
// 200: StatusOK
// 500: StatusInternalServerError

func (s *ProjectApi) GetProjectById(ctx *gin.Context) {
	uriParams := &project.GetProjectByIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	// query service
	project, err := s.projectService.GetProjectById(ctx, &project.GetProjectByIdQuery{
		OrgId:       uriParams.OrgId,
		WorkspaceId: uriParams.WorkspaceId,
		ProjectId:   uriParams.ProjectId,
	})
	if err != nil {
		api.AbortWithInternalServerErrorJSON(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, project)
}
