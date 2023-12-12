package table_record

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/project_table"
)

type ListTableRecordsUriParams struct {
	ProjectId string `uri:"projectId" binding:"required"`
}

type ListTableRecordsResponse struct {
	api.ApiResponse
	Data []model.TableWithRelationCount `json:"data"`
}

func (s *TableRecordApi) ListTableRecords(c *gin.Context) {
	uriParams := &ListTableRecordsUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	// query service
	projectList, err := s.tableRecord.ListTableRecords(
		&project_table.GetTableListParams{
			Context:   c,
			ProjectId: uriParams.ProjectId,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	res := &ListTableRecordsResponse{
		ApiResponse: api.ApiResponse{
			Status:       http.StatusOK,
			StatusString: "OK",
		},
		Data: projectList,
	}
	c.JSON(http.StatusOK, res)
}

type GetTableRecordByIdUriParams struct {
	ProjectId     string `uri:"projectId" binding:"required"`
	TableRecordId string `uri:"tableId" binding:"required"`
}

type GetTableRecordByIdResponse struct {
	api.ApiResponse
	Data model.TableWithRelationCount `json:"data,omitempty"`
}

func (s *TableRecordApi) GetTableRecordById(c *gin.Context) {
	uriParams := &GetTableRecordByIdUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	// query service
	project_table, err := s.tableRecord.GetTableRecordById(
		&project_table.GetTableByIdParams{
			Context:   c,
			ProjectId: uriParams.ProjectId,
			TableId:   uriParams.TableRecordId,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	res := &GetTableRecordByIdResponse{
		ApiResponse: api.ApiResponse{
			Status:       http.StatusOK,
			StatusString: "OK",
		},
		Data: *project_table,
	}
	c.JSON(http.StatusOK, res)
}
