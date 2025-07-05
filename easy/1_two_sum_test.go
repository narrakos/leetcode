package easy_test

import (
	"reflect"
	"testing"

	"github.com/narrakos/leetcode/easy"
)

func TestTwoSum(t *testing.T) {
	implementations := []struct {
		name string
		impl easy.TwoSumFunc
	}{
		{name: "TwoSum_BruteForce", impl: easy.TwoSum},
		{name: "TwoSum_Map", impl: easy.TwoSumMap},
	}

	for _, impl := range implementations {
		t.Run(impl.name, func(t *testing.T) {
			testTwoSumHelper(impl.name, impl.impl, t)
		})
	}
}

func testTwoSumHelper(funcName string, f easy.TwoSumFunc, t *testing.T) {
	t.Helper()
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "Example 1: [2,7,11,15] target=9",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			name:   "Example 2: [3,2,4] target=6",
			nums:   []int{3, 2, 4},
			target: 6,
			want:   []int{1, 2},
		},
		{
			name:   "Example 3: [3,3] target=6",
			nums:   []int{3, 3},
			target: 6,
			want:   []int{0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := f(tt.nums, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%s(%v, %d) = %v, want %v", funcName, tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
