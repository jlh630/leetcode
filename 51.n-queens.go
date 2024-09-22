/*
 * @lc app=leetcode.cn id=51 lang=golang
 * @lcpr version=30204
 *
 * [51] N 皇后
 */

// @lc code=start
func solveNQueens(n int) [][]string {
	ans := make([][]string, 0)
	use := make([][]int, 0)
	var dfs func(row int)
	isPut := func(curri, currj int) bool {
		for _, v := range use {
			if currj == v[1] || math.Abs(float64(curri-v[0])) == math.Abs(float64(currj-v[1])) {
				return false
			}
		}
		return true
	}
	curr := make([]string, 0)
	dfs = func(row int) {
		if len(use) == n {
			ans = append(ans, append([]string(nil), curr...))
			return
		}
		tmp := make([]byte, n)
		for i := 0; i < n; i++ {
			tmp[i] = '.'
		}
		for j := 0; j < n; j++ {
			if isPut(row, j) {
				tmp[j] = 'Q'
				curr = append(curr, string(tmp))
				use = append(use, []int{row, j})
				dfs(row + 1)
				curr = curr[:len(curr)-1]
				use = use[:len(use)-1]
				tmp[j] = '.'
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

