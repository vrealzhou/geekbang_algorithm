package data

type Trie struct {
	root *Node
}

type Node struct {
	Count    int
	Children map[rune]*Node
}

func newNode() *Node {
	return &Node{
		Count:    0,
		Children: make(map[rune]*Node),
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
