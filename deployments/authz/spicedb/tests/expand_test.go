package expand_test

import (
	"context"
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

func TestQueryExpand(t *testing.T) {
	t.Run("expand user relationships", func(t *testing.T) {
		parent := "org"
		parentId := "201"
		child := "workspace"
		childId := "36"

		relationships := []*v1.RelationshipUpdate{
			{
				Operation: v1.RelationshipUpdate_OPERATION_CREATE,
				Relationship: &v1.Relationship{
					Resource: &v1.ObjectReference{
						ObjectType: child,
						ObjectId:   childId,
					},
					Relation: "parent",
					Subject: &v1.SubjectReference{
						Object: &v1.ObjectReference{
							ObjectType: parent,
							ObjectId:   parentId,
						},
					},
				},
			},
		}
		_, err := authz.WriteRelationships(ctx, &v1.WriteRelationshipsRequest{Updates: relationships})
		require.NoError(t, err)

		resp, err := authz.ExpandPermissionTree(context.Background(), &v1.ExpandPermissionTreeRequest{
			Consistency: &v1.Consistency{
				Requirement: &v1.Consistency_FullyConsistent{
					FullyConsistent: true,
				},
			},
			Resource: &v1.ObjectReference{
				ObjectType: "org",
				ObjectId:   "201",
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
