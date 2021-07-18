package main

// 例题做法
func relativeSortArray1(arr1 []int, arr2 []int) []int {
	arr := make([]int, 1001) // 从0开始算，包含1000，可取值一共1001个
	for _, v := range arr1 {
		arr[v]++
	}
	result := make([]int, len(arr1))
	index := 0
	for _, v := range arr2 {
		for arr[v] > 0 {
			result[index] = v
			index++
			arr[v]--
		}
	}
	for i := range arr {
		for arr[i] > 0 {
			result[index] = i
			index++
			arr[i]--
		}
	}
	return result
}
