package easy

import "strconv"

// https://leetcode.com/problems/palindrome-number/description/
// Given an integer x, return true if x is a palindrome, and false otherwise.
//
// Constraints: -2^31 <= x <= 2^31 - 1

type IsPalindromeFunc func(int) bool

func IsPalindromeString(x int) bool {
	//Time complexity: O(n)
	//Space complexity: O(n)
	if x < 0 {
		return false
	}

	str := strconv.Itoa(x)
	strLen := len(str)

	for i := 0; i < strLen/2; i++ {
		if str[i] != str[strLen-i-1] {
			return false
		}
	}

	return true
}

func IsPalindromeMod(x int) bool {
	//Time complexity: O(log n )
	//Space complexity: O(1)
	if x < 0 {
		return false
	}

	top := 1
	for x/top >= 10 {
		top *= 10
	}

	for x > 0 {
		if x/top != x%10 {
			return false
		}

		x = (x % top) / 10
		top /= 100
	}

	return true
}
