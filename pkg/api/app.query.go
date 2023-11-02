package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAppListResponse struct{}

func (s *ServerApi) GetAppList(c *gin.Context) {
	res := &GetAppListResponse{}
	c.JSON(http.StatusOK, res)
}

type GetAppByIdRequest struct{}
type GetAppByIdResponse struct{}

func (s *ServerApi) GetAppById(c *gin.Context) {
	req := &GetAppByIdRequest{}

	if err := c.BindJSON(req); err != nil {
		AbortWithBadRequestJSON(c, err)
		return
	}

	res := &GetAppByIdResponse{}
	c.JSON(http.StatusOK, res)
}
