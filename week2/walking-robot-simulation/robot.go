package main

func robotSim(commands []int, obstacles [][]int) int {
	ds := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	di := 0
	x := 0
	y := 0
	result := 0
	om := make(map[int]struct{})
	for _, obstacle := range obstacles {
		om[(obstacle[0]+30000)<<16+(obstacle[1]+30000)] = struct{}{}
	}
	for i := 0; i < len(commands); i++ {
		if commands[i] == -1 {
			di = (di + 1) % 4
		} else if commands[i] == -2 {
			di = (4 + di - 1) % 4
		} else {
			for j := 0; j < commands[i]; j++ {
				x1 := x + ds[di][0]
				y1 := y + ds[di][1]
				if _, ok := om[(x1+30000)<<16+(y1+30000)]; ok {
					break
				}
				x = x1
				y = y1
			}
			tmp := x*x + y*y
			if tmp > result {
				result = tmp
			}
		}
	}
	return result
}
