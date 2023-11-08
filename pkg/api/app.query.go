package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/db/model"
	"github.com/quarkloop/quarkloop/pkg/db/repository"
)

type GetAppListUriParams struct {
	OsId        []string `uri:"osId"`
	WorkspaceId []string `uri:"workspaceId"`
}

type GetAppListResponse struct {
	ApiResponse
	Data []model.App `json:"data,omitempty"`
}

func (s *ServerApi) GetAppList(c *gin.Context) {
	uriParams := &GetAppListUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		AbortWithBadRequestJSON(c, err)
		return
	}

	// query database
	appList, err := s.dataStore.ListApps(
		&repository.ListAppsParams{
			Context:     c,
			OsId:        uriParams.OsId,
			WorkspaceId: uriParams.WorkspaceId,
		},
	)
	if err != nil {
		AbortWithInternalServerErrorJSON(c, err)
		return
	}

	res := &GetAppListResponse{
		ApiResponse: ApiResponse{
			Status:       http.StatusOK,
			StatusString: "OK",
		},
		Data: appList,
	}
	c.JSON(http.StatusOK, res)
}

type GetAppByIdUriParams struct {
	OsId        string `uri:"osId" binding:"required"`
	WorkspaceId string `uri:"workspaceId" binding:"required"`
	AppId       string `uri:"appId" binding:"required"`
}

type GetAppByIdResponse struct{}

func (s *ServerApi) GetAppById(c *gin.Context) {
	uriParams := &GetAppByIdUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		AbortWithBadRequestJSON(c, err)
		return
	}

	// query database

	res := &GetAppByIdResponse{}
	c.JSON(http.StatusOK, res)
}
