/*
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。



来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

func letterCombinations(digits string) []string {
	result := make([]string, 0)
	combine(&result, digits)
	return result
}

func combine(r *[]string, digits string) {
	if len(digits) == 0 {
		return
	}
	head := digits[0]
	rest := digits[1:]
	if len(rest) == 0 { // 提前判断是否余下的digits已经为空，可以减少一次递归调用
		for _, c := range keyNums[head] {
			*r = append(*r, c)
		}
		return
	}
	combine(r, rest)
	for i, s := range *r {
		for j, c := range keyNums[head] {
			if j == 0 {
				(*r)[i] = c + s
			} else {
				*r = append(*r, c+s)
			}
		}
	}
}

var keyNums = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}
