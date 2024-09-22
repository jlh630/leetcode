/*
 * @lc app=leetcode.cn id=38 lang=golang
 * @lcpr version=30204
 *
 * [38] 外观数列
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func countAndSay(n int) string {
	ans := "1"
	for i := 1; i < n; i++ {
		l := len(ans)
		curr := ans[0]
		size := 1
		var res strings.Builder
		for j := 1; j < l; j++ {
			if ans[j] == curr {
				size++
			} else {
				res.WriteString(strconv.Itoa(size))
				res.WriteByte(curr)
				curr = ans[j]
				size = 1
			}
		}
		res.WriteString(strconv.Itoa(size))
		res.WriteByte(curr)
		ans = res.String()
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// 4\n
// @lcpr case=end

// @lcpr case=start
// 1\n
// @lcpr case=end

*/

