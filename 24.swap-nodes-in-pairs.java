/*
 * @lc app=leetcode.cn id=24 lang=java
 * @lcpr version=30204
 *
 * [24] 两两交换链表中的节点
 */


// @lcpr-template-start

// @lcpr-template-end
// @lc code=start

import java.util.List;

/**
 * Definition for singly-linked list.
 */
//    class ListNode {
//      int val;
//       ListNode next;
//       ListNode() {}
//       ListNode(int val) { this.val = val; }
//       ListNode(int val, ListNode next) { this.val = val; this.next = next; }
//   }
  
      

class Solution {
    public ListNode swapPairs(ListNode head) {
        if (head==null||head.next==null) {
            return head;
        }
        ListNode listHead=new ListNode(0);
        listHead.next=head;
        ListNode pre=listHead;
        ListNode node=head;
        while (node!=null&&node.next!=null) {
            ListNode next=node.next;
            node.next = next.next;
            next.next = node;
            pre.next = next;
            
            pre = node;
            node = node.next;
        }
        return listHead.next;
        
    }
}
// @lc code=end



/*
// @lcpr case=start
// [1,2,3,4]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

 */

