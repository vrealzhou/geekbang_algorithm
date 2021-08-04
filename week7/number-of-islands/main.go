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
	m, n := len(grid), len(grid[0])
	ds := make([]int, m*n)
	for i := 0; i < m*n; i++ {
		ds[i] = i + 1
	}
	var find func(i int) int
	find = func(i int) int {
		if ds[i] != i+1 {
			ds[i] = find(ds[i] - 1)
		}
		return ds[i]
	}
	for y, row := range grid {
		for x, v := range row {
			pos := index(x, y, n)
			if v == '0' {
				ds[pos] = 0
				continue
			}
			for _, step := range steps {
				xi, yi := x+step[0], y+step[1]
				if xi < 0 || yi < 0 || xi >= n || yi >= m {
					continue
				} else if grid[yi][xi] == '1' {
					posi := index(xi, yi, n)
					// fmt.Println("1:", x, y, xi, yi, ds)
					parent := find(pos)
					parenti := find(posi)
					// fmt.Println("2:", x, y, parent, xi, yi, parenti, ds)
					if parent < parenti {
						ds[parenti-1] = parent
					} else {
						ds[parent-1] = parenti
					}
					// fmt.Println("3:", x, y, parent, xi, yi, parenti, ds)
				}
			}
		}
	}
	// fmt.Println(ds)
	islands := make(map[int]struct{})
	for i, v := range ds {
		if v == 0 {
			continue
		}
		if _, ok := islands[find(i)]; !ok {
			islands[v] = struct{}{}
		}
	}
	return len(islands)
}

func index(x, y, col int) int {
	return y*col + x
}

var steps = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
