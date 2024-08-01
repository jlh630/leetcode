/*
 * @lc app=leetcode.cn id=226 lang=c
 * @lcpr version=30204
 *
 * [226] 翻转二叉树
 */


// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     struct TreeNode *left;
 *     struct TreeNode *right;
 * };
 */
void dfs(struct TreeNode* root,struct TreeNode * left,struct TreeNode* right){
    
    root->left=right;
    root->right=left;

    if (left!=NULL)
    {
        dfs(left,left->left,left->right);
    }
    if (right!=NULL)
    {
        dfs(right,right->left,right->right);
    }
   
}

struct TreeNode* invertTree(struct TreeNode* root) {
    if (root!=NULL)
    {
        dfs(root,root->left,root->right);
    }
    
    return  root; 
}
// @lc code=end



/*
// @lcpr case=start
// [4,2,7,1,3,6,9]\n
// @lcpr case=end

// @lcpr case=start
// [2,1,3]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

 */

