package workspace

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"testing"

	v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/authzed/authzed-go/v1"
	"github.com/authzed/grpcutil"
	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/org"
	"github.com/quarkloop/quarkloop/pkg/service/workspace"
	"github.com/quarkloop/quarkloop/pkg/test"
	"github.com/quarkloop/quarkloop/service/v1/system"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var baseUrl = "http://localhost:8000"

type testCase struct {
	description    string
	path           string
	expectedStatus int
	expectedBody   any
	authenticate   bool
	toUrl          func(path string) string
	decodeBody     func(r io.Reader) (any, error)
	encodeBody     func(body any) (*bytes.Buffer, error)
}

var (
	ctx   context.Context
	authz *authzed.Client
	orgId int32 = 0
)

var schema = `
definition user {}

definition org {
    relation owner: user

    relation admin: user
    relation contributor: user
    relation viewer: user

    permission all = admins + contributors + viewer
    permission admins = owner + admin
    permission contributors = contributor
    permission viewers = viewer

    // org
    permission read = all
    permission update = admins + contributors
    permission delete = admins
    permission settings_read = admins
    permission settings_update = admins
    permission quota_read = all
    permission quota_create = admins
    permission quota_update = admins
    permission quota_delete = admins
    permission user_read = all
    permission user_create = admins
    permission user_update = admins
    permission user_delete = admins

    // workspace
    permission create_workspace = admins
}

definition workspace {
    relation parent: org

    relation admin: user
    relation contributor: user
    relation viewer: user

    permission all = admins + contributors + viewers
    permission admins = admin + parent->admins
    permission contributors = contributor + parent->contributors
    permission viewers = viewer + parent->viewers

    // workspace
    permission read = all
    permission update = admins + contributors
    permission delete = admins
    permission settings_read = admins
    permission settings_update = admins
    permission quota_read = all
    permission quota_create = admins
    permission quota_update = admins
    permission quota_delete = admins
    permission user_read = all
    permission user_create = admins
    permission user_update = admins
    permission user_delete = admins

    // project
    permission create_project = admins
}

definition project {
    relation parent: workspace

    relation admin: user
    relation contributor: user
    relation viewer: user

    permission all = admins + contributors + viewer
    permission admins = admin + parent->admins
    permission contributors = contributor + parent->contributors
    permission viewers = viewer + parent->viewers

    // project
    permission read = all
    permission update = admins + contributors
    permission delete = admins
    permission settings_read = admins
    permission settings_update = admins
    permission quota_read = all
    permission quota_create = admins
    permission quota_update = admins
    permission quota_delete = admins
    permission user_read = all
    permission user_create = admins
    permission user_update = admins
    permission user_delete = admins
}
`

func init() {
	client, err := authzed.NewClient(
		"localhost:50051",
		grpcutil.WithInsecureBearerToken("my_passphrase_key"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("unable to initialize client: %s", err)
	}

	authz = client
	ctx = context.Background()
}

func TestMutationCreateSchema(t *testing.T) {
	t.Run("create schema", func(t *testing.T) {
		resp, err := authz.WriteSchema(ctx, &v1.WriteSchemaRequest{
			Schema: schema,
		})
		token := resp.GetWrittenAt()

		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, token)
	})

	t.Run("read schema", func(t *testing.T) {
		resp, err := authz.ReadSchema(ctx, &v1.ReadSchemaRequest{})
		token := resp.GetReadAt()

		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, token)
	})
}

