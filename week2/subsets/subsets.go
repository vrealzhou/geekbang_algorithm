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
	f := &finder{
		result: make([][]int, 0),
	}
	f.findSubsets(nums, make([]int, 0), 0)
	return f.result
}

type finder struct {
	result [][]int
}

func (f *finder) findSubsets(nums, set []int, index int) {
	if index == len(nums) {
		s := make([]int, len(set))
		copy(s, set)
		f.result = append(f.result, s)
		return
	}
	// not add
	f.findSubsets(nums, set, index+1)
	// add
	set = append(set, nums[index])
	f.findSubsets(nums, set, index+1)
}
