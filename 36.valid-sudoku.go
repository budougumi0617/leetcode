package main

import "fmt"

/*
 * @lc app=leetcode id=36 lang=golang
 *
 * [36] Valid Sudoku
 *
 * https://leetcode.com/problems/valid-sudoku/description/
 *
 * algorithms
 * Medium (48.40%)
 * Likes:    1696
 * Dislikes: 448
 * Total Accepted:    369.3K
 * Total Submissions: 760K
 * Testcase Example:  '[["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]'
 *
 * Determine if a 9x9 Sudoku board is valid. Only the filled cells need to be
 * validated according to the following rules:
 *
 *
 * Each row must contain the digits 1-9 without repetition.
 * Each column must contain the digits 1-9 without repetition.
 * Each of the 9 3x3 sub-boxes of the grid must contain the digits 1-9 without
 * repetition.
 *
 *
 *
 * A partially filled sudoku which is valid.
 *
 * The Sudoku board could be partially filled, where empty cells are filled
 * with the character '.'.
 *
 * Example 1:
 *
 *
 * Input:
 * [
 * ⁠ ["5","3",".",".","7",".",".",".","."],
 * ⁠ ["6",".",".","1","9","5",".",".","."],
 * ⁠ [".","9","8",".",".",".",".","6","."],
 * ⁠ ["8",".",".",".","6",".",".",".","3"],
 * ⁠ ["4",".",".","8",".","3",".",".","1"],
 * ⁠ ["7",".",".",".","2",".",".",".","6"],
 * ⁠ [".","6",".",".",".",".","2","8","."],
 * ⁠ [".",".",".","4","1","9",".",".","5"],
 * ⁠ [".",".",".",".","8",".",".","7","9"]
 * ]
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input:
 * [
 * ["8","3",".",".","7",".",".",".","."],
 * ["6",".",".","1","9","5",".",".","."],
 * [".","9","8",".",".",".",".","6","."],
 * ["8",".",".",".","6",".",".",".","3"],
 * ["4",".",".","8",".","3",".",".","1"],
 * ["7",".",".",".","2",".",".",".","6"],
 * [".","6",".",".",".",".","2","8","."],
 * [".",".",".","4","1","9",".",".","5"],
 * [".",".",".",".","8",".",".","7","9"]
 * ]
 * Output: false
 * Explanation: Same as Example 1, except with the 5 in the top left corner
 * being
 * ⁠   modified to 8. Since there are two 8's in the top left 3x3 sub-box, it
 * is invalid.
 *
 *
 * Note:
 *
 *
 * A Sudoku board (partially filled) could be valid but is not necessarily
 * solvable.
 * Only the filled cells need to be validated according to the mentioned
 * rules.
 * The given board contain only digits 1-9 and the character '.'.
 * The given board size is always 9x9.
 *
 *
 */

// @lc code=start
func isValidSudoku(board [][]byte) bool {
	for i := 0; i < len(board); i++ {
		bitsRow := 0
		bitsColumn := 0

		for j := 0; j < len(board); j++ {
			if board[i][j] != '.' {
				if (1 << board[i][j] & bitsRow) != 0 {
					fmt.Printf("rows %d, %d : %q was false\n", i, j, board[i][j])
					return false
				}
				bitsRow |= 1 << board[i][j]
			}

			if board[j][i] != '.' {
				if (1 << board[j][i] & bitsColumn) != 0 {
					fmt.Printf("column %d, %d : %q was false\n", j, i, board[j][i])
					return false
				}
				bitsColumn |= 1 << board[j][i]
			}
		}
	}

	for i := 0; i < len(board); i += 3 {
		for j := 0; j < len(board); j += 3 {
			bitsRow := 0
			for ii := 0; ii < 3; ii++ {
				for jj := 0; jj < 3; jj++ {
					fmt.Println("loop3")
					iii := i + ii
					jjj := j + jj
					if board[iii][jjj] != '.' {
						if (1 << board[iii][jjj] & bitsRow) != 0 {
							fmt.Printf("matrix %d,%d %d %d: %q was false\n", i, ii, j, jj, board[iii][jjj])
							return false
						}
						fmt.Printf("matrix %d, %d : %q was ok\n", iii, jjj, board[iii][jjj])
						bitsRow |= 1 << board[iii][jjj]
					}
				}
			}
		}
	}

	return true
}

// @lc code=end
