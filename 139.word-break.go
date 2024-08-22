/*
 * @lc app=leetcode.cn id=139 lang=golang
 * @lcpr version=30204
 *
 * [139] 单词拆分
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func wordBreak(s string, wordDict []string) bool {
	wordSet := make(map[string]bool)
	for _, v := range wordDict {
		wordSet[v] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for r := 1; r <= len(s); r++ {
		for l := 0; l < r; l++ {
			if dp[l] && wordSet[s[l:r]] {
				dp[r] = true
				break
			}
		}
	}
	return dp[len(s)]
}

// @lc code=end

/*
// @lcpr case=start
// "leetcode"\n["leet", "code"]\n
// @lcpr case=end

// @lcpr case=start
// "applepenapple"\n["apple", "pen"]\n
// @lcpr case=end

// @lcpr case=start
// "catsandog"\n["cats", "dog", "sand", "and", "cat"]\n
// @lcpr case=end

// @lcpr case=start
// "leetcode"\n["leetcode","a","b","c"]\n
// @lcpr case=end

*/

