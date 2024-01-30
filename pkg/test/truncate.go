package test

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

const truncateAuthDbTablesQuery = `
TRUNCATE
    "auth"."VerificationToken",
    "auth"."Session",
    "auth"."Account",
    "auth"."User";
`

func TruncateAuthDBTables(ctx context.Context, conn *pgx.Conn) error {
	_, err := conn.Exec(ctx, truncateAuthDbTablesQuery)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[TRUNCATE] failed: %v\n", err)
		return err
	}

	return nil
}

// const truncateAuthzDbTablesQuery = `
// TRUNCATE
// 	"public"."alembic_version",
// 	"public"."metadata",
// 	"public"."relation_tuple_transaction",
// 	"public"."relation_tuple",
// 	"public"."namespace_config",
// 	"public"."caveat";
// `

// func TruncateAuthzDBTables(ctx context.Context, conn *pgx.Conn) error {
// 	_, err := conn.Exec(ctx, truncateAuthzDbTablesQuery)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "[TRUNCATE] failed: %v\n", err)
// 		return err
// 	}

// 	return nil
// }

const truncateSystemDbTablesQuery = `
TRUNCATE
    "system"."Role",
    "system"."ProjectMember",
	"system"."WorkspaceMember",
	"system"."OrganizationMember",
    "system"."Project",
    "system"."Workspace",
    "system"."Organization";
`

func TruncateSystemDBTables(ctx context.Context, conn *pgx.Conn) error {
	_, err := conn.Exec(ctx, truncateSystemDbTablesQuery)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[TRUNCATE] failed: %v\n", err)
		return err
	}

	return nil
}

const truncateEngineDbTablesQuery = `
TRUNCATE
    "engine"."TableColumn",
	"engine"."TableRecord",
    "engine"."Table";
`

func TruncateEngineDbTables(ctx context.Context, conn *pgx.Conn) error {
	_, err := conn.Exec(ctx, truncateEngineDbTablesQuery)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[TRUNCATE] failed: %v\n", err)
		return err
	}

	return nil
}
