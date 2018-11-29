package kiecommands

import "github.com/lzol/KIEClient/enums"

var setGlobalCommand = enums.SET_GLOBAL

type SetGlobalCommand struct {
	Identifier    string `json:"identifier,omitempty"`
	Object        Object `json:"object"`
	Out           bool   `json:"out"`
	OutIdentifier string `json:"out-identifier,omitempty"`
}

func (s SetGlobalCommand) GetName() string {
	return setGlobalCommand
}

func NewSetGlobalCommand(object Object, identifier, outIdentifier string, out bool) SetGlobalCommand {
	setGlobalCommand := SetGlobalCommand{}
	setGlobalCommand.Object = object
	setGlobalCommand.Identifier = identifier
	setGlobalCommand.Out = out
	setGlobalCommand.OutIdentifier = outIdentifier
	return setGlobalCommand
}
