package commands

type Command interface {
	GetName() string
}

type Object struct {
	m map[string]interface{}
}

type Commands struct {
	commands []Command `json:"commands"`
}
