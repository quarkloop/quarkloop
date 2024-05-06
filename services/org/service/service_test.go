package org

import (
	"context"
	"fmt"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"

	"github.com/quarkloop/quarkloop/pkg/grpc/v1/system"
	orgGrpc "github.com/quarkloop/quarkloop/pkg/grpc/v1/system/org"
	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/test"
	grpcserver "github.com/quarkloop/quarkloop/pkg/test/grpcserver"
	"github.com/quarkloop/quarkloop/services/org/store"
)

var (
	ctx      context.Context
	conn     *pgxpool.Pool
	client   orgGrpc.OrgServiceClient
	orgCount int = 10
)

func init() {
	grpcServer := grpcserver.NewSystemGrpcServer()
	{
		orgSvc := NewOrgService(store.NewOrgStore(grpcServer.GetTestDbConn()))
		orgSvc.RegisterService(grpcServer.GetServer())

		ctx = grpcServer.GetContext()
		conn = grpcServer.GetTestDbConn()
	}
	grpcServer.StartGrpcServer()
	grpcServer.ConnectGrpcServer()
	{
		client = orgGrpc.NewOrgServiceClient(grpcServer.GetClient())
	}
}

// func init() {
// 	ctx, conn = test.InitTestSystemDB()
// 	orgSvc := NewOrgService(store.NewOrgStore(conn))

// 	lis = bufconn.Listen(bufSize)
// 	s := grpc.NewServer()
// 	orgSvc.RegisterService(s)
// 	go func() {
// 		if err := s.Serve(lis); err != nil {
// 			log.Fatalf("Server exited with error: %v", err)
// 		}
// 	}()

// 	clientConn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
// 	if err != nil {
// 		log.Fatalf("Failed to dial bufnet: %v", err)
// 	}

// 	client = orgGrpc.NewOrgServiceClient(clientConn)
// }

// func bufDialer(context.Context, string) (net.Conn, error) {
// 	return lis.Dial()
// }

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
				CreatedBy: fmt.Sprintf("admin_%d", i),
				Payload: &system.OrgMutation{
					ScopeId:     fmt.Sprintf("quarkloop_%d", i),
					Name:        fmt.Sprintf("Quarkloop_%d", i),
					Description: fmt.Sprintf("Quarkloop Corporation #%d", i),
					Visibility:  model.PublicVisibility.ToString(),
				},
			}
			reply, err := client.CreateOrg(ctx, cmd)
			org := reply.Data

			require.NoError(t, err)
			require.NotNil(t, org)
			require.NotZero(t, org.Id)
			require.NotEmpty(t, org.ScopeId)
			require.NotEmpty(t, org.Name)
			require.NotEmpty(t, org.Description)
			require.NotZero(t, org.Visibility)
			require.Equal(t, cmd.Payload.ScopeId, org.ScopeId)
			require.Equal(t, cmd.Payload.Name, org.Name)
			require.Equal(t, cmd.Payload.Description, org.Description)
			require.Equal(t, cmd.Payload.Visibility, org.Visibility)
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
