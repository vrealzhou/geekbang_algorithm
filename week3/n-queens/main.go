/*
n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。

给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。

每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。



来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/n-queens
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

func solveNQueens(n int) [][]string {
	combinations = make([][]int, 0)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i
	}
	backtrack(nums, 0, make(map[int]bool), make(map[int]bool))
	// fmt.Println(combinations)
	result := make([][]string, len(combinations))
	for i, set := range combinations {
		plate := make([]string, n)
		for r, loc := range set {
			row := make([]byte, n)
			for col := 0; col < n; col++ {
				if col == loc {
					row[col] = 'Q'
				} else {
					row[col] = '.'
				}
			}
			plate[r] = string(row)
		}
		result[i] = plate
	}
	return result
}

var combinations [][]int

func backtrack(nums []int, index int, cond1, cond2 map[int]bool) {
	if index == len(nums) {
		combinations = append(combinations, clone(nums))
		return
	}
	for i := index; i < len(nums); i++ {
		if cond1[nums[i]-index] { // 如果斜向已经有皇后了就不用往下走了
			continue
		}
		if cond2[nums[i]+index] {
			continue
		}
		cond1[nums[i]-index] = true
		cond2[nums[i]+index] = true
		swap(nums, index, i)
		backtrack(nums, index+1, cond1, cond2)
		swap(nums, index, i)
		cond1[nums[i]-index] = false
		cond2[nums[i]+index] = false
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
