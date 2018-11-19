package kieresult

import (
	"encoding/json"
)

type ServiceResponse struct {
	ResponseType string      `json:"type"`
	Msg          string      `json:"msg"`
	Result       interface{} `json:"result"`
}

func (s *ServiceResponse) GetResults() (interface{}, error) {
	switch s.Result.(type) {
	case map[string]interface{}:
		r := s.Result.(map[string]interface{})
		if _, ok := r["execution-results"]; ok {
			b, _ := json.Marshal(r["execution-results"])
			str := string(b)
			e := ExecutionResult{}
			err := json.Unmarshal([]byte(str), &e)
			return e, err
		} else {
			return s.Result, nil
		}

	default:
		return s.Result, nil
	}
}

type ExecutionResult struct {
	Results []map[string]interface{} `json:"results"`
	Facts   []Facts                  `json:"facts"`
}

func (e *ExecutionResult) GetValue(key, structName string) interface{} {
	pojoMap := e.Results[0]["value"].(map[string]interface{})

	return pojoMap[key]
}

//func (e *ExecutionResult)GetFactHandle(key string)interface{}{
//	return e.Facts[key]
//}

type Facts struct {
	Key   string `json:"key"`
	Value map[string]interface{}
}
