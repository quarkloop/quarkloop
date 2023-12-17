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

func (s *TableBranchApi) ListTableBranches(ctx *gin.Context) {
	uriParams := &ListTableBranchesUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	// query service
	branchList, err := s.tableBranchService.ListTableBranches(ctx, &table_branch.GetTableBranchListParams{
		ProjectId: uriParams.ProjectId,
	},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, &branchList)
}

type GetTableBranchByIdUriParams struct {
	ProjectId int `uri:"projectId" binding:"required"`
	BranchId  int `uri:"branchId" binding:"required"`
}

func (s *TableBranchApi) GetTableBranchById(ctx *gin.Context) {
	uriParams := &GetTableBranchByIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	// query service
	branch, err := s.tableBranchService.GetTableBranchById(ctx, &table_branch.GetTableBranchByIdParams{
		ProjectId: uriParams.ProjectId,
		BranchId:  uriParams.BranchId,
	},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, branch)
}
