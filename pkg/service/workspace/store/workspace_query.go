package store

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/project"
	"github.com/quarkloop/quarkloop/pkg/service/workspace"
)

/// GetWorkspaceList

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
    ws."orgId" = ANY (@orgId)
%s	
`

func (store *workspaceStore) GetWorkspaceList(ctx context.Context, visibility model.ScopeVisibility, orgId []int) ([]*workspace.Workspace, error) {
	var finalQuery strings.Builder
	if visibility == model.PublicVisibility || visibility == model.PrivateVisibility {
		finalQuery.WriteString(fmt.Sprintf(listWorkspacesQuery, `AND ws."visibility" = @visibility;`))
	} else {
		finalQuery.WriteString(fmt.Sprintf(listWorkspacesQuery, ";"))
	}

	rows, err := store.Conn.Query(ctx, finalQuery.String(), pgx.NamedArgs{
		"orgId":      orgId,
		"visibility": visibility,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var wsList []*workspace.Workspace = []*workspace.Workspace{}

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

		wsList = append(wsList, &workspace)
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

/// GetProjectList

const listProjectsQuery = `
SELECT 
    p."id",
    p."sid",
    p."orgId",
    p."workspaceId",
    org."sid",
    ws."sid",
    p."name",
    p."description",
    p."visibility",
    p."createdAt",
    p."createdBy",
    p."updatedAt",
    p."updatedBy"
FROM 
    "system"."Project" AS p
LEFT JOIN 
    system."Organization" AS org ON org."id" = p."orgId"
LEFT JOIN 
    system."Workspace" AS ws ON ws."id" = p."workspaceId"
WHERE
    p."orgId" = @orgId
AND
    p."workspaceId" = @workspaceId	
%s
`

func (store *workspaceStore) GetProjectList(ctx context.Context, visibility model.ScopeVisibility, orgId int, workspaceId int) ([]*project.Project, error) {
	whereClause := ";"
	if visibility == model.PublicVisibility || visibility == model.PrivateVisibility {
		whereClause = `AND p."visibility" = @visibility;`
	}

	finalQuery := fmt.Sprintf(listProjectsQuery, whereClause)

	rows, err := store.Conn.Query(ctx, finalQuery, pgx.NamedArgs{
		"orgId":       orgId,
		"workspaceId": workspaceId,
		"visibility":  visibility,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var projectList []*project.Project = []*project.Project{}

	for rows.Next() {
		var project project.Project
		err := rows.Scan(
			&project.Id,
			&project.ScopedId,
			&project.OrgId,
			&project.WorkspaceId,
			&project.OrgScopedId,
			&project.WorkspaceScopedId,
			&project.Name,
			&project.Description,
			&project.Visibility,
			&project.CreatedAt,
			&project.CreatedBy,
			&project.UpdatedAt,
			&project.UpdatedBy,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
			return nil, err
		}

		projectList = append(projectList, &project)
	}

	if err := rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
		return nil, err
	}

	return projectList, nil
}
