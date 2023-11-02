package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateAppRequest struct{}
type CreateAppResponse struct{}

func (s *ServerApi) CreateApp(c *gin.Context) {
	req := &CreateAppRequest{}

	if err := c.BindJSON(req); err != nil {
		AbortWithBadRequestJSON(c, err)
		return
	}

	res := &CreateAppResponse{}
	c.JSON(http.StatusCreated, res)
}

type UpdateAppByIdRequest struct{}
type UpdateAppByIdResponse struct{}

func (s *ServerApi) UpdateAppById(c *gin.Context) {
	req := &UpdateAppByIdRequest{}

	if err := c.BindJSON(req); err != nil {
		AbortWithBadRequestJSON(c, err)
		return
	}

	res := &UpdateAppByIdResponse{}
	c.JSON(http.StatusOK, res)
}

type DeleteAppByIdRequest struct{}

func (s *ServerApi) DeleteAppById(c *gin.Context) {
	req := &DeleteAppByIdRequest{}

	if err := c.BindJSON(req); err != nil {
		AbortWithBadRequestJSON(c, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
