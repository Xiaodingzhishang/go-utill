package slice

// Delete 删除指定位置的元素
// 如果下标不是合法的下标，返回 ErrIndexOutOfRange
func Delete[T any](src []T, index int) ([]T, T, error) {
	length := len(src)
	if index < 0 || index >= length {
		var zero T
		return nil, zero, ErrIndexOutOfRange
	}
	j := 0
	res := src[index]
	for i, v := range src {
		if i != index {
			src[j] = v
			j++
		}
	}
	src = src[:j]
	return src, res, nil
}
