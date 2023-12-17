package table_branch

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/service/table_branch"
)

type ListTableBranchesUriParams struct {
	ProjectId int `uri:"projectId" binding:"required"`
}

func (s *TableBranchApi) ListTableBranches(c *gin.Context) {
	uriParams := &ListTableBranchesUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	// query service
	branchList, err := s.tableBranchService.ListTableBranches(
		&table_branch.GetTableBranchListParams{
			Context:   c,
			ProjectId: uriParams.ProjectId,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	c.JSON(http.StatusOK, &branchList)
}

type GetTableBranchByIdUriParams struct {
	ProjectId int `uri:"projectId" binding:"required"`
	BranchId  int `uri:"branchId" binding:"required"`
}

func (s *TableBranchApi) GetTableBranchById(c *gin.Context) {
	uriParams := &GetTableBranchByIdUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	// query service
	branch, err := s.tableBranchService.GetTableBranchById(
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

	c.JSON(http.StatusOK, branch)
}
