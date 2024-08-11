/*
 * @lc app=leetcode.cn id=739 lang=golang
 * @lcpr version=30204
 *
 * [739] 每日温度
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
package main

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	res := make([]int, n)
	stack := []int{}
	for i := 0; i < n; i++ {
		curr := temperatures[i]
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < curr {
			top := stack[len(stack)-1]
			res[top] = i - top
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return res
}

// func main() {
// 	dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73})
// }

// @lc code=end

/*
// @lcpr case=start
// [73,74,75,71,69,72,76,73]\n
// @lcpr case=end

// @lcpr case=start
// [30,40,50,60]\n
// @lcpr case=end

// @lcpr case=start
// [30,60,90]\n
// @lcpr case=end

*/
