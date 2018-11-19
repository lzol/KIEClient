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

func NewBatchExecutionCommand(kieSession string, commands Commands) BatchExecutionCommand {
	batchExecutionCommand := BatchExecutionCommand{}
	batchExecutionCommand.Lookup = kieSession
	batchExecutionCommand.Commands = commands.Commands
	return batchExecutionCommand
}
