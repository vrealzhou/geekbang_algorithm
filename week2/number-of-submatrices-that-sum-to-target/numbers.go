/*

	给出矩阵 matrix 和目标值 target，返回元素总和等于目标值的非空子矩阵的数量。

	子矩阵 x1, y1, x2, y2 是满足 x1 <= x <= x2 且 y1 <= y <= y2 的所有单元 matrix[x][y] 的集合。

	如果 (x1, y1, x2, y2) 和 (x1', y1', x2', y2') 两个子矩阵中部分坐标不同（如：x1 != x1'），那么这两个子矩阵也不同。

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/number-of-submatrices-that-sum-to-target
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/
package main

import "fmt"

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	// 求前缀和
	s := make([][]int, len(matrix)+1)
	length := len(matrix[0])
	s[0] = make([]int, length+1) // zero line
	for i := 1; i <= length; i++ {
		s[i] = make([]int, length+1)
		for j := 1; j <= length; j++ {
			s[i][j] = s[i-1][j] + s[i][j-1] - s[i-1][j-1] + matrix[i-1][j-1]
		}
	}
	fmt.Println(s)
	// 枚举
	for p := 1; p < len(s); p++ {
		for q := 1; q < len(s[p]); q++ {

		}
	}
	return 0
}

func main() {

}
