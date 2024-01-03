package server

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/api/org"
	"github.com/quarkloop/quarkloop/pkg/api/project"
	"github.com/quarkloop/quarkloop/pkg/api/project_submission"
	table_branch "github.com/quarkloop/quarkloop/pkg/api/table_branch"
	table_record "github.com/quarkloop/quarkloop/pkg/api/table_record"
	table_schema "github.com/quarkloop/quarkloop/pkg/api/table_schema"
	"github.com/quarkloop/quarkloop/pkg/api/workspace"
	"github.com/quarkloop/quarkloop/pkg/contextdata"
	acl_impl "github.com/quarkloop/quarkloop/pkg/service/accesscontrol/impl"
	acl_store "github.com/quarkloop/quarkloop/pkg/service/accesscontrol/store"
	org_impl "github.com/quarkloop/quarkloop/pkg/service/org/impl"
	org_store "github.com/quarkloop/quarkloop/pkg/service/org/store"
	project_impl "github.com/quarkloop/quarkloop/pkg/service/project/impl"
	project_store "github.com/quarkloop/quarkloop/pkg/service/project/store"
	project_submission_impl "github.com/quarkloop/quarkloop/pkg/service/project_submission/impl"
	quota_impl "github.com/quarkloop/quarkloop/pkg/service/quota/impl"
	quota_store "github.com/quarkloop/quarkloop/pkg/service/quota/store"
	table_branch_impl "github.com/quarkloop/quarkloop/pkg/service/table_branch/impl"
	table_branch_store "github.com/quarkloop/quarkloop/pkg/service/table_branch/store"
	table_record_impl "github.com/quarkloop/quarkloop/pkg/service/table_record/impl"
	table_record_store "github.com/quarkloop/quarkloop/pkg/service/table_record/store"
	table_schema_impl "github.com/quarkloop/quarkloop/pkg/service/table_schema/impl"
	table_schema_store "github.com/quarkloop/quarkloop/pkg/service/table_schema/store"
	"github.com/quarkloop/quarkloop/pkg/service/user"
	user_impl "github.com/quarkloop/quarkloop/pkg/service/user/impl"
	user_store "github.com/quarkloop/quarkloop/pkg/service/user/store"
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

	userService := user_impl.NewUserService(user_store.NewOrgStore(ds.SystemDbConn))
	aclService := acl_impl.NewAccessControlService(acl_store.NewAccessControlStore(ds.SystemDbConn))
	quotaService := quota_impl.NewQuotaService(quota_store.NewQuotaStore(ds.SystemDbConn))

	tableBranchService := table_branch_impl.NewTableBranchService(table_branch_store.NewTableBranchStore(ds.ProjectDbConn))
	tableSchemaService := table_schema_impl.NewTableSchemaService(table_schema_store.NewTableSchemaStore(ds.ProjectDbConn))
	tableRecordService := table_record_impl.NewTableRecordService(table_record_store.NewTableRecordStore(ds.ProjectDbConn))

	orgService := org_impl.NewOrgService(org_store.NewOrgStore(ds.SystemDbConn))
	workspaceService := ws_impl.NewWorkspaceService(ws_store.NewWorkspaceStore(ds.SystemDbConn))
	projectTableService := project_impl.NewProjectService(project_store.NewProjectStore(ds.SystemDbConn))

	projectSubmissionService := project_submission_impl.NewAppSubmissionService(ds)

	serve := Server{
		router:    router,
		dataStore: ds,

		orgApi:       org.NewOrgApi(orgService, userService, aclService, quotaService),
		workspaceApi: workspace.NewWorkspaceApi(workspaceService, userService, aclService, quotaService),
		projectApi:   project.NewProjectApi(projectTableService, userService, aclService, quotaService, tableBranchService),

		tableBranchApi: table_branch.NewTableBranchApi(tableBranchService),
		tableSchemaApi: table_schema.NewTableSchemaApi(tableSchemaService),
		tableRecordApi: table_record.NewTableRecordApi(tableRecordService),

		projectSubmissionApi: project_submission.NewAppSubmissionApi(projectSubmissionService),
	}

	return serve
}

func (s *Server) UserAuth(ctx *gin.Context) {
	req, _ := http.NewRequest("GET", "http://localhost:3001/api/auth/check", nil)
	req.Header = ctx.Request.Header.Clone()

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		ctx.AbortWithError(resp.StatusCode, errors.New(resp.Status))
		return
	}

	u := &user.User{}
	err = json.NewDecoder(resp.Body).Decode(u)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// set user context data
	contextdata.SetUser(ctx, u)
	ctx.Next()
}

func (s *Server) AbortAnonymousUserRequest(ctx *gin.Context) {
	if contextdata.IsUserAnonymous(ctx) {
		ctx.AbortWithError(http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	ctx.Next()
}

func (s *Server) TestApi(ctx *gin.Context) {
	u := contextdata.GetUser(ctx)
	ctx.JSON(http.StatusOK, u)
}

func (s *Server) BindHandlers(api *api.ServerApi) {
	router := s.router.Group("/api/v1")

	testGroup := router.Group("/test")
	testGroup.Use(s.UserAuth)
	testGroup.GET("", s.TestApi)

	// org apis
	// org query
	orgGroup := router.Group("/orgs")
	orgGroup.Use(s.UserAuth)
	{
		orgGroup.GET("", s.orgApi.GetOrgList)
		orgGroup.GET("/:orgId", s.orgApi.GetOrgById)
		// TODO: first must be a reserved name
		// TODO: rewrite
		// orgGroup.GET("/first", s.orgApi.GetOrg)
	}
	// org mutation
	orgMutationGroup := orgGroup.Group("")
	orgMutationGroup.Use(s.AbortAnonymousUserRequest)
	{
		orgMutationGroup.POST("", s.orgApi.CreateOrg)
		orgMutationGroup.PUT("/:orgId", s.orgApi.UpdateOrgById)
		orgMutationGroup.DELETE("/:orgId", s.orgApi.DeleteOrgById)
	}

	// workspace apis
	// workspace query
	wsGroup := router.Group("/workspaces")
	wsGroup.Use(s.UserAuth)
	{
		wsGroup.GET("", s.workspaceApi.GetWorkspaceList)
		wsGroup.GET("/:workspaceId", s.workspaceApi.GetWorkspaceById)
		// TODO: first must be a reserved name
		// TODO: rewrite
		// wsGroup.GET("/first", s.workspaceApi.GetWorkspace)
	}
	// workspace mutation
	workspaceMutationGroup := wsGroup.Group("")
	workspaceMutationGroup.Use(s.AbortAnonymousUserRequest)
	{
		workspaceMutationGroup.POST("", s.workspaceApi.CreateWorkspace)
		workspaceMutationGroup.PUT("/:workspaceId", s.workspaceApi.UpdateWorkspaceById)
		workspaceMutationGroup.DELETE("/:workspaceId", s.workspaceApi.DeleteWorkspaceById)
	}

	// project apis
	// project query
	projectGroup := router.Group("/projects")
	projectGroup.Use(s.UserAuth)
	{
		projectGroup.GET("", s.projectApi.GetProjectList)
		projectGroup.GET("/:projectId", s.projectApi.GetProjectById)
	}
	// project mutation
	projectMutationGroup := projectGroup.Group("")
	projectMutationGroup.Use(s.AbortAnonymousUserRequest)
	{
		projectMutationGroup.POST("", s.projectApi.CreateProject)
		projectMutationGroup.PUT("/:projectId", s.projectApi.UpdateProjectById)
		projectMutationGroup.DELETE("/:projectId", s.projectApi.DeleteProjectById)
	}

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
