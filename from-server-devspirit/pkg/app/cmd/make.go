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

// makeCmd represents the make command
var makeCmd = &cobra.Command{
	Use:   "make",
	Short: "A brief description of your command",
	Long: `app make -rpc GoFuncToRpc -use myplugin -doc "makes an endpoint that connects to the go func"
    # creates a container and inserts boilerplate code for GoFuncToRpc.
	# make is different than 'turn' in that it makes boilerplate code for an endpoint that reaches
	# an unimplemented function named GoFuncToRpc; it stops here if there is no -use flag
	# -use myplugin makes the client for myplugin avaliable to this unimplemented function.  
    # the files that make up this function are returned to the client so the user can 
	# implement the function`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("make called")
	},
}

func init() {
	rootCmd.AddCommand(makeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// makeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// makeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
