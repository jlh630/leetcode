/*
 * @lc app=leetcode.cn id=42 lang=golang
 * @lcpr version=30204
 *
 * [42] 接雨水
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func trap(height []int) int {
	n := len(height)
	ans := 0
	left := 0
	for left < n && height[left] <= 0 {
		left++
	}
	right := left + 1
	flage := true
	for left < n && right > left && right < n {
		count := 0
		for flage && right < n && height[right] < height[left] {
			count += height[right]
			right++
		}
		//峰顶
		if right >= n || flage == false {
			if flage == true {
				right = n - 1
			}
			flage = false

			count = 0
			l := right - 1
			for l >= 0 && height[right] > height[l] {
				count += height[l]
				l--
			}
			ans += (right-l-1)*min(height[l], height[right]) - count
			right = l
		} else {
			ans += (right-left-1)*min(height[left], height[right]) - count
			left = right
			right = left + 1
		}
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [0,1,0,2,1,0,1,3,2,1,2,1]\n
// @lcpr case=end

// @lcpr case=start
// [4,2,0,3,2,5]\n
// @lcpr case=end

*/

