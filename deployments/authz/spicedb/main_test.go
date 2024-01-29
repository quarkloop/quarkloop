package main_test

import (
	"context"
	_ "embed"
	"errors"
	"io"
	"log"
	"testing"

	v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/authzed/authzed-go/v1"
	"github.com/authzed/grpcutil"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	ctx   context.Context
	authz *authzed.Client
)

//go:embed model/schema.zed
var schema string

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

// func TestMutationCreateRelationship(t *testing.T) {
// 	t.Run("create relationships", func(t *testing.T) {
// 		relationships := []*v1.RelationshipUpdate{
// 			{
// 				Operation: v1.RelationshipUpdate_OPERATION_CREATE,
// 				Relationship: &v1.Relationship{
// 					Resource: &v1.ObjectReference{
// 						ObjectType: "workspace",
// 						ObjectId:   "it_department",
// 					},
// 					Relation: "parent",
// 					Subject: &v1.SubjectReference{
// 						Object: &v1.ObjectReference{
// 							ObjectType: "org",
// 							ObjectId:   "acme",
// 						},
// 					},
// 				},
// 			},
// 			{
// 				Operation: v1.RelationshipUpdate_OPERATION_CREATE,
// 				Relationship: &v1.Relationship{
// 					Resource: &v1.ObjectReference{
// 						ObjectType: "org",
// 						ObjectId:   "acme",
// 					},
// 					Relation: "admin",
// 					Subject: &v1.SubjectReference{
// 						Object: &v1.ObjectReference{
// 							ObjectType: "user",
// 							ObjectId:   "org_admin",
// 						},
// 					},
// 				},
// 			},
// 			{
// 				Operation: v1.RelationshipUpdate_OPERATION_CREATE,
// 				Relationship: &v1.Relationship{
// 					Resource: &v1.ObjectReference{
// 						ObjectType: "org",
// 						ObjectId:   "acme",
// 					},
// 					Relation: "viewer",
// 					Subject: &v1.SubjectReference{
// 						Object: &v1.ObjectReference{
// 							ObjectType: "user",
// 							ObjectId:   "org_viewer",
// 						},
// 					},
// 				},
// 			},
// 			{
// 				Operation: v1.RelationshipUpdate_OPERATION_CREATE,
// 				Relationship: &v1.Relationship{
// 					Resource: &v1.ObjectReference{
// 						ObjectType: "workspace",
// 						ObjectId:   "it_department",
// 					},
// 					Relation: "contributor",
// 					Subject: &v1.SubjectReference{
// 						Object: &v1.ObjectReference{
// 							ObjectType: "user",
// 							ObjectId:   "ws_contributor",
// 						},
// 					},
// 				},
// 			},
// 			{
// 				Operation: v1.RelationshipUpdate_OPERATION_CREATE,
// 				Relationship: &v1.Relationship{
// 					Resource: &v1.ObjectReference{
// 						ObjectType: "workspace",
// 						ObjectId:   "it_department",
// 					},
// 					Relation: "admin",
// 					Subject: &v1.SubjectReference{
// 						Object: &v1.ObjectReference{
// 							ObjectType: "user",
// 							ObjectId:   "org_viewer",
// 						},
// 					},
// 				},
// 			},
// 		}
// 		resp, err := authz.WriteRelationships(ctx, &v1.WriteRelationshipsRequest{Updates: relationships})
// 		require.NoError(t, err)

// 		token := resp.GetWrittenAt()
// 		require.NotNil(t, resp)
// 		require.NotEmpty(t, token)
// 	})
// }

func TestQueryReadRelationship(t *testing.T) {
	t.Run("read org relationships", func(t *testing.T) {
		resp, err := authz.ReadRelationships(ctx, &v1.ReadRelationshipsRequest{
			Consistency: &v1.Consistency{
				Requirement: &v1.Consistency_FullyConsistent{
					FullyConsistent: true,
				},
			},
			RelationshipFilter: &v1.RelationshipFilter{
				ResourceType: "org",
			},
		})

		require.NoError(t, err)
		require.NotNil(t, resp)

		for {
			relationship, err := resp.Recv()
			if errors.Is(err, io.EOF) {
				break
			}

			require.NoError(t, err)
			require.NotNil(t, relationship)
		}
		require.NoError(t, resp.CloseSend())
	})

	t.Run("read workspace relationships", func(t *testing.T) {
		resp, err := authz.ReadRelationships(ctx, &v1.ReadRelationshipsRequest{
			Consistency: &v1.Consistency{
				Requirement: &v1.Consistency_FullyConsistent{
					FullyConsistent: true,
				},
			},
			RelationshipFilter: &v1.RelationshipFilter{
				ResourceType: "workspace",
			},
		})

		require.NoError(t, err)
		require.NotNil(t, resp)

		for {
			relationship, err := resp.Recv()
			if errors.Is(err, io.EOF) {
				t.Log("\nReadRelationships EOF", relationship, "\n")
				break
			}

			t.Log("\nReadRelationships", relationship, "\n")

			require.NoError(t, err)
			require.NotNil(t, relationship)
		}
		require.NoError(t, resp.CloseSend())
	})
}

