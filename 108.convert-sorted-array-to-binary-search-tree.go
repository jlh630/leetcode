/*
 * @lc app=leetcode.cn id=108 lang=golang
 * @lcpr version=30204
 *
 * [108] 将有序数组转换为二叉搜索树
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {

	var merge func([]int) *TreeNode
	merge = func(nums []int) *TreeNode {
		if len(nums) == 0 {
			return nil
		}
		var root TreeNode
		mid := len(nums) / 2
		root.Val = nums[mid]

		root.Left = merge(nums[:mid])
		root.Right = merge(nums[mid+1:])
		return &root
	}
	return merge(nums)
}

// @lc code=end

/*
// @lcpr case=start
// [-10,-3,0,5,9]\n
// @lcpr case=end

// @lcpr case=start
// [1,3]\n
// @lcpr case=end

*/

