package server

import (
	"encoding/json"
	"errors"
	"flag"
	"log"
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
	"github.com/quarkloop/quarkloop/pkg/api/user"
	"github.com/quarkloop/quarkloop/pkg/api/workspace"
	"github.com/quarkloop/quarkloop/pkg/contextdata"
	acl_impl "github.com/quarkloop/quarkloop/pkg/service/accesscontrol/impl"
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
	user_service "github.com/quarkloop/quarkloop/pkg/service/user"
	user_impl "github.com/quarkloop/quarkloop/pkg/service/user/impl"
	user_store "github.com/quarkloop/quarkloop/pkg/service/user/store"
	ws_impl "github.com/quarkloop/quarkloop/pkg/service/workspace/impl"
	ws_store "github.com/quarkloop/quarkloop/pkg/service/workspace/store"
	"github.com/quarkloop/quarkloop/pkg/store/repository"
	"github.com/quarkloop/quarkloop/service/system"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	authzServiceAddr = flag.String("authzServiceAddr", "localhost:50051", "the address to connect to")
	orgServiceAddr   = flag.String("orgServiceAddr", "localhost:50095", "the address to connect to")
)

type Server struct {
	router    *gin.Engine
	dataStore *repository.Repository

	orgService system.OrgServiceClient

	userApi              user.Api
	orgApi               org.Api
	workspaceApi         workspace.Api
	projectApi           project.Api
	tableBranchApi       table_branch.Api
	tableSchemaApi       table_schema.Api
	tableRecordApi       table_record.Api
	projectSubmissionApi project_submission.Api
}

func NewDefaultServer(ds *repository.Repository) Server {
	//gin.SetMode("release")
	router := gin.Default()
	router.RedirectFixedPath = false
	router.RedirectTrailingSlash = true
	router.RemoveExtraSlash = true
	router.Use(gin.Recovery())

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

	_, err := grpc.Dial(*authzServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("[grpc][authz] could not connect: %v", err)
	}

	orgServiceConn, err := grpc.Dial(*orgServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("[grpc][org] could not connect: %v", err)
	}

	userService := user_impl.NewUserService(user_store.NewUserStore(ds.AuthDbConn))
	aclService := acl_impl.NewAccessControlService()
	quotaService := quota_impl.NewQuotaService(quota_store.NewQuotaStore(ds.SystemDbConn))

	tableBranchService := table_branch_impl.NewTableBranchService(table_branch_store.NewTableBranchStore(ds.ProjectDbConn))
	tableSchemaService := table_schema_impl.NewTableSchemaService(table_schema_store.NewTableSchemaStore(ds.ProjectDbConn))
	tableRecordService := table_record_impl.NewTableRecordService(table_record_store.NewTableRecordStore(ds.ProjectDbConn))

	orgService := system.NewOrgServiceClient(orgServiceConn)
	workspaceService := ws_impl.NewWorkspaceService(ws_store.NewWorkspaceStore(ds.SystemDbConn))
	projectTableService := project_impl.NewProjectService(project_store.NewProjectStore(ds.SystemDbConn))

	projectSubmissionService := project_submission_impl.NewAppSubmissionService(ds)

	serve := Server{
		router:    router,
		dataStore: ds,

		orgService: orgService,

		userApi:              user.NewUserApi(userService, aclService),
		orgApi:               org.NewOrgApi(orgService, userService, aclService, quotaService),
		workspaceApi:         workspace.NewWorkspaceApi(workspaceService, userService, aclService, quotaService),
		projectApi:           project.NewProjectApi(projectTableService, userService, aclService, quotaService, tableBranchService),
		tableBranchApi:       table_branch.NewTableBranchApi(tableBranchService),
		tableSchemaApi:       table_schema.NewTableSchemaApi(tableSchemaService),
		tableRecordApi:       table_record.NewTableRecordApi(tableRecordService),
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
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		ctx.AbortWithStatusJSON(resp.StatusCode, errors.New(resp.Status))
		return
	}

	u := &user_service.User{}
	err = json.NewDecoder(resp.Body).Decode(u)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	// set user context data
	contextdata.SetUser(ctx, u)
	ctx.Next()
}

