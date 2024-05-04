package store

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	gonanoid "github.com/matoous/go-nanoid/v2"

	"github.com/quarkloop/quarkloop/pkg/model"
	orgErrors "github.com/quarkloop/quarkloop/services/org/errors"
)

/// CreateOrg

const createOrgMutation = `
INSERT INTO "system"."Org" (
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

type CreateOrgCommand struct {
	CreatedBy   string
	ScopeId     string
	Name        string
	Description string
	Visibility  model.ScopeVisibility
}

func (store *orgStore) CreateOrg(ctx context.Context, cmd *CreateOrgCommand) (*model.Org, error) {
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

	var o model.Org
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
		return nil, orgErrors.HandleError(err)
	}

	return &o, nil
}

/// UpdateOrgById

const updateOrgByIdMutation = `
UPDATE
    "system"."Org"
SET
    "sid"         = COALESCE (NULLIF(@sid, ''), "sid"),
    "name"        = COALESCE (NULLIF(@name, ''), "name"),
    "description" = COALESCE (NULLIF(@description, ''), "description"),
    "visibility"  = COALESCE (NULLIF(@visibility, ''), "visibility"),
    "updatedAt"   = @updatedAt,
    "updatedBy"   = @updatedBy
WHERE
    "id" = @id;
`

type UpdateOrgByIdCommand struct {
	UpdatedBy   string
	OrgId       int64
	ScopeId     string
	Name        string
	Description string
	Visibility  model.ScopeVisibility
}

func (store *orgStore) UpdateOrgById(ctx context.Context, cmd *UpdateOrgByIdCommand) error {
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
		return orgErrors.HandleError(err)
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
    "system"."Org"
WHERE
    "id" = @id;
`

type DeleteOrgByIdCommand struct {
	OrgId int64
}

func (store *orgStore) DeleteOrgById(ctx context.Context, cmd *DeleteOrgByIdCommand) error {
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
