/*
 * @lc app=leetcode.cn id=6 lang=golang
 * @lcpr version=30204
 *
 * [6] Z 字形变换
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func convert(s string, numRows int) string {
	if numRows == 1 || len(s) <= numRows {
		return s
	}
	arr := make([]strings.Builder, numRows)
	for i := 0; i < numRows; i++ {
		arr[i] = strings.Builder{}
	}
	flag, row := -1, 0

	for i := 0; i < len(s); i++ {
		arr[row].WriteByte(s[i])
		if row == 0 || row == numRows-1 {
			flag = -flag
		}
		row += flag
	}
	var ans strings.Builder
	for i := 0; i < numRows; i++ {
		ans.WriteString(arr[i].String())
	}
	return ans.String()
}

// @lc code=end

/*
// @lcpr case=start
// "PAYPALISHIRING"\n3\n
// @lcpr case=end

// @lcpr case=start
// "PAYPALISHIRING"\n4\n
// @lcpr case=end

// @lcpr case=start
// "A"\n1\n
// @lcpr case=end

*/

