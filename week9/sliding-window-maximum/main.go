/*
给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。

返回滑动窗口中的最大值。

**尝试用平衡树解决**
Go没有自带平衡树，自己写个ALV树

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sliding-window-maximum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import "github.com/vrealzhou/geekbang_algorithm/internal/data"

func maxSlidingWindow(nums []int, k int) []int {
	t := data.NewAVLTree()
	result := make([]int, 0)
	for i, v := range nums {
		t.Insert(data.IntKey(v))
		if i < k-1 { // 窗口未满
			continue
		}
		result = append(result, int(t.Max().(data.IntKey)))
		// fmt.Println("i", i, "remove", i-k+1)
		t.Remove(data.IntKey(nums[i-k+1]))
	}
	return result
}
