package server

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

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
	// App
	{
		app.GET("", s.HandleGetApps)
		app.POST("", s.HandleCreateApp)

		app.GET("/:appId", s.HandleGetApp)
		app.PUT("/:appId", s.HandleUpdateApp)
		app.DELETE("/:appId", s.HandleDeleteApp)
	}
	// App Instance
	{
		app.GET("/:appId/instances", s.HandleGetInstances)
		app.POST("/:appId/instances", s.HandleCreateInstance)

		app.GET("/:appId/instances/:instanceId", s.HandleGetInstance)
		app.PUT("/:appId/instances/:instanceId", s.HandleUpdateInstance)
		app.DELETE("/:appId/instances/:instanceId", s.HandleDeleteInstance)
	}
	// AppFileSettings
	{
		app.GET("/:appId/settings/files", s.HandleGetFileSettings)
		app.POST("/:appId/settings/files", s.HandleCreateFileSettings)
		app.PUT("/:appId/settings/files", s.HandleUpdateFileSettings)
		app.DELETE("/:appId/settings/files", s.HandleDeleteFileSettings)
	}

	// AppThreadSettings
	{
		app.GET("/:appId/settings/threads", s.HandleGetThreadSettings)
		app.POST("/:appId/settings/threads", s.HandleCreateThreadSettings)
		app.PUT("/:appId/settings/threads", s.HandleUpdateThreadSettings)
		app.DELETE("/:appId/settings/threads", s.HandleDeleteThreadSettings)
	}

	// AppFormSettings
	{
		app.GET("/:appId/settings/forms", s.HandleGetFormSettings)
		app.POST("/:appId/settings/forms", s.HandleCreateFormSettings)
		app.PUT("/:appId/settings/forms", s.HandleUpdateFormSettings)
		app.DELETE("/:appId/settings/forms", s.HandleDeleteFormSettings)
	}
	// AppPageSettings
	{
		app.GET("/:appId/settings/pages", s.HandleGetPageSettings)
		app.POST("/:appId/settings/pages", s.HandleCreatePageSettings)
		app.PUT("/:appId/settings/pages", s.HandleUpdatePageSettings)
		app.DELETE("/:appId/settings/pages", s.HandleDeletePageSettings)
	}
	// AppThread
	{
		app.GET("/:appId/instances/:instanceId/thread", s.HandleGetThreadOps)
		app.POST("/:appId/instances/:instanceId/thread", s.HandleCallThreadOp)

		// app.GET("/:appId/instances/:instanceId/threads", s.HandleGetThreads)
		// app.POST("/:appId/instances/:instanceId/threads", s.HandleCreateThread)

		// app.GET("/:appId/instances/:instanceId/threads/:threadId", s.HandleGetThread)
		// app.PUT("/:appId/instances/:instanceId/threads/:threadId", s.HandleUpdateThread)
		// app.DELETE("/:appId/instances/:instanceId/threads/:threadId", s.HandleDeleteThread)
	}
	// AppFile
	{
		app.GET("/:appId/instances/:instanceId/file", s.HandleGetFileOps)
		app.POST("/:appId/instances/:instanceId/file", s.HandleCallFileOp)

		// app.GET("/:appId/instances/:instanceId/files", s.HandleGetFiles)
		// app.POST("/:appId/instances/:instanceId/files", s.HandleCreateFile)

		// app.GET("/:appId/instances/:instanceId/files/:fileId", s.HandleGetFile)
		// app.PUT("/:appId/instances/:instanceId/files/:fileId", s.HandleUpdateFile)
		// app.DELETE("/:appId/instances/:instanceId/files/:fileId", s.HandleDeleteFile)
	}
	// AppForm
	{
		app.GET("/:appId/instances/:instanceId/form", s.HandleGetFormOps)
		app.POST("/:appId/instances/:instanceId/form", s.HandleCallFormOp)

		// app.GET("/:appId/instances/:instanceId/forms", s.HandleGetForms)
		// app.POST("/:appId/instances/:instanceId/forms", s.HandleCreateForm)

		// app.GET("/:appId/instances/:instanceId/forms/:formId", s.HandleGetForm)
		// app.PUT("/:appId/instances/:instanceId/forms/:formId", s.HandleUpdateForm)
		// app.DELETE("/:appId/instances/:instanceId/forms/:formId", s.HandleDeleteForm)
	}
	// AppPage
	{
		app.GET("/:appId/instances/:instanceId/page", s.HandleGetPageOps)
		app.POST("/:appId/instances/:instanceId/page", s.HandleCallPageOp)

		// app.GET("/:appId/instances/:instanceId/pages", s.HandleGetPages)
		// app.POST("/:appId/instances/:instanceId/pages", s.HandleCreatePage)

		// app.GET("/:appId/instances/:instanceId/pages/:pageId", s.HandleGetPage)
		// app.PUT("/:appId/instances/:instanceId/pages/:pageId", s.HandleUpdatePage)
		// app.DELETE("/:appId/instances/:instanceId/pages/:pageId", s.HandleDeletePage)
	}
}

func (s *Server) Database(
	c *gin.Context,
	method, path string,
	buf []byte,
	queryParams *url.Values,
) (*http.Response, error) {
	req, err := http.NewRequest(method, path, bytes.NewBuffer(buf))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, AppResponsePayload{
			Status:       http.StatusInternalServerError,
			StatusString: "InternalServerError",
			Error:        err,
			ErrorString:  fmt.Sprintf("[NewRequest] %s", err.Error()),
		})
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	if queryParams != nil {
		req.URL.RawQuery = queryParams.Encode()
	}

	client := http.Client{Timeout: 10 * time.Second}
	res, err := client.Do(req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, AppResponsePayload{
			Status:       http.StatusBadRequest,
			StatusString: "BadRequest",
			Error:        err,
			ErrorString:  fmt.Sprintf("[Client] %s", err.Error()),
		})
		return nil, err
	}

	return res, nil
}
