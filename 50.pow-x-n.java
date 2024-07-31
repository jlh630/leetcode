/*
 * @lc app=leetcode.cn id=50 lang=java
 * @lcpr version=30204
 *
 * [50] Pow(x, n)
 */


// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
class Solution {
    public double myPow(double x, long n) {
        if(n==0||x==1){
            return 1;
        }
        if (n==1) {
            return x;
        }
        double res=1.0;
        long num=Math.abs(n);
        while (num > 0) {
            if (num % 2 == 1) {
                res *= x;
            }
            x *= x;
            num /= 2;
        }
        if (n<0) {
            res=1/res;
        }
        return res;
    }
}
// @lc code=end



/*
// @lcpr case=start
// 2.00000\n10\n
// @lcpr case=end

// @lcpr case=start
// 2.10000\n3\n
// @lcpr case=end

// @lcpr case=start
// 2.00000\n-2\n
// @lcpr case=end

 */

