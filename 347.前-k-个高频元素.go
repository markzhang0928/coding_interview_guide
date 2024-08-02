/*
 * @lc app=leetcode.cn id=347 lang=golang
 *
 * [347] 前 K 个高频元素
 *
 * https://leetcode.cn/problems/top-k-frequent-elements/description/
 *
 * algorithms
 * Medium (63.82%)
 * Likes:    1867
 * Dislikes: 0
 * Total Accepted:    582.1K
 * Total Submissions: 911.7K
 * Testcase Example:  '[1,1,1,2,2,3]\n2'
 *
 * 给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: nums = [1,1,1,2,2,3], k = 2
 * 输出: [1,2]
 *
 *
 * 示例 2:
 *
 *
 * 输入: nums = [1], k = 1
 * 输出: [1]
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * k 的取值范围是 [1, 数组中不相同的元素的个数]
 * 题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的
 *
 *
 *
 *
 * 进阶：你所设计算法的时间复杂度 必须 优于 O(n log n) ，其中 n 是数组大小。
 *
 */

// @lc code=start
func topKFrequent(nums []int, k int) []int {
	ans := []int{}
	map_num := map[int]int{}
	for _, item := range nums {
		map_num[item]++
	}
	for key, _ := range map_num {
		ans = append(ans, key)
	}
	//核心思想：排序
	//可以不用包函数，自己实现快排
	sort.Slice(ans, func(a, b int) bool {
		return map_num[ans[a]] > map_num[ans[b]]
	})
	return ans[:k]
}

// @lc code=end

