package project

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/service/project"
)

type CreateProjectRequest struct {
	OrgId       int             `json:"orgId" binding:"required"`
	WorkspaceId int             `json:"workspaceId" binding:"required"`
	Project     project.Project `json:"project" binding:"required"`
}

func (s *ProjectApi) CreateProject(ctx *gin.Context) {
	req := &CreateProjectRequest{}
	if err := ctx.BindJSON(req); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	// query service
	ws, err := s.projectService.CreateProject(ctx, &project.CreateProjectParams{
		OrgId:       req.OrgId,
		WorkspaceId: req.WorkspaceId,
		Project:     req.Project,
	},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, ws)
}

type UpdateProjectByIdUriParams struct {
	ProjectId int `uri:"projectId" binding:"required"`
}

type UpdateProjectByIdRequest struct {
	project.Project
}

func (s *ProjectApi) UpdateProjectById(ctx *gin.Context) {
	uriParams := &UpdateProjectByIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	req := &UpdateProjectByIdRequest{}
	if err := ctx.BindJSON(req); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	// query service
	err := s.projectService.UpdateProjectById(ctx, &project.UpdateProjectByIdParams{
		ProjectId: uriParams.ProjectId,
		Project:   req.Project,
	},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

type DeleteProjectByIdUriParams struct {
	ProjectId int `uri:"projectId" binding:"required"`
}

func (s *ProjectApi) DeleteProjectById(ctx *gin.Context) {
	uriParams := &DeleteProjectByIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	// query service
	err := s.projectService.DeleteProjectById(ctx, &project.DeleteProjectByIdParams{
		ProjectId: uriParams.ProjectId,
	},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
