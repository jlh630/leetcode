/*
 * @lc app=leetcode.cn id=101 lang=c
 * @lcpr version=30204
 *
 * [101] 对称二叉树
 */


// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
/**
 * Definition for a binary tree node.

 */
#include<stdbool.h>
#include<stdio.h>
// struct TreeNode {
//       int val;
//       struct TreeNode *left;
//       struct TreeNode *right;
//   };
bool dfs(struct TreeNode*,struct TreeNode*);
bool isSymmetric(struct TreeNode* root) {
   return  dfs(root->left,root->right);
}
bool dfs(struct TreeNode* left,struct TreeNode* rigfht){
    if (left==NULL&&rigfht==NULL)
    {
        return true;
    }
    if (left==NULL||rigfht==NULL)
    {
        return false;
    }
    if (left->val !=rigfht->val)
    {
        return false;
    }
    return dfs(left->left,rigfht->right) && dfs(left->right,rigfht->left);
    
}
// @lc code=end



/*
// @lcpr case=start
// [1,2,2,3,4,4,3]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,2,null,3,null,3]\n
// @lcpr case=end

 */

