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
)

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
    %s;
`

func (store *projectStore) GetProjectList(ctx context.Context, visibility model.ScopeVisibility, orgId []int, workspaceId []int) ([]*project.Project, error) {
	whereClause := []string{}
	if len(orgId) != 0 {
		whereClause = append(whereClause, `p."orgId" = ANY (@orgId)`)
	}
	if len(workspaceId) != 0 {
		whereClause = append(whereClause, `p."workspaceId" = ANY (@workspaceId)`)
	}
	if visibility == model.PublicVisibility || visibility == model.PrivateVisibility {
		whereClause = append(whereClause, `p."visibility" = @visibility`)
	}

	finalQuery := fmt.Sprintf(listProjectsQuery, strings.Join(whereClause[:], " AND "))

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

/// GetProjectById

const getProjectByIdQuery = `
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
    p."id" = @id;
`

func (store *projectStore) GetProjectById(ctx context.Context, projectId int) (*project.Project, error) {
	row := store.Conn.QueryRow(ctx, getProjectByIdQuery, pgx.NamedArgs{"id": projectId})

	var project project.Project
	err := row.Scan(
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
		fmt.Fprintf(os.Stderr, "[READ] failed: %v\n", err)
		return nil, err
	}

	return &project, nil
}

/// GetProject

const getProjectQuery = `
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
    %s
ORDER BY "updatedAt" ASC
LIMIT 1;
`

func (store *projectStore) GetProject(ctx context.Context, p *project.Project) (*project.Project, error) {
	availableFields := []string{}
	projectFields := map[string]interface{}{
		"sid":        p.ScopedId,
		"name":       p.Name,
		"visibility": p.Visibility,
		"createdAt":  p.CreatedAt,
		"updatedAt":  p.UpdatedAt,
	}
	for key, value := range projectFields {
		switch v := value.(type) {
		case int:
			if v != 0 {
				availableFields = append(availableFields, fmt.Sprintf("p.\"%s\" = '%d'", key, v))
			}
		case float64:
			if v != 0.0 {
				availableFields = append(availableFields, fmt.Sprintf("p.\"%s\" = '%f'", key, v))
			}
		case string:
			if v != "" {
				availableFields = append(availableFields, fmt.Sprintf("p.\"%s\" = '%s'", key, v))
			}
		case *time.Time:
			if v != nil {
				availableFields = append(availableFields, fmt.Sprintf("p.\"%s\" = '%s'", key, v))
			}
		}
	}
	finalQuery := fmt.Sprintf(getProjectQuery, strings.Join(availableFields, " AND "))

	row := store.Conn.QueryRow(ctx, finalQuery)

	var project project.Project
	err := row.Scan(
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
		fmt.Fprintf(os.Stderr, "[READ] failed: %v\n", err)
		return nil, err
	}

	return &project, nil
}
