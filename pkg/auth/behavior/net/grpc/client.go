package grpc

import (
	"apppathway.com/pkg/net/grpc/tools"
	"apppathway.com/pkg/user/behavior/api/behaviorpb"
	"context"
	"encoding/gob"
	"google.golang.org/grpc"
	"io"
)

type Client struct {
	behaviorpb.BehaviorServiceClient
	conn *grpc.ClientConn
}

// NewClient returns a new instance of Client.
func NewClient(connStr, caFilePath string) *Client {
	connection := tools.Connect(connStr, caFilePath)
	return &Client{
		BehaviorServiceClient: behaviorpb.NewBehaviorServiceClient(connection),
		conn:                  connection,
	}
}

func (c *Client) LogCmd(ctx context.Context, rw io.ReadWriter) error {
	serve := decodeEncode[*behaviorpb.LogCmdRequest, *behaviorpb.LogCmdResponse]
	return serve(rw, ctx, c.conn, c.BehaviorServiceClient.LogCmd)
}

func decodeEncode[r ReqReader, w ResWriter](
	rw io.ReadWriter, ctx context.Context, conn *grpc.ClientConn,
	handler func(context.Context, r, ...grpc.CallOption) (w, error)) error {

	var err error
	var req r
	if err := gob.NewDecoder(rw).Decode(&req); err != nil {
		return err
	}

	var res w
	if res, err = handler(ctx, req); err != nil {
		return err
	}

	tools.Close(conn)
	if err := gob.NewEncoder(rw).Encode(res); err != nil {
		return err
	}
	return nil
}
