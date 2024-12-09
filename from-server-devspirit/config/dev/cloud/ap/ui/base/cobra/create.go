package cobra

import (
	"apppathway.com/pkg/builder/base"
	"apppathway.com/pkg/net/grpc/tools"
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"github.com/google/goterm/term"
	"github.com/spf13/cobra"
	"io"
)

func (g *CmdSet) Create() *cobra.Command {
	cmd := &cobra.Command{
		Use: "create",
		Run: func(cmd *cobra.Command, args []string) {
			name, err := cmd.Flags().GetString("name")
			if err != nil {
				panic(err)
			}
			if name == "" {
				fmt.Println(term.Red("name for plugin is required"))
				return
			}
			userId, tok := tools.Creds()
			md := tools.NewMeta(map[string]string{
				"authorization": tok,
				"ownerId":       userId,
			})
			ctx := tools.NewOutCtx(context.Background(), md)
			serve := encodeDecode[base.CreateReq, base.CreateRes]
			req := base.CreateReq{
				Name:    name,
				Content: []byte{},
			}
			resp := &base.CreateRes{}
			if err := serve(ctx, g.client.Create, req, resp); err != nil {
				panic(err)
			}
			fmt.Println(term.Greenf("\n **** %v **** \n", resp.Status))
		}}

	cmd.PersistentFlags().String("name", "", "The name for the containerized plugin being created.")

	return cmd
}

func encodeDecode[Req base.ReqReader, Res base.ResWriter](ctx context.Context,
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
