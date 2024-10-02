/*
 * @lc app=leetcode.cn id=37 lang=golang
 * @lcpr version=30204
 *
 * [37] 解数独
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func solveSudoku(board [][]byte) {
	var rows, cols [9][9]bool
	var box [3][3][9]bool
	count := 9 * 9
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			num := board[i][j] - '0' - 1
			rows[i][num] = true
			cols[j][num] = true
			box[i/3][j/3][num] = true
			count--
		}
	}
	isfill := func(row, col, num int) bool {
		if rows[row][num] || cols[col][num] || box[row/3][col/3][num] {
			return false
		}
		return true
	}
	var dfs func(row, col int)
	next := func(row, col int) {
		if col == 8 {
			dfs(row+1, 0)
		} else {
			dfs(row, col+1)
		}
	}
	dfs = func(row, col int) {
		if count == 0 {
			return
		}
		if board[row][col] == '.' {
			for i := 0; i < 9; i++ {
				if isfill(row, col, i) {
					count--
					board[row][col] = byte(i+1) + '0'
					rows[row][i] = true
					cols[col][i] = true
					box[row/3][col/3][i] = true
					next(row, col)
					if count == 0 {
						return
					}
					board[row][col] = '.'
					rows[row][i] = false
					cols[col][i] = false
					box[row/3][col/3][i] = false
					count++
				}
			}
		} else {
			next(row, col)
		}

	}
	dfs(0, 0)
}

// @lc code=end

/*
// @lcpr case=start
// [["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]\n
// @lcpr case=end

*/

