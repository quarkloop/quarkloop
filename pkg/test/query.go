package test

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/quarkloop/quarkloop/pkg/service/org"
	"github.com/quarkloop/quarkloop/pkg/service/workspace"
)

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
ORDER BY id ASC;
`

func GetFullOrgList(ctx context.Context, conn *pgx.Conn) ([]*org.Org, error) {
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

const getWorkspaceListQuery = `
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
    "system"."Workspace"
ORDER BY id ASC;
`

func GetFullWorkspaceList(ctx context.Context, conn *pgx.Conn) ([]*workspace.Workspace, error) {
	rows, err := conn.Query(ctx, getWorkspaceListQuery)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var wsList []*workspace.Workspace = []*workspace.Workspace{}
	for rows.Next() {
		var ws workspace.Workspace
		err := rows.Scan(
			&ws.Id,
			&ws.ScopeId,
			&ws.Name,
			&ws.Description,
			&ws.Visibility,
			&ws.CreatedAt,
			&ws.CreatedBy,
			&ws.UpdatedAt,
			&ws.UpdatedBy,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
			return nil, err
		}
		wsList = append(wsList, &ws)
	}
	if err := rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
		return nil, err
	}

	return wsList, nil
}
