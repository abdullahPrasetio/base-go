package json

/********************************************************************************
* Temancode Example Minify Package                                              *
*                                                                               *
* Version: 1.0.0                                                                *
* Date:    2023-01-05                                                           *
* Author:  Waluyo Ade Prasetio                                                  *
* Github:  https://github.com/abdullahPrasetio                                  *
********************************************************************************/

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
