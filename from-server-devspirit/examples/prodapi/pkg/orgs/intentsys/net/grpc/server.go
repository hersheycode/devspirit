package grpc

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net"

	"apppathway.com/examples/prodapi/pkg/orgs/intentsys"
	"apppathway.com/examples/prodapi/pkg/orgs/intentsys/api/intentsyspb"
	"apppathway.com/examples/prodapi/pkg/plugins/intent/api/intentpb"
	"apppathway.com/examples/prodapi/pkg/plugins/scheduler/api/schedulerpb"
	"apppathway.com/examples/prodapi/pkg/plugins/sms/api/smspb"
	"apppathway.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type IntentSysServer struct {
	server *grpc.Server
	Addr   string
	intentsys.IntentSysService
	IntentService    intentpb.IntentClient
	SchedulerService schedulerpb.SchedulerClient
	SMSService       smspb.SMSClient
}

type handleFunc func(context.Context, io.ReadWriter) error

func NewIntentSysServer(conf *tls.Config) *IntentSysServer {
	creds := credentials.NewTLS(conf)
	return &IntentSysServer{
		server: grpc.NewServer(grpc.Creds(creds)),
	}
}

func Open(is *IntentSysServer) error {
	intentsyspb.RegisterIntentSysServer(is.server, is)
	lis, err := net.Listen("tcp", is.Addr)
	if err != nil {
		return errors.UnexpectedError(fmt.Errorf("error while listening: %v", err))
	}
	log.Printf("[dev-server] builder/intentsys listening at %v \n", is.Addr)

	if err := is.server.Serve(lis); err != nil {
		return errors.UnexpectedError(fmt.Errorf("error while serving: %v", err))
	}
	return nil
}

// Close gracefully shuts down the server.
func (is *IntentSysServer) Close() {
	is.server.GracefulStop()
}

func (is *IntentSysServer) SetIntent(ctx context.Context, req *intentsyspb.SetIntentRequest) (*intentsyspb.SetIntentResponse, error) {
	intentRes, err := is.IntentService.Register(ctx, &intentpb.RegisterRequest{
		Name: req.Intent.Name,
	})
	if err != nil {
		return &intentsyspb.SetIntentResponse{}, err
	}

	schedRes, err := is.SchedulerService.Register(ctx, &schedulerpb.RegisterRequest{
		Time: req.Schedule.Time,
	})
	if err != nil {
		return &intentsyspb.SetIntentResponse{}, err
	}

	smsRes, err := is.SMSService.Send(ctx, &smspb.SendRequest{
		PhoneNum: req.Sms.Recipient.PhoneNum,
		Email:    req.Sms.Recipient.Email,
		Msg: &smspb.Message{
			Body: req.Sms.Msg.Body,
		},
	})

	if err != nil {
		return &intentsyspb.SetIntentResponse{}, err
	}

	status := fmt.Sprintf("Intent Status: %s, Scheduler Status: %s, SMSStatus: %s",
		intentRes.Status,
		schedRes.Status,
		smsRes.Status,
	)
	return &intentsyspb.SetIntentResponse{Status: status}, nil
}