package kiecommands

import "github.com/lzol/KIEClient/enums"

var insertCommandName = enums.INSERT_OBJECT

type InsertCommand struct {
	Object        Object `json:"object"`
	OutIdentifier string `json:"out-identifier,omitempty"`
	ReturnObject  bool   `json:"return-object,omitempty"`
	EntryPoint    string `json:"entry-point,omitempty"`
	Disconnected  bool   `json:"disconnected,omitempty"`
}

func (i InsertCommand) GetName() string {
	return insertCommandName
}

func NewInsertCommand(outIdentifier string, objcet Object) InsertCommand {
	insertCommand := InsertCommand{}
	insertCommand.Object = objcet
	insertCommand.Disconnected = false
	insertCommand.OutIdentifier = outIdentifier
	insertCommand.EntryPoint = "DEFAULT"
	return insertCommand
}
