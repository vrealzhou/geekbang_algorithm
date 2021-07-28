/*
给定一个未排序的整数数组，找到最长递增子序列的个数。
https://leetcode-cn.com/problems/number-of-longest-increasing-subsequence/
*/
package main

func findNumberOfLIS(nums []int) int {
	lengths := make([]lengthCount, len(nums))
	lengths[0] = lengthCount{1, 1} // length, count
	for i := 1; i < len(lengths); i *= 2 {
		copy(lengths[i:], lengths[:i])
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > nums[i] {
				lc := lengths[i]
				lc.length = lc.length + 1
				lengths[j] = lengths[j].merge(lc)
			}
		}
	}
	// fmt.Println(lengths)
	var max lengthCount
	for _, item := range lengths {
		max = max.merge(item)
	}
	return max.count
}

type lengthCount struct {
	length int
	count  int
}

func (c lengthCount) merge(c1 lengthCount) lengthCount {
	if c.length < c1.length {
		c.length = c1.length
		c.count = c1.count
	} else if c.length == c1.length {
		c.count += c1.count
	}
	return c
}
