package commands

type InsertCommand struct {
	Object        map[string]interface{} `json:"object"`
	OutIdentifier string                 `json:"out-identifier"`
	ReturnObject  bool                   `json:"return-object"`
	EntryPoint    string                 `json:"entry-point"`
	Disconnected  bool                   `json:"disconnected"`
}
