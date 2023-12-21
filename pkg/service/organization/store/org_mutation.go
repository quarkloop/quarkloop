package store

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	gonanoid "github.com/matoous/go-nanoid/v2"
	org "github.com/quarkloop/quarkloop/pkg/service/organization"
)

/// CreateOrganization

const createOrganizationMutation = `
INSERT INTO
  "system"."Organization" ("sid", "name", "description", "visibility", "createdBy")
VALUES
  (@sid, @name, @description, @visibility, @createdBy)
RETURNING 
  "id", "sid", "name", "description", "visibility", "createdAt", "createdBy";
`

func (store *orgStore) CreateOrganization(ctx context.Context, organization *org.Organization) (*org.Organization, error) {
	if organization.ScopedId == "" {
		sid, err := gonanoid.New()
		if err != nil {
			return nil, err
		}
		organization.ScopedId = sid
	}

	row := store.Conn.QueryRow(
		ctx,
		createOrganizationMutation,
		pgx.NamedArgs{
			"sid":         organization.ScopedId,
			"name":        organization.Name,
			"description": organization.Description,
			"visibility":  organization.Visibility,
			"createdBy":   organization.CreatedBy,
		},
	)

	var org org.Organization
	rowErr := row.Scan(
		&org.Id,
		&org.ScopedId,
		&org.Name,
		&org.Description,
		&org.Visibility,
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
  "visibility"  = @visibility,
  "updatedAt"   = @updatedAt,
  "updatedBy"   = @updatedBy,
WHERE
  "id" = @id;
`

func (store *orgStore) UpdateOrganizationById(ctx context.Context, orgId int, org *org.Organization) error {
	commandTag, err := store.Conn.Exec(
		ctx,
		updateOrganizationByIdMutation,
		pgx.NamedArgs{
			"id":          orgId,
			"sid":         org.ScopedId,
			"name":        org.Name,
			"description": org.Description,
			"visibility":  org.Visibility,
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

func (store *orgStore) DeleteOrganizationById(ctx context.Context, orgId int) error {
	commandTag, err := store.Conn.Exec(ctx, deleteOrganizationByIdMutation, pgx.NamedArgs{"id": orgId})
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
