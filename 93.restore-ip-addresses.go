/*
 * @lc app=leetcode.cn id=93 lang=golang
 * @lcpr version=30204
 *
 * [93] 复原 IP 地址
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func restoreIpAddresses(s string) []string {
	ans := make([]string, 0)
	n := len(s)
	if n < 4 || n > 12 {
		return ans
	}
	atoi := func(str string) (int, bool) {
		l := len(str)
		if l > 3 {
			return 0, false
		}
		if str[0] == '0' && l > 1 {
			return 0, false
		}
		res := 0
		for i := 0; i < l; i++ {
			if str[i] >= '0' && str[i] <= '9' && res*10+int(str[i]-'0') <= 255 {
				res = res*10 + int(str[i]-'0')
			} else {
				return 0, false
			}
		}
		return res, true
	}
	count := 4
	var dfs func(int, string)
	dfs = func(start int, curr string) {

		if len(curr)-3 == n {
			ans = append(ans, curr)
			return
		}
		if count == 0 {
			return
		}
		for i := 1; i <= 3; i++ {
			if i+start <= n {
				if num, ok := atoi(s[start : i+start]); ok {
					count--
					if curr == "" {
						dfs(i+start, strconv.Itoa(num))
					} else {
						dfs(i+start, curr+"."+strconv.Itoa(num))
					}
					count++
				}
			}
		}

	}
	dfs(0, "")
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// "25525511135"\n
// @lcpr case=end

// @lcpr case=start
// "0000"\n
// @lcpr case=end

// @lcpr case=start
// "101023"\n
// @lcpr case=end

// @lcpr case=start
// "111111111111"\n
// @lcpr case=end
*/

