package slice

// Delete 删除指定位置的元素
// 如果下标不是合法的下标，返回 ErrIndexOutOfRange
func Delete[T any](src []T, index int) ([]T, T, error) {
	length := len(src)
	if index < 0 || index >= length {
		var zero T
		return nil, zero, ErrIndexOutOfRange
	}
	res := src[index]
	for i := index; i+1 < length; i++ {
		src[i] = src[i+1]
	}
	src = src[:length-1]

	return src, res, nil
}
