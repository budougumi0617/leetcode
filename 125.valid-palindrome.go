package main

import (
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
func isAlphanumeric(r rune) bool {
	if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') && (r < '0' || r > '9') {
		return false
	}
	return true
}

func trim(s string) []byte {
	b := make([]byte, len(s))

	c := 0
	for i, r := range s {
		if isAlphanumeric(r) {
			b[c] = s[i]
			c++
		}
	}
	if c == 0 {
		return []byte{}
	}
	b = b[:c]
	return b
}

func isPalindrome(s string) bool {
	if len(s) < 2 {
		return true
	}
	s = strings.ToLower(s)

	b := trim(s)
	if len(b) < 1 {
		return true
	}
	// fmt.Printf("modified %q\n", b)
	n := len(b) - 1

	for i := 0; i < len(b)/2; i++ {
		// fmt.Printf("%q vs %q\n", b[i], b[n])
		if b[i] != b[n] {
			return false
		}
		n--
	}

	return true
}

// @lc code=end
