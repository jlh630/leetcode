/*
 * @lc app=leetcode.cn id=63 lang=golang
 * @lcpr version=30204
 *
 * [63] 不同路径 II
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {

	n := len(obstacleGrid)
	m := len(obstacleGrid[0])
	if obstacleGrid[0][0] == 1 || obstacleGrid[n-1][m-1] == 1 {
		return 0
	}
	dp := make([]int, m)
	dp[0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {

			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
				continue
			}
			if j != 0 && obstacleGrid[i][j-1] == 0 {
				dp[j] += dp[j-1]
			}
		}
	}

	return dp[m-1]
}

// @lc code=end

/*
// @lcpr case=start
// [[0,0,0],[0,1,0],[0,0,0]]\n
// @lcpr case=end

// @lcpr case=start
// [[0,1],[0,0]]\n
// @lcpr case=end

// @lcpr case=start
// [[0,0],[0,1]]\n
// @lcpr case=end

*/
