/*
 * @lc app=leetcode.cn id=39 lang=java
 * @lcpr version=30204
 *
 * [39] 组合总和
 */


// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
import java.util.*;
class Solution {
    List<List<Integer>> res=new ArrayList<>();
    public List<List<Integer>> combinationSum(int[] candidates, int target) {
        dfs(new ArrayList<>(), 0, candidates, target,0);
        return res;
    }
    void dfs(List<Integer>curList,int curTarget,int[] candidates,int target,int start){
        if (curTarget==target) {
            res.add(new ArrayList<>(curList));
            return;
        }
        if (curTarget>target) {
            return;
        }
        for(int i=start;i<candidates.length;i++){
            curList.add(candidates[i]);
            dfs(curList, curTarget+candidates[i], candidates, target,i);
            curList.remove(curList.size()-1);
        }
    }
}
// @lc code=end



/*
// @lcpr case=start
// [2,3,6,7]\n7\n
// @lcpr case=end

// @lcpr case=start
// [2,3,5]\n8\n
// @lcpr case=end

// @lcpr case=start
// [2]\n1\n
// @lcpr case=end

 */

