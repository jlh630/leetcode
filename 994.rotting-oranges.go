/*
 * @lc app=leetcode.cn id=994 lang=golang
 * @lcpr version=30204
 *
 * [994] 腐烂的橘子
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
// bfs
func orangesRotting(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])
	fresh := 0
	//队列
	bad_oranges := [][]int{}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				fresh++
			} else if grid[i][j] == 2 {
				bad_oranges = append(bad_oranges, []int{i, j})
			}
		}
	}
	if len(bad_oranges) == 0 && fresh == 0 {
		return 0
	}
	res := -1
	bad_fun := func(i, j int) {
		bad_oranges = append(bad_oranges, []int{i, j})
		grid[i][j] = 2
		fresh--
	}
	for len(bad_oranges) > 0 {
		level := len(bad_oranges)
		res++
		for count := 0; count < level; count++ {
			i, j := bad_oranges[0][0], bad_oranges[0][1]
			if i-1 >= 0 && grid[i-1][j] == 1 {
				bad_fun(i-1, j)
			}
			if j-1 >= 0 && grid[i][j-1] == 1 {
				bad_fun(i, j-1)
			}
			if i+1 < n && grid[i+1][j] == 1 {
				bad_fun(i+1, j)
			}
			if j+1 < m && grid[i][j+1] == 1 {
				bad_fun(i, j+1)
			}
			bad_oranges = bad_oranges[1:]
		}
	}
	if fresh != 0 {
		res = -1
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [[2,1,1],[1,1,0],[0,1,1]]\n
// @lcpr case=end

// @lcpr case=start
// [[2,1,1],[0,1,1],[1,0,1]]\n
// @lcpr case=end

// @lcpr case=start
// [[0,2]]\n
// @lcpr case=end

*/

