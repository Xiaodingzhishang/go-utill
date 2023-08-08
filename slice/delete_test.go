package slice

import (
	"github.com/Xiaodingzhishang/go-utill/internal/slice"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDetele(t *testing.T) {
	testCases := []struct {
		name      string
		slice     []int
		index     int
		wantSlice []int
		wantErr   error
	}{
		{
			name:      "index 0",
			slice:     []int{123, 100},
			index:     0,
			wantSlice: []int{100},
		},
		{
			name:    "index -1",
			slice:   []int{123, 100},
			index:   -1,
			wantErr: slice.ErrIndexOutOfRange,
		},
	}
	for _, ts := range testCases {
		t.Run(ts.name, func(t *testing.T) {
			res, err := Detele(ts.slice, ts.index)
			assert.Equal(t, ts.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, ts.wantSlice, res)
		})
	}
}
