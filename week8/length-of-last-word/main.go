/*
给你一个字符串 s，由若干单词组成，单词前后用一些空格字符隔开。返回字符串中最后一个单词的长度。

单词 是指仅由字母组成、不包含任何空格字符的最大子字符串。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/length-of-last-word
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

func lengthOfLastWord(s string) int {
	reset := false
	length := 0
	for _, c := range s {
		if c == ' ' {
			reset = true
		} else {
			if reset {
				length = 0
				reset = false
			}
			length++
		}
	}
	return length
}
