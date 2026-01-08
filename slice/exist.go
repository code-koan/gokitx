package slice

// Exist 数组是否存在元素
func Exist[T comparable](slice []T, search T) bool {
	for _, v := range slice {
		if v == search {
			return true
		}
	}
	return false
}
