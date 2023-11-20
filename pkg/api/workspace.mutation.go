package api

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/quarkloop/quarkloop/pkg/model"
// 	"github.com/quarkloop/quarkloop/pkg/store/repository"
// )

// type CreateWorkspaceRequest struct {
// 	OrgId     string          `json:"orgId" binding:"required"`
// 	Workspace model.Workspace `json:"workspace" binding:"required"`
// }

// type CreateWorkspaceResponse struct {
// 	ApiResponse
// 	Data model.Workspace `json:"data,omitempty"`
// }

// func (s *ServerApi) CreateWorkspace(c *gin.Context) {
// 	req := &CreateWorkspaceRequest{}
// 	if err := c.BindJSON(req); err != nil {
// 		AbortWithBadRequestJSON(c, err)
// 		return
// 	}

// 	// query database
// 	ws, err := s.dataStore.CreateWorkspace(
// 		&repository.CreateWorkspaceParams{
// 			Context:   c,
// 			OrgId:     req.OrgId,
// 			Workspace: req.Workspace,
// 		},
// 	)
// 	if err != nil {
// 		AbortWithInternalServerErrorJSON(c, err)
// 		return
// 	}

// 	res := &CreateWorkspaceResponse{
// 		ApiResponse: ApiResponse{
// 			Status:       http.StatusCreated,
// 			StatusString: "Created",
// 		},
// 		Data: *ws,
// 	}
// 	c.JSON(http.StatusCreated, res)
// }

// type UpdateWorkspaceByIdUriParams struct {
// 	WorkspaceId string `uri:"workspaceId" binding:"required"`
// }

// type UpdateWorkspaceByIdRequest struct {
// 	model.Workspace
// }

// func (s *ServerApi) UpdateWorkspaceById(c *gin.Context) {
// 	uriParams := &UpdateWorkspaceByIdUriParams{}
// 	if err := c.ShouldBindUri(uriParams); err != nil {
// 		AbortWithBadRequestJSON(c, err)
// 		return
// 	}

// 	req := &UpdateWorkspaceByIdRequest{}
// 	if err := c.BindJSON(req); err != nil {
// 		AbortWithBadRequestJSON(c, err)
// 		return
// 	}

// 	// query database
// 	err := s.dataStore.UpdateWorkspaceById(
// 		&repository.UpdateWorkspaceByIdParams{
// 			Context:     c,
// 			WorkspaceId: uriParams.WorkspaceId,
// 			Workspace:   req.Workspace,
// 		},
// 	)
// 	if err != nil {
// 		AbortWithInternalServerErrorJSON(c, err)
// 		return
// 	}

// 	c.JSON(http.StatusOK, nil)
// }

// type DeleteWorkspaceByIdUriParams struct {
// 	WorkspaceId string `uri:"workspaceId" binding:"required"`
// }

// func (s *ServerApi) DeleteWorkspaceById(c *gin.Context) {
// 	uriParams := &DeleteWorkspaceByIdUriParams{}
// 	if err := c.ShouldBindUri(uriParams); err != nil {
// 		AbortWithBadRequestJSON(c, err)
// 		return
// 	}

// 	// query database
// 	err := s.dataStore.DeleteWorkspaceById(
// 		&repository.DeleteWorkspaceByIdParams{
// 			Context:     c,
// 			WorkspaceId: uriParams.WorkspaceId,
// 		},
// 	)
// 	if err != nil {
// 		AbortWithInternalServerErrorJSON(c, err)
// 		return
// 	}

// 	c.JSON(http.StatusNoContent, nil)
// }
