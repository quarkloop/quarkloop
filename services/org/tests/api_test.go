package tests

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
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"

	orgGrpc "github.com/quarkloop/quarkloop/pkg/grpc/v1/system/org"
	orgApi "github.com/quarkloop/quarkloop/services/org/api"
	orgErrors "github.com/quarkloop/quarkloop/services/org/errors"

	"github.com/quarkloop/quarkloop/apiserver/tests"
	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/test"
	"github.com/quarkloop/quarkloop/services/quota"
)

var (
	ctx    context.Context
	conn   *pgxpool.Pool
	client orgGrpc.OrgServiceClient
	authz  *authzed.Client
)

func init() {
	authz = tests.InitAuthzClient()
	ctx, conn, client, _ = tests.StartGrpcServer()
	tests.StartApiServer()
}

func TestMutationTruncateTables(t *testing.T) {
	t.Run("should truncate tables", func(t *testing.T) {
		err := test.TruncateSystemDBTables(ctx, conn)
		require.NoError(t, err)
	})

	t.Run("should get org list return empty after truncating tables", func(t *testing.T) {
		orgList, err := test.GetFullOrgList(ctx, conn)
		require.NoError(t, err)
		require.Zero(t, len(orgList))
		require.Equal(t, 0, len(orgList))
	})
}

