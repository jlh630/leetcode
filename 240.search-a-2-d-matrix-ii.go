/*
 * @lc app=leetcode.cn id=240 lang=golang
 * @lcpr version=30204
 *
 * [240] 搜索二维矩阵 II
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
package main

func searchMatrix(matrix [][]int, target int) bool {
	curri := 0
	curry := len(matrix[0]) - 1
	for curri < len(matrix) && curry >= 0 {
		if matrix[curri][curry] == target {
			return true
		} else if matrix[curri][curry] < target {
			curri++
		} else {
			curry--
		}
	}
	return false
}

// @lc code=end

/*
// @lcpr case=start
// [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]]\n5\n
// @lcpr case=end

// @lcpr case=start
// [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]]\n20\n
// @lcpr case=end

*/
