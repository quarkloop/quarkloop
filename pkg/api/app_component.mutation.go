package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/db/model"
	"github.com/quarkloop/quarkloop/pkg/db/repository"
)

type CreateAppComponentUriParams struct {
	AppId string `uri:"appId" binding:"required"`
}

type CreateAppComponentRequest struct {
	model.AppComponent
}

type CreateAppComponentResponse struct {
	ApiResponse
	Data model.SheetInstance `json:"data,omitempty"`
}

func (s *ServerApi) CreateAppComponent(c *gin.Context) {
	uriParams := &CreateAppComponentUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		AbortWithBadRequestJSON(c, err)
		return
	}

	req := &CreateAppComponentRequest{}
	if err := c.BindJSON(req); err != nil {
		AbortWithBadRequestJSON(c, err)
		return
	}

	// query database
	ws, err := s.dataStore.CreateSheetInstance(
		&repository.CreateSheetInstanceParams{
			Context:       c,
			AppId:         uriParams.AppId,
			SheetInstance: req.SheetInstance,
		},
	)
	if err != nil {
		AbortWithInternalServerErrorJSON(c, err)
		return
	}

	res := &CreateAppComponentResponse{
		ApiResponse: ApiResponse{
			Status:       http.StatusCreated,
			StatusString: "Created",
		},
		Data: *ws,
	}
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
