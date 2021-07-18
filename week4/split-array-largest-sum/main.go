/*
给定一个非负整数数组 nums 和一个整数 m ，你需要将这个数组分成 m 个非空的连续子数组。

设计一个算法使得这 m 个子数组各自和的最大值最小。
示例 1：

	```
	输入：nums = [7,2,5,10,8], m = 2
	输出：18
	解释：
	一共有四种方法将 nums 分割为 2 个子数组。 其中最好的方式是将其分为 [7,2,5] 和 [10,8] 。
	因为此时这两个子数组各自的和的最大值为18，在所有情况中最小。
	```

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/split-array-largest-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

/*
二分的变体。当取值范围比较大，但是判断值结果true/false的逻辑是确定的，而且复杂度不高的话。
可以将判断值的函数作为二分判断条件。用二分在取值范围内进行查找
*/
func splitArray(nums []int, m int) int {
	l, r := edge(nums)
	// 套用lower_bound二分模版，找到第一个>=n的值
	for l < r {
		mid := (l + r) / 2
		exist := matchGroups(nums, m, mid)
		// fmt.Println(l, mid, r, exist)
		if exist { // 如果存在分组说明给定的值偏大或者正好，需要调整右边界向左收敛
			r = mid
		} else { // 不存在分组说明给定值偏小，需要调整左边界向右收敛。结果值不在l内所以要mid+1
			l = mid + 1
		}
	}
	return r // 最终收敛在r里
}

// 寻找边界值。左边界是所有元素里最大值，右边界是所有元素的和。
func edge(nums []int) (int, int) {
	max := 0
	sum := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
		sum += nums[i]
	}
	return max, sum // 结果肯定存在所以sum不需要+1
}

// 判断是否能在指定分组数量内每个分组的和不大于n。结果必须是可以判断为二分条件的（true/false）
func matchGroups(nums []int, m, n int) bool {
	sum := 0
	groups := 1
	for i := 0; i < len(nums); i++ {
		if sum+nums[i] > n {
			groups++
			sum = 0
			if groups > m { // 超出分组数限制，直接返回false
				// fmt.Println("n", n, "groups", groups, "m", m)
				return false
			}
		}
		sum += nums[i]
	}
	// fmt.Println("n", n, "groups", groups, "m", m)
	return true
}
