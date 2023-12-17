package store

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/quarkloop/quarkloop/pkg/service/project"
)

/// ListProjects

const listProjectsQuery = `
SELECT 
  "id", "sid", "orgId", "workspaceId",
  "name", "description", "accessType",
  "createdAt", "createdBy", "updatedAt", "updatedBy"
FROM 
  "system"."Project"
WHERE
  %s;
`

func (store *projectStore) ListProjects(ctx context.Context, orgId []int, workspaceId []int) ([]project.Project, error) {
	var whereClause string = `"accessType" = 1` // TODO: 1 => public, 2 => private
	if len(orgId) != 0 {
		whereClause = `"orgId" = ANY (@orgId)`
	}
	if len(workspaceId) != 0 {
		whereClause = `"workspaceId" = ANY (@workspaceId)`
	}
	finalQuery := fmt.Sprintf(listProjectsQuery, whereClause)

	rows, err := store.Conn.Query(ctx, finalQuery, pgx.NamedArgs{
		"orgId":       orgId,
		"workspaceId": workspaceId,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var projectList []project.Project = []project.Project{}

	for rows.Next() {
		var project project.Project
		err := rows.Scan(
			&project.Id,
			&project.ScopedId,
			&project.OrgId,
			&project.WorkspaceId,
			&project.Name,
			&project.Description,
			&project.AccessType,
			&project.CreatedAt,
			&project.CreatedBy,
			&project.UpdatedAt,
			&project.UpdatedBy,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
			return nil, err
		}

		projectList = append(projectList, project)
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
  "id", "sid", "orgId", "workspaceId",
  "name", "description", "accessType",
  "createdAt", "createdBy", "updatedAt", "updatedBy"
FROM
  "system"."Project"
WHERE
  "id" = @id;
`

func (store *projectStore) GetProjectById(ctx context.Context, projectId int) (*project.Project, error) {
	row := store.Conn.QueryRow(ctx, getProjectByIdQuery, pgx.NamedArgs{"id": projectId})

	var project project.Project
	err := row.Scan(
		&project.Id,
		&project.ScopedId,
		&project.OrgId,
		&project.WorkspaceId,
		&project.Name,
		&project.Description,
		&project.AccessType,
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
  "id", "sid", "orgId", "workspaceId",
  "name", "description", "accessType",
  "createdAt", "createdBy", "updatedAt", "updatedBy"
FROM
  "system"."Project"
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
		"accessType": p.AccessType,
		"createdAt":  p.CreatedAt,
		"updatedAt":  p.UpdatedAt,
	}
	for key, value := range projectFields {
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
	finalQuery := fmt.Sprintf(getProjectQuery, strings.Join(availableFields, " AND "))

	row := store.Conn.QueryRow(ctx, finalQuery)

	var project project.Project
	err := row.Scan(
		&project.Id,
		&project.ScopedId,
		&project.OrgId,
		&project.WorkspaceId,
		&project.Name,
		&project.Description,
		&project.AccessType,
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
