/*
给定两个字符串 s 和 t，判断它们是否是同构的。

如果 s 中的字符可以按某种映射关系替换得到 t ，那么这两个字符串是同构的。

每个出现的字符都应当映射到另一个字符，同时不改变字符的顺序。不同字符不能映射到同一个字符上，相同字符只能映射到同一个字符上，字符可以映射到自己本身。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/isomorphic-strings
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import "strings"

func isIsomorphic(s string, t string) bool {
	ms := make(map[rune]int)
	mt := make(map[rune]int)
	aryT := []rune(t)
	for i, vs := range s {
		vt := aryT[i]
		idxs, oks := ms[vs]
		idxt, okt := mt[vt]
		if !oks && !okt {
			ms[vs] = i
			mt[vt] = i
		} else if !(oks && okt) || idxs != idxt {
			return false
		}
	}
	return true
}

func solution2(s string, t string) bool {
	for k, v := range s {
		if strings.IndexRune(s, v) != strings.IndexRune(t, rune(t[k])) {
			return false
		}
	}
	return true
}
