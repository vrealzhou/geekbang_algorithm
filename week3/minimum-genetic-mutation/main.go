/*
一条基因序列由一个带有8个字符的字符串表示，其中每个字符都属于 "A", "C", "G", "T"中的任意一个。

假设我们要调查一个基因序列的变化。一次基因变化意味着这个基因序列中的一个字符发生了变化。

例如，基因序列由"AACCGGTT" 变化至 "AACCGGTA" 即发生了一次基因变化。

与此同时，每一次基因变化的结果，都需要是一个合法的基因串，即该结果属于一个基因库。

现在给定3个参数 — start, end, bank，分别代表起始基因序列，目标基因序列及基因库，请找出能够使起始基因序列变化为目标基因序列所需的最少变化次数。如果无法实现目标变化，请返回 -1。

注意：

    起始基因序列默认是合法的，但是它并不一定会出现在基因库中。
    如果一个起始基因序列需要多次变化，那么它每一次变化之后的基因序列都必须是合法的。
    假定起始基因序列与目标基因序列是不一样的。



来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-genetic-mutation
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

// 每次变换一个状态，形成一个状态变化树。求最小变化次数就是求层数
// 适用于用广度优先搜索
func minMutation(start string, end string, bank []string) int {
	queue := []string{start}
	depth := 1
	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			seq := queue[i]
			for j := 0; j < len(seq); j++ {
				for _, item := range "ACGT" {
					if rune(seq[j]) == item {
						continue
					}
					s := []rune(seq)
					s[j] = item
					mutated := string(s)
					index := inBank(mutated, bank)
					if index >= 0 {
						if mutated == end {
							return depth
						}
						bank[index] = ""
						queue = append(queue, mutated)
					}
				}
			}
		}
		depth++
		queue = queue[length:]
	}
	return -1
}

func inBank(v string, bank []string) int {
	for i, s := range bank {
		if s == v {
			return i
		}
	}
	return -1
}
