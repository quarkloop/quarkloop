package api

import (
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
	osId := c.Param("osId")

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
	ws, err := s.dataStore.GetWorkspaceById(
		&repository.GetWorkspaceByIdParams{
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
