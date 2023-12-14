package table_branch

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/model"
	table_branch "github.com/quarkloop/quarkloop/pkg/service/project_table_branch"
)

type ListTableBranchesUriParams struct {
	ProjectId string `uri:"projectId" binding:"required"`
}

type ListTableBranchesResponse struct {
	api.ApiResponse
	Data []model.TableBranch `json:"data"`
}

func (s *TableBranchApi) ListTableBranches(c *gin.Context) {
	uriParams := &ListTableBranchesUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	// query service
	projectList, err := s.tableBranchService.ListTableBranches(
		&table_branch.GetTableBranchListParams{
			Context:   c,
			ProjectId: uriParams.ProjectId,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	res := &ListTableBranchesResponse{
		ApiResponse: api.ApiResponse{
			Status:       http.StatusOK,
			StatusString: "OK",
		},
		Data: projectList,
	}
	c.JSON(http.StatusOK, res)
}

type GetTableBranchByIdUriParams struct {
	ProjectId string `uri:"projectId" binding:"required"`
	BranchId  string `uri:"branchId" binding:"required"`
}

type GetTableBranchByIdResponse struct {
	api.ApiResponse
	Data model.TableBranch `json:"data,omitempty"`
}

func (s *TableBranchApi) GetTableBranchById(c *gin.Context) {
	uriParams := &GetTableBranchByIdUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	// query service
	project_table_branch, err := s.tableBranchService.GetTableBranchById(
		&table_branch.GetTableBranchByIdParams{
			Context:   c,
			ProjectId: uriParams.ProjectId,
			BranchId:  uriParams.BranchId,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	res := &GetTableBranchByIdResponse{
		ApiResponse: api.ApiResponse{
			Status:       http.StatusOK,
			StatusString: "OK",
		},
		Data: *project_table_branch,
	}
	c.JSON(http.StatusOK, res)
}
