/*
 * @lc app=leetcode.cn id=12 lang=java
 * @lcpr version=30204
 *
 * [12] 整数转罗马数字
 */


// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
class Solution {
    public String intToRoman(int num) {
        // int[] values={1000,900,500,400,100,90,50,40,10,9,5,4,1};
        // String[] symbols={"M","CM","D","CD","C","XC","L","XL","X","IX","V","IV","I"};
        // StringBuilder ans=new StringBuilder();
        // for (int i=0;i<values.length;i++){
        //     while (num>=values[i]){
        //             num-=values[i];
        //             ans.append(symbols[i]);
        //     }
        //     if(num==0){
        //         break;
        //     }
        // }
        // return ans.toString();
        String[] thousands={"","M","MM","MMM"};
        String[] hundreds={"","C","CC","CCC","CD","D","DC","DCC","DCCC","CM"};
        String[] tens={"","X","XX","XXX","XL","L","LX","LXX","LXXX","XC"};
        String[] nums={"","I","II","III","IV","V","VI","VII","VIII","IX"};
        return thousands[num/1000]+hundreds[num/100%10]+tens[num/10%10]+nums[num%10];

    }
}
// @lc code=end



/*
// @lcpr case=start
// 3999\n
// @lcpr case=end

// @lcpr case=start
// 58\n
// @lcpr case=end

// @lcpr case=start
// 1994\n
// @lcpr case=end

 */

