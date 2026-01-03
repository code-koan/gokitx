package slice

import "github.com/code-koan/gokitx/errs"

// Add 在 index 下标添加元素
func Add[T any](src []T, ele T, index int) ([]T, error) {
	length := len(src)
	if index < 0 || index > length {
		return nil, errs.NewErrIndexOutOfRange(length, index)
	}

	var zero T
	src = append(src, zero)
	for i := length; i > index; i-- {
		if i > 0 {
			src[i] = src[i-1]
		}
	}
	src[index] = ele
	return src, nil
}
