package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExist(t *testing.T) {
	testCases := []struct {
		name   string
		slices []int
		search int
		want   bool
	}{
		{
			name:   "find item",
			slices: []int{1, 2, 3, 4, 5},
			search: 1,
			want:   true,
		},
		{
			name:   "not find item",
			slices: []int{1, 2, 3, 4, 5},
			search: 7,
			want:   false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			exist := Exist(tc.slices, tc.search)
			assert.Equal(t, tc.want, exist)
		})
	}
}
