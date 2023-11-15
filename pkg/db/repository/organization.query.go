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

/// ListOrganizations

type ListOrganizationsParams struct {
	Context context.Context
}

const listOrganizationsQuery = `
SELECT 
  "id", "name", "description", "path", "createdAt", "updatedAt"
FROM 
  "system"."Organization";
`

func (r *Repository) ListOrganizations(p *ListOrganizationsParams) ([]model.Organization, error) {
	rows, err := r.SystemDbConn.Query(p.Context, listOrganizationsQuery)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var osList []model.Organization

	for rows.Next() {
		var operatingSystem model.Organization
		err := rows.Scan(
			&operatingSystem.Id,
			&operatingSystem.Name,
			&operatingSystem.Description,
			&operatingSystem.Path,
			&operatingSystem.CreatedAt,
			&operatingSystem.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		osList = append(osList, operatingSystem)
	}

	if err := rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
		return nil, err
	}

	return osList, nil
}

/// FindUniqueOrganization

type FindUniqueOrganizationParams struct {
	Context context.Context
	Id      string
}

const findUniqueOrganizationQuery = `
SELECT 
  "id", "name", "description", "path", "createdAt", "updatedAt"
FROM 
  "system"."Organization" 
WHERE 
  "id" = @id;
`

func (r *Repository) FindUniqueOrganization(p *FindUniqueOrganizationParams) (*model.Organization, error) {
	row := r.SystemDbConn.QueryRow(p.Context, findUniqueOrganizationQuery, pgx.NamedArgs{"id": p.Id})

	var operatingSystem model.Organization
	err := row.Scan(
		&operatingSystem.Id,
		&operatingSystem.Name,
		&operatingSystem.Description,
		&operatingSystem.Path,
		&operatingSystem.CreatedAt,
		&operatingSystem.UpdatedAt,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[READ] failed: %v\n", err)
		return nil, err
	}

	return &operatingSystem, nil
}

/// FindFirstOrganization

type FindFirstOrganizationParams struct {
	Context      context.Context
	Organization model.Organization
}

const findFirstOrganizationQuery = `
SELECT 
  "id", "name", "description", "path", "createdAt", "updatedAt"
FROM 
  "system"."Organization" 
WHERE
`

func (r *Repository) FindFirstOrganization(p *FindFirstOrganizationParams) (*model.Organization, error) {
	availableFields := []string{}
	operatingSystemFields := map[string]interface{}{
		"id":        p.Organization.Id,
		"name":      p.Organization.Name,
		"path":      p.Organization.Path,
		"createdAt": p.Organization.CreatedAt,
		"updatedAt": p.Organization.UpdatedAt,
	}
	for key, value := range operatingSystemFields {
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
	finalQuery := findFirstOrganizationQuery + strings.Join(availableFields, " AND ")

	row := r.SystemDbConn.QueryRow(p.Context, finalQuery)

	var operatingSystem model.Organization
	err := row.Scan(
		&operatingSystem.Id,
		&operatingSystem.Name,
		&operatingSystem.Description,
		&operatingSystem.Path,
		&operatingSystem.CreatedAt,
		&operatingSystem.UpdatedAt,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[READ] failed: %v\n", err)
		return nil, err
	}

	return &operatingSystem, nil
}
