package accesscontrol_impl

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strconv"

	v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/quarkloop/quarkloop/pkg/service/accesscontrol"
)

func (s *aclService) GetOrgList(ctx context.Context, query *accesscontrol.GetOrgListQuery) ([]int32, error) {
	return s.getResources(ctx, query.Permission, "org", strconv.FormatInt(int64(query.UserId), 10))
}

func (s *aclService) GetWorkspaceList(ctx context.Context, query *accesscontrol.GetWorkspaceListQuery) ([]int32, error) {
	return s.getResources(ctx, query.Permission, "workspace", strconv.FormatInt(int64(query.UserId), 10))
}

func (s *aclService) GetProjectList(ctx context.Context, query *accesscontrol.GetProjectListQuery) ([]int32, error) {
	return s.getResources(ctx, query.Permission, "project", strconv.FormatInt(int64(query.UserId), 10))
}

func (s *aclService) getResources(ctx context.Context, permission, resource, userId string) ([]int32, error) {
	var resourceList []int32 = []int32{}
	resp, err := s.authz.LookupResources(ctx, &v1.LookupResourcesRequest{
		Permission:         permission,
		ResourceObjectType: resource,
		Subject: &v1.SubjectReference{
			Object: &v1.ObjectReference{
				ObjectType: "user",
				ObjectId:   userId,
			},
		},
		Consistency: &v1.Consistency{
			Requirement: &v1.Consistency_FullyConsistent{FullyConsistent: true},
		},
	})
	if err != nil {
		return resourceList, err
	}

	for {
		res, err := resp.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			fmt.Printf("\n[LookupResources] (%+v, %+v, %+v) => %+v\n", permission, resource, userId, err)
		}

		resourceId, err := strconv.ParseInt(res.ResourceObjectId, 10, 32)
		if err != nil {
			panic(err)
		}

		resourceList = append(resourceList, int32(resourceId))
	}
	return resourceList, nil
}
