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

	insertCommand := kiecommands.NewInsertCommand("com.qchery.harper.fact.ApplyInfo", "applyInfo", applyInfo)

	fireAllRulsCommand := kiecommands.NewFireAllRulsCommand(-1)

	commands := kiecommands.Commands{}
	commands.AddCommand(insertCommand)
	commands.AddCommand(fireAllRulsCommand)
	batchExecutionCommand := kiecommands.NewBatchExecutionCommand("ksessionId", commands)

	ruleServiceClient := serviceclient.RuleServiceClient{}
	ruleServiceClient.Url = "http://localhost:8180/kie-server/services/rest/server"
	ruleServiceClient.UserName = "admin"
	ruleServiceClient.Password = "admin"
	ruleServiceClient.ContainerId = "harper_1.0.0"

	resp, err := ruleServiceClient.ExecWithResponse(batchExecutionCommand)
	if err == nil {
		if resp.ResponseType == enums.RESPONSE_SUCCESS {
			result := kieresult.ExecutionResult{}
			err = resp.GetResults(enums.EXECUTION_RESULTS, &result)
			applyInfo := fact.ApplyInfo{}
			err = result.GetValue("com.qchery.harper.fact.ApplyInfo", &applyInfo)
			fmt.Println(applyInfo.Name, applyInfo.Age)
		}
	}

}
