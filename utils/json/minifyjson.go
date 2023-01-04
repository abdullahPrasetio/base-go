package json

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
)

func MinifyJson(r interface{}) map[string]interface{} {
	js, _ := jsoniter.Marshal(r)
	var m map[string]interface{}
	_ = jsoniter.Unmarshal(js, &m)
	for k, v := range m {
		s := fmt.Sprintf("%v", v)
		if len(s) > 10000 {
			delete(m, k)
		}
	}
	return m
}
