/*
	给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
	https://leetcode-cn.com/problems/permutations-ii/
*/

package permutationsii

func permuteUnique(nums []int) [][]int {
	f := &finder{
		result: make(map[uint][]int),
		used:   make([]bool, len(nums)),
		set:    make([]int, 0),
		setVal: 0,
	}
	f.permute(nums, 0)
	result := make([][]int, len(f.result))
	i := 0
	for _, v := range f.result {
		result[i] = v
		i++
	}
	return result
}

type finder struct {
	result map[uint][]int
	used   []bool
	set    []int
	setVal uint
}

func (f *finder) permute(nums []int, index int) {
	if index == len(nums) {
		f.result[f.setVal] = copyIntAry(f.set)
		return
	}
	for i := 0; i < len(nums); i++ {
		if !f.used[i] {
			f.set = append(f.set, nums[i])
			f.setVal = f.setVal*100 + uint(nums[i]+10)
			f.used[i] = true
			f.permute(nums, index+1)
			f.used[i] = false
			f.set = f.set[:len(f.set)-1]
			f.setVal = f.setVal / 100
		}
	}
}

func copyIntAry(a []int) []int {
	r := make([]int, len(a))
	copy(r, a)
	return r
}
func copyBoolAry(a []bool) []bool {
	r := make([]bool, len(a))
	copy(r, a)
	return r
}
