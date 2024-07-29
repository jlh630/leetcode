/*
 * @lc app=leetcode.cn id=438 lang=java
 * @lcpr version=30204
 *
 * [438] 找到字符串中所有字母异位词
 */


// @lcpr-template-start

// @lcpr-template-end
// @lc code=start

import java.util.*;
//滑动窗口
class Solution {
    public List<Integer> findAnagrams(String s, String p) {
        List<Integer> res=new ArrayList<>();
        int m=p.length(); //窗口大小
        int n=s.length(); 
        if (n<m) {
            return res;
        }
        int[] sCount = new int[26]; 
        int[] pCount = new int[26]; 
        
        for (int i = 0; i < m; ++i) {
            ++sCount[s.charAt(i) - 'a'];
            ++pCount[p.charAt(i) - 'a'];
        }
        //第初次窗口
        if (Arrays.equals(sCount, pCount)) {
            res.add(0);
        }
        int left=0;
        int right=left+m-1;
        while (true) {
            sCount[s.charAt(left) - 'a']--;
            left++;
            right++;
            if(right>=n){
                break;
            }
            sCount[s.charAt(right) - 'a']++;

            if (Arrays.equals(sCount, pCount)) {
                res.add(left);
            }
        }
        return res;
    }
}
// @lc code=end



/*
// @lcpr case=start
// "cbaebabacd"\n"abc"\n
// @lcpr case=end

// @lcpr case=start
// "abab"\n"ab"\n
// @lcpr case=end

 */