func TestMutationCreateOrg(t *testing.T) {
	var testCases []struct {
		testCase
		org *org.CreateOrgCommand
	} = []struct {
		testCase
		org *org.CreateOrgCommand
	}{
		{
			testCase: testCase{
				description:    "should create org 1 and return it with success",
				path:           "/api/v1/orgs",
				expectedStatus: http.StatusCreated,
				expectedBody: system.Org{
					ScopeId:     "org_1",
					Name:        "Org_1",
					Description: "Org_1 description",
					Visibility:  int32(model.PublicVisibility),
				},
				authenticate: true,
				toUrl:        func(path string) string { return fmt.Sprint(baseUrl, path) },
				decodeBody: func(r io.Reader) (any, error) {
					res := &system.Org{}
					err := json.NewDecoder(r).Decode(res)
					return res, err
				},
				encodeBody: func(body any) (*bytes.Buffer, error) {
					payloadBuf := &bytes.Buffer{}
					err := json.NewEncoder(payloadBuf).Encode(body)
					return payloadBuf, err
				},
			},
			org: &org.CreateOrgCommand{
				CreatedBy:   "admin",
				ScopeId:     "org_1",
				Name:        "Org_1",
				Description: "Org_1 description",
				Visibility:  model.PublicVisibility,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			reqBody, err := testCase.encodeBody(testCase.org)
			require.NoError(t, err)
			require.NotNil(t, reqBody)

			client := &http.Client{}
			req, err := http.NewRequest("POST", testCase.toUrl(testCase.path), reqBody)
			if err != nil {
				t.Fatal(err)
			}
			if testCase.authenticate {
				req.Header.Set("Cookie", test.GetCookie())
			}

			resp, err := client.Do(req)
			if err != nil {
				log.Fatal(err)
			}
			t.Cleanup(func() { require.NoError(t, resp.Body.Close()) })

			require.NoError(t, err)
			require.Equal(t, testCase.expectedStatus, resp.StatusCode)

			respBody, err := testCase.decodeBody(resp.Body)
			require.NoError(t, err)
			switch body := respBody.(type) {
			case string:
				require.Equal(t, testCase.expectedBody, body)
			case *system.Org:
				switch exp := testCase.expectedBody.(type) {
				case system.Org:
					require.NotNil(t, body)
					require.NotZero(t, body.Id)
					require.NotEmpty(t, body.Name)
					require.NotEmpty(t, body.Description)
					require.NotZero(t, body.Visibility)
					require.Equal(t, exp.ScopeId, body.ScopeId)
					require.Equal(t, exp.Name, body.Name)
					require.Equal(t, exp.Description, body.Description)
					require.Equal(t, exp.Visibility, body.Visibility)
					require.Equal(t, fmt.Sprint("/org/", exp.ScopeId), body.Path)

					orgId = body.Id
				}
			}
		})
	}
}

