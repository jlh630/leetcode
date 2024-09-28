/*
 * @lc app=leetcode.cn id=1143 lang=golang
 * @lcpr version=30204
 *
 * [1143] 最长公共子序列
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
// 动态规划
//
//	| dfs(i-1,j-1)+1  s1[i]==s2[i]
//	|
//	| max(dfs(i-1,j),dfs(i,j-1))  s1[i]!=s[i]
package main

func longestCommonSubsequence(text1 string, text2 string) int {
	n := len(text1) + 1
	m := len(text2) + 1
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[n-1][m-1]
}

// @lc code=end

/*
// @lcpr case=start
// "abcde"\n"ace"\n
// @lcpr case=end

// @lcpr case=start
// "abc"\n"abc"\n
// @lcpr case=end

// @lcpr case=start
// "abc"\n"def"\n
// @lcpr case=end

*/
