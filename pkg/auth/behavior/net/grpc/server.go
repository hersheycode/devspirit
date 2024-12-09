package grpc

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"net"

	"apppathway.com/pkg/errors"
	"apppathway.com/pkg/user/behavior"
	"apppathway.com/pkg/user/behavior/api/behaviorpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Server struct {
	server *grpc.Server
	Addr   string
	behavior.BehaviorService
}

type ReqReader interface {
	*behaviorpb.LogCmdRequest
}

type ResWriter interface {
	*behaviorpb.LogCmdResponse
}

type handleFunc func(context.Context, io.ReadWriter) error

func NewServer(conf *tls.Config) *Server {
	creds := credentials.NewTLS(conf)
	return &Server{
		server: grpc.NewServer(grpc.Creds(creds)),
	}
}

func (s *Server) Open() error {
	behaviorpb.RegisterBehaviorServiceServer(s.server, s)
	lis, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return errors.UnexpectedError(fmt.Errorf("error while listening: %v", err))
	}
	log.Printf("[dev-server] user/behavior listening at %v \n", s.Addr)

	if err := s.server.Serve(lis); err != nil {
		return errors.UnexpectedError(fmt.Errorf("error while serving: %v", err))
	}
	return nil
}

// Close gracefully shuts down the server.
func (s *Server) Close() {
	s.server.GracefulStop()
}

func (s *Server) LogCmd(ctx context.Context, req *behaviorpb.LogCmdRequest) (*behaviorpb.LogCmdResponse, error) {
	serve := encodeDecode[*behaviorpb.LogCmdRequest, *behaviorpb.LogCmdResponse]
	resp := &behaviorpb.LogCmdResponse{}
	if err := serve(ctx, s.BehaviorService.LogCmd, req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

//Template for next rpc
// func (s *Server) LogCmd(ctx context.Context, req *behaviorpb.LogCmdRequest) (*behaviorpb.LogCmdResponse, error) {
// 	serve := encodeDecode[*behaviorpb.LogCmdRequest, *behaviorpb.LogCmdResponse]
// 	resp := &behaviorpb.LogCmdResponse{}
// 	if err := serve(ctx, s.BehaviorService.LogCmd, req, resp); err != nil {
// 		return nil, err
// 	}
// 	return resp, nil
// }

func encodeDecode[Req ReqReader, Resp ResWriter](ctx context.Context,
	handler handleFunc, r Req, w Resp) error {

	rw := &bytes.Buffer{}
	if err := gob.NewEncoder(rw).Encode(r); err != nil {
		return err
	}
	err := handler(ctx, rw)
	if err != nil {
		return err
	}
	return gob.NewDecoder(rw).Decode(w)
}
