package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateAppIssueRequest struct{}
type CreateAppIssueResponse struct{}

func (s *ServerApi) CreateAppIssue(c *gin.Context) {
	req := &CreateAppIssueRequest{}
	if err := c.BindJSON(req); err != nil {
		AbortWithBadRequestJSON(c, err)
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

type UpdateAppIssueByIdRequest struct{}
type UpdateAppIssueByIdResponse struct{}

func (s *ServerApi) UpdateAppIssueById(c *gin.Context) {
	req := &UpdateAppIssueByIdRequest{}
	if err := c.BindJSON(req); err != nil {
		AbortWithBadRequestJSON(c, err)
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

func (s *ServerApi) DeleteAppIssueById(c *gin.Context) {
	osId := c.Param("osId")
	workspaceId := c.Param("workspaceId")
	appId := c.Param("appId")
	issueId := c.Param("issueId")
	_ = osId + workspaceId + appId + issueId

	// query database

	c.JSON(http.StatusNoContent, nil)
}
