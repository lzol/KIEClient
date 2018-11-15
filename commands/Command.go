package commands

type Command struct {
	CommandType enums.EnumsCommand
}

type Commands struct {
	commands []Command `json:"commands"`
}
