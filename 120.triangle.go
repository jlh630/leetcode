/*
 * @lc app=leetcode.cn id=120 lang=golang
 * @lcpr version=30204
 *
 * [120] 三角形最小路径和
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp := make([]int, n)
	dp[0] = triangle[0][0]
	for i := 1; i < n; i++ {
		m := len(triangle[i])
		dp[m-1] = triangle[i][m-1] + dp[m-2]
		for j := m - 2; j >= 1; j-- {
			dp[j] = min(dp[j-1], dp[j]) + triangle[i][j]
		}
		dp[0] += triangle[i][0]

	}
	ans := math.MaxInt32
	for i := range dp {
		ans = min(ans, dp[i])
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [[2],[3,4],[6,5,7],[4,1,8,3]]\n
// @lcpr case=end

// @lcpr case=start
// [[-10]]\n
// @lcpr case=end

*/

