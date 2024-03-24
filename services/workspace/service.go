package workspace

import (
	"google.golang.org/grpc"

	wsGrpc "github.com/quarkloop/quarkloop/pkg/grpc/v1/system/workspace"
	"github.com/quarkloop/quarkloop/services/workspace/store"
)

type WorkspaceService interface {
	wsGrpc.WorkspaceServiceServer

	RegisterService(s *grpc.Server)
}

type workspaceService struct {
	store store.WorkspaceStore

	wsGrpc.UnimplementedWorkspaceServiceServer
}

func NewWorkspaceService(ds store.WorkspaceStore) WorkspaceService {
	return &workspaceService{store: ds}
}

func (service *workspaceService) RegisterService(s *grpc.Server) {
	wsGrpc.RegisterWorkspaceServiceServer(s, service)
}
