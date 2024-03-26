package project

import (
	"google.golang.org/grpc"

	projectGrpc "github.com/quarkloop/quarkloop/pkg/grpc/v1/system/project"
	"github.com/quarkloop/quarkloop/services/project/store"
)

type ProjectService interface {
	projectGrpc.ProjectServiceServer

	RegisterService(s *grpc.Server)
}

type projectService struct {
	store store.ProjectStore

	projectGrpc.UnimplementedProjectServiceServer
}

func NewProjectService(ds store.ProjectStore) ProjectService {
	return &projectService{store: ds}
}

func (service *projectService) RegisterService(s *grpc.Server) {
	projectGrpc.RegisterProjectServiceServer(s, service)
}
