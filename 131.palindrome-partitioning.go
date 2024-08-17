/*
 * @lc app=leetcode.cn id=131 lang=golang
 * @lcpr version=30204
 *
 * [131] 分割回文串
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func partition(s string) (ans [][]string) {
	n := len(s)
	dp := make([][]bool, n)
	for index := range dp {
		dp[index] = make([]bool, n)
		dp[index][index] = true
	}
	for j := 1; j < n; j++ {
		for i := 0; i < j; i++ {
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j-i < 2 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
		}
	}
	res := make([][]string, 0)
	path := []string{}
	var dfs func(count int)
	dfs = func(count int) {
		if count == n {
			res = append(res, append([]string{}, path...))
			return
		}
		for i := count; i < n; i++ {
			if dp[count][i] {
				path = append(path, s[count:i+1])
				dfs(i + 1)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)
	return res
}

// @lc code=end

/*
// @lcpr case=start
// "aab"\n
// @lcpr case=end

// @lcpr case=start
// "a"\n
// @lcpr case=end

*/

