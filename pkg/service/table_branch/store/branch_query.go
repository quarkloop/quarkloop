package store

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/quarkloop/quarkloop/pkg/service/table_branch"
)

/// ListTableBranches

const listTableBranchesQuery = `
SELECT 
    "id",
    "name",
    "description",
    "createdAt",
    "createdBy",
    "updatedAt",
    "updatedBy"
FROM 
    "project"."TableBranch"
WHERE 
    "projectId" = @projectId;
`

func (store *tableBranchStore) ListTableBranches(ctx context.Context, projectId int32) ([]table_branch.TableBranch, error) {
	rows, err := store.Conn.Query(ctx, listTableBranchesQuery, pgx.NamedArgs{"projectId": projectId})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var branchList []table_branch.TableBranch

	for rows.Next() {
		var branch table_branch.TableBranch
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
    "id",
    "name",
    "description",
    "createdAt",
    "createdBy",
    "updatedAt",
    "updatedBy"
FROM 
    "project"."TableBranch"
WHERE 
    "id" = @id 
AND 
    "projectId" = @projectId;
`

func (store *tableBranchStore) GetTableBranchById(ctx context.Context, projectId int32, branchId int32) (*table_branch.TableBranch, error) {
	row := store.Conn.QueryRow(ctx, getTableBranchByIdQuery, pgx.NamedArgs{
		"projectId": projectId,
		"id":        branchId,
	})

	var branch table_branch.TableBranch
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
