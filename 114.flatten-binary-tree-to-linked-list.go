/*
 * @lc app=leetcode.cn id=114 lang=golang
 * @lcpr version=30204
 *
 * [114] 二叉树展开为链表
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
func flatten(root *TreeNode) {
	var dfs func(*TreeNode) *TreeNode
	dfs = func(root *TreeNode) *TreeNode {
		if root == nil {
			return nil
		}
		next := dfs(root.Left)
		next_next := dfs(root.Right)
		if next != nil {
			next.Left = nil
			root.Right = next
			for next.Right != nil {
				next = next.Right
			}
			next.Right = next_next
		}
		if next_next != nil {
			next_next.Left = nil
		}
		root.Left = nil
		return root
	}
	dfs(root)
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,5,3,4,null,6]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

// @lcpr case=start
// [0]\n
// @lcpr case=end

*/

