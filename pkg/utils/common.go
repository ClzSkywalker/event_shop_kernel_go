package utils

import (
	"strconv"
	"strings"
)

func ToString(v interface{}) (result string) {
	switch value := v.(type) {
	case string:
		return value
	case int:
		return strconv.Itoa(value)
	case float64:
		return strconv.FormatFloat(value, 'e', -1, 64)
	case []string:
		return strings.Join(value, ",")
	}
	return
}
