/*
 * @lc app=leetcode.cn id=53 lang=java
 * @lcpr version=30204
 *
 * [53] 最大子数组和
 */


// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
class Solution {
    //动态规划
    //f=| nums[i] i=0
    //  | max{f(i-1)+nums[i],nums[i]} i>0
    public int maxSubArray(int[] nums) {
        int n=nums.length;
        int res=nums[0];
        int pre=nums[0];
        for(int i=1;i<n;i++){
            pre=Math.max(pre+nums[i], nums[i]);
            res=Math.max(pre, res);
        }
        return res;
    }
}
// @lc code=end



/*
// @lcpr case=start
// [-2,1,-3,4,-1,2,1,-5,4]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

// @lcpr case=start
// [5,4,-1,7,8]\n
// @lcpr case=end

 */

