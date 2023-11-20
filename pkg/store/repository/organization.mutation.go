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
  "id", "name", "accessType", "description", "path", "createdAt";
`

func (r *Repository) CreateOrganization(ctx context.Context, org *model.Organization) (*model.Organization, error) {
	id, err := gonanoid.New()
	if err != nil {
		return nil, err
	}

	org.Id = id
	org.Path = fmt.Sprintf("/org/%s", id)

	commandTag, err := r.SystemDbConn.Exec(
		ctx,
		createOrganizationMutation,
		pgx.NamedArgs{
			"id":          org.Id,
			"name":        org.Name,
			"accessType":  org.AccessType,
			"description": org.Description,
			"path":        org.Path,
		},
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[CREATE] failed: %v\n", err)
		return nil, err
	}

	if commandTag.RowsAffected() != 1 {
		notFoundErr := errors.New("cannot find to create")
		fmt.Fprintf(os.Stderr, "[CREATE] failed: %v\n", notFoundErr)
		return nil, notFoundErr
	}

	return org, nil
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
