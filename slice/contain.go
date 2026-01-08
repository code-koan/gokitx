package slice

// Contains 是否包含相应数组
func Contains[T comparable](src, target []T) bool {
	mp := toMap(src)
	for _, v := range target {
		if _, ok := mp[v]; !ok {
			return false
		}
	}
	return true
}
