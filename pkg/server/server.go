package server

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/api/org"
	"github.com/quarkloop/quarkloop/pkg/api/project"
	"github.com/quarkloop/quarkloop/pkg/api/project_submission"
	table_branch "github.com/quarkloop/quarkloop/pkg/api/project_table_branch"
	table_record "github.com/quarkloop/quarkloop/pkg/api/project_table_record"
	table_schema "github.com/quarkloop/quarkloop/pkg/api/project_table_schema"
	"github.com/quarkloop/quarkloop/pkg/api/workspace"
	org_impl "github.com/quarkloop/quarkloop/pkg/service/organization/impl"
	org_store "github.com/quarkloop/quarkloop/pkg/service/organization/store"
	project_impl "github.com/quarkloop/quarkloop/pkg/service/project/impl"
	project_submission_impl "github.com/quarkloop/quarkloop/pkg/service/project_submission/impl"
	table_branch_impl "github.com/quarkloop/quarkloop/pkg/service/project_table_branch/impl"
	table_record_impl "github.com/quarkloop/quarkloop/pkg/service/project_table_record/impl"
	table_schema_impl "github.com/quarkloop/quarkloop/pkg/service/project_table_schema/impl"
	ws_impl "github.com/quarkloop/quarkloop/pkg/service/workspace/impl"
	ws_store "github.com/quarkloop/quarkloop/pkg/service/workspace/store"
	"github.com/quarkloop/quarkloop/pkg/store/repository"
)

type Server struct {
	router    *gin.Engine
	dataStore *repository.Repository

	orgApi       org.Api
	workspaceApi workspace.Api
	projectApi   project.Api

	tableBranchApi table_branch.Api
	tableSchemaApi table_schema.Api
	tableRecordApi table_record.Api

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

	orgService := org_impl.NewOrganizationService(org_store.NewOrgStore(ds.SystemDbConn))
	workspaceService := ws_impl.NewWorkspaceService(ws_store.NewWorkspaceStore(ds.SystemDbConn))
	projectTableService := project_impl.NewProjectService(ds)

	tableBranchService := table_branch_impl.NewTableBranchService(ds)
	tableSchemaService := table_schema_impl.NewTableSchemaService(ds)
	tableRecordService := table_record_impl.NewTableRecordService(ds)

	projectSubmissionService := project_submission_impl.NewAppSubmissionService(ds)

	serve := Server{
		router:    router,
		dataStore: ds,

		orgApi:       org.NewOrganizationApi(orgService),
		workspaceApi: workspace.NewWorkspaceApi(workspaceService),
		projectApi:   project.NewProjectApi(projectTableService),

		tableBranchApi: table_branch.NewTableBranchApi(tableBranchService),
		tableSchemaApi: table_schema.NewTableSchemaApi(tableSchemaService),
		tableRecordApi: table_record.NewTableRecordApi(tableRecordService),

		projectSubmissionApi: project_submission.NewAppSubmissionApi(projectSubmissionService),
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

	// // Tables apis
	// projectGroup.GET("/:projectId/tables", s.projectTableApi.ListTableRecords)
	// projectGroup.POST("/:projectId/tables", s.projectTableApi.CreateProjectTable)
	// projectGroup.DELETE("/:projectId/tables/:tableType", s.projectTableApi.DeleteProjectTableById)

	// Branches apis
	projectGroup.GET("/:projectId/tables/main/branches", s.tableBranchApi.ListTableBranches)
	projectGroup.GET("/:projectId/tables/main/branches/:branchId", s.tableBranchApi.GetTableBranchById)

	// Schemas apis
	projectGroup.GET("/:projectId/tables/form/schemas", s.tableSchemaApi.ListTableSchemas)
	projectGroup.POST("/:projectId/tables/form/schemas", s.tableSchemaApi.CreateTableSchema)
	projectGroup.GET("/:projectId/tables/form/schemas/:schemaId", s.tableSchemaApi.GetTableSchemaById)
	projectGroup.PUT("/:projectId/tables/form/schemas/:schemaId", s.tableSchemaApi.UpdateTableSchemaById)
	projectGroup.DELETE("/:projectId/tables/form/schemas/:schemaId", s.tableSchemaApi.DeleteTableSchemaById)

	// Records apis
	projectGroup.GET("/:projectId/tables/:tableType/branches/:branchId/records", s.tableRecordApi.ListTableRecords)
	//projectGroup.GET("/:projectId/tables/:tableType/branches/:branchId/records/count", s.tableRecordApi.ListTableRecords)
	projectGroup.POST("/:projectId/tables/:tableType/branches/:branchId/records", s.tableRecordApi.CreateTableRecord)
	projectGroup.GET("/:projectId/tables/:tableType/branches/:branchId/records/:recordId", s.tableRecordApi.GetTableRecordById)
	projectGroup.PUT("/:projectId/tables/:tableType/branches/:branchId/records/:recordId", s.tableRecordApi.UpdateTableRecordById)
	projectGroup.DELETE("/:projectId/tables/:tableType/branches/:branchId/records/:recordId", s.tableRecordApi.DeleteTableRecordById)

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
