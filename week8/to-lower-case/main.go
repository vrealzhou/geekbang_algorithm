/*
给你一个字符串 s ，将该字符串中的大写字母转换成相同的小写字母，返回新的字符串。
https://leetcode-cn.com/problems/to-lower-case/
*/
package main

func toLowerCase(s string) string {
	buf := []rune(s)
	for i, c := range buf {
		if c >= 'A' && c <= 'Z' {
			buf[i] = c + ('a' - 'A')
		}
	}
	return string(buf)
}