func TestQueryWorkspaceListApi(t *testing.T) {
	testCases := []testCase{
		{
			description:    "should return empty public and private workspaces for authenticated user",
			path:           "/api/v1/workspaces",
			expectedStatus: http.StatusOK,
			expectedBody:   []*system.Workspace{},
			authenticate:   true,
			toUrl:          func(path string) string { return fmt.Sprint(baseUrl, path) },
			decodeBody: func(r io.Reader) (any, error) {
				var res []*system.Workspace
				err := json.NewDecoder(r).Decode(&res)
				return res, err
			},
		},
		{
			description:    "should return empty public workspaces for unauthenticated user",
			path:           "/api/v1/workspaces",
			expectedStatus: http.StatusOK,
			expectedBody:   []*system.Workspace{},
			authenticate:   false,
			toUrl:          func(path string) string { return fmt.Sprint(baseUrl, path) },
			decodeBody: func(r io.Reader) (any, error) {
				var res []*system.Workspace
				err := json.NewDecoder(r).Decode(&res)
				return res, err
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			client := &http.Client{}
			req, err := http.NewRequest("GET", testCase.toUrl(testCase.path), nil)
			if err != nil {
				t.Fatal(err)
			}
			if testCase.authenticate {
				req.Header.Set("Cookie", test.GetCookie())
			}

			resp, err := client.Do(req)
			if err != nil {
				t.Fatal(err)
			}
			t.Cleanup(func() { require.NoError(t, resp.Body.Close()) })

			require.NoError(t, err)
			require.Equal(t, testCase.expectedStatus, resp.StatusCode)

			body, err := testCase.decodeBody(resp.Body)
			require.NoError(t, err)
			require.Empty(t, body)
			require.Equal(t, testCase.expectedBody, body)
		})
	}

	testCases = []testCase{
		{
			description:    "should return empty public and private workspaces for authenticated user",
			path:           "/api/v1/orgs/%d/workspaces",
			expectedStatus: http.StatusOK,
			expectedBody:   []*system.Workspace{},
			authenticate:   true,
			toUrl:          func(path string) string { return fmt.Sprint(baseUrl, fmt.Sprintf(path, 10)) }, // TODO: %d args
			decodeBody: func(r io.Reader) (any, error) {
				var res []*system.Workspace
				err := json.NewDecoder(r).Decode(&res)
				return res, err
			},
		},
		{
			description:    "should return empty public workspaces for unauthenticated user",
			path:           "/api/v1/orgs/%d/workspaces",
			expectedStatus: http.StatusOK,
			expectedBody:   []*system.Workspace{},
			authenticate:   false,
			toUrl:          func(path string) string { return fmt.Sprint(baseUrl, fmt.Sprintf(path, 10)) }, // TODO: %d args
			decodeBody: func(r io.Reader) (any, error) {
				var res []*system.Workspace
				err := json.NewDecoder(r).Decode(&res)
				return res, err
			},
		},
		{
			description:    "should return bad request with org id 0 for authenticated user",
			path:           "/api/v1/orgs/0/workspaces",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Key: 'GetWorkspaceListQuery.OrgId' Error:Field validation for 'OrgId' failed on the 'required' tag",
			authenticate:   true,
			toUrl:          func(path string) string { return fmt.Sprint(baseUrl, path) },
			decodeBody: func(r io.Reader) (any, error) {
				var res string
				err := json.NewDecoder(r).Decode(&res)
				return res, err
			},
		},
		{
			description:    "should return bad request with org id 0 for unauthenticated user",
			path:           "/api/v1/orgs/0/workspaces",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Key: 'GetWorkspaceListQuery.OrgId' Error:Field validation for 'OrgId' failed on the 'required' tag",
			authenticate:   false,
			toUrl:          func(path string) string { return fmt.Sprint(baseUrl, path) },
			decodeBody: func(r io.Reader) (any, error) {
				var res string
				err := json.NewDecoder(r).Decode(&res)
				return res, err
			},
		},
		{
			description:    "should return bad request with org workspace id as string for authenticated user",
			path:           "/api/v1/orgs/org_id/workspaces/workspace_id",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `field validation for 'org_id' failed`,
			authenticate:   true,
			toUrl:          func(path string) string { return fmt.Sprint(baseUrl, path) },
			decodeBody: func(r io.Reader) (any, error) {
				var res string
				err := json.NewDecoder(r).Decode(&res)
				return res, err
			},
		},
		{
			description:    "should return bad request with org and workspace id as string for unauthenticated user",
			path:           "/api/v1/orgs/org_id/workspaces/workspace_id",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `field validation for 'org_id' failed`,
			authenticate:   false,
			toUrl:          func(path string) string { return fmt.Sprint(baseUrl, path) },
			decodeBody: func(r io.Reader) (any, error) {
				var res string
				err := json.NewDecoder(r).Decode(&res)
				return res, err
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			client := &http.Client{}
			req, err := http.NewRequest("GET", testCase.toUrl(testCase.path), nil)
			if err != nil {
				t.Fatal(err)
			}
			if testCase.authenticate {
				req.Header.Set("Cookie", test.GetCookie())
			}

			resp, err := client.Do(req)
			if err != nil {
				t.Fatal(err)
			}
			t.Cleanup(func() { require.NoError(t, resp.Body.Close()) })

			require.NoError(t, err)
			require.Equal(t, testCase.expectedStatus, resp.StatusCode)

			body, err := testCase.decodeBody(resp.Body)
			require.NoError(t, err)
			require.Equal(t, testCase.expectedBody, body)
		})
	}
}

