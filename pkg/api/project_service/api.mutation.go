package project_service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/project_service"
)

type CreateProjectServiceUriParams struct {
	ProjectId string `uri:"projectId" binding:"required"`
}

type CreateProjectServiceRequest struct {
	model.ProjectService
}

type CreateProjectServiceResponse struct {
	api.ApiResponse
	Data model.ProjectService `json:"data,omitempty"`
}

func (s *ProjectServiceApi) CreateProjectService(c *gin.Context) {
	uriParams := &CreateProjectServiceUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	req := &CreateProjectServiceRequest{}
	if err := c.BindJSON(req); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	// query service
	ws, err := s.projectService.CreateProjectService(
		&project_service.CreateProjectServiceParams{
			Context:        c,
			ProjectId:      uriParams.ProjectId,
			ProjectService: &req.ProjectService,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	res := &CreateProjectServiceResponse{
		ApiResponse: api.ApiResponse{
			Status:       http.StatusCreated,
			StatusString: "Created",
		},
		Data: *ws,
	}
	c.JSON(http.StatusCreated, res)
}

type UpdateProjectServiceByIdUriParams struct {
	ProjectServiceId string `uri:"projectServiceId" binding:"required"`
}

type UpdateProjectServiceByIdRequest struct {
	model.ProjectService
}

func (s *ProjectServiceApi) UpdateProjectServiceById(c *gin.Context) {
	uriParams := &UpdateProjectServiceByIdUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	req := &UpdateProjectServiceByIdRequest{}
	if err := c.BindJSON(req); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	// query service
	err := s.projectService.UpdateProjectServiceById(
		&project_service.UpdateProjectServiceByIdParams{
			Context:          c,
			ProjectServiceId: uriParams.ProjectServiceId,
			ProjectService:   &req.ProjectService,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}

type DeleteProjectServiceByIdUriParams struct {
	ProjectServiceId string `uri:"projectServiceId" binding:"required"`
}

func (s *ProjectServiceApi) DeleteProjectServiceById(c *gin.Context) {
	uriParams := &DeleteProjectServiceByIdUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	// query service
	err := s.projectService.DeleteProjectServiceById(
		&project_service.DeleteProjectServiceByIdParams{
			Context:          c,
			ProjectServiceId: uriParams.ProjectServiceId,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
