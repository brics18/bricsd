package server

import (
	"context"
	"github.com/brics18/bricsd/cmd/bricswallet/daemon/pb"
	"github.com/brics18/bricsd/version"
)

func (s *server) GetVersion(_ context.Context, _ *pb.GetVersionRequest) (*pb.GetVersionResponse, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return &pb.GetVersionResponse{
		Version: version.Version(),
	}, nil
}
