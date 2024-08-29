/*
 * @lc app=leetcode.cn id=113 lang=golang
 *
 * [113] 路径总和 II
 *
 * https://leetcode.cn/problems/path-sum-ii/description/
 *
 * algorithms
 * Medium (63.36%)
 * Likes:    1134
 * Dislikes: 0
 * Total Accepted:    426.1K
 * Total Submissions: 671.9K
 * Testcase Example:  '[5,4,8,11,null,13,4,7,2,null,null,5,1]\n22'
 *
 * 给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。
 *
 * 叶子节点 是指没有子节点的节点。
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
 * 输出：[[5,4,11,2],[5,8,4,5]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [1,2,3], targetSum = 5
 * 输出：[]
 *
 *
 * 示例 3：
 *
 *
 * 输入：root = [1,2], targetSum = 0
 * 输出：[]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点总数在范围 [0, 5000] 内
 * -1000
 * -1000
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
func pathSum(root *TreeNode, targetSum int) [][]int {
	result := make([][]int, 0)
	traverse(root, &result, new([]int), targetSum)
	return result
}

func traverse(node *TreeNode, result *[][]int, currPath *[]int, targetSum int) {
	if node == nil { // 这个判空也可以挪到递归遍历左右子树时去判断
		return
	}

	targetSum -= node.Val                   // 将targetSum在遍历每层的时候都减去本层节点的值
	*currPath = append(*currPath, node.Val) // 把当前节点放到路径记录里

	if node.Left == nil && node.Right == nil && targetSum == 0 { // 如果剩余的targetSum为0, 则正好就是符合的结果
		// 不能直接将currPath放到result里面, 因为currPath是共享的, 每次遍历子树时都会被修改
		pathCopy := make([]int, len(*currPath))
		for i, element := range *currPath {
			pathCopy[i] = element
		}
		*result = append(*result, pathCopy) // 将副本放到结果集里
	}

	traverse(node.Left, result, currPath, targetSum)
	traverse(node.Right, result, currPath, targetSum)
	*currPath = (*currPath)[:len(*currPath)-1] // 当前节点遍历完成, 从路径记录里删除掉
}

// @lc code=end

