/*
 * @lc app=leetcode.cn id=22 lang=golang
 * @lcpr version=30204
 *
 * [22] 括号生成
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
package main

import "strings"

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	var curr strings.Builder
	var dfs func(int, int, int)
	dfs = func(left int, right int, n int) {
		if curr.Len() == n*2 {
			res = append(res, curr.String())
		}
		if left < n {
			curr.WriteString("(")
			dfs(left+1, right, n)
			t := curr.String()[:curr.Len()-1]
			curr.Reset()
			curr.WriteString(t)
		}
		if right < left {
			curr.WriteString(")")
			dfs(left, right+1, n)
			t := curr.String()[:curr.Len()-1]
			curr.Reset()
			curr.WriteString(t)
		}
	}
	dfs(0, 0, n)
	return res
}

// @lc code=end

/*
// @lcpr case=start
// 3\n
// @lcpr case=end

// @lcpr case=start
// 1\n
// @lcpr case=end

*/
