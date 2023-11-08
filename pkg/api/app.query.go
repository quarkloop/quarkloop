package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAppListUriParams struct {
	OsId        string `uri:"osId" binding:"required"`
	WorkspaceId string `uri:"workspaceId" binding:"required"`
}

type GetAppListResponse struct{}

func (s *ServerApi) GetAppList(c *gin.Context) {
	uriParams := &GetAppListUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		AbortWithBadRequestJSON(c, err)
		return
	}

	// query database

	res := &GetAppListResponse{}
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
