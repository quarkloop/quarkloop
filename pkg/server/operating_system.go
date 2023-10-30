package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/quarkloop/quarkloop/pkg/db/model"
	"github.com/quarkloop/quarkloop/pkg/db/repository"
)

type GetOperatingSystemListResponse struct {
	Data []model.OperatingSystem `json:"data"`
}
type GetOperatingSystemByIdResponse struct{}
type CreateOperatingSystemRequest struct{ model.OperatingSystem }
type CreateOperatingSystemResponse struct {
	Status       int                   `json:"status"`
	StatusString string                `json:"statusString"`
	Data         model.OperatingSystem `json:"data"`
}
type UpdateOperatingSystemByIdRequest struct{ model.OperatingSystem }

func (s *Server) GetOperatingSystemList(c *gin.Context) {
	app, err := s.dataStore.ListOperatingSystems(
		&repository.ListOperatingSystemsParams{
			Context: c,
		},
	)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			AppResponsePayload{
				Status:       http.StatusInternalServerError,
				StatusString: "InternalServerError",
				Error:        err,
				ErrorString:  fmt.Sprintf("[BindJSON] %s", err.Error()),
			})
		return
	}

	res := &GetOperatingSystemListResponse{Data: app}
	c.JSON(http.StatusOK, res)
}

func (s *Server) CreateOperatingSystem(c *gin.Context) {
	req := &CreateOperatingSystemRequest{}
	if err := c.BindJSON(req); err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			AppResponsePayload{
				Status:       http.StatusBadRequest,
				StatusString: "BadRequest",
				Error:        err,
				ErrorString:  fmt.Sprintf("[BindJSON] %s", err.Error()),
			})
		return
	}

	// query database
	app, err := s.dataStore.CreateOperatingSystem(
		&repository.CreateOperatingSystemParams{
			Context: c,
			Os:      req.OperatingSystem,
		},
	)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			AppResponsePayload{
				Status:       http.StatusInternalServerError,
				StatusString: "InternalServerError",
				Error:        err,
				ErrorString:  fmt.Sprintf("[BindJSON] %s", err.Error()),
			})
		return
	}

	fmt.Printf("Req: %v", app)

	res := &CreateOperatingSystemResponse{
		Status: http.StatusCreated,
		Data:   *app,
	}
	c.JSON(http.StatusCreated, res)
}

func (s *Server) GetOperatingSystemById(c *gin.Context) {
	osId := c.Param("osId")
	_ = osId

	// query database

	res := &GetOperatingSystemByIdResponse{}
	c.JSON(http.StatusOK, res)
}

func (s *Server) UpdateOperatingSystemById(c *gin.Context) {
	req := &UpdateOperatingSystemByIdRequest{}
	if err := c.BindJSON(req); err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			AppResponsePayload{
				Status:       http.StatusBadRequest,
				StatusString: "BadRequest",
				Error:        err,
				ErrorString:  fmt.Sprintf("[BindJSON] %s", err.Error()),
			})
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
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			AppResponsePayload{
				Status:       http.StatusInternalServerError,
				StatusString: "InternalServerError",
				Error:        err,
				ErrorString:  fmt.Sprintf("[BindJSON] %s", err.Error()),
			})
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (s *Server) DeleteOperatingSystemById(c *gin.Context) {
	osId := c.Param("osId")

	// query database
	err := s.dataStore.DeleteOperatingSystemById(
		&repository.DeleteOperatingSystemByIdParams{
			Context: c,
			Id:      osId,
		},
	)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			AppResponsePayload{
				Status:       http.StatusInternalServerError,
				StatusString: "InternalServerError",
				Error:        err,
				ErrorString:  fmt.Sprintf("[BindJSON] %s", err.Error()),
			})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
