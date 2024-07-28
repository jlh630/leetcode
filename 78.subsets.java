/*
 * @lc app=leetcode.cn id=78 lang=java
 * @lcpr version=30204
 *
 * [78] 子集
 */


// @lcpr-template-start

// @lcpr-template-end
// @lc code=start

import java.util.*;

class Solution {
    
    public List<List<Integer>> subsets(int[] nums) {
     int n=nums.length;
     List<List<Integer>> res=new ArrayList<>();
     for(int i=1;i< (1<<n);i++){
        //0 0 1开始
        int bit=1;
        List<Integer> addList=new ArrayList<>();
        for(int j=0;j<n;j++){
            if ((i & bit) != 0 ) {
                addList.add(nums[j]);
            }
            bit =bit << 1;
        }
        res.add(addList);
     }
     res.add(new ArrayList<>());
     return res;
    }
}
// @lc code=end



/*
// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [0]\n
// @lcpr case=end

 */

