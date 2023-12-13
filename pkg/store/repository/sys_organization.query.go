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

const listOrganizationsQuery = `
SELECT 
  "id", "name", "accessType", "description", "path", "createdAt", "updatedAt"
FROM 
  "system"."Organization";
`

func (r *Repository) ListOrganizations(ctx context.Context) ([]model.Organization, error) {
	rows, err := r.SystemDbConn.Query(ctx, listOrganizationsQuery)
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

/// GetOrganizationById

const getOrganizationByIdQuery = `
SELECT 
  "id", "name", "accessType", "description", "path", "createdAt", "updatedAt"
FROM 
  "system"."Organization" 
WHERE 
  "id" = @id;
`

func (r *Repository) GetOrganizationById(ctx context.Context, orgId string) (*model.Organization, error) {
	row := r.SystemDbConn.QueryRow(ctx, getOrganizationByIdQuery, pgx.NamedArgs{"id": orgId})

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

const findFirstOrganizationQuery = `
SELECT 
  "id", "name", "accessType", "description", "path", "createdAt", "updatedAt"
FROM 
  "system"."Organization" 
WHERE
`

func (r *Repository) FindFirstOrganization(ctx context.Context, org *model.Organization) (*model.Organization, error) {
	availableFields := []string{}
	organizationFields := map[string]interface{}{
		"id":         org.Id,
		"name":       org.Name,
		"accessType": org.AccessType,
		"path":       org.Path,
		"createdAt":  org.CreatedAt,
		"updatedAt":  org.UpdatedAt,
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

	row := r.SystemDbConn.QueryRow(ctx, finalQuery)

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
