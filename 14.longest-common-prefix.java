/*
 * @lc app=leetcode.cn id=14 lang=java
 * @lcpr version=30204
 *
 * [14] 最长公共前缀
 */


// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
class Solution {
    public String longestCommonPrefix(String[] strs) {
        if (strs.length==1) {
            return strs[0];
        }
        StringBuilder builder=new StringBuilder();
        for(int i=0;i<strs[0].length();i++){
            for(int j=1;j<strs.length;j++){
                if (i>strs[j].length()-1||strs[j].charAt(i)!=strs[0].charAt(i)) {
                    return builder.toString();
                }
            }
            builder.append(strs[0].charAt(i));
        }
        return builder.toString();
    }
}
// @lc code=end



/*
// @lcpr case=start
// ["flower","flow","flight"]\n
// @lcpr case=end

// @lcpr case=start
// ["dog","racecar","car"]\n
// @lcpr case=end

 */

