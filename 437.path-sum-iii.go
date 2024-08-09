/*
 * @lc app=leetcode.cn id=437 lang=golang
 * @lcpr version=30204
 *
 * [437] 路径总和 III
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
func pathSum(root *TreeNode, targetSum int) int {
	var dfs func(*TreeNode, int)
	var dfsii func(*TreeNode)
	var sum = 0
	dfs = func(root *TreeNode, curNum int) {
		if root == nil {
			return
		}
		curNum += root.Val
		if curNum == targetSum {
			sum++
		}

		dfs(root.Left, curNum)
		dfs(root.Right, curNum)

	}
	dfsii = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root, 0)
		dfsii(root.Left)
		dfsii(root.Right)

	}
	dfsii(root)
	return sum
}

// @lc code=end

/*
// @lcpr case=start
// [10,5,-3,3,2,null,11,3,-2,null,1]\n8\n
// @lcpr case=end

// @lcpr case=start
// [5,4,8,11,null,13,4,7,2,null,null,5,1]\n22\n
// @lcpr case=end

*/

