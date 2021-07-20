/*
传送带上的包裹必须在 D 天内从一个港口运送到另一个港口。
传送带上的第 i 个包裹的重量为 weights[i]。每一天，我们都会按给出重量的顺序往传送带上装载包裹。我们装载的重量不会超过船的最大运载重量。
返回能在 D 天内将传送带上的所有包裹送达的船的最低运载能力。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/capacity-to-ship-packages-within-d-days
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

func shipWithinDays(weights []int, days int) int {
	l, r := edge(weights) // 取边界值
	for l < r {
		mid := (l + r) / 2
		if canCarry(weights, days, mid) { // 如果满足说明mid大，向左收敛
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}

func edge(weights []int) (int, int) {
	max := 0
	sum := 0
	for _, weight := range weights {
		if max < weight {
			max = weight
		}
		sum += weight
	}
	return max, sum
}

func canCarry(weights []int, days, capability int) bool { // 利用贪心算法求给定重量能否在指定天内完成
	sum := 0
	dayCount := 1
	for _, weight := range weights {
		if sum+weight > capability {
			dayCount++
			if dayCount > days {
				return false
			}
			sum = weight
		} else {
			sum += weight
		}
	}
	return dayCount <= days
}
