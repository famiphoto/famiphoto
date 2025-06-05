package cast

import (
	"fmt"
	"reflect"
	"strconv"
)

func BoolToInt8(b bool) int8 {
	if b {
		return 1
	}
	return 0
}

func Array[T any, V any](list []*T, castFunc func(*T) *V) []*V {
	dst := make([]*V, len(list))
	for i, v := range list {
		dst[i] = castFunc(v)
	}
	return dst
}

func ArrayValues[T any, V any](list []*T, castFunc func(*T) V) []V {
	dst := make([]V, len(list))
	for i, v := range list {
		dst[i] = castFunc(v)
	}
	return dst
}

func IntToBool[T ~int8 | ~int | ~int64](v T) bool {
	return v > 0
}

func IntPtrToInt64Ptr(intVal *int) *int64 {
	if intVal == nil {
		return nil
	}
	val := *intVal
	dst := int64(val)
	return &dst
}

func Ptr[T any](val T) *T {
	return &val
}

func PtrToVal[T any](ptr *T, defaultValue T) T {
	if ptr == nil {
		return defaultValue
	}
	return *ptr
}

func ToInt64(val any) (int64, error) {
	if dst, ok := val.(int64); ok {
		return dst, nil
	}
	if dst, ok := val.(float64); ok {
		return int64(dst), nil
	}
	if dst, ok := val.(uint32); ok {
		return int64(dst), nil
	}
	if dst, ok := val.(uint16); ok {
		return int64(dst), nil
	}
	if dst, ok := val.(uint8); ok {
		return int64(dst), nil
	}
	if dst, ok := val.(uint); ok {
		return int64(dst), nil
	}
	if dst, ok := val.(int); ok {
		return int64(dst), nil
	}
	if dst, ok := val.(uint64); ok {
		return int64(dst), nil
	}
	if dst, ok := val.(int8); ok {
		return int64(dst), nil
	}
	if dst, ok := val.(int16); ok {
		return int64(dst), nil
	}
	if dst, ok := val.(int32); ok {
		return int64(dst), nil
	}
	if dst, ok := val.(float32); ok {
		return int64(dst), nil
	}
	if dst, ok := val.(string); ok {
		if dst2, err := strconv.Atoi(dst); err == nil {
			return int64(dst2), nil
		}
	}
	if dst, ok := val.([]byte); ok {
		if dst2, err := strconv.Atoi(string(dst)); err == nil {
			return int64(dst2), nil
		}
	}
	return 0, fmt.Errorf("failed to ToInt64, val is %s %#v", reflect.TypeOf(val).String(), val)
}

func ToString(val any) (string, error) {
	if dst, ok := val.([]uint8); ok {
		return string(dst), nil
	}
	if dst, ok := val.(string); ok {
		return dst, nil
	}
	return "", fmt.Errorf("failed to ToString, valu is %s %#v", reflect.TypeOf(val).String(), val)
}
