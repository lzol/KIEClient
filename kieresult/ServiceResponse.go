package kieresult

type ServiceResponse struct {
	ResponseType string      `json:"type"`
	Msg          string      `json:"msg"`
	Result       interface{} `json:"result"`
}

type ExecutionResult struct {
	Results map[string]interface{} `json:"results"`
	Facts   map[string]interface{} `json:"facts"`
}
