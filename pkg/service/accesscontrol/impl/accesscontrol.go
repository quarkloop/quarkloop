package accesscontrol_impl

import (
	"context"
	"fmt"
	"log"

	v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/authzed/authzed-go/v1"
	"github.com/authzed/grpcutil"
	"github.com/quarkloop/quarkloop/pkg/service/accesscontrol"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type aclService struct {
	authz *authzed.Client
}

func NewAccessControlService() accesscontrol.Service {
	client, err := authzed.NewClient(
		"localhost:50051",
		grpcutil.WithInsecureBearerToken("my_passphrase_key"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("unable to initialize client: %s", err)
	}

	return &aclService{
		authz: client,
	}
}

func (s *aclService) EvaluateUserAccess(ctx context.Context, query *accesscontrol.EvaluateQuery) (bool, error) {
	objectType := ""
	objectId := ""
	if query.OrgId != 0 && query.WorkspaceId != 0 && query.ProjectId != 0 {
		objectType = "project"
		objectId = fmt.Sprintf("%d", query.ProjectId)
	} else if query.OrgId != 0 && query.WorkspaceId != 0 {
		objectType = "workspace"
		objectId = fmt.Sprintf("%d", query.WorkspaceId)
	} else if query.OrgId != 0 {
		objectType = "org"
		objectId = fmt.Sprintf("%d", query.OrgId)
	}

	fmt.Printf("\nresource: => %+v => %+v => %+v\n\n", objectType, objectId, query)

	rResp, err := s.authz.CheckPermission(context.Background(), &v1.CheckPermissionRequest{
		Resource: &v1.ObjectReference{
			ObjectType: objectType,
			ObjectId:   objectId,
		},
		Permission: query.Permission,
		Subject: &v1.SubjectReference{
			Object: &v1.ObjectReference{
				ObjectType: "user",
				ObjectId:   fmt.Sprintf("%d", query.UserId),
			},
		},
	})
	if err != nil {
		log.Fatalf("failed to check permission: %s", err)
		return false, err
	}

	if rResp.Permissionship == v1.CheckPermissionResponse_PERMISSIONSHIP_HAS_PERMISSION {
		return true, nil
	}
	if rResp.Permissionship == v1.CheckPermissionResponse_PERMISSIONSHIP_NO_PERMISSION {
		return false, nil
	}

	return false, nil
}

func (s *aclService) GrantUserAccess(ctx context.Context, cmd *accesscontrol.GrantUserAccessCommand) error {
	objectType := ""
	objectId := ""
	relation := cmd.Role
	userId := fmt.Sprintf("%d", cmd.UserId)

	if cmd.OrgId != 0 && cmd.WorkspaceId != 0 && cmd.ProjectId != 0 {
		objectType = "project"
		objectId = fmt.Sprintf("%d", cmd.ProjectId)
	} else if cmd.OrgId != 0 && cmd.WorkspaceId != 0 {
		objectType = "workspace"
		objectId = fmt.Sprintf("%d", cmd.WorkspaceId)
	} else if cmd.OrgId != 0 {
		objectType = "org"
		objectId = fmt.Sprintf("%d", cmd.OrgId)
	}

	relationships := []*v1.RelationshipUpdate{
		{
			Operation: v1.RelationshipUpdate_OPERATION_CREATE,
			Relationship: &v1.Relationship{
				Resource: &v1.ObjectReference{
					ObjectType: objectType,
					ObjectId:   objectId,
				},
				Relation: relation,
				Subject: &v1.SubjectReference{
					Object: &v1.ObjectReference{
						ObjectType: "user",
						ObjectId:   userId,
					},
				},
			},
		},
	}
	_, err := s.authz.WriteRelationships(ctx, &v1.WriteRelationshipsRequest{Updates: relationships})
	if err != nil {
		return err
	}

	return nil

	// if cmd.OrgId != 0 {
	// 	query := &accesscontrol.GetOrgMemberByUserIdQuery{UserId: cmd.UserId, OrgId: 0}
	// 	_, err := s.store.GetOrgMemberByUserId(ctx, query)
	// 	if err != nil {
	// 		return false, err
	// 	}

	// 	cmd := &accesscontrol.CreateOrgMemberCommand{
	// 		OrgId:     cmd.OrgId,
	// 		UserId:    cmd.UserId,
	// 		RoleId:    cmd.RoleId,
	// 		CreatedBy: "admin",
	// 	}
	// 	_, err = s.store.CreateOrgMember(ctx, cmd)
	// 	if err != nil {
	// 		return false, err
	// 	}
	// } else if cmd.WorkspaceId != 0 {
	// 	query := &accesscontrol.GetOrgMemberByUserIdQuery{UserId: cmd.UserId, OrgId: 0}
	// 	orgMember, err := s.store.GetOrgMemberByUserId(ctx, query)
	// 	if err != nil && err != org.ErrOrgMemberNotFound {
	// 		return false, err
	// 	}

	// 	roleId := cmd.RoleId
	// 	membership := accesscontrol.DirectMembership
	// 	sourceId := 0

	// 	if orgMember != nil {
	// 		roleId = orgMember.RoleId
	// 		membership = accesscontrol.InheritedMembership
	// 		sourceId = orgMember.OrgId
	// 	}

	// 	cmd := &accesscontrol.CreateWorkspaceMemberCommand{
	// 		WorkspaceId: cmd.WorkspaceId,
	// 		UserId:      cmd.UserId,
	// 		RoleId:      roleId,
	// 		Type:        membership, // direct, inherited
	// 		Source:      sourceId,   // orgId if membership is inherited
	// 		CreatedBy:   "admin",
	// 	}
	// 	_, err = s.store.CreateWorkspaceMember(ctx, cmd)
	// 	if err != nil {
	// 		return false, err
	// 	}
	// } else if cmd.ProjectId != 0 {
	// 	orgQuery := &accesscontrol.GetOrgMemberByUserIdQuery{UserId: cmd.UserId, OrgId: 0}
	// 	orgMember, err := s.store.GetOrgMemberByUserId(ctx, orgQuery)
	// 	if err != nil && err != org.ErrOrgMemberNotFound {
	// 		return false, err
	// 	}

	// 	wsQuery := &accesscontrol.GetWorkspaceMemberByUserIdQuery{UserId: cmd.UserId, WorkspaceId: 0}
	// 	wsMember, err := s.store.GetWorkspaceMemberByUserId(ctx, wsQuery)
	// 	if err != nil && err != org.ErrOrgMemberNotFound {
	// 		return false, err
	// 	}

	// 	roleId := cmd.RoleId
	// 	membership := accesscontrol.DirectMembership
	// 	sourceId := 0

	// 	if orgMember != nil {
	// 		roleId = orgMember.RoleId
	// 		membership = accesscontrol.InheritedMembership
	// 		sourceId = orgMember.OrgId
	// 	} else if wsMember != nil && wsMember.Type == accesscontrol.DirectMembership {
	// 		roleId = wsMember.RoleId
	// 		membership = accesscontrol.InheritedMembership
	// 		sourceId = wsMember.WorkspaceId
	// 	}

	// 	cmd := &accesscontrol.CreateProjectMemberCommand{
	// 		ProjectId: cmd.ProjectId,
	// 		UserId:    cmd.UserId,
	// 		RoleId:    roleId,
	// 		Type:      membership, // direct, inherited
	// 		Source:    sourceId,   // orgId if membership is inherited
	// 		CreatedBy: "admin",
	// 	}
	// 	_, err = s.store.CreateProjectMember(ctx, cmd)
	// 	if err != nil {
	// 		return false, err
	// 	}
	// }

	// return true, nil
}

func (s *aclService) RevokeUserAccess(ctx context.Context, cmd *accesscontrol.RevokeUserAccessCommand) error {
	objectType := ""
	objectId := ""
	relation := cmd.Role
	userId := fmt.Sprintf("%d", cmd.UserId)

	if cmd.OrgId != 0 && cmd.WorkspaceId != 0 && cmd.ProjectId != 0 {
		objectType = "project"
		objectId = fmt.Sprintf("%d", cmd.ProjectId)
	} else if cmd.OrgId != 0 && cmd.WorkspaceId != 0 {
		objectType = "workspace"
		objectId = fmt.Sprintf("%d", cmd.WorkspaceId)
	} else if cmd.OrgId != 0 {
		objectType = "org"
		objectId = fmt.Sprintf("%d", cmd.OrgId)
	}

	req := &v1.DeleteRelationshipsRequest{
		RelationshipFilter: &v1.RelationshipFilter{
			ResourceType:       objectType,
			OptionalResourceId: objectId,
			OptionalRelation:   relation,
			OptionalSubjectFilter: &v1.SubjectFilter{
				SubjectType:       "user",
				OptionalSubjectId: userId,
			},
		},
	}
	_, err := s.authz.DeleteRelationships(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
