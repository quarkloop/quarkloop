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

	orgId := c.Param("orgId")
	workspaceId := c.Param("workspaceId")
	projectId := c.Param("projectId")
	_ = orgId + workspaceId + projectId

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

	orgId := c.Param("orgId")
	workspaceId := c.Param("workspaceId")
	projectId := c.Param("projectId")
	issueId := c.Param("issueId")
	_ = orgId + workspaceId + projectId + issueId

	// query database

	res := &UpdateAppIssueByIdResponse{}
	c.JSON(http.StatusOK, res)
}

func (s *ServerApi) DeleteAppIssueById(c *gin.Context) {
	orgId := c.Param("orgId")
	workspaceId := c.Param("workspaceId")
	projectId := c.Param("projectId")
	issueId := c.Param("issueId")
	_ = orgId + workspaceId + projectId + issueId

	// query database

	c.JSON(http.StatusNoContent, nil)
}
