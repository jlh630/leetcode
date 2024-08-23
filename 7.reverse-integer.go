/*
 * @lc app=leetcode.cn id=7 lang=golang
 * @lcpr version=30204
 *
 * [7] 整数反转
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func reverse(x int) int {
	sign := 1
	if x < 0 {
		sign = -1
		x = -1 * x
	}
	ans := 0
	for x > 0 {
		if ans*sign < math.MinInt32/10 || ans*sign > math.MaxInt32/10 {
			return 0
		}
		ans = ans*10 + x%10
		x /= 10
	}
	return sign * ans
}

// @lc code=end

/*
// @lcpr case=start
// 123\n
// @lcpr case=end

// @lcpr case=start
// -123\n
// @lcpr case=end

// @lcpr case=start
// 120\n
// @lcpr case=end

// @lcpr case=start
// 0\n
// @lcpr case=end

*/

