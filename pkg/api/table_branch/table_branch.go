package table_branch

import (
	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/service/table_branch"
)

type Api interface {
	ListTableBranches(c *gin.Context)
	GetTableBranchById(c *gin.Context)
}

type TableBranchApi struct {
	tableBranchService table_branch.Service
}

func NewTableBranchApi(service table_branch.Service) *TableBranchApi {
	return &TableBranchApi{
		tableBranchService: service,
	}
}
