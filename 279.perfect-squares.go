/*
 * @lc app=leetcode.cn id=279 lang=golang
 * @lcpr version=30204
 *
 * [279] 完全平方数
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
// 动态规划 跟零钱兑换一样 不过需要自己算能兑换的零钱
func numSquares(n int) int {
	squareNum := int(math.Sqrt(float64(n)))
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = math.MaxInt / 2
	}
	dp[0] = 0
	for i := 1; i <= squareNum; i++ {
		for j := i * i; j < n+1; j++ {
			dp[j] = min(dp[j-i*i]+1, dp[j])
		}
	}
	return dp[n]
}

// @lc code=end

/*
// @lcpr case=start
// 12\n
// @lcpr case=end

// @lcpr case=start
// 13\n
// @lcpr case=end

@lcpr case=start
// 100\n
// @lcpr case=end

*/

