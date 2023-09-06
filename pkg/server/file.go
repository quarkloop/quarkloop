package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) HandleGetFileOps(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func (s *Server) HandleCallFileOp(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func (s *Server) HandleGetFile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func (s *Server) HandleGetFiles(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func (s *Server) HandleCreateFile(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"status": "Created"})
}

func (s *Server) HandleUpdateFile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func (s *Server) HandleDeleteFile(c *gin.Context) {
	c.JSON(http.StatusNoContent, gin.H{"status": "NoContent"})
}
