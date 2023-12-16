package table_schema

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/model"
	table_schema "github.com/quarkloop/quarkloop/pkg/service/project_table_schema"
)

type CreateTableSchemaUriParams struct {
	ProjectId int `uri:"projectId" binding:"required"`
}

type CreateTableSchemaRequest struct {
	model.TableSchema
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
	ws, err := s.tableSchemaService.CreateTableSchema(
		&table_schema.CreateTableSchemaParams{
			Context:   c,
			ProjectId: uriParams.ProjectId,
			Schema:    &req.TableSchema,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	c.JSON(http.StatusCreated, ws)
}

type UpdateTableSchemaByIdUriParams struct {
	ProjectId int    `uri:"projectId" binding:"required"`
	SchemaId  string `uri:"schemaId" binding:"required"`
}

type UpdateTableSchemaByIdRequest struct {
	model.TableSchema
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
	err := s.tableSchemaService.UpdateTableSchemaById(
		&table_schema.UpdateTableSchemaByIdParams{
			Context:   c,
			ProjectId: uriParams.ProjectId,
			SchemaId:  uriParams.SchemaId,
			Schema:    &req.TableSchema,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}

type DeleteTableSchemaByIdUriParams struct {
	ProjectId int    `uri:"projectId" binding:"required"`
	SchemaId  string `uri:"schemaId" binding:"required"`
}

func (s *TableSchemaApi) DeleteTableSchemaById(c *gin.Context) {
	uriParams := &DeleteTableSchemaByIdUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	// query service
	err := s.tableSchemaService.DeleteTableSchemaById(
		&table_schema.DeleteTableSchemaByIdParams{
			Context:   c,
			ProjectId: uriParams.ProjectId,
			SchemaId:  uriParams.SchemaId,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
