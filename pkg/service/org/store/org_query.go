package store

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/org"
	"github.com/quarkloop/quarkloop/pkg/service/project"
)

/// GetOrgList

const listOrgsQuery = `
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
%s	
`

func (store *orgStore) GetOrgList(ctx context.Context, visibility model.ScopeVisibility) ([]*org.Org, error) {
	var finalQuery strings.Builder
	if visibility == model.PublicVisibility || visibility == model.PrivateVisibility {
		finalQuery.WriteString(fmt.Sprintf(listOrgsQuery, `WHERE "visibility" = @visibility;`))
	} else {
		finalQuery.WriteString(fmt.Sprintf(listOrgsQuery, ";"))
	}

	rows, err := store.Conn.Query(ctx, finalQuery.String(), pgx.NamedArgs{"visibility": visibility})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var orgList []*org.Org = []*org.Org{}

	for rows.Next() {
		var org org.Org
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

		orgList = append(orgList, &org)
	}

	if err := rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
		return nil, err
	}

	return orgList, nil
}

/// GetOrgById

const getOrgByIdQuery = `
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

func (store *orgStore) GetOrgById(ctx context.Context, orgId int) (*org.Org, error) {
	row := store.Conn.QueryRow(ctx, getOrgByIdQuery, pgx.NamedArgs{"id": orgId})

	var org org.Org
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

/// GetOrg

const getOrgQuery = `
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

func (store *orgStore) GetOrg(ctx context.Context, organization *org.Org) (*org.Org, error) {
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
	finalQuery := getOrgQuery + strings.Join(availableFields, " AND ")

	row := store.Conn.QueryRow(ctx, finalQuery)

	var org org.Org
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
    "system"."Organization" AS org ON org."id" = p."orgId"
WHERE
    p."orgId" = @orgId
%s
`

func (store *orgStore) GetProjectList(ctx context.Context, visibility model.ScopeVisibility, orgId int) ([]*project.Project, error) {
	whereClause := ";"
	if visibility == model.PublicVisibility || visibility == model.PrivateVisibility {
		whereClause = `AND p."visibility" = @visibility;`
	}

	finalQuery := fmt.Sprintf(listProjectsQuery, whereClause)

	rows, err := store.Conn.Query(ctx, finalQuery, pgx.NamedArgs{
		"orgId":      orgId,
		"visibility": visibility,
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
