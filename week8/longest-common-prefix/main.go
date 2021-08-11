/*
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

提示：

    1 <= strs.length <= 200
    0 <= strs[i].length <= 200
    strs[i] 仅由小写英文字母组成

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-common-prefix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

func longestCommonPrefix(strs []string) string {
	for i, v := range strs[0] {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || rune(strs[j][i]) != v {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