func TestMutationCreateWorkspaceApi(t *testing.T) {
	var testCases []struct {
		testCase
		workspace *workspace.CreateWorkspaceCommand
	} = []struct {
		testCase
		workspace *workspace.CreateWorkspaceCommand
	}{
		{
			testCase: testCase{
				description:    "should fail to create workspace for unauthorized user",
				path:           "/api/v1/orgs/%d/workspaces",
				expectedStatus: http.StatusUnauthorized,
				expectedBody:   "401 Unauthorized",
				authenticate:   false,
				toUrl:          func(path string) string { return fmt.Sprint(baseUrl, fmt.Sprintf(path, orgId)) },
				decodeBody: func(r io.Reader) (any, error) {
					var res string
					err := json.NewDecoder(r).Decode(&res)
					return res, err
				},
				encodeBody: func(body any) (*bytes.Buffer, error) {
					payloadBuf := new(bytes.Buffer)
					err := json.NewEncoder(payloadBuf).Encode(body)
					return payloadBuf, err
				},
			},
			workspace: &workspace.CreateWorkspaceCommand{},
		},
		{
			testCase: testCase{
				description:    "should create workspace 1 and return it with success",
				path:           "/api/v1/orgs/%d/workspaces",
				expectedStatus: http.StatusCreated,
				expectedBody: system.Workspace{
					ScopeId:     "workspace_1",
					Name:        "Workspace_1",
					Description: "Workspace_1 description",
					Visibility:  int32(model.PublicVisibility),
				},
				authenticate: true,
				toUrl:        func(path string) string { return fmt.Sprint(baseUrl, fmt.Sprintf(path, orgId)) },
				decodeBody: func(r io.Reader) (any, error) {
					res := &system.Workspace{}
					err := json.NewDecoder(r).Decode(&res)
					return res, err
				},
				encodeBody: func(body any) (*bytes.Buffer, error) {
					payloadBuf := &bytes.Buffer{}
					err := json.NewEncoder(payloadBuf).Encode(body)
					return payloadBuf, err
				},
			},
			workspace: &workspace.CreateWorkspaceCommand{
				OrgId:       orgId,
				CreatedBy:   "admin",
				ScopeId:     "workspace_1",
				Name:        "Workspace_1",
				Description: "Workspace_1 description",
				Visibility:  model.PublicVisibility,
			},
		},
		{
			testCase: testCase{
				description:    "should return already exists error for creating an workspace with same scope id",
				path:           "/api/v1/orgs/%d/workspaces",
				expectedStatus: http.StatusConflict,
				expectedBody:   "rpc error: code = AlreadyExists desc = workspace with same scopeId already exists",
				authenticate:   true,
				toUrl:          func(path string) string { return fmt.Sprint(baseUrl, fmt.Sprintf(path, orgId)) },
				decodeBody: func(r io.Reader) (any, error) {
					var res string
					err := json.NewDecoder(r).Decode(&res)
					return res, err
				},
				encodeBody: func(body any) (*bytes.Buffer, error) {
					payloadBuf := &bytes.Buffer{}
					err := json.NewEncoder(payloadBuf).Encode(body)
					return payloadBuf, err
				},
			},
			workspace: &workspace.CreateWorkspaceCommand{
				OrgId:       orgId,
				CreatedBy:   "admin",
				ScopeId:     "workspace_1",
				Name:        "Workspace_1",
				Description: "Workspace_1 description",
				Visibility:  model.PublicVisibility,
			},
		},
		{
			testCase: testCase{
				description:    "should create workspace 2 and return it with success",
				path:           "/api/v1/orgs/%d/workspaces",
				expectedStatus: http.StatusCreated,
				expectedBody: system.Workspace{
					ScopeId:     "workspace_2",
					Name:        "Workspace_2",
					Description: "Workspace_2 description",
					Visibility:  int32(model.PublicVisibility),
				},
				authenticate: true,
				toUrl:        func(path string) string { return fmt.Sprint(baseUrl, fmt.Sprintf(path, orgId)) },
				decodeBody: func(r io.Reader) (any, error) {
					res := &system.Workspace{}
					err := json.NewDecoder(r).Decode(res)
					return res, err
				},
				encodeBody: func(body any) (*bytes.Buffer, error) {
					payloadBuf := &bytes.Buffer{}
					err := json.NewEncoder(payloadBuf).Encode(body)
					return payloadBuf, err
				},
			},
			workspace: &workspace.CreateWorkspaceCommand{
				OrgId:       orgId,
				CreatedBy:   "admin",
				ScopeId:     "workspace_2",
				Name:        "Workspace_2",
				Description: "Workspace_2 description",
				Visibility:  model.PublicVisibility,
			},
		},
		{
			testCase: testCase{
				description:    "should return too many requests error for quota limit exceeded",
				path:           "/api/v1/orgs/%d/workspaces",
				expectedStatus: http.StatusTooManyRequests,
				expectedBody:   "workspace quota reached",
				authenticate:   true,
				toUrl:          func(path string) string { return fmt.Sprint(baseUrl, fmt.Sprintf(path, orgId)) },
				decodeBody: func(r io.Reader) (any, error) {
					var res string
					err := json.NewDecoder(r).Decode(&res)
					return res, err
				},
				encodeBody: func(body any) (*bytes.Buffer, error) {
					payloadBuf := &bytes.Buffer{}
					err := json.NewEncoder(payloadBuf).Encode(body)
					return payloadBuf, err
				},
			},
			workspace: &workspace.CreateWorkspaceCommand{
				OrgId:       orgId,
				CreatedBy:   "admin",
				ScopeId:     "workspace_3",
				Name:        "Workspace_3",
				Description: "Workspace_3 description",
				Visibility:  model.PublicVisibility,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			reqBody, err := testCase.encodeBody(testCase.workspace)
			require.NoError(t, err)
			require.NotNil(t, reqBody)

			client := &http.Client{}
			req, err := http.NewRequest("POST", testCase.toUrl(testCase.path), reqBody)
			if err != nil {
				t.Fatal(err)
			}
			if testCase.authenticate {
				req.Header.Set("Cookie", test.GetCookie())
			}

			resp, err := client.Do(req)
			if err != nil {
				log.Fatal(err)
			}
			t.Cleanup(func() { require.NoError(t, resp.Body.Close()) })

			require.NoError(t, err)
			require.Equal(t, testCase.expectedStatus, resp.StatusCode)

			respBody, err := testCase.decodeBody(resp.Body)
			require.NoError(t, err)
			switch body := respBody.(type) {
			case string:
				require.Equal(t, testCase.expectedBody, body)
			case *system.Workspace:
				switch exp := testCase.expectedBody.(type) {
				case system.Workspace:
					require.NotNil(t, body)
					require.NotZero(t, body.Id)
					require.NotEmpty(t, body.Name)
					require.NotEmpty(t, body.Description)
					require.NotZero(t, body.Visibility)
					require.Equal(t, exp.ScopeId, body.ScopeId)
					require.Equal(t, exp.Name, body.Name)
					require.Equal(t, exp.Description, body.Description)
					require.Equal(t, exp.Visibility, body.Visibility)
					require.Equal(t, fmt.Sprintf("/org/%s/%s", exp.OrgScopeId, exp.ScopeId), body.Path)
				}
			default:
			}
		})
	}
}

