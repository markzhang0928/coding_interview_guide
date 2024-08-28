去·/*
 * @lc app=leetcode.cn id=404 lang=golang
 *
 * [404] 左叶子之和
 *
 * https://leetcode.cn/problems/sum-of-left-leaves/description/
 *
 * algorithms
 * Easy (63.05%)
 * Likes:    731
 * Dislikes: 0
 * Total Accepted:    323.2K
 * Total Submissions: 511.7K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定二叉树的根节点 root ，返回所有左叶子之和。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 
 * 
 * 输入: root = [3,9,20,null,null,15,7] 
 * 输出: 24 
 * 解释: 在这个二叉树中，有两个左叶子，分别是 9 和 15，所以返回 24
 * 
 * 
 * 示例 2:
 * 
 * 
 * 输入: root = [1]
 * 输出: 0
 * 
 * 
 * 
 * 
 * 提示:
 * 
 * 
 * 节点数在 [1, 1000] 范围内
 * -1000 <= Node.val <= 1000
 * 
 * 
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
func sumOfLeftLeaves(root *TreeNode) int {

	if root == nil {
		return 0
	}
	st := list.New()
	st.PushBack(root)
	result := 0

	for st.Len() > 0 {
		node := st.Remove(st.Back()).(*TreeNode)
		if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
			result += node.Left.Val
		}

		if node.Right != nil {
			st.PushBack(node.Right)
		}
		if node.Left != nil {
			st.PushBack(node.Left)
		}
	}
	return result
}
// @lc code=end
