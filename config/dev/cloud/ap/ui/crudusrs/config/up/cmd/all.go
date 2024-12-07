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

	"strings"

	"github.com/spf13/cobra"
)

// allCmd represents the all command
var allCmd = &cobra.Command{
	Use:   "all",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		//Deploy dgraph databases with docker swarm
		// command := "sudo docker stack rm auth_db_stack"
		// run(command, wd)
		// command = "sudo docker stack rm app_api_generator_db_stack"
		// run(command, wd)
		fmt.Println("Deploy databases? ")
		if ask() {
			command := "sudo docker stack deploy -c /home/nate/code/app-pathway/pkg/user/auth/deployments/docker-compose.db.yml auth_db_stack"
			run(command, wd)
			command = "sudo docker stack deploy -c /home/nate/code/app-pathway/pkg/cert/deployments/docker-compose.db.yml app_api_generator_db_stack"
			run(command, wd)
		}
		fmt.Println("Bring up services? ")
		if ask() {
			dockerignore := ""
			//fmt.Println("Build all services?")
			//buildAll := ask()
			for _, service := range services {
				dockerignore = service.name
				// command := fmt.Sprintf("sudo docker rm /%s", service)
				// run(command, wd)
				command := "sudo rm .dockerignore"
				fmt.Println(command)
				run(command, wd)
				command = fmt.Sprintf("sudo cp /home/nate/code/app-pathway/config/dev/dockerignore/%s .dockerignore", dockerignore)
				run(command, wd)

				if service.name == "nginx" {
					command = fmt.Sprintf("sudo docker-compose -f /home/nate/code/app-pathway/config/dev/dev/compose/%s up -d ", service.compose)
					run(command, wd)
					continue
				}
				command = fmt.Sprintf("sudo docker-compose -f /home/nate/code/app-pathway/config/dev/dev/compose/%s up -d %s", service.compose, service.name)
				run(command, wd)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(allCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// allCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// allCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func ask() bool {
	var s string

	fmt.Printf("(y/N): ")
	_, err := fmt.Scan(&s)
	if err != nil {
		panic(err)
	}

	s = strings.TrimSpace(s)
	s = strings.ToLower(s)

	if s == "y" || s == "yes" {
		return true
	}
	return false
}
