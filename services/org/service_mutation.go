package org

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	grpc "github.com/quarkloop/quarkloop/pkg/grpc/v1/system/org"
	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/services/org/errors"
	"github.com/quarkloop/quarkloop/services/org/store"
)

func (s *orgService) CreateOrg(ctx context.Context, cmd *grpc.CreateOrgCommand) (*grpc.CreateOrgReply, error) {
	if cmd.Name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "request missing required field: Name")
	} else if cmd.Description == "" {
		return nil, status.Errorf(codes.InvalidArgument, "request missing required field: Description")
	} else if cmd.CreatedBy == "" {
		return nil, status.Errorf(codes.InvalidArgument, "request missing required field: CreatedBy")
	}

	o, err := s.store.CreateOrg(ctx, &store.CreateOrgCommand{
		CreatedBy:   cmd.CreatedBy,
		ScopeId:     cmd.ScopeId,
		Name:        cmd.Name,
		Description: cmd.Description,
		Visibility:  model.ScopeVisibility(cmd.Visibility),
	})
	if err != nil {
		if err == errors.ErrOrgAlreadyExists {
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
	err := s.store.UpdateOrgById(ctx, &store.UpdateOrgByIdCommand{
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
	err := s.store.DeleteOrgById(ctx, &store.DeleteOrgByIdCommand{OrgId: cmd.OrgId})
	return &emptypb.Empty{}, err
}
