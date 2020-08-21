package jsonutils

import (
	"errors"
	"fmt"
)

func Flatten(json interface{}) interface{} {
	res, _ := flattenHelper(json)
	return res
}

func flattenHelper(root interface{}) (interface{}, error) {
	fmt.Println("helper: ", root)
	result := map[string]interface{}{}
	rootMap, ok := root.(map[string]interface{})
	if !ok {
		return nil, errors.New("input is not of abstract json type (map[string]interface{})")
	}
	for k, v := range rootMap {
		if v == nil {
			result[k] = v
		} else if _, ok := v.(bool); ok {
			result[k] = v
		} else if _, ok := v.(float64); ok {
			result[k] = v
		} else if _, ok := v.(string); ok {
			result[k] = v
		} else {
			flattenedChild, err := flattenHelper(v)
			if err != nil {
				return nil, err
			}
			flattenedChildMap := flattenedChild.(map[string]interface{})
			for childK, childV := range flattenedChildMap {
				result[k+"."+childK] = childV
			}
		}
	}
	return result, nil
}
