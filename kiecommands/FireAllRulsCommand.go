package kiecommands

import "github.com/lzol/KIEClient/enums"

var fireAllRulsCommand = enums.FIREALLRULES

type FireAllRulsCommand struct {
	Max           int64 `json:"max"`
	agendaFilter  string
	OutIdentifier string `json:"out-identifier"`
}

func (f FireAllRulsCommand) GetName() string {
	return fireAllRulsCommand
}

func NewFireAllRulsCommand(max int64) FireAllRulsCommand {
	fireAllRulsCommand := FireAllRulsCommand{}
	fireAllRulsCommand.Max = max
	fireAllRulsCommand.OutIdentifier = "out-identifier"
	return fireAllRulsCommand
}
