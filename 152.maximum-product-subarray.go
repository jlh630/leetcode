/*
 * @lc app=leetcode.cn id=152 lang=golang
 * @lcpr version=30204
 *
 * [152] 乘积最大子数组
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maxProduct(nums []int) int {
	ans := nums[0]
	pre := nums[0]
	n := len(nums)
	for i := 1; i < n; i++ {
		if pre*nums[i] > pre {
			ans = max(ans, pre*nums[i])
		} else {
			pre = nums[i]
		}

	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [2,3,-2,4]\n
// @lcpr case=end

// @lcpr case=start
// [-2,0,-1]\n
// @lcpr case=end

*/

