package main

type DFSearcher struct {
	matrix  [][]int
	visited [][]int
}

func NewDFSearcher(matrix [][]int) *DFSearcher {
	visited := make([][]int, len(matrix))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]int, len(matrix[i]))
	}
	return &DFSearcher{
		matrix:  matrix,
		visited: visited,
	}
}

func (s *DFSearcher) Search() int {
	count := 0
	for y, row := range s.matrix {
		for x := range row {
			if s.visited[y][x] > 0 {
				continue
			}
			c := s.dfs(x, y)
			if c > count {
				count = c
			}
		}
	}
	return count
}

func (s *DFSearcher) dfs(x, y int) int {
	count := 1
	checkSteps(x, y, s.matrix, func(xi, yi int) {
		d := s.visited[yi][xi]
		if d > 0 {
			if d+1 > count {
				count = d + 1
			}
			return
		}
		d = s.dfs(xi, yi)
		if d+1 > count {
			count = d + 1
		}
	})
	s.visited[y][x] = count
	return count
}
