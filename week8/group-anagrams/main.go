/*
给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。

字母异位词 是由重新排列源单词的字母得到的一个新单词，所有源单词中的字母都恰好只用一次。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/group-anagrams
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

func groupAnagrams(strs []string) [][]string {
	m := make(map[uint]int)
	l := make([][]string, len(strs))
	tail := 0
	for _, s := range strs {
		key := hash(s)
		i, ok := m[key]
		if !ok {
			l[tail] = []string{s}
			m[key] = tail
			tail++
		} else {
			l[i] = append(l[i], s)
		}
	}
	return l[:tail]
}

func hash(s string) uint {
	h := make([]int, 26)
	for _, c := range s {
		h[int(c-'a')]++
	}
	var hash uint = 0
	for i, c := range h {
		if c == 0 {
			continue
		}
		for j := 0; j < c; j++ {
			hash = hash*131 + uint(i+1)
		}
	}
	return hash
}
