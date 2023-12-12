package project_table

import (
	"github.com/gin-gonic/gin"
	table "github.com/quarkloop/quarkloop/pkg/service/project_table"
)

type Api interface {
	ListTableRecords(c *gin.Context)
	GetTableRecordById(c *gin.Context)
	CreateProjectTable(c *gin.Context)
	UpdateProjectTableById(c *gin.Context)
	DeleteProjectTableById(c *gin.Context)
}

type ProjectTableApi struct {
	projectTable table.Service
}

func NewProjectTableApi(service table.Service) *ProjectTableApi {
	return &ProjectTableApi{
		projectTable: service,
	}
}
