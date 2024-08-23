/*
 * @lc app=leetcode.cn id=73 lang=golang
 * @lcpr version=30204
 *
 * [73] 矩阵置零
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func setZeroes(matrix [][]int) {
	n := len(matrix)
	m := len(matrix[0])
	rows := make([]int, n)
	cloms := make([]int, m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				rows[i] = 1
				cloms[j] = 1
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if rows[i] == 1 || cloms[j] == 1 {
				matrix[i][j] = 0
			}
		}
	}
}

// @lc code=end

/*
// @lcpr case=start
// [[1,1,1],[1,0,1],[1,1,1]]\n
// @lcpr case=end

// @lcpr case=start
// [[0,1,2,0],[3,4,5,2],[1,3,1,5]]\n
// @lcpr case=end

*/

