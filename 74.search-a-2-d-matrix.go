/*
 * @lc app=leetcode.cn id=74 lang=golang
 * @lcpr version=30204
 *
 * [74] 搜索二维矩阵
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
package main

func searchMatrix(matrix [][]int, target int) bool {

	left := 0
	right := len(matrix) - 1
	row := -1
	for left <= right {
		mid := (left + right) / 2
		if matrix[mid][0] == target {
			row = mid
			break
		} else if matrix[mid][0] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if row < 0 {
		if right < 0 {
			right = 0
		}
		row = right

	}
	left = 0
	right = len(matrix[0]) - 1
	for left <= right {
		mid := (left + right) / 2
		if matrix[row][mid] == target {
			return true
		} else if matrix[row][mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false

}

// @lc code=end

/*
// @lcpr case=start
// [[1,3,5,7],[10,11,16,20],[23,30,34,60]]\n3\n
// @lcpr case=end

// @lcpr case=start
// [[1,3,5,7],[10,11,16,20],[23,30,34,60]]\n13\n
// @lcpr case=end

*/
