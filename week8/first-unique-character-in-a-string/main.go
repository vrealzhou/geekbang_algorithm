/*
给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。

示例：

	s = "leetcode"
	返回 0

	s = "loveleetcode"
	返回 2

提示：你可以假定该字符串只包含小写字母。
https://leetcode-cn.com/problems/first-unique-character-in-a-string/
*/
package main

func firstUniqChar(s string) int {
	counts := make([]int, 26)
	for i := 0; i < 26; i++ {
		counts[i] = -1
	}
	for i, v := range s {
		if counts[int(v-'a')] == -1 {
			counts[int(v-'a')] = i
		} else {
			counts[int(v-'a')] = -2
		}
	}
	ans := -1
	for _, count := range counts {
		if count < 0 {
			continue
		}
		if ans < 0 || ans > count {
			ans = count
		}
	}
	return ans
}
