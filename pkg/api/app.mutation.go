package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/db/model"
	"github.com/quarkloop/quarkloop/pkg/db/repository"
)

type CreateAppRequest struct {
	OsId        string    `json:"osId" binding:"required"`
	WorkspaceId string    `json:"workspaceId" binding:"required"`
	App         model.App `json:"app" binding:"required"`
}

type CreateAppResponse struct {
	ApiResponse
	Data model.App `json:"data,omitempty"`
}

func (s *ServerApi) CreateApp(c *gin.Context) {
	req := &CreateAppRequest{}
	if err := c.BindJSON(req); err != nil {
		AbortWithBadRequestJSON(c, err)
		return
	}

	// query database
	ws, err := s.dataStore.CreateApp(
		&repository.CreateAppParams{
			Context:     c,
			OsId:        req.OsId,
			WorkspaceId: req.WorkspaceId,
			App:         req.App,
		},
	)
	if err != nil {
		AbortWithInternalServerErrorJSON(c, err)
		return
	}

	res := &CreateAppResponse{
		ApiResponse: ApiResponse{
			Status:       http.StatusCreated,
			StatusString: "Created",
		},
		Data: *ws,
	}
	c.JSON(http.StatusCreated, res)
}

type UpdateAppByIdUriParams struct {
	AppId string `uri:"appId" binding:"required"`
}

type UpdateAppByIdRequest struct {
	model.App
}

func (s *ServerApi) UpdateAppById(c *gin.Context) {
	uriParams := &UpdateAppByIdUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		AbortWithBadRequestJSON(c, err)
		return
	}

	req := &UpdateAppByIdRequest{}
	if err := c.BindJSON(req); err != nil {
		AbortWithBadRequestJSON(c, err)
		return
	}

	// query database
	err := s.dataStore.UpdateAppById(
		&repository.UpdateAppByIdParams{
			Context: c,
			AppId:   uriParams.AppId,
			App:     req.App,
		},
	)
	if err != nil {
		AbortWithInternalServerErrorJSON(c, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}

type DeleteAppByIdUriParams struct {
	AppId string `uri:"appId" binding:"required"`
}

func (s *ServerApi) DeleteAppById(c *gin.Context) {
	uriParams := &DeleteAppByIdUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		AbortWithBadRequestJSON(c, err)
		return
	}

	// query database
	err := s.dataStore.DeleteAppById(
		&repository.DeleteAppByIdParams{
			Context: c,
			AppId:   uriParams.AppId,
		},
	)
	if err != nil {
		AbortWithInternalServerErrorJSON(c, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
