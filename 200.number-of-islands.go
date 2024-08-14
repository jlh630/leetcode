/*
 * @lc app=leetcode.cn id=200 lang=golang
 * @lcpr version=30204
 *
 * [200] 岛屿数量
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
package main

// 可以加入一个字典记录扫过的地方记录一下优化
func numIslands(grid [][]byte) int {
	n := len(grid)
	m := len(grid[0])
	res := 0
	var dfs func(int, int, int, int)
	dfs = func(i, j, addi, addj int) {
		ni, nj := i+addi, j+addj
		if ni >= 0 && nj >= 0 && ni < n && nj < m && grid[ni][nj] == '1' {
			grid[ni][nj] = '0'
		} else {
			return
		}
		dfs(ni, nj, 1, 0)
		dfs(ni, nj, -1, 0)
		dfs(ni, nj, 0, 1)
		dfs(ni, nj, 0, -1)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '1' {
				res++
				dfs(i, j, 0, 0)
			}
		}
	}

	return res
}

// @lc code=end

/*
// @lcpr case=start
// [["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]]\n
// @lcpr case=end

// @lcpr case=start
// [["1","1","0","0","0"],["1","1","0","0","0"],["0","0","1","0","0"],["0","0","0","1","1"]]\n
// @lcpr case=end

*/
