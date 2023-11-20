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

/// ListOrganizations

type ListOrganizationsParams struct {
	Context context.Context
}

const listOrganizationsQuery = `
SELECT 
  "id", "name", "accessType", "description", "path", "createdAt", "updatedAt"
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

	var orgList []model.Organization = []model.Organization{}

	for rows.Next() {
		var organization model.Organization
		err := rows.Scan(
			&organization.Id,
			&organization.Name,
			&organization.AccessType,
			&organization.Description,
			&organization.Path,
			&organization.CreatedAt,
			&organization.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		orgList = append(orgList, organization)
	}

	if err := rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
		return nil, err
	}

	return orgList, nil
}

/// FindUniqueOrganization

type FindUniqueOrganizationParams struct {
	Context context.Context
	Id      string
}

const findUniqueOrganizationQuery = `
SELECT 
  "id", "name", "accessType", "description", "path", "createdAt", "updatedAt"
FROM 
  "system"."Organization" 
WHERE 
  "id" = @id;
`

func (r *Repository) FindUniqueOrganization(p *FindUniqueOrganizationParams) (*model.Organization, error) {
	row := r.SystemDbConn.QueryRow(p.Context, findUniqueOrganizationQuery, pgx.NamedArgs{"id": p.Id})

	var organization model.Organization
	err := row.Scan(
		&organization.Id,
		&organization.Name,
		&organization.AccessType,
		&organization.Description,
		&organization.Path,
		&organization.CreatedAt,
		&organization.UpdatedAt,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[READ] failed: %v\n", err)
		return nil, err
	}

	return &organization, nil
}

/// FindFirstOrganization

type FindFirstOrganizationParams struct {
	Context      context.Context
	Organization model.Organization
}

const findFirstOrganizationQuery = `
SELECT 
  "id", "name", "accessType", "description", "path", "createdAt", "updatedAt"
FROM 
  "system"."Organization" 
WHERE
`

func (r *Repository) FindFirstOrganization(p *FindFirstOrganizationParams) (*model.Organization, error) {
	availableFields := []string{}
	organizationFields := map[string]interface{}{
		"id":         p.Organization.Id,
		"name":       p.Organization.Name,
		"accessType": p.Organization.AccessType,
		"path":       p.Organization.Path,
		"createdAt":  p.Organization.CreatedAt,
		"updatedAt":  p.Organization.UpdatedAt,
	}
	for key, value := range organizationFields {
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

	var organization model.Organization
	err := row.Scan(
		&organization.Id,
		&organization.Name,
		&organization.AccessType,
		&organization.Description,
		&organization.Path,
		&organization.CreatedAt,
		&organization.UpdatedAt,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[READ] failed: %v\n", err)
		return nil, err
	}

	return &organization, nil
}
