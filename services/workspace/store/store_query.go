package store

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/quarkloop/quarkloop/pkg/model"
	wsErrors "github.com/quarkloop/quarkloop/services/workspace/errors"
)

/// GetWorkspaceId

const getWorkspaceIdQuery = `
SELECT 
    org."id",
    ws."id"
FROM 
    "system"."Workspace" AS ws
LEFT JOIN 
    "system"."Org" AS org ON org."id" = ws."orgId"
WHERE (
	org."sid" = @orgSid
	AND
	ws."sid" = @workspaceSid
);
`

type GetWorkspaceIdQuery struct {
	OrgSid       string
	WorkspaceSid string
}

func (store *workspaceStore) GetWorkspaceId(ctx context.Context, query *GetWorkspaceIdQuery) (int64, int64, error) {
	row := store.Conn.QueryRow(ctx, getWorkspaceIdQuery, pgx.NamedArgs{
		"orgSid":       query.OrgSid,
		"workspaceSid": query.WorkspaceSid,
	})

	var orgId, workspaceId int64
	err := row.Scan(&orgId, &workspaceId)
	if err != nil {
		if err == pgx.ErrNoRows {
			return 0, 0, wsErrors.ErrWorkspaceNotFound
		}
		fmt.Fprintf(os.Stderr, "[READ] failed: %v\n", err)
		return 0, 0, err
	}

	return orgId, workspaceId, nil
}

/// GetWorkspaceById

const getWorkspaceByIdQuery = `
SELECT 
    ws."id",
    ws."sid",
    ws."orgId",
    org."sid",
    ws."name",
    ws."description",
    ws."visibility",
    ws."createdAt",
    ws."createdBy",
    ws."updatedAt",
    ws."updatedBy"
FROM 
    "system"."Workspace" AS ws
LEFT JOIN 
    "system"."Org" AS org ON org."id" = ws."orgId"
WHERE (
	ws."orgId" = @orgId
	AND
	ws."id" = @id
);
`

type GetWorkspaceByIdQuery struct {
	OrgId       int64
	WorkspaceId int64
}

func (store *workspaceStore) GetWorkspaceById(ctx context.Context, query *GetWorkspaceByIdQuery) (*model.Workspace, error) {
	row := store.Conn.QueryRow(ctx, getWorkspaceByIdQuery, pgx.NamedArgs{
		"orgId": query.OrgId,
		"id":    query.WorkspaceId,
	})

	var ws model.Workspace
	err := row.Scan(
		&ws.Id,
		&ws.ScopeId,
		&ws.OrgId,
		&ws.OrgScopeId,
		&ws.Name,
		&ws.Description,
		&ws.Visibility,
		&ws.CreatedAt,
		&ws.CreatedBy,
		&ws.UpdatedAt,
		&ws.UpdatedBy,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, wsErrors.ErrWorkspaceNotFound
		}
		fmt.Fprintf(os.Stderr, "[READ] failed: %v\n", err)
		return nil, err
	}

	return &ws, nil
}

/// GetWorkspaceVisibilityById

const getWorkspaceVisibilityByIdQuery = `
SELECT 
    "visibility"
FROM 
    "system"."Workspace"
WHERE (
	"orgId" = @orgId
	AND
	"id" = @id
);
`

type GetWorkspaceVisibilityByIdQuery struct {
	OrgId       int64
	WorkspaceId int64
}

func (store *workspaceStore) GetWorkspaceVisibilityById(ctx context.Context, query *GetWorkspaceVisibilityByIdQuery) (model.ScopeVisibility, error) {
	row := store.Conn.QueryRow(ctx, getWorkspaceVisibilityByIdQuery, pgx.NamedArgs{
		"orgId": query.OrgId,
		"id":    query.WorkspaceId,
	})

	var visibility model.ScopeVisibility
	if err := row.Scan(&visibility); err != nil {
		if err == pgx.ErrNoRows {
			return "", wsErrors.ErrWorkspaceNotFound
		}
		fmt.Fprintf(os.Stderr, "[READ] failed: %v\n", err)
		return "", err
	}

	return visibility, nil
}

/// GetWorkspace

