package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/store/repository"
)

type GetWorkspaceListResponse struct {
	ApiResponse
	Data []model.Workspace `json:"data"`
}

func (s *ServerApi) GetWorkspaceList(c *gin.Context) {
	orgId, exists := c.GetQueryArray("orgId")
	if !exists {
		AbortWithBadRequestJSON(c, errors.New("missing orgId parameter"))
		return
	}

	// query database
	wsList, err := s.dataStore.ListWorkspaces(
		&repository.ListWorkspacesParams{
			Context: c,
			OrgId:   orgId,
		},
	)
	if err != nil {
		AbortWithInternalServerErrorJSON(c, err)
		return
	}

	res := &GetWorkspaceListResponse{
		ApiResponse: ApiResponse{
			Status:       http.StatusOK,
			StatusString: "OK",
		},
		Data: wsList,
	}
	c.JSON(http.StatusOK, res)
}

type GetWorkspaceByIdResponse struct {
	ApiResponse
	Data model.Workspace `json:"data,omitempty"`
}

func (s *ServerApi) GetWorkspaceById(c *gin.Context) {
	wsId := c.Param("workspaceId")

	// query database
	ws, err := s.dataStore.FindUniqueWorkspace(
		&repository.FindUniqueWorkspaceParams{
			Context: c,
			Id:      wsId,
		},
	)
	if err != nil {
		AbortWithInternalServerErrorJSON(c, err)
		return
	}

	res := &GetWorkspaceByIdResponse{
		ApiResponse: ApiResponse{
			Status:       http.StatusOK,
			StatusString: "OK",
		},
		Data: *ws,
	}
	c.JSON(http.StatusOK, res)
}

type GetWorkspaceQueryParams struct {
	OrgId string `form:"orgId" binding:"required"`
	model.Workspace
}

type GetWorkspaceResponse struct {
	ApiResponse
	OrgId string          `json:"orgId" binding:"required"`
	Data  model.Workspace `json:"data,omitempty"`
}

func (s *ServerApi) GetWorkspace(c *gin.Context) {
	queryParams := &GetWorkspaceQueryParams{}
	if err := c.ShouldBindQuery(queryParams); err != nil {
		AbortWithBadRequestJSON(c, err)
		return
	}

	// query database
	ws, err := s.dataStore.FindFirstWorkspace(
		&repository.FindFirstWorkspaceParams{
			Context:   c,
			OrgId:     queryParams.OrgId,
			Workspace: queryParams.Workspace,
		},
	)
	if err != nil {
		AbortWithInternalServerErrorJSON(c, err)
		return
	}

	res := &GetWorkspaceResponse{
		ApiResponse: ApiResponse{
			Status:       http.StatusOK,
			StatusString: "OK",
		},
		Data: *ws,
	}
	c.JSON(http.StatusOK, res)
}
