/*
	给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
*/
package combinations

func combine(n int, k int) [][]int {
	f := &finder{
		result: make([][]int, 0),
	}
	f.findSubsets(make([]int, 0), 1, n, k)
	return f.result
}

type finder struct {
	result [][]int
}

func (f *finder) findSubsets(set []int, index, n, k int) {
	if len(set) > k || len(set)+n-index+1 < k {
		return
	}
	if index == n+1 {
		s := make([]int, len(set))
		copy(s, set)
		f.result = append(f.result, s)
		return
	}
	// not add
	f.findSubsets(set, index+1, n, k)
	// add
	set = append(set, index)
	f.findSubsets(set, index+1, n, k)
}
