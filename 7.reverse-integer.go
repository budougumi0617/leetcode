package main

import (
	"math"
	"strconv"
)

/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 *
 * https://leetcode.com/problems/reverse-integer/description/
 *
 * algorithms
 * Easy (25.77%)
 * Likes:    3397
 * Dislikes: 5361
 * Total Accepted:    1.1M
 * Total Submissions: 4.3M
 * Testcase Example:  '123'
 *
 * Given a 32-bit signed integer, reverse digits of an integer.
 *
 * Example 1:
 *
 *
 * Input: 123
 * Output: 321
 *
 *
 * Example 2:
 *
 *
 * Input: -123
 * Output: -321
 *
 *
 * Example 3:
 *
 *
 * Input: 120
 * Output: 21
 *
 *
 * Note:
 * Assume we are dealing with an environment which could only store integers
 * within the 32-bit signed integer range: [−2^31,  2^31 − 1]. For the purpose
 * of this problem, assume that your function returns 0 when the reversed
 * integer overflows.
 *
 */

// @lc code=start
func reverse(x int) int {
	b := []byte(strconv.Itoa(x))
	if x < 0 {
		b = b[1:]
	}
	max := len(b)
	for i := 0; i < max/2; i++ {
		b[i], b[max-1-i] = b[max-1-i], b[i]
	}

	num, _ := strconv.Atoi(string(b))
	if x < 0 {
		num = -num
	}
	if math.MaxInt32 < num || num < math.MinInt32 {
		return 0
	}
	return num
}

// @lc code=end
