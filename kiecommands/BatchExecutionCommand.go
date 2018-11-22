package kiecommands

import "github.com/lzol/KIEClient/enums"

var batchExecutionCommandName = enums.BATCH_EXECUTION

type BatchExecutionCommand struct {
	Lookup   string               `json:"lookup"`
	Commands []map[string]Command `json:"commands"`
}

func (b BatchExecutionCommand) GetName() string {
	return batchExecutionCommandName
}

func NewBatchExecutionCommand(kieSession string) BatchExecutionCommand {
	batchExecutionCommand := BatchExecutionCommand{}
	batchExecutionCommand.Lookup = kieSession
	return batchExecutionCommand
}

func (b *BatchExecutionCommand) AddCommand(command Command) {
	tmpMap := make(map[string]Command)
	tmpMap[command.GetName()] = command
	b.Commands = append(b.Commands, tmpMap)
}
