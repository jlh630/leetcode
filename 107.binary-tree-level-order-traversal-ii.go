/*
 * @lc app=leetcode.cn id=107 lang=golang
 * @lcpr version=30204
 *
 * [107] 二叉树的层序遍历 II
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
func levelOrderBottom(root *TreeNode) [][]int {
	var dfs func(*TreeNode, int)
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if len(ans) == level {
			ans = append(ans, make([]int, 0))
		}
		ans[level] = append(ans[level], node.Val)
		dfs(node.Left, level+1)
		dfs(node.Right, level+1)
	}
	dfs(root, 0)
	n := len(ans)
	right := len(ans) - 1
	for i := 0; i < n/2; i++ {
		ans[i], ans[right] = ans[right], ans[i]
		right--
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [3,9,20,null,null,15,7]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

*/