func TestQueryWorkspaceListApiAfterCreation(t *testing.T) {
	testCases := []testCase{
		{
			description:    "should return two previously created workspaces for authenticated user",
			path:           "/api/v1/workspaces",
			expectedStatus: http.StatusOK,
			expectedBody:   2, // returned slice len
			authenticate:   true,
			toUrl:          func(path string) string { return fmt.Sprint(baseUrl, path) },
			decodeBody: func(r io.Reader) (any, error) {
				var res []*system.Workspace
				err := json.NewDecoder(r).Decode(&res)
				return res, err
			},
		},
		{
			description:    "should return empty public workspaces if any for unauthenticated user",
			path:           "/api/v1/workspaces",
			expectedStatus: http.StatusOK,
			expectedBody:   0, // returned slice len
			authenticate:   false,
			toUrl:          func(path string) string { return fmt.Sprint(baseUrl, path) },
			decodeBody: func(r io.Reader) (any, error) {
				var res []*system.Workspace
				err := json.NewDecoder(r).Decode(&res)
				return res, err
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			client := &http.Client{}
			req, err := http.NewRequest("GET", testCase.toUrl(testCase.path), nil)
			if err != nil {
				t.Fatal(err)
			}
			if testCase.authenticate {
				req.Header.Set("Cookie", test.GetCookie())
			}

			resp, err := client.Do(req)
			if err != nil {
				t.Fatal(err)
			}
			t.Cleanup(func() { require.NoError(t, resp.Body.Close()) })

			require.NoError(t, err)
			require.Equal(t, testCase.expectedStatus, resp.StatusCode)

			respBody, err := testCase.decodeBody(resp.Body)
			require.NoError(t, err)
			switch body := respBody.(type) {
			case []*system.Workspace:
				require.Equal(t, testCase.expectedBody, len(body))
			}
		})
	}
}

