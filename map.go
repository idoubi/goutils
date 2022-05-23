package goutils

// MGetString: get string from map
func MGetStr(m map[string]interface{}, key string) string {
	if v, ok := m[key]; ok {
		if str, ok := v.(string); ok {
			return str
		}
	}

	return ""
}

// MGetInt: get int64 from map
func MGetInt(m map[string]interface{}, key string) int64 {
	if v, ok := m[key]; ok {
		if i, ok := v.(int64); ok {
			return i
		}
	}

	return 0
}
