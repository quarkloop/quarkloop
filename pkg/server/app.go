package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/db"
	"github.com/quarkloop/quarkloop/pkg/types"
)

type AppRequestPayload struct {
	App types.App `json:"app"`
}

type AppResponsePayload struct {
	Status       int         `json:"status,omitempty"`
	StatusString string      `json:"statusText,omitempty"`
	Error        error       `json:"error,omitempty"`
	ErrorString  string      `json:"errorString,omitempty"`
	App          interface{} `json:"app,omitempty"`
}

func (s *Server) HandleGetApp(c *gin.Context) {
	payload := AppResponsePayload{
		Status:       http.StatusOK,
		StatusString: "OK",
		App:          types.App{Name: "UploadFileToS3"},
	}

	c.JSON(http.StatusOK, payload)
}

func (s *Server) HandleGetApps(c *gin.Context) {
	payload := AppResponsePayload{
		Status:       http.StatusOK,
		StatusString: "OK",
		App:          types.App{Name: "UploadFileToS3"},
	}

	c.JSON(http.StatusOK, payload)
}

func (s *Server) HandleCreateApp(c *gin.Context) {
	payload := &AppRequestPayload{}

	if err := c.BindJSON(payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, AppResponsePayload{
			Status:       http.StatusBadRequest,
			StatusString: "BadRequest",
			Error:        err,
			ErrorString:  fmt.Sprintf("[BindJSON] %s", err.Error()),
		})
		return
	}

	marshalled, err := json.Marshal(payload.App)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, AppResponsePayload{
			Status:       http.StatusBadRequest,
			StatusString: "BadRequest",
			Error:        err,
			ErrorString:  fmt.Sprintf("[Marshal] %s", err.Error()),
		})
		return
	}

	req, err := http.NewRequest("POST", "http://localhost:3000/api/v1/tables/app", bytes.NewBuffer(marshalled))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, AppResponsePayload{
			Status:       http.StatusInternalServerError,
			StatusString: "InternalServerError",
			Error:        err,
			ErrorString:  fmt.Sprintf("[NewRequest] %s", err.Error()),
		})
		return
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := http.Client{Timeout: 10 * time.Second}
	res, err := client.Do(req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, AppResponsePayload{
			Status:       http.StatusBadRequest,
			StatusString: "BadRequest",
			Error:        err,
			ErrorString:  fmt.Sprintf("[Client] %s", err.Error()),
		})
		return
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, AppResponsePayload{
			Status:       http.StatusBadRequest,
			StatusString: "BadRequest",
			Error:        err,
			ErrorString:  fmt.Sprintf("[BufferReader] %s", err.Error()),
		})
		return
	}

	databasePayload := db.DatabaseResponsePayload{}
	if err := json.Unmarshal(resBody, &databasePayload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, AppResponsePayload{
			Status:       http.StatusBadRequest,
			StatusString: "BadRequest",
			Error:        err,
			ErrorString:  fmt.Sprintf("[Unmarshal] %s", err.Error()),
		})
		return
	}

	resPayload := AppResponsePayload{
		Status:       http.StatusCreated,
		StatusString: "Created",
		App:          databasePayload.Database.App.Records,
	}

	c.JSON(http.StatusCreated, resPayload)
}

func (s *Server) HandleUpdateApp(c *gin.Context) {
	appId := c.Param("appId")

	payload := &AppRequestPayload{}
	payload.App.Id = appId

	if err := c.BindJSON(payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, AppResponsePayload{
			Status:       http.StatusBadRequest,
			StatusString: "BadRequest",
			Error:        err,
			ErrorString:  fmt.Sprintf("[BindJSON] %s", err.Error()),
		})
		return
	}

	marshalled, err := json.Marshal(payload.App)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, AppResponsePayload{
			Status:       http.StatusBadRequest,
			StatusString: "BadRequest",
			Error:        err,
			ErrorString:  fmt.Sprintf("[Marshal] %s", err.Error()),
		})
		return
	}

	req, err := http.NewRequest("PUT", "http://localhost:3000/api/v1/tables/app", bytes.NewBuffer(marshalled))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, AppResponsePayload{
			Status:       http.StatusInternalServerError,
			StatusString: "InternalServerError",
			Error:        err,
			ErrorString:  fmt.Sprintf("[NewRequest] %s", err.Error()),
		})
		return
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := http.Client{Timeout: 10 * time.Second}
	res, err := client.Do(req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, AppResponsePayload{
			Status:       http.StatusBadRequest,
			StatusString: "BadRequest",
			Error:        err,
			ErrorString:  fmt.Sprintf("[Client] %s", err.Error()),
		})
		return
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, AppResponsePayload{
			Status:       http.StatusBadRequest,
			StatusString: "BadRequest",
			Error:        err,
			ErrorString:  fmt.Sprintf("[BufferReader] %s", err.Error()),
		})
		return
	}

	databasePayload := db.DatabaseResponsePayload{}
	if err := json.Unmarshal(resBody, &databasePayload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, AppResponsePayload{
			Status:       http.StatusBadRequest,
			StatusString: "BadRequest",
			Error:        err,
			ErrorString:  fmt.Sprintf("[Unmarshal] %s", err.Error()),
		})
		return
	}

	resPayload := AppResponsePayload{
		Status:       http.StatusOK,
		StatusString: "OK",
		App:          databasePayload.Database.App.Records,
	}

	c.JSON(http.StatusOK, resPayload)
}

func (s *Server) HandleDeleteApp(c *gin.Context) {
	appId := c.Param("appId")

	req, err := http.NewRequest("DELETE", "http://localhost:3000/api/v1/tables/app", nil)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, AppResponsePayload{
			Status:       http.StatusInternalServerError,
			StatusString: "InternalServerError",
			Error:        err,
			ErrorString:  fmt.Sprintf("[NewRequest] %s", err.Error()),
		})
		return
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	q := req.URL.Query()
	q.Add("id", appId)
	req.URL.RawQuery = q.Encode()

	client := http.Client{Timeout: 10 * time.Second}
	res, err := client.Do(req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, AppResponsePayload{
			Status:       http.StatusBadRequest,
			StatusString: "BadRequest",
			Error:        err,
			ErrorString:  fmt.Sprintf("[Client] %s", err.Error()),
		})
		return
	}
	defer res.Body.Close()

	c.JSON(res.StatusCode, nil)
}
