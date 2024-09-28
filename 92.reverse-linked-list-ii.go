/*
 * @lc app=leetcode.cn id=92 lang=golang
 * @lcpr version=30204
 *
 * [92] 反转链表 II
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package main

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	thead := &ListNode{Val: -1}
	thead.Next = head
	pre := thead

	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	node := pre.Next
	for i := 0; i < right-left; i++ {
		next := node.Next
		node.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return thead.Next
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5]\n2\n4\n
// @lcpr case=end

// @lcpr case=start
// [5]\n1\n1\n
// @lcpr case=end

*/
