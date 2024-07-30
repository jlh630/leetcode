/*
 * @lc app=leetcode.cn id=20 lang=java
 * @lcpr version=30204
 *
 * [20] 有效的括号
 */


// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
import java.util.*;
class Solution {
    public boolean isValid(String s) {
        Stack<Character> stack=new Stack<>();
        int n=s.length();
        if (n%2!=0) {
            return false;
        }
        for(int i=0;i<n;i++){
            Character ch=s.charAt(i);
            if (ch=='('||ch=='{'||ch=='[') {
                stack.push(ch);
            }else{
                if (stack.isEmpty()) {
                    return false;
                }    
                Character left=stack.pop();
                if ((left=='('&&ch!=')')||(left=='{'&&ch!='}')||(left=='['&&ch!=']')) {
                    return false;                   
                }
            }      
        }

        return stack.isEmpty();
    }
}
// @lc code=end



/*
// @lcpr case=start
// "()"\n
// @lcpr case=end

// @lcpr case=start
// "()[]{}"\n
// @lcpr case=end

// @lcpr case=start
// "(]"\n
// @lcpr case=end

 */

