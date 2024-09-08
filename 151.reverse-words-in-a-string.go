/*
 * @lc app=leetcode.cn id=151 lang=golang
 * @lcpr version=30204
 *
 * [151] 反转字符串中的单词
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func reverseWords(s string) string {
	stack := make([]string, 0)
	var curr strings.Builder
	for _, v := range s {
		if v == ' ' {
			if curr.Len() == 0 {
				continue
			} else {
				stack = append(stack, curr.String())
				curr.Reset()
			}
		} else {
			curr.WriteString(string(v))
		}
	}
	if curr.Len() != 0 {
		stack = append(stack, curr.String())
	}

	var ans strings.Builder
	for i := len(stack) - 1; i > 0; i-- {
		ans.WriteString(stack[i])
		ans.WriteByte(' ')
	}
	ans.WriteString(stack[0])
	return ans.String()
}

// @lc code=end

/*
// @lcpr case=start
// "the sky is blue"\n
// @lcpr case=end

// @lcpr case=start
// "  hello world  "\n
// @lcpr case=end

// @lcpr case=start
// "a good   example"\n
// @lcpr case=end

*/

