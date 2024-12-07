package ci

import (
	"apppathway.com/pkg/builder/ci"
	"apppathway.com/pkg/client/user/behavior"
	"apppathway.com/pkg/net/grpc/tools"
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"github.com/google/goterm/term"
	"github.com/spf13/cobra"
	"io"
	"strings"
)

func (g *CmdSet) Play() *cobra.Command {
	cmd := &cobra.Command{
		Use: "play",
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
			serve := encodeDecode[ci.PlayReq, ci.PlayRes]
			req := ci.PlayReq{
				ID:   "containerID!",
				Name: name,
			}
			resp := &ci.PlayRes{}
			if err := serve(ctx, g.client.Play, req, resp); err != nil {
				panic(err)
			}
			slic := []string{"play"}
			slic = append(slic, args...)
			behavior.LogCmd(strings.Join(slic, " "))
			fmt.Println(term.Greenf("\n **** %v **** \n", resp.Status))
		}}

	cmd.PersistentFlags().String("name", "", "The name for the containerized plugin being created.")
	return cmd
}

func encodeDecode[Req ci.ReqReader, Res ci.ResWriter](ctx context.Context,
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
