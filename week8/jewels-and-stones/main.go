/*
 给定字符串J 代表石头中宝石的类型，和字符串 S代表你拥有的石头。 S 中每个字符代表了一种你拥有的石头的类型，你想知道你拥有的石头中有多少是宝石。

J 中的字母不重复，J 和 S中的所有字符都是字母。字母区分大小写，因此"a"和"A"是不同类型的石头。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/jewels-and-stones
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

func numJewelsInStones(jewels string, stones string) int {
	m := make([]bool, 52)
	for _, v := range jewels {
		m[id(v)] = true
	}
	count := 0
	for _, v := range stones {
		if m[id(v)] {
			count++
		}
	}
	return count
}

func id(v rune) int {
	if v >= 'a' {
		return int(v-'a') + 26
	}
	return int(v - 'A')
}
