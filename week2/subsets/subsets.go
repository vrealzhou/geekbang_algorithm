/*
给你一个整数数组 nums ，数组中的元素互不相同 。返回该数组所有可能的子集（幂集）。

解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。

提示：
	- 1 <= nums.length <= 10
	- -10 <= nums[i] <= 10
	- nums 中的所有元素互不相同
*/
package subset

func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		result = append(result, inner(nums[i], nums[i+1:])...)
	}
	return append([][]int{{}}, result...)
}

func inner(head int, tail []int) [][]int {
	result := [][]int{{head}}
	for i := 0; i < len(tail); i++ {
		rest := inner(tail[i], tail[i+1:])
		for j := 0; j < len(rest); j++ {
			result = append(result, append([]int{head}, rest[j]...))
		}
	}
	return result
}
