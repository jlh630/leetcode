/*
 * @lc app=leetcode.cn id=199 lang=golang
 * @lcpr version=30204
 *
 * [199] 二叉树的右视图
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

// 广度优先
func rightSideView(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			curr := queue[0]
			queue = queue[1:]
			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
			if i == n-1 {
				res = append(res, curr.Val)
			}
		}

	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,null,5,null,4]\n
// @lcpr case=end

// @lcpr case=start
// [1,null,3]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

*/

