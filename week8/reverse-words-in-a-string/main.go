/*
给你一个字符串 s ，逐个翻转字符串中的所有 单词 。

单词 是由非空格字符组成的字符串。s 中使用至少一个空格将字符串中的 单词 分隔开。

请你返回一个翻转 s 中单词顺序并用单个空格相连的字符串。

说明：

    输入字符串 s 可以在前面、后面或者单词间包含多余的空格。
    翻转后单词间应当仅用一个空格分隔。
    翻转后的字符串中不应包含额外的空格。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-words-in-a-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import "strings"

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	buf := strings.Builder{}
	begin := true
	for i := len(words) - 1; i >= 0; i-- {
		if len(words[i]) == 0 {
			continue
		}
		if begin {
			begin = false
		} else {
			buf.WriteRune(' ')
		}
		buf.WriteString(words[i])
	}
	return buf.String()
}
