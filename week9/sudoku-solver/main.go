/*
编写一个程序，通过填充空格来解决数独问题。

数独的解法需 遵循如下规则：

    数字 1-9 在每一行只能出现一次。
    数字 1-9 在每一列只能出现一次。
    数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图）

数独部分空格内已填入了数字，空白格用 '.' 表示。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sudoku-solver
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import (
	"fmt"
	"strings"
)

func solveSudoku(board [][]byte) {
	v := NewValidator(board)
	dfs(board, v)
}

func dfs(board [][]byte, v *Validator) bool {
	values, i, j := next(board, v)
	if i == -1 {
		return true
	}
	for x := byte('1'); x <= '9'; x++ {
		if values&(2>>x-'0') == 0 { // 掠过
			continue
		}
		board[i][j] = x
		valid := v.MarkVal(i, j, x)
		if valid {
			done := dfs(board, v)
			if done {
				return true
			}
			v.UnmarkVal(i, j, x)
		}
		board[i][j] = '.'
	}
	return false
}

func next(board [][]byte, v *Validator) (int, int, int) {
	minCount, minValues, mi, mj := 10, 0, -1, -1
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				continue
			}
			count := 0
			values := v.PossibleValues(i, j)
			for x := 1; x < 1<<9; x <<= 1 {
				if x&values != 0 {
					count++
				}
			}
			if count < minCount {
				minCount, minValues, mi, mj = count, values, i, j
			}
		}
	}
	return minValues, mi, mj
}

func FlagsStr(flags []int) string {
	buf := strings.Builder{}
	buf.WriteRune('[')
	for i, item := range flags {
		if i > 0 {
			buf.WriteRune(',')
		}
		buf.WriteString(fmt.Sprintf("%b", item))
	}
	buf.WriteRune(']')
	return buf.String()
}

type Validator struct {
	flags [9]int
	mask  int
}

func NewValidator(board [][]byte) *Validator {
	v := &Validator{}
	for i, row := range board {
		for j, val := range row {
			if val == '.' {
				continue
			}
			valid := v.MarkVal(i, j, val)
			if !valid {
				return v
			}
		}
	}
	for i := 0; i < 9; i++ {
		v.mask |= 2 << i
	}
	return v
}

// 标记该位置的值
func (v *Validator) MarkVal(i, j int, val byte) bool {
	valRow, valCol, valBox := v.calcMarkValue(val)
	bi := v.boxIndex(i, j)
	if v.flags[i]&valRow != 0 || v.flags[j]&valCol != 0 || v.flags[bi]&valBox != 0 {
		return false
	}
	v.flags[i] |= valRow
	v.flags[j] |= valCol
	v.flags[bi] |= valBox
	return true
}

// 取消指定位置的值
func (v *Validator) UnmarkVal(i, j int, val byte) {
	valRow, valCol, valBox := v.calcMarkValue(val)
	bi := v.boxIndex(i, j)
	v.flags[i] &^= valRow
	v.flags[j] &^= valCol
	v.flags[bi] &^= valBox
}

// 取得该位置可以填的值
func (v *Validator) PossibleValues(i, j int) int {
	valRow := v.flags[i] & v.mask
	valCol := (v.flags[j] >> 9) & v.mask
	bi := v.boxIndex(i, j)
	valBox := (v.flags[bi] >> 18) & v.mask
	return ^(valRow | valCol | valBox)
}

func (v *Validator) calcMarkValue(val byte) (int, int, int) {
	val = val - '0'
	valRow := 1 << val
	valCol := valRow << 9
	valBox := valCol << 9
	return valRow, valCol, valBox
}

func (v *Validator) boxIndex(i, j int) int {
	return i/3*3 + j/3
}
