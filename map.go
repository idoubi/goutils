package goutils

import (
	"strings"
)

// StrMap: map with string value
type StrMap = map[string]string

// IntMap: map with int value
type IntMap = map[string]int

// Int64Map: map with int64 value
type Int64Map = map[string]int64

// BoolMap: map with bool value
type BoolMap = map[string]bool

// MixMap: map with mix value
type MixMap = map[string]interface{}

// MGetString: get string from map
func MGetStr(m MixMap, key string) string {
	if v, ok := m[key]; ok {
		if str, ok := v.(string); ok {
			return str
		}
	}

	return ""
}

// MGetInt: get int from map
func MGetInt(m MixMap, key string) int {
	if v, ok := m[key]; ok {
		if i, ok := v.(int); ok {
			return i
		}
	}

	return 0
}

// MGetInt64: get int64 from map
func MGetInt64(m MixMap, key string) int64 {
	if v, ok := m[key]; ok {
		if i, ok := v.(int64); ok {
			return i
		}
	}

	return 0
}

// IsInMap: is str key in map
func IsInMap(m MixMap, key string) bool {
	if _, ok := m[key]; ok {
		return true
	}

	return false
}

// Multiline2Map: transfer multi line content to mixmap
func Multiline2Map(str string) MixMap {
	m := MixMap{}

	lines := strings.Split(str, "\r\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		line := strings.TrimSpace(line)
		arr := strings.SplitN(line, ":", 2)
		if len(arr) != 2 {
			continue
		}
		m[strings.TrimSpace(arr[0])] = strings.TrimSpace(arr[1])
	}

	return m
}
