package store

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/quarkloop/quarkloop/pkg/service/table_branch"
)

type TableBranchStore interface {
	ListTableBranches(ctx context.Context, projectId int32) ([]table_branch.TableBranch, error)
	GetTableBranchById(ctx context.Context, projectId int32, branchId int32) (*table_branch.TableBranch, error)
	CreateTableBranch(ctx context.Context, projectId int32, branch *table_branch.TableBranch) (*table_branch.TableBranch, error)
	UpdateTableBranchById(ctx context.Context, projectId int32, branchId int32, branch *table_branch.TableBranch) error
	DeleteTableBranchById(ctx context.Context, projectId int32, branchId int32) error
}

type tableBranchStore struct {
	Conn *pgx.Conn
}

func NewTableBranchStore(conn *pgx.Conn) *tableBranchStore {
	return &tableBranchStore{
		Conn: conn,
	}
}
