/*
 * @lc app=leetcode.cn id=17 lang=java
 * @lcpr version=30204
 *
 * [17] 电话号码的字母组合
 */


// @lcpr-template-start

// @lcpr-template-end
// @lc code=start

import java.util.ArrayList;
import java.util.*;
import java.util.HashMap;
import java.util.List;

class Solution {
    // @SuppressWarnings("unchecked")
    // public static final List<String>[] buttons = new ArrayList[]{
    //     new ArrayList<>(),
    //     new ArrayList<>(Arrays.asList("a", "b","c")),
    //     new ArrayList<>(Arrays.asList("d","e","f")),
    //     new ArrayList<>(Arrays.asList("g","h","i")),
    //     new ArrayList<>(Arrays.asList("j", "k","l")),
    //     new ArrayList<>(Arrays.asList("m", "n","o")),
    //     new ArrayList<>(Arrays.asList("p", "q","r","s")),
    //     new ArrayList<>(Arrays.asList("t", "u","v")),
    //     new ArrayList<>(Arrays.asList("w", "x","y","z"))
    // };
    public List<String> letterCombinations(String digits) {
        int len =digits.length();
        List<String>res =new ArrayList<>();
        if (len==0) {
            return res;
        }
        Map<Character,String> button=new HashMap<Character,String>(){{
            put('2', "abc");
            put('3', "def");
            put('4', "ghi");
            put('5', "jkl");
            put('6', "mno");
            put('7', "pqrs");
            put('8', "tuv");
            put('9', "wxyz");
        }}; 
        
        fun(res,button, digits, 0, new StringBuffer());
        return res;
    }
    public void fun(List<String> res,Map<Character,String> button,String digits,int count,StringBuffer word){
        if (word.length()==digits.length()) {
            res.add(word.toString());
            return;
        }

        Character character=digits.charAt(count);
        String buff=button.get(character);
        for(int i=0;i<buff.length();i++){
            word.append(buff.charAt(i));
            fun(res,button, digits, count+1, word);
            word.deleteCharAt(word.length()-1);
        }
    }
}
// @lc code=end



/*
// @lcpr case=start
// "23"\n
// @lcpr case=end

// @lcpr case=start
// ""\n
// @lcpr case=end

// @lcpr case=start
// "2"\n
// @lcpr case=end

 */

