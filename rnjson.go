package rnjson

import (
	"encoding/json"
	"strings"
)

/**
 * Rn's JSON Type Model
 */
type Rnjson map[string]interface{}

func Get(obj Rnjson, path string) (interface{}, bool) {
	spt := strings.Split(path, ".")
	temp := obj
	n := len(spt)
	for _, v := range spt[:n-1] {
		switch temp[v].(type) {
		case map[string]interface{}:
			temp = temp[v].(map[string]interface{})
		default:
			return nil, false
		}
	}
	v, ok := temp[spt[n-1]]
	return v, ok
}

func Unmarshal(dat string) (Rnjson, error) {
	var result Rnjson
	err := json.Unmarshal([]byte(dat), &result)
	return result, err
}
