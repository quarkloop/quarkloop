package workspace

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/workspace"
)

type GetWorkspaceListUriParams struct {
	OrgId []int `uri:"orgId" binding:"required"`
}

type GetWorkspaceListResponse struct {
	api.ApiResponse
	Data []model.Workspace `json:"data"`
}

func (s *WorkspaceApi) GetWorkspaceList(c *gin.Context) {
	uriParams := &GetWorkspaceListUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	// query service
	wsList, err := s.workspaceService.GetWorkspaceList(
		&workspace.GetWorkspaceListParams{
			Context: c,
			OrgId:   uriParams.OrgId,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	res := &GetWorkspaceListResponse{
		ApiResponse: api.ApiResponse{
			Status:       http.StatusOK,
			StatusString: "OK",
		},
		Data: wsList,
	}
	c.JSON(http.StatusOK, res)
}

type GetWorkspaceByIdUriParams struct {
	workspaceId int `uri:"orgId" binding:"required"`
}

type GetWorkspaceByIdResponse struct {
	api.ApiResponse
	Data model.Workspace `json:"data,omitempty"`
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
			WorkspaceId: uriParams.workspaceId,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	res := &GetWorkspaceByIdResponse{
		ApiResponse: api.ApiResponse{
			Status:       http.StatusOK,
			StatusString: "OK",
		},
		Data: *ws,
	}
	c.JSON(http.StatusOK, res)
}

type GetWorkspaceQueryParams struct {
	OrgId int `form:"orgId" binding:"required"`
	model.Workspace
}

type GetWorkspaceResponse struct {
	api.ApiResponse
	OrgId int             `json:"orgId" binding:"required"`
	Data  model.Workspace `json:"data,omitempty"`
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

	res := &GetWorkspaceResponse{
		ApiResponse: api.ApiResponse{
			Status:       http.StatusOK,
			StatusString: "OK",
		},
		Data: *ws,
	}
	c.JSON(http.StatusOK, res)
}
