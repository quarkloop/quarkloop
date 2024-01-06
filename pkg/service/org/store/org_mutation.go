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

func (store *orgStore) CreateOrg(ctx context.Context, cmd *org.CreateOrgCommand) (*org.Org, error) {
	if cmd.ScopeId == "" {
		sid, err := gonanoid.New()
		if err != nil {
			return nil, err
		}
		cmd.ScopeId = sid
	}

	row := store.Conn.QueryRow(ctx, createOrgMutation, pgx.NamedArgs{
		"sid":         cmd.ScopeId,
		"name":        cmd.Name,
		"description": cmd.Description,
		"visibility":  cmd.Visibility,
		"createdBy":   cmd.CreatedBy,
	})

	var o org.Org
	err := row.Scan(
		&o.Id,
		&o.ScopeId,
		&o.Name,
		&o.Description,
		&o.Visibility,
		&o.CreatedAt,
		&o.CreatedBy,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[CREATE] failed: %v\n", err)
		return nil, org.HandleError(err)
	}

	return &o, nil
}

/// UpdateOrgById

const updateOrgByIdMutation = `
UPDATE
    "system"."Organization"
SET
    "sid"         = COALESCE (NULLIF(@sid, ''), "sid"),
    "name"        = COALESCE (NULLIF(@name, ''), "name"),
    "description" = COALESCE (NULLIF(@description, ''), "description"),
    "visibility"  = COALESCE (NULLIF(@visibility, 0), "visibility"),
    "updatedAt"   = @updatedAt,
    "updatedBy"   = @updatedBy
WHERE
    "id" = @id;
`

func (store *orgStore) UpdateOrgById(ctx context.Context, cmd *org.UpdateOrgByIdCommand) error {
	commandTag, err := store.Conn.Exec(ctx, updateOrgByIdMutation, pgx.NamedArgs{
		"id":          cmd.OrgId,
		"sid":         cmd.ScopeId,
		"name":        cmd.Name,
		"description": cmd.Description,
		"visibility":  cmd.Visibility,
		"updatedBy":   cmd.UpdatedBy,
		"updatedAt":   time.Now(),
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[UPDATE] failed: %v\n", err)
		return org.HandleError(err)
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

func (store *orgStore) DeleteOrgById(ctx context.Context, cmd *org.DeleteOrgByIdCommand) error {
	commandTag, err := store.Conn.Exec(ctx, deleteOrgByIdMutation, pgx.NamedArgs{"id": cmd.OrgId})
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
