/*
 * @lc app=leetcode.cn id=57 lang=golang
 * @lcpr version=30204
 *
 * [57] 插入区间
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func insert(intervals [][]int, newInterval []int) [][]int {
	ans := make([][]int, 0)
	left, right := newInterval[0], newInterval[1]
	i := 0
	n := len(intervals)
	//左边不重叠的
	for i < n && intervals[i][1] < left {
		ans = append(ans, intervals[i])
		i++
	}
	//重叠部分
	for i < n && intervals[i][0] <= right {
		//并集
		left = min(left, intervals[i][0])
		right = max(right, intervals[i][1])
		i++
	}
	ans = append(ans, []int{left, right})
	//右边不重叠
	for i < n {
		ans = append(ans, intervals[i])
		i++
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [[1,3],[6,9]]\n[2,5]\n
// @lcpr case=end

// @lcpr case=start
// [[1,2],[3,5],[6,7],[8,10],[12,16]]\n[4,8]\n
// @lcpr case=end

*/

