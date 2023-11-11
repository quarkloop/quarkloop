package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/db/model"
	"github.com/quarkloop/quarkloop/pkg/db/repository"
)

type CreateSheetInstanceUriParams struct {
	AppId string `uri:"appId" binding:"required"`
}

type CreateSheetInstanceRequest struct {
	model.SheetInstance
}

type CreateSheetInstanceResponse struct {
	ApiResponse
	Data model.SheetInstance `json:"data,omitempty"`
}

func (s *ServerApi) CreateSheetInstance(c *gin.Context) {
	uriParams := &CreateSheetInstanceUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		AbortWithBadRequestJSON(c, err)
		return
	}

	req := &CreateSheetInstanceRequest{}
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

	res := &CreateSheetInstanceResponse{
		ApiResponse: ApiResponse{
			Status:       http.StatusCreated,
			StatusString: "Created",
		},
		Data: *ws,
	}
	c.JSON(http.StatusCreated, res)
}

type UpdateSheetInstanceByIdUriParams struct {
	AppId      string `uri:"appId" binding:"required"`
	InstanceId int    `uri:"instanceId" binding:"required"`
}

type UpdateSheetInstanceByIdRequest struct {
	model.SheetInstance
}

func (s *ServerApi) UpdateSheetInstanceById(c *gin.Context) {
	uriParams := &UpdateSheetInstanceByIdUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		AbortWithBadRequestJSON(c, err)
		return
	}

	req := &UpdateSheetInstanceByIdRequest{}
	if err := c.BindJSON(req); err != nil {
		AbortWithBadRequestJSON(c, err)
		return
	}

	// query database
	err := s.dataStore.UpdateSheetInstanceById(
		&repository.UpdateSheetInstanceByIdParams{
			Context:       c,
			AppId:         uriParams.AppId,
			Id:            uriParams.InstanceId,
			SheetInstance: req.SheetInstance,
		},
	)
	if err != nil {
		AbortWithInternalServerErrorJSON(c, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}

type DeleteSheetInstanceByIdUriParams struct {
	AppId      string `uri:"appId" binding:"required"`
	InstanceId int    `uri:"instanceId" binding:"required"`
}

func (s *ServerApi) DeleteSheetInstanceById(c *gin.Context) {
	uriParams := &DeleteSheetInstanceByIdUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		AbortWithBadRequestJSON(c, err)
		return
	}

	// query database
	err := s.dataStore.DeleteSheetInstanceById(
		&repository.DeleteSheetInstanceByIdParams{
			Context: c,
			AppId:   uriParams.AppId,
			Id:      uriParams.InstanceId,
		},
	)
	if err != nil {
		AbortWithInternalServerErrorJSON(c, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
