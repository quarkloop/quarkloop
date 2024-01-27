package accesscontrol_impl

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"strconv"

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
}

func (s *aclService) RevokeUserAccess(ctx context.Context, cmd *accesscontrol.RevokeUserAccessCommand) error {
	objectType := ""
	objectId := ""
	relation := cmd.Role

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

	filter := &v1.RelationshipFilter{
		ResourceType:       objectType,
		OptionalResourceId: objectId,
	}

	if relation != "" {
		filter.OptionalRelation = relation
	}
	if cmd.UserId != 0 {
		filter.OptionalSubjectFilter = &v1.SubjectFilter{
			SubjectType:       "user",
			OptionalSubjectId: fmt.Sprintf("%d", cmd.UserId),
		}
	}

	req := &v1.DeleteRelationshipsRequest{
		RelationshipFilter:            filter,
		OptionalAllowPartialDeletions: true,
	}
	_, err := s.authz.DeleteRelationships(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (s *aclService) RevokeUserAccessByResourceType(ctx context.Context, cmd *accesscontrol.RevokeUserAccessCommand) error {
	objectType := ""
	objectId := ""
	//relation := cmd.Role
	//userId := fmt.Sprintf("%d", cmd.UserId)

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

	userList, err := s.getResourceUsers(ctx, objectType, objectId)
	if err != nil {
		return err
	}

	for _, userId := range userList {
		req := &v1.DeleteRelationshipsRequest{
			RelationshipFilter: &v1.RelationshipFilter{
				ResourceType:       objectType,
				OptionalResourceId: objectId,
				//OptionalRelation:   relation,
				OptionalSubjectFilter: &v1.SubjectFilter{
					SubjectType:       "user",
					OptionalSubjectId: strconv.FormatInt(int64(userId), 10),
				},
			},
		}
		_, err := s.authz.DeleteRelationships(ctx, req)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *aclService) getResourceUsers(ctx context.Context, resource, resourceId string) ([]int32, error) {
	var userList []int32 = []int32{}
	resp, err := s.authz.LookupSubjects(ctx, &v1.LookupSubjectsRequest{
		Permission:        "all",
		SubjectObjectType: "user",
		Resource: &v1.ObjectReference{
			ObjectType: resource,
			ObjectId:   resourceId,
		},
		Consistency: &v1.Consistency{
			Requirement: &v1.Consistency_FullyConsistent{FullyConsistent: true},
		},
	})
	if err != nil {
		return userList, err
	}

	for {
		res, err := resp.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			fmt.Printf("\n[LookupSubjects] (%+v, %+v) => %+v\n", resource, resourceId, err)
		}

		userId, err := strconv.ParseInt(res.Subject.SubjectObjectId, 10, 32)
		if err != nil {
			panic(err)
		}

		userList = append(userList, int32(userId))
	}
	return userList, nil
}

func (s *aclService) MakeParentResource(ctx context.Context, cmd *accesscontrol.MakeParentResourceCommand) error {
	parent := cmd.ParentResource
	parentId := strconv.FormatInt(int64(cmd.ParentResourceId), 10)
	child := cmd.ChildResource
	childId := strconv.FormatInt(int64(cmd.ChildResourceId), 10)

	relationships := []*v1.RelationshipUpdate{
		{
			Operation: v1.RelationshipUpdate_OPERATION_CREATE,
			Relationship: &v1.Relationship{
				Resource: &v1.ObjectReference{
					ObjectType: child,
					ObjectId:   childId,
				},
				Relation: "parent",
				Subject: &v1.SubjectReference{
					Object: &v1.ObjectReference{
						ObjectType: parent,
						ObjectId:   parentId,
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
}
