/*
 * @lc app=leetcode id=697 lang=golang
 *
 * [697] Degree of an Array
 */

package leetcode

import "math"

// @lc code=start
func findShortestSubArray(nums []int) int {
	// count nums and store nums locaion
	locations := make(map[int][]int)
	for i := 0; i < len(nums); i++ { // O(n)
		locations[nums[i]] = append(locations[nums[i]], i)
	}
	// get degree
	degree := 0
	length := int(math.MaxInt32)
	for _, location := range locations { // O(n)
		count := len(location)
		if count < degree {
			continue
		}
		l := location[len(location)-1] - location[0] + 1
		if count > degree {
			degree = count
			length = l
		} else if l < length {
			length = l
		}
	}

	return length
} // @lc code=end
