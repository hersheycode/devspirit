package cmd

import (
	"fmt"
	"github.com/google/goterm/term"
	"github.com/spf13/cobra"
)

var tidyCmd = &cobra.Command{Use: "tidy", Short: "A brief description of your command", Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`, Run: func(cmd *cobra.Command, args []string) {
	if len(args) != 1 || args[0] == "" {
		fmt.Println(term.Red("ONE service abbr is required"))
	}
	abbr := args[0]
	if _, ok := services[abbr]; !ok {
		fmt.Println(term.Red("abbr no found"))
		return
	}
	srv := services[abbr]
	if srv.lang != "go" {
		fmt.Println(term.Red("only go lang"))
		return
	}
	tidy(srv.name, srv.root)
}}

func init() {
	rootCmd.AddCommand(tidyCmd)
}
