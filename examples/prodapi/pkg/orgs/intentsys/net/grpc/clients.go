package grpc

import (
	"fmt"

	"apppathway.com/examples/prodapi/pkg/orgs/intentsys/api/intentsyspb"
	"apppathway.com/examples/prodapi/pkg/plugins/intent/api/intentpb"
	"apppathway.com/examples/prodapi/pkg/plugins/scheduler/api/schedulerpb"
	"apppathway.com/examples/prodapi/pkg/plugins/sms/api/smspb"
	"apppathway.com/pkg/net/grpc/tools"
	"google.golang.org/grpc"
)

type IntentClient struct {
	intentpb.IntentClient
	conn *grpc.ClientConn
}

type SchedulerClient struct {
	schedulerpb.SchedulerClient
	conn *grpc.ClientConn
}

type SMSClient struct {
	smspb.SMSClient
	conn *grpc.ClientConn
}

type IntentSysClient struct {
	intentsyspb.IntentSysClient
	conn *grpc.ClientConn
}

func NewIntentClient(connStr, caFilePath string) *IntentClient {
	connection := tools.Connect(connStr, caFilePath)
	return &IntentClient{
		IntentClient: intentpb.NewIntentClient(connection),
		conn:         connection,
	}
}

func NewSchedulerClient(connStr, caFilePath string) *SchedulerClient {
	connection := tools.Connect(connStr, caFilePath)
	return &SchedulerClient{
		SchedulerClient: schedulerpb.NewSchedulerClient(connection),
		conn:            connection,
	}
}

func NewSMSClient(connStr, caFilePath string) *SMSClient {
	connection := tools.Connect(connStr, caFilePath)
	return &SMSClient{
		SMSClient: smspb.NewSMSClient(connection),
		conn:      connection,
	}
}

func NewIntentSysClient(connStr, caFilePath string) *IntentSysClient {
	fmt.Println("Creating client for intentsysd...", connStr)
	connection := tools.Connect(connStr, caFilePath)
	return &IntentSysClient{
		IntentSysClient: intentsyspb.NewIntentSysClient(connection),
		conn:            connection,
	}
}
