package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) HandleGetForm(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func (s *Server) HandleGetForms(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func (s *Server) HandleCreateForm(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"status": "Created"})
}

func (s *Server) HandleUpdateForm(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func (s *Server) HandleDeleteForm(c *gin.Context) {
	c.JSON(http.StatusNoContent, gin.H{"status": "NoContent"})
}
