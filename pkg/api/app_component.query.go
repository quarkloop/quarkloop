package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/db/model"
	"github.com/quarkloop/quarkloop/pkg/db/repository"
)

type GetAppComponentListUriParams struct {
	AppId string `uri:"appId" binding:"required"`
}

type GetAppComponentListResponse struct {
	ApiResponse
	Data []model.AppComponent `json:"data,omitempty"`
}

func (s *ServerApi) GetAppComponentList(c *gin.Context) {
	uriParams := &GetAppComponentListUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		AbortWithBadRequestJSON(c, err)
		return
	}

	// query database
	compList, err := s.dataStore.ListAppComponents(
		&repository.ListAppComponentsParams{
			Context: c,
			AppId:   uriParams.AppId,
		},
	)
	if err != nil {
		AbortWithInternalServerErrorJSON(c, err)
		return
	}

	res := &GetAppComponentListResponse{
		ApiResponse: ApiResponse{
			Status:       http.StatusOK,
			StatusString: "OK",
		},
		Data: compList,
	}
	c.JSON(http.StatusOK, res)
}

type GetAppComponentByIdUriParams struct {
	AppId       string `uri:"appId" binding:"required"`
	ComponentId string `uri:"componentId" binding:"required"`
}

type GetAppComponentByIdResponse struct {
	ApiResponse
	Data model.AppComponent `json:"data,omitempty"`
}

func (s *ServerApi) GetAppComponentById(c *gin.Context) {
	uriParams := &GetAppComponentByIdUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		AbortWithBadRequestJSON(c, err)
		return
	}

	// query database
	comp, err := s.dataStore.FindUniqueAppComponent(
		&repository.FindUniqueAppComponentParams{
			Context:     c,
			AppId:       uriParams.AppId,
			ComponentId: uriParams.ComponentId,
		},
	)
	if err != nil {
		AbortWithInternalServerErrorJSON(c, err)
		return
	}

	res := &GetAppComponentByIdResponse{
		ApiResponse: ApiResponse{
			Status:       http.StatusOK,
			StatusString: "OK",
		},
		Data: *comp,
	}
	c.JSON(http.StatusOK, res)
}
