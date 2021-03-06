package serviceclient

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lzol/KIEClient/kiecommands"
	"github.com/lzol/KIEClient/kieresult"
	"io/ioutil"
	"net/http"
)

type RuleServiceClient struct {
	Url         string
	UserName    string
	Password    string
	ContainerId string
}

func (r *RuleServiceClient) ExecWithResponse(command kiecommands.Command) (response kieresult.ServiceResponse, err error) {
	response = kieresult.ServiceResponse{}
	if r.UserName == "" || r.Password == "" || r.ContainerId == "" || r.Url == "" {
		err = errors.New("Url、UserName、Password、ContainerId某项为空")
		return response, err
	}
	url := r.Url + "/containers/instances/" + r.ContainerId
	str := r.UserName + ":" + r.Password
	basic := base64.StdEncoding.EncodeToString([]byte(str))
	payload, err := json.Marshal(command)
	if err != nil {
		return response, err
	}
	fmt.Println(string(payload))
	reader := bytes.NewReader(payload)
	request, _ := http.NewRequest("POST", url, reader)
	request.Header.Set("Authorization", "basic "+basic)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-KIE-ContentType", "json")
	client := &http.Client{}
	postResp, err := client.Do(request)
	if err != nil {
		return response, err
	}
	defer postResp.Body.Close()

	retMsg, err := ioutil.ReadAll(postResp.Body)
	if err != nil {
		return response, err
	}
	json.Unmarshal(retMsg, &response)
	return response, nil
}
