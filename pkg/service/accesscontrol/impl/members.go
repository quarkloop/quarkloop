package accesscontrol_impl

import (
	"context"
	"errors"
	"strconv"

	v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/quarkloop/quarkloop/pkg/service/accesscontrol"
)

func (s *aclService) GetOrgMemberList(ctx context.Context, query *accesscontrol.GetOrgMemberListQuery) ([]*accesscontrol.Member, error) {
	return s.getMembers(ctx, "org", strconv.FormatInt(int64(query.OrgId), 10))
}

func (s *aclService) GetWorkspaceMemberList(ctx context.Context, query *accesscontrol.GetWorkspaceMemberListQuery) ([]*accesscontrol.Member, error) {
	return s.getMembers(ctx, "workspace", strconv.FormatInt(int64(query.WorkspaceId), 10))
}

func (s *aclService) GetProjectMemberList(ctx context.Context, query *accesscontrol.GetProjectMemberListQuery) ([]*accesscontrol.Member, error) {
	return s.getMembers(ctx, "project", strconv.FormatInt(int64(query.ProjectId), 10))
}

func (s *aclService) getMembers(ctx context.Context, resource string, resourceId string) ([]*accesscontrol.Member, error) {
	resp, err := s.authz.ExpandPermissionTree(ctx, &v1.ExpandPermissionTreeRequest{
		Permission: "all",
		Resource: &v1.ObjectReference{
			ObjectType: resource,
			ObjectId:   resourceId,
		},
		Consistency: &v1.Consistency{
			Requirement: &v1.Consistency_FullyConsistent{FullyConsistent: true},
		},
	})
	if err != nil {
		return nil, err
	}

	return s.expandPermissionTreeRecursive(resp.GetTreeRoot()), nil
}

func (s *aclService) expandPermissionTreeRecursive(root *v1.PermissionRelationshipTree) []*accesscontrol.Member {
	var memberList []*accesscontrol.Member = []*accesscontrol.Member{}
	if root == nil {
		return memberList
	}

	intermediate := root.GetIntermediate()
	for _, child := range intermediate.GetChildren() {
		leaf := child.GetLeaf()
		if leaf != nil {
			for _, subject := range leaf.Subjects {
				resource := child.GetExpandedObject()

				userId, err := strconv.ParseInt(subject.Object.ObjectId, 10, 32)
				if err != nil {
					panic(err)
				}
				objectId, err := strconv.ParseInt(resource.ObjectId, 10, 32)
				if err != nil {
					panic(err)
				}

				member := &accesscontrol.Member{
					UserId: int32(userId),
					Role:   child.GetExpandedRelation(),
				}
				switch resource.ObjectType {
				case "org":
					member.OrgId = int32(objectId)
				case "workspace":
					member.WorkspaceId = int32(objectId)
				case "project":
					member.ProjectId = int32(objectId)
				default:
					panic(errors.New("unknown resource type"))
				}
				memberList = append(memberList, member)
			}
		}
		memberList = append(memberList, s.expandPermissionTreeRecursive(child)...)
	}

	return memberList
}
