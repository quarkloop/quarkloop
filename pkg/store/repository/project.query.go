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

/// ListProjects

const listProjectsQuery = `
SELECT 
  "id", "name", "accessType", "path", "description", "createdAt", "updatedAt"
FROM 
  "system"."Project"
WHERE
  %s;
`

func (r *Repository) ListProjects(ctx context.Context, orgId []string, workspaceId []string) ([]model.Project, error) {
	var whereClause string = `"accessType" = 1` // TODO: 1 => public, 2 => private
	if len(orgId) != 0 {
		whereClause = `"orgId" = ANY (@orgId)`
	}
	if len(workspaceId) != 0 {
		whereClause = `"workspaceId" = ANY (@workspaceId)`
	}
	finalQuery := fmt.Sprintf(listProjectsQuery, whereClause)

	rows, err := r.SystemDbConn.Query(ctx, finalQuery, pgx.NamedArgs{
		"orgId":       orgId,
		"workspaceId": workspaceId,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var projectList []model.Project = []model.Project{}

	for rows.Next() {
		var project model.Project
		err := rows.Scan(
			&project.Id,
			&project.Name,
			&project.AccessType,
			&project.Path,
			&project.Description,
			&project.CreatedAt,
			&project.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		projectList = append(projectList, project)
	}

	if err := rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "[LIST]:Rows failed: %v\n", err)
		return nil, err
	}

	return projectList, nil
}

/// FindUniqueProject

const findUniqueProjectQuery = `
SELECT
  "id", "name", "accessType", "path", "description", "createdAt", "updatedAt"
FROM
  "system"."Project"
WHERE
  "id" = @id;
`

func (r *Repository) FindUniqueProject(ctx context.Context, projectId string) (*model.Project, error) {
	row := r.SystemDbConn.QueryRow(ctx, findUniqueProjectQuery, pgx.NamedArgs{"id": projectId})

	var project model.Project
	err := row.Scan(
		&project.Id,
		&project.Name,
		&project.AccessType,
		&project.Path,
		&project.Description,
		&project.CreatedAt,
		&project.UpdatedAt,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[READ] failed: %v\n", err)
		return nil, err
	}

	return &project, nil
}

/// FindFirstProject

const findFirstProjectQuery = `
SELECT
  "id", "name", "accessType", "path", "description", "createdAt", "updatedAt"
FROM
  "system"."Project"
WHERE 
  %s
ORDER BY "updatedAt" ASC 
LIMIT 1;
`

func (r *Repository) FindFirstProject(ctx context.Context, project *model.Project) (*model.Project, error) {
	availableFields := []string{}
	projectFields := map[string]interface{}{
		"id":         project.Id,
		"name":       project.Name,
		"accessType": project.AccessType,
		"createdAt":  project.CreatedAt,
		"updatedAt":  project.UpdatedAt,
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
	finalQuery := fmt.Sprintf(findFirstProjectQuery, strings.Join(availableFields, " AND "))

	row := r.SystemDbConn.QueryRow(ctx, finalQuery)

	var p model.Project
	err := row.Scan(
		&p.Id,
		&p.Name,
		&p.AccessType,
		&p.Path,
		&p.Description,
		&p.CreatedAt,
		&p.UpdatedAt,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[READ] failed: %v\n", err)
		return nil, err
	}

	return &p, nil
}
