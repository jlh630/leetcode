/*
 * @lc app=leetcode.cn id=130 lang=golang
 * @lcpr version=30204
 *
 * [130] 被围绕的区域
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func solve(board [][]byte) {
	n := len(board)
	m := len(board[0])
	road := make([][]bool, n)
	for i := range road {
		road[i] = make([]bool, m)
	}

	var dfs func(curri, curry int)
	dfs = func(curri, curry int) {
		if curri < 0 || curri >= n || curry < 0 || curry >= m || board[curri][curry] == 'X' || road[curri][curry] == true {
			return
		}
		road[curri][curry] = true
		dfs(curri+1, curry)
		dfs(curri-1, curry)
		dfs(curri, curry-1)
		dfs(curri, curry+1)
	}

	for j := 0; j < m; j++ {
		dfs(0, j)
		dfs(n-1, j)
	}
	for i := 1; i < n-1; i++ {
		dfs(i, 0)
		dfs(i, m-1)
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if !road[i][j] && board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

// @lc code=end

/*
// @lcpr case=start
// [["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]\n
// @lcpr case=end

// @lcpr case=start
// [["X"]]\n
// @lcpr case=end

// @lcpr case=start
// [["O","O"],["O","O"]]\n
// @lcpr case=end
*/

