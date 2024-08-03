/*
 * @lc app=leetcode.cn id=198 lang=java
 * @lcpr version=30204
 *
 * [198] 打家劫舍
 */


// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
class Solution {
    //动态规划
    //f=max(dfs(i-1),dfs(i-2)+money(i))
    //选 算i-2 之前的金额和
    //不选 算i-1 之前的金额和
    //超时 优化加入一个数组记录

    public int rob(int[] nums) {
        int n=nums.length;
        int[] moneys=new int[n+2];
        for(int i=2;i<n+2;i++){
            moneys[i]=Math.max(moneys[i-1],moneys[i-2]+nums[i-2]);
        }    
        return moneys[n+1];
    }   
}
// @lc code=end



/*
// @lcpr case=start
// [1,2,3,1]\n
// @lcpr case=end

// @lcpr case=start
// [2,7,9,3,1]\n
// @lcpr case=end

 */

