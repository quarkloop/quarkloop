package table_schema

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/project_table"
)

type CreateTableSchemaUriParams struct {
	ProjectId string `uri:"projectId" binding:"required"`
}

type CreateTableSchemaRequest struct {
	model.Table
}

type CreateTableSchemaResponse struct {
	api.ApiResponse
	Data model.Table `json:"data,omitempty"`
}

func (s *TableSchemaApi) CreateTableSchema(c *gin.Context) {
	uriParams := &CreateTableSchemaUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	req := &CreateTableSchemaRequest{}
	if err := c.BindJSON(req); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	// query service
	ws, err := s.tableSchema.CreateTable(
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

	res := &CreateTableSchemaResponse{
		ApiResponse: api.ApiResponse{
			Status:       http.StatusCreated,
			StatusString: "Created",
		},
		Data: *ws,
	}
	c.JSON(http.StatusCreated, res)
}

type UpdateTableSchemaByIdUriParams struct {
	ProjectId     string `uri:"projectId" binding:"required"`
	TableSchemaId string `uri:"tableId" binding:"required"`
}

type UpdateTableSchemaByIdRequest struct {
	model.Table
}

func (s *TableSchemaApi) UpdateTableSchemaById(c *gin.Context) {
	uriParams := &UpdateTableSchemaByIdUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	req := &UpdateTableSchemaByIdRequest{}
	if err := c.BindJSON(req); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	// query service
	err := s.tableSchema.UpdateTableById(
		&project_table.UpdateTableByIdParams{
			Context: c,
			TableId: uriParams.TableSchemaId,
			Table:   &req.Table,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}

type DeleteTableSchemaByIdUriParams struct {
	TableSchemaId string `uri:"tableSchemaId" binding:"required"`
}

func (s *TableSchemaApi) DeleteTableSchemaById(c *gin.Context) {
	uriParams := &DeleteTableSchemaByIdUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	// query service
	err := s.tableSchema.DeleteTableById(
		&project_table.DeleteTableByIdParams{
			Context: c,
			TableId: uriParams.TableSchemaId,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
