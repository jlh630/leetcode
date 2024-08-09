/*
 * @lc app=leetcode.cn id=100 lang=c
 * @lcpr version=30204
 *
 * [100] 相同的树
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
// #define bool _Bool
// #define true 1
// #define false 0
// #define NULL ((void*)0)
 struct TreeNode {
      int val;
      struct TreeNode *left;
      struct TreeNode *right;
  };
bool dfs(struct TreeNode*,struct TreeNode*);
bool isSameTree(struct TreeNode* p, struct TreeNode* q) {
    return dfs(p,q);
}

bool dfs(struct TreeNode* p,struct TreeNode* q){
     if (p==NULL&&q==NULL){
        return true;
     } 
     if(p==NULL||q ==NULL){
        return false;
     }
    if(p->val!=q->val){
        return false;
    }
     return dfs(p->left,q->left) && dfs(p->right,q->right);

}
// @lc code=end



/*
// @lcpr case=start
// [1,2,3]\n[1,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n[1,null,2]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,1]\n[1,1,2]\n
// @lcpr case=end

 */

