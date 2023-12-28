package workspace

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/service/workspace"
)

type GetWorkspaceListQueryParams struct {
	OrgId []int `form:"orgId"`
}

func (s *WorkspaceApi) GetWorkspaceList(ctx *gin.Context) {
	queryParams := &GetWorkspaceListQueryParams{}
	if err := ctx.ShouldBindQuery(queryParams); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	// query service
	wsList, err := s.workspaceService.GetWorkspaceList(ctx, &workspace.GetWorkspaceListParams{
		OrgId: queryParams.OrgId,
	},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, &wsList)
}

type GetWorkspaceByIdUriParams struct {
	WorkspaceId int `uri:"workspaceId" binding:"required"`
}

func (s *WorkspaceApi) GetWorkspaceById(ctx *gin.Context) {
	uriParams := &GetWorkspaceByIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	// query service
	ws, err := s.workspaceService.GetWorkspaceById(ctx, &workspace.GetWorkspaceByIdParams{
		WorkspaceId: uriParams.WorkspaceId,
	},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, ws)
}

// type GetWorkspaceQueryParams struct {
// 	OrgId int `form:"orgId" binding:"required"`
// 	workspace.Workspace
// }

// func (s *WorkspaceApi) GetWorkspace(ctx *gin.Context) {
// 	queryParams := &GetWorkspaceQueryParams{}
// 	if err := ctx.ShouldBindQuery(queryParams); err != nil {
// 		api.AbortWithBadRequestJSON(ctx, err)
// 		return
// 	}

// 	// query service
// 	ws, err := s.workspaceService.GetWorkspace(ctx, &workspace.GetWorkspaceParams{
// 		OrgId:     queryParams.OrgId,
// 		Workspace: queryParams.Workspace,
// 	},
// 	)
// 	if err != nil {
// 		api.AbortWithInternalServerErrorJSON(ctx, err)
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, ws)
// }
