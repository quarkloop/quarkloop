package server

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/types"
)

func (s *Server) HandleGetApp(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func (s *Server) HandleCreateApp(c *gin.Context) {
	app := &types.App{}

	err := c.BindJSON(app)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "BadRequest"})
		return
	}

	marshalled, err := json.Marshal(app)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "Marshalling failed"})
		return
	}

	req, err := http.NewRequest("POST", "http://localhost:3001/db", bytes.NewBuffer(marshalled))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "InternalServerError"})
		return
	}

	client := http.Client{Timeout: 10 * time.Second}
	res, err := client.Do(req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "Failed to send request"})
		return
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "Failed to read body"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "Created",
		"message": string(resBody),
	})
}

func (s *Server) HandleUpdateApp(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func (s *Server) HandleDeleteApp(c *gin.Context) {
	c.JSON(http.StatusNoContent, gin.H{"status": "NoContent"})
}
