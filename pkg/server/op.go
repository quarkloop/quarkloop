package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/ops"
)

// type OpArg struct {
// 	AppID      string          `json:"appId" binding:"required"`
// 	InstanceID string          `json:"instanceId" binding:"required"`
// 	Args       json.RawMessage `json:"args" binding:"required"`
// }

type OpCallRequestPayload struct {
	Op   string          `json:"op" binding:"required"`
	Args json.RawMessage `json:"args" binding:"required"`
}

type BatchOpCallRequestPayload = []OpCallRequestPayload

// type OpCallResponsePayload struct {
// 	Status       int         `json:"status,omitempty"`
// 	StatusString string      `json:"statusText,omitempty"`
// 	Error        error       `json:"error,omitempty"`
// 	ErrorString  string      `json:"errorString,omitempty"`
// 	AppInstance  interface{} `json:"appInstance,omitempty"`
// }

func (s *Server) HandleCallOp(c *gin.Context) {
	payload := &OpCallRequestPayload{}

	if err := c.BindJSON(payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, AppResponsePayload{
			Status:       http.StatusBadRequest,
			StatusString: "BadRequest",
			Error:        err,
			ErrorString:  fmt.Sprintf("[BindJSON] %s", err.Error()),
		})
		return
	}

	catalog, err := ops.FindOp(payload.Op, payload.Args)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, AppResponsePayload{
			Status:       http.StatusNotFound,
			StatusString: "NotFound",
			Error:        err,
			ErrorString:  fmt.Sprintf("[FindOp] %s", err.Error()),
		})
		return
	}

	res, err := catalog.Exec()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusExpectationFailed, AppResponsePayload{
			Status:       http.StatusExpectationFailed,
			StatusString: "ExpectationFailed",
			Error:        err,
			ErrorString:  fmt.Sprintf("[Exec] %s", err.Error()),
		})
		return
	}
	fmt.Printf("\n--------- %+v\n", res)

	c.JSON(http.StatusOK, res)
}
