package kiecommands

import "fmt"

type Command interface {
	GetName() string
}

type Object map[string]interface{}

func NewObject(className string, obj interface{}) Object {
	object := make(map[string]interface{})
	fmt.Println(object)
	object[className] = obj

	return object
}

type Commands struct {
	Commands []map[string]Command
}

//func (b *Commands) AddCommand(command Command) {
//	tmpMap := make(map[string]Command)
//	tmpMap[command.GetName()] = command
//	b.Commands = append(b.Commands, tmpMap)
//}
