package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateAppComponentRequest struct{}
type CreateAppComponentResponse struct{}

func (s *ServerApi) CreateAppComponent(c *gin.Context) {
	req := &CreateAppComponentRequest{}
	if err := c.BindJSON(req); err != nil {
		AbortWithBadRequestJSON(c, err)
		return
	}

	osId := c.Param("osId")
	workspaceId := c.Param("workspaceId")
	appId := c.Param("appId")
	_ = osId + workspaceId + appId

	// query database

	res := &CreateAppComponentResponse{}
	c.JSON(http.StatusCreated, res)
}

type UpdateAppComponentByIdRequest struct{}
type UpdateAppComponentByIdResponse struct{}

func (s *ServerApi) UpdateAppComponentById(c *gin.Context) {
	req := &UpdateAppComponentByIdRequest{}
	if err := c.BindJSON(req); err != nil {
		AbortWithBadRequestJSON(c, err)
		return
	}

	osId := c.Param("osId")
	workspaceId := c.Param("workspaceId")
	appId := c.Param("appId")
	componentId := c.Param("componentId")
	_ = osId + workspaceId + appId + componentId

	// query database

	res := &UpdateAppComponentByIdResponse{}
	c.JSON(http.StatusOK, res)
}

func (s *ServerApi) DeleteAppComponentById(c *gin.Context) {
	osId := c.Param("osId")
	workspaceId := c.Param("workspaceId")
	appId := c.Param("appId")
	componentId := c.Param("componentId")
	_ = osId + workspaceId + appId + componentId

	// query database

	c.JSON(http.StatusNoContent, nil)
}
