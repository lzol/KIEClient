package kieresult

import (
	"encoding/json"
	"github.com/kataras/iris/core/errors"
)

type ServiceResponse struct {
	ResponseType string      `json:"type"`
	Msg          string      `json:"msg"`
	Result       interface{} `json:"result"`
}

func (s *ServiceResponse) GetResults(resultKey string, i interface{}) error {
	r := s.Result.(map[string]interface{})
	if _, ok := r[resultKey]; ok {
		b, _ := json.Marshal(r[resultKey])
		str := string(b)
		err := json.Unmarshal([]byte(str), &i)
		return err
	} else {
		return errors.New("没有key为" + resultKey + "的数据")
	}

}

type ExecutionResult struct {
	Results []map[string]interface{} `json:"results"`
	Facts   []Facts                  `json:"facts"`
}

func (e *ExecutionResult) GetValue(key string, i interface{}) error {
	mapLen := len(e.Results)
	//如果FireAllRulsCommand命令OutIdentifier不为空，则返回值中会添加一个调用规则数量的Value值，所以统一取最后一个Results的value值
	pojoMap := e.Results[mapLen-1]["value"].(map[string]interface{})
	b, _ := json.Marshal(pojoMap[key])
	err := json.Unmarshal(b, &i)
	return err
}

type Facts struct {
	Key   string `json:"key"`
	Value map[string]interface{}
}
