/*
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

如果数组中不存在目标值 target，返回 [-1, -1]。

进阶：

    你可以设计并实现时间复杂度为 O(log n) 的算法解决此问题吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

func searchRange(nums []int, target int) []int {
	left := lowerBound(nums, target)
	if left == -1 {
		return []int{-1, -1}
	}
	right := upperBound(nums, target)
	return []int{left, right}
}

func lowerBound(nums []int, target int) int {
	left := 0
	right := len(nums)
	for left < right {
		mid := (left + right) / 2
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if right >= len(nums) || nums[right] != target {
		return -1
	}
	return right
}

func upperBound(nums []int, target int) int {
	left := -1
	right := len(nums) - 1
	for left < right {
		mid := (left + right + 1) / 2
		if nums[mid] <= target {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return right
}
