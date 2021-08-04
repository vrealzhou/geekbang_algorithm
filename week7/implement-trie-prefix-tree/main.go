/*
Trie（发音类似 "try"）或者说 前缀树 是一种树形数据结构，用于高效地存储和检索字符串数据集中的键。这一数据结构有相当多的应用情景，例如自动补完和拼写检查。

请你实现 Trie 类：

    Trie() 初始化前缀树对象。
    void insert(String word) 向前缀树中插入字符串 word 。
    boolean search(String word) 如果字符串 word 在前缀树中，返回 true（即，在检索之前已经插入）；否则，返回 false 。
    boolean startsWith(String prefix) 如果之前已经插入的字符串 word 的前缀之一为 prefix ，返回 true ；否则，返回 false 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/implement-trie-prefix-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

type Trie struct {
	root *node
}

type node struct {
	Count    int
	Children map[rune]*node
}

func newNode() *node {
	return &node{
		Count:    0,
		Children: make(map[rune]*node),
	}
}

/** Initialize your data structure here. */
func NewTrie() *Trie {
	return &Trie{
		root: newNode(),
	}
}

/** Inserts a word into the trie. */
func (t *Trie) Insert(word string) {
	t.find([]rune(word), true, true)
}

/** Returns if the word is in the trie. */
func (t *Trie) Search(word string) bool {
	return t.find([]rune(word), true, false)
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (t *Trie) StartsWith(prefix string) bool {
	return t.find([]rune(prefix), false, false)
}

func (t *Trie) find(s []rune, exactMatch, insertIfNotExist bool) bool {
	n := t.root
	for _, c := range s {
		child, ok := n.Children[c]
		if !ok {
			if !insertIfNotExist {
				return false
			}
			child = newNode()
			n.Children[c] = child
		}
		n = child
	}
	if insertIfNotExist {
		n.Count++
	}
	if exactMatch {
		return n.Count > 0
	} else {
		return true
	}
}
