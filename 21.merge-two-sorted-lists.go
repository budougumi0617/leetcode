package main

/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 *
 * https://leetcode.com/problems/merge-two-sorted-lists/description/
 *
 * algorithms
 * Easy (53.10%)
 * Likes:    4441
 * Dislikes: 603
 * Total Accepted:    1M
 * Total Submissions: 1.9M
 * Testcase Example:  '[1,2,4]\n[1,3,4]'
 *
 * Merge two sorted linked lists and return it as a new sorted list. The new
 * list should be made by splicing together the nodes of the first two lists.
 *
 * Example:
 *
 *
 * Input: 1->2->4, 1->3->4
 * Output: 1->1->2->3->4->4
 *
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
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var res *ListNode
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 != nil {
		if l2 != nil && l1.Val > l2.Val {
			res = l2
			l2 = l2.Next
		} else {
			res = l1
			l1 = l1.Next
		}
	} else {
		res = l2
		l2 = l2.Next
	}
	root := res

	for {
		if l1 != nil && l2 != nil {
			// fmt.Printf("current status l1 %d, l2 %d\n", l1.Val, l2.Val)
			if l1.Val > l2.Val {
				res.Next = l2
				l2 = l2.Next
			} else {
				res.Next = l1
				l1 = l1.Next
			}
		} else if l1 != nil {
			res.Next = l1
			l1 = l1.Next
		} else if l2 != nil {
			res.Next = l2
			l2 = l2.Next
		} else {
			return root
		}
		res = res.Next
	}
}

// @lc code=end
