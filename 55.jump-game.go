/*
 * @lc app=leetcode.cn id=55 lang=golang
 * @lcpr version=30204
 *
 * [55] 跳跃游戏
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
// 贪心
package main

func canJump(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return true
	}
	//最远能到达的下标
	maxIndex := 0
	for index, num := range nums {
		if index > maxIndex {
			return false
		}
		maxIndex = max(maxIndex, num+index)
	}
	return true
}

// @lc code=end

/*
// @lcpr case=start
// [2,3,1,1,4]\n
// @lcpr case=end

// @lcpr case=start
// [3,2,1,0,4]\n
// @lcpr case=end

*/
