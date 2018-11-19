package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/lzol/KIEClient/kieresult"
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
	resp := kieresult.ServiceResponse{}
	resp.Result = kieresult.ExecutionResult{}
	json.Unmarshal(retMsg, &resp)
	//result := resp.Result
	//fmt.Println(result.GetValue("applyInfo"))
	tmp, err := resp.GetResults()
	result := tmp.(kieresult.ExecutionResult)
	fmt.Println(result.GetValue("com.qchery.harper.fact.ApplyInfo"), err)
	//fmt.Println(result.GetValue("value"))
	//fmt.Println(resp.Msg)
	//result := kieresult.ExecutionResult{}

}
