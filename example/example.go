package main

import (
	"encoding/json"
	"fmt"
	"github.com/lzol/KIEClient/example/pojo"
	"github.com/lzol/KIEClient/kiecommands"
)

func main() {
	address := pojo.Address{}
	address.Province = "aa"

	applyInfo := pojo.ApplyInfo{}
	applyInfo.Age = 19
	applyInfo.FamilyAddress = address

	insertCommand := kiecommands.NewInsertCommand("com.qchery.harper.fact.ApplyInfo", "applyInfo", applyInfo)

	fireAllRulsCommand := kiecommands.NewFireAllRulsCommand(-1)

	commands := kiecommands.Commands{}
	commands.AddCommand(insertCommand)
	commands.AddCommand(fireAllRulsCommand)
	batchExecutionCommand := kiecommands.NewBatchExecutionCommand("ksessionId", commands)

	fmt.Println(batchExecutionCommand.Commands)
	b, _ := json.Marshal(batchExecutionCommand)
	fmt.Println(string(b))

}
