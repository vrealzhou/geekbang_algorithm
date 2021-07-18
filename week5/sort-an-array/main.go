/*
给你一个整数数组 nums，请你将该数组升序排列。

提示：

    1 <= nums.length <= 50000
    -50000 <= nums[i] <= 50000

https://leetcode-cn.com/problems/sort-an-array/
*/
package main

import (
	"math/rand"
)

func sortArray(nums []int) []int {
	qs(nums, 0, len(nums)-1)
	return nums
}

// 快速排序
func qs(nums []int, l, r int) {
	if l >= r {
		return
	}
	pivot := partition(nums, l, r)
	qs(nums, l, pivot)
	qs(nums, pivot+1, r)
}

// Hoare Partition算法，交换次数最少。https://www.bilibili.com/video/BV1q64y1S7Ax
// 左右两个指针从两端扫描数组，左指针停在第一个不小于中值的位置，右指针停在第一个不大于中值的位置
// 如果左右指针没有重合或者交叉，那么交换左右指针的值
// 继续操作直到左右指针重合或者交叉
func partition(nums []int, l, r int) int {
	pivot := l + int(rand.Int31n(int32(r-l+1))) // 随机中间值
	pivotVal := nums[pivot]
	// fmt.Println(nums, l, r, pivotVal)
	for l <= r {
		for nums[l] < pivotVal {
			l++
		}
		for nums[r] > pivotVal {
			r--
		}
		if l <= r {
			nums[l], nums[r] = nums[r], nums[l] // 交换左右指针的值
			l++
			r--
		}
	}
	// 中值一起交换。如果qs()里pivot是在右侧里排序则这里return l。
	// 上一层变成：
	// qs(nums, l, pivot-1)
	// qs(nums, pivot, r)
	return r
}
