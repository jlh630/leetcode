/*
 * @lc app=leetcode.cn id=5 lang=golang
 * @lcpr version=30204
 *
 * [5] 最长回文子串
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func longestPalindrome(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	// 回文串去掉头和尾还是子串  dp[i][j]=dp[i+1][j-1]&&s[i] == s[j]
	/*
			i:串的左边下标 j串的右边下标
			0   1  2  3
		0	[t] [] [] []
		1	[] [t] [] []
		2	[] [] [t] []
		3	[] [] [] [t]
	*/
	dp := make([][]bool, n)
	for index := range dp {
		dp[index] = make([]bool, n)
	}
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}
	start := 0
	max_len := 1
	for j := 1; j < n; j++ {
		for i := 0; i < j; i++ {
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j-i < 2 {
					//这种情况肯定是true
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}

				if dp[i][j] && max_len < j-i+1 {
					max_len = j - i + 1
					start = i
				}

			}
		}
	}
	return s[start : start+max_len]
}

// @lc code=end

/*
// @lcpr case=start
// "babad"\n
// @lcpr case=end

// @lcpr case=start
// "cbbd"\n
// @lcpr case=end

*/

