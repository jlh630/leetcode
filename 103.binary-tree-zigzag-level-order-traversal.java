/*
 * @lc app=leetcode.cn id=103 lang=java
 * @lcpr version=30204
 *
 * [103] 二叉树的锯齿形层序遍历
 */


// @lcpr-template-start

// @lcpr-template-end
// @lc code=start

/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
import java.util.*;
class Solution {
    List<List<Integer>> res=new ArrayList<>();
    public List<List<Integer>> zigzagLevelOrder(TreeNode root) {
        dfs(root, 0);
        return res;
    }
    public void dfs(TreeNode node,int level){
        if (node==null){
            return;
        } 
        if (res.size()==level) {
            res.add(new ArrayList<Integer>());
        }
        if (level%2==0) {
            res.get(level).add(node.val);
        }else{
            res.get(level).add(0, node.val);
        }
        dfs(node.left, level+1);
        dfs(node.right, level+1);
    }
}
// @lc code=end



/*
// @lcpr case=start
// [3,9,20,null,null,15,7]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

 */

