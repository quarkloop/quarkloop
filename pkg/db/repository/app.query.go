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

/// ListApps

type ListAppsParams struct {
	Context     context.Context
	OsId        []string
	WorkspaceId []string
}

const listAppsQuery = `
SELECT 
  "id", "name", "path", "description", "createdAt", "updatedAt"
FROM 
  "system"."App"
WHERE
  %s;
`

func (r *Repository) ListApps(p *ListAppsParams) ([]model.App, error) {
	var whereClause string = `"accessType" = 1` // TODO: 1 => public, 2 => private
	if len(p.OsId) != 0 {
		whereClause = `"osId" = ANY (@osId)`
	}
	if len(p.WorkspaceId) != 0 {
		whereClause = `"workspaceId" = ANY (@workspaceId)`
	}
	finalQuery := fmt.Sprintf(listAppsQuery, whereClause)

	rows, err := r.SystemDbConn.Query(p.Context, finalQuery, pgx.NamedArgs{
		"osId":        p.OsId,
		"workspaceId": p.WorkspaceId,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var appList []model.App

	for rows.Next() {
		var app model.App
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

/// FindUniqueApp

type FindUniqueAppParams struct {
	Context context.Context
	AppId   string
}

const findUniqueAppQuery = `
SELECT
  "id", "name", "path", "description", "createdAt", "updatedAt"
FROM
  "system"."App"
WHERE
  "id" = @id;
`

func (r *Repository) FindUniqueApp(p *FindUniqueAppParams) (*model.App, error) {
	row := r.SystemDbConn.QueryRow(p.Context, findUniqueAppQuery, pgx.NamedArgs{"id": p.AppId})

	var app model.App
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

/// FindFirstApp

type FindFirstAppParams struct {
	Context context.Context
	App     model.App
}

const findFirstAppQuery = `
SELECT
  "id", "name", "path", "description", "createdAt", "updatedAt"
FROM
  "system"."App"
WHERE 
  %s
ORDER BY "updatedAt" ASC 
LIMIT 1;
`

func (r *Repository) FindFirstApp(p *FindFirstAppParams) (*model.App, error) {
	availableFields := []string{}
	appFields := map[string]interface{}{
		"id":        p.App.Id,
		"name":      p.App.Name,
		"createdAt": p.App.CreatedAt,
		"updatedAt": p.App.UpdatedAt,
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
	finalQuery := fmt.Sprintf(findFirstAppQuery, strings.Join(availableFields, " AND "))

	row := r.SystemDbConn.QueryRow(p.Context, finalQuery)

	var app model.App
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
