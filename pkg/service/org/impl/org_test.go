package org_impl_test

import (
	"context"
	"fmt"
	"log"
	"net"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/org"
	org_impl "github.com/quarkloop/quarkloop/pkg/service/org/impl"
	"github.com/quarkloop/quarkloop/pkg/service/org/store"
	"github.com/quarkloop/quarkloop/pkg/test"
	"github.com/quarkloop/quarkloop/service/system"
	"github.com/stretchr/testify/require"
)

const bufSize = 1024 * 1024

var (
	ctx      context.Context
	conn     *pgx.Conn
	lis      *bufconn.Listener
	client   system.OrgServiceClient
	orgCount int = 10
)

func init() {
	ctx, conn = test.InitTestSystemDB()
	orgService := org_impl.NewOrgService(store.NewOrgStore(conn))

	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	system.RegisterOrgServiceServer(s, orgService)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()

	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial bufnet: %v", err)
	}
	client = system.NewOrgServiceClient(conn)
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func timer(t *testing.T, name string) func() {
	start := time.Now()
	return func() {
		finish := time.Since(start)
		t.Log("\n", name, " took ", finish, "\n")
	}
}

func TestMutationTruncateTables(t *testing.T) {
	t.Run("truncate tables", func(t *testing.T) {
		err := test.TruncateSystemDBTables(ctx, conn)
		require.NoError(t, err)
	})

	t.Run("get org list return empty after truncating tables", func(t *testing.T) {
		orgList, err := test.GetFullOrgList(ctx, conn)
		require.NoError(t, err)
		require.Zero(t, len(orgList))
		require.Equal(t, 0, len(orgList))
	})
}

func TestMutationCreateOrg(t *testing.T) {
	store := store.NewOrgStore(conn)

	t.Run("create org with scopeId", func(t *testing.T) {
		for i := 0; i < orgCount; i++ {
			cmd := &org.CreateOrgCommand{
				ScopeId:     fmt.Sprintf("quarkloop_%d", i),
				Name:        fmt.Sprintf("Quarkloop_%d", i),
				Description: fmt.Sprintf("Quarkloop Corporation #%d", i),
				CreatedBy:   fmt.Sprintf("admin_%d", i),
				Visibility:  model.PublicVisibility,
			}
			org, err := store.CreateOrg(ctx, cmd)

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

	t.Run("get org list return full", func(t *testing.T) {
		orgList, err := test.GetFullOrgList(ctx, conn)

		require.NoError(t, err)
		require.NotZero(t, len(orgList))
		require.Equal(t, orgCount, len(orgList))
	})
}

func TestGetOrgList(t *testing.T) {
	defer timer(t, "TestGetOrgList")()

	resp, err := client.GetOrgList(ctx, &system.GetOrgListQuery{UserId: 0})
	if err != nil {
		t.Fatalf("TestGetOrgList failed: %v", err)
	}

	require.Equal(t, len(resp.OrgList), orgCount)
}
