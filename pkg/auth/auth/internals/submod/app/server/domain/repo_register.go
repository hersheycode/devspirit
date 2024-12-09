package domain

import (
	dt "apppathway.com/pkg/user/auth/internals/project/pkg/types"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

const tok = "token 1bf17aae6745d4545e7bd1cc18decce9fd4908bd"

func createRepoAccount(usr dt.User) {
	username := strings.Split(usr.Email, "@")[0]
	postBody, _ := json.Marshal(map[string]interface{}{"email": usr.Email, "full_name": username, "login_name": username, "must_change_password": false, "password": usr.Password, "send_notify": false, "source_id": 0, "username": username, "visibility": "private"})
	body := bytes.NewBuffer(postBody)
	err := call("https://apppathway.com/pkg/api/v1/admin/users", "POST", body)
	if err != nil {
		log.Printf("err: %v \n", err)
	}
}
func call(url, method string, body io.Reader) error {
	client := &http.Client{Timeout: time.Second * 10}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return fmt.Errorf("Got error %s", err.Error())
	}
	req.Header.Set("user-agent", "apclivm")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", tok)
	response, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Got error %s", err.Error())
	}
	defer response.Body.Close()
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	log.Printf("response: %v \n", bodyString)
	return nil
}
