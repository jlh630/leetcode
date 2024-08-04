/*
 * @lc app=leetcode.cn id=9 lang=golang
 * @lcpr version=30204
 *
 * [9] 回文数
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	sum := 0
	tmp := x
	for tmp > 0 {
		sum = sum*10 + tmp%10
		tmp /= 10
	}
	if x == sum {
		return true
	}
	return false
}

// @lc code=end

/*
// @lcpr case=start
// 121\n
// @lcpr case=end

// @lcpr case=start
// -121\n
// @lcpr case=end

// @lcpr case=start
// 10\n
// @lcpr case=end

*/