const getWorkspaceQuery = `
SELECT 
    ws."id",
    ws."sid",
    ws."orgId",
    org."sid",
    ws."name",
    ws."description",
    ws."visibility",
    ws."createdAt",
    ws."createdBy",
    ws."updatedAt",
    ws."updatedBy"
FROM 
    "system"."Workspace" AS ws
LEFT JOIN 
    "system"."Org" AS org ON org."id" = ws."orgId"
WHERE 
    %s
ORDER BY 
    "updatedAt" ASC
LIMIT 1;
`

func (store *workspaceStore) GetWorkspace(ctx context.Context, orgId int, ws *model.Workspace) (*model.Workspace, error) {
	availableFields := []string{}
	workspaceFields := map[string]interface{}{
		"name":       ws.Name,
		"visibility": ws.Visibility,
		"createdAt":  ws.CreatedAt,
		"updatedAt":  ws.UpdatedAt,
	}
	for key, value := range workspaceFields {
		switch v := value.(type) {
		case int:
			if v != 0 {
				availableFields = append(availableFields, fmt.Sprintf("ws.\"%s\" = '%d'", key, v))
			}
		case float64:
			if v != 0.0 {
				availableFields = append(availableFields, fmt.Sprintf("ws.\"%s\" = '%f'", key, v))
			}
		case string:
			if v != "" {
				availableFields = append(availableFields, fmt.Sprintf("ws.\"%s\" = '%s'", key, v))
			}
		case *time.Time:
			if v != nil {
				availableFields = append(availableFields, fmt.Sprintf("ws.\"%s\" = '%s'", key, v))
			}
		}
	}
	finalQuery := fmt.Sprintf(getWorkspaceQuery, strings.Join(availableFields, " AND "))

	row := store.Conn.QueryRow(ctx, finalQuery)

	var workspace model.Workspace
	err := row.Scan(
		&workspace.Id,
		&workspace.ScopeId,
		&workspace.OrgId,
		&workspace.OrgScopeId,
		&workspace.Name,
		&workspace.Description,
		&workspace.Visibility,
		&workspace.CreatedAt,
		&workspace.CreatedBy,
		&workspace.UpdatedAt,
		&workspace.UpdatedBy,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[READ] failed: %v\n", err)
		return nil, err
	}

	return &workspace, nil
}

/// GetWorkspaceList

const listWorkspacesQuery = `
SELECT 
    ws."id",
    ws."sid",
    ws."orgId",
    org."sid",
    ws."name",
    ws."description",
    ws."visibility",
    ws."createdAt",
    ws."createdBy",
    ws."updatedAt",
    ws."updatedBy"
FROM 
    "system"."Workspace" AS ws
LEFT JOIN 
    "system"."Org" AS org ON org."id" = ws."orgId"
WHERE 
    ws."id" = ANY (@workspaceIdList)
%s	
`

type GetWorkspaceListQuery struct {
	WorkspaceIdList []int64
	Visibility      model.ScopeVisibility
}

