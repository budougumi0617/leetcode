package main

import "fmt"

/*
 * @lc app=leetcode id=189 lang=golang
 *
 * [189] Rotate Array
 *
 * https://leetcode.com/problems/rotate-array/description/
 *
 * algorithms
 * Easy (34.37%)
 * Likes:    2841
 * Dislikes: 816
 * Total Accepted:    495.2K
 * Total Submissions: 1.4M
 * Testcase Example:  '[1,2,3,4,5,6,7]\n3'
 *
 * Given an array, rotate the array to the right by k steps, where k is
 * non-negative.
 *
 * Follow up:
 *
 *
 * Try to come up as many solutions as you can, there are at least 3 different
 * ways to solve this problem.
 * Could you do it in-place with O(1) extra space?
 *
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [1,2,3,4,5,6,7], k = 3
 * Output: [5,6,7,1,2,3,4]
 * Explanation:
 * rotate 1 steps to the right: [7,1,2,3,4,5,6]
 * rotate 2 steps to the right: [6,7,1,2,3,4,5]
 * rotate 3 steps to the right: [5,6,7,1,2,3,4]
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [-1,-100,3,99], k = 2
 * Output: [3,99,-1,-100]
 * Explanation:
 * rotate 1 steps to the right: [99,-1,-100,3]
 * rotate 2 steps to the right: [3,99,-1,-100]
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 2 * 10^4
 * It's guaranteed that nums[i] fits in a 32 bit-signed integer.
 * k >= 0
 *
 *
 */

// @lc code=start
func rotate(nums []int, k int) {
	// use 'Approach 3: Using Cyclic Replacements'
	k %= len(nums)
	var count int

	for s := 0; count < len(nums); s++ {
		current := s
		prev := nums[s]

		for {
			n := (current + k) % len(nums)
			tmp := nums[n]
			nums[n] = prev
			prev = tmp
			current = n
			count++

			if s != current {
				break
			}
		}
		fmt.Printf("current %v\n", nums)
	}
}

//   public void rotate(int[] nums, int k) {
//     k = k % nums.length;
//     int count = 0;
//     for (int start = 0; count < nums.length; start++) {
//       int current = start;
//       int prev = nums[start];
//       do {
//         int next = (current + k) % nums.length;
//         int temp = nums[next];
//         nums[next] = prev;
//         prev = temp;
//         current = next;
//         count++;
//       } while (start != current);
//     }
//   }
// }

// @lc code=end
