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
  "system"."Organization" ("sid", "name", "description", "accessType", "createdBy")
VALUES
  (@sid, @name, @description, @accessType, @createdBy)
RETURNING 
  "id", "sid", "name", "description", "accessType", "createdAt", "createdBy";
`

func (r *Repository) CreateOrganization(ctx context.Context, organization *model.Organization) (*model.Organization, error) {
	if organization.ScopedId == "" {
		sid, err := gonanoid.New()
		if err != nil {
			return nil, err
		}
		organization.ScopedId = sid
	}

	row := r.SystemDbConn.QueryRow(
		ctx,
		createOrganizationMutation,
		pgx.NamedArgs{
			"sid":         organization.ScopedId,
			"name":        organization.Name,
			"description": organization.Description,
			"accessType":  organization.AccessType,
			"createdBy":   organization.CreatedBy,
		},
	)

	var org model.Organization
	rowErr := row.Scan(
		&org.Id,
		&org.ScopedId,
		&org.Name,
		&org.Description,
		&org.AccessType,
		&org.CreatedAt,
		&org.CreatedBy,
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
  "sid"         = @sid, 
  "name"        = @name,
  "description" = @description,
  "accessType"  = @accessType,
  "updatedAt"   = @updatedAt,
  "updatedBy"   = @updatedBy,
WHERE
  "id" = @id;
`

func (r *Repository) UpdateOrganizationById(ctx context.Context, orgId int, org *model.Organization) error {
	commandTag, err := r.SystemDbConn.Exec(
		ctx,
		updateOrganizationByIdMutation,
		pgx.NamedArgs{
			"id":          orgId,
			"sid":         org.ScopedId,
			"name":        org.Name,
			"description": org.Description,
			"accessType":  org.AccessType,
			"updatedAt":   time.Now(),
			"updatedBy":   org.UpdatedBy,
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

func (r *Repository) DeleteOrganizationById(ctx context.Context, orgId int) error {
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
