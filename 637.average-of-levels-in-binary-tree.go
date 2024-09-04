/*
 * @lc app=leetcode.cn id=637 lang=golang
 * @lcpr version=30204
 *
 * [637] 二叉树的层平均值
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
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
	count := []int{}
	sums := []int{}
	ans := []float64{}
	var dfs func(int, *TreeNode)
	dfs = func(level int, node *TreeNode) {
		if node == nil {
			return
		}
		if len(sums) == level {
			sums = append(sums, 0)
			count = append(count, 0)
		}
		sums[level] += node.Val
		count[level]++
		dfs(level+1, node.Left)
		dfs(level+1, node.Right)
	}
	dfs(0, root)
	for i := 0; i < len(sums); i++ {
		ans = append(ans, float64(sums[i])/float64(count[i]))
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [3,9,20,null,null,15,7]\n
// @lcpr case=end

// @lcpr case=start
// [3,9,20,15,7]\n
// @lcpr case=end

*/

