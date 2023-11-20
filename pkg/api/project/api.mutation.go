package project

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/project"
)

type CreateProjectRequest struct {
	OrgId       string        `json:"orgId" binding:"required"`
	WorkspaceId string        `json:"workspaceId" binding:"required"`
	Project     model.Project `json:"project" binding:"required"`
}

type CreateProjectResponse struct {
	api.ApiResponse
	Data model.Project `json:"data,omitempty"`
}

func (s *ProjectApi) CreateProject(c *gin.Context) {
	req := &CreateProjectRequest{}
	if err := c.BindJSON(req); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	// query service
	ws, err := s.projectService.CreateProject(
		&project.CreateProjectParams{
			Context:     c,
			OrgId:       req.OrgId,
			WorkspaceId: req.WorkspaceId,
			Project:     req.Project,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	res := &CreateProjectResponse{
		ApiResponse: api.ApiResponse{
			Status:       http.StatusCreated,
			StatusString: "Created",
		},
		Data: *ws,
	}
	c.JSON(http.StatusCreated, res)
}

type UpdateProjectByIdUriParams struct {
	ProjectId string `uri:"projectId" binding:"required"`
}

type UpdateProjectByIdRequest struct {
	model.Project
}

func (s *ProjectApi) UpdateProjectById(c *gin.Context) {
	uriParams := &UpdateProjectByIdUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	req := &UpdateProjectByIdRequest{}
	if err := c.BindJSON(req); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	// query service
	err := s.projectService.UpdateProjectById(
		&project.UpdateProjectByIdParams{
			Context:   c,
			ProjectId: uriParams.ProjectId,
			Project:   req.Project,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}

type DeleteProjectByIdUriParams struct {
	ProjectId string `uri:"projectId" binding:"required"`
}

func (s *ProjectApi) DeleteProjectById(c *gin.Context) {
	uriParams := &DeleteProjectByIdUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	// query service
	err := s.projectService.DeleteProjectById(
		&project.DeleteProjectByIdParams{
			Context:   c,
			ProjectId: uriParams.ProjectId,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
