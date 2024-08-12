/*
 * @lc app=leetcode.cn id=189 lang=golang
 * @lcpr version=30204
 *
 * [189] 轮转数组
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
// 方法一 额外数组
//
//	func rotate(nums []int, k int) {
//		n := len(nums)
//		arr := make([]int, n)
//		for index, _ := range nums {
//			t := index + k%n
//			if t >= n {
//				arr[t-n] = nums[index]
//			} else {
//				arr[t] = nums[index]
//			}
//		}
//			copy(nums, arr)
//	}
//
// 方法三 翻转数组
func rotate(nums []int, k int) {
	reverse := func(l, r int) {
		for l < r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}
	n := len(nums)
	k = k % n
	reverse(0, n-1)
	reverse(0, k-1)
	reverse(k, n-1)

}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5,6,7]\n3\n
// @lcpr case=end

// @lcpr case=start
// [-1,-100,3,99]\n2\n
// @lcpr case=end

*/

