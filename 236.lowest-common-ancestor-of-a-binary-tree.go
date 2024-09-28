/*
 * @lc app=leetcode.cn id=236 lang=golang
 * @lcpr version=30204
 *
 * [236] 二叉树的最近公共祖先
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

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	/**
	    1. 当前是空节点直接返回当前节点
		2. q节点、q节点直接返回当前节点
		3. 左右子树都有（当前节点就是最近的公共祖先） 返回当前节点
		4. 左子树、右子树 有(当前节点的左、右就是最近的公共祖先) 返回左节点、右节点
		5. 左、右子树都没有返回空节点就好了
	*/

	if root == nil || root == q || root == p {
		return root
	}
	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)
	if l != nil && r != nil {
		return root
	}
	if l != nil {
		return l
	} else {
		//两边都没的反正都是返回nil
		//右边有
		return r
	}

}

// @lc code=end

/*
// @lcpr case=start
// [3,5,1,6,2,0,8,null,null,7,4]\n5\n1\n
// @lcpr case=end

// @lcpr case=start
// [3,5,1,6,2,0,8,null,null,7,4]\n5\n4\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n1\n2\n
// @lcpr case=end

*/
