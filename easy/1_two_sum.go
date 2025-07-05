package easy

// https://leetcode.com/problems/two-sum/description/
// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.

type TwoSumFunc func([]int, int) []int

func TwoSum(nums []int, target int) []int {
	//Time complexity: O(n^2)
	//Space complexity: O(1)
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func TwoSumMap(nums []int, target int) []int {
	//Time complexity: O(n)
	//Space complexity: O(n)
	m := make(map[int]int)

	for i, num := range nums {
		diff := target - num
		if _, ok := m[diff]; ok {
			return []int{m[diff], i}
		}
		m[num] = i
	}
	return nil
}
