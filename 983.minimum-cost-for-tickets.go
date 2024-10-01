/*
 * @lc app=leetcode.cn id=983 lang=golang
 * @lcpr version=30204
 *
 * [983] 最低票价
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func mincostTickets(days []int, costs []int) int {
	n := len(days)
	maxDay := days[n-1] + 31
	dp := make([]int, maxDay)
	daysbool := make([]bool, maxDay)
	for i := range days {
		daysbool[days[i]+30] = true
	}
	for i := 31; i < maxDay; i++ {
		if daysbool[i] {
			dp[i] = min(dp[i-1]+costs[0], dp[i-7]+costs[1], dp[i-30]+costs[2])
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[maxDay-1]
}

// @lc code=end

/*
// @lcpr case=start
// [1,4,6,7,8,20]\n[2,7,15]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4,5,6,7,8,9,10,30,31]\n[2,7,15]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n[7,2,15]\n
// @lcpr case=end

*/

