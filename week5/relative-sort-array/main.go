/*
给你两个数组，arr1 和 arr2，

    arr2 中的元素各不相同
    arr2 中的每个元素都出现在 arr1 中

对 arr1 中的元素进行排序，使 arr1 中项的相对顺序和 arr2 中的相对顺序相同。未在 arr2 中出现过的元素需要按照升序放在 arr1 的末尾。

示例：

输入：arr1 = [2,3,1,3,2,4,6,7,9,2,19], arr2 = [2,1,4,3,9,6]
输出：[2,2,2,1,4,3,3,9,6,7,19]

提示：

    1 <= arr1.length, arr2.length <= 1000
    0 <= arr1[i], arr2[i] <= 1000
    arr2 中的元素 arr2[i] 各不相同
    arr2 中的每个元素 arr2[i] 都出现在 arr1 中

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/relative-sort-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import "math/rand"

// 和例题不同解法。先建立arr2的值->index的map
// 随机快速排序，判断大小使用compare函数比较在v2里的index
func relativeSortArray(arr1 []int, arr2 []int) []int {
	m := make(map[int]int)
	// M=len(arr1), N=len(arr2)
	for i, v := range arr2 { // O(N)
		m[v] = i
	}
	qs(arr1, m, 0, len(arr1)-1) // O(MlogM)
	return arr1
}

func qs(arr1 []int, m map[int]int, l, r int) {
	if l >= r {
		return
	}
	pivot := partition(arr1, m, l, r)
	qs(arr1, m, l, pivot)
	qs(arr1, m, pivot+1, r)
}

func partition(nums []int, m map[int]int, l, r int) int {
	pivot := l + int(rand.Int31n(int32(r-l+1)))
	pivotVal := nums[pivot]
	for l <= r {
		for compare(nums[l], pivotVal, m) < 0 {
			l++
		}
		for compare(nums[r], pivotVal, m) > 0 {
			r--
		}
		if l <= r {
			nums[l], nums[r] = nums[r], nums[l] // 交换左右指针的值
			l++
			r--
		}
	}
	return r
}

func compare(v1, v2 int, m map[int]int) int {
	index1, v1ok := m[v1]
	index2, v2ok := m[v2]
	if !v1ok && !v2ok {
		return v1 - v2
	} else if !v1ok {
		return 1
	} else if !v2ok {
		return -1
	} else {
		return index1 - index2
	}
}
