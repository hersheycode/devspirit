package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

func run(command, dir string) {
	cmdArgs := strings.Fields(command)
	cmd := exec.Command(cmdArgs[0],
		cmdArgs[1:]...)
	cmd.Dir = dir
	var stdBuffer bytes.Buffer
	mw := io.MultiWriter(os.Stdout, &stdBuffer)
	cmd.Stdout = mw
	cmd.Stderr = mw
	if err := cmd.Run(); err != nil {
		log.Fatalf("%s %v", command, err)
	}
}

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	//Deploy dgraph databases with docker swarm
	// command := "sudo docker stack rm auth_db_stack"
	// run(command, wd)
	// command = "sudo docker stack rm app_api_generator_db_stack"
	// run(command, wd)
	fmt.Println("Deploy databases? ")
	if ask() {
		command := "sudo docker stack deploy -c ../../pkg/user/auth/deployments/docker-compose.db.yml auth_db_stack"
		run(command, wd)
		command = "sudo docker stack deploy -c ../../pkg/cert/deployments/docker-compose.db.yml app_api_generator_db_stack"
		run(command, wd)
	}

	services := map[string]string{
		"nodes":     "builder.yml",
		"cplugind":  "builder.yml",
		"cid":       "builder.yml",
		"cdd":       "builder.yml",
		"behaviord": "user.yml",
		"code":      "third_party.yml",

		// "hoppscotch":    "docker-compose.yml",
		// "gitea":         "docker-compose.yml",
		"auth":        "user.yml",
		"auth_db_api": "db.yml",
		// "api_generator": "docker-compose.yml",
		"ap": "ctl.yml",
		// "ui":            "docker-compose.yml",
		// "envoy":         "docker-compose.yml",
		"nginx": "cert.yml",
	}

	dockerignore := ""
	//fmt.Println("Build all services?")
	//buildAll := ask()
	for service, file := range services {
		dockerignore = service
		// command := fmt.Sprintf("sudo docker rm /%s", service)
		// run(command, wd)
		command := "sudo rm .dockerignore"
		fmt.Println(command)
		run(command, wd)
		command = fmt.Sprintf("sudo cp ./scripts/build/dev/dockerignore/%s .dockerignore", dockerignore)
		run(command, wd)

		if service == "nginx" {
			command = fmt.Sprintf("sudo docker-compose -f ./config/dev/compose/%s up -d ", file)
			run(command, wd)
			continue
		}
		command = fmt.Sprintf("sudo docker-compose -f ./config/dev/compose/%s up -d %s", file, service)
		run(command, wd)
	}
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
