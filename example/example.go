package main

import (
	"fmt"
	"github.com/lzol/KIEClient/enums"
	"github.com/lzol/KIEClient/example/fact"
	"github.com/lzol/KIEClient/kiecommands"
	"github.com/lzol/KIEClient/kieresult"
	"github.com/lzol/KIEClient/serviceclient"
)

func main() {
	address := fact.Address{}
	address.Province = "aa"

	applyInfo := fact.ApplyInfo{}
	applyInfo.Age = 19
	applyInfo.FamilyAddress = address

	insertObject := kiecommands.NewObject("com.qchery.harper.fact.ApplyInfo", applyInfo)
	insertCommand := kiecommands.NewInsertCommand("applyInfo", insertObject)
	fireAllRulsCommand := kiecommands.NewFireAllRulsCommand(-1)
	fireAllRulsCommand.AgendaFilter = "aaa"
	fireAllRulsCommand.OutIdentifier = "bbb"

	batchExecutionCommand := kiecommands.NewBatchExecutionCommand("ksessionId")
	batchExecutionCommand.AddCommand(insertCommand)
	batchExecutionCommand.AddCommand(fireAllRulsCommand)
	//objAddr := kiecommands.NewObject("aaa",address)
	//fmt.Println(&objAddr)
	//setGlobalCommand := kiecommands.NewSetGlobalCommand(objAddr,"addr","addr",true)
	//batchExecutionCommand.AddCommand(setGlobalCommand)

	ruleServiceClient := serviceclient.RuleServiceClient{}
	ruleServiceClient.Url = "http://localhost:8180/kie-server/services/rest/server"
	ruleServiceClient.UserName = "admin"
	ruleServiceClient.Password = "admin"
	ruleServiceClient.ContainerId = "harper_1.0.0"

	resp, err := ruleServiceClient.ExecWithResponse(batchExecutionCommand)
	fmt.Println(resp)
	if err == nil {
		if resp.ResponseType == enums.RESPONSE_SUCCESS {
			result := kieresult.ExecutionResult{}
			err = resp.GetResults(enums.EXECUTION_RESULTS, &result)
			fmt.Println(result)
			applyInfo := fact.ApplyInfo{}
			err = result.GetValue("com.qchery.harper.fact.ApplyInfo", &applyInfo)
			fmt.Println(applyInfo.Name, applyInfo.Age)
		}
	}

}
