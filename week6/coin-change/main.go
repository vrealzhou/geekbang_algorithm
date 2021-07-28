/*
给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。

计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。

你可以认为每种硬币的数量是无限的。



来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/coin-change
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import "math"

/*
动态规划，从底开始计算最优子结构，保存每一步结果。
状态转移方程 s(i) = min(s(i-1), s(i-9), s(i-10))+1
*/
func coinChange(coins []int, amount int) int {
	if len(coins) == 0 {
		return -1
	}
	m := make([]int, amount+1)
	m[0] = 0
	for amt := 1; amt <= amount; amt++ {
		min := int(math.MaxInt32)
		for _, coin := range coins {
			if amt-coin < 0 {
				continue
			}
			count := m[amt-coin] + 1
			if min > count {
				min = count
			}
		}
		m[amt] = min
	}
	if m[amount] == int(math.MaxInt32) {
		return -1
	}
	return m[amount]
}
