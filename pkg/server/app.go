package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAppListResponse struct{}
type CreateAppRequest struct{}
type CreateAppResponse struct{}
type GetAppByIdRequest struct{}
type GetAppByIdResponse struct{}
type UpdateAppByIdRequest struct{}
type UpdateAppByIdResponse struct{}
type DeleteAppByIdRequest struct{}

func (s *Server) GetAppList(c *gin.Context) {
	res := &GetAppListResponse{}
	c.JSON(http.StatusOK, res)
}

func (s *Server) CreateApp(c *gin.Context) {
	req := &CreateAppRequest{}

	if err := c.BindJSON(req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, AppResponsePayload{
			Status:       http.StatusBadRequest,
			StatusString: "BadRequest",
			Error:        err,
			ErrorString:  fmt.Sprintf("[BindJSON] %s", err.Error()),
		})
		return
	}

	res := &CreateAppResponse{}
	c.JSON(http.StatusCreated, res)
}

func (s *Server) GetAppById(c *gin.Context) {
	req := &GetAppByIdRequest{}

	if err := c.BindJSON(req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, AppResponsePayload{
			Status:       http.StatusBadRequest,
			StatusString: "BadRequest",
			Error:        err,
			ErrorString:  fmt.Sprintf("[BindJSON] %s", err.Error()),
		})
		return
	}

	res := &GetAppByIdResponse{}
	c.JSON(http.StatusOK, res)
}

func (s *Server) UpdateAppById(c *gin.Context) {
	req := &UpdateAppByIdRequest{}

	if err := c.BindJSON(req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, AppResponsePayload{
			Status:       http.StatusBadRequest,
			StatusString: "BadRequest",
			Error:        err,
			ErrorString:  fmt.Sprintf("[BindJSON] %s", err.Error()),
		})
		return
	}

	res := &UpdateAppByIdResponse{}
	c.JSON(http.StatusOK, res)
}

func (s *Server) DeleteAppById(c *gin.Context) {
	req := &DeleteAppByIdRequest{}

	if err := c.BindJSON(req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, AppResponsePayload{
			Status:       http.StatusBadRequest,
			StatusString: "BadRequest",
			Error:        err,
			ErrorString:  fmt.Sprintf("[BindJSON] %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
