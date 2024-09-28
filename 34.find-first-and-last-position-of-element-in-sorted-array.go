/*
 * @lc app=leetcode.cn id=34 lang=golang
 * @lcpr version=30204
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
package main

func searchRange(nums []int, target int) []int {
	left := 0
	right := len(nums) - 1
	resL := -1
	resM := -1
	resR := -1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			if resM == -1 {
				resM = mid
			}
			resL = mid
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if resL == -1 {
		return []int{-1, -1}
	}
	left = resM
	right = len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			resR = mid
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return []int{resL, resR}
}

// @lc code=end

/*
// @lcpr case=start
// [5,7,7,8,8,10]\n8\n
// @lcpr case=end

// @lcpr case=start
// [5,7,7,8,8,10]\n6\n
// @lcpr case=end

// @lcpr case=start
// []\n0\n
// @lcpr case=end

*/
