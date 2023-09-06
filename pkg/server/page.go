package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) HandleGetPageOps(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func (s *Server) HandleCallPageOp(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func (s *Server) HandleGetPage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func (s *Server) HandleGetPages(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func (s *Server) HandleCreatePage(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"status": "Created"})
}

func (s *Server) HandleUpdatePage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func (s *Server) HandleDeletePage(c *gin.Context) {
	c.JSON(http.StatusNoContent, gin.H{"status": "NoContent"})
}
