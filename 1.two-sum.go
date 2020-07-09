package main

/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	l := make(map[int]int, len(nums))
	// answer - x = y
	for i, n := range nums {
		if j, ok := l[n]; ok {
			return []int{j, i}
		}
		l[target-n] = i
	}
	return []int{}
}

// @lc code=end
