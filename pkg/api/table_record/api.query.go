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

func (s *TableRecordApi) ListTableRecords(c *gin.Context) {
	uriParams := &ListTableRecordsUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	// query service
	recordList, err := s.tableRecordService.ListTableRecords(
		&table_record.GetTableRecordListParams{
			Context:   c,
			ProjectId: uriParams.ProjectId,
			BranchId:  uriParams.BranchId,
			TableType: uriParams.TableType,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	c.JSON(http.StatusOK, &recordList)
}

type GetTableRecordByIdUriParams struct {
	ProjectId int    `uri:"projectId" binding:"required"`
	BranchId  int    `uri:"branchId" binding:"required"`
	RecordId  string `uri:"recordId" binding:"required"`
}

func (s *TableRecordApi) GetTableRecordById(c *gin.Context) {
	uriParams := &GetTableRecordByIdUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	// query service
	record, err := s.tableRecordService.GetTableRecordById(
		&table_record.GetTableRecordByIdParams{
			Context:   c,
			ProjectId: uriParams.ProjectId,
			BranchId:  uriParams.BranchId,
			RecordId:  uriParams.RecordId,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	c.JSON(http.StatusOK, record)
}
