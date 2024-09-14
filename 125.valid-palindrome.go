/*
 * @lc app=leetcode.cn id=125 lang=golang
 * @lcpr version=30204
 *
 * [125] 验证回文串
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	s = strings.ToLower(s)
	for left < right {
		b1, b2 := (s[left] >= 'a' && s[left] <= 'z') || (s[left] >= '0' && s[left] <= '9'), (s[right] >= '0' && s[right] <= '9') || (s[right] >= 'a' && s[right] <= 'z')
		if b1 && b2 {
			if s[left] != s[right] {
				return false
			}
			left++
			right--
		} else if b1 {
			right--
		} else {
			left++
		}
	}
	return true
}

// @lc code=end

/*
// @lcpr case=start
// "A man, a plan, a canal: Panama"\n
// @lcpr case=end

// @lcpr case=start
// "race a car"\n
// @lcpr case=end

// @lcpr case=start
// " "\n
// @lcpr case=end
// @lcpr case=start
// "0p"\n
// @lcpr case=end

*/

