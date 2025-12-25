/*
 * @lc app=leetcode.cn id=3074 lang=java
 * @lcpr version=30204
 *
 * [3074] 重新分装苹果
 */


// @lcpr-template-start

// @lcpr-template-end
// @lc code=start

import java.util.ArrayList;
import java.util.Arrays;

class Solution {
    public int minimumBoxes(int[] apple, int[] capacity) {
        int n=apple.length;
        int appleNum=0;
        for (int i=0;i<n;i++){
            appleNum+=apple[i];
        }
        Arrays.sort(capacity);
        int ans=0;
        int capacityNum=capacity.length;
        for (int i=capacityNum-1;i>=0&&appleNum>0;i--) {
            appleNum-=capacity[i];
            ans++;
        }
        return ans;
    }
}
// @lc code=end



/*
// @lcpr case=start
// [1,3,2]\n[4,3,1,5,2]\n
// @lcpr case=end

// @lcpr case=start
// [5,5,5]\n[2,4,2,7]\n
// @lcpr case=end

// @lcpr case=start
// [19,1,2]\n[10,11,1,1]\n
// @lcpr case=end

 */

