package rnjson

import (
	"encoding/json"
	"strings"
)

func Get(obj interface{}, path string) interface{} {
	if path == "" {
		return obj
	}
	spt := strings.Split(path, ".")
	temp, ok := obj.(map[string]interface{})
	if !ok {
		return nil
	}
	n := len(spt)
	for _, v := range spt[:n-1] {
		tmp, ok := temp[v].(map[string]interface{})
		if !ok {
			return nil
		}
		temp = tmp
	}
	return temp[spt[n-1]]
}

func Unmarshal(dat string) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := json.Unmarshal([]byte(dat), &result)
	return result, err
}
