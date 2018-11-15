package commands

type BatchExecutionCommand struct {
	Lookup   string   `json:"lookup"`
	commands Commands `json:"commands"`
}
