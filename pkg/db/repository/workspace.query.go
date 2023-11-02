package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"

	"github.com/quarkloop/quarkloop/pkg/db/model"
)

/// ListWorkspaces

type ListWorkspacesParams struct {
	Context context.Context
	OsId    string
}

const listWorkspacesQuery = `
SELECT 
  "id", "name", "description", "path", "createdAt", "updatedAt"
FROM 
  "app"."Workspace"
WHERE
  "osId" = @osId;
`

func (r *Repository) ListWorkspaces(p *ListWorkspacesParams) ([]model.Workspace, error) {
	rows, err := r.Conn.Query(p.Context, listWorkspacesQuery, pgx.NamedArgs{"osId": p.OsId})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var wsList []model.Workspace

	for rows.Next() {
		var workspace model.Workspace
		err := rows.Scan(
			&workspace.Id,
			&workspace.Name,
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

type GetWorkspaceByIdParams struct {
	Context context.Context
	Id      string
}

const getWorkspaceByIdQuery = `
SELECT 
  "id", "name", "description", "path", "createdAt", "updatedAt"
FROM 
  "app"."Workspace" 
WHERE 
  "id" = @id;
`

func (r *Repository) GetWorkspaceById(p *GetWorkspaceByIdParams) (*model.Workspace, error) {
	row := r.Conn.QueryRow(p.Context, getWorkspaceByIdQuery, pgx.NamedArgs{"id": p.Id})

	var workspace model.Workspace
	err := row.Scan(
		&workspace.Id,
		&workspace.Name,
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
