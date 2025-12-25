/*
 * @lc app=leetcode.cn id=3075 lang=java
 * @lcpr version=30204
 *
 * [3075] 幸福值最大化的选择方案
 */


// @lcpr-template-start

// @lcpr-template-end
// @lc code=start

import java.util.Arrays;

class Solution {
    public long maximumHappinessSum(int[] happiness, int k) {
        Arrays.sort(happiness);
        long ans=0;
        int sum=0;
        int n=happiness.length;
        for(int i=n-1;i>=n-k;i--){
            long t=happiness[i]-sum;
            if (t<=0) {
                break;
            }
            sum++;
            ans+=t;
        }
        return ans;
    }
}
// @lc code=end



/*
// @lcpr case=start
// [1,2,3]\n2\n
// @lcpr case=end

// @lcpr case=start
// [1,1,1,1]\n2\n
// @lcpr case=end

// @lcpr case=start
// [2,3,4,5]\n1\n
// @lcpr case=end
// @lcpr case=start
// [2135218,73431904,92495076,77528042,82824634,3036629,28375907,65220365,40948869,58914871,57169530,89783499,19582915,19676695,11932465,21770144,49740276,22303751,80746555,97391584,95775653,43396943,47271136,43935930,59643137,64183008,8892641,39587569,85086654,5663585,82925096,24868817,95900395,48155864,74447380,7618448,63299623,91141186,33347112,81951555,52867615,92184410,7024265,85525916,29846922,59532692,47267934,6514603,1137830,97807470]\n41\n
// @lcpr case=end


 */

