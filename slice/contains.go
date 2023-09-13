package slice

// Contains 检查slice中是否包含某个数值
func Contains[T comparable](s []T, v T) bool {
	for _, item := range s {
		if item == v {
			return true
		}
	}

	return false
}
