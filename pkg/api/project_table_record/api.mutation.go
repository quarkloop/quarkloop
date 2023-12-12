package table_record

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/project_table"
)

type CreateTableRecordUriParams struct {
	ProjectId string `uri:"projectId" binding:"required"`
}

type CreateTableRecordRequest struct {
	model.Table
}

type CreateTableRecordResponse struct {
	api.ApiResponse
	Data model.Table `json:"data,omitempty"`
}

func (s *TableRecordApi) CreateTableRecord(c *gin.Context) {
	uriParams := &CreateTableRecordUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	req := &CreateTableRecordRequest{}
	if err := c.BindJSON(req); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	// query service
	ws, err := s.tableRecord.CreateTable(
		&project_table.CreateTableParams{
			Context:   c,
			ProjectId: uriParams.ProjectId,
			Table:     &req.Table,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	res := &CreateTableRecordResponse{
		ApiResponse: api.ApiResponse{
			Status:       http.StatusCreated,
			StatusString: "Created",
		},
		Data: *ws,
	}
	c.JSON(http.StatusCreated, res)
}

type UpdateTableRecordByIdUriParams struct {
	ProjectId     string `uri:"projectId" binding:"required"`
	TableRecordId string `uri:"tableId" binding:"required"`
}

type UpdateTableRecordByIdRequest struct {
	model.Table
}

func (s *TableRecordApi) UpdateTableRecordById(c *gin.Context) {
	uriParams := &UpdateTableRecordByIdUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	req := &UpdateTableRecordByIdRequest{}
	if err := c.BindJSON(req); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	// query service
	err := s.tableRecord.UpdateTableById(
		&project_table.UpdateTableByIdParams{
			Context: c,
			TableId: uriParams.TableRecordId,
			Table:   &req.Table,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}

type DeleteTableRecordByIdUriParams struct {
	TableRecordId string `uri:"tableRecordId" binding:"required"`
}

func (s *TableRecordApi) DeleteTableRecordById(c *gin.Context) {
	uriParams := &DeleteTableRecordByIdUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	// query service
	err := s.tableRecord.DeleteTableById(
		&project_table.DeleteTableByIdParams{
			Context: c,
			TableId: uriParams.TableRecordId,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
