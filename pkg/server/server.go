package server

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
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

type Server struct {
	Status int
	Router *gin.Engine
}

func NewDefaultServer() Server {
	router := gin.Default()
	router.RedirectFixedPath = false
	router.RedirectTrailingSlash = false

	serve := Server{
		Status: 0,
		Router: router,
	}

	return serve
}

func NewRequest() []byte {
	requestURL := fmt.Sprintf("http://localhost:%d", 3000)
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		os.Exit(1)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("client: response body: %s\n", resBody)

	return resBody
}

func (s *Server) BindHandlers() {
	router := s.Router

	app := router.Group("/api/v1/apps")
	app.POST("/call", s.HandleCallOp)
}
