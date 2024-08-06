/*
 * @lc app=leetcode.cn id=128 lang=golang
 * @lcpr version=30204
 *
 * [128] 最长连续序列
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
// 哈希
func longestConsecutive(nums []int) int {
	n := len(nums)
	set := make(map[int]bool, n)
	for _, tmp := range nums {
		set[tmp] = true
	}
	res := 0
	for num := range set {
		//找到连续数的开头
		if !set[num-1] {
			count := 1

			for j := num + 1; set[j]; j, count = j+1, count+1 {
			}
			res = max(count, res)
		}
	}
	return res

}

// @lc code=end

/*
// @lcpr case=start
// [100,4,200,1,3,2]\n
// @lcpr case=end

// @lcpr case=start
// [0,3,7,2,5,8,4,6,0,1]\n
// @lcpr case=end

*/

