/*
现在你总共有 n 门课需要选，记为 0 到 n-1。

在选修某些课程之前需要一些先修课程。 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们: [0,1]

给定课程总量以及它们的先决条件，返回你为了学完所有课程所安排的学习顺序。

可能会有多个正确的顺序，你只要返回一种就可以了。如果不可能完成所有课程，返回一个空数组。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/course-schedule-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

func findOrder(numCourses int, prerequisites [][]int) []int {
	g := newGraph(numCourses)
	for _, p := range prerequisites {
		g.AddEdge(p[1], p[0])
	}
	learned := g.TopSort()
	if len(learned) == numCourses {
		return learned
	}
	return []int{}
}

type Graph struct {
	edges [][]int
	inDeg []int
}

func newGraph(size int) *Graph {
	return &Graph{
		edges: make([][]int, size),
		inDeg: make([]int, size),
	}
}

func (g *Graph) AddEdge(x, y int) {
	g.edges[x] = append(g.edges[x], y)
	g.inDeg[y]++
}

func (g *Graph) Edges(x int) []int {
	return g.edges[x]
}

func (g *Graph) TopSort() []int {
	q := &Queue{}
	for x := range g.edges { // O(n). n = total nodes
		if g.inDeg[x] == 0 {
			q.Push(x)
		}
	}
	learned := make([]int, 0)
	for !q.Empty() { // O(m). m = total edges
		x := q.Pop()
		learned = append(learned, x)
		for _, y := range g.Edges(x) {
			g.inDeg[y]--
			if g.inDeg[y] == 0 {
				q.Push(y)
			}
		}
	}
	return learned
}

type Queue struct {
	data []int
}

func (q *Queue) Push(v int) {
	q.data = append(q.data, v)
}

func (q *Queue) Pop() int {
	v := q.data[0]
	q.data = q.data[1:]
	return v
}

func (q *Queue) Empty() bool {
	return len(q.data) == 0
}
