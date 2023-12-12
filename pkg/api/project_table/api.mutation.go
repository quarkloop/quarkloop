package project_table

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/project_table"
)

type CreateProjectTableUriParams struct {
	ProjectId string `uri:"projectId" binding:"required"`
}

type CreateProjectTableRequest struct {
	model.Table
}

type CreateProjectTableResponse struct {
	api.ApiResponse
	Data model.Table `json:"data,omitempty"`
}

func (s *ProjectTableApi) CreateProjectTable(c *gin.Context) {
	uriParams := &CreateProjectTableUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	req := &CreateProjectTableRequest{}
	if err := c.BindJSON(req); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	// query service
	ws, err := s.projectTable.CreateTable(
		&project_table.CreateTableParams{
			Context:   c,
			ProjectId: uriParams.ProjectId,
			Table:     &req.Table,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	res := &CreateProjectTableResponse{
		ApiResponse: api.ApiResponse{
			Status:       http.StatusCreated,
			StatusString: "Created",
		},
		Data: *ws,
	}
	c.JSON(http.StatusCreated, res)
}

type UpdateProjectTableByIdUriParams struct {
	ProjectId      string `uri:"projectId" binding:"required"`
	ProjectTableId string `uri:"tableId" binding:"required"`
}

type UpdateProjectTableByIdRequest struct {
	model.Table
}

func (s *ProjectTableApi) UpdateProjectTableById(c *gin.Context) {
	uriParams := &UpdateProjectTableByIdUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	req := &UpdateProjectTableByIdRequest{}
	if err := c.BindJSON(req); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	// query service
	err := s.projectTable.UpdateTableById(
		&project_table.UpdateTableByIdParams{
			Context: c,
			TableId: uriParams.ProjectTableId,
			Table:   &req.Table,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}

type DeleteProjectTableByIdUriParams struct {
	ProjectTableId string `uri:"projectTableId" binding:"required"`
}

func (s *ProjectTableApi) DeleteProjectTableById(c *gin.Context) {
	uriParams := &DeleteProjectTableByIdUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	// query service
	err := s.projectTable.DeleteTableById(
		&project_table.DeleteTableByIdParams{
			Context: c,
			TableId: uriParams.ProjectTableId,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
