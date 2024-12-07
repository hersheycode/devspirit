package grpc

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net"

	"apppathway.com/examples/prodapi/pkg/plugins/sms"
	"apppathway.com/examples/prodapi/pkg/plugins/sms/api/smspb"
	"apppathway.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type SMSServer struct {
	server *grpc.Server
	Addr   string
	sms.SMSService
}

func NewSMSServer(conf *tls.Config) *SMSServer {
	creds := credentials.NewTLS(conf)
	return &SMSServer{
		server: grpc.NewServer(grpc.Creds(creds)),
	}
}

func Open(s *SMSServer) error {
	smspb.RegisterSMSServer(s.server, s)
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
func (s *SMSServer) Close() {
	s.server.GracefulStop()
}

func (s *SMSServer) Send(ctx context.Context, req *smspb.SendRequest) (*smspb.SendResponse, error) {
	res, err := s.SMSService.Send(ctx, sms.SendReq{
		PhoneNum: req.PhoneNum,
		Email:    req.Email,
		Message: sms.Message{
			Body: req.Msg.Body,
		},
	})
	if err != nil {
		return nil, err
	}
	return &smspb.SendResponse{Status: res.Status}, nil
}
