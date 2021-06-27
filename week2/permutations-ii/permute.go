/*
	给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
	https://leetcode-cn.com/problems/permutations-ii/
*/

package permutationsii

func permuteUnique(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}
	result := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		v := nums[i]
		rest := append([]int{}, nums[0:i]...)
		subResult := permuteUnique(append(rest, nums[i+1:]...))
		for j := 0; j < len(subResult); j++ {
			result = append(result, append([]int{v}, subResult[j]...))
		}
	}
	return result
}
