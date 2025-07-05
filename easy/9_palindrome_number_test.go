package easy_test

import (
	"testing"

	"github.com/narrakos/leetcode/easy"
)

func TestIsPalindrome(t *testing.T) {
	implementations := []struct {
		name string
		impl easy.IsPalindromeFunc
	}{
		{name: "IsPalindromeString", impl: easy.IsPalindromeString},
		{name: "IsPalindromeMod", impl: easy.IsPalindromeMod},
	}

	for _, impl := range implementations {
		t.Run(impl.name, func(t *testing.T) {
			testIsPalindromeHelper(impl.name, impl.impl, t)
		})
	}
}

func testIsPalindromeHelper(name string, f easy.IsPalindromeFunc, t *testing.T) {
	t.Helper()
	tests := []struct {
		name string
		x    int
		want bool
	}{
		{name: "Example 1", x: 121, want: true},
		{name: "Example 2", x: -121, want: false},
		{name: "Example 3", x: 10, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := f(tt.x)
			if got != tt.want {
				t.Errorf("%s(%d) = %v, want %v", name, tt.x, got, tt.want)
			}
		})
	}
}
