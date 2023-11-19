package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/store/repository"
)

type ServerApi struct {
	dataStore *repository.Repository
}

func NewServerApi(ds *repository.Repository) ServerApi {
	return ServerApi{dataStore: ds}
}

type ApiResponse struct {
	Status       int    `json:"status"`
	StatusString string `json:"statusString"`
	Error        error  `json:"error,omitempty"`
	ErrorString  string `json:"errorString,omitempty"`
}

func AbortWithBadRequestJSON(c *gin.Context, err error) {
	response := ApiResponse{
		Status:       http.StatusBadRequest,
		StatusString: "BadRequest",
		Error:        err,
		ErrorString:  fmt.Sprintf("[BindJSON] %s", err.Error()),
	}

	c.AbortWithStatusJSON(http.StatusBadRequest, response)
}

func AbortWithInternalServerErrorJSON(c *gin.Context, err error) {
	response := ApiResponse{
		Status:       http.StatusInternalServerError,
		StatusString: "InternalServerError",
		Error:        err,
		ErrorString:  fmt.Sprintf("[BindJSON] %s", err.Error()),
	}

	c.AbortWithStatusJSON(http.StatusInternalServerError, response)
}
