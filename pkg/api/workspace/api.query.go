package workspace

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/service/workspace"
)

type GetWorkspaceListQueryParams struct {
	OrgId []int `form:"orgId"`
}

func (s *WorkspaceApi) GetWorkspaceList(c *gin.Context) {
	queryParams := &GetWorkspaceListQueryParams{}
	if err := c.ShouldBindQuery(queryParams); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	// query service
	wsList, err := s.workspaceService.GetWorkspaceList(
		&workspace.GetWorkspaceListParams{
			Context: c,
			OrgId:   queryParams.OrgId,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	c.JSON(http.StatusOK, &wsList)
}

type GetWorkspaceByIdUriParams struct {
	WorkspaceId int `uri:"workspaceId" binding:"required"`
}

func (s *WorkspaceApi) GetWorkspaceById(c *gin.Context) {
	uriParams := &GetWorkspaceByIdUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	// query service
	ws, err := s.workspaceService.GetWorkspaceById(
		&workspace.GetWorkspaceByIdParams{
			Context:     c,
			WorkspaceId: uriParams.WorkspaceId,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	c.JSON(http.StatusOK, ws)
}

type GetWorkspaceQueryParams struct {
	OrgId int `form:"orgId" binding:"required"`
	workspace.Workspace
}

func (s *WorkspaceApi) GetWorkspace(c *gin.Context) {
	queryParams := &GetWorkspaceQueryParams{}
	if err := c.ShouldBindQuery(queryParams); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	// query service
	ws, err := s.workspaceService.GetWorkspace(
		&workspace.GetWorkspaceParams{
			Context:   c,
			OrgId:     queryParams.OrgId,
			Workspace: queryParams.Workspace,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	c.JSON(http.StatusOK, ws)
}
