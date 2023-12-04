package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func NewTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.RedirectFixedPath = false
	router.RedirectTrailingSlash = false
	return router
}

func PerformRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	return w
}

func PerformRequestWithBody(r http.Handler, method, path string, body any) *httptest.ResponseRecorder {
	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(body)

	req, _ := http.NewRequest(method, path, payloadBuf)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	return w
}

// func SetUpServer() *Server {
// 	router := NewTestRouter()

// 	s := Server{
// 		Status: 0,
// 		Router: router,
// 	}
// 	s.BindHandlers()

// 	return &s
// }
