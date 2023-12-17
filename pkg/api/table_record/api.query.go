package table_record

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/service/table_record"
)

type ListTableRecordsUriParams struct {
	ProjectId int    `uri:"projectId" binding:"required"`
	BranchId  int    `uri:"branchId" binding:"required"`
	TableType string `uri:"tableType" binding:"required"`
}

func (s *TableRecordApi) ListTableRecords(ctx *gin.Context) {
	uriParams := &ListTableRecordsUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	// query service
	recordList, err := s.tableRecordService.ListTableRecords(ctx, &table_record.GetTableRecordListParams{
		ProjectId: uriParams.ProjectId,
		BranchId:  uriParams.BranchId,
		TableType: uriParams.TableType,
	},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, &recordList)
}

type GetTableRecordByIdUriParams struct {
	ProjectId int    `uri:"projectId" binding:"required"`
	BranchId  int    `uri:"branchId" binding:"required"`
	RecordId  string `uri:"recordId" binding:"required"`
}

func (s *TableRecordApi) GetTableRecordById(ctx *gin.Context) {
	uriParams := &GetTableRecordByIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	// query service
	record, err := s.tableRecordService.GetTableRecordById(ctx, &table_record.GetTableRecordByIdParams{
		ProjectId: uriParams.ProjectId,
		BranchId:  uriParams.BranchId,
		RecordId:  uriParams.RecordId,
	},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, record)
}