func TestMutationDeleteWorkspaceApi(t *testing.T) {
	var workspaceList []*system.Workspace = []*system.Workspace{}

	t.Run("should get all available workspace list for authenticated user", func(t *testing.T) {
		client := &http.Client{}
		req, err := http.NewRequest("GET", fmt.Sprint(baseUrl, "/api/v1/workspaces"), nil)
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Set("Cookie", test.GetCookie())

		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		t.Cleanup(func() { require.NoError(t, resp.Body.Close()) })

		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)

		err = json.NewDecoder(resp.Body).Decode(&workspaceList)
		require.NoError(t, err)
		require.NotNil(t, workspaceList)
		require.Len(t, workspaceList, 2)
	})

	var testCases = []testCase{
		{
			description:    "should delete workspace for authorized user",
			path:           "/api/v1/orgs/%d/workspaces/%d",
			expectedStatus: http.StatusNoContent,
			expectedBody:   nil,
			authenticate:   true,
			toUrl:          func(path string) string { return fmt.Sprint(baseUrl, path) },
		},
	}

	for _, ws := range workspaceList {
		for _, testCase := range testCases {
			t.Run(testCase.description, func(t *testing.T) {
				client := &http.Client{}
				req, err := http.NewRequest("DELETE", fmt.Sprintf(testCase.toUrl(testCase.path), ws.OrgId, ws.Id), nil)
				if err != nil {
					t.Fatal(err)
				}
				if testCase.authenticate {
					req.Header.Set("Cookie", test.GetCookie())
				}

				resp, err := client.Do(req)
				if err != nil {
					log.Fatal(err)
				}
				t.Cleanup(func() { require.NoError(t, resp.Body.Close()) })

				require.NoError(t, err)
				require.Equal(t, testCase.expectedStatus, resp.StatusCode)
			})
		}
	}
}

func TestMutationDeleteOrgApi(t *testing.T) {
	var testCases = []testCase{
		{
			description:    "should delete org for aithenticated user",
			path:           "/api/v1/orgs/%v",
			expectedStatus: http.StatusNoContent,
			expectedBody:   nil,
			authenticate:   true,
			toUrl:          func(path string) string { return fmt.Sprint(baseUrl, path) },
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			client := &http.Client{}
			req, err := http.NewRequest("DELETE", fmt.Sprintf(testCase.toUrl(testCase.path), orgId), nil)
			if err != nil {
				t.Fatal(err)
			}
			if testCase.authenticate {
				req.Header.Set("Cookie", test.GetCookie())
			}

			resp, err := client.Do(req)
			if err != nil {
				log.Fatal(err)
			}
			t.Cleanup(func() { require.NoError(t, resp.Body.Close()) })

			require.NoError(t, err)
			require.Equal(t, testCase.expectedStatus, resp.StatusCode)
		})
	}
}