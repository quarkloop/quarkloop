package table_record

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/service/table_record"
)

type CreateTableRecordUriParams struct {
	ProjectId int    `uri:"projectId" binding:"required"`
	BranchId  int    `uri:"branchId" binding:"required"`
	TableType string `uri:"tableType" binding:"required"`
	RecordId  string `uri:"recordId" binding:"required"`
}

func (s *TableRecordApi) CreateTableRecord(ctx *gin.Context) {
	uriParams := &CreateTableRecordUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	var req interface{}
	if uriParams.TableType == "main" {
		req = table_record.MainRecordWithRelationCount{}
		if err := ctx.ShouldBindJSON(req); err != nil {
			api.AbortWithBadRequestJSON(ctx, err)
			return
		}
	} else if uriParams.TableType == "document" {
		req = table_record.DocumentRecord{}
		if err := ctx.ShouldBindJSON(req); err != nil {
			api.AbortWithBadRequestJSON(ctx, err)
			return
		}
	}

	// query service
	record, err := s.tableRecordService.CreateTableRecord(ctx, &table_record.CreateTableRecordParams{
		ProjectId: uriParams.ProjectId,
		BranchId:  uriParams.BranchId,
		TableType: uriParams.TableType,
		Record:    req,
	},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, record)
}

type UpdateTableRecordByIdUriParams struct {
	ProjectId int    `uri:"projectId" binding:"required"`
	BranchId  int    `uri:"branchId" binding:"required"`
	RecordId  string `uri:"recordId" binding:"required"`
	TableType string `uri:"tableType" binding:"required"`
}

func (s *TableRecordApi) UpdateTableRecordById(ctx *gin.Context) {
	uriParams := &UpdateTableRecordByIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	var req interface{}
	if uriParams.TableType == "main" {
		req = table_record.MainRecordWithRelationCount{}
		if err := ctx.ShouldBindJSON(req); err != nil {
			api.AbortWithBadRequestJSON(ctx, err)
			return
		}
	} else if uriParams.TableType == "document" {
		req = table_record.DocumentRecord{}
		if err := ctx.ShouldBindJSON(req); err != nil {
			api.AbortWithBadRequestJSON(ctx, err)
			return
		}
	}

	// query service
	err := s.tableRecordService.UpdateTableRecordById(ctx, &table_record.UpdateTableRecordByIdParams{
		ProjectId: uriParams.ProjectId,
		BranchId:  uriParams.BranchId,
		RecordId:  uriParams.RecordId,
		TableType: uriParams.TableType,
		Record:    req,
	},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

type DeleteTableRecordByIdUriParams struct {
	ProjectId int    `uri:"projectId" binding:"required"`
	BranchId  int    `uri:"branchId" binding:"required"`
	RecordId  string `uri:"recordId" binding:"required"`
	TableType string `uri:"tableType" binding:"required"`
}

func (s *TableRecordApi) DeleteTableRecordById(ctx *gin.Context) {
	uriParams := &DeleteTableRecordByIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	// query service
	err := s.tableRecordService.DeleteTableRecordById(ctx, &table_record.DeleteTableRecordByIdParams{
		ProjectId: uriParams.ProjectId,
		BranchId:  uriParams.BranchId,
		RecordId:  uriParams.RecordId,
		TableType: uriParams.TableType,
	},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
