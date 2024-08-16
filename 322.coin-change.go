/*
 * @lc app=leetcode.cn id=322 lang=golang
 * @lcpr version=30204
 *
 * [322] 零钱兑换
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
// 动态规划
// |   min (dfs(amount-coins[i],coins[i])+1,dfs(amount,coins[i-1]))  amount>0
// |       0                                                         amount==0
// |       math.MaxInt                                             coins[i] >amount
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i < amount+1; i++ {
		dp[i] = math.MaxInt / 2
	}
	dp[0] = 0
	for _, money := range coins {
		for j := money; j < amount+1; j++ {
			dp[j] = min(dp[j], dp[j-money]+1)
		}
	}
	res := dp[amount]
	if res < math.MaxInt/2 {
		return res
	} else {
		return -1
	}

}

// @lc code=end

/*
// @lcpr case=start
// [1, 2, 5]\n11\n
// @lcpr case=end

// @lcpr case=start
// [2]\n3\n
// @lcpr case=end

// @lcpr case=start
// [1]\n0\n
// @lcpr case=end

*/

