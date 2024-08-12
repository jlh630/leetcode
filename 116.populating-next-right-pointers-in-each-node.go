/*
 * @lc app=leetcode.cn id=116 lang=golang
 * @lcpr version=30204
 *
 * [116] 填充每个节点的下一个右侧节点指针
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
//广度优先
// func connect(root *Node) *Node {
// 	if root == nil {
// 		return nil
// 	}
// 	queue := []*Node{root}
// 	for len(queue) > 0 {
// 		n := len(queue)
// 		var pre *Node = nil
// 		for i := 0; i < n; i++ {
// 			curr := queue[0]
// 			queue = queue[1:]
// 			if pre != nil {
// 				pre.Next = curr
// 			}
// 			if curr.Left != nil {
// 				queue = append(queue, curr.Left)
// 			}
// 			if curr.Right != nil {
// 				queue = append(queue, curr.Right)
// 			}
// 			pre = curr
// 		}
// 	}
// 	return root
// }
//深度优先
func connect(root *Node) *Node {
	var dfs func(root *Node)
	dfs = func(root *Node) {
		if root == nil || root.Left == nil || root.Right == nil {
			return
		}
		root.Left.Next = root.Right
		if root.Next != nil {
			root.Right.Next = root.Next.Left
		}
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return root
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5,6,7]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

*/

