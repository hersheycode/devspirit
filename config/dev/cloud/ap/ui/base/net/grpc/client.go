package grpc

import (
	"apppathway.com/pkg/builder/base/api/cpluginpb"
	"apppathway.com/pkg/net/grpc/tools"
	"context"
	"encoding/gob"
	"google.golang.org/grpc"
	"io"
)

type Client struct {
	cpluginpb.CPluginServiceClient
	conn *grpc.ClientConn
}

// NewClient returns a new instance of Client.
func NewClient(connStr, caFilePath string) *Client {
	connection := tools.Connect(connStr, caFilePath)
	return &Client{
		CPluginServiceClient: cpluginpb.NewCPluginServiceClient(connection),
		conn:                 connection,
	}
}

func (c *Client) Create(ctx context.Context, rw io.ReadWriter) error {
	serve := decodeEncode[*cpluginpb.CreateRequest, *cpluginpb.CreateResponse]
	return serve(rw, ctx, c.conn, c.CPluginServiceClient.Create)
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
