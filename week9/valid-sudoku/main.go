// https://leetcode-cn.com/problems/valid-sudoku/
package main

func isValidSudoku(board [][]byte) bool {
	v := &validator{}
	return v.Init(board)
}

type validator struct {
	flags [9]int
}

func (v *validator) Init(board [][]byte) bool {
	for i, row := range board {
		for j, val := range row {
			if val == '.' {
				continue
			}
			valid := v.AddVal(i, j, val)
			if !valid {
				return false
			}
		}
	}
	return true
}

func (v *validator) AddVal(i, j int, val byte) bool {
	val = val - '0'
	valRow := 2 << val
	valCol := 2 << (val + 9)
	valBox := 2 << (val + 18)
	boxIndex := i/3*3 + j/3
	if v.flags[i]&valRow != 0 || v.flags[j]&valCol != 0 || v.flags[boxIndex]&valBox != 0 {
		return false
	}
	v.flags[i] |= valRow
	v.flags[j] |= valCol
	v.flags[boxIndex] |= valBox
	return true
}

func (v *validator) RemoveVal(i, j int, val byte) {
	val = val - '0'
	valRow := 2 << val
	valCol := 2 << (val + 9)
	valBox := 2 << (val + 18)
	boxIndex := i/3*3 + j/3
	v.flags[i] &^= valRow
	v.flags[j] &^= valCol
	v.flags[boxIndex] &^= valBox
}