func (s *Server) AbortAnonymousUserRequest(ctx *gin.Context) {
	if contextdata.IsUserAnonymous(ctx) {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	ctx.Next()
}

// TODO: rewrite
// func (s *Server) ValidateOrgIdUriParam(ctx *gin.Context) {
// 	type OrgIdUriParam struct {
// 		OrgId int `uri:"orgId" binding:"required"`
// 	}

// 	uriParam := &OrgIdUriParam{}
// 	if err := ctx.ShouldBindUri(&uriParam); err != nil {
// 		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
// 		return
// 	}

// 	_, err := s.orgService.GetOrgById(ctx, &orgService.GetOrgByIdQuery{OrgId: uriParam.OrgId})
// 	if err != nil {
// 		ctx.AbortWithStatusJSON(http.StatusNotFound, "org not found")
// 		return
// 	}

// 	ctx.Next()
// }

func (s *Server) TODOHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "This is a TODO handler")
}

func (s *Server) BindHandlers(api *api.ServerApi) {
	// query apis
	query := s.router.Group("/api/v1")
	query.Use(s.UserAuth)
	{
		query.GET("/user", s.userApi.GetUser)
		query.GET("/user/username", s.userApi.GetUsername)
		query.GET("/user/status", s.userApi.GetStatus)
		query.GET("/user/email", s.userApi.GetEmail)
		query.GET("/user/preferences", s.userApi.GetPreferences)
		query.GET("/user/sessions", s.userApi.GetSessions)
		query.GET("/user/accounts", s.userApi.GetAccounts)

		query.GET("/users", s.userApi.GetUsers)
		query.GET("/users/:userId_or_username", s.userApi.GetUserById)
		query.GET("/users/:userId_or_username/username", s.userApi.GetUsernameByUserId)
		query.GET("/users/:userId_or_username/status", s.userApi.GetStatusByUserId)
		query.GET("/users/:userId_or_username/email", s.userApi.GetEmailByUserId)
		query.GET("/users/:userId_or_username/preferences", s.userApi.GetPreferencesByUserId)
		query.GET("/users/:userId_or_username/sessions", s.userApi.GetSessionsByUserId)
		query.GET("/users/:userId_or_username/accounts", s.userApi.GetAccountsByUserId)
	}
	{
		query.GET("/orgs", s.orgApi.GetOrgList)
		query.GET("/workspaces", s.workspaceApi.GetWorkspaceList)
		query.GET("/projects", s.projectApi.GetProjectList)

		query.GET("/orgs/:orgId", s.orgApi.GetOrgById)
		query.GET("/orgs/:orgId/members", s.orgApi.GetMemberList)
		query.GET("/orgs/:orgId/workspaces", s.orgApi.GetWorkspaceList)
		query.GET("/orgs/:orgId/workspaces/:workspaceId", s.workspaceApi.GetWorkspaceById)
		query.GET("/orgs/:orgId/workspaces/:workspaceId/projects", s.workspaceApi.GetProjectList)
		query.GET("/orgs/:orgId/workspaces/:workspaceId/members", s.workspaceApi.GetMemberList)
		query.GET("/orgs/:orgId/projects", s.orgApi.GetProjectList)
		query.GET("/orgs/:orgId/projects/:projectId", s.projectApi.GetProjectById)
		query.GET("/orgs/:orgId/projects/:projectId/members", s.projectApi.GetMemberList)
	}

	// mutation apis
	mutation := s.router.Group("/api/v1")
	mutation.Use(s.UserAuth)
	mutation.Use(s.AbortAnonymousUserRequest)
	{
		mutation.PUT("/user", s.userApi.UpdateUser)
		mutation.PUT("/user/username", s.userApi.UpdateUsername)
		mutation.PUT("/user/password", s.userApi.UpdatePassword)
		mutation.PUT("/user/preferences", s.userApi.UpdatePreferences)

		mutation.PUT("/users/:userId", s.userApi.UpdateUserById)
		mutation.PUT("/users/:userId/username", s.userApi.UpdateUsernameByUserId)
		mutation.PUT("/users/:userId/password", s.userApi.UpdatePasswordByUserId)
		mutation.PUT("/users/:userId/preferences", s.userApi.UpdatePreferencesByUserId)
		mutation.PUT("/users/:userId/activate", s.TODOHandler)
		mutation.PUT("/users/:userId/deactivate", s.TODOHandler)
		mutation.PUT("/users/:userId/block", s.TODOHandler)
		mutation.PUT("/users/:userId/unblock", s.TODOHandler)

		mutation.DELETE("/users/:userId", s.userApi.DeleteUserById)
		mutation.DELETE("/users/:userId/sessions/:sessionId", s.userApi.DeleteSessionById)
		mutation.DELETE("/users/:userId/accounts/:accountId", s.userApi.DeleteAccountById)
	}
	{
		mutation.POST("/orgs", s.orgApi.CreateOrg)
		mutation.POST("/orgs/:orgId/workspaces", s.workspaceApi.CreateWorkspace)
		mutation.POST("/orgs/:orgId/workspaces/:workspaceId/projects", s.projectApi.CreateProject)

		mutation.PUT("/orgs/:orgId", s.orgApi.UpdateOrgById)
		mutation.PUT("/orgs/:orgId/workspaces/:workspaceId", s.workspaceApi.UpdateWorkspaceById)
		mutation.PUT("/orgs/:orgId/workspaces/:workspaceId/projects/:projectId", s.projectApi.UpdateProjectById)

		mutation.DELETE("/orgs/:orgId", s.orgApi.DeleteOrgById)
		mutation.DELETE("/orgs/:orgId/workspaces/:workspaceId", s.workspaceApi.DeleteWorkspaceById)
		mutation.DELETE("/orgs/:orgId/workspaces/:workspaceId/projects/:projectId", s.projectApi.DeleteProjectById)
	}
}

