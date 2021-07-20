/*
珂珂喜欢吃香蕉。这里有 N 堆香蕉，第 i 堆中有 piles[i] 根香蕉。警卫已经离开了，将在 H 小时后回来。
珂珂可以决定她吃香蕉的速度 K （单位：根/小时）。每个小时，她将会选择一堆香蕉，从中吃掉 K 根。如果这堆香蕉少于 K 根，她将吃掉这堆的所有香蕉，然后这一小时内不会再吃更多的香蕉。
珂珂喜欢慢慢吃，但仍然想在警卫回来前吃掉所有的香蕉。
返回她可以在 H 小时内吃掉所有香蕉的最小速度 K（K 为整数）。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/koko-eating-bananas
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

func minEatingSpeed(piles []int, h int) int {
	l := 1
	r := sum(piles)
	for l < r {
		mid := (l + r) / 2
		if canFinish(piles, h, mid) { // 能满足条件说明mid偏大，向左收敛
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}

func sum(piles []int) int {
	sum := 0
	for _, count := range piles {
		sum += count
	}
	return sum
}

func canFinish(piles []int, h, k int) bool {
	timeConsume := 0
	for _, count := range piles {
		if k >= count {
			timeConsume++
		} else {
			v := count / k
			if count%k > 0 {
				v++
			}
			timeConsume += v
		}
		if timeConsume > h {
			return false
		}
	}
	return true
}
