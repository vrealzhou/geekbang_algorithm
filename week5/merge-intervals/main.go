/*
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。



示例 1：

输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].

示例 2：

输入：intervals = [[1,4],[4,5]]
输出：[[1,5]]
解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。



提示：

    1 <= intervals.length <= 104
    intervals[i].length == 2
    0 <= starti <= endi <= 104

*/
package main

import "math/rand"

func merge(intervals [][]int) [][]int {
	qs(intervals, 0, len(intervals)-1)
	result := [][]int{intervals[0]}
	index := 0
	for _, interval := range intervals[1:] {
		interval1, interval2 := mergeInterval(result[index], interval)
		if interval2 != nil {
			result = append(result, interval2)
			index++
		} else {
			result[index] = interval1
		}
	}
	return result
}

func mergeInterval(interval1, interval2 []int) ([]int, []int) {
	if interval1[1] < interval2[0] {
		return interval1, interval2
	}
	if interval1[1] > interval2[1] {
		return interval1, nil
	}
	interval1[1] = interval2[1]
	return interval1, nil
}

func qs(intervals [][]int, l, r int) {
	if l >= r {
		return
	}
	pivot := partition(intervals, l, r)
	qs(intervals, l, pivot)
	qs(intervals, pivot+1, r)
}

func partition(intervals [][]int, l, r int) int {
	pivot := l + int(rand.Int31n(int32(r-l+1)))
	pivotVal := intervals[pivot]
	for l <= r {
		for compare(intervals[l], pivotVal) < 0 {
			l++
		}
		for compare(intervals[r], pivotVal) > 0 {
			r--
		}
		if l <= r {
			intervals[l], intervals[r] = intervals[r], intervals[l]
			l++
			r--
		}
	}
	return r
}

func compare(interval1, interval2 []int) int {
	front := interval1[0] - interval2[0]
	if front != 0 {
		return front
	}
	return interval1[1] - interval2[1]
}
