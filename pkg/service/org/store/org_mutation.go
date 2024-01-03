package store

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/quarkloop/quarkloop/pkg/service/org"
)

/// CreateOrg

const createOrgMutation = `
INSERT INTO "system"."Organization" (
    "sid",
    "name",
    "description",
    "visibility",
    "createdBy"
)
VALUES (
    @sid,
    @name,
    @description,
    @visibility,
    @createdBy
)
RETURNING 
    "id",
    "sid",
    "name",
    "description",
    "visibility",
    "createdAt",
    "createdBy";
`

func (store *orgStore) CreateOrg(ctx context.Context, organization *org.Org) (*org.Org, error) {
	if organization.ScopeId == "" {
		sid, err := gonanoid.New()
		if err != nil {
			return nil, err
		}
		organization.ScopeId = sid
	}

	row := store.Conn.QueryRow(ctx, createOrgMutation, pgx.NamedArgs{
		"sid":         organization.ScopeId,
		"name":        organization.Name,
		"description": organization.Description,
		"visibility":  organization.Visibility,
		"createdBy":   organization.CreatedBy,
	})

	var org org.Org
	rowErr := row.Scan(
		&org.Id,
		&org.ScopeId,
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

/// UpdateOrgById

const updateOrgByIdMutation = `
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

func (store *orgStore) UpdateOrgById(ctx context.Context, orgId int, org *org.Org) error {
	commandTag, err := store.Conn.Exec(ctx, updateOrgByIdMutation, pgx.NamedArgs{
		"id":          orgId,
		"sid":         org.ScopeId,
		"name":        org.Name,
		"description": org.Description,
		"visibility":  org.Visibility,
		"updatedAt":   time.Now(),
		"updatedBy":   org.UpdatedBy,
	})
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

/// DeleteOrgById

const deleteOrgByIdMutation = `
DELETE FROM
    "system"."Organization"
WHERE
    "id" = @id;
`

func (store *orgStore) DeleteOrgById(ctx context.Context, orgId int) error {
	commandTag, err := store.Conn.Exec(ctx, deleteOrgByIdMutation, pgx.NamedArgs{"id": orgId})
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
