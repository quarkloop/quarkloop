package org

import (
	"context"
	"fmt"
	"log"
	"net"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	orgGrpc "github.com/quarkloop/quarkloop/pkg/grpc/v1/system/org"
	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/test"
	"github.com/quarkloop/quarkloop/services/org/store"
)

const bufSize = 1024 * 1024

var (
	ctx      context.Context
	conn     *pgx.Conn
	lis      *bufconn.Listener
	client   orgGrpc.OrgServiceClient
	orgCount int = 10
)

func init() {
	ctx, conn = test.InitTestSystemDB()
	orgSvc := NewOrgService(store.NewOrgStore(conn))

	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	orgSvc.RegisterService(s)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()

	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial bufnet: %v", err)
	}

	client = orgGrpc.NewOrgServiceClient(conn)
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
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

func TestMutationCreateOrg(t *testing.T) {
	t.Run("should create org with scopeId", func(t *testing.T) {
		for i := 0; i < orgCount; i++ {
			cmd := &orgGrpc.CreateOrgCommand{
				ScopeId:     fmt.Sprintf("quarkloop_%d", i),
				Name:        fmt.Sprintf("Quarkloop_%d", i),
				Description: fmt.Sprintf("Quarkloop Corporation #%d", i),
				CreatedBy:   fmt.Sprintf("admin_%d", i),
				Visibility:  int32(model.PublicVisibility),
			}
			reply, err := client.CreateOrg(ctx, cmd)
			org := reply.Org

			require.NoError(t, err)
			require.NotNil(t, org)
			require.NotEmpty(t, org.ScopeId)
			require.NotEmpty(t, org.Name)
			require.NotEmpty(t, org.Description)
			require.NotEmpty(t, org.CreatedBy)
			require.NotZero(t, org.Visibility)
			require.Equal(t, cmd.ScopeId, org.ScopeId)
			require.Equal(t, cmd.Name, org.Name)
			require.Equal(t, cmd.Description, org.Description)
			require.Equal(t, cmd.Visibility, org.Visibility)
			require.Equal(t, cmd.CreatedBy, org.CreatedBy)
		}
	})

	t.Run("should  get org list return full", func(t *testing.T) {
		orgList, err := test.GetFullOrgList(ctx, conn)

		require.NoError(t, err)
		require.NotZero(t, len(orgList))
		require.Equal(t, orgCount, len(orgList))
	})
}

// func TestGetOrgList(t *testing.T) {
// 	resp, err := client.GetOrgList(ctx, &orgGrpc.GetOrgListQuery{OrgIdList: []int64{}})
// 	if err != nil {
// 		t.Fatalf("TestGetOrgList failed: %v", err)
// 	}

// 	require.Equal(t, orgCount, len(resp.OrgList))
// }

func TestMutationDeleteOrg(t *testing.T) {
	t.Run("should delete org with scopeId", func(t *testing.T) {
		orgList, err := test.GetFullOrgList(ctx, conn)
		require.NoError(t, err)

		for _, org := range orgList {
			cmd := &orgGrpc.DeleteOrgByIdCommand{OrgId: org.Id}
			_, err := client.DeleteOrgById(ctx, cmd)

			require.NoError(t, err)
			{
				cmd := &orgGrpc.GetOrgByIdQuery{OrgId: org.Id}
				_, err := client.GetOrgById(ctx, cmd)
				require.Error(t, err)
			}
		}
	})

	t.Run("should get org list return full", func(t *testing.T) {
		orgList, err := test.GetFullOrgList(ctx, conn)

		require.NoError(t, err)
		require.Zero(t, len(orgList))
	})
}
