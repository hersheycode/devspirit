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
	"fmt"

	"github.com/google/goterm/term"
	"github.com/spf13/cobra"
)

// dbCmd represents the db command
var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(term.Green("command"))
		abbr, err := cmd.Flags().GetString("n")
		if err != nil {
			panic(err)
		}
		shouldBuild, err := cmd.Flags().GetBool("b")
		if err != nil {
			panic(err)
		}
		if abbr == "" {
			fmt.Println(term.Red("abbr for service is required"))
			return
		}

		if _, ok := databases[abbr]; !ok {
			fmt.Println(term.Red("abbr no found"))
			return
		}

		srv := databases[abbr]

		if shouldBuild {
			command := fmt.Sprintf("sudo docker stack rm %s", srv.name)
			fmt.Println(term.Green(command))
			run(command, wd)
		}

		if databases[abbr].dbType == "dgraph" {
			command := fmt.Sprintf("sudo docker stack deploy -c %s %s", srv.compose, srv.name)
			fmt.Println(term.Green(command))
			run(command, wd)
		}

	},
}

func init() {
	dbCmd.Flags().Bool("b", false, "build")
	dbCmd.Flags().String("n", "", "name")
	rootCmd.AddCommand(dbCmd)
}
