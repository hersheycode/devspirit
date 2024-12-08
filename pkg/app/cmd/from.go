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

	"github.com/spf13/cobra"
)

// fromCmd represents the from command
var fromCmd = &cobra.Command{
	Use:   "from",
	Short: "A brief description of your command",
	Long: `app from "golang:latest" -copy ./template:/app -tag default -save
	# sends files over network to srcfiled 
	# srcfiled creates a new container and adds the files sent by the app cli
	# saves container in db, tagged as default
	
	app from default -require myplugin
	# attaches client for myplugin to default`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("from called")
	},
}

func init() {
	rootCmd.AddCommand(fromCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fromCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fromCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
