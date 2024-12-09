package nodes

import (
	"apppathway.com/pkg/builder/gonodes"
	"apppathway.com/pkg/builder/gonodes/net/grpc"
	"apppathway.com/pkg/client/user/behavior"
	"apppathway.com/pkg/net/grpc/tools"
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

type CmdSet struct {
	client grpc.Client
}

func NewCmdSet(connStr, caFilePath string) []*cobra.Command {
	behavior.Init(os.Getenv("BEHAVIORD_ADDRESS"), caFilePath)
	cmd := &CmdSet{
		client: *grpc.NewClient(connStr, caFilePath),
	}
	return []*cobra.Command{
		cmd.Load(),
		cmd.List(),
		cmd.Link(),
		cmd.Build(),
		cmd.Deploy(),
		cmd.Req(),
	}
}

func (g *CmdSet) Load() *cobra.Command {
	return &cobra.Command{
		Use: "load",
		Run: func(cmd *cobra.Command, args []string) {
			userId, tok := tools.Creds()
			md := tools.NewMeta(map[string]string{
				"authorization": tok,
				"ownerId":       userId,
			})
			ctx := tools.NewOutCtx(context.Background(), md)
			serve := encodeDecode[nodes.Req, nodes.Res]
			req := nodes.Req{Nodes: []nodes.Node{}}
			resp := &nodes.Res{}
			if err := serve(ctx, g.client.Load, req, resp); err != nil {
				panic(err)
			}
			for _, result := range resp.Nodes {
				fmt.Printf("\n *********************************** \n")
				out(result.Data)
				fmt.Printf("\n *********************************** \n")
			}
			slic := []string{"app load"}
			slic = append(slic, args...)
			behavior.LogCmd(strings.Join(slic, " "))
		}}
}

func (g *CmdSet) List() *cobra.Command {
	return &cobra.Command{
		Use: "list",
		Run: func(cmd *cobra.Command, args []string) {
			userId, tok := tools.Creds()
			md := tools.NewMeta(map[string]string{
				"authorization": tok,
				"ownerId":       userId,
			})
			ctx := tools.NewOutCtx(context.Background(), md)
			request := encodeDecode[nodes.Req, nodes.Res]
			if len(args) != 1 {
				panic(fmt.Errorf("arg must be 1"))
			}
			prop := args[0]
			req := nodes.Req{Nodes: []nodes.Node{{Data: prop}}}
			resp := &nodes.Res{}
			if err := request(ctx, g.client.List, req, resp); err != nil {
				panic(err)
			}
			for _, result := range resp.Nodes {
				fmt.Printf("\n *********************************** \n")
				out(result.Data)
				fmt.Printf("\n *********************************** \n")
			}
			slic := []string{"app list"}
			slic = append(slic, args...)
			behavior.LogCmd(strings.Join(slic, " "))
		}}
}

func (g *CmdSet) Link() *cobra.Command {
	return &cobra.Command{
		Use: "link",
		Run: func(cmd *cobra.Command, args []string) {
			userId, tok := tools.Creds()
			md := tools.NewMeta(map[string]string{
				"authorization": tok,
				"ownerId":       userId,
			})
			ctx := tools.NewOutCtx(context.Background(), md)
			serve := encodeDecode[nodes.Req, nodes.Res]
			req := nodes.Req{Nodes: []nodes.Node{}}
			resp := &nodes.Res{}
			if err := serve(ctx, g.client.Link, req, resp); err != nil {
				panic(err)
			}
			fmt.Println(resp)
		}}
}

func (g *CmdSet) Build() *cobra.Command {
	return &cobra.Command{
		Use: "build",
		Run: func(cmd *cobra.Command, args []string) {
			userId, tok := tools.Creds()
			md := tools.NewMeta(map[string]string{
				"authorization": tok,
				"ownerId":       userId,
			})
			ctx := tools.NewOutCtx(context.Background(), md)
			serve := encodeDecode[nodes.Req, nodes.Res]
			req := nodes.Req{Nodes: []nodes.Node{}}
			resp := &nodes.Res{}
			if err := serve(ctx, g.client.Build, req, resp); err != nil {
				panic(err)
			}
			fmt.Println(resp)
		}}
}

func (g *CmdSet) Deploy() *cobra.Command {
	return &cobra.Command{
		Use: "deploy",
		Run: func(cmd *cobra.Command, args []string) {
			userId, tok := tools.Creds()
			md := tools.NewMeta(map[string]string{
				"authorization": tok,
				"ownerId":       userId,
			})
			ctx := tools.NewOutCtx(context.Background(), md)
			serve := encodeDecode[nodes.Req, nodes.Res]
			req := nodes.Req{Nodes: []nodes.Node{}}
			resp := &nodes.Res{}
			if err := serve(ctx, g.client.Deploy, req, resp); err != nil {
				panic(err)
			}
			fmt.Println(resp)
		}}
}

func (g *CmdSet) Req() *cobra.Command {
	return &cobra.Command{
		Use: "req",
		Run: func(cmd *cobra.Command, args []string) {
			userId, tok := tools.Creds()
			md := tools.NewMeta(map[string]string{
				"authorization": tok,
				"ownerId":       userId,
			})
			ctx := tools.NewOutCtx(context.Background(), md)
			serve := encodeDecode[nodes.Req, nodes.Res]
			req := nodes.Req{Nodes: []nodes.Node{}}
			resp := &nodes.Res{}
			if err := serve(ctx, g.client.Req, req, resp); err != nil {
				panic(err)
			}
			fmt.Println(resp)
		}}
}

func encodeDecode[Req nodes.ReqReader, Res nodes.ResWriter](ctx context.Context,
	handler grpc.HandleFunc, r Req, w *Res) error {
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
