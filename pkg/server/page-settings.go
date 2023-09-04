package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) HandleGetPageSettings(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func (s *Server) HandleCreatePageSettings(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"status": "Created"})
}

func (s *Server) HandleUpdatePageSettings(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func (s *Server) HandleDeletePageSettings(c *gin.Context) {
	c.JSON(http.StatusNoContent, gin.H{"status": "NoContent"})
}
