package store

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/quarkloop/quarkloop/pkg/service/workspace"
)

/// ListWorkspaces

const listWorkspacesQuery = `
SELECT 
	ws."id",
    ws."sid",
    ws."orgId",
    org."sid",
    ws."name",
    ws."description",
    ws."visibility",
    ws."createdAt",
    ws."createdBy",
    ws."updatedAt",
    ws."updatedBy"
FROM 
	"system"."Workspace" AS ws
LEFT JOIN 
	system."Organization" AS org ON org."id" = ws."orgId"
WHERE 
	ws."orgId" = ANY (@orgId);
`

func (store *workspaceStore) ListWorkspaces(ctx context.Context, orgId []int) ([]workspace.Workspace, error) {
	rows, err := store.Conn.Query(ctx, listWorkspacesQuery, pgx.NamedArgs{"orgId": orgId})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var wsList []workspace.Workspace = []workspace.Workspace{}

	for rows.Next() {
		var workspace workspace.Workspace
		err := rows.Scan(
			&workspace.Id,
			&workspace.ScopedId,
			&workspace.OrgId,
			&workspace.OrgScopedId,
			&workspace.Name,
			&workspace.Description,
			&workspace.Visibility,
			&workspace.CreatedAt,
			&workspace.CreatedBy,
			&workspace.UpdatedAt,
			&workspace.UpdatedBy,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
			return nil, err
		}

		wsList = append(wsList, workspace)
	}

	if err := rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
		return nil, err
	}

	return wsList, nil
}

/// GetWorkspaceById

const getWorkspaceByIdQuery = `
SELECT 
	ws."id",
    ws."sid",
    ws."orgId",
    org."sid",
    ws."name",
    ws."description",
    ws."visibility",
    ws."createdAt",
    ws."createdBy",
    ws."updatedAt",
    ws."updatedBy"
FROM 
	"system"."Workspace" AS ws
LEFT JOIN 
	system."Organization" AS org ON org."id" = ws."orgId"
WHERE 
	ws."id" = @id;
`

func (store *workspaceStore) GetWorkspaceById(ctx context.Context, workspaceId int) (*workspace.Workspace, error) {
	row := store.Conn.QueryRow(ctx, getWorkspaceByIdQuery, pgx.NamedArgs{"id": workspaceId})

	var workspace workspace.Workspace
	err := row.Scan(
		&workspace.Id,
		&workspace.ScopedId,
		&workspace.OrgId,
		&workspace.OrgScopedId,
		&workspace.Name,
		&workspace.Description,
		&workspace.Visibility,
		&workspace.CreatedAt,
		&workspace.CreatedBy,
		&workspace.UpdatedAt,
		&workspace.UpdatedBy,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[READ] failed: %v\n", err)
		return nil, err
	}

	return &workspace, nil
}

/// GetWorkspace

const getWorkspaceQuery = `
SELECT 
	ws."id",
    ws."sid",
    ws."orgId",
    org."sid",
    ws."name",
    ws."description",
    ws."visibility",
    ws."createdAt",
    ws."createdBy",
    ws."updatedAt",
    ws."updatedBy"
FROM 
	"system"."Workspace" AS ws
LEFT JOIN 
	system."Organization" AS org ON org."id" = ws."orgId"
WHERE 
	%s
ORDER BY 
	"updatedAt" ASC
LIMIT 1;
`

func (store *workspaceStore) GetWorkspace(ctx context.Context, orgId int, ws *workspace.Workspace) (*workspace.Workspace, error) {
	availableFields := []string{}
	workspaceFields := map[string]interface{}{
		"name":       ws.Name,
		"visibility": *ws.Visibility,
		"createdAt":  ws.CreatedAt,
		"updatedAt":  ws.UpdatedAt,
	}
	for key, value := range workspaceFields {
		switch v := value.(type) {
		case int:
			if v != 0 {
				availableFields = append(availableFields, fmt.Sprintf("ws.\"%s\" = '%d'", key, v))
			}
		case float64:
			if v != 0.0 {
				availableFields = append(availableFields, fmt.Sprintf("ws.\"%s\" = '%f'", key, v))
			}
		case string:
			if v != "" {
				availableFields = append(availableFields, fmt.Sprintf("ws.\"%s\" = '%s'", key, v))
			}
		case *time.Time:
			if v != nil {
				availableFields = append(availableFields, fmt.Sprintf("ws.\"%s\" = '%s'", key, v))
			}
		}
	}
	finalQuery := fmt.Sprintf(getWorkspaceQuery, strings.Join(availableFields, " AND "))

	row := store.Conn.QueryRow(ctx, finalQuery)

	var workspace workspace.Workspace
	err := row.Scan(
		&workspace.Id,
		&workspace.ScopedId,
		&workspace.OrgId,
		&workspace.OrgScopedId,
		&workspace.Name,
		&workspace.Description,
		&workspace.Visibility,
		&workspace.CreatedAt,
		&workspace.CreatedBy,
		&workspace.UpdatedAt,
		&workspace.UpdatedBy,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[READ] failed: %v\n", err)
		return nil, err
	}

	return &workspace, nil
}
