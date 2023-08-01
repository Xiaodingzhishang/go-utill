package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteAt(t *testing.T) {

	testCases := []struct {
		name      string
		slice     []int
		index     int
		wantSlice []int
		wantVal   int
		wantErr   error
	}{
		{
			name:    "nil",
			slice:   nil,
			wantErr: ErrIndexOutOfRange,
		},
		{
			name:      "index 0",
			slice:     []int{1, 2, 3},
			index:     0,
			wantVal:   1,
			wantSlice: []int{2, 3},
		},
		{
			name:    "boundary",
			index:   3,
			slice:   []int{1, 2, 3},
			wantErr: ErrIndexOutOfRange,
		},
		{
			name:      "index last",
			slice:     []int{255, 369, 581, 732, 693},
			index:     4,
			wantVal:   693,
			wantSlice: []int{255, 369, 581, 732},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, val, err := Delete(tc.slice, tc.index)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}

			assert.Equal(t, tc.wantSlice, res)
			assert.Equal(t, tc.wantVal, val)
		})
	}
}