func (store *workspaceStore) GetWorkspaceList(ctx context.Context, query *GetWorkspaceListQuery) ([]*model.Workspace, error) {
	var finalQuery strings.Builder
	if query.Visibility == model.PublicVisibility || query.Visibility == model.PrivateVisibility {
		finalQuery.WriteString(fmt.Sprintf(listWorkspacesQuery, `AND ws."visibility" = @visibility;`))
	} else {
		finalQuery.WriteString(fmt.Sprintf(listWorkspacesQuery, ";"))
	}

	rows, err := store.Conn.Query(ctx, finalQuery.String(), pgx.NamedArgs{
		"workspaceIdList": query.WorkspaceIdList,
		"visibility":      query.Visibility,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var wsList []*model.Workspace = []*model.Workspace{}
	for rows.Next() {
		var workspace model.Workspace
		err := rows.Scan(
			&workspace.Id,
			&workspace.ScopeId,
			&workspace.OrgId,
			&workspace.OrgScopeId,
			&workspace.Name,
			&workspace.Description,
			&workspace.Visibility,
			&workspace.CreatedAt,
			&workspace.CreatedBy,
			&workspace.UpdatedAt,
			&workspace.UpdatedBy,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
			return nil, err
		}

		wsList = append(wsList, &workspace)
	}

	if err := rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
		return nil, err
	}

	return wsList, nil
}

// /// GetProjectList

// const listProjectsQuery = `
// SELECT
//     p."id",
//     p."sid",
//     p."orgId",
//     p."workspaceId",
//     org."sid",
//     ws."sid",
//     p."name",
//     p."description",
//     p."visibility",
//     p."createdAt",
//     p."createdBy",
//     p."updatedAt",
//     p."updatedBy"
// FROM
//     "system"."Project" AS p
// LEFT JOIN
//     "system"."Org" AS org ON org."id" = p."orgId"
// LEFT JOIN
//     system."Workspace" AS ws ON ws."id" = p."workspaceId"
// WHERE
//     p."orgId" = @orgId
// AND
//     p."workspaceId" = @workspaceId
// %s
// `

// type GetProjectListQuery struct {
// 	OrgId       int64
// 	WorkspaceId int64
// 	Visibility  model.ScopeVisibility
// }

// func (store *workspaceStore) GetProjectList(ctx context.Context, query *GetProjectListQuery) ([]*model.Project, error) {
// 	whereClause := ";"
// 	if query.Visibility == model.PublicVisibility || query.Visibility == model.PrivateVisibility {
// 		whereClause = `AND p."visibility" = @visibility;`
// 	}

// 	finalQuery := fmt.Sprintf(listProjectsQuery, whereClause)

// 	rows, err := store.Conn.Query(ctx, finalQuery, pgx.NamedArgs{
// 		"orgId":       query.OrgId,
// 		"workspaceId": query.WorkspaceId,
// 		"visibility":  query.Visibility,
// 	})
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var projectList []*model.Project = []*model.Project{}

// 	for rows.Next() {
// 		var project model.Project
// 		err := rows.Scan(
// 			&project.Id,
// 			&project.ScopeId,
// 			&project.OrgId,
// 			&project.WorkspaceId,
// 			&project.OrgScopeId,
// 			&project.WorkspaceScopeId,
// 			&project.Name,
// 			&project.Description,
// 			&project.Visibility,
// 			&project.CreatedAt,
// 			&project.CreatedBy,
// 			&project.UpdatedAt,
// 			&project.UpdatedBy,
// 		)
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
// 			return nil, err
// 		}

// 		projectList = append(projectList, &project)
// 	}

// 	if err := rows.Err(); err != nil {
// 		fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
// 		return nil, err
// 	}

// 	return projectList, nil
// }

// /// GetUserAssignmentList

// const getUserAssignmentListQuery = `
// SELECT
//     ua.id,
//     ua."userId",
//     ur."name",
//     ua."createdAt",
//     ua."createdBy",
//     ua."updatedAt",
//     ua."updatedBy"
// FROM
//     "system"."UserAssignment" AS ua
// LEFT JOIN
//     "system"."UserRole" AS ur ON ur.id = ua."userRoleId"
// WHERE (
// 	ua."orgId" = @orgId
// 	AND
// 	ua."workspaceId" = @workspaceId
// )
// GROUP BY
//     ua.id,
//     ur."name"
// ORDER BY id ASC;
// `

// type GetUserAssignmentListQuery struct {
// 	OrgId       int64
// 	WorkspaceId int64
// }

// func (store *workspaceStore) GetUserAssignmentList(ctx context.Context, query *GetUserAssignmentListQuery) ([]*user.UserAssignment, error) {
// 	rows, err := store.Conn.Query(ctx, getUserAssignmentListQuery, pgx.NamedArgs{
// 		"orgId":       query.OrgId,
// 		"workspaceId": query.WorkspaceId,
// 	})
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var uaList []*user.UserAssignment = []*user.UserAssignment{}

// 	for rows.Next() {
// 		var ua user.UserAssignment
// 		err := rows.Scan(
// 			&ua.Id,
// 			&ua.UserId,
// 			&ua.Role,
// 			&ua.CreatedAt,
// 			&ua.CreatedBy,
// 			&ua.UpdatedAt,
// 			&ua.UpdatedBy,
// 		)
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
// 			return nil, err
// 		}

// 		uaList = append(uaList, &ua)
// 	}

// 	if err := rows.Err(); err != nil {
// 		fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
// 		return nil, err
// 	}

// 	return uaList, nil
// }
