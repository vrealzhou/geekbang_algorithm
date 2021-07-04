/*
给定一个非空且只包含非负数的整数数组 nums，数组的度的定义是指数组里任一元素出现频数的最大值。

你的任务是在 nums 中找到与 nums 拥有相同大小的度的最短连续子数组，返回其长度。





来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/degree-of-an-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

func findShortestSubArray(nums []int) int {
	// count nums and store nums locaion
	metas := make(map[int]*meta)
	var degree *meta
	for i := 0; i < len(nums); i++ { // O(n)
		item, ok := metas[nums[i]]
		if !ok {
			item = &meta{
				start:    i,
				count:    0,
				distance: 0,
			}
			metas[nums[i]] = item
		}
		item.count++
		item.distance = i - item.start + 1
		if degree == nil || degree.count < item.count ||
			(degree.count == item.count && degree.distance > item.distance) {
			degree = item
		}
	}
	return degree.distance
}

type meta struct {
	start    int
	count    int
	distance int
}
