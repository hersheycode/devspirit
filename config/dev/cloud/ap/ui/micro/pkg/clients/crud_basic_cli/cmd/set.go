package cmd

import (
	"apppathway.com/examples/prodapi/pkg/orgs/intentsys/api/intentsyspb"
	"apppathway.com/examples/prodapi/pkg/orgs/intentsys/net/grpc"
	"apppathway.com/pkg/net/grpc/tools"
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var setCmd = &cobra.Command{Use: "set", Short: "A brief description of your command", Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`, Run: func(cmd *cobra.Command, args []string) {
	caFilePath := os.Getenv("CA_FILE")
	fmt.Println(os.Getenv("INTENTSYSD_ADDRESS"))
	client := grpc.NewIntentSysClient(os.Getenv("INTENTSYSD_ADDRESS"), caFilePath)
	fmt.Println(client)
	md := tools.NewMeta(map[string]string{})
	ctx := tools.NewOutCtx(context.Background(), md)
	res, err := client.SetIntent(ctx, &intentsyspb.SetIntentRequest{Schedule: &intentsyspb.Schedule{Time: "now"}, Sms: &intentsyspb.SMSInfo{Recipient: &intentsyspb.Recipient{PhoneNum: "5085213315", Email: "datadrivenpath@gmail.com"}, Msg: &intentsyspb.Message{Body: "my message"}}, Intent: &intentsyspb.Intent{Name: "get up at 6am"}})
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Status)
}}

func init() {
	rootCmd.AddCommand(setCmd)
}
