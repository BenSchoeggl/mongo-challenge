package jsonutils

// Flatten takes a map[string]interface{} representation of a json object and returns
// a flattened version, meaning that the keys of the parents of all the leaves in the
// object are joined with a ".". For example the json {"a":{"b":"c"}} will become
// {"a.b":"c"}
func Flatten(json map[string]interface{}) interface{} {
	result := map[string]interface{}{}
	for k, v := range json {
		// Our base case is when we have reached a key that has a value that is not an object
		if v == nil {
			result[k] = v
		} else if _, ok := v.(bool); ok {
			result[k] = v
		} else if _, ok := v.(float64); ok {
			result[k] = v
		} else if _, ok := v.(string); ok {
			result[k] = v
		} else {
			// At this point we know that v has to be a map[string]interface{} because we have
			// tested for all the possible types (except arrays which are explicitly forbidden)
			// so it's safe to cast
			vMap := v.(map[string]interface{})
			flattenedChild := Flatten(vMap)
			flattenedChildMap := flattenedChild.(map[string]interface{})
			for childK, childV := range flattenedChildMap {
				result[k+"."+childK] = childV
			}
		}
	}
	return result
}
