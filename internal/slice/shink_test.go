package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShrink(t *testing.T) {
	testCases := []struct {
		name        string
		originCap   int
		enqueueLoop int
		expectCap   int
	}{
		{
			name:        "小于64",
			originCap:   32,
			enqueueLoop: 6,
			expectCap:   32,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			l := make([]int, 0, tc.originCap)
			for i := 0; i < tc.enqueueLoop; i++ {
				l = append(l, i)
			}
			l = Shrink[int](l)
			assert.Equal(t, tc.expectCap, cap(l))
		})
	}
}
