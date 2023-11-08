package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/quarkloop/quarkloop/pkg/db/model"
	"github.com/quarkloop/quarkloop/pkg/db/repository"
)

type GetWorkspaceListResponse struct {
	ApiResponse
	Data []model.Workspace `json:"data,omitempty"`
}

func (s *ServerApi) GetWorkspaceList(c *gin.Context) {
	osId, exists := c.GetQueryArray("osId")
	if !exists {
		AbortWithBadRequestJSON(c, errors.New("missing osId parameter"))
		return
	}

	// query database
	wsList, err := s.dataStore.ListWorkspaces(
		&repository.ListWorkspacesParams{
			Context: c,
			OsId:    osId,
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
	OsId string `form:"osId" binding:"required"`
	model.Workspace
}

type GetWorkspaceResponse struct {
	ApiResponse
	OsId string          `json:"osId" binding:"required"`
	Data model.Workspace `json:"data,omitempty"`
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
			OsId:      queryParams.OsId,
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
