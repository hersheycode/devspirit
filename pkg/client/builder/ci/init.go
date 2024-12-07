package ci

import (
	"apppathway.com/pkg/builder/ci/net/grpc"
	"apppathway.com/pkg/client/user/behavior"
	"github.com/spf13/cobra"
	"os"
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
		cmd.Play(),
	}
}
