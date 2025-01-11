package cast

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

func IntToBool[T ~int8 | ~int](v T) bool {
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
