package slice

import "errors"

var ErrIndexOutOfRange = errors.New("下标超出范围")

// DeleteAt 删除指定位置的元素
// 如果下标不是合法的下标，返回 ErrIndexOutOfRange
func DeleteAt(s []int, idx int) ([]int, error) {
	length := len(s)
	if idx < 0 || idx >= length {
		return nil, ErrIndexOutOfRange
	}
	j := 0
	for i, v := range s {
		if i != idx {
			s[j] = v
			j++
		}
	}
	s = s[:j]
	return s, nil
}
