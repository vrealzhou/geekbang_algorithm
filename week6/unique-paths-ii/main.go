/*
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。

现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？



来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/unique-paths-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

/*
   只能往下，往右移动，转移方程是：f(i,j) = f(x-1,y)+f(x,y-1)
   用一维数组保存最优子结构可以减少内存分配开销
*/
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	n := len(obstacleGrid[0])
	paths := make([]int, len(obstacleGrid)*n)
	for y, row := range obstacleGrid {
		for x, v := range row {
			if v == 1 {
				continue
			}
			if x == 0 && y == 0 {
				paths[0] = 1
				continue
			}
			sum := 0
			if x > 0 {
				sum += paths[pathId(x-1, y, n)]
			}
			if y > 0 {
				sum += paths[pathId(x, y-1, n)]
			}
			paths[pathId(x, y, n)] = sum
		}
	}
	return paths[len(paths)-1]
}

func pathId(x, y, n int) int {
	return y*n + x
}
