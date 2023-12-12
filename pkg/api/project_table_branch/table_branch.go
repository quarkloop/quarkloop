package table_branch

import (
	"github.com/gin-gonic/gin"
	table_branch "github.com/quarkloop/quarkloop/pkg/service/project_table_branch"
)

type Api interface {
	ListTableRecords(c *gin.Context)
	GetTableRecordById(c *gin.Context)
}

type TableBranchApi struct {
	tableBranch table_branch.Service
}

func NewTableBranchApi(service table_branch.Service) *TableBranchApi {
	return &TableBranchApi{
		tableBranch: service,
	}
}
