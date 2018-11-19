package kiecommands

type Command interface {
	GetName() string
}

type Object map[string]interface{}

type Commands struct {
	Commands []map[string]Command
}

func (b *Commands) AddCommand(command Command) {
	tmpMap := make(map[string]Command)
	tmpMap[command.GetName()] = command
	b.Commands = append(b.Commands, tmpMap)
}
