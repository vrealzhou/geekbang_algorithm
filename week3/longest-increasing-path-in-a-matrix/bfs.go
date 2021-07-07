package main

type BFSearcher struct {
	matrix [][]int
	queue  []int
}

func NewBFSearcher(matrix [][]int) *BFSearcher {
	queue := make([]int, 0)
	for y, row := range matrix {
		for x := range row {
			queue = append(queue, combinePoint(x, y))
		}
	}
	return &BFSearcher{
		matrix: matrix,
		queue:  queue,
	}
}

func (s *BFSearcher) Search() int {
	count := 0
	visited := make(map[int]int)
	for len(s.queue) > 0 {
		length := len(s.queue)
		for i := 0; i < length; i++ {
			x, y := splitPoint(s.queue[i])
			checkSteps(x, y, s.matrix, func(xi, yi int) {
				v := combinePoint(xi, yi)
				depth, ok := visited[v]
				if ok && depth >= count {
					return
				}
				visited[v] = count
				s.queue = append(s.queue, v)
			})
		}
		count++
		s.queue = s.queue[length:]
	}
	return count
}

func splitPoint(v int) (int, int) {
	x := v / 1000
	y := v - x*1000
	return x, y
}

func combinePoint(x, y int) int {
	return x*1000 + y
}
