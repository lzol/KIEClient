package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	str := "admin:admin"
	encodeStr := base64.StdEncoding.EncodeToString([]byte(str))
	fmt.Println(encodeStr)
	url := "http://localhost:8180/kie-server/services/rest/server/containers/instances/harper_1.0.0"
	jsonStr := "{	\"lookup\" : \"ksessionId\",	\"commands\" : [ {	\"insert\" : {	\"object\" : {\"com.qchery.harper.fact.ApplyInfo\":{	\"name\" : null,	\"age\" : 19,	\"familyAddress\" : {	\"city\" : null,	\"detail\" : null,	\"distract\" : null,	\"province\" : \"aa\"}}},\"out-identifier\" : \"applyInfo\",\"return-object\" : true,\"entry-point\" : \"DEFAULT\",\"disconnected\" : false}}, {\"fire-all-rules\" : {\"max\" : -1,\"out-identifier\" : null}} ]}"
	reader := bytes.NewReader([]byte(jsonStr))
	request, _ := http.NewRequest("POST", url, reader)
	request.Header.Set("Authorization", "basic "+encodeStr)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-KIE-ContentType", "json")
	client := &http.Client{}
	response, _ := client.Do(request)
	defer response.Body.Close()
	retMsg, _ := ioutil.ReadAll(response.Body)
	resp := make(map[string]map[string]interface{})
	json.Unmarshal(retMsg, &resp)
	strMsg := string(retMsg)
	fmt.Println(strMsg)

	result := resp["result"]
	fmt.Println(result["execution-results"])
}
