/*
 * @lc app=leetcode.cn id=77 lang=golang
 * @lcpr version=30204
 *
 * [77] 组合
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
package main

func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	var dfs func(start int)
	dfs = func(start int) {
		if len(path) == k {
			res = append(res, append([]int{}, path...))
			return
		}
		for i := start; i <= n; i++ {
			path = append(path, i)
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(1)
	return res
}

// @lc code=end

/*
// @lcpr case=start
// 4\n2\n
// @lcpr case=end

// @lcpr case=start
// 1\n1\n
// @lcpr case=end

*/
