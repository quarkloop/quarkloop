package accesscontrol_impl

import (
	"context"

	"github.com/quarkloop/quarkloop/pkg/service/accesscontrol"
	"github.com/quarkloop/quarkloop/pkg/service/accesscontrol/store"
	"github.com/quarkloop/quarkloop/pkg/service/org"
)

type aclService struct {
	store store.AccessControlStore
}

func NewAccessControlService(ds store.AccessControlStore) accesscontrol.Service {
	return &aclService{
		store: ds,
	}
}

func (s *aclService) EvaluateUserAccess(ctx context.Context, query *accesscontrol.EvaluateQuery) error {
	panic("not implemented")
}

func (s *aclService) GrantUserAccess(ctx context.Context, cmd *accesscontrol.GrantUserAccessCommand) (bool, error) {
	// prepare role permissions

	if cmd.OrgId != 0 {
		query := &accesscontrol.GetOrgMemberByUserIdQuery{UserId: cmd.UserId, OrgId: 0}
		_, err := s.store.GetOrgMemberByUserId(ctx, query)
		if err != nil {
			return false, err
		}

		cmd := &accesscontrol.CreateOrgMemberCommand{
			OrgId:     cmd.OrgId,
			UserId:    cmd.UserId,
			RoleId:    cmd.RoleId,
			CreatedBy: "admin",
		}
		_, err = s.store.CreateOrgMember(ctx, cmd)
		if err != nil {
			return false, err
		}
	} else if cmd.WorkspaceId != 0 {
		query := &accesscontrol.GetOrgMemberByUserIdQuery{UserId: cmd.UserId, OrgId: 0}
		orgMember, err := s.store.GetOrgMemberByUserId(ctx, query)
		if err != nil && err != org.ErrOrgMemberNotFound {
			return false, err
		}

		roleId := cmd.RoleId
		membership := accesscontrol.DirectMembership
		sourceId := 0

		if orgMember != nil {
			roleId = orgMember.RoleId
			membership = accesscontrol.InheritedMembership
			sourceId = orgMember.OrgId
		}

		cmd := &accesscontrol.CreateWorkspaceMemberCommand{
			WorkspaceId: cmd.WorkspaceId,
			UserId:      cmd.UserId,
			RoleId:      roleId,
			Type:        membership, // direct, inherited
			Source:      sourceId,   // orgId if membership is inherited
			CreatedBy:   "admin",
		}
		_, err = s.store.CreateWorkspaceMember(ctx, cmd)
		if err != nil {
			return false, err
		}
	} else if cmd.ProjectId != 0 {
		orgQuery := &accesscontrol.GetOrgMemberByUserIdQuery{UserId: cmd.UserId, OrgId: 0}
		orgMember, err := s.store.GetOrgMemberByUserId(ctx, orgQuery)
		if err != nil && err != org.ErrOrgMemberNotFound {
			return false, err
		}

		wsQuery := &accesscontrol.GetWorkspaceMemberByUserIdQuery{UserId: cmd.UserId, WorkspaceId: 0}
		wsMember, err := s.store.GetWorkspaceMemberByUserId(ctx, wsQuery)
		if err != nil && err != org.ErrOrgMemberNotFound {
			return false, err
		}

		roleId := cmd.RoleId
		membership := accesscontrol.DirectMembership
		sourceId := 0

		if orgMember != nil {
			roleId = orgMember.RoleId
			membership = accesscontrol.InheritedMembership
			sourceId = orgMember.OrgId
		} else if wsMember != nil && wsMember.Type == accesscontrol.DirectMembership {
			roleId = wsMember.RoleId
			membership = accesscontrol.InheritedMembership
			sourceId = wsMember.WorkspaceId
		}

		cmd := &accesscontrol.CreateProjectMemberCommand{
			ProjectId: cmd.ProjectId,
			UserId:    cmd.UserId,
			RoleId:    roleId,
			Type:      membership, // direct, inherited
			Source:    sourceId,   // orgId if membership is inherited
			CreatedBy: "admin",
		}
		_, err = s.store.CreateProjectMember(ctx, cmd)
		if err != nil {
			return false, err
		}
	}

	return true, nil
}

func (s *aclService) RevokeUserAccess(ctx context.Context, cmd *accesscontrol.RevokeUserAccessCommand) error {
	orgId := 0
	wsId := 0
	projectId := 0

	err := s.store.DeleteOrgMemberById()
}
