package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/db"
	"github.com/quarkloop/quarkloop/pkg/types"
)

type AppInstanceRequestPayload struct {
	AppInstance types.AppInstance `json:"appInstance"`
}

type AppInstanceResponsePayload struct {
	Status       int         `json:"status,omitempty"`
	StatusString string      `json:"statusText,omitempty"`
	Error        error       `json:"error,omitempty"`
	ErrorString  string      `json:"errorString,omitempty"`
	AppInstance  interface{} `json:"appInstance,omitempty"`
}

func (s *Server) HandleGetAppInstance(c *gin.Context) {
	payload := AppInstanceResponsePayload{
		Status:       http.StatusOK,
		StatusString: "OK",
		AppInstance:  types.AppInstance{Name: "UploadFileToS3"},
	}

	c.JSON(http.StatusOK, payload)
}

func (s *Server) HandleGetAppInstances(c *gin.Context) {
	payload := AppInstanceResponsePayload{
		Status:       http.StatusOK,
		StatusString: "OK",
		AppInstance:  types.AppInstance{Name: "UploadFileToS3"},
	}

	c.JSON(http.StatusOK, payload)
}

func (s *Server) HandleCreateAppInstance(c *gin.Context) {
	payload := &AppInstanceRequestPayload{}

	if err := c.BindJSON(payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, AppInstanceResponsePayload{
			Status:       http.StatusBadRequest,
			StatusString: "BadRequest",
			Error:        err,
			ErrorString:  fmt.Sprintf("[BindJSON] %s", err.Error()),
		})
		return
	}

	marshalled, err := json.Marshal(payload.AppInstance)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, AppInstanceResponsePayload{
			Status:       http.StatusBadRequest,
			StatusString: "BadRequest",
			Error:        err,
			ErrorString:  fmt.Sprintf("[Marshal] %s", err.Error()),
		})
		return
	}

	res, err := s.Database(c, "POST", "http://localhost:3000/api/v1/tables/appInstance", marshalled, nil)
	if err != nil {
		return
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, AppInstanceResponsePayload{
			Status:       http.StatusBadRequest,
			StatusString: "BadRequest",
			Error:        err,
			ErrorString:  fmt.Sprintf("[BufferReader] %s", err.Error()),
		})
		return
	}

	databasePayload := db.DatabaseResponsePayload{}
	if err := json.Unmarshal(resBody, &databasePayload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, AppInstanceResponsePayload{
			Status:       http.StatusBadRequest,
			StatusString: "BadRequest",
			Error:        err,
			ErrorString:  fmt.Sprintf("[Unmarshal] %s", err.Error()),
		})
		return
	}

	resPayload := AppInstanceResponsePayload{
		Status:       http.StatusCreated,
		StatusString: "Created",
		AppInstance:  databasePayload.Database.AppInstance.Records,
	}

	c.JSON(http.StatusCreated, resPayload)
}

func (s *Server) HandleUpdateAppInstance(c *gin.Context) {
	appId := c.Param("appId")
	appInstanceId := c.Param("appInstanceId")

	payload := &AppInstanceRequestPayload{}
	payload.AppInstance.Id = appInstanceId
	payload.AppInstance.AppId = appId

	if err := c.BindJSON(payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, AppInstanceResponsePayload{
			Status:       http.StatusBadRequest,
			StatusString: "BadRequest",
			Error:        err,
			ErrorString:  fmt.Sprintf("[BindJSON] %s", err.Error()),
		})
		return
	}

	marshalled, err := json.Marshal(payload.AppInstance)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, AppInstanceResponsePayload{
			Status:       http.StatusBadRequest,
			StatusString: "BadRequest",
			Error:        err,
			ErrorString:  fmt.Sprintf("[Marshal] %s", err.Error()),
		})
		return
	}

	res, err := s.Database(c, "PUT", "http://localhost:3000/api/v1/tables/appInstance", marshalled, nil)
	if err != nil {
		return
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, AppInstanceResponsePayload{
			Status:       http.StatusBadRequest,
			StatusString: "BadRequest",
			Error:        err,
			ErrorString:  fmt.Sprintf("[BufferReader] %s", err.Error()),
		})
		return
	}

	databasePayload := db.DatabaseResponsePayload{}
	if err := json.Unmarshal(resBody, &databasePayload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, AppInstanceResponsePayload{
			Status:       http.StatusBadRequest,
			StatusString: "BadRequest",
			Error:        err,
			ErrorString:  fmt.Sprintf("[Unmarshal] %s", err.Error()),
		})
		return
	}

	resPayload := AppInstanceResponsePayload{
		Status:       http.StatusOK,
		StatusString: "OK",
		AppInstance:  databasePayload.Database.AppInstance.Records,
	}

	c.JSON(http.StatusOK, resPayload)
}

func (s *Server) HandleDeleteAppInstance(c *gin.Context) {
	appId := c.Param("appId")
	appInstanceId := c.Param("appInstanceId")

	q := url.Values{}
	q.Add("id", appInstanceId)
	q.Add("appId", appId)

	res, err := s.Database(c, "DELETE", "http://localhost:3000/api/v1/tables/appInstance", nil, &q)
	if err != nil {
		return
	}
	defer res.Body.Close()

	c.JSON(res.StatusCode, nil)
}
