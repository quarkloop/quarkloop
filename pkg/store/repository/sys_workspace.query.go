package repository

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/quarkloop/quarkloop/pkg/model"
)

/// ListWorkspaces

const listWorkspacesQuery = `
SELECT 
  "id", "name", "description", "accessType",  "createdAt", "createdBy", "updatedAt", "updatedBy"
FROM 
  "system"."Workspace"
WHERE
  "orgId" = ANY (@orgId);
`

func (r *Repository) ListWorkspaces(ctx context.Context, orgId []int) ([]model.Workspace, error) {
	rows, err := r.SystemDbConn.Query(ctx, listWorkspacesQuery, pgx.NamedArgs{"orgId": orgId})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var wsList []model.Workspace = []model.Workspace{}

	for rows.Next() {
		var workspace model.Workspace
		err := rows.Scan(
			&workspace.Id,
			&workspace.Name,
			&workspace.Description,
			&workspace.AccessType,
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
  "id", "name", "description", "accessType",  "createdAt", "createdBy", "updatedAt", "updatedBy"
FROM 
  "system"."Workspace" 
WHERE 
  "id" = @id;
`

func (r *Repository) GetWorkspaceById(ctx context.Context, workspaceId int) (*model.Workspace, error) {
	row := r.SystemDbConn.QueryRow(ctx, getWorkspaceByIdQuery, pgx.NamedArgs{"id": workspaceId})

	var workspace model.Workspace
	err := row.Scan(
		&workspace.Id,
		&workspace.Name,
		&workspace.Description,
		&workspace.AccessType,
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
  "id", "name", "description", "accessType",  "createdAt", "createdBy", "updatedAt", "updatedBy"
FROM 
  "system"."Workspace" 
WHERE
%s
ORDER BY "updatedAt" ASC
LIMIT 1;
`

func (r *Repository) GetWorkspace(ctx context.Context, orgId int, workspace *model.Workspace) (*model.Workspace, error) {
	availableFields := []string{}
	workspaceFields := map[string]interface{}{
		"orgId":      orgId,
		"id":         workspace.Id,
		"name":       workspace.Name,
		"accessType": *workspace.AccessType,
		"createdAt":  workspace.CreatedAt,
		"updatedAt":  workspace.UpdatedAt,
	}
	for key, value := range workspaceFields {
		switch v := value.(type) {
		case int:
			if v != 0 {
				availableFields = append(availableFields, fmt.Sprintf("\"%s\" = '%d'", key, v))
			}
		case float64:
			if v != 0.0 {
				availableFields = append(availableFields, fmt.Sprintf("\"%s\" = '%f'", key, v))
			}
		case string:
			if v != "" {
				availableFields = append(availableFields, fmt.Sprintf("\"%s\" = '%s'", key, v))
			}
		case *time.Time:
			if v != nil {
				availableFields = append(availableFields, fmt.Sprintf("\"%s\" = '%s'", key, v))
			}
		}
	}
	finalQuery := fmt.Sprintf(getWorkspaceQuery, strings.Join(availableFields, " AND "))

	row := r.SystemDbConn.QueryRow(ctx, finalQuery)

	var ws model.Workspace
	err := row.Scan(
		&ws.Id,
		&ws.Name,
		&ws.Description,
		&ws.AccessType,
		&ws.CreatedAt,
		&ws.CreatedBy,
		&ws.UpdatedAt,
		&ws.UpdatedBy,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[READ] failed: %v\n", err)
		return nil, err
	}

	return &ws, nil
}
