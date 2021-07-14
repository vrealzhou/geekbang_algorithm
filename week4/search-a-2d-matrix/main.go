/*
编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：

    每行中的整数从左到右按升序排列。
    每行的第一个整数大于前一行的最后一个整数。

m == matrix.length
n == matrix[i].length
1 <= m, n <= 100
-104 <= matrix[i][j], target <= 104

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/search-a-2d-matrix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

func searchMatrix(matrix [][]int, target int) bool {
	left := 0
	right := len(matrix) * len(matrix[0])
	for left < right {
		mid := (left + right) / 2
		if matrixValue(matrix, mid) >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if right >= len(matrix)*len(matrix[0]) || matrixValue(matrix, right) != target {
		return false
	}
	return true
}

func matrixValue(matrix [][]int, index int) int {
	row := (index + 1) / len(matrix[0])
	col := (index+1)%len(matrix[0]) - 1
	if col < 0 {
		row--
		col = len(matrix[0]) - 1
	}
	return matrix[row][col]
}
