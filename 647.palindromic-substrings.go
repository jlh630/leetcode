/*
 * @lc app=leetcode.cn id=647 lang=golang
 * @lcpr version=30204
 *
 * [647] 回文子串
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func countSubstrings(s string) int {
	n := len(s)
	if n == 1 {
		return 1
	}
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}
	ans := n
	for j := 0; j < n; j++ {
		for i := 0; i < j; i++ {
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j-i < 2 {
					dp[i][j] = true
					ans++
					continue
				}
				dp[i][j] = dp[i+1][j-1]
				if dp[i+1][j-1] == true {
					ans++
				}
			}
		}
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// "abc"\n
// @lcpr case=end

// @lcpr case=start
// "aaa"\n
// @lcpr case=end

// @lcpr case=start
// "abcabc"\n
// @lcpr case=end
*/

