package cobra

import (
	"apppathway.com/pkg/builder/base/net/grpc"
	"github.com/spf13/cobra"
)

type CmdSet struct {
	client grpc.Client
}

func NewCmdSet(connStr, caFilePath string) []*cobra.Command {
	cmd := &CmdSet{
		client: *grpc.NewClient(connStr, caFilePath),
	}

	return []*cobra.Command{
		cmd.Create(),
	}
}
