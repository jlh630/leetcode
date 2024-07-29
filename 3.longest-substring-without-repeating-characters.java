/*
 * @lc app=leetcode.cn id=3 lang=java
 * @lcpr version=30204
 *
 * [3] 无重复字符的最长子串
 */


// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
//滑动窗口

import java.util.*;

class Solution {
    public int lengthOfLongestSubstring(String s) {
        int n=s.length();
        int max=0;
        int left=0;
        Map<Character,Integer> map=new HashMap<>(); //k:字符 v:下标
        for(int i=0;i<n;i++){
            if (map.containsKey(s.charAt(i))) {
                //1.有重复的窗口往右left+1
                //2.如果是窗口之前的key left不变
                left=Math.max(left, map.get(s.charAt(i))+1);
            }
            //添加or更新下标
            map.put(s.charAt(i),i);
            max=Math.max(max, i-left+1);
        }
    return max;
    }
}
// @lc code=end



/*
// @lcpr case=start
// "abcabcbb"\n
// @lcpr case=end

// @lcpr case=start
// "bbbbb"\n
// @lcpr case=end

// @lcpr case=start
// "pwwkew"\n
// @lcpr case=end

 */

