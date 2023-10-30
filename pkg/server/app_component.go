package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAppComponentListResponse struct{}
type CreateAppComponentRequest struct{}
type CreateAppComponentResponse struct{}
type GetAppComponentByIdResponse struct{}
type UpdateAppComponentByIdRequest struct{}
type UpdateAppComponentByIdResponse struct{}

func (s *Server) GetAppComponentList(c *gin.Context) {
	osId := c.Param("osId")
	workspaceId := c.Param("workspaceId")
	appId := c.Param("appId")
	_ = osId + workspaceId + appId

	// query database

	res := &GetAppComponentListResponse{}
	c.JSON(http.StatusOK, res)
}

func (s *Server) CreateAppComponent(c *gin.Context) {
	req := &CreateAppComponentRequest{}
	if err := c.BindJSON(req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, AppResponsePayload{
			Status:       http.StatusBadRequest,
			StatusString: "BadRequest",
			Error:        err,
			ErrorString:  fmt.Sprintf("[BindJSON] %s", err.Error()),
		})
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

func (s *Server) GetAppComponentById(c *gin.Context) {
	osId := c.Param("osId")
	workspaceId := c.Param("workspaceId")
	appId := c.Param("appId")
	componentId := c.Param("componentId")
	_ = osId + workspaceId + appId + componentId

	// query database

	res := &GetAppComponentByIdResponse{}
	c.JSON(http.StatusOK, res)
}

func (s *Server) UpdateAppComponentById(c *gin.Context) {
	req := &UpdateAppComponentByIdRequest{}
	if err := c.BindJSON(req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, AppResponsePayload{
			Status:       http.StatusBadRequest,
			StatusString: "BadRequest",
			Error:        err,
			ErrorString:  fmt.Sprintf("[BindJSON] %s", err.Error()),
		})
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

func (s *Server) DeleteAppComponentById(c *gin.Context) {
	osId := c.Param("osId")
	workspaceId := c.Param("workspaceId")
	appId := c.Param("appId")
	componentId := c.Param("componentId")
	_ = osId + workspaceId + appId + componentId

	// query database

	c.JSON(http.StatusNoContent, nil)
}
