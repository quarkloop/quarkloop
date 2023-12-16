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

func (s *ProjectApi) GetProjectList(c *gin.Context) {
	queryParams := &GetProjectListQueryParams{}
	if err := c.ShouldBindQuery(queryParams); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	// query service
	projectList, err := s.projectService.GetProjectList(
		&project.GetProjectListParams{
			Context:     c,
			OrgId:       queryParams.OrgId,
			WorkspaceId: queryParams.WorkspaceId,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	c.JSON(http.StatusOK, &projectList)
}

type GetProjectByIdUriParams struct {
	ProjectId int `uri:"projectId" binding:"required"`
}

func (s *ProjectApi) GetProjectById(c *gin.Context) {
	uriParams := &GetProjectByIdUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	// query service
	project, err := s.projectService.GetProjectById(
		&project.GetProjectByIdParams{
			Context:   c,
			ProjectId: uriParams.ProjectId,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	c.JSON(http.StatusOK, project)
}
