package api

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/quarkloop/quarkloop/pkg/model"
// 	"github.com/quarkloop/quarkloop/pkg/store/repository"
// )

// type CreateAppComponentUriParams struct {
// 	AppId string `uri:"projectId" binding:"required"`
// }

// type CreateAppComponentRequest struct {
// 	model.AppComponent
// }

// type CreateAppComponentResponse struct {
// 	ApiResponse
// 	Data model.SheetInstance `json:"data,omitempty"`
// }

// func (s *ServerApi) CreateAppComponent(c *gin.Context) {
// 	uriParams := &CreateAppComponentUriParams{}
// 	if err := c.ShouldBindUri(uriParams); err != nil {
// 		AbortWithBadRequestJSON(c, err)
// 		return
// 	}

// 	req := &CreateAppComponentRequest{}
// 	if err := c.BindJSON(req); err != nil {
// 		AbortWithBadRequestJSON(c, err)
// 		return
// 	}

// 	// query database
// 	ws, err := s.dataStore.CreateAppComponent(
// 		&repository.CreateAppComponentParams{
// 			Context:       c,
// 			AppId:         uriParams.AppId,
// 			SheetInstance: req.SheetInstance,
// 		},
// 	)
// 	if err != nil {
// 		AbortWithInternalServerErrorJSON(c, err)
// 		return
// 	}

// 	res := &CreateAppComponentResponse{
// 		ApiResponse: ApiResponse{
// 			Status:       http.StatusCreated,
// 			StatusString: "Created",
// 		},
// 		Data: *ws,
// 	}
// 	c.JSON(http.StatusCreated, res)
// }

// type UpdateAppComponentByIdRequest struct{}
// type UpdateAppComponentByIdResponse struct{}

// func (s *ServerApi) UpdateAppComponentById(c *gin.Context) {
// 	req := &UpdateAppComponentByIdRequest{}
// 	if err := c.BindJSON(req); err != nil {
// 		AbortWithBadRequestJSON(c, err)
// 		return
// 	}

// 	orgId := c.Param("orgId")
// 	workspaceId := c.Param("workspaceId")
// 	projectId := c.Param("projectId")
// 	componentId := c.Param("componentId")
// 	_ = orgId + workspaceId + projectId + componentId

// 	// query database

// 	res := &UpdateAppComponentByIdResponse{}
// 	c.JSON(http.StatusOK, res)
// }

// func (s *ServerApi) DeleteAppComponentById(c *gin.Context) {
// 	orgId := c.Param("orgId")
// 	workspaceId := c.Param("workspaceId")
// 	projectId := c.Param("projectId")
// 	componentId := c.Param("componentId")
// 	_ = orgId + workspaceId + projectId + componentId

// 	// query database

// 	c.JSON(http.StatusNoContent, nil)
// }
