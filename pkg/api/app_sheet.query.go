package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/quarkloop/quarkloop/pkg/db/model"
	"github.com/quarkloop/quarkloop/pkg/db/repository"
)

type GetSheetInstanceListUriParams struct {
	AppId string `uri:"appId" binding:"required"`
}

type GetSheetInstanceListResponse struct {
	ApiResponse
	Data []model.SheetInstance `json:"data,omitempty"`
}

func (s *ServerApi) GetSheetInstanceList(c *gin.Context) {
	uriParams := &GetSheetInstanceListUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		AbortWithBadRequestJSON(c, err)
		return
	}

	// query database
	instanceList, err := s.dataStore.ListSheetInstances(
		&repository.ListSheetInstancesParams{
			Context: c,
			AppId:   uriParams.AppId,
		},
	)
	if err != nil {
		AbortWithInternalServerErrorJSON(c, err)
		return
	}

	res := &GetSheetInstanceListResponse{
		ApiResponse: ApiResponse{
			Status:       http.StatusOK,
			StatusString: "OK",
		},
		Data: instanceList,
	}
	c.JSON(http.StatusOK, res)
}

type GetSheetInstanceByIdUriParams struct {
	AppId      string `uri:"appId" binding:"required"`
	InstanceId int    `uri:"instanceId" binding:"required"`
}

type GetSheetInstanceByIdResponse struct {
	ApiResponse
	Data model.SheetInstance `json:"data,omitempty"`
}

func (s *ServerApi) GetSheetInstanceById(c *gin.Context) {
	uriParams := &GetSheetInstanceByIdUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		AbortWithBadRequestJSON(c, err)
		return
	}

	// query database
	instance, err := s.dataStore.FindUniqueSheetInstance(
		&repository.FindUniqueSheetInstanceParams{
			Context: c,
			AppId:   uriParams.AppId,
			Id:      uriParams.InstanceId,
		},
	)
	if err != nil {
		AbortWithInternalServerErrorJSON(c, err)
		return
	}

	res := &GetSheetInstanceByIdResponse{
		ApiResponse: ApiResponse{
			Status:       http.StatusOK,
			StatusString: "OK",
		},
		Data: *instance,
	}
	c.JSON(http.StatusOK, res)
}

type GetSheetInstanceUriParams struct {
	AppId string `uri:"appId" binding:"required"`
}

type GetSheetInstanceQueryParams struct {
	model.SheetInstance
}

type GetSheetInstanceResponse struct {
	ApiResponse
	Data model.SheetInstance `json:"data,omitempty"`
}

func (s *ServerApi) GetSheetInstance(c *gin.Context) {
	uriParams := &GetSheetInstanceUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		AbortWithBadRequestJSON(c, err)
		return
	}

	queryParams := &GetSheetInstanceQueryParams{}
	if err := c.ShouldBindQuery(queryParams); err != nil {
		AbortWithBadRequestJSON(c, err)
		return
	}

	// query database
	instance, err := s.dataStore.FindFirstSheetInstance(
		&repository.FindFirstSheetInstanceParams{
			Context:       c,
			AppId:         uriParams.AppId,
			SheetInstance: queryParams.SheetInstance,
		},
	)
	if err != nil {
		AbortWithInternalServerErrorJSON(c, err)
		return
	}

	res := &GetSheetInstanceResponse{
		ApiResponse: ApiResponse{
			Status:       http.StatusOK,
			StatusString: "OK",
		},
		Data: *instance,
	}
	c.JSON(http.StatusOK, res)
}
