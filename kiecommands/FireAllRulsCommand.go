package kiecommands

import "github.com/lzol/KIEClient/enums"

var fireAllRulsCommand = enums.FIREALLRULES

type FireAllRulsCommand struct {
	Max           int64  `json:"max"`
	AgendaFilter  string `json:"agenda-filter,omitempty"`
	OutIdentifier string `json:"out-identifier,omitempty"` //Add the number of rules activations fired on the execution results
}

func (f FireAllRulsCommand) GetName() string {
	return fireAllRulsCommand
}

func NewFireAllRulsCommand(max int64) FireAllRulsCommand {
	fireAllRulsCommand := FireAllRulsCommand{}
	fireAllRulsCommand.Max = max
	//fireAllRulsCommand.OutIdentifier = "out-identifier"
	return fireAllRulsCommand
}
