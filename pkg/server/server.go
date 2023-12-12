package server

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/api/org"
	"github.com/quarkloop/quarkloop/pkg/api/project"
	"github.com/quarkloop/quarkloop/pkg/api/project_submission"
	"github.com/quarkloop/quarkloop/pkg/api/project_table"
	"github.com/quarkloop/quarkloop/pkg/api/workspace"
	organization_impl "github.com/quarkloop/quarkloop/pkg/service/organization/impl"
	project_impl "github.com/quarkloop/quarkloop/pkg/service/project/impl"
	project_submission_impl "github.com/quarkloop/quarkloop/pkg/service/project_submission/impl"
	project_table_impl "github.com/quarkloop/quarkloop/pkg/service/project_table/impl"
	workspace_impl "github.com/quarkloop/quarkloop/pkg/service/workspace/impl"
	"github.com/quarkloop/quarkloop/pkg/store/repository"
)

type Server struct {
	router               *gin.Engine
	dataStore            *repository.Repository
	orgApi               org.Api
	workspaceApi         workspace.Api
	projectApi           project.Api
	projectTableApi      project_table.Api
	projectSubmissionApi project_submission.Api
}

func NewDefaultServer(ds *repository.Repository) Server {
	router := gin.Default()
	router.RedirectFixedPath = false
	router.RedirectTrailingSlash = true
	router.RemoveExtraSlash = true

	router.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
			"accept",
			"origin",
			"Cache-Control",
			"X-Requested-With",
		},
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:        12 * time.Hour,
	}))

	orgService := organization_impl.NewOrganizationService(ds)
	workspaceService := workspace_impl.NewWorkspaceService(ds)
	projectTableService := project_table_impl.NewTableService(ds)
	projectTable := project_impl.NewProjectService(ds, projectTableService)
	projectSubmission := project_submission_impl.NewAppSubmissionService(ds)

	serve := Server{
		router:               router,
		dataStore:            ds,
		orgApi:               org.NewOrganizationApi(orgService),
		workspaceApi:         workspace.NewWorkspaceApi(workspaceService),
		projectApi:           project.NewProjectApi(projectTable),
		projectTableApi:      project_table.NewProjectTableApi(projectTableService),
		projectSubmissionApi: project_submission.NewAppSubmissionApi(projectSubmission),
	}

	return serve
}

func (s *Server) BindHandlers(api *api.ServerApi) {
	router := s.router.Group("/api/v1")

	// Organizations apis
	orgGroup := router.Group("/orgs")
	orgGroup.GET("", s.orgApi.GetOrganizationList)
	orgGroup.POST("", s.orgApi.CreateOrganization)
	// TODO: first must be a reserved name
	orgGroup.GET("/first", s.orgApi.GetOrganization)
	orgGroup.GET("/:orgId", s.orgApi.GetOrganizationById)
	orgGroup.PUT("/:orgId", s.orgApi.UpdateOrganizationById)
	orgGroup.DELETE("/:orgId", s.orgApi.DeleteOrganizationById)

	// Workspaces apis
	wsGroup := router.Group("/workspaces")
	wsGroup.GET("", s.workspaceApi.GetWorkspaceList)
	wsGroup.POST("", s.workspaceApi.CreateWorkspace)
	// TODO: first must be a reserved name
	wsGroup.GET("/first", s.workspaceApi.GetWorkspace)
	wsGroup.GET("/:workspaceId", s.workspaceApi.GetWorkspaceById)
	wsGroup.PUT("/:workspaceId", s.workspaceApi.UpdateWorkspaceById)
	wsGroup.DELETE("/:workspaceId", s.workspaceApi.DeleteWorkspaceById)

	// Projects apis
	projectGroup := router.Group("/projects")
	projectGroup.GET("", s.projectApi.GetProjectList)
	projectGroup.POST("", s.projectApi.CreateProject)
	projectGroup.GET("/:projectId", s.projectApi.GetProjectById)
	projectGroup.PUT("/:projectId", s.projectApi.UpdateProjectById)
	projectGroup.DELETE("/:projectId", s.projectApi.DeleteProjectById)

	// Tables apis
	projectGroup.GET("/:projectId/tables", s.projectTableApi.ListTableRecords)
	projectGroup.POST("/:projectId/tables", s.projectTableApi.CreateProjectTable)
	projectGroup.DELETE("/:projectId/tables/:tableType", s.projectTableApi.DeleteProjectTableById)

	// Branches apis
	projectGroup.GET("/:projectId/tables/main/branches", s.projectTableApi.ListTableRecords)
	projectGroup.POST("/:projectId/tables/main/branches", s.projectTableApi.CreateProjectTable)
	projectGroup.GET("/:projectId/tables/main/branches/:branchId", s.projectTableApi.GetTableRecordById)
	projectGroup.PUT("/:projectId/tables/main/branches/:branchId", s.projectTableApi.UpdateProjectTableById)
	projectGroup.DELETE("/:projectId/tables/main/branches/:branchId", s.projectTableApi.DeleteProjectTableById)

	// Records apis
	projectGroup.GET("/:projectId/tables/:tableType/branches/:branchId/records", s.projectTableApi.ListTableRecords)
	projectGroup.GET("/:projectId/tables/:tableType/branches/:branchId/records/count", s.projectTableApi.ListTableRecords)
	projectGroup.POST("/:projectId/tables/:tableType/branches/:branchId/records", s.projectTableApi.CreateProjectTable)
	projectGroup.GET("/:projectId/tables/:tableType/branches/:branchId/records/:recordId", s.projectTableApi.GetTableRecordById)
	projectGroup.PUT("/:projectId/tables/:tableType/branches/:branchId/records/:recordId", s.projectTableApi.UpdateProjectTableById)
	projectGroup.DELETE("/:projectId/tables/:tableType/branches/:branchId/records/:recordId", s.projectTableApi.DeleteProjectTableById)

	// Schemas apis
	projectGroup.GET("/:projectId/tables/form/schemas", s.projectTableApi.ListTableRecords)
	projectGroup.POST("/:projectId/tables/form/schemas", s.projectTableApi.CreateProjectTable)
	projectGroup.GET("/:projectId/tables/form/schemas/:schemaId", s.projectTableApi.GetTableRecordById)
	projectGroup.PUT("/:projectId/tables/form/schemas/:schemaId", s.projectTableApi.UpdateProjectTableById)
	projectGroup.DELETE("/:projectId/tables/form/schemas/:schemaId", s.projectTableApi.DeleteProjectTableById)

	// // submissions apis
	// projectGroup.GET("/:projectId/submissions", s.projectSubmissionApi.GetAppSubmissionList)
	// projectGroup.POST("/:projectId/submissions", s.projectSubmissionApi.CreateAppSubmission)
	// projectGroup.GET("/:projectId/submissions/:submissionId", s.projectSubmissionApi.GetAppSubmissionById)
	// projectGroup.PUT("/:projectId/submissions/:submissionId", s.projectSubmissionApi.UpdateAppSubmissionById)
	// projectGroup.DELETE("/:projectId/submissions/:submissionId", s.projectSubmissionApi.DeleteAppSubmissionById)
}

func (server *Server) Router() *gin.Engine {
	return server.router
}

func (server *Server) DataStore() *repository.Repository {
	return server.dataStore
}
