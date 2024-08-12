/*
 * @lc app=leetcode.cn id=148 lang=golang
 * @lcpr version=30204
 *
 * [148] 排序链表
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
func sortList(head *ListNode) *ListNode {
	var partition func(*ListNode) *ListNode
	var merge func(*ListNode, *ListNode) *ListNode
	merge = func(l1 *ListNode, l2 *ListNode) *ListNode {
		var thread = &ListNode{Val: -1}
		node := thread
		for l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				node.Next = l1
				l1 = l1.Next
			} else {
				node.Next = l2
				l2 = l2.Next
			}
			node = node.Next
		}
		if l1 != nil {
			node.Next = l1
		} else {
			node.Next = l2
		}
		return thread.Next
	}
	partition = func(left *ListNode) *ListNode {
		if left == nil {
			return nil
		}
		if left.Next == nil {
			return left
		}
		slow := left
		fast := left
		for fast.Next != nil && fast.Next.Next != nil {
			slow = slow.Next
			fast = fast.Next.Next
		}
		if fast.Next == nil {
			fast.Next = nil
		} else {
			fast.Next.Next = nil
		}
		mid := slow.Next
		slow.Next = nil
		l1 := partition(left)
		l2 := partition(mid)
		return merge(l1, l2)
	}
	return partition(head)

}

// @lc code=end

/*
// @lcpr case=start
// [4,2,1,3]\n
// @lcpr case=end

// @lcpr case=start
// [-1,5,3,4,0]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

*/

