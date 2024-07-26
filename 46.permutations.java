/*
 * @lc app=leetcode.cn id=46 lang=java
 * @lcpr version=30204
 *
 * [46] 全排列
 */


// @lcpr-template-start

// @lcpr-template-end
// @lc code=start

import java.util.ArrayList;
import java.util.List;

class Solution {
    List<List<Integer>> res=new ArrayList<>();
    public List<List<Integer>> permute(int[] nums) {
        dfs(new ArrayList<>(), nums, new boolean[nums.length]);
        return res;

    }
    public void dfs(List<Integer> currList,int[] nums,boolean [] road){
        if (currList.size()==nums.length) {
            res.add(new ArrayList<>(currList));
        }
        for(int i=0;i<nums.length;i++){
            if (!road[i]) {
                currList.add(nums[i]);
                road[i]=true;
                dfs(currList, nums, road);
                road[i]=false;
                currList.remove(currList.size()-1);
            }
        }
    } 
}
// @lc code=end



/*
// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [0,1]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

 */

