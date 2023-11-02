package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/quarkloop/quarkloop/pkg/db/model"
	"github.com/quarkloop/quarkloop/pkg/db/repository"
)

type GetOperatingSystemListResponse struct {
	ApiResponse
	Data []model.OperatingSystem `json:"data,omitempty"`
}

func (s *ServerApi) GetOperatingSystemList(c *gin.Context) {
	// query database
	osList, err := s.dataStore.ListOperatingSystems(
		&repository.ListOperatingSystemsParams{
			Context: c,
		},
	)
	if err != nil {
		AbortWithInternalServerErrorJSON(c, err)
		return
	}

	res := &GetOperatingSystemListResponse{
		ApiResponse: ApiResponse{
			Status:       http.StatusOK,
			StatusString: "OK",
		},
		Data: osList,
	}
	c.JSON(http.StatusOK, res)
}

type GetOperatingSystemByIdResponse struct {
	ApiResponse
	Data model.OperatingSystem `json:"data,omitempty"`
}

func (s *ServerApi) GetOperatingSystemById(c *gin.Context) {
	osId := c.Param("osId")

	// query database
	os, err := s.dataStore.GetOperatingSystemById(
		&repository.GetOperatingSystemByIdParams{
			Context: c,
			Id:      osId,
		},
	)
	if err != nil {
		AbortWithInternalServerErrorJSON(c, err)
		return
	}

	res := &GetOperatingSystemByIdResponse{
		ApiResponse: ApiResponse{
			Status:       http.StatusOK,
			StatusString: "OK",
		},
		Data: *os,
	}
	c.JSON(http.StatusOK, res)
}
