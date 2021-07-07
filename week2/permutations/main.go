/*
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
https://leetcode-cn.com/problems/permutations/
*/
package main

func permute(nums []int) [][]int {
	result = make([][]int, 0)
	backtrack(nums, 0)
	return result
}

var result [][]int

func backtrack(nums []int, index int) {
	if index == len(nums) {
		result = append(result, clone(nums))
		return
	}
	for i := index; i < len(nums); i++ {
		swap(nums, index, i)
		backtrack(nums, index+1)
		swap(nums, index, i)
	}
}

func clone(nums []int) []int {
	output := make([]int, len(nums))
	copy(output, nums)
	return output
}

func swap(nums []int, a, b int) {
	if a == b {
		return
	}
	v := nums[a]
	nums[a] = nums[b]
	nums[b] = v
}
