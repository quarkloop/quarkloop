package org_impl

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/org"
	"github.com/quarkloop/quarkloop/pkg/service/org/store"
	"github.com/quarkloop/quarkloop/service/v1/system"
	grpc "github.com/quarkloop/quarkloop/service/v1/system/org"
)

type orgService struct {
	store store.OrgStore

	grpc.UnimplementedOrgServiceServer
}

func NewOrgService(ds store.OrgStore) org.Service {
	return &orgService{store: ds}
}

func (s *orgService) GetOrgList(ctx context.Context, query *grpc.GetOrgListQuery) (*grpc.GetOrgListReply, error) {
	orgList, err := s.store.GetOrgList(ctx, &org.GetOrgListQuery{
		OrgIdList:  query.OrgIdList,
		Visibility: model.ScopeVisibility(query.Visibility),
	})
	if err != nil {
		return nil, err
	}

	res := &grpc.GetOrgListReply{OrgList: make([]*system.Org, len(orgList))}
	for i, org := range orgList {
		if org == nil {
			continue
		}

		org.GeneratePath()
		proto := org.Proto()
		res.OrgList[i] = proto
	}

	return res, nil
}

func (s *orgService) GetOrgById(ctx context.Context, query *grpc.GetOrgByIdQuery) (*grpc.GetOrgByIdReply, error) {
	org, err := s.store.GetOrgById(ctx, &org.GetOrgByIdQuery{OrgId: query.OrgId})
	if err != nil {
		return nil, err
	}

	org.GeneratePath()
	proto := org.Proto()

	reply := &grpc.GetOrgByIdReply{Org: proto}
	return reply, nil
}

func (s *orgService) GetOrgVisibilityById(ctx context.Context, query *grpc.GetOrgVisibilityByIdQuery) (*grpc.GetOrgVisibilityByIdReply, error) {
	visibility, err := s.store.GetOrgVisibilityById(ctx, &org.GetOrgVisibilityByIdQuery{OrgId: query.OrgId})
	if err != nil {
		return nil, err
	}

	reply := &grpc.GetOrgVisibilityByIdReply{Visibility: int32(visibility)}
	return reply, nil
}

func (s *orgService) CreateOrg(ctx context.Context, cmd *grpc.CreateOrgCommand) (*grpc.CreateOrgReply, error) {
	if cmd.Name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "request missing required field: Name")
	} else if cmd.Description == "" {
		return nil, status.Errorf(codes.InvalidArgument, "request missing required field: Description")
	} else if cmd.CreatedBy == "" {
		return nil, status.Errorf(codes.InvalidArgument, "request missing required field: CreatedBy")
	}

	o, err := s.store.CreateOrg(ctx, &org.CreateOrgCommand{
		CreatedBy:   cmd.CreatedBy,
		ScopeId:     cmd.ScopeId,
		Name:        cmd.Name,
		Description: cmd.Description,
		Visibility:  model.ScopeVisibility(cmd.Visibility),
	})
	if err != nil {
		if err == org.ErrOrgAlreadyExists {
			return nil, status.Errorf(codes.AlreadyExists, err.Error())
		}
		return nil, status.Errorf(codes.Internal, "something went wrong in server")
	}

	o.GeneratePath()
	proto := o.Proto()

	reply := &grpc.CreateOrgReply{Org: proto}
	return reply, nil
}

func (s *orgService) UpdateOrgById(ctx context.Context, cmd *grpc.UpdateOrgByIdCommand) (*emptypb.Empty, error) {
	err := s.store.UpdateOrgById(ctx, &org.UpdateOrgByIdCommand{
		OrgId:       cmd.OrgId,
		UpdatedBy:   cmd.UpdatedBy,
		ScopeId:     cmd.ScopeId,
		Name:        cmd.Name,
		Description: cmd.Description,
		Visibility:  model.ScopeVisibility(cmd.Visibility),
	})
	return &emptypb.Empty{}, err
}

func (s *orgService) DeleteOrgById(ctx context.Context, cmd *grpc.DeleteOrgByIdCommand) (*emptypb.Empty, error) {
	err := s.store.DeleteOrgById(ctx, &org.DeleteOrgByIdCommand{OrgId: cmd.OrgId})
	return &emptypb.Empty{}, err
}

func (s *orgService) GetWorkspaceList(ctx context.Context, query *grpc.GetWorkspaceListQuery) (*grpc.GetWorkspaceListReply, error) {
	workspaceList, err := s.store.GetWorkspaceList(ctx, &org.GetWorkspaceListQuery{
		OrgId:      query.OrgId,
		Visibility: model.ScopeVisibility(query.Visibility),
	})
	if err != nil {
		return nil, err
	}

	reply := &grpc.GetWorkspaceListReply{WorkspaceList: make([]*system.Workspace, len(workspaceList))}
	for i, ws := range workspaceList {
		if ws == nil {
			continue
		}

		ws.GeneratePath()
		proto := ws.Proto()
		reply.WorkspaceList[i] = proto
	}

	return reply, nil
}

func (s *orgService) GetProjectList(ctx context.Context, query *grpc.GetProjectListQuery) (*grpc.GetProjectListReply, error) {
	projectList, err := s.store.GetProjectList(ctx, &org.GetProjectListQuery{
		OrgId:      query.OrgId,
		Visibility: model.ScopeVisibility(query.Visibility),
	})
	if err != nil {
		return nil, err
	}

	reply := &grpc.GetProjectListReply{ProjectList: make([]*system.Project, len(projectList))}
	for i, project := range projectList {
		if project == nil {
			continue
		}

		project.GeneratePath()
		proto := project.Proto()
		reply.ProjectList[i] = proto
	}

	return reply, nil
}

// func (s *orgService) GetUserAssignmentList(ctx context.Context, query *grpc.GetUserAssignmentListQuery) (*grpc.GetUserAssignmentListReply, error) {
// 	uaList, err := s.store.GetUserAssignmentList(ctx, query)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return uaList, nil
// }
