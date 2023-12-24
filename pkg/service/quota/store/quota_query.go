package store

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/quarkloop/quarkloop/pkg/service/quota"
)

const getOrgQuotaMetricQuery = `
SELECT 
    COUNT(DISTINCT o.id) AS org_count,
    COUNT(DISTINCT w.id) AS workspace_count,
    COUNT(DISTINCT p.id) AS project_count,
    COUNT(DISTINCT ou.id) AS org_user_count,
    COUNT(DISTINCT wu.id) AS workspace_user_count,
    COUNT(DISTINCT pu.id) AS project_user_count
FROM 
    "system"."Organization" As o
LEFT JOIN 
    "system"."Workspace" AS w ON w."orgId" = o."id"
LEFT JOIN 
    "system"."Project" AS p ON p."orgId" = o."id" AND p."workspaceId" = w."id"
LEFT JOIN 
    "system"."User" AS ou ON ou."orgId" = o."id"
LEFT JOIN 
    "system"."User" AS wu ON wu."orgId" = o."id" AND wu."workspaceId" = w."id"
LEFT JOIN 
    "system"."User" AS pu ON pu."orgId" = o."id" AND pu."workspaceId" = w."id" AND pu."projectId" = p."id"
WHERE 
    o."id" = 59;
`

const getOrgUserQuotaMetricQuery = `
SELECT feature, metric 
FROM (
    SELECT 
       --w.id AS workspaceId,
       --COUNT(DISTINCT o.id)  AS org_count,
       COUNT(DISTINCT w.id)  AS workspace_count,
       --COUNT(DISTINCT p.id)  AS project_count,
       --COUNT(DISTINCT ou.id) AS org_user_count,
       COUNT(DISTINCT wu.id) AS workspace_user_count
       --COUNT(DISTINCT pu.id) AS project_user_count
    FROM 
      "system"."Organization"      As o
    LEFT JOIN "system"."Workspace" AS w  ON w."orgId"  = o."id"
    LEFT JOIN "system"."Project"   AS p  ON p."orgId"  = o."id" AND p."workspaceId" = w."id"
    LEFT JOIN "system"."User"      AS ou ON ou."orgId" = o."id"
    LEFT JOIN "system"."User"      AS wu ON wu."orgId" = o."id" AND wu."workspaceId" = w."id"
    LEFT JOIN "system"."User"      AS pu ON pu."orgId" = o."id" AND pu."workspaceId" = w."id" AND pu."projectId" = p."id"
    WHERE 
      o."id" = 59
) AS metrics
CROSS JOIN jsonb_each_text(to_jsonb(metrics)) as cols(feature, metric);
`

const getWorkspacePerOrgQuotaMetricQuery = `
SELECT feature, metric 
FROM (
    SELECT 
       COUNT(DISTINCT w.id)  AS workspace_count,
       COUNT(DISTINCT wu.id) AS workspace_user_count
    FROM 
      "system"."Workspace"    As w
    LEFT JOIN "system"."User" AS wu ON wu."orgId" = w."orgId" AND wu."workspaceId" = w."id"
    WHERE 
      w."orgId" = 59
) AS metrics
CROSS JOIN jsonb_each_text(to_jsonb(metrics)) as cols(feature, metric);
`

/// GetProjectPerWorkspaceQuotaMetric

const getProjectPerWorkspaceQuotaMetricQuery = `
SELECT feature, metric
FROM (
    (
        SELECT 
           COUNT(DISTINCT o.id)  AS org_count,
           COUNT(DISTINCT ou.id) AS org_user_count
        FROM 
          "system"."Organization" As o
        LEFT JOIN "system"."User" AS ou ON ou."orgId" = o."id"
        WHERE 
          o."id" = 59
    ) AS org
    CROSS JOIN 
    (
        SELECT 
           COUNT(DISTINCT w.id)  AS workspace_count,
           COUNT(DISTINCT wu.id) AS workspace_user_count
        FROM 
          "system"."Workspace"    As w
        LEFT JOIN "system"."User" AS wu ON wu."orgId" = w."orgId" AND wu."workspaceId" = w."id"
        WHERE 
          w."orgId" = 59
    ) AS workspace
    CROSS JOIN 
    (
        SELECT 
           COUNT(DISTINCT p.id)  AS project_count,
           COUNT(DISTINCT pu.id) AS project_user_count
        FROM 
          "system"."Project"    As p
        LEFT JOIN "system"."User" AS pu ON pu."orgId" = p."orgId" AND pu."workspaceId" = p."workspaceId" AND pu."projectId" = p."id"
        WHERE 
          p."orgId" = 59
    ) AS project
) AS metrics
CROSS JOIN jsonb_each_text(to_jsonb(metrics)) as cols(feature, metric)
ORDER BY feature;
`

/// GetQuotasByUserId

const getQuotasByUserIdQuery = `
SELECT 
    feature, 
    metric 
FROM (
    SELECT 
       COUNT(DISTINCT id) AS org_count,
    FROM 
      "system"."Organization"
    WHERE
      "userId" = @userId
) AS metrics
CROSS JOIN 
    jsonb_each_text(to_jsonb(metrics)) as cols(feature, metric);
`

func (store *quotaStore) GetQuotasByUserId(ctx context.Context, userId int) (quota.Quota, error) {
	row := store.Conn.QueryRow(ctx, getQuotasByUserIdQuery, pgx.NamedArgs{
		"userId": userId,
	})

	var q quota.Quota
	if err := row.Scan(&q); err != nil {
		fmt.Fprintf(os.Stderr, "[READ] failed: %v\n", err)
		return quota.Quota{}, err
	}

	return q, nil
}

/// GetQuotasByOrgId

const getQuotasByOrgIdQuery = `
SELECT 
    feature, 
    metric 
FROM (
    SELECT 
       COUNT(DISTINCT w.id) AS workspace_count,
       COUNT(DISTINCT p.id) AS project_count,
       COUNT(DISTINCT u.id) AS org_user_count
    FROM 
      "system"."Organization" As o
    LEFT JOIN 
        "system"."Workspace" AS w ON w."orgId" = o."id"
    LEFT JOIN 
        "system"."Project" AS p ON p."orgId" = o."id"
    LEFT JOIN 
        "system"."User" AS u ON u."orgId" = o."id"
    WHERE
      o."id" = @orgId
) AS metrics
CROSS JOIN 
    jsonb_each_text(to_jsonb(metrics)) as cols(feature, metric);
`

func (store *quotaStore) GetQuotasByOrgId(ctx context.Context, orgId int) ([]quota.Quota, error) {
	rows, err := store.Conn.Query(ctx, getQuotasByOrgIdQuery, pgx.NamedArgs{
		"orgId": orgId,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var quotaList []quota.Quota = []quota.Quota{}

	for rows.Next() {
		var quota quota.Quota
		if err := rows.Scan(&quota.Feature, &quota.Metric); err != nil {
			fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
			return nil, err
		}

		quotaList = append(quotaList, quota)
	}

	if err := rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
		return nil, err
	}

	return quotaList, nil
}
