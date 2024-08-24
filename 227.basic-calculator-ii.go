/*
 * @lc app=leetcode.cn id=227 lang=golang
 * @lcpr version=30204
 *
 * [227] 基本计算器 II
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start\
func calculate(s string) int {
	stackNum := make([]int, 0)
	preNum := 0
	preSign := '+'
	for index, v := range s {
		isDigit := v >= '0' && v <= '9'
		if isDigit {
			preNum = preNum*10 + int(v-'0')
		}
		if !isDigit && v != ' ' || index == len(s)-1 {
			if preSign == '+' {
				stackNum = append(stackNum, preNum)
			}
			if preSign == '-' {
				stackNum = append(stackNum, -preNum)
			}
			if preSign == '*' {
				stackNum[len(stackNum)-1] *= preNum
			}
			if preSign == '/' {
				stackNum[len(stackNum)-1] /= preNum
			}
			preNum = 0
			preSign = v
		}
	}
	ans := 0
	for _, num := range stackNum {
		ans += num
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// "3+2*2"\n
// @lcpr case=end

// @lcpr case=start
// " 3/2 "\n
// @lcpr case=end

// @lcpr case=start
// " 3+5 / 2 "\n
// @lcpr case=end

*/

