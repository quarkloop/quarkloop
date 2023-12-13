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
  "id", "name", "accessType", "description", "path", "createdAt", "updatedAt"
FROM 
  "system"."Workspace"
WHERE
  "orgId" = ANY (@orgId);
`

func (r *Repository) ListWorkspaces(ctx context.Context, orgId []string) ([]model.Workspace, error) {
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
			&workspace.AccessType,
			&workspace.Description,
			&workspace.Path,
			&workspace.CreatedAt,
			&workspace.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		wsList = append(wsList, workspace)
	}

	if err := rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
		return nil, err
	}

	return wsList, nil
}

/// GetWorkspaceById

const getWorkspaceByIdQuery = `
SELECT 
  "id", "name", "accessType", "description", "path", "createdAt", "updatedAt"
FROM 
  "system"."Workspace" 
WHERE 
  "id" = @id;
`

func (r *Repository) GetWorkspaceById(ctx context.Context, workspaceId string) (*model.Workspace, error) {
	row := r.SystemDbConn.QueryRow(ctx, getWorkspaceByIdQuery, pgx.NamedArgs{"id": workspaceId})

	var workspace model.Workspace
	err := row.Scan(
		&workspace.Id,
		&workspace.Name,
		&workspace.AccessType,
		&workspace.Description,
		&workspace.Path,
		&workspace.CreatedAt,
		&workspace.UpdatedAt,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[READ] failed: %v\n", err)
		return nil, err
	}

	return &workspace, nil
}

/// FindFirstWorkspace

const findFirstWorkspaceQuery = `
SELECT 
  "id", "name", "accessType", "description", "path", "createdAt", "updatedAt"
FROM 
  "system"."Workspace" 
WHERE
%s
ORDER BY "updatedAt" ASC
LIMIT 1;
`

func (r *Repository) FindFirstWorkspace(ctx context.Context, orgId string, workspace *model.Workspace) (*model.Workspace, error) {
	availableFields := []string{}
	workspaceFields := map[string]interface{}{
		"orgId":      orgId,
		"id":         workspace.Id,
		"name":       workspace.Name,
		"accessType": *workspace.AccessType,
		"path":       workspace.Path,
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
	finalQuery := fmt.Sprintf(findFirstWorkspaceQuery, strings.Join(availableFields, " AND "))

	row := r.SystemDbConn.QueryRow(ctx, finalQuery)

	var ws model.Workspace
	err := row.Scan(
		&ws.Id,
		&ws.Name,
		&ws.AccessType,
		&ws.Description,
		&ws.Path,
		&ws.CreatedAt,
		&ws.UpdatedAt,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[READ] failed: %v\n", err)
		return nil, err
	}

	return &ws, nil
}
