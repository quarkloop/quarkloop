package store

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"

	org "github.com/quarkloop/quarkloop/pkg/service/organization"
)

/// ListOrganizations

const listOrganizationsQuery = `
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
	"system"."Organization";
`

func (store *orgStore) ListOrganizations(ctx context.Context) ([]org.Organization, error) {
	rows, err := store.Conn.Query(ctx, listOrganizationsQuery)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var orgList []org.Organization = []org.Organization{}

	for rows.Next() {
		var org org.Organization
		err := rows.Scan(
			&org.Id,
			&org.ScopedId,
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

		orgList = append(orgList, org)
	}

	if err := rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
		return nil, err
	}

	return orgList, nil
}

/// GetOrganizationById

const getOrganizationByIdQuery = `
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
WHERE 
	"id" = @id;
`

func (store *orgStore) GetOrganizationById(ctx context.Context, orgId int) (*org.Organization, error) {
	row := store.Conn.QueryRow(ctx, getOrganizationByIdQuery, pgx.NamedArgs{"id": orgId})

	var org org.Organization
	err := row.Scan(
		&org.Id,
		&org.ScopedId,
		&org.Name,
		&org.Description,
		&org.Visibility,
		&org.CreatedAt,
		&org.CreatedBy,
		&org.UpdatedAt,
		&org.UpdatedBy,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[READ] failed: %v\n", err)
		return nil, err
	}

	return &org, nil
}

/// GetOrganization

const getOrganizationQuery = `
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
WHERE
`

func (store *orgStore) GetOrganization(ctx context.Context, organization *org.Organization) (*org.Organization, error) {
	availableFields := []string{}
	organizationFields := map[string]interface{}{
		"sid":        organization.ScopedId,
		"name":       organization.Name,
		"visibility": organization.Visibility,
		"createdAt":  organization.CreatedAt,
		"updatedAt":  organization.UpdatedAt,
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
	finalQuery := getOrganizationQuery + strings.Join(availableFields, " AND ")

	row := store.Conn.QueryRow(ctx, finalQuery)

	var org org.Organization
	err := row.Scan(
		&org.Id,
		&org.ScopedId,
		&org.Name,
		&org.Description,
		&org.Visibility,
		&org.CreatedAt,
		&org.CreatedBy,
		&org.UpdatedAt,
		&org.UpdatedBy,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[READ] failed: %v\n", err)
		return nil, err
	}

	return &org, nil
}
