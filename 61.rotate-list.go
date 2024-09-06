/*
 * @lc app=leetcode.cn id=61 lang=golang
 * @lcpr version=30204
 *
 * [61] 旋转链表
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
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	size := 1
	node := head
	for node.Next != nil {
		size++
		node = node.Next
	}
	step := size - k%size
	if step == size {
		return head
	}
	node.Next = head
	for step != 0 {
		node = node.Next
		step--
	}
	ans := node.Next
	node.Next = nil
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5]\n2\n
// @lcpr case=end

// @lcpr case=start
// [0,1,2]\n4\n
// @lcpr case=end

*/

