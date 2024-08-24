/*
 * @lc app=leetcode.cn id=47 lang=golang
 * @lcpr version=30204
 *
 * [47] 全排列 II
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	used := make([]bool, n)
	res := make([][]int, 0)
	path := make([]int, 0)
	var dfs func()
	dfs = func() {
		if len(path) == n {
			res = append(res, append([]int{}, path...))
			return
		}
		for i := 0; i < n; i++ {
			if used[i] || (i > 0 && nums[i] == nums[i-1] && !used[i-1]) {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			dfs()
			used[i] = false
			path = path[:len(path)-1]
		}
	}
	dfs()
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [1,1,2]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

*/

