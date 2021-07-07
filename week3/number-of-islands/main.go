/*
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。

岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。

此外，你可以假设该网格的四条边均被水包围。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/number-of-islands
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

func numIslands(grid [][]byte) int {
	islands := 0
	for y, row := range grid {
		for x, v := range row {
			if v == '0' {
				continue
			}
			islands++
			dfs(grid, x, y)
		}
	}
	return islands
}

func dfs(grid [][]byte, x, y int) {
	grid[y][x] = '0'
	for i := 0; i < 4; i++ {
		xi := x + ds[i][0]
		yi := y + ds[i][1]
		if xi < 0 || yi < 0 || xi >= len(grid[0]) || yi >= len(grid) {
			continue
		}
		if grid[yi][xi] == '1' {
			dfs(grid, xi, yi)
		}
	}
	return
}

var ds = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
