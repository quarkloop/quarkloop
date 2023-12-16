package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/quarkloop/quarkloop/pkg/model"
)

/// ListTableBranches

const listTableBranchesQuery = `
SELECT
  "id", "name", "description", "createdAt", "createdBy", "updatedAt", "updatedBy"
FROM
  "project"."TableBranch"
WHERE
  "projectId" = @projectId;
`

func (r *Repository) ListTableBranches(ctx context.Context, projectId int) ([]model.TableBranch, error) {
	rows, err := r.ProjectDbConn.Query(ctx, listTableBranchesQuery, pgx.NamedArgs{"projectId": projectId})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var branchList []model.TableBranch

	for rows.Next() {
		var branch model.TableBranch
		err := rows.Scan(
			&branch.Id,
			&branch.Name,
			&branch.Description,
			&branch.CreatedAt,
			&branch.CreatedBy,
			&branch.UpdatedAt,
			&branch.UpdatedBy,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
			return nil, err
		}

		branchList = append(branchList, branch)
	}

	if err := rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
		return nil, err
	}

	return branchList, nil
}

/// GetTableBranchById

const getTableBranchByIdQuery = `
SELECT
  "id", "name", "description", "createdAt", "createdBy", "updatedAt", "updatedBy"
FROM
  "project"."TableBranch"
WHERE
  "id" = @id
AND
  "projectId" = @projectId;  
`

func (r *Repository) GetTableBranchById(ctx context.Context, projectId int, branchId int) (*model.TableBranch, error) {
	row := r.ProjectDbConn.QueryRow(ctx, getTableBranchByIdQuery, pgx.NamedArgs{
		"projectId": projectId,
		"id":        branchId,
	})

	var branch model.TableBranch
	err := row.Scan(
		&branch.Id,
		&branch.Name,
		&branch.Description,
		&branch.CreatedAt,
		&branch.CreatedBy,
		&branch.UpdatedAt,
		&branch.UpdatedBy,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[READ] failed: %v\n", err)
		return nil, err
	}

	return &branch, nil
}
