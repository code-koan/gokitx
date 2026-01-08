package slice

// ToMap slice 转换为 map
func ToMap[Src, Key comparable](src []Src, fn func(s Src) Key) map[Key]Src {
	return ToMapV(src, func(s Src) (Key, Src) {
		return fn(s), s
	})
}

// ToMapV slice 转换为 map
func ToMapV[Src, Key, Value comparable](src []Src, fn func(s Src) (Key, Value)) map[Key]Value {
	res := make(map[Key]Value, len(src))
	for _, v := range src {
		k, val := fn(v)
		res[k] = val
	}
	return res
}

// MapX 数组类型转换函数
func MapX[Src any, Target any](src []Src, f func(Src) Target) (res []Target) {
	res = make([]Target, len(src))
	for i, s := range src {
		res[i] = f(s)
	}
	return
}

func toMap[T comparable](src []T) map[T]struct{} {
	res := make(map[T]struct{}, len(src))
	for _, s := range src {
		res[s] = struct{}{}
	}
	return res
}
