/*
 * @lc app=leetcode.cn id=45 lang=golang
 * @lcpr version=30204
 *
 * [45] 跳跃游戏 II
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func jump(nums []int) int {
	step := 0
	end := 0
	far := 0
	for i := 0; i < len(nums)-1; i++ {
		far = max(far, i+nums[i])
		if i == end {
			end = far
			step++
		}
	}
	return step
}

// @lc code=end

/*
// @lcpr case=start
// [2,3,1,1,4]\n
// @lcpr case=end

// @lcpr case=start
// [2,3,0,1,4]\n
// @lcpr case=end

*/

