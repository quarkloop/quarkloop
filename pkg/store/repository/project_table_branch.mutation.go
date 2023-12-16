package repository

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/quarkloop/quarkloop/pkg/model"
)

/// CreateTableBranch

const createTableBranchMutation = `
INSERT INTO
  "project"."TableBranch" ("projectId", "name", "description", "createdBy")
VALUES
  (@projectId, @name, @description, @createdBy)
RETURNING
  "id", "name", "description", "createdAt", "createdBy";
`

func (r *Repository) CreateTableBranch(ctx context.Context, projectId int, branch *model.TableBranch) (*model.TableBranch, error) {
	commandTag, err := r.ProjectDbConn.Exec(
		ctx,
		createTableBranchMutation,
		pgx.NamedArgs{
			"projectId":   projectId,
			"name":        branch.Name,
			"description": branch.Description,
			"createdBy":   branch.CreatedBy,
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

	return branch, nil
}

/// UpdateTableBranchById

const updateTableBranchByIdMutation = `
UPDATE
  "project"."TableBranch"
set
  "name"        = @name,
  "description" = @description,
  "updatedAt"   = @updatedAt,
  "updatedBy"   = @updatedBy,
WHERE
  "id" = @id
AND
  "projectId" = @projectId;
`

func (r *Repository) UpdateTableBranchById(ctx context.Context, projectId int, branchId int, branch *model.TableBranch) error {
	commandTag, err := r.ProjectDbConn.Exec(
		ctx,
		updateTableBranchByIdMutation,
		pgx.NamedArgs{
			"projectId":   projectId,
			"id":          branchId,
			"name":        branch.Name,
			"description": branch.Description,
			"updatedAt":   time.Now(),
			"updatedBy":   branch.UpdatedBy,
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

/// DeleteTableBranchById

const deleteTableBranchByIdMutation = `
DELETE FROM
  "project"."TableBranch"
WHERE
  "id" = @id
AND
  "projectId" = @projectId;
`

func (r *Repository) DeleteTableBranchById(ctx context.Context, projectId int, branchId int) error {
	commandTag, err := r.ProjectDbConn.Exec(ctx, deleteTableBranchByIdMutation, pgx.NamedArgs{
		"projectId": projectId,
		"id":        branchId,
	})
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
