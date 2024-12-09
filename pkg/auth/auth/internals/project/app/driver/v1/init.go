package authdriver

import (
	apitools "apppathway.com/pkg/net/grpc/tools"
	authPb "apppathway.com/pkg/user/auth/internals/project/api/v1"
	"google.golang.org/grpc"
)

type AuthPathway struct {
	api  authPb.AuthClient
	conn *grpc.ClientConn
}

func Init(connStr, caFilePath string) AuthPathway {
	connection := apitools.Connect(connStr, caFilePath)
	return AuthPathway{api: authPb.NewAuthClient(connection), conn: connection}
}
