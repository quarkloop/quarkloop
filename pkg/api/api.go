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

type ErrorResponse struct {
	Status       int    `json:"status"`
	StatusString string `json:"statusString"`
	Error        error  `json:"error,omitempty"`
	ErrorString  string `json:"errorString,omitempty"`
}

func AbortWithBadRequestJSON(ctx *gin.Context, err error) {
	response := ErrorResponse{
		Status:       http.StatusBadRequest,
		StatusString: "BadRequest",
		Error:        err,
		ErrorString:  fmt.Sprintf("[BindJSON] %s", err.Error()),
	}

	ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
}

func AbortWithInternalServerErrorJSON(ctx *gin.Context, err error) {
	response := ErrorResponse{
		Status:       http.StatusInternalServerError,
		StatusString: "InternalServerError",
		Error:        err,
		ErrorString:  fmt.Sprintf("[BindJSON] %s", err.Error()),
	}

	ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
}

type Reponse interface {
	Status() int
	Body() any
}

type response struct {
	status int
	body   any
}

func (r *response) Status() int {
	return r.status
}

func (r *response) Body() any {
	return r.body
}

func Error(status int, err error) Reponse {
	return &response{
		status: status,
		body:   err.Error(),
	}
}

func Success(status int, message any) Reponse {
	return &response{
		status: status,
		body:   message,
	}
}
