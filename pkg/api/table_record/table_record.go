package table_record

import (
	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/service/table_record"
)

type Api interface {
	ListTableRecords(c *gin.Context)
	GetTableRecordById(c *gin.Context)
	CreateTableRecord(c *gin.Context)
	UpdateTableRecordById(c *gin.Context)
	DeleteTableRecordById(c *gin.Context)
}

type TableRecordApi struct {
	tableRecordService table_record.Service
}

func NewTableRecordApi(service table_record.Service) *TableRecordApi {
	return &TableRecordApi{
		tableRecordService: service,
	}
}
