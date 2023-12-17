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

func (s *TableSchemaApi) ListTableSchemas(c *gin.Context) {
	uriParams := &ListTableSchemasUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	// query service
	projectList, err := s.tableSchemaService.ListTableSchemas(
		&table_schema.GetTableSchemaListParams{
			Context:   c,
			ProjectId: uriParams.ProjectId,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	c.JSON(http.StatusOK, &projectList)
}

type GetTableSchemaByIdUriParams struct {
	ProjectId int    `uri:"projectId" binding:"required"`
	SchemaId  string `uri:"schemaId" binding:"required"`
}

func (s *TableSchemaApi) GetTableSchemaById(c *gin.Context) {
	uriParams := &GetTableSchemaByIdUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	// query service
	project_table, err := s.tableSchemaService.GetTableSchemaById(
		&table_schema.GetTableSchemaByIdParams{
			Context:   c,
			ProjectId: uriParams.ProjectId,
			SchemaId:  uriParams.SchemaId,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	c.JSON(http.StatusOK, project_table)
}
