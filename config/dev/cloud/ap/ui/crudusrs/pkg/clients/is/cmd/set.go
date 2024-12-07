/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"fmt"
	"os"

	"apppathway.com/examples/prodapi/pkg/orgs/intentsys/api/intentsyspb"
	"apppathway.com/examples/prodapi/pkg/orgs/intentsys/net/grpc"
	"apppathway.com/pkg/net/grpc/tools"
	"github.com/spf13/cobra"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		caFilePath := os.Getenv("CA_FILE")
		fmt.Println(os.Getenv("INTENTSYSD_ADDRESS"))
		client := grpc.NewIntentSysClient(os.Getenv("INTENTSYSD_ADDRESS"), caFilePath)
		fmt.Println(client)
		// userId, tok := tools.Creds()
		md := tools.NewMeta(map[string]string{
			// "authorization": tok,
			// "ownerId":       userId,
		})
		ctx := tools.NewOutCtx(context.Background(), md)
		res, err := client.SetIntent(ctx, &intentsyspb.SetIntentRequest{
			Schedule: &intentsyspb.Schedule{
				Time: "now",
			},
			Sms: &intentsyspb.SMSInfo{
				Recipient: &intentsyspb.Recipient{
					PhoneNum: "5085213315",
					Email:    "datadrivenpath@gmail.com",
				},
				Msg: &intentsyspb.Message{
					Body: "my message",
				},
			},
			Intent: &intentsyspb.Intent{
				Name: "get up at 6am",
			},
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(res.Status)
	},
}

func init() {
	rootCmd.AddCommand(setCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
