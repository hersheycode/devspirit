package grpc

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net"

	"apppathway.com/examples/prodapi/pkg/plugins/scheduler"
	"apppathway.com/examples/prodapi/pkg/plugins/scheduler/api/schedulerpb"
	"apppathway.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type SchedulerServer struct {
	server *grpc.Server
	Addr   string
	scheduler.SchedulerService
}

func NewSchedulerServer(conf *tls.Config) *SchedulerServer {
	creds := credentials.NewTLS(conf)
	return &SchedulerServer{
		server: grpc.NewServer(grpc.Creds(creds)),
	}
}

func Open(s *SchedulerServer) error {
	schedulerpb.RegisterSchedulerServer(s.server, s)
	lis, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return errors.UnexpectedError(fmt.Errorf("error while listening: %v", err))
	}
	log.Printf("[dev-server] builder/cd listening at %v \n", s.Addr)

	if err := s.server.Serve(lis); err != nil {
		return errors.UnexpectedError(fmt.Errorf("error while serving: %v", err))
	}
	return nil
}

// Close gracefully shuts down the server.
func (s *SchedulerServer) Close() {
	s.server.GracefulStop()
}

func (s *SchedulerServer) Register(ctx context.Context, req *schedulerpb.RegisterRequest) (*schedulerpb.RegisterResponse, error) {
	res, err := s.SchedulerService.Register(ctx, scheduler.RegisterSchedulerReq{
		Time: req.Time,
	})
	if err != nil {
		return nil, err
	}
	return &schedulerpb.RegisterResponse{Status: res.Status}, nil
}
