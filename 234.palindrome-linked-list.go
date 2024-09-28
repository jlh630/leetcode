/*
 * @lc app=leetcode.cn id=234 lang=golang
 * @lcpr version=30204
 *
 * [234] 回文链表
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

func isPalindrome(head *ListNode) bool {
	slow := head
	fast := head
	//1.找到回文的中间节点
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	//2.反转链表
	returnList := func(node *ListNode) *ListNode {
		var pre *ListNode = nil
		for node != nil {
			next := node.Next
			node.Next = pre
			pre = node
			node = next
		}
		return pre
	}
	l1 := head
	l2 := returnList(slow.Next)
	t := l2
	res := true
	for res && t != nil {
		if t.Val != l1.Val {
			res = false
		}
		t = t.Next
		l1 = l1.Next
	}
	//恢复链表
	returnList(l2)
	return res

}

// @lc code=end

/*
// @lcpr case=start
// [1,2,2,1]\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n
// @lcpr case=end

*/
