/*
 * @lc app=leetcode.cn id=19 lang=c
 * @lcpr version=30204
 *
 * [19] 删除链表的倒数第 N 个结点
 */


// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
/**
 * Definition for singly-linked list.
 */
#include<stdio.h>
//  struct ListNode {
//       int val;
//       struct ListNode *next;
//   };
//双指针
struct ListNode* removeNthFromEnd(struct ListNode* node, int n) {
    struct ListNode* pre=(struct ListNode*)malloc(sizeof(struct ListNode));
    pre->next=node;
    struct ListNode* fast =node;
    struct ListNode* slow =pre;
    for (size_t i = 0; i < n; i++)
    {
        fast=fast->next;
    }
    while (fast!=NULL)
    {
        fast=fast->next;
        slow=slow->next;
    }
    slow->next=slow->next->next;
    struct ListNode* res=pre->next;
    free(pre);
    
    return res;
}
// @lc code=end



/*
// @lcpr case=start
// [1,2,3,4,5]\n2\n
// @lcpr case=end

// @lcpr case=start
// [1]\n1\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n1\n
// @lcpr case=end

 */

