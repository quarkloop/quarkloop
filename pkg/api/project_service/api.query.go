package project_service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/project_service"
)

type GetProjectServiceListUriParams struct {
	ProjectId string `uri:"projectId"`
}

type GetProjectServiceListResponse struct {
	api.ApiResponse
	Data []model.ProjectService `json:"data"`
}

func (s *ProjectServiceApi) GetProjectServiceList(c *gin.Context) {
	uriParams := &GetProjectServiceListUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	// query service
	projectList, err := s.projectService.GetProjectServiceList(
		&project_service.GetProjectServiceListParams{
			Context:   c,
			ProjectId: uriParams.ProjectId,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	res := &GetProjectServiceListResponse{
		ApiResponse: api.ApiResponse{
			Status:       http.StatusOK,
			StatusString: "OK",
		},
		Data: projectList,
	}
	c.JSON(http.StatusOK, res)
}

type GetProjectServiceByIdUriParams struct {
	ProjectId        string `uri:"projectId"`
	ProjectServiceId string `uri:"projectServiceId" binding:"required"`
}

type GetProjectServiceByIdResponse struct {
	api.ApiResponse
	Data model.ProjectService `json:"data,omitempty"`
}

func (s *ProjectServiceApi) GetProjectServiceById(c *gin.Context) {
	uriParams := &GetProjectServiceByIdUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	// query service
	project_service, err := s.projectService.GetProjectServiceById(
		&project_service.GetProjectServiceByIdParams{
			Context:          c,
			ProjectId:        uriParams.ProjectId,
			ProjectServiceId: uriParams.ProjectServiceId,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	res := &GetProjectServiceByIdResponse{
		ApiResponse: api.ApiResponse{
			Status:       http.StatusOK,
			StatusString: "OK",
		},
		Data: *project_service,
	}
	c.JSON(http.StatusOK, res)
}
