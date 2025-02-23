package sessions

import (
	"fmt"
	"github.com/famiphoto/famiphoto/api/utils/cast"
)

func getInt64(values map[any]any, key string, defaultValue int64) int64 {
	val, ok := values[key]
	if !ok {
		return defaultValue
	}
	dst, err := cast.ToInt64(val)
	if err != nil {
		fmt.Println("Failed to cast session data. key: ", key, err)
		return defaultValue
	}
	return dst
}
