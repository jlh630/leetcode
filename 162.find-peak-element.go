/*
 * @lc app=leetcode.cn id=162 lang=golang
 * @lcpr version=30204
 *
 * [162] 寻找峰值
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func findPeakElement(nums []int) int {
	n := len(nums)
	left := 0
	right := n - 1
	top := func(curr int) (bool, bool) {
		res := nums[curr]
		l, r := false, false
		if curr-1 >= 0 && nums[curr-1] > res {
			l = true
			res = nums[curr-1]
		}
		if curr+1 < n && nums[curr+1] > res {
			r = true
			l = false
		}
		return l, r
	}
	for left <= right {
		mid := (left + right) / 2
		l, r := top(mid)
		if l == true {
			right = mid - 1
		} else if r == true {
			left = mid + 1
		} else {
			return mid
		}
	}
	return 0
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,1]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,1,3,5,6,4]\n
// @lcpr case=end

*/

