/*
 * @lc app=leetcode.cn id=112 lang=golang
 * @lcpr version=30204
 *
 * [112] 路径总和
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
func hasPathSum(root *TreeNode, targetSum int) bool {
	return dfs(root, 0, targetSum)

}
func dfs(root *TreeNode, sum int, targetSum int) bool {
	if root == nil {
		return false
	}
	sum += root.Val
	if sum == targetSum && root.Left == nil && root.Right == nil {
		return true
	}
	return dfs(root.Left, sum, targetSum) || dfs(root.Right, sum, targetSum)
}

// @lc code=end

/*
// @lcpr case=start
// [5,4,8,11,null,13,4,7,2,null,null,null,1]\n22\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3]\n5\n
// @lcpr case=end

// @lcpr case=start
// []\n0\n
// @lcpr case=end

*/

