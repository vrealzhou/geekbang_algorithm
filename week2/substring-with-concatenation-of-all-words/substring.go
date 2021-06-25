package substring

// 基本思路
// 暴力解法是每次子串都需要重新分词然后生成map
// 如果之前的分词map能继续使用，那么只需要去掉头一个词再添加后一个词就可以了
// 那么可以分开求，按照一个词的长度分组，每组开始分别对应第一个字符到第n个字符（n==len(word)).
// 每一组循环每次递增一个词长的字符。每组里的分词map就可以延续使用到该组结束。每递增一次只需要增加一个新词，减少一个旧词
// 设n=len(s), m = len(words). 总体复杂度是O(n)。wordsMap复杂度是O(m), 扫描的复杂度是O(n log m) (m不是每个循环都有), mapEquals复杂度是O(m)
func findSubstring(s string, words []string) []int {
	wl, m := wordsMap(words)
	indices := make([]int, 0)
	for i := 0; i < wl; i++ { // 按单词字符数分组
		pool := make(map[string]int)
		j := i
		for ; j < (len(words)-1)*wl; j += wl { // 预先填充满 len(words)-1 个分词 map
			word := s[j : j+wl]
			if _, ok := m[word]; ok { // 判断word是否在map中，如果不在不需要添加到pool中
				pool[word]++
			}
		}
		slow := i
		length := len(s) - i
		if length%wl != 0 {
			length = length - length%wl
		}
		for ; j < length; j += wl { // 继续扫描分词，每扫一个判断map是否满足
			word := s[j : j+wl]
			if n, ok := m[word]; ok { // 判断word是否在map中。如果在的话，进行pool添加和比较操作
				v := pool[word]
				v++
				pool[word] = v
				if n == v { // 只有当pool[word]和m[word]相等的时候才有必要做全map比较
					if mapEquals(m, pool) {
						indices = append(indices, slow)
					}
				}
			}
			prev := s[slow : slow+wl]
			if _, ok := pool[prev]; ok { // 判断word是否在pool中，如果在的话，从pool里删除。也可以用m判断，在这里没什么区别。
				remove(pool, prev)
			}
			slow += wl
		}
	}
	return indices
}

func remove(m map[string]int, word string) {
	if v, ok := m[word]; ok {
		if v > 1 {
			m[word]--
		} else {
			delete(m, word)
		}
	}
}

func mapEquals(a, b map[string]int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if v != b[k] {
			return false
		}
	}
	return true
}

func wordsMap(words []string) (int, map[string]int) {
	m := make(map[string]int)
	length := len(words[0])
	for _, word := range words { // O(m)
		m[word]++
	}
	return length, m
}
