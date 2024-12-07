package behavior

import (
	"apppathway.com/pkg/net/grpc/tools"
	"apppathway.com/pkg/user/behavior"
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"github.com/google/goterm/term"
	"io"
)

func LogCmd(command string) {
	userId, tok := tools.Creds()
	md := tools.NewMeta(map[string]string{
		"authorization": tok,
		"ownerId":       userId,
	})
	ctx := tools.NewOutCtx(context.Background(), md)
	serve := encodeDecode[behavior.LogCmdReq, behavior.LogCmdRes]
	req := behavior.LogCmdReq{
		Command: command,
	}
	resp := &behavior.LogCmdRes{}
	if err := serve(ctx, client.LogCmd, req, resp); err != nil {
		panic(err)
	}
	fmt.Println(term.Greenf("\n **** %v **** \n", resp.Status))
}

func encodeDecode[Req behavior.ReqReader, Res behavior.ResWriter](ctx context.Context,
	handler func(context.Context, io.ReadWriter) error, r Req, w *Res) error {

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