func TestQueryLookups(t *testing.T) {
	t.Run("lookup user relationships", func(t *testing.T) {
		resp, err := authz.LookupSubjects(context.Background(), &v1.LookupSubjectsRequest{
			Consistency: &v1.Consistency{
				Requirement: &v1.Consistency_FullyConsistent{
					FullyConsistent: true,
				},
			},
			Resource: &v1.ObjectReference{
				ObjectType: "workspace",
				ObjectId:   "it_department",
			},
			Permission:        "all",
			SubjectObjectType: "user",
			//OptionalSubjectRelation: "admin",
		})

		require.NoError(t, err)
		require.NotNil(t, resp)

		for {
			relationship, err := resp.Recv()
			if errors.Is(err, io.EOF) {
				break
			}

			//t.Log("\n", relationship, "\n")

			require.NoError(t, err)
			require.NotNil(t, relationship)
		}
		require.NoError(t, resp.CloseSend())
	})

	t.Run("lookup resource relationships", func(t *testing.T) {
		resp, err := authz.LookupResources(context.Background(), &v1.LookupResourcesRequest{
			Consistency: &v1.Consistency{
				Requirement: &v1.Consistency_FullyConsistent{
					FullyConsistent: true,
				},
			},
			ResourceObjectType: "org",
			Permission:         "all",
			Subject: &v1.SubjectReference{
				Object: &v1.ObjectReference{
					ObjectType: "user",
					ObjectId:   "org_admin",
				},
			},
			//OptionalSubjectRelation: "admin",
		})

		require.NoError(t, err)
		require.NotNil(t, resp)

		for {
			relationship, err := resp.Recv()
			if errors.Is(err, io.EOF) {
				break
			}

			t.Log("\n", relationship, "\n")

			require.NoError(t, err)
			require.NotNil(t, relationship)
		}
		require.NoError(t, resp.CloseSend())
	})

	t.Run("lookup user relationship expands", func(t *testing.T) {
		resp, err := authz.ExpandPermissionTree(context.Background(), &v1.ExpandPermissionTreeRequest{
			Consistency: &v1.Consistency{
				Requirement: &v1.Consistency_FullyConsistent{
					FullyConsistent: true,
				},
			},
			Resource: &v1.ObjectReference{
				// ObjectType: "org",
				// ObjectId:   "acme",
				ObjectType: "workspace",
				ObjectId:   "it_department",
			},
			Permission: "all",
		})

		require.NoError(t, err)
		require.NotNil(t, resp)

		expandPermissionTreeRecursive(t, resp.TreeRoot)
	})
}

func expandPermissionTreeRecursive(t *testing.T, root *v1.PermissionRelationshipTree) {
	if root == nil {
		return
	}

	inter := root.GetIntermediate()
	for _, c := range inter.GetChildren() {
		leaf := c.GetLeaf()
		if leaf != nil && len(leaf.Subjects) != 0 {
			t.Log("\n\n", c.GetExpandedRelation(), " ", leaf, " ", c.ExpandedObject)
		}
		expandPermissionTreeRecursive(t, c)
	}
}

// func TestMutationDeleteRelationship(t *testing.T) {
// 	t.Run("read org relationships", func(t *testing.T) {
// 		resources := []string{"org", "workspace", "project"}
// 		for _, resource := range resources {
// 			resp, err := authz.DeleteRelationships(ctx, &v1.DeleteRelationshipsRequest{
// 				RelationshipFilter: &v1.RelationshipFilter{
// 					ResourceType: resource,
// 				},
// 			})
// 			require.NoError(t, err)

// 			token := resp.GetDeletedAt()
// 			require.NotNil(t, resp)
// 			require.NotEmpty(t, token)
// 		}
// 	})
// }
