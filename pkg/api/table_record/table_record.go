package table_record

import (
	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/service/table_record"
)

type Api interface {
	ListTableRecords(*gin.Context)
	GetTableRecordById(*gin.Context)
	CreateTableRecord(*gin.Context)
	UpdateTableRecordById(*gin.Context)
	DeleteTableRecordById(*gin.Context)
}

type TableRecordApi struct {
	tableRecordService table_record.Service
}

func NewTableRecordApi(service table_record.Service) *TableRecordApi {
	return &TableRecordApi{
		tableRecordService: service,
	}
}
