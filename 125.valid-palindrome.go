package main

import (
	"fmt"
	"strings"
)

/*
 * @lc app=leetcode id=125 lang=golang
 *
 * [125] Valid Palindrome
 *
 * https://leetcode.com/problems/valid-palindrome/description/
 *
 * algorithms
 * Easy (35.70%)
 * Likes:    1212
 * Dislikes: 2931
 * Total Accepted:    591K
 * Total Submissions: 1.7M
 * Testcase Example:  '"A man, a plan, a canal: Panama"'
 *
 * Given a string, determine if it is a palindrome, considering only
 * alphanumeric characters and ignoring cases.
 *
 * Note: For the purpose of this problem, we define empty string as valid
 * palindrome.
 *
 * Example 1:
 *
 *
 * Input: "A man, a plan, a canal: Panama"
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: "race a car"
 * Output: false
 *
 *
 *
 * Constraints:
 *
 *
 * s consists only of printable ASCII characters.
 *
 *
 */

// @lc code=start
func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	s = strings.TrimSpace(s)
	fmt.Printf("modified %q\n", s)
	s = strings.Trim(s, " !&%$#:")
	fmt.Printf("modified %q\n", s)
	b := []byte(s)
	n := len(s) - 1
	fmt.Printf("modified %q\n", s)

	for i := 0; i < len(s)/2; i++ {
		if b[i] != b[n] {
			return false
		}
		n--
	}

	return false
}

// @lc code=end
