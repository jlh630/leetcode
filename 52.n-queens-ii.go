/*
 * @lc app=leetcode.cn id=52 lang=golang
 * @lcpr version=30204
 *
 * [52] N 皇后 II
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func totalNQueens(n int) int {
	chess := make([][]int, 0)
	ans := 0
	isExits := func(i int, j int) bool {
		for _, v := range chess {
			if j == v[1] || math.Abs(float64(i-v[0])) == math.Abs(float64(j-v[1])) {
				return false
			}
		}
		return true
	}
	var dfs func(row int)
	dfs = func(row int) {
		if len(chess) == n {
			ans++
			return
		}
		for j := 0; j < n; j++ {
			if isExits(row, j) {
				chess = append(chess, []int{row, j})
				dfs(row + 1)
				chess = chess[:len(chess)-1]
			}
		}
	}
	dfs(0)
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// 4\n
// @lcpr case=end

// @lcpr case=start
// 1\n
// @lcpr case=end

*/

