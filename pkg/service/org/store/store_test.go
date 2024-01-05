package store_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"

	"github.com/quarkloop/quarkloop/pkg/db"
	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/org"
	"github.com/quarkloop/quarkloop/pkg/service/org/store"
)

var conn *pgx.Conn
var ctx context.Context

func init() {
	err := godotenv.Load("/home/reza/dev/quarkloop/submodules/quarkloop/.env.development")
	if err != nil {
		log.Fatal("Error loading .env file", err.Error())
	}

	database := db.NewSystemDatabase()
	database.Connect()
	conn = database.Connection()
	ctx = context.Background()
}

const getOrgListQuery = `
SELECT 
    "id",
    "sid",
    "name",
    "description",
    "visibility",
    "createdAt",
    "createdBy",
    "updatedAt",
    "updatedBy"
FROM 
    "system"."Organization"
`

func getAllOrgList(ctx context.Context) ([]*org.Org, error) {
	rows, err := conn.Query(ctx, getOrgListQuery)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var orgList []*org.Org = []*org.Org{}
	for rows.Next() {
		var org org.Org
		err := rows.Scan(
			&org.Id,
			&org.ScopeId,
			&org.Name,
			&org.Description,
			&org.Visibility,
			&org.CreatedAt,
			&org.CreatedBy,
			&org.UpdatedAt,
			&org.UpdatedBy,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
			return nil, err
		}
		orgList = append(orgList, &org)
	}
	if err := rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
		return nil, err
	}

	return orgList, nil
}

const truncateTablesQuery = `
TRUNCATE
    "system"."Permission",
    "system"."UserRole",
    "system"."UserGroup",
    "system"."UserAssignment",
    "system"."Project",
    "system"."Workspace",
    "system"."Organization";
`

func truncateTables(ctx context.Context) error {
	_, err := conn.Exec(ctx, truncateTablesQuery)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[TRUNCATE] failed: %v\n", err)
		return err
	}
	return nil
}

