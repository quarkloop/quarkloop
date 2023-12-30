package table_schema

import (
	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/service/table_schema"
)

type Api interface {
	ListTableSchemas(*gin.Context)
	GetTableSchemaById(*gin.Context)
	CreateTableSchema(*gin.Context)
	UpdateTableSchemaById(*gin.Context)
	DeleteTableSchemaById(*gin.Context)
}

type TableSchemaApi struct {
	tableSchemaService table_schema.Service
}

func NewTableSchemaApi(service table_schema.Service) *TableSchemaApi {
	return &TableSchemaApi{
		tableSchemaService: service,
	}
}
