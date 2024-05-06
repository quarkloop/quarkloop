package org

import (
	"google.golang.org/grpc"

	orgGrpc "github.com/quarkloop/quarkloop/pkg/grpc/v1/system/org"
	"github.com/quarkloop/quarkloop/services/org/store"
)

type OrgService interface {
	orgGrpc.OrgServiceServer

	RegisterService(s *grpc.Server)
}

type orgService struct {
	store store.OrgStore

	orgGrpc.UnimplementedOrgServiceServer
}

func NewOrgService(s store.OrgStore) OrgService {
	return &orgService{store: s}
}

func (service *orgService) RegisterService(s *grpc.Server) {
	orgGrpc.RegisterOrgServiceServer(s, service)
}
