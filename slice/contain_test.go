package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainAll(t *testing.T) {
	testCases := []struct {
		name     string
		src      []int
		target   []int
		expected bool
	}{
		{
			name:     "all contained",
			src:      []int{1, 2, 3, 4},
			target:   []int{2, 3},
			expected: true,
		},
		{
			name:     "partial",
			src:      []int{1, 2, 3},
			target:   []int{2, 4},
			expected: false,
		},
		{
			name:     "empty target",
			src:      []int{1, 2, 3},
			target:   []int{},
			expected: true,
		},
		{
			name:     "empty src",
			src:      []int{},
			target:   []int{1},
			expected: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Contains(tc.src, tc.target)
			assert.Equal(t, tc.expected, res)
		})
	}
}
