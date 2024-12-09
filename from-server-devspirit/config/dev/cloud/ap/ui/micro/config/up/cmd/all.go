package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var allCmd = &cobra.Command{Use: "all", Short: "A brief description of your command", Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`, Run: func(cmd *cobra.Command, args []string) {
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
		for _, service := range services {
			dockerignore = service.name
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
}}

func init() {
	rootCmd.AddCommand(allCmd)
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
