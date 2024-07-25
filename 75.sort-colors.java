/*
 * @lc app=leetcode.cn id=75 lang=java
 * @lcpr version=30204
 *
 * [75] 颜色分类
 */


// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
class Solution {
    public void sortColors(int[] nums) {
        int n=nums.length;
        int ptr0=0,ptr1=0;
        for(int i=0;i<n;i++){
            if (nums[i]==0) {
                change(nums, ptr0, i);
                if (ptr0<ptr1) {
                    change(nums, ptr1, i);
                }
                ptr0++;
                ptr1++;
            }else if (nums[i]==1) {
                change(nums,ptr1 , i);
                ptr1++;
            }
        }
    }
    public void change(int[] nums,int i,int j){
        int t=nums[i];
        nums[i]=nums[j];
        nums[j]=t;
    }
}
// @lc code=end



/*
// @lcpr case=start
// [2,0,2,1,1,0]\n
// @lcpr case=end

// @lcpr case=start
// [2,0,1]\n
// @lcpr case=end

 */

