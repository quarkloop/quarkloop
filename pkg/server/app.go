package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/db/api"
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

	res, err := s.Database(c, "POST", "http://localhost:3000/api/v1/tables/app", marshalled, nil)
	if err != nil {
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

	fmt.Printf("\n*********** %s", string(resBody))

	databasePayload := api.DatabaseResponsePayload{}
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
		Status:       res.StatusCode,
		StatusString: res.Status,
		App:          databasePayload.Database.App.Records,
	}

	c.JSON(res.StatusCode, resPayload)
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

	res, err := s.Database(c, "PUT", "http://localhost:3000/api/v1/tables/app", marshalled, nil)
	if err != nil {
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

	databasePayload := api.DatabaseResponsePayload{}
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
		Status:       res.StatusCode,
		StatusString: res.Status,
		App:          databasePayload.Database.App.Records,
	}

	c.JSON(res.StatusCode, resPayload)
}

func (s *Server) HandleDeleteApp(c *gin.Context) {
	appId := c.Param("appId")

	q := url.Values{}
	q.Add("id", appId)

	res, err := s.Database(c, "DELETE", "http://localhost:3000/api/v1/tables/app", nil, &q)
	if err != nil {
		return
	}
	defer res.Body.Close()

	c.JSON(res.StatusCode, nil)
}
