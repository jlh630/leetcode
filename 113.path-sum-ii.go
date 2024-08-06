/*
 * @lc app=leetcode.cn id=113 lang=golang
 * @lcpr version=30204
 *
 * [113] 路径总和 II
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

func pathSum(root *TreeNode, targetSum int) [][]int {
	var res = make([][]int, 0)

	dfs(root, &res, targetSum, 0, make([]int, 0))

	return res
}
func dfs(root *TreeNode, res *[][]int, targetSum int, sum int, path []int) {
	if root == nil {
		return
	}
	sum += root.Val
	path = append(path, root.Val)

	if sum == targetSum && root.Right == nil && root.Left == nil {
		*res = append(*res, append([]int{}, path...))
		return
	}
	dfs(root.Left, res, targetSum, sum, path)
	dfs(root.Right, res, targetSum, sum, path)
	path = path[:len(path)-1]
}

// @lc code=end

/*
// @lcpr case=start
// [5,4,8,11,null,13,4,7,2,null,null,5,1]\n22\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3]\n5\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n0\n
// @lcpr case=end

*/

