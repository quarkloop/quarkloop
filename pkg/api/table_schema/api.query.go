package table_schema

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/service/table_schema"
)

type ListTableSchemasUriParams struct {
	ProjectId int `uri:"projectId" binding:"required"`
}

func (s *TableSchemaApi) ListTableSchemas(ctx *gin.Context) {
	uriParams := &ListTableSchemasUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	// query service
	projectList, err := s.tableSchemaService.ListTableSchemas(ctx, &table_schema.GetTableSchemaListParams{
		ProjectId: uriParams.ProjectId,
	},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, &projectList)
}

type GetTableSchemaByIdUriParams struct {
	ProjectId int    `uri:"projectId" binding:"required"`
	SchemaId  string `uri:"schemaId" binding:"required"`
}

func (s *TableSchemaApi) GetTableSchemaById(ctx *gin.Context) {
	uriParams := &GetTableSchemaByIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	// query service
	project_table, err := s.tableSchemaService.GetTableSchemaById(ctx, &table_schema.GetTableSchemaByIdParams{
		ProjectId: uriParams.ProjectId,
		SchemaId:  uriParams.SchemaId,
	},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, project_table)
}
