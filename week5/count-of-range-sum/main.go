/*
给你一个整数数组 nums 以及两个整数 lower 和 upper 。求数组中，值位于范围 [lower, upper] （包含 lower 和 upper）之内的 区间和的个数 。
区间和 S(i, j) 表示在 nums 中，位置从 i 到 j 的元素之和，包含 i 和 j (i ≤ j)。


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/count-of-range-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

func countRangeSum(nums []int, lower int, upper int) int {
	s := make([]int64, len(nums)+1)
	for i, num := range nums {
		s[i+1] = int64(num) + s[i]
	}
	// fmt.Println(s)
	sl := &solution{}
	sl.mergeSort(s, lower, upper)
	return sl.result
}

type solution struct {
	result int
}

func (sl *solution) mergeSort(s []int64, lower, upper int) {
	if len(s) <= 1 {
		return
	}
	mid := (len(s) - 1) / 2
	sl.mergeSort(s[:mid+1], lower, upper)
	sl.mergeSort(s[mid+1:], lower, upper)
	sl.calculate(s[:mid+1], s[mid+1:], lower, upper)
	sl.merge(s, mid)
}

func (sl *solution) calculate(s1, s2 []int64, lower, upper int) {
	l, r := 0, 0
	for _, v := range s1 { // 左边值越大，右侧值应该相应也变大
		for l < len(s2) && int(s2[l]-v) < lower {
			l++
		}
		for r < len(s2) && int(s2[r]-v) <= upper {
			r++
		}
		sl.result += r - l
	}
}

func (sl *solution) merge(s []int64, mid int) {
	tmp := make([]int64, len(s))
	l := 0
	r := mid + 1
	for i := 0; i < len(s); i++ {
		if r >= len(s) || (l <= mid && s[l] <= s[r]) {
			tmp[i] = s[l]
			l++
		} else {
			tmp[i] = s[r]
			r++
		}
	}
	for i, v := range tmp {
		s[i] = v
	}
}
