/*
给定一个 m x n 整数矩阵 matrix ，找出其中 最长递增路径 的长度。

对于每个单元格，你可以往上，下，左，右四个方向移动。 你 不能 在 对角线 方向上移动或移动到 边界外（即不允许环绕）。

 提示：

    m == matrix.length
    n == matrix[i].length
    1 <= m, n <= 200
    0 <= matrix[i][j] <= 231 - 1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-increasing-path-in-a-matrix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import (
	"fmt"

	"github.com/vrealzhou/geekbang_algorithm/internal/test"
)

func longestIncreasingPath(matrix [][]int) int {
	s := NewDFSearcher(matrix)
	return s.Search()
}

var steps = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func checkSteps(x, y int, matrix [][]int, f func(x, y int)) {
	for _, step := range steps {
		xi := x + step[0]
		yi := y + step[1]
		if xi < 0 || yi < 0 || xi >= len(matrix[0]) || yi >= len(matrix) {
			continue
		}
		if matrix[yi][xi] <= matrix[y][x] {
			continue
		}
		f(xi, yi)
	}
}

func main() {
	data := `[[0,1,2,3,4,5,6,7,8,9],[19,18,17,16,15,14,13,12,11,10],[20,21,22,23,24,25,26,27,28,29],[39,38,37,36,35,34,33,32,31,30],[40,41,42,43,44,45,46,47,48,49],[59,58,57,56,55,54,53,52,51,50],[60,61,62,63,64,65,66,67,68,69],[79,78,77,76,75,74,73,72,71,70],[80,81,82,83,84,85,86,87,88,89],[99,98,97,96,95,94,93,92,91,90],[100,101,102,103,104,105,106,107,108,109],[119,118,117,116,115,114,113,112,111,110],[120,121,122,123,124,125,126,127,128,129],[139,138,137,136,135,134,133,132,131,130],[0,0,0,0,0,0,0,0,0,0]]`
	// data := `[[3,4,5],[3,2,6],[2,2,1]]`
	ary := test.ParseStringArray(data)
	input := make([][]int, len(ary))
	for i := 0; i < len(input); i++ {
		row := test.ParseIntArray(ary[i])
		input[i] = row
	}
	fmt.Println(longestIncreasingPath(input))
}
