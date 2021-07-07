/*
给你一个 m x n 的矩阵 board ，由若干字符 'X' 和 'O' ，找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。



示例 1：

输入：board = [["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]
输出：[["X","X","X","X"],["X","X","X","X"],["X","X","X","X"],["X","O","X","X"]]
解释：被围绕的区间不会存在于边界上，换句话说，任何边界上的 'O' 都不会被填充为 'X'。 任何不在边界上，或不与边界上的 'O' 相连的 'O' 最终都会被填充为 'X'。如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。
提示：

    m == board.length
    n == board[i].length
    1 <= m, n <= 200
    board[i][j] 为 'X' 或 'O'

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/surrounded-regions
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

func solve(board [][]byte) {
	visited := make([][]bool, len(board))
	for y, row := range board { // 初始化visited矩阵
		visited[y] = make([]bool, len(row))
	}
	for y, row := range board { // 从边缘做深度优先搜索
		if y == 0 || y == len(board)-1 {
			for x := range row {
				dfs(board, x, y, visited)
			}
		} else {
			dfs(board, 0, y, visited)
			dfs(board, len(row)-1, y, visited)
		}
	}
	for y := 1; y < len(board)-1; y++ {
		for x := 1; x < len(board[y])-1; x++ {
			if !visited[y][x] { // 如果从边缘深度优先搜索不到就说明和边缘没有连接，可以改成‘X’
				board[y][x] = 'X'
			}
		}
	}
}

func dfs(board [][]byte, x, y int, visited [][]bool) {
	if board[y][x] == 'X' {
		return
	}
	if visited[y][x] {
		return
	}
	visited[y][x] = true
	for i := 0; i < 4; i++ {
		xi := x + ds[i][0]
		yi := y + ds[i][1]
		if xi < 0 || yi < 0 || xi >= len(board[0]) || yi >= len(board) {
			continue
		}
		if board[yi][xi] == 'X' {
			continue
		}
		dfs(board, xi, yi, visited)
	}
}

var ds = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
