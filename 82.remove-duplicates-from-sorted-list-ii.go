/*
 * @lc app=leetcode.cn id=82 lang=golang
 * @lcpr version=30204
 *
 * [82] 删除排序链表中的重复元素 II
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
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	thead := &ListNode{0, head}
	fast := head
	slow := thead

	for fast != nil {
		if fast.Next != nil && fast.Val == fast.Next.Val {
			for fast.Next != nil && fast.Val == fast.Next.Val {
				fast = fast.Next
			}
			slow.Next = fast.Next

		} else {
			slow = slow.Next

		}
		fast = fast.Next
	}
	return thead.Next
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,3,4,4,5]\n
// @lcpr case=end

// @lcpr case=start
// [1,1,1,2,3]\n
// @lcpr case=end

*/

