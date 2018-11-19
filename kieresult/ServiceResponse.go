package kieresult

import (
	"encoding/json"
)

type ServiceResponse struct {
	ResponseType string      `json:"type"`
	Msg          string      `json:"msg"`
	Result       interface{} `json:"result"`
}

func (s *ServiceResponse) GetResults(resultKey string) (interface{}, error) {
	r := s.Result.(map[string]interface{})
	if _, ok := r[resultKey]; ok {
		b, _ := json.Marshal(r[resultKey])
		str := string(b)
		e := ExecutionResult{}
		err := json.Unmarshal([]byte(str), &e)
		return e, err
	} else {
		return s.Result, nil
	}

}

type ExecutionResult struct {
	Results []map[string]interface{} `json:"results"`
	Facts   []Facts                  `json:"facts"`
}

func (e *ExecutionResult) GetValue(key string, i interface{}) error {
	pojoMap := e.Results[0]["value"].(map[string]interface{})
	b, _ := json.Marshal(pojoMap[key])
	err := json.Unmarshal(b, &i)
	return err
}

type Facts struct {
	Key   string `json:"key"`
	Value map[string]interface{}
}
