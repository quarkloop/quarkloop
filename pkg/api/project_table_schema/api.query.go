package table_schema

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/model"
	table_schema "github.com/quarkloop/quarkloop/pkg/service/project_table_schema"
)

type ListTableSchemasUriParams struct {
	ProjectId string `uri:"projectId" binding:"required"`
}

type ListTableSchemasResponse struct {
	api.ApiResponse
	Data []model.TableSchema `json:"data"`
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

	res := &ListTableSchemasResponse{
		ApiResponse: api.ApiResponse{
			Status:       http.StatusOK,
			StatusString: "OK",
		},
		Data: projectList,
	}
	c.JSON(http.StatusOK, res)
}

type GetTableSchemaByIdUriParams struct {
	ProjectId string `uri:"projectId" binding:"required"`
	SchemaId  string `uri:"schemaId" binding:"required"`
}

type GetTableSchemaByIdResponse struct {
	api.ApiResponse
	Data model.TableSchema `json:"data,omitempty"`
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

	res := &GetTableSchemaByIdResponse{
		ApiResponse: api.ApiResponse{
			Status:       http.StatusOK,
			StatusString: "OK",
		},
		Data: *project_table,
	}
	c.JSON(http.StatusOK, res)
}
