/*
 * @lc app=leetcode.cn id=21 lang=c
 * @lcpr version=30204
 *
 * [21] 合并两个有序链表
 */


// @lcpr-template-start

// @lcpr-template-end
#include<stdio.h>
// @lc code=start
/**
 * Definition for singly-linked list.*/
//   struct ListNode {
//       int val;
//       struct ListNode *next;
//   };
 /**/
struct ListNode* mergeTwoLists(struct ListNode* list1, struct ListNode* list2) {
    struct ListNode head;
    struct ListNode* node=&head;
    while (list1!=NULL && list2!=NULL)
    {
        if (list1->val < list2->val)
        {
            node->next=list1;
            list1=list1->next;
        }else{
            node->next=list2;
            list2=list2->next;
        }
        node=node->next;
    }
    node->next=list1==NULL? list2:list1;
    return head.next;
}
// @lc code=end



/*
// @lcpr case=start
// [1,2,4]\n[1,3,4]\n
// @lcpr case=end

// @lcpr case=start
// []\n[]\n
// @lcpr case=end

// @lcpr case=start
// []\n[0]\n
// @lcpr case=end

 */

