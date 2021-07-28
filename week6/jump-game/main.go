/*
给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。

数组中的每个元素代表你在该位置可以跳跃的最大长度。

判断你是否能够到达最后一个下标。



来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/jump-game
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import "math"

// 动归
func canJump(nums []int) bool {
	steps := make([]bool, len(nums))
	steps[0] = true
	for i := 0; i < len(nums)-1; i++ {
		if steps[i] == false {
			continue
		}
		for n := nums[i]; n > 0; n-- {
			if i+n == len(nums)-1 {
				return true
			} else if i+n > len(nums)-1 {
				continue
			}
			steps[i+n] = true
		}
	}
	// fmt.Println(steps)
	return steps[len(nums)-1]
}

/*
   f(i) = max(nums[i]+i, max(nums[i-1]+i-1, max(...)))
*/
func canJump1(nums []int) bool {
	furthest := nums[0]
	for i := 1; i < len(nums); i++ {
		if i <= furthest { // 如果大于代表到furthest这个点过不去了
			furthest = int(math.Max(float64(nums[i]+i), float64(furthest)))
		} else {
			return false
		}
		if furthest >= len(nums)-1 {
			return true
		}
		// fmt.Println(furthest)
	}
	return furthest >= len(nums)-1
}
