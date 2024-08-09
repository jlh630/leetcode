/*
 * @lc app=leetcode.cn id=62 lang=golang
 * @lcpr version=30204
 *
 * [62] 不同路径
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
// 动态规划
func uniquePaths(m int, n int) int {
	if n == 1 || m == 1 {
		return 1
	}
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := 1; i < m; i++ {
		dp[i][0] = 1
	}
	for j := 1; j < n; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

// @lc code=end

/*
// @lcpr case=start
// 3\n7\n
// @lcpr case=end

// @lcpr case=start
// 3\n2\n
// @lcpr case=end

// @lcpr case=start
// 7\n3\n
// @lcpr case=end

// @lcpr case=start
// 3\n3\n
// @lcpr case=end

*/

