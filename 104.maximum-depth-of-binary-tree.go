/*
 * @lc app=leetcode.cn id=104 lang=golang
 * @lcpr version=30204
 *
 * [104] 二叉树的最大深度
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
func maxDepth(root *TreeNode) int {
	var dfs func(*TreeNode, int)
	res := 0

	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
		res = max(res, level)
	}
	dfs(root, 1)
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [3,9,20,null,null,15,7]\n
// @lcpr case=end

// @lcpr case=start
// [1,null,2]\n
// @lcpr case=end

*/

