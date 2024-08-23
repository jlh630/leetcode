/*
 * @lc app=leetcode.cn id=221 lang=golang
 * @lcpr version=30204
 *
 * [221] 最大正方形
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maximalSquare(matrix [][]byte) int {
	n := len(matrix)
	m := len(matrix[0])
	ans := 0
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = 1
				ans = 1
			}
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if dp[i][j] == 1 {
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
				ans = max(ans, dp[i][j])
			}
		}
	}
	return ans * ans
}

// @lc code=end

/*
// @lcpr case=start
// [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]\n
// @lcpr case=end

// @lcpr case=start
// [["0","1"],["1","0"]]\n
// @lcpr case=end

// @lcpr case=start
// [["1"]]\n
// @lcpr case=end

// @lcpr case=start
// [["1","1","1","1","0"],["0","1","1","1","1"],["1","1","1","1","1"],["1","1","1","1","0"]]\n
// @lcpr case=end
*/

