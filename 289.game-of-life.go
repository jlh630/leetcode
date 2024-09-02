/*
 * @lc app=leetcode.cn id=289 lang=golang
 * @lcpr version=30204
 *
 * [289] 生命游戏
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func gameOfLife(board [][]int) {
	// -1 活变死  1 活  0死  2死变活
	n := len(board)
	m := len(board[0])
	currLifeNum := func(curri, currj int) int {
		count := 0
		arr := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
		for i := 0; i < 8; i++ {
			addi := curri + arr[i][0]
			addj := currj + arr[i][1]
			if addi >= 0 && addj >= 0 && addi < n && addj < m && (board[addi][addj] == -1 || board[addi][addj] == 1) {
				count++
			}
		}
		return count
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			count := currLifeNum(i, j)
			if board[i][j] == 1 && (count < 2 || count > 3) {
				board[i][j] = -1
			}
			if board[i][j] == 0 && count == 3 {
				board[i][j] = 2
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] > 0 {
				board[i][j] = 1
			} else {
				board[i][j] = 0
			}
		}
	}
}

// @lc code=end

/*
// @lcpr case=start
// [[0,1,0],[0,0,1],[1,1,1],[0,0,0]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,1],[1,0]]\n
// @lcpr case=end

*/

