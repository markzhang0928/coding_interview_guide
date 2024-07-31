/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 *
 * https://leetcode.cn/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (44.06%)
 * Likes:    4497
 * Dislikes: 0
 * Total Accepted:    1.9M
 * Total Submissions: 4.3M
 * Testcase Example:  '"()"'
 *
 * 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
 *
 * 有效字符串需满足：
 *
 *
 * 左括号必须用相同类型的右括号闭合。
 * 左括号必须以正确的顺序闭合。
 * 每个右括号都有一个对应的相同类型的左括号。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "()"
 * 输出：true
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "()[]{}"
 * 输出：true
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = "(]"
 * 输出：false
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 10^4
 * s 仅由括号 '()[]{}' 组成
 *
 *
 */

// @lc code=start
func isValid(s string) bool {
	length := len(s)
	if length == 0 {
		return true
	} else if length%2 != 0 {
		return false
	}

	stack := make([]rune, 0)
	m := make(map[rune]rune)
	m[')'] = '('
	m['}'] = '{'
	m[']'] = '['

	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
		} else {
			// 如果是右括号，先判断栈内是否还有元素
			if len(stack) == 0 {
				return false
			}
			// 再判断栈顶元素是否能够匹配
			peek := stack[len(stack)-1]
			if peek != m[c] {
				return false
			}
			// 模拟栈顶弹出
			stack = stack[:len(stack)-1]
		}
	}

	// 若栈中不再包含元素，则能完全匹配
	return len(stack) == 0
}

// @lc code=end

