/*
 * @lc app=leetcode.cn id=160 lang=java
 * @lcpr version=30204
 *
 * [160] 相交链表
 */


// @lcpr-template-start

// @lcpr-template-end
// @lc code=start

import java.util.HashMap;
import java.util.HashSet;
import java.util.Map;
import java.util.Set;


//  class ListNode {
//       int val;
//       ListNode next;
//       ListNode(int x) {
//           val = x;
//           next = null;
//       }
//   }

public class Solution {
    public ListNode getIntersectionNode(ListNode headA, ListNode headB) {
        //case1:hash
    //     Set<ListNode> set=new HashSet<>();
    //     while (headA!=null) {
    //         set.add(headA);
    //         headA=headA.next;
    //     }
    //     while (headB!=null) {
    //         if (set.contains(headB)) {
    //             return headB;
    //         }
    //         headB=headB.next;
    //     }
    //     return null;
        //case2:双指针
        //a：A全长 b: B全长 c:公共部分
        //A---------\
        //          node---------
        //  B-------/
        //   a+（b-c）=b+（a-c）
        //
        ListNode a=headA,b=headB;
        while (a!=b) {
            a=a!=null?a.next:headB;
            b=b!=null?b.next:headA;

        }
        return a;

}

}
// @lc code=end



/*
// @lcpr case=start
// 8\n[4,1,8,4,5]\n[5,6,1,8,4,5]\n2\n3\n
// @lcpr case=end

// @lcpr case=start
// 2\n[1,9,1,2,4]\n[3,2,4]\n3\n1\n
// @lcpr case=end

// @lcpr case=start
// 0\n[2,6,4]\n[1,5]\n3\n2\n
// @lcpr case=end

 */

