/*
 * @lc app=leetcode.cn id=98 lang=golang
 * @lcpr version=30204
 *
 * [98] 验证二叉搜索树
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

import "math"

func isValidBST(root *TreeNode) bool {
	var dfs func(*TreeNode, int, int) bool
	dfs = func(root *TreeNode, minNum int, maxNum int) bool {
		if root == nil {
			return true
		}
		if root.Val <= minNum || root.Val >= maxNum {
			return false
		}
		return dfs(root.Left, minNum, root.Val) && dfs(root.Right, root.Val, maxNum)
	}
	return dfs(root, math.MinInt32-1, math.MaxInt32+1)
}

// @lc code=end

/*
// @lcpr case=start
// [2,1,3]\n
// @lcpr case=end

// @lcpr case=start
// [5,1,4,null,null,3,6]\n
// @lcpr case=end

*/
