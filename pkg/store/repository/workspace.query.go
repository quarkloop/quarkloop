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

type ListWorkspacesParams struct {
	Context context.Context
	OrgId   []string
}

const listWorkspacesQuery = `
SELECT 
  "id", "name", "accessType", "description", "path", "createdAt", "updatedAt"
FROM 
  "system"."Workspace"
WHERE
  "orgId" = ANY (@orgId);
`

func (r *Repository) ListWorkspaces(p *ListWorkspacesParams) ([]model.Workspace, error) {
	rows, err := r.SystemDbConn.Query(p.Context, listWorkspacesQuery, pgx.NamedArgs{"orgId": p.OrgId})
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

/// FindUniqueWorkspace

type FindUniqueWorkspaceParams struct {
	Context     context.Context
	WorkspaceId string
}

const findUniqueWorkspaceQuery = `
SELECT 
  "id", "name", "accessType", "description", "path", "createdAt", "updatedAt"
FROM 
  "system"."Workspace" 
WHERE 
  "id" = @id;
`

func (r *Repository) FindUniqueWorkspace(p *FindUniqueWorkspaceParams) (*model.Workspace, error) {
	row := r.SystemDbConn.QueryRow(p.Context, findUniqueWorkspaceQuery, pgx.NamedArgs{"id": p.WorkspaceId})

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

type FindFirstWorkspaceParams struct {
	Context   context.Context
	OrgId     string
	Workspace model.Workspace
}

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

func (r *Repository) FindFirstWorkspace(p *FindFirstWorkspaceParams) (*model.Workspace, error) {
	availableFields := []string{}
	workspaceFields := map[string]interface{}{
		"orgId":      p.OrgId,
		"id":         p.Workspace.Id,
		"name":       p.Workspace.Name,
		"accessType": *p.Workspace.AccessType,
		"path":       p.Workspace.Path,
		"createdAt":  p.Workspace.CreatedAt,
		"updatedAt":  p.Workspace.UpdatedAt,
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

	row := r.SystemDbConn.QueryRow(p.Context, finalQuery)

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
