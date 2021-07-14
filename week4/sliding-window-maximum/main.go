/*
给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。

返回滑动窗口中的最大值。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sliding-window-maximum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import "github.com/vrealzhou/geekbang_algorithm/internal/data"

func maxSlidingWindow(nums []int, k int) []int {
	h := data.NewHeap(0, func(source, target interface{}) int {
		return source.(*Item).Val - target.(*Item).Val
	})
	result := make([]int, 0)
	for i, v := range nums {
		h.Add(&Item{v, i})
		if i < k-1 { // 窗口未满
			continue
		}
		for {
			max := h.Pop()
			if max == nil {
				break
			}
			if max.(*Item).Index+k >= i {
				result = append(result, max.(*Item).Val)
				break
			}
		}
	}
	return result
}

type Item struct {
	Val   int
	Index int
}
