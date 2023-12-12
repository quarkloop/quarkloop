package repository

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	gonanoid "github.com/matoous/go-nanoid/v2"

	"github.com/quarkloop/quarkloop/pkg/model"
)

/// CreateOrganization

const createOrganizationMutation = `
INSERT INTO
  "system"."Organization" ("id", "name", "accessType", "description", "path")
VALUES
  (@id, @name, @accessType, @description, @path)
RETURNING 
  "id", "name", "accessType", "path", "description", "createdAt";
`

func (r *Repository) CreateOrganization(ctx context.Context, organization *model.Organization) (*model.Organization, error) {
	if organization.Id == "" {
		id, err := gonanoid.New()
		if err != nil {
			return nil, err
		}
		organization.Id = id
	}
	organization.Path = fmt.Sprintf("/org/%s", organization.Id)

	row := r.SystemDbConn.QueryRow(
		ctx,
		createOrganizationMutation,
		pgx.NamedArgs{
			"id":          organization.Id,
			"name":        organization.Name,
			"accessType":  organization.AccessType,
			"description": organization.Description,
			"path":        organization.Path,
		},
	)

	var org model.Organization
	rowErr := row.Scan(
		&org.Id,
		&org.Name,
		&org.AccessType,
		&org.Path,
		&org.Description,
		&org.CreatedAt,
	)
	if rowErr != nil {
		fmt.Fprintf(os.Stderr, "[CREATE] failed: %v\n", rowErr)
		return nil, rowErr
	}

	return &org, nil
}

/// UpdateOrganizationById

const updateOrganizationByIdMutation = `
UPDATE
  "system"."Organization"
SET
  "name"        = @name,
  "accessType"  = @accessType,
  "description" = @description,
  "path"        = @path,
  "updatedAt"   = @updatedAt
WHERE
  "id" = @id;
`

func (r *Repository) UpdateOrganizationById(ctx context.Context, orgId string, org *model.Organization) error {
	commandTag, err := r.SystemDbConn.Exec(
		ctx,
		updateOrganizationByIdMutation,
		pgx.NamedArgs{
			"id":          orgId,
			"name":        org.Name,
			"accessType":  org.AccessType,
			"description": org.Description,
			"path":        org.Path,
			"updatedAt":   time.Now(),
		},
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[UPDATE] failed: %v\n", err)
		return err
	}

	if commandTag.RowsAffected() != 1 {
		notFoundErr := errors.New("cannot find to update")
		fmt.Fprintf(os.Stderr, "[UPDATE] failed: %v\n", notFoundErr)
		return notFoundErr
	}

	return nil
}

/// DeleteOrganizationById

const deleteOrganizationByIdMutation = `
DELETE FROM
  "system"."Organization"
WHERE
  "id" = @id;
`

func (r *Repository) DeleteOrganizationById(ctx context.Context, orgId string) error {
	commandTag, err := r.SystemDbConn.Exec(ctx, deleteOrganizationByIdMutation, pgx.NamedArgs{"id": orgId})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[DELETE] failed: %v\n", err)
		return err
	}

	if commandTag.RowsAffected() != 1 {
		notFoundErr := errors.New("cannot find to delete")
		fmt.Fprintf(os.Stderr, "[DELETE] failed: %v\n", notFoundErr)
		return notFoundErr
	}

	return nil
}