func TestOrgService(t *testing.T) {
	store := store.NewOrgStore(conn)

	t.Run("truncate tables", func(t *testing.T) {
		err := truncateTables(ctx)
		require.NoError(t, err)
	})

	t.Run("CreateOrg", func(t *testing.T) {
		for i := 0; i < 10; i++ {
			cmd := &org.CreateOrgCommand{
				Name:        fmt.Sprintf("Quarkloop_%d", i),
				ScopeId:     fmt.Sprintf("quarkloop_%d", i),
				Description: fmt.Sprintf("Quarkloop Corporation #%d", i),
				CreatedBy:   fmt.Sprintf("admin_%d", i),
				Visibility:  model.PublicVisibility,
			}
			org, err := store.CreateOrg(ctx, cmd)
			require.NoError(t, err)
			require.NotEmpty(t, org)
		}
	})

	t.Run("GetOrgById after creation", func(t *testing.T) {
		orgList, err := getAllOrgList(ctx)
		require.NoError(t, err)
		require.NotZero(t, len(orgList))

		for idx, o := range orgList {
			org, err := store.GetOrgById(ctx, &org.GetOrgByIdQuery{OrgId: o.Id})
			require.NoError(t, err)
			require.NotEmpty(t, org)
			require.Equal(t, org.Name, fmt.Sprintf("Quarkloop_%d", idx))
			require.Equal(t, org.ScopeId, fmt.Sprintf("quarkloop_%d", idx))
			require.Equal(t, org.Description, fmt.Sprintf("Quarkloop Corporation #%d", idx))
			require.Equal(t, org.CreatedBy, fmt.Sprintf("admin_%d", idx))
			require.Equal(t, org.Visibility, model.PublicVisibility)
		}
	})

	t.Run("GetOrgVisibilityById after creation", func(t *testing.T) {
		orgList, err := getAllOrgList(ctx)
		require.NoError(t, err)
		require.NotZero(t, len(orgList))

		for _, o := range orgList {
			org, err := store.GetOrgVisibilityById(ctx, &org.GetOrgVisibilityByIdQuery{OrgId: o.Id})
			require.NoError(t, err)
			require.NotEmpty(t, org)
			require.Equal(t, org, model.PublicVisibility)
		}
	})

	t.Run("UpdateOrgById", func(t *testing.T) {
		orgList, err := getAllOrgList(ctx)
		require.NoError(t, err)
		require.NotZero(t, len(orgList))

		for idx, o := range orgList {
			cmd := &org.UpdateOrgByIdCommand{
				OrgId:       o.Id,
				Name:        fmt.Sprintf("Quarkloop_Updated_%d", idx),
				ScopeId:     fmt.Sprintf("quarkloop_Updated_%d", idx),
				Description: fmt.Sprintf("Quarkloop Corporation Updated #%d", idx),
				UpdatedBy:   fmt.Sprintf("admin_Updated_%d", idx),
				Visibility:  model.PrivateVisibility,
			}
			err := store.UpdateOrgById(ctx, cmd)
			require.NoError(t, err)
		}
	})

	t.Run("GetOrgById after update", func(t *testing.T) {
		orgList, err := getAllOrgList(ctx)
		require.NoError(t, err)
		require.NotZero(t, len(orgList))

		for idx, o := range orgList {
			org, err := store.GetOrgById(ctx, &org.GetOrgByIdQuery{OrgId: o.Id})
			require.NoError(t, err)
			require.NotEmpty(t, org)
			require.Equal(t, org.Name, fmt.Sprintf("Quarkloop_Updated_%d", idx))
			require.Equal(t, org.ScopeId, fmt.Sprintf("quarkloop_Updated_%d", idx))
			require.Equal(t, org.Description, fmt.Sprintf("Quarkloop Corporation Updated #%d", idx))
			require.Equal(t, org.CreatedBy, fmt.Sprintf("admin_%d", idx))
			require.Equal(t, *org.UpdatedBy, fmt.Sprintf("admin_Updated_%d", idx))
			require.Equal(t, org.Visibility, model.PrivateVisibility)
		}
	})

	t.Run("GetOrgVisibilityById after update", func(t *testing.T) {
		orgList, err := getAllOrgList(ctx)
		require.NoError(t, err)
		require.NotZero(t, len(orgList))

		for _, o := range orgList {
			org, err := store.GetOrgVisibilityById(ctx, &org.GetOrgVisibilityByIdQuery{OrgId: o.Id})
			require.NoError(t, err)
			require.NotEmpty(t, org)
			require.Equal(t, org, model.PrivateVisibility)
		}
	})

	t.Run("GetWorkspaceList", func(t *testing.T) {
		orgList, err := getAllOrgList(ctx)
		require.NoError(t, err)
		require.NotZero(t, len(orgList))

		for _, o := range orgList {
			query := &org.GetWorkspaceListQuery{
				OrgId:      o.Id,
				Visibility: model.AllVisibility,
			}
			list, err := store.GetWorkspaceList(ctx, query)
			require.NoError(t, err)
			require.Empty(t, list)
			require.Equal(t, len(list), 0)
		}
	})

	t.Run("GetProjectList", func(t *testing.T) {
		orgList, err := getAllOrgList(ctx)
		require.NoError(t, err)
		require.NotZero(t, len(orgList))

		for _, o := range orgList {
			query := &org.GetProjectListQuery{
				OrgId:      o.Id,
				Visibility: model.AllVisibility,
			}
			list, err := store.GetProjectList(ctx, query)
			require.NoError(t, err)
			require.Empty(t, list)
			require.Equal(t, len(list), 0)
		}
	})

	t.Run("GetUserAssignmentList", func(t *testing.T) {
		orgList, err := getAllOrgList(ctx)
		require.NoError(t, err)
		require.NotZero(t, len(orgList))

		for _, o := range orgList {
			query := &org.GetUserAssignmentListQuery{OrgId: o.Id}
			list, err := store.GetUserAssignmentList(ctx, query)
			require.NoError(t, err)
			require.Empty(t, list)
			require.Equal(t, len(list), 0)
		}
	})

	t.Run("DeleteOrgById", func(t *testing.T) {
		orgList, err := getAllOrgList(ctx)
		require.NoError(t, err)
		require.NotZero(t, len(orgList))

		for _, o := range orgList {
			err := store.DeleteOrgById(ctx, &org.DeleteOrgByIdCommand{OrgId: o.Id})
			require.NoError(t, err)
		}
	})
}
