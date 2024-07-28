/*
 * @lc app=leetcode.cn id=58 lang=java
 * @lcpr version=30204
 *
 * [58] 最后一个单词的长度
 */


// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
class Solution {
    public int lengthOfLastWord(String s) {
        s=s.trim();
        if (s.length()==0) {
            return 0;
        }
        if (s.length()==1) {
            return 1;
        }
        int count=0;
        
       for(int i=s.length()-1;i>=0&&s.charAt(i)!=' ';i--){
            count++;
       }
       return count;
    }
}
// @lc code=end



/*
// @lcpr case=start
// "Hello World"\n
// @lcpr case=end

// @lcpr case=start
// "   fly me   to   the moon  "\n
// @lcpr case=end

// @lcpr case=start
// "luffy is still joyboy"\n
// @lcpr case=end

 */

