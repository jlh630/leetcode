/*
 * @lc app=leetcode.cn id=543 lang=golang
 * @lcpr version=30204
 *
 * [543] 二叉树的直径
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
package main

func diameterOfBinaryTree(root *TreeNode) int {
	var dfs func(root *TreeNode) int
	res := 0
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		l_len := dfs(root.Left)
		r_len := dfs(root.Right)
		res = max(l_len+r_len, res)
		return max(l_len, r_len) + 1
	}
	dfs(root)
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5]\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n
// @lcpr case=end

*/
