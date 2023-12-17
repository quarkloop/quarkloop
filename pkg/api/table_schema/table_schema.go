package table_schema

import (
	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/service/table_schema"
)

type Api interface {
	ListTableSchemas(c *gin.Context)
	GetTableSchemaById(c *gin.Context)
	CreateTableSchema(c *gin.Context)
	UpdateTableSchemaById(c *gin.Context)
	DeleteTableSchemaById(c *gin.Context)
}

type TableSchemaApi struct {
	tableSchemaService table_schema.Service
}

func NewTableSchemaApi(service table_schema.Service) *TableSchemaApi {
	return &TableSchemaApi{
		tableSchemaService: service,
	}
}
