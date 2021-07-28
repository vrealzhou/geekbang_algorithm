/*
给定一个三角形 triangle ，找出自顶向下的最小路径和。

每一步只能移动到下一行中相邻的结点上。相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。也就是说，如果正位于当前行的下标 i ，那么下一步可以移动到下一行的下标 i 或 i + 1 。



来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/triangle
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import "math"

/*
	l: 步数，层数
	i: 下标

	f(l, i) = min(f(l-1, i), f(l-1, i-1))
*/
func minimumTotal(triangle [][]int) int {
	layers := len(triangle)
	n := len(triangle[layers-1])
	opts := make([]int, n) // 可以覆盖，只需要底边长度的数组
	opts[0] = int(math.MaxInt32)
	for i := 1; i < n; i *= 2 { // 设置初始值
		copy(opts[i:], opts[:i])
	}
	opts[0] = triangle[0][0]
	for l := 1; l < layers; l++ {
		for i := len(triangle[l]) - 1; i >= 0; i-- {
			v := triangle[l][i]
			min := opts[i]
			if i > 0 && opts[i-1] < min {
				min = opts[i-1]
			}
			opts[i] = min + v
		}
	}
	min := int(math.MaxInt32)
	for i := 0; i < n; i++ {
		if opts[i] < min {
			min = opts[i]
		}
	}
	return min
}
