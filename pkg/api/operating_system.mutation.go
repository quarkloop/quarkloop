package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/quarkloop/quarkloop/pkg/db/model"
	"github.com/quarkloop/quarkloop/pkg/db/repository"
)

type CreateOperatingSystemRequest struct {
	model.OperatingSystem
}

type CreateOperatingSystemResponse struct {
	ApiResponse
	Data model.OperatingSystem `json:"data,omitempty"`
}

func (s *ServerApi) CreateOperatingSystem(c *gin.Context) {
	req := &CreateOperatingSystemRequest{}
	if err := c.BindJSON(req); err != nil {
		AbortWithBadRequestJSON(c, err)
		return
	}

	// query database
	os, err := s.dataStore.CreateOperatingSystem(
		&repository.CreateOperatingSystemParams{
			Context: c,
			Os:      req.OperatingSystem,
		},
	)
	if err != nil {
		AbortWithInternalServerErrorJSON(c, err)
		return
	}

	res := &CreateOperatingSystemResponse{
		ApiResponse: ApiResponse{
			Status:       http.StatusCreated,
			StatusString: "Created",
		},
		Data: *os,
	}
	c.JSON(http.StatusCreated, res)
}

type UpdateOperatingSystemByIdUriParams struct {
	OsId string `uri:"osId" binding:"required"`
}

type UpdateOperatingSystemByIdRequest struct {
	model.OperatingSystem
}

func (s *ServerApi) UpdateOperatingSystemById(c *gin.Context) {
	uriParams := &UpdateOperatingSystemByIdUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		AbortWithBadRequestJSON(c, err)
		return
	}

	req := &UpdateOperatingSystemByIdRequest{}
	if err := c.BindJSON(req); err != nil {
		AbortWithBadRequestJSON(c, err)
		return
	}

	// query database
	err := s.dataStore.UpdateOperatingSystemById(
		&repository.UpdateOperatingSystemByIdParams{
			Context: c,
			OsId:    uriParams.OsId,
			Os:      req.OperatingSystem,
		},
	)
	if err != nil {
		AbortWithInternalServerErrorJSON(c, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}

type DeleteOperatingSystemByIdUriParams struct {
	OsId string `uri:"osId" binding:"required"`
}

func (s *ServerApi) DeleteOperatingSystemById(c *gin.Context) {
	uriParams := &DeleteOperatingSystemByIdUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		AbortWithBadRequestJSON(c, err)
		return
	}

	// query database
	err := s.dataStore.DeleteOperatingSystemById(
		&repository.DeleteOperatingSystemByIdParams{
			Context: c,
			OsId:    uriParams.OsId,
		},
	)
	if err != nil {
		AbortWithInternalServerErrorJSON(c, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
