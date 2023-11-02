package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAppComponentListResponse struct{}

func (s *ServerApi) GetAppComponentList(c *gin.Context) {
	osId := c.Param("osId")
	workspaceId := c.Param("workspaceId")
	appId := c.Param("appId")
	_ = osId + workspaceId + appId

	// query database

	res := &GetAppComponentListResponse{}
	c.JSON(http.StatusOK, res)
}

type GetAppComponentByIdResponse struct{}

func (s *ServerApi) GetAppComponentById(c *gin.Context) {
	osId := c.Param("osId")
	workspaceId := c.Param("workspaceId")
	appId := c.Param("appId")
	componentId := c.Param("componentId")
	_ = osId + workspaceId + appId + componentId

	// query database

	res := &GetAppComponentByIdResponse{}
	c.JSON(http.StatusOK, res)
}
