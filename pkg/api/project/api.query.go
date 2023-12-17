package project

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/service/project"
)

type GetProjectListQueryParams struct {
	OrgId       []int `form:"orgId"`
	WorkspaceId []int `form:"workspaceId"`
}

func (s *ProjectApi) GetProjectList(ctx *gin.Context) {
	queryParams := &GetProjectListQueryParams{}
	if err := ctx.ShouldBindQuery(queryParams); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	// query service
	projectList, err := s.projectService.GetProjectList(ctx, &project.GetProjectListParams{
		OrgId:       queryParams.OrgId,
		WorkspaceId: queryParams.WorkspaceId,
	},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, &projectList)
}

type GetProjectByIdUriParams struct {
	ProjectId int `uri:"projectId" binding:"required"`
}

func (s *ProjectApi) GetProjectById(ctx *gin.Context) {
	uriParams := &GetProjectByIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	// query service
	project, err := s.projectService.GetProjectById(ctx, &project.GetProjectByIdParams{
		ProjectId: uriParams.ProjectId,
	},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, project)
}
