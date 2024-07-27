/*
 * @lc app=leetcode.cn id=141 lang=c
 * @lcpr version=30204
 *
 * [141] 环形链表
 */


// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
#include<stdio.h>
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
bool hasCycle(struct ListNode *head) {
    if (head==NULL)
    {
        return false;
    }
    
    struct ListNode * slow=head;
    struct ListNode * fast=head->next;
    while (slow!=fast)
    {
        if (fast==NULL||fast->next==NULL)
        {
            return false;
        }
        slow=slow->next;
        fast=fast->next->next;
        
    }
    return true;
    
}
// @lc code=end



/*
// @lcpr case=start
// [3,2,0,-4]\n1\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n0\n
// @lcpr case=end

// @lcpr case=start
// [1]\n-1\n
// @lcpr case=end

 */

