package org

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/quarkloop/quarkloop/pkg/grpc/v1/system"
	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/services/org/errors"
	"github.com/quarkloop/quarkloop/services/org/store"

	grpc "github.com/quarkloop/quarkloop/pkg/grpc/v1/system/org"
)

func (s *orgService) GetOrgId(ctx context.Context, query *grpc.GetOrgIdQuery) (*grpc.GetOrgIdReply, error) {
	orgId, err := s.store.GetOrgId(ctx, query.OrgSid)
	if err != nil {
		switch err {
		case errors.ErrOrgNotFound:
			return nil, status.Errorf(codes.NotFound, err.Error())
		case errors.ErrOrgAlreadyExists:
			return nil, status.Errorf(codes.AlreadyExists, err.Error())
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &grpc.GetOrgIdReply{OrgId: orgId}, nil
}

func (s *orgService) GetOrgById(ctx context.Context, query *grpc.GetOrgByIdQuery) (*grpc.GetOrgByIdReply, error) {
	org, err := s.store.GetOrgById(ctx, &store.GetOrgByIdQuery{OrgId: query.OrgId})
	if err != nil {
		return nil, err
	}

	org.GeneratePath()
	reply := &grpc.GetOrgByIdReply{Data: org.ToProto()}
	return reply, nil
}

func (s *orgService) GetOrgVisibilityById(ctx context.Context, query *grpc.GetOrgVisibilityByIdQuery) (*grpc.GetOrgVisibilityByIdReply, error) {
	visibility, err := s.store.GetOrgVisibilityById(ctx, &store.GetOrgVisibilityByIdQuery{OrgId: query.OrgId})
	if err != nil {
		return nil, err
	}

	reply := &grpc.GetOrgVisibilityByIdReply{Visibility: visibility.ToString()}
	return reply, nil
}

func (s *orgService) GetOrgList(ctx context.Context, query *grpc.GetOrgListQuery) (*grpc.GetOrgListReply, error) {
	data, err := s.store.GetOrgList(ctx, &store.GetOrgListQuery{
		OrgIdList:  query.OrgIdList,
		Visibility: model.ScopeVisibility(query.Visibility),
	})
	if err != nil {
		return nil, err
	}

	reply := &grpc.GetOrgListReply{Data: make([]*system.Org, len(data))}
	for i, org := range data {
		if org == nil {
			continue
		}

		org.GeneratePath()
		reply.Data[i] = org.ToProto()
	}

	return reply, nil
}

func (s *orgService) GetWorkspaceList(ctx context.Context, query *grpc.GetWorkspaceListQuery) (*grpc.GetWorkspaceListReply, error) {
	data, err := s.store.GetWorkspaceList(ctx, &store.GetWorkspaceListQuery{
		OrgId:      query.OrgId,
		Visibility: model.ScopeVisibility(query.Visibility),
	})
	if err != nil {
		return nil, err
	}

	reply := &grpc.GetWorkspaceListReply{Data: make([]*system.Workspace, len(data))}
	for i, ws := range data {
		if ws == nil {
			continue
		}

		ws.GeneratePath()
		reply.Data[i] = ws.ToProto()
	}

	return reply, nil
}

// func (s *orgService) GetProjectList(ctx context.Context, query *grpc.GetProjectListQuery) (*grpc.GetProjectListReply, error) {
// 	data, err := s.store.GetProjectList(ctx, &store.GetProjectListQuery{
// 		OrgId:      query.OrgId,
// 		Visibility: model.ScopeVisibility(query.Visibility),
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	reply := &grpc.GetProjectListReply{Data: make([]*system.Project, len(data))}
// 	for i, project := range data {
// 		if project == nil {
// 			continue
// 		}

// 		project.GeneratePath()
// 		reply.Data[i] = project.Proto()
// 	}

// 	return reply, nil
// }

// func (s *orgService) GetUserAssignmentList(ctx context.Context, query *grpc.GetUserAssignmentListQuery) (*grpc.GetUserAssignmentListReply, error) {
// 	uaList, err := s.store.GetUserAssignmentList(ctx, query)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return uaList, nil
// }
