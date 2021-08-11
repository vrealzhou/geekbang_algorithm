/*
给定一个字符串 S，返回 “反转后的” 字符串，其中不是字母的字符都保留在原地，而所有字母的位置发生反转。

https://leetcode-cn.com/problems/reverse-only-letters/
*/
package main

func reverseOnlyLetters(s string) string {
	b := []byte(s)
	l, r := 0, len(s)-1
	for l < r {
		for !isLetter(b[l]) && l < r {
			l++
		}
		for !isLetter(b[r]) && l < r {
			r--
		}
		b[l], b[r] = b[r], b[l]
		l++
		r--
	}
	return string(b)
}

func isLetter(v byte) bool {
	return (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z')
}
