/*
 * @lc app=leetcode.cn id=13 lang=java
 * @lcpr version=30204
 *
 * [13] 罗马数字转整数
 */


// @lcpr-template-start

// @lcpr-template-end
// @lc code=start

import java.util.*;

class Solution {
    //哈希
    public int romanToInt(String s) {
        Map<Character,Integer> map=new HashMap<Character,Integer>(){{
            put('I',1);
            put('V',5);
            put('X',10);
            put('L',50);
            put('C',100);
            put('D',500);
            put('M',1000);
        }};
        int index=s.length()-1;
        int sum=0;
        while (index>=0) {
            Character ch=s.charAt(index);
            int addNum=map.get(ch);
            if (index+1<s.length()&&addNum<map.get(s.charAt(index+1))) {
                addNum=-1*addNum;
            }
            sum+=addNum;
            index--;
        }
        return sum;
    }
}
// @lc code=end



/*
// @lcpr case=start
// "III"\n
// @lcpr case=end

// @lcpr case=start
// "IV"\n
// @lcpr case=end

// @lcpr case=start
// "IX"\n
// @lcpr case=end

// @lcpr case=start
// "LVIII"\n
// @lcpr case=end

// @lcpr case=start
// "MCMXCIV"\n
// @lcpr case=end

 */

