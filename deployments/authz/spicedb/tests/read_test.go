package read_test

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

//go:embed schema.zed
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

			t.Log("\n=======> ReadRelationships (org) =>", relationship, "\n")

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
				break
			}

			t.Log("\n=======> ReadRelationships (workspace) =>", relationship, "\n")

			require.NoError(t, err)
			require.NotNil(t, relationship)
		}
		require.NoError(t, resp.CloseSend())
	})
}
