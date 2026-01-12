package slice

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapX(t *testing.T) {
	testCases := []struct {
		name       string
		src        []int
		fn         func(int) int
		wantTarget []int
	}{
		{
			name: "success",
			src:  []int{1, 2, 3},
			fn: func(i int) int {
				return i + 1
			},
			wantTarget: []int{2, 3, 4},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := MapX(tc.src, tc.fn)
			assert.Equal(t, res, tc.wantTarget)
		})
	}
}

func TestToMap(t *testing.T) {
	testCases := []struct {
		name string
		src  []int
		fn   func(int) string
		want map[string]int
	}{
		{
			name: "normal",
			src:  []int{1, 2, 3},
			fn:   func(i int) string { return fmt.Sprintf("k%d", i) },
			want: map[string]int{"k1": 1, "k2": 2, "k3": 3},
		},
		{
			name: "duplicate key",
			src:  []int{1, 1, 2},
			fn:   func(i int) string { return fmt.Sprintf("k%d", i%2) },
			want: map[string]int{"k1": 1, "k0": 2}, // last wins
		},
		{
			name: "empty",
			src:  []int{},
			fn:   func(i int) string { return "" },
			want: map[string]int{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := ToMap(tc.src, tc.fn)
			assert.Equal(t, tc.want, res)
		})
	}
}

func TestToMapV(t *testing.T) {
	testCases := []struct {
		name string
		src  []int
		fn   func(int) (string, int)
		want map[string]int
	}{
		{
			name: "normal",
			src:  []int{1, 2, 3},
			fn:   func(i int) (string, int) { return fmt.Sprintf("k%d", i), i * 10 },
			want: map[string]int{"k1": 10, "k2": 20, "k3": 30},
		},
		{
			name: "duplicate key",
			src:  []int{1, 2, 3},
			fn:   func(i int) (string, int) { return "k", i },
			want: map[string]int{"k": 3}, // last wins
		},
		{
			name: "empty",
			src:  []int{},
			fn:   func(i int) (string, int) { return "", 0 },
			want: map[string]int{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := ToMapV(tc.src, tc.fn)
			assert.Equal(t, tc.want, res)
		})
	}
}

func TestFilterMap(t *testing.T) {
	testCases := []struct {
		name       string
		src        []int
		fn         func(int) (int, bool)
		wantTarget []int
	}{
		{
			name: "filter even and add one",
			src:  []int{1, 2, 3, 4, 5},
			fn: func(i int) (int, bool) {
				if i%2 == 0 {
					return i + 1, true
				}
				return 0, false
			},
			wantTarget: []int{3, 5},
		},
		{
			name:       "empty input",
			src:        []int{},
			fn:         func(i int) (int, bool) { return i, true },
			wantTarget: []int{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := FilterMap(tc.src, tc.fn)
			assert.Equal(t, res, tc.wantTarget)
		})
	}
}
