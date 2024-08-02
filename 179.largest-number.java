/*
 * @lc app=leetcode.cn id=179 lang=java
 * @lcpr version=30204
 *
 * [179] 最大数
 */


// @lcpr-template-start

// @lcpr-template-end
// @lc code=start

import java.util.Arrays;
//贪心-排序
class Solution {
    public String largestNumber(int[] nums) {
        int n=nums.length;
        String[] strs=new String[n];
        for(int i=0;i<n;i++){
            strs[i]=String.valueOf( nums[i]);
        }
        Arrays.sort(strs,(x,y)->(y+x).compareTo(x+y));
        StringBuilder res=new StringBuilder();
        if (strs[0].equals("0")) {
            return "0";
        }
        for(String s:strs){
            res.append(s);
        }
        
        return res.toString();
    }
}
// @lc code=end



/*
// @lcpr case=start
// [10,2]\n
// @lcpr case=end

// @lcpr case=start
// [3,30,34,5,9]\n
// @lcpr case=end

 */

