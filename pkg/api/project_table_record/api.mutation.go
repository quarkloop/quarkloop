package table_record

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/model"
	table_record "github.com/quarkloop/quarkloop/pkg/service/project_table_record"
)

type CreateTableRecordUriParams struct {
	ProjectId int    `uri:"projectId" binding:"required"`
	BranchId  int    `uri:"branchId" binding:"required"`
	TableType string `uri:"tableType" binding:"required"`
	RecordId  string `uri:"recordId" binding:"required"`
}

func (s *TableRecordApi) CreateTableRecord(c *gin.Context) {
	uriParams := &CreateTableRecordUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	var req interface{}
	if uriParams.TableType == "main" {
		req = model.MainRecordWithRelationCount{}
		if err := c.BindJSON(req); err != nil {
			api.AbortWithBadRequestJSON(c, err)
			return
		}
	} else if uriParams.TableType == "document" {
		req = model.DocumentRecord{}
		if err := c.BindJSON(req); err != nil {
			api.AbortWithBadRequestJSON(c, err)
			return
		}
	}

	// query service
	record, err := s.tableRecordService.CreateTableRecord(
		&table_record.CreateTableRecordParams{
			Context:   c,
			ProjectId: uriParams.ProjectId,
			BranchId:  uriParams.BranchId,
			TableType: uriParams.TableType,
			Record:    req,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	c.JSON(http.StatusCreated, record)
}

type UpdateTableRecordByIdUriParams struct {
	ProjectId int    `uri:"projectId" binding:"required"`
	BranchId  int    `uri:"branchId" binding:"required"`
	RecordId  string `uri:"recordId" binding:"required"`
	TableType string `uri:"tableType" binding:"required"`
}

func (s *TableRecordApi) UpdateTableRecordById(c *gin.Context) {
	uriParams := &UpdateTableRecordByIdUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	var req interface{}
	if uriParams.TableType == "main" {
		req = model.MainRecordWithRelationCount{}
		if err := c.BindJSON(req); err != nil {
			api.AbortWithBadRequestJSON(c, err)
			return
		}
	} else if uriParams.TableType == "document" {
		req = model.DocumentRecord{}
		if err := c.BindJSON(req); err != nil {
			api.AbortWithBadRequestJSON(c, err)
			return
		}
	}

	// query service
	err := s.tableRecordService.UpdateTableRecordById(
		&table_record.UpdateTableRecordByIdParams{
			Context:   c,
			ProjectId: uriParams.ProjectId,
			BranchId:  uriParams.BranchId,
			RecordId:  uriParams.RecordId,
			TableType: uriParams.TableType,
			Record:    req,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}

type DeleteTableRecordByIdUriParams struct {
	ProjectId int    `uri:"projectId" binding:"required"`
	BranchId  int    `uri:"branchId" binding:"required"`
	RecordId  string `uri:"recordId" binding:"required"`
	TableType string `uri:"tableType" binding:"required"`
}

func (s *TableRecordApi) DeleteTableRecordById(c *gin.Context) {
	uriParams := &DeleteTableRecordByIdUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	// query service
	err := s.tableRecordService.DeleteTableRecordById(
		&table_record.DeleteTableRecordByIdParams{
			Context:   c,
			ProjectId: uriParams.ProjectId,
			BranchId:  uriParams.BranchId,
			RecordId:  uriParams.RecordId,
			TableType: uriParams.TableType,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
