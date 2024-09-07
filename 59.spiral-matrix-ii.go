/*
 * @lc app=leetcode.cn id=59 lang=golang
 * @lcpr version=30204
 *
 * [59] 螺旋矩阵 II
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func generateMatrix(n int) [][]int {
	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	ans[0][0] = 1
	if n == 1 {
		return ans
	}
	pre := 1
	curri, currj := 0, 0
	fill := func(addi, addj int) {
		tmpi := curri + addi
		tmpj := currj + addj
		for tmpi >= 0 && tmpi < n && tmpj >= 0 && tmpj < n && ans[tmpi][tmpj] == 0 {
			curri = tmpi
			currj = tmpj
			pre++
			ans[curri][currj] = pre
			tmpi = curri + addi
			tmpj = currj + addj
		}
	}
	for pre < n*n {
		fill(0, 1)
		fill(1, 0)
		fill(0, -1)
		fill(-1, 0)
	}
	return ans
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

