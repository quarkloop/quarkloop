package table_branch

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/project_table_branch"
)

type ListTableRecordsUriParams struct {
	ProjectId string `uri:"projectId" binding:"required"`
}

type ListTableRecordsResponse struct {
	api.ApiResponse
	Data []model.TableWithRelationCount `json:"data"`
}

func (s *TableBranchApi) ListTableRecords(c *gin.Context) {
	uriParams := &ListTableRecordsUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	// query service
	projectList, err := s.tableBranch.ListTableRecords(
		&project_table_branch.GetTableListParams{
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
	TableBranchId string `uri:"tableId" binding:"required"`
}

type GetTableRecordByIdResponse struct {
	api.ApiResponse
	Data model.TableWithRelationCount `json:"data,omitempty"`
}

func (s *TableBranchApi) GetTableRecordById(c *gin.Context) {
	uriParams := &GetTableRecordByIdUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	// query service
	project_table_branch, err := s.tableBranch.GetTableRecordById(
		&project_table_branch.GetTableByIdParams{
			Context:   c,
			ProjectId: uriParams.ProjectId,
			TableId:   uriParams.TableBranchId,
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
		Data: *project_table_branch,
	}
	c.JSON(http.StatusOK, res)
}
