package array

func IsContain[T comparable](val T, list []T) bool {
	for _, item := range list {
		if item == val {
			return true
		}
	}
	return false
}

func Map[T any, V any](list []T, fc func(T) V) []V {
	dst := make([]V, len(list))
	for i, val := range list {
		dst[i] = fc(val)
	}
	return dst
}

func Filter[T comparable](list []T, fc func(T) bool) []T {
	dst := make([]T, 0)
	for _, v := range list {
		if fc(v) {
			dst = append(dst, v)
		}
	}
	return dst
}

func RemoveDuplicates[T comparable](list []T) []T {
	dst := make([]T, 0)
	for _, v := range list {
		if !IsContain(v, dst) {
			dst = append(dst, v)
		}
	}
	return dst
}
