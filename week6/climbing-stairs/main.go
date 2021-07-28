/*
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

注意：给定 n 是一个正整数。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/climbing-stairs
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import "fmt"

/*
	i: step
	f(i) = f(i-1)+f(i-2)

	init: steps[1]=1, steps[2]=2
*/
func climbStairs(n int) int {
	steps := make([]int, 3)
	for i := 1; i <= n; i++ {
		if i <= 2 {
			steps[i] = i
		} else {
			steps[i%3] = steps[(i-1)%3] + steps[(i-2)%3]
		}
	}
	return steps[n%3]
}

func main() {
	fmt.Println(climbStairs(4))
}
