/*
 * @lc app=leetcode.cn id=202 lang=golang
 * @lcpr version=30204
 *
 * [202] 快乐数
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
// 哈希表
// func isHappy(n int) bool {
// 	set := make(map[int]bool)

// 	for n != 1 && !set[n] {
// 		sum := 0
// 		set[n] = true
// 		for n > 0 {
// 			sum += (n % 10) * (n % 10)
// 			n /= 10
// 		}
// 		n = sum
// 	}
// 	return n == 1
// }

func getNum(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}
	return sum
}

// 双指针
func isHappy(n int) bool {
	slow := n
	fast := getNum(n)
	for fast != 1 && slow != fast {
		slow = getNum(slow)
		fast = getNum(getNum(fast))
	}
	return fast == 1
}

// @lc code=end

/*
// @lcpr case=start
// 19\n
// @lcpr case=end

// @lcpr case=start
// 2\n
// @lcpr case=end

*/

