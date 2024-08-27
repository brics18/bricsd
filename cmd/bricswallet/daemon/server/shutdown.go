package server

import (
	"context"
	"github.com/brics18/bricsd/cmd/bricswallet/daemon/pb"
)

func (s *server) Shutdown(ctx context.Context, request *pb.ShutdownRequest) (*pb.ShutdownResponse, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	close(s.shutdown)
	return &pb.ShutdownResponse{}, nil
}
