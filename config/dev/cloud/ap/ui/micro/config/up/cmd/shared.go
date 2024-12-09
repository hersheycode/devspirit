package cmd

import (
	"bytes"
	"fmt"
	expect "github.com/google/goexpect"
	"golang.org/x/crypto/ssh"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

var databases = map[string]database{"au": {dbType: "dgraph", compose: "/workspaces/devspirit/deployments/db/auth/docker-compose.db.yml", name: "auth_db_stack"}, "cd": {dbType: "dgraph", compose: "/workspaces/devspirit/deployments/db/builder/insecure/docker-compose.db.yml", name: "no_tls_builder_db_stack"}, "cd-tls": {dbType: "dgraph", compose: "/workspaces/devspirit/deployments/db/builder/docker-compose.db.yml", name: "builder_db_stack"}}
var services = map[string]service{"is": {cmd: "is/up", root: composeRoot}, "id": {cmd: "intentsysd/up", root: composeRoot}, "in": {cmd: "intentd/up", root: composeRoot}, "sch": {cmd: "schedulerd/up", root: composeRoot}, "sms": {cmd: "smsd/up", root: composeRoot}}
var composeRoot = "/workspaces/devspirit/config/dev/cloud/ap/ui/crudusrs/config/compose"
var wd string

type service struct {
	cmd     string
	root    string
	name    string
	compose string
	lang    string
}
type database struct {
	dbType  string
	compose string
	name    string
}
type remote struct {
	IP       string
	VM       string
	StartCue string
	StartCmd string
	StartMsg string
	EndMsg   string
	Root     string
	Batch    []expect.Batcher
}

func init() {
	var err error
	wd, err = os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
}
func tidy(container, root string) {
	rt := remote{StartMsg: "start", EndMsg: "end", Root: root}
	rt.Batch = []expect.Batcher{&expect.BExp{R: " "}, &expect.BSnd{S: fmt.Sprintf("cd %s && go mod tidy && echo 'done'\n", root)}, &expect.BExp{R: "done"}, &expect.BSnd{S: "\n"}}
	rt.execute()
	rt.Batch = []expect.Batcher{&expect.BExp{R: " "}, &expect.BSnd{S: fmt.Sprintf("sudo docker exec %s echo 'package %s' > %s/temp.go && rm %s/temp.go \n", container, container, root, root)}, &expect.BExp{R: "nate:"}, &expect.BSnd{S: "s42go@p*T1SG*p\n"}, &expect.BExp{R: " "}, &expect.BSnd{S: "\n"}}
	rt.execute()
}
func run(command, dir string) {
	cmdArgs := strings.Fields(command)
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	cmd.Dir = dir
	var stdBuffer bytes.Buffer
	mw := io.MultiWriter(os.Stdout, &stdBuffer)
	cmd.Stdout = mw
	cmd.Stderr = mw
	if err := cmd.Run(); err != nil {
		log.Fatalf("%s %v", command, err)
	}
}
func (rt remote) execute() {
	timeout := 10 * time.Minute
	user := "nate"
	host := "10.0.0.186" + ":22"
	cloudKey := "/workspaces/devspirit/secrets/ssh/id_apppathway"
	key, err := ioutil.ReadFile(cloudKey)
	if err != nil {
		log.Fatalf("unable to read private key: %v", err)
	}
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("unable to parse private key: %v", err)
	}
	config := &ssh.ClientConfig{User: user, Auth: []ssh.AuthMethod{ssh.PublicKeys(signer)}, HostKeyCallback: ssh.InsecureIgnoreHostKey()}
	sshClt, err := ssh.Dial("tcp", host, config)
	if err != nil {
		log.Fatalf("unable to connect: %v", err)
	}
	defer sshClt.Close()
	e, _, err := expect.SpawnSSH(sshClt, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer e.Close()
	res, err := e.ExpectBatch(rt.Batch, timeout)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("OUTPUT: %+v\n", res)
}
