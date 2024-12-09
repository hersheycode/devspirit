package behavior

import (
	"apppathway.com/pkg/user/behavior/net/grpc"
)

var client grpc.Client

func Init(connStr, caFilePath string) {
	client = *grpc.NewClient(connStr, caFilePath)
}
