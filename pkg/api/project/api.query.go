package project

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/project"
)

type GetProjectListQueryParams struct {
	OrgId       []string `form:"orgId"`
	WorkspaceId []string `form:"workspaceId"`
}

type GetProjectListResponse struct {
	api.ApiResponse
	Data []model.Project `json:"data"`
}

func (s *ProjectApi) GetProjectList(c *gin.Context) {
	uriParams := &GetProjectListQueryParams{}
	if err := c.ShouldBindQuery(uriParams); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	// query service
	projectList, err := s.projectService.GetProjectList(
		&project.GetProjectListParams{
			Context:     c,
			OrgId:       uriParams.OrgId,
			WorkspaceId: uriParams.WorkspaceId,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	res := &GetProjectListResponse{
		ApiResponse: api.ApiResponse{
			Status:       http.StatusOK,
			StatusString: "OK",
		},
		Data: projectList,
	}
	c.JSON(http.StatusOK, res)
}

type GetProjectByIdUriParams struct {
	ProjectId string `uri:"projectId" binding:"required"`
}

type GetProjectByIdResponse struct {
	api.ApiResponse
	Data model.Project `json:"data,omitempty"`
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

	res := &GetProjectByIdResponse{
		ApiResponse: api.ApiResponse{
			Status:       http.StatusOK,
			StatusString: "OK",
		},
		Data: *project,
	}
	c.JSON(http.StatusOK, res)
}
