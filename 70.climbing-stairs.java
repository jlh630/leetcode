/*
 * @lc app=leetcode.cn id=70 lang=java
 * @lcpr version=30204
 *
 * [70] 爬楼梯
 */


// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
class Solution {
    //动态规划
    // public int climbStairs(int n) {
    //     return dfs(n);
    // }
    // public int  dfs(int n){
    //     if (n<0) {
    //         return 0;
    //     }else if(n==0){
    //         return 1;
    //     }else{
    //         return dfs(n-1)+dfs(n-2);
    //     }
    // }
    //使用数组优化成递推
    public int climbStairs(int n){
        int[] dp=new int[n+2];
        dp[0]=0;
        dp[1]=1;
        for(int i=0;i<n;i++){
            dp[i+2]=dp[i]+dp[i+1];
        }
        return dp[n+1];
        
    }
}
// @lc code=end



/*
// @lcpr case=start
// 2\n
// @lcpr case=end

// @lcpr case=start
// 3\n
// @lcpr case=end

 */

