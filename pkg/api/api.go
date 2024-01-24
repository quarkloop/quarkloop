package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/store/repository"
)

type ServerApi struct {
	dataStore *repository.Repository
}

func NewServerApi(ds *repository.Repository) ServerApi {
	return ServerApi{dataStore: ds}
}

func AbortWithStatusJSON(ctx *gin.Context, code int, err error) {
	if e, ok := err.(*strconv.NumError); ok && e.Func == "ParseInt" {
		ctx.AbortWithStatusJSON(code, fmt.Errorf("field validation for '%s' failed", e.Num).Error())
		return
	}
	ctx.AbortWithStatusJSON(code, err.Error())
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

type Response interface {
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

func Error(status int, err error) Response {
	if err != nil {
		return &response{
			status: status,
			body:   err.Error(),
		}
	}

	return &response{
		status: status,
		body:   nil,
	}
}

func Success(status int, message any) Response {
	return &response{
		status: status,
		body:   message,
	}
}