func TestMutationCreateSchema(t *testing.T) {
	t.Run("create schema", func(t *testing.T) {
		resp, err := authz.WriteSchema(ctx, &v1.WriteSchemaRequest{
			Schema: tests.Schema,
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

func TestQueryOrgListApi(t *testing.T) {
	testCases := []tests.TestCase{
		{
			Description:    "should return empty public and private orgs for authenticated user",
			Path:           "/api/v1/manage/orgs",
			ExpectedStatus: http.StatusOK,
			ExpectedBody:   &orgApi.GetOrgListDTO{Data: []*model.Org{}},
			Authenticate:   true,
			ToUrl:          func(path string) string { return fmt.Sprint(tests.BaseUrl, path) },
			DecodeBody: func(r io.Reader) (any, error) {
				var res *orgApi.GetOrgListDTO
				err := json.NewDecoder(r).Decode(&res)
				return res, err
			},
		},
		{
			Description:    "should return empty public orgs for unauthenticated user",
			Path:           "/api/v1/manage/orgs",
			ExpectedStatus: http.StatusOK,
			ExpectedBody:   &orgApi.GetOrgListDTO{Data: []*model.Org{}},
			Authenticate:   false,
			ToUrl:          func(path string) string { return fmt.Sprint(tests.BaseUrl, path) },
			DecodeBody: func(r io.Reader) (any, error) {
				var res *orgApi.GetOrgListDTO
				err := json.NewDecoder(r).Decode(&res)
				return res, err
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Description, func(t *testing.T) {
			client := &http.Client{}
			req, err := http.NewRequest("GET", testCase.ToUrl(testCase.Path), nil)
			if err != nil {
				t.Fatal(err)
			}
			if testCase.Authenticate {
				req.Header.Set("Cookie", test.GetCookie())
			}

			resp, err := client.Do(req)
			if err != nil {
				t.Fatal(err)
			}
			t.Cleanup(func() { require.NoError(t, resp.Body.Close()) })

			require.NoError(t, err)
			require.Equal(t, testCase.ExpectedStatus, resp.StatusCode)

			body, err := testCase.DecodeBody(resp.Body)
			require.NoError(t, err)
			require.Equal(t, testCase.ExpectedBody, body)
		})
	}

	testCases = []tests.TestCase{
		{
			Description:    "should return bad request with org sid 0 for authenticated user",
			Path:           "/api/v1/manage/0/workspaces",
			ExpectedStatus: http.StatusBadRequest,
			ExpectedBody:   `Key: 'GetWorkspaceListQuery.OrgSid' Error:Field validation for 'OrgSid' failed on the 'sid' tag`,
			Authenticate:   true,
			ToUrl:          func(path string) string { return fmt.Sprint(tests.BaseUrl, path) },
			DecodeBody: func(r io.Reader) (any, error) {
				var res string
				err := json.NewDecoder(r).Decode(&res)
				return res, err
			},
		},
		{
			Description:    "should return bad request with org sid 0 for unauthenticated user",
			Path:           "/api/v1/manage/0/workspaces",
			ExpectedStatus: http.StatusBadRequest,
			ExpectedBody:   `Key: 'GetWorkspaceListQuery.OrgSid' Error:Field validation for 'OrgSid' failed on the 'sid' tag`,
			Authenticate:   false,
			ToUrl:          func(path string) string { return fmt.Sprint(tests.BaseUrl, path) },
			DecodeBody: func(r io.Reader) (any, error) {
				var res string
				err := json.NewDecoder(r).Decode(&res)
				return res, err
			},
		},
		{
			Description:    "should return not found for unknown org sid and authenticated user",
			Path:           "/api/v1/manage/%s/workspaces",
			ExpectedStatus: http.StatusNotFound,
			ExpectedBody:   orgErrors.ErrOrgNotFound.Error(),
			Authenticate:   true,
			ToUrl:          func(path string) string { return fmt.Sprint(tests.BaseUrl, fmt.Sprintf(path, "org_1")) },
			DecodeBody: func(r io.Reader) (any, error) {
				var res string
				err := json.NewDecoder(r).Decode(&res)
				return res, err
			},
		},
		{
			Description:    "should return not found for unknown org sid and unauthenticated user",
			Path:           "/api/v1/manage/%s/workspaces",
			ExpectedStatus: http.StatusNotFound,
			ExpectedBody:   orgErrors.ErrOrgNotFound.Error(),
			Authenticate:   false,
			ToUrl:          func(path string) string { return fmt.Sprint(tests.BaseUrl, fmt.Sprintf(path, "org_1")) },
			DecodeBody: func(r io.Reader) (any, error) {
				var res string
				err := json.NewDecoder(r).Decode(&res)
				return res, err
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Description, func(t *testing.T) {
			client := &http.Client{}
			req, err := http.NewRequest("GET", testCase.ToUrl(testCase.Path), nil)
			if err != nil {
				t.Fatal(err)
			}
			if testCase.Authenticate {
				req.Header.Set("Cookie", test.GetCookie())
			}

			resp, err := client.Do(req)
			if err != nil {
				t.Fatal(err)
			}
			t.Cleanup(func() { require.NoError(t, resp.Body.Close()) })

			require.NoError(t, err)
			require.Equal(t, testCase.ExpectedStatus, resp.StatusCode)

			body, err := testCase.DecodeBody(resp.Body)
			require.NoError(t, err)
			require.Equal(t, testCase.ExpectedBody, body)
		})
	}
}

func TestMutationCreateOrgApi(t *testing.T) {
	var testCases []struct {
		tests.TestCase
		org *orgApi.CreateOrgCommand
	} = []struct {
		tests.TestCase
		org *orgApi.CreateOrgCommand
	}{
		{
			TestCase: tests.TestCase{
				Description:    "should fail to create org for unauthenticated user",
				Path:           "/api/v1/manage/orgs",
				ExpectedStatus: http.StatusUnauthorized,
				ExpectedBody:   "401 Unauthenticated",
				Authenticate:   false,
				ToUrl:          func(path string) string { return fmt.Sprint(tests.BaseUrl, path) },
				DecodeBody: func(r io.Reader) (any, error) {
					var res string
					err := json.NewDecoder(r).Decode(&res)
					return res, err
				},
				EncodeBody: func(body any) (*bytes.Buffer, error) {
					payloadBuf := new(bytes.Buffer)
					err := json.NewEncoder(payloadBuf).Encode(body)
					return payloadBuf, err
				},
			},
			org: &orgApi.CreateOrgCommand{},
		},
		{
			TestCase: tests.TestCase{
				Description:    "should create org 1 and return it with success",
				Path:           "/api/v1/manage/orgs",
				ExpectedStatus: http.StatusCreated,
				ExpectedBody: &orgApi.CreateOrgDTO{
					Data: &model.Org{
						ScopeId:     "org_1",
						Name:        "Org_1",
						Description: "Org_1 description",
						Visibility:  model.PublicVisibility,
					},
				},
				Authenticate: true,
				ToUrl:        func(path string) string { return fmt.Sprint(tests.BaseUrl, path) },
				DecodeBody: func(r io.Reader) (any, error) {
					res := &model.Org{}
					err := json.NewDecoder(r).Decode(res)
					return res, err
				},
				EncodeBody: func(body any) (*bytes.Buffer, error) {
					payloadBuf := &bytes.Buffer{}
					err := json.NewEncoder(payloadBuf).Encode(body)
					return payloadBuf, err
				},
			},
			org: &orgApi.CreateOrgCommand{
				Payload: &model.OrgMutation{
					ScopeId:     "org_1",
					Name:        "Org_1",
					Description: "Org_1 description",
					Visibility:  model.PublicVisibility,
				},
			},
		},
		{
			TestCase: tests.TestCase{
				Description:    "should return already exists error for creating an org with same scope id",
				Path:           "/api/v1/manage/orgs",
				ExpectedStatus: http.StatusConflict,
				ExpectedBody:   "rpc error: code = AlreadyExists desc = org with same scopeId already exists",
				Authenticate:   true,
				ToUrl:          func(path string) string { return fmt.Sprint(tests.BaseUrl, path) },
				DecodeBody: func(r io.Reader) (any, error) {
					var res string
					err := json.NewDecoder(r).Decode(&res)
					return res, err
				},
				EncodeBody: func(body any) (*bytes.Buffer, error) {
					payloadBuf := &bytes.Buffer{}
					err := json.NewEncoder(payloadBuf).Encode(body)
					return payloadBuf, err
				},
			},
			org: &orgApi.CreateOrgCommand{
				Payload: &model.OrgMutation{
					ScopeId:     "org_1",
					Name:        "Org_1",
					Description: "Org_1 description",
					Visibility:  model.PublicVisibility,
				},
			},
		},
		{
			TestCase: tests.TestCase{
				Description:    "should create org 2 and return it with success",
				Path:           "/api/v1/manage/orgs",
				ExpectedStatus: http.StatusCreated,
				ExpectedBody: model.Org{
					ScopeId:     "org_2",
					Name:        "Org_2",
					Description: "Org_2 description",
					Visibility:  model.PublicVisibility,
				},
				Authenticate: true,
				ToUrl:        func(path string) string { return fmt.Sprint(tests.BaseUrl, path) },
				DecodeBody: func(r io.Reader) (any, error) {
					res := &model.Org{}
					err := json.NewDecoder(r).Decode(res)
					return res, err
				},
				EncodeBody: func(body any) (*bytes.Buffer, error) {
					payloadBuf := &bytes.Buffer{}
					err := json.NewEncoder(payloadBuf).Encode(body)
					return payloadBuf, err
				},
			},
			org: &orgApi.CreateOrgCommand{
				Payload: &model.OrgMutation{
					ScopeId:     "org_2",
					Name:        "Org_2",
					Description: "Org_2 description",
					Visibility:  model.PublicVisibility,
				},
			},
		},
		{
			TestCase: tests.TestCase{
				Description:    "should return too many requests error for quota limit exceeded",
				Path:           "/api/v1/manage/orgs",
				ExpectedStatus: http.StatusTooManyRequests,
				ExpectedBody:   "org quota reached",
				Authenticate:   true,
				ToUrl:          func(path string) string { return fmt.Sprint(tests.BaseUrl, path) },
				DecodeBody: func(r io.Reader) (any, error) {
					var res string
					err := json.NewDecoder(r).Decode(&res)
					return res, err
				},
				EncodeBody: func(body any) (*bytes.Buffer, error) {
					payloadBuf := &bytes.Buffer{}
					err := json.NewEncoder(payloadBuf).Encode(body)
					return payloadBuf, err
				},
			},
			org: &orgApi.CreateOrgCommand{
				Payload: &model.OrgMutation{
					ScopeId:     "org_3",
					Name:        "Org_3",
					Description: "Org_3 description",
					Visibility:  model.PublicVisibility,
				},
			},
		},
	}

	// adapt quota limit with test (after 2 successful post request should throw quota error)
	quota.OrgQuotaLimit = 2

	for _, testCase := range testCases {
		t.Run(testCase.Description, func(t *testing.T) {
			reqBody, err := testCase.EncodeBody(testCase.org)
			require.NoError(t, err)
			require.NotNil(t, reqBody)

			client := &http.Client{}
			req, err := http.NewRequest("POST", testCase.ToUrl(testCase.Path), reqBody)
			if err != nil {
				t.Fatal(err)
			}
			if testCase.Authenticate {
				req.Header.Set("Cookie", test.GetCookie())
			}

			resp, err := client.Do(req)
			if err != nil {
				log.Fatal(err)
			}
			t.Cleanup(func() { require.NoError(t, resp.Body.Close()) })

			require.NoError(t, err)
			require.Equal(t, testCase.ExpectedStatus, resp.StatusCode)

			respBody, err := testCase.DecodeBody(resp.Body)
			require.NoError(t, err)

			switch body := respBody.(type) {
			case string:
				require.Equal(t, testCase.ExpectedBody, body)
			case *orgApi.CreateOrgDTO:
				switch exp := testCase.ExpectedBody.(type) {
				case *orgApi.CreateOrgDTO:
					require.NotNil(t, body)
					require.NotZero(t, body.Data.Id)
					require.NotEmpty(t, body.Data.Name)
					require.NotEmpty(t, body.Data.Description)
					require.NotZero(t, body.Data.Visibility)
					require.Equal(t, exp.Data.ScopeId, body.Data.ScopeId)
					require.Equal(t, exp.Data.Name, body.Data.Name)
					require.Equal(t, exp.Data.Description, body.Data.Description)
					require.Equal(t, exp.Data.Visibility, body.Data.Visibility)
					require.Equal(t, fmt.Sprint("/org/", exp.Data.ScopeId), body.Data.Path)
				}
			default:
			}
		})
	}
}

func TestQueryOrgListApiAfterCreation(t *testing.T) {
	testCases := []tests.TestCase{
		{
			Description:    "should return two previously created orgs for authenticated user",
			Path:           "/api/v1/manage/orgs",
			ExpectedStatus: http.StatusOK,
			ExpectedBody:   2, // returned slice len
			Authenticate:   true,
			ToUrl:          func(path string) string { return fmt.Sprint(tests.BaseUrl, path) },
			DecodeBody: func(r io.Reader) (any, error) {
				var res *orgApi.GetOrgListDTO
				err := json.NewDecoder(r).Decode(&res)
				return res, err
			},
		},
		{
			Description:    "should return empty public orgs if any for unauthenticated user",
			Path:           "/api/v1/manage/orgs",
			ExpectedStatus: http.StatusOK,
			ExpectedBody:   0, // returned slice len
			Authenticate:   false,
			ToUrl:          func(path string) string { return fmt.Sprint(tests.BaseUrl, path) },
			DecodeBody: func(r io.Reader) (any, error) {
				var res *orgApi.GetOrgListDTO
				err := json.NewDecoder(r).Decode(&res)
				return res, err
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Description, func(t *testing.T) {
			client := &http.Client{}
			req, err := http.NewRequest("GET", testCase.ToUrl(testCase.Path), nil)
			if err != nil {
				t.Fatal(err)
			}
			if testCase.Authenticate {
				req.Header.Set("Cookie", test.GetCookie())
			}

			resp, err := client.Do(req)
			if err != nil {
				t.Fatal(err)
			}
			t.Cleanup(func() { require.NoError(t, resp.Body.Close()) })

			require.NoError(t, err)
			require.Equal(t, testCase.ExpectedStatus, resp.StatusCode)

			respBody, err := testCase.DecodeBody(resp.Body)
			require.NoError(t, err)
			switch body := respBody.(type) {
			case *orgApi.GetOrgListDTO:
				require.Equal(t, testCase.ExpectedBody, len(body.Data))
			}
		})
	}
}

func Test_Org_MutationDeleteOrgApi(t *testing.T) {
	var orgList *orgApi.GetOrgListDTO

	t.Run("should get all available org list for authenticated user", func(t *testing.T) {
		client := &http.Client{}
		req, err := http.NewRequest("GET", fmt.Sprint(tests.BaseUrl, "/api/v1/manage/orgs"), nil)
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

		err = json.NewDecoder(resp.Body).Decode(&orgList)
		require.NoError(t, err)
		require.NotNil(t, orgList)
		require.Len(t, orgList.Data, 2)
	})

	var testCases = []tests.TestCase{
		{
			Description:    "should delete org for authorized user",
			Path:           "/api/v1/manage/%s",
			ExpectedStatus: http.StatusNoContent,
			ExpectedBody:   nil,
			Authenticate:   true,
			ToUrl:          func(path string) string { return fmt.Sprint(tests.BaseUrl, path) },
		},
	}

	for _, o := range orgList.Data {
		for _, testCase := range testCases {
			t.Run(testCase.Description, func(t *testing.T) {
				client := &http.Client{}
				req, err := http.NewRequest("DELETE", fmt.Sprintf(testCase.ToUrl(testCase.Path), o.ScopeId), nil)
				if err != nil {
					t.Fatal(err)
				}
				if testCase.Authenticate {
					req.Header.Set("Cookie", test.GetCookie())
				}

				resp, err := client.Do(req)
				if err != nil {
					log.Fatal(err)
				}
				t.Cleanup(func() { require.NoError(t, resp.Body.Close()) })

				require.NoError(t, err)
				require.Equal(t, testCase.ExpectedStatus, resp.StatusCode)
			})
		}
	}
}

// var schema = `
// definition user {}

// definition org {
//     relation owner: user

//     relation admin: user
//     relation contributor: user
//     relation viewer: user

//     permission all = admins + contributors + viewer
//     permission admins = owner + admin
//     permission contributors = contributor
//     permission viewers = viewer

//     // org
//     permission read = all
//     permission update = admins + contributors
//     permission delete = admins
//     permission settings_read = admins
//     permission settings_update = admins
//     permission quota_read = all
//     permission quota_create = admins
//     permission quota_update = admins
//     permission quota_delete = admins
//     permission user_read = all
//     permission user_create = admins
//     permission user_update = admins
//     permission user_delete = admins

//     // workspace
//     permission create_workspace = admins
// }

// definition workspace {
//     relation parent: org

//     relation admin: user
//     relation contributor: user
//     relation viewer: user

//     permission all = admins + contributors + viewers
//     permission admins = admin + parent->admins
//     permission contributors = contributor + parent->contributors
//     permission viewers = viewer + parent->viewers

//     // workspace
//     permission read = all
//     permission update = admins + contributors
//     permission delete = admins
//     permission settings_read = admins
//     permission settings_update = admins
//     permission quota_read = all
//     permission quota_create = admins
//     permission quota_update = admins
//     permission quota_delete = admins
//     permission user_read = all
//     permission user_create = admins
//     permission user_update = admins
//     permission user_delete = admins

//     // project
//     permission create_project = admins
// }

// definition project {
//     relation parent: workspace

//     relation admin: user
//     relation contributor: user
//     relation viewer: user

//     permission all = admins + contributors + viewer
//     permission admins = admin + parent->admins
//     permission contributors = contributor + parent->contributors
//     permission viewers = viewer + parent->viewers

//     // project
//     permission read = all
//     permission update = admins + contributors
//     permission delete = admins
//     permission settings_read = admins
//     permission settings_update = admins
//     permission quota_read = all
//     permission quota_create = admins
//     permission quota_update = admins
//     permission quota_delete = admins
//     permission user_read = all
//     permission user_create = admins
//     permission user_update = admins
//     permission user_delete = admins
// }
// `
