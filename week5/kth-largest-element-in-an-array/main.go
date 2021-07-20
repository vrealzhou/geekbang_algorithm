/*
给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。

请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

示例 1:
输入: [3,2,1,5,6,4] 和 k = 2
输出: 5

示例 2:
输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
输出: 4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/kth-largest-element-in-an-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import "math/rand"

func findKthLargest(nums []int, k int) int {
	qs(nums, 0, len(nums)-1, k)
	return nums[k-1]
}

func qs(nums []int, l, r, k int) {
	if l >= r {
		return
	}
	pivot := partition(nums, l, r)
	k1, k2 := splitK(pivot, k)

	// 如果中值左侧超出k个则不需要对右侧排序了。如果需要对右侧排序则只需要对最大的k2个进行排序。
	// PS：leetcode提交的话总是在第30个用例失败，但是直接运行第30个用例是成功的。
	if k2 > 0 {
		qs(nums, pivot+1, r, k2)
	} else {
		qs(nums, l, pivot, k1)
	}
}

// 如果中值左侧不足k个那么需要对右侧排序，但是应该截取到k-pivot个
func splitK(pivot, k int) (int, int) {
	if pivot >= k {
		return k, 0
	}
	return pivot, k - pivot
}

func partition(nums []int, l, r int) int {
	pivot := l + int(rand.Int31n(int32(r-l+1)))
	pivotVal := nums[pivot]
	for l <= r {
		for nums[l] > pivotVal {
			l++
		}
		for nums[r] < pivotVal {
			r--
		}
		if l <= r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}
	return r
}
