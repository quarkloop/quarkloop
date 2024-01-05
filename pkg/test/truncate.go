package test

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

const truncateTablesQuery = `
TRUNCATE
    "system"."Permission",
    "system"."UserRole",
    "system"."UserGroup",
    "system"."UserAssignment",
    "system"."Project",
    "system"."Workspace",
    "system"."Organization";
`

func TruncateSystemDBTables(ctx context.Context, conn *pgx.Conn) error {
	_, err := conn.Exec(ctx, truncateTablesQuery)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[TRUNCATE] failed: %v\n", err)
		return err
	}

	return nil
}
