/*
 * @lc app=leetcode.cn id=209 lang=golang
 * @lcpr version=30204
 *
 * [209] 长度最小的子数组
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start

func minSubArrayLen(target int, nums []int) int {
	ans := math.MaxInt32
	n := len(nums)
	left := 0
	right := 0
	sum := 0
	for right < n {
		sum += nums[right]
		for sum >= target {
			ans = min(ans, right-left+1)
			sum -= nums[left]
			left++
		}
		right++
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// 7\n[2,3,1,2,4,3]\n
// @lcpr case=end

// @lcpr case=start
// 4\n[1,4,4]\n
// @lcpr case=end

// @lcpr case=start
// 11\n[1,1,1,1,1,1,1,1]\n
// @lcpr case=end

*/