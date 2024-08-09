/*
 * @lc app=leetcode.cn id=56 lang=golang
 * @lcpr version=30204
 *
 * [56] 合并区间
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	result := make([][]int, 0)
	for i := 0; i < len(intervals); i++ {
		left := intervals[i][0]
		right := intervals[i][1]

		for i+1 < len(intervals) && right >= intervals[i+1][0] {
			right = max(right, intervals[i+1][1])
			i++
		}
		result = append(result, []int{left, right})
	}
	return result
}

// @lc code=end

/*
// @lcpr case=start
// [[1,3],[2,6],[8,10],[15,18]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,4],[4,5]]\n
// @lcpr case=end

*/

