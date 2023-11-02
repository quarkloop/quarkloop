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

type UpdateOperatingSystemByIdRequest struct {
	model.OperatingSystem
}

func (s *ServerApi) UpdateOperatingSystemById(c *gin.Context) {
	req := &UpdateOperatingSystemByIdRequest{}
	if err := c.BindJSON(req); err != nil {
		AbortWithBadRequestJSON(c, err)
		return
	}

	osId := c.Param("osId")

	// query database
	err := s.dataStore.UpdateOperatingSystemById(
		&repository.UpdateOperatingSystemByIdParams{
			Context: c,
			OsId:    osId,
			Os:      req.OperatingSystem,
		},
	)
	if err != nil {
		AbortWithInternalServerErrorJSON(c, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (s *ServerApi) DeleteOperatingSystemById(c *gin.Context) {
	osId := c.Param("osId")

	// query database
	err := s.dataStore.DeleteOperatingSystemById(
		&repository.DeleteOperatingSystemByIdParams{
			Context: c,
			Id:      osId,
		},
	)
	if err != nil {
		AbortWithInternalServerErrorJSON(c, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
