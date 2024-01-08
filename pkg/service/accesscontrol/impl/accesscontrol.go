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

func (s *aclService) Evaluate(ctx context.Context, query *accesscontrol.EvaluateQuery) error {
	hasPermission, err := s.store.Evaluate(ctx, query)
	if err != nil {
		return err
	}

	if !hasPermission {
		return accesscontrol.ErrPermissionDenied
	}

	return nil
}

func (s *aclService) GetUserAccessList(ctx context.Context, query *accesscontrol.GetUserAssignmentListQuery) ([]accesscontrol.UserAssignment, error) {
	uaList, err := s.store.GetUserAssignmentList(ctx, query)
	return uaList, err
}

func (s *aclService) GetUserAccessById(ctx context.Context, query *accesscontrol.GetUserAssignmentByIdQuery) (*accesscontrol.UserAssignment, error) {
	ua, err := s.store.GetUserAssignmentById(ctx, query)
	return ua, err
}

func (s *aclService) GrantUserAccess(ctx context.Context, cmd *accesscontrol.CreateUserAssignmentCommand) (*accesscontrol.UserAssignment, error) {
	ua, err := s.store.CreateUserAssignment(ctx, cmd)
	return ua, err
}

func (s *aclService) UpdateUserAccessById(ctx context.Context, cmd *accesscontrol.UpdateUserAssignmentByIdCommand) error {
	err := s.store.UpdateUserAssignmentById(ctx, cmd)
	return err
}

func (s *aclService) RevokeUserAccessById(ctx context.Context, cmd *accesscontrol.DeleteUserAssignmentByIdCommand) error {
	err := s.store.DeleteUserAssignmentById(ctx, cmd)
	return err
}
