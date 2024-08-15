/*
 * @lc app=leetcode.cn id=230 lang=golang
 * @lcpr version=30204
 *
 * [230] 二叉搜索树中第K小的元素
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
func kthSmallest(root *TreeNode, k int) int {
	res := -1
	num := 0
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		num++
		if k == num {
			res = root.Val
		}
		if res != -1 {
			return
		}
		dfs(root.Right)
	}
	dfs(root)
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [3,1,4,null,2]\n1\n
// @lcpr case=end

// @lcpr case=start
// [5,3,6,2,4,null,null,1]\n3\n
// @lcpr case=end

*/

