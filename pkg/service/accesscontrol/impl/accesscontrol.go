package accesscontrol_impl

import (
	"context"

	"github.com/quarkloop/quarkloop/pkg/service/accesscontrol"
	"github.com/quarkloop/quarkloop/pkg/service/accesscontrol/store"
)

type aclService struct {
	store store.AccessControlStore
}

func NewAccessControlService(ds store.AccessControlStore) accesscontrol.Service {
	return &aclService{
		store: ds,
	}
}

func (s *aclService) Evaluate(ctx context.Context, permission string, p *accesscontrol.EvaluateFilterQuery) error {
	hasPermission, err := s.store.Evaluate(ctx, permission, p.OrgId, p.WorkspaceId, p.ProjectId, p.UserId)
	if err != nil {
		return err
	}

	if !hasPermission {
		return accesscontrol.ErrPermissionDenied
	}

	return nil
}

func (s *aclService) ListUserAccesses(ctx context.Context, orgId int) ([]accesscontrol.UserAssignment, error) {
	uaList, err := s.store.ListUserAssignments(ctx, orgId)
	return uaList, err
}

func (s *aclService) GetUserAccessById(ctx context.Context, userAssignmentId int) (*accesscontrol.UserAssignment, error) {
	ua, err := s.store.GetUserAssignmentById(ctx, userAssignmentId)
	return ua, err
}

func (s *aclService) GrantUserAccess(ctx context.Context, orgId int, userRole *accesscontrol.UserAssignment) (*accesscontrol.UserAssignment, error) {
	ua, err := s.store.CreateUserAssignment(ctx, orgId, userRole)
	return ua, err
}

func (s *aclService) UpdateUserAccessById(ctx context.Context, userAssignmentId int, userRole *accesscontrol.UserAssignment) error {
	err := s.store.UpdateUserAssignmentById(ctx, userAssignmentId, userRole)
	return err
}

func (s *aclService) RevokeUserAccessById(ctx context.Context, orgId int, userAssignmentId int) error {
	err := s.store.DeleteUserAssignmentById(ctx, orgId, userAssignmentId)
	return err
}
