package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAppIssueListResponse struct{}
type CreateAppIssueRequest struct{}
type CreateAppIssueResponse struct{}
type GetAppIssueByIdResponse struct{}
type UpdateAppIssueByIdRequest struct{}
type UpdateAppIssueByIdResponse struct{}

func (s *Server) GetAppIssueList(c *gin.Context) {
	osId := c.Param("osId")
	workspaceId := c.Param("workspaceId")
	appId := c.Param("appId")
	_ = osId + workspaceId + appId

	// query database

	res := &GetAppIssueListResponse{}
	c.JSON(http.StatusOK, res)
}

func (s *Server) CreateAppIssue(c *gin.Context) {
	req := &CreateAppIssueRequest{}
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

	res := &CreateAppIssueResponse{}
	c.JSON(http.StatusCreated, res)
}

func (s *Server) GetAppIssueById(c *gin.Context) {
	osId := c.Param("osId")
	workspaceId := c.Param("workspaceId")
	appId := c.Param("appId")
	issueId := c.Param("issueId")
	_ = osId + workspaceId + appId + issueId

	// query database

	res := &GetAppIssueByIdResponse{}
	c.JSON(http.StatusOK, res)
}

func (s *Server) UpdateAppIssueById(c *gin.Context) {
	req := &UpdateAppIssueByIdRequest{}
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
	issueId := c.Param("issueId")
	_ = osId + workspaceId + appId + issueId

	// query database

	res := &UpdateAppIssueByIdResponse{}
	c.JSON(http.StatusOK, res)
}

func (s *Server) DeleteAppIssueById(c *gin.Context) {
	osId := c.Param("osId")
	workspaceId := c.Param("workspaceId")
	appId := c.Param("appId")
	issueId := c.Param("issueId")
	_ = osId + workspaceId + appId + issueId

	// query database

	c.JSON(http.StatusNoContent, nil)
}
