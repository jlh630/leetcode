/*
 * @lc app=leetcode.cn id=129 lang=golang
 * @lcpr version=30204
 *
 * [129] 求根节点到叶节点数字之和
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
func sumNumbers(root *TreeNode) int {
	res := 0
	var dfs func(node *TreeNode, sum int)
	dfs = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}
		t := sum*10 + node.Val
		dfs(node.Left, t)
		dfs(node.Right, t)
		if node.Left == nil && node.Right == nil {
			res += t
		}

	}
	dfs(root, 0)
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [4,9,0,5,1]\n
// @lcpr case=end
// @lcpr case=start
// [1,2,3,null,4]\n
// @lcpr case=end

*/

