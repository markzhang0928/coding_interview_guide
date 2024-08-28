/*
 * @lc app=leetcode.cn id=513 lang=golang
 *
 * [513] 找树左下角的值
 *
 * https://leetcode.cn/problems/find-bottom-left-tree-value/description/
 *
 * algorithms
 * Medium (73.25%)
 * Likes:    592
 * Dislikes: 0
 * Total Accepted:    265.8K
 * Total Submissions: 362.4K
 * Testcase Example:  '[2,1,3]'
 *
 * 给定一个二叉树的 根节点 root，请找出该二叉树的 最底层 最左边 节点的值。
 *
 * 假设二叉树中至少有一个节点。
 *
 *
 *
 * 示例 1:
 *
 *
 *
 *
 * 输入: root = [2,1,3]
 * 输出: 1
 *
 *
 * 示例 2:
 *
 * ⁠
 *
 *
 * 输入: [1,2,3,4,null,5,6,null,null,7]
 * 输出: 7
 *
 *
 *
 *
 * 提示:
 *
 *
 * 二叉树的节点个数的范围是 [1,10^4]
 * -2^31
 *
 *
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findBottomLeftValue(root *TreeNode) int {
	var result int
	// 层序遍历的基础上，记录最后一行第一个节点的数值
	st := list.New()
	st.PushBack(root)
	for st.Len() > 0 {
		length := st.Len()
		for i := 0; i < length; i++ {
			node := st.Remove(st.Front()).(*TreeNode)
			if i == 0 {
				result = node.Val
			}
			if node.Left != nil {
				st.PushBack(node.Left)
			}
			if node.Right != nil {
				st.PushBack(node.Right)
			}
		}
	}
	return result
}

// @lc code=end

