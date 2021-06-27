/*
	给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
*/
package combinations

func combine(n int, k int) [][]int {
	return sub(1, n, k)
}

func sub(start, end, k int) [][]int {
	result := make([][]int, 0)
	for i := start; i <= end; i++ {
		if k > 1 {
			rest := sub(i+1, end, k-1)
			for j := 0; j < len(rest); j++ {
				result = append(result, append([]int{i}, rest[j]...))
			}
		} else {
			result = append(result, append([]int{i}))
		}
	}
	return result
}
