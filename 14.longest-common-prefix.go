package main

import "strings"

/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 *
 * https://leetcode.com/problems/longest-common-prefix/description/
 *
 * algorithms
 * Easy (35.80%)
 * Likes:    3371
 * Dislikes: 2044
 * Total Accepted:    886.8K
 * Total Submissions: 2.5M
 * Testcase Example:  '["flower","flow","flight"]'
 *
 * Write a function to find the longest common prefix string amongst an array
 * of strings.
 *
 * If there is no common prefix, return an empty string "".
 *
 *
 * Example 1:
 *
 *
 * Input: strs = ["flower","flow","flight"]
 * Output: "fl"
 *
 *
 * Example 2:
 *
 *
 * Input: strs = ["dog","racecar","car"]
 * Output: ""
 * Explanation: There is no common prefix among the input strings.
 *
 *
 *
 * Constraints:
 *
 *
 * 0 <= strs.length <= 200
 * 0 <= strs[i].length <= 200
 * strs[i] consists of only lower-case English letters.
 *
 *
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	ans := strs[0]
	if len(strs) == 1 {
		return ans
	}

	for i := 0; i <= len(ans); i++ {
		for _, str := range strs[1:] {
			if !strings.HasPrefix(str, ans[:i]) {
				return ans[:i-1]
			}
		}
	}
	return ans
}

// @lc code=end
