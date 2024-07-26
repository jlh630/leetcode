/*
 * @lc app=leetcode.cn id=206 lang=c
 * @lcpr version=30204
 *
 * [206] 反转链表
 */


// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
/**
  Definition for singly-linked list.*/
// #include <stdio.h>
//   struct ListNode {
//       int val;
//       struct ListNode *next;
//   };
  /** */
struct ListNode* reverseList(struct ListNode* node) {
    if (node ==NULL)
    {
        return node;
    }
     struct ListNode* pre=NULL;
     
     while(node!=NULL){
        struct ListNode* next= node->next;
        node->next=pre;
        pre= node;
        node= next;
     }

     return pre;
}


// @lc code=end



/*
// @lcpr case=start
// [1,2,3,4,5]\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

 */

