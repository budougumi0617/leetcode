package main

/*
 * @lc app=leetcode id=234 lang=golang
 *
 * [234] Palindrome Linked List
 *
 * https://leetcode.com/problems/palindrome-linked-list/description/
 *
 * algorithms
 * Easy (39.01%)
 * Likes:    4025
 * Dislikes: 396
 * Total Accepted:    501.5K
 * Total Submissions: 1.3M
 * Testcase Example:  '[1,2]'
 *
 * Given a singly linked list, determine if it is a palindrome.
 *
 * Example 1:
 *
 *
 * Input: 1->2
 * Output: false
 *
 * Example 2:
 *
 *
 * Input: 1->2->2->1
 * Output: true
 *
 * Follow up:
 * Could you do it in O(n) time and O(1) space?
 *
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	var n int
	nodes := []*ListNode{}
	c := head
	for {
		if c == nil {
			break
		}
		nodes = append(nodes, c)
		c = c.Next
		n++
	}
	j := n - 1
	i := 0
	m := n / 2
	for {
		if nodes[i].Val != nodes[j].Val {
			return false
		}
		i++
		j--
		if i > m {
			break
		}
	}
	return true
}

// @lc code=end
