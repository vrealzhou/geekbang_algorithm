/*
给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。

异位词 指字母相同，但排列不同的字符串。

    1 <= s.length, p.length <= 3 * 104
    s 和 p 仅包含小写字母

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-all-anagrams-in-a-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

func findAnagrams(s string, p string) []int {
	lp := len(p)
	var hp [26]int
	for i := 0; i < lp; i++ {
		hp[int(p[i]-'a')]++
	}
	var h [26]int
	ans := make([]int, 0)
	for i := 0; i < len(s); i++ {
		h[int(s[i]-'a')]++
		if i < lp-1 {
			continue
		} else if i >= lp {
			h[int(s[i-lp]-'a')]--
		}
		equal := true
		for i, c := range hp {
			if h[i] != c {
				equal = false
				break
			}
		}
		if equal {
			ans = append(ans, i-lp+1)
		}
	}
	return ans
}