func (s *Server) BindHandlers_old(api *api.ServerApi) {
	router := s.router.Group("/api/v1")

	testGroup := router.Group("/test")
	testGroup.Use(s.UserAuth)

	router.GET("/orgs", s.orgApi.GetOrgList)
	router.GET("/workspaces", s.workspaceApi.GetWorkspaceList)
	router.GET("/projects", s.projectApi.GetProjectList)

	// org query apis
	orgGroup := router.Group("/orgs")
	orgGroup.Use(s.UserAuth)
	//orgGroup.Use(s.ValidateOrgIdUriParam)
	{
		orgGroup.GET("/:orgId", s.orgApi.GetOrgById)
		orgGroup.GET("/:orgId/workspaces", s.orgApi.GetWorkspaceList)
		orgGroup.GET("/:orgId/projects", s.orgApi.GetProjectList)
		orgGroup.GET("/:orgId/members", s.orgApi.GetMemberList)
	}
	// org mutation apis
	orgMutationGroup := orgGroup.Group("")
	orgMutationGroup.Use(s.AbortAnonymousUserRequest)
	{
		orgMutationGroup.POST("", s.orgApi.CreateOrg)
		orgMutationGroup.PUT("/:orgId", s.orgApi.UpdateOrgById)
		orgMutationGroup.DELETE("/:orgId", s.orgApi.DeleteOrgById)
	}

	// workspace query apis
	wsGroup := router.Group("/orgs/:orgId/workspaces")
	wsGroup.Use(s.UserAuth)
	{
		wsGroup.GET("/:workspaceId", s.workspaceApi.GetWorkspaceById)
		wsGroup.GET("/:workspaceId/projects", s.workspaceApi.GetProjectList)
		wsGroup.GET("/:workspaceId/members", s.workspaceApi.GetMemberList)
	}
	// workspace mutation apis
	workspaceMutationGroup := wsGroup.Group("")
	workspaceMutationGroup.Use(s.AbortAnonymousUserRequest)
	{
		workspaceMutationGroup.POST("", s.workspaceApi.CreateWorkspace)
		workspaceMutationGroup.PUT("/:workspaceId", s.workspaceApi.UpdateWorkspaceById)
		workspaceMutationGroup.DELETE("/:workspaceId", s.workspaceApi.DeleteWorkspaceById)
	}

	// project query apis
	projectGroup := router.Group("/orgs/:orgId/workspaces/:workspaceId/projects")
	projectGroup.Use(s.UserAuth)
	{
		projectGroup.GET("/:projectId", s.projectApi.GetProjectById)
		projectGroup.GET("/:projectId/projects", s.projectApi.GetProjectList)
		projectGroup.GET("/:projectId/members", s.projectApi.GetMemberList)
	}
	// project mutation apis
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
