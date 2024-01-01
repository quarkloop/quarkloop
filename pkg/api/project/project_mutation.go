package project

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/service/project"
)

// POST /orgs/:orgId/workspaces/:workspaceId/projects
//
// Create project.
//
// Response status:
// 201: StatusCreated
// 400: StatusBadRequest
// 500: StatusInternalServerError

func (s *ProjectApi) CreateProject(ctx *gin.Context) {
	uriParams := &project.CreateProjectUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	cmd := &project.CreateProjectCommand{}
	if err := ctx.ShouldBindJSON(cmd); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	project, err := s.createProject(ctx, &project.CreateProjectCommand{
		OrgId:       uriParams.OrgId,
		WorkspaceId: uriParams.WorkspaceId,
		Project:     cmd.Project,
	})
	if err != nil {
		api.AbortWithInternalServerErrorJSON(ctx, err)
		return
	}

	res := s.CreateProject(ctx, cmd)
	ctx.JSON(res.Status(), res.Body())
}

// PUT /orgs/:orgId/workspaces/:workspaceId/projects/:projectId
//
// Update project by id.
//
// Response status:
// 200: StatusOK
// 400: StatusBadRequest
// 500: StatusInternalServerError

func (s *ProjectApi) UpdateProjectById(ctx *gin.Context) {
	uriParams := &project.UpdateProjectByIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	cmd := &project.UpdateProjectByIdCommand{}
	if err := ctx.ShouldBindJSON(cmd); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// query service
	err := s.projectService.UpdateProjectById(ctx, &project.UpdateProjectByIdCommand{
		OrgId:       uriParams.OrgId,
		WorkspaceId: uriParams.WorkspaceId,
		ProjectId:   uriParams.ProjectId,
		Project:     cmd.Project,
	})
	if err != nil {
		api.AbortWithInternalServerErrorJSON(ctx, err)
		return
	}

	res := s.UpdateProjectById(ctx, cmd)
	ctx.JSON(res.Status(), res.Body())
}

// DELETE /orgs/:orgId/workspaces/:workspaceId/projects/:projectId
//
// Delete project by id.
//
// Response status:
// 204: StatusNoContent
// 400: StatusBadRequest
// 500: StatusInternalServerError

func (s *ProjectApi) DeleteProjectById(ctx *gin.Context) {
	uriParams := &project.DeleteProjectByIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// query service
	err := s.projectService.DeleteProjectById(ctx, &project.DeleteProjectByIdCommand{
		OrgId:       uriParams.OrgId,
		WorkspaceId: uriParams.WorkspaceId,
		ProjectId:   uriParams.ProjectId,
	})
	if err != nil {
		api.AbortWithInternalServerErrorJSON(ctx, err)
		return
	}

	res := s.DeleteProjectById(ctx, cmd)
	ctx.JSON(res.Status(), res.Body())
}
