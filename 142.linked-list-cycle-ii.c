/*
 * @lc app=leetcode.cn id=142 lang=c
 * @lcpr version=30204
 *
 * [142] 环形链表 II
 */


// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
/**
 * Definition for singly-linked list.
 */
//  struct ListNode {
//       int val;
//       struct ListNode *next;
//   };
 //双指针
struct ListNode *detectCycle(struct ListNode *head) {

    struct ListNode* slow=head;
    struct ListNode* fast=head;
    while (fast!=NULL&&fast->next!=NULL)
    {
        fast=fast->next->next;
        slow=slow->next;

        if (slow==fast)
        {
            struct ListNode* ptr=head;
            while (ptr!=slow)
            {
                ptr=ptr->next;
                slow=slow->next;
            }
            return ptr;
        }

    }
    return NULL;
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

