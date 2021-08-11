/*
你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。

给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，今晚能够偷窃到的最高金额。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/house-robber-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return max(sub(nums, 1, len(nums)), sub(nums, 0, len(nums)-1))
}

func sub(nums []int, l, r int) int {
	if r-l == 1 {
		return nums[l]
	}
	dp1, dp2 := nums[l], max(nums[l], nums[l+1])
	for i := l + 2; i < r; i++ {
		v := max(dp1+nums[i], dp2)
		dp1 = dp2
		dp2 = v
	}
	return dp2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
