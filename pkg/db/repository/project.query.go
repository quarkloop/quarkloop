package repository

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/quarkloop/quarkloop/pkg/db/model"
)

/// ListProjects

type ListProjectsParams struct {
	Context     context.Context
	OrgId       []string
	WorkspaceId []string
}

const listProjectsQuery = `
SELECT 
  "id", "name", "path", "description", "createdAt", "updatedAt"
FROM 
  "system"."Project"
WHERE
  %s;
`

func (r *Repository) ListProjects(p *ListProjectsParams) ([]model.Project, error) {
	var whereClause string = `"accessType" = 1` // TODO: 1 => public, 2 => private
	if len(p.OrgId) != 0 {
		whereClause = `"orgId" = ANY (@orgId)`
	}
	if len(p.WorkspaceId) != 0 {
		whereClause = `"workspaceId" = ANY (@workspaceId)`
	}
	finalQuery := fmt.Sprintf(listProjectsQuery, whereClause)

	rows, err := r.SystemDbConn.Query(p.Context, finalQuery, pgx.NamedArgs{
		"orgId":       p.OrgId,
		"workspaceId": p.WorkspaceId,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var appList []model.Project

	for rows.Next() {
		var app model.Project
		err := rows.Scan(
			&app.Id,
			&app.Name,
			&app.Path,
			&app.Description,
			&app.CreatedAt,
			&app.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		appList = append(appList, app)
	}

	if err := rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "[LIST]:Rows failed: %v\n", err)
		return nil, err
	}

	return appList, nil
}

/// FindUniqueProject

type FindUniqueProjectParams struct {
	Context   context.Context
	ProjectId string
}

const findUniqueProjectQuery = `
SELECT
  "id", "name", "path", "description", "createdAt", "updatedAt"
FROM
  "system"."Project"
WHERE
  "id" = @id;
`

func (r *Repository) FindUniqueProject(p *FindUniqueProjectParams) (*model.Project, error) {
	row := r.SystemDbConn.QueryRow(p.Context, findUniqueProjectQuery, pgx.NamedArgs{"id": p.ProjectId})

	var app model.Project
	err := row.Scan(
		&app.Id,
		&app.Name,
		&app.Path,
		&app.Description,
		&app.CreatedAt,
		&app.UpdatedAt,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[READ] failed: %v\n", err)
		return nil, err
	}

	return &app, nil
}

/// FindFirstProject

type FindFirstProjectParams struct {
	Context context.Context
	Project model.Project
}

const findFirstProjectQuery = `
SELECT
  "id", "name", "path", "description", "createdAt", "updatedAt"
FROM
  "system"."Project"
WHERE 
  %s
ORDER BY "updatedAt" ASC 
LIMIT 1;
`

func (r *Repository) FindFirstProject(p *FindFirstProjectParams) (*model.Project, error) {
	availableFields := []string{}
	appFields := map[string]interface{}{
		"id":        p.Project.Id,
		"name":      p.Project.Name,
		"createdAt": p.Project.CreatedAt,
		"updatedAt": p.Project.UpdatedAt,
	}
	for key, value := range appFields {
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

	row := r.SystemDbConn.QueryRow(p.Context, finalQuery)

	var app model.Project
	err := row.Scan(
		&app.Id,
		&app.Name,
		&app.Path,
		&app.Description,
		&app.CreatedAt,
		&app.UpdatedAt,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[READ] failed: %v\n", err)
		return nil, err
	}

	return &app, nil
}
