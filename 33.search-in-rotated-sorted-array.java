/*
 * @lc app=leetcode.cn id=33 lang=java
 * @lcpr version=30204
 *
 * [33] 搜索旋转排序数组
 */


// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
class Solution {
    public int search(int[] nums, int target) {
       //2分
       int n=nums.length;
       if (n==1){
        return nums[0]==target?  0: -1;
       } 
       int left=0;
       int right=n-1;
       while (left<=right) {
        int mid=(left+right)/2;
        if (nums[mid]==target) {
            return mid;
        }
        //左边升序还是右边升序
        if (nums[0]<=nums[mid]) {
            if (target>=nums[0]&&target<nums[mid]) {
                right=mid-1;
            }else{
                left=mid+1;
            }
        }else{
            if (target>nums[mid]&&target<=nums[n-1]) {
                left=mid+1;
            }else{
                right=mid-1;
            }
        }
       }
       return -1;
    }
}
// @lc code=end



/*
// @lcpr case=start
// [4,5,6,7,0,1,2]\n0\n
// @lcpr case=end

// @lcpr case=start
// [3,1]\n1\n
// @lcpr case=end

// @lcpr case=start
// [1]\n0\n
// @lcpr case=end

 */

