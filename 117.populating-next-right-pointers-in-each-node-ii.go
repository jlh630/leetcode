/*
 * @lc app=leetcode.cn id=117 lang=golang
 * @lcpr version=30204
 *
 * [117] 填充每个节点的下一个右侧节点指针 II
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	//空间复杂度 O(1)
	start := root
	for start != nil {
		var nextStart, last *Node
		connected := func(curr *Node) {
			if curr == nil {
				return
			}
			if nextStart == nil {
				nextStart = curr
			}
			if last != nil {
				last.Next = curr
			}
			last = curr
		}
		for node := start; node != nil; node = node.Next {
			connected(node.Left)
			connected(node.Right)
		}
		start = nextStart

	}
	return root
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5,null,7]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

*/

