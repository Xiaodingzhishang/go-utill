package slice

import (
	"reflect"
	"testing"
)

func TestDeleteAt(t *testing.T) {

	tests := []struct {
		name    string
		s       []int
		idx     int
		want    []int
		wantErr bool
	}{
		{
			name: "index 0",
			s:    []int{1, 2, 3},
			idx:  0,
			want: []int{2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeleteAt(tt.s, tt.idx)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteAt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteAt() got = %v, want %v", got, tt.want)
			}
		})
	}
}
