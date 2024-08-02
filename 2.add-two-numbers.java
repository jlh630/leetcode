/*
 * @lc app=leetcode.cn id=2 lang=java
 * @lcpr version=30204
 *
 * [2] 两数相加
 */


// @lcpr-template-start

// @lcpr-template-end
// @lc code=start

import java.util.ArrayList;
import java.util.List;

/**
 * Definition for singly-linked list.
 */ 
//  class ListNode {
//       int val;
//       ListNode next;
//       ListNode() {}
//       ListNode(int val) { this.val = val; }
//       ListNode(int val, ListNode next) { this.val = val; this.next = next; }
//   }
 
class Solution {
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        ListNode thead=new ListNode(0,null);
        ListNode preNode=thead;
        int pre=0;
        while (l1!=null&&l2!=null) {
            int sum=pre+l1.val+l2.val;
            pre=sum/10;
            ListNode node=new ListNode(sum%10,null);
            preNode.next=node;
            preNode=node;

            l1=l1.next;
            l2=l2.next;
        }
        while (l1!=null) {
            int sum=pre+l1.val;
            pre=sum/10;
            ListNode node=new ListNode(sum%10,null);
            preNode.next=node;
            preNode=node;
            l1=l1.next;
        }
        while (l2!=null) {
            int sum=pre+l2.val;
            pre=sum/10;
            ListNode node=new ListNode(sum%10,null);
            preNode.next=node;
            preNode=node;
            l2=l2.next;
        }
        if (pre!=0) {
            ListNode node=new ListNode(pre,null);
            preNode.next=node;
        }
        return thead.next;
    }

}
// @lc code=end



/*
// @lcpr case=start
// [2,4,3]\n[5,6,4]\n
// @lcpr case=end

// @lcpr case=start
// [0]\n[0]\n
// @lcpr case=end

// @lcpr case=start
// [9,9,9,9,9,9,9]\n[9,9,9,9]\n
// @lcpr case=end

 */

