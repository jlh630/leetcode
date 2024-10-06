/*
 * @lc app=leetcode.cn id=529 lang=golang
 * @lcpr version=30204
 *
 * [529] 扫雷游戏
 */

// @lc code=start
func updateBoard(board [][]byte, click []int) [][]byte {
	row, col := click[0], click[1]
	n, m := len(board), len(board[0])
	if board[row][col] == 'M' {
		board[row][col] = 'X'
		return board
	}
	flag := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

	around := func(row, col int) int {
		if row >= 0 && col >= 0 && row < n && col < m && board[row][col] == 'M' {
			return 1
		}
		return 0
	}
	var dfs func(row, col int)
	dfs = func(row, col int) {
		if row < 0 || col < 0 || row >= n || col >= m || board[row][col] != 'E' {
			return
		}
		count := 0
		for i := range flag {
			count += around(row+flag[i][0], col+flag[i][1])
		}
		if count > 0 {
			board[row][col] = byte(count) + '0'
			return
		} else {
			board[row][col] = 'B'
			for i := range flag {
				dfs(row+flag[i][0], col+flag[i][1])
			}
		}
	}
	dfs(row, col)
	return board
}

// @lc code=end

/*
// @lcpr case=start
// [["E","E","E","E","E"],["E","E","M","E","E"],["E","E","E","E","E"],["E","E","E","E","E"]]\n[3,0]\n
// @lcpr case=end

// @lcpr case=start
// [["B","1","E","1","B"],["B","1","M","1","B"],["B","1","1","1","B"],["B","B","B","B","B"]]\n[1,2]\n
// @lcpr case=end

// @lcpr case=start
// [["E","E","E","E","E","E","E","E"],["E","E","E","E","E","E","E","M"],["E","E","M","E","E","E","E","E"],["M","E","E","E","E","E","E","E"],["E","E","E","E","E","E","E","E"],["E","E","E","E","E","E","E","E"],["E","E","E","E","E","E","E","E"],["E","E","M","M","E","E","E","E"]]\n[0,0]\n
// @lcpr case=end
*/

