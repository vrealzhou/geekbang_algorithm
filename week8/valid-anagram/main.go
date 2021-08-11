/*
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。

    1 <= s.length, t.length <= 5 * 104
    s 和 t 仅包含小写字母

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-anagram
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var ms [26]int
	for i, vs := range s {
		vt := t[i]
		ms[int(vs-'a')]++
		ms[int(vt-'a')]--
	}
	for _, c := range ms {
		if c != 0 {
			return false
		}
	}
	return true
}
