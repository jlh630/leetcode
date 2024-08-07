/*
 * @lc app=leetcode.cn id=8 lang=golang
 * @lcpr version=30204
 *
 * [8] 字符串转换整数 (atoi)
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func myAtoi(s string) int {
	flag := 1
	zero := false
	sign := false
	sum := 0
	for _, c := range s {
		if c == ' ' && sum == 0 && zero == false && sign == false {
			continue
		}
		if c == '-' && sign == false && sum == 0 && zero == false {
			flag = -1
			sign = true
			continue
		}
		if c == '+' && sum == 0 && sign == false && zero == false {
			sign = true
			continue
		}
		if c == '0' && sum == 0 {
			zero = true
			continue
		}
		if c < '0' || c > '9' {
			break
		}
		sum = sum*10 + int(c-'0')
		if flag*sum >= math.MaxInt32 {
			return math.MaxInt32
		} else if flag*sum <= math.MinInt32 {
			return math.MinInt32
		}
	}

	return sum * flag
}

// @lc code=end

/*
// @lcpr case=start
// "42"\n
// @lcpr case=end

// @lcpr case=start
// " -042"\n
// @lcpr case=end

// @lcpr case=start
// "1337c0d3"\n
// @lcpr case=end

// @lcpr case=start
// "0-1"\n
// @lcpr case=end

// @lcpr case=start
// "words and 987"\n
// @lcpr case=end

*/

