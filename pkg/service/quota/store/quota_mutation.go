package store

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
)

/// UpdateQuotaLimit

const updateQuotaLimitMutation = `
UPDATE
  "system"."Quota" AS q
SET
  "orgLimit"                 = CASE WHEN  @orgLimit                 > 0 THEN  @orgLimit                 ELSE q."orgLimit" END,
  "orgUserLimit"             = CASE WHEN  @orgUserLimit             > 0 THEN  @orgUserLimit             ELSE q."orgUserLimit" END,
  "workspacePerOrgLimit"     = CASE WHEN  @workspacePerOrgLimit     > 0 THEN  @workspacePerOrgLimit     ELSE q."workspacePerOrgLimit" END,
  "projectPerWorkspaceLimit" = CASE WHEN  @projectPerWorkspaceLimit > 0 THEN  @projectPerWorkspaceLimit ELSE q."projectPerWorkspaceLimit" END,
  "updatedAt"                = @updatedAt,
  "updatedBy"                = @updatedBy
WHERE
  "userId" = @userId;
`

func (store *quotaStore) UpdateQuotaLimits(ctx context.Context, userId int, limit QuoataLimit) error {
	commandTag, err := store.Conn.Exec(ctx, updateQuotaLimitMutation, pgx.NamedArgs{
		"userId":                   userId,
		"orgLimit":                 limit.OrgLimit,
		"orgUserLimit":             limit.OrgUserLimit,
		"workspacePerOrgLimit":     limit.WorkspacePerOrgLimit,
		"projectPerWorkspaceLimit": limit.ProjectPerWorkspaceLimit,
		"updatedAt":                time.Now(),
		"updatedBy":                "TODO", // TODO: userId
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
