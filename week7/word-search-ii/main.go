/*
给定一个 m x n 二维字符网格 board 和一个单词（字符串）列表 words，找出所有同时在二维网格和字典中出现的单词。

单词必须按照字母顺序，通过 相邻的单元格 内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母在一个单词中不允许被重复使用。



来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/word-search-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

func findWords(board [][]byte, words []string) []string {
	m, n := len(board), len(board[0])
	s := &Solution{
		result:  make([]string, 0),
		root:    newNode(),
		board:   board,
		visited: make([]bool, m*n),
		n:       n,
	}
	for _, word := range words {
		s.Insert(word)
	}
	for i, row := range board {
		for j := range row {
			s.dfs(i, j, s.root)
		}
	}
	return s.result
}

type Solution struct {
	result  []string
	root    *Node
	board   [][]byte
	visited []bool
	n       int
}

func (s *Solution) dfs(i, j int, n *Node) {
	key := index(i, j, s.n)
	if s.visited[key] {
		return
	}
	child := n.Children[s.board[i][j]-byte('a')]
	if child == nil {
		return
	} else if child.Word != "" && child.Matched == 0 { // only matched once
		s.result = append(s.result, child.Word)
		child.Matched++
	}
	s.visited[key] = true
	// next
	if i > 0 {
		s.dfs(i-1, j, child)
	}
	if j > 0 {
		s.dfs(i, j-1, child)
	}
	if i < len(s.board)-1 {
		s.dfs(i+1, j, child)
	}
	if j < len(s.board[i])-1 {
		s.dfs(i, j+1, child)
	}
	s.visited[key] = false
}

func index(i, j, n int) int {
	return i*n + j
}

type Node struct {
	Word     string
	Matched  int
	Children []*Node
}

func newNode() *Node {
	return &Node{
		Children: make([]*Node, 26),
	}
}

/** Inserts a word into the trie. */
func (s *Solution) Insert(word string) {
	n := s.root
	for _, c := range []byte(word) {
		index := c - byte('a')
		child := n.Children[index]
		if child == nil {
			child = newNode()
			n.Children[index] = child
		}
		n = child
	}
	n.Word = word
}
