package grpc

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net"

	"apppathway.com/examples/prodapi/pkg/plugins/intent"
	"apppathway.com/pkg/errors"
	"codestore.localhost/crudusrs/crud_basic/api/crudbasicsic/api/intentpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type IntentServer struct {
	server *grpc.Server
	Addr   string
	intent.IntentService
}

func NewIntentServer(conf *tls.Config) *IntentServer {
	creds := credentials.NewTLS(conf)
	return &IntentServer{
		server: grpc.NewServer(grpc.Creds(creds)),
	}
}

func Open(ts *IntentServer) error {
	intentpb.RegisterIntentServer(ts.server, ts)
	lis, err := net.Listen("tcp", ts.Addr)
	if err != nil {
		return errors.UnexpectedError(fmt.Errorf("error while listening: %v", err))
	}
	log.Printf("[dev-server] builder/cd listening at %v \n", ts.Addr)

	if err := ts.server.Serve(lis); err != nil {
		return errors.UnexpectedError(fmt.Errorf("error while serving: %v", err))
	}
	return nil
}

// Close gracefully shuts down the server.
func (s *IntentServer) Close() {
	s.server.GracefulStop()
}

func (is *IntentServer) Register(ctx context.Context, req *intentpb.RegisterRequest) (*intentpb.RegisterResponse, error) {
	res, err := is.IntentService.Register(ctx, intent.RegisterIntentReq{
		Name: req.Name,
	})
	if err != nil {
		return nil, err
	}
	return &intentpb.RegisterResponse{Status: res.Status}, nil
}
