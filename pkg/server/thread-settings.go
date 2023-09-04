package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) HandleGetThreadSettings(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func (s *Server) HandleCreateThreadSettings(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"status": "Created"})
}

func (s *Server) HandleUpdateThreadSettings(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func (s *Server) HandleDeleteThreadSettings(c *gin.Context) {
	c.JSON(http.StatusNoContent, gin.H{"status": "NoContent"})
}
