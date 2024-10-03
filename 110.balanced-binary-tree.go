/*
 * @lc app=leetcode.cn id=110 lang=golang
 * @lcpr version=30204
 *
 * [110] 平衡二叉树
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
func isBalanced(root *TreeNode) bool {
	var dfs func(*TreeNode) (int, bool)
	dfs = func(node *TreeNode) (int, bool) {
		if node == nil {
			return 0, true
		}
		lHeight, lok := dfs(node.Left)
		rHeight, rok := dfs(node.Right)
		if lok && rok {
			if math.Abs(float64(lHeight-rHeight)) > 1 {
				return 0, false
			}
			return max(lHeight, rHeight) + 1, true
		} else {
			return 0, false
		}
	}
	_, ans := dfs(root)
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [3,9,20,null,null,15,7]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,2,3,3,null,null,4,4]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

*/

