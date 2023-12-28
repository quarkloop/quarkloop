package table_schema

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/service/table_schema"
)

type CreateTableSchemaUriParams struct {
	ProjectId int `uri:"projectId" binding:"required"`
}

type CreateTableSchemaRequest struct {
	table_schema.TableSchema
}

func (s *TableSchemaApi) CreateTableSchema(ctx *gin.Context) {
	uriParams := &CreateTableSchemaUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	req := &CreateTableSchemaRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	// query service
	ws, err := s.tableSchemaService.CreateTableSchema(ctx, &table_schema.CreateTableSchemaParams{
		ProjectId: uriParams.ProjectId,
		Schema:    &req.TableSchema,
	},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, ws)
}

type UpdateTableSchemaByIdUriParams struct {
	ProjectId int    `uri:"projectId" binding:"required"`
	SchemaId  string `uri:"schemaId" binding:"required"`
}

type UpdateTableSchemaByIdRequest struct {
	table_schema.TableSchema
}

func (s *TableSchemaApi) UpdateTableSchemaById(ctx *gin.Context) {
	uriParams := &UpdateTableSchemaByIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	req := &UpdateTableSchemaByIdRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	// query service
	err := s.tableSchemaService.UpdateTableSchemaById(ctx, &table_schema.UpdateTableSchemaByIdParams{
		ProjectId: uriParams.ProjectId,
		SchemaId:  uriParams.SchemaId,
		Schema:    &req.TableSchema,
	},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

type DeleteTableSchemaByIdUriParams struct {
	ProjectId int    `uri:"projectId" binding:"required"`
	SchemaId  string `uri:"schemaId" binding:"required"`
}

func (s *TableSchemaApi) DeleteTableSchemaById(ctx *gin.Context) {
	uriParams := &DeleteTableSchemaByIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	// query service
	err := s.tableSchemaService.DeleteTableSchemaById(ctx, &table_schema.DeleteTableSchemaByIdParams{
		ProjectId: uriParams.ProjectId,
		SchemaId:  uriParams.SchemaId,
	},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
