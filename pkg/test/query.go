package test

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/project"
	"github.com/quarkloop/quarkloop/pkg/service/user"
	"github.com/quarkloop/quarkloop/pkg/service/workspace"
)

const getOrgListQuery = `
SELECT 
    "id",
    "sid",
    "name",
    "description",
    "visibility",
    "createdAt",
    "createdBy",
    "updatedAt",
    "updatedBy"
FROM 
    "system"."Organization"
ORDER BY id ASC;
`

func GetFullOrgList(ctx context.Context, conn *pgx.Conn) ([]*model.Org, error) {
	rows, err := conn.Query(ctx, getOrgListQuery)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var orgList []*model.Org = []*model.Org{}
	for rows.Next() {
		var org model.Org
		err := rows.Scan(
			&org.Id,
			&org.ScopeId,
			&org.Name,
			&org.Description,
			&org.Visibility,
			&org.CreatedAt,
			&org.CreatedBy,
			&org.UpdatedAt,
			&org.UpdatedBy,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
			return nil, err
		}
		orgList = append(orgList, &org)
	}
	if err := rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
		return nil, err
	}

	return orgList, nil
}

const getWorkspaceListQuery = `
SELECT 
    "id",
    "sid",
    "name",
    "description",
    "visibility",
    "createdAt",
    "createdBy",
    "updatedAt",
    "updatedBy"
FROM 
    "system"."Workspace"
ORDER BY id ASC;
`

func GetFullWorkspaceList(ctx context.Context, conn *pgx.Conn) ([]*workspace.Workspace, error) {
	rows, err := conn.Query(ctx, getWorkspaceListQuery)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var wsList []*workspace.Workspace = []*workspace.Workspace{}
	for rows.Next() {
		var ws workspace.Workspace
		err := rows.Scan(
			&ws.Id,
			&ws.ScopeId,
			&ws.Name,
			&ws.Description,
			&ws.Visibility,
			&ws.CreatedAt,
			&ws.CreatedBy,
			&ws.UpdatedAt,
			&ws.UpdatedBy,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
			return nil, err
		}
		wsList = append(wsList, &ws)
	}
	if err := rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
		return nil, err
	}

	return wsList, nil
}

const getProjectListQuery = `
SELECT 
    "id",
    "sid",
    "name",
    "description",
    "visibility",
    "createdAt",
    "createdBy",
    "updatedAt",
    "updatedBy"
FROM 
    "system"."Project"
ORDER BY id ASC;
`

func GetFullProjectList(ctx context.Context, conn *pgx.Conn) ([]*project.Project, error) {
	rows, err := conn.Query(ctx, getProjectListQuery)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var projectList []*project.Project = []*project.Project{}
	for rows.Next() {
		var ws project.Project
		err := rows.Scan(
			&ws.Id,
			&ws.ScopeId,
			&ws.Name,
			&ws.Description,
			&ws.Visibility,
			&ws.CreatedAt,
			&ws.CreatedBy,
			&ws.UpdatedAt,
			&ws.UpdatedBy,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
			return nil, err
		}
		projectList = append(projectList, &ws)
	}
	if err := rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
		return nil, err
	}

	return projectList, nil
}

const getUserListQuery = `
SELECT 
    "id",
    "username",
    "name",
    "email",
    "birthdate",
    "country",
    "image",
    "status",
    "createdAt",
    "createdBy",
    "updatedAt",
    "updatedBy"
FROM 
    "auth"."User"
ORDER BY id ASC;
`

func GetFullUserList(ctx context.Context, conn *pgx.Conn) ([]*user.User, error) {
	rows, err := conn.Query(ctx, getUserListQuery)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var userList []*user.User = []*user.User{}
	for rows.Next() {
		var u user.User
		err := rows.Scan(
			&u.Id,
			&u.Username,
			&u.Name,
			&u.Email,
			&u.Birthdate,
			&u.Country,
			&u.Image,
			&u.Status,
			&u.CreatedAt,
			&u.CreatedBy,
			&u.UpdatedAt,
			&u.UpdatedBy,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
			return nil, err
		}
		userList = append(userList, &u)
	}
	if err := rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
		return nil, err
	}

	return userList, nil
}
