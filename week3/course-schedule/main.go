/*
你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。

在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi] ，表示如果要学习课程 ai 则 必须 先学习课程  bi 。

    例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。

请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。

示例 1：

输入：numCourses = 2, prerequisites = [[1,0]]
输出：true
解释：总共有 2 门课程。学习课程 1 之前，你需要完成课程 0 。这是可能的。

示例 2：

输入：numCourses = 2, prerequisites = [[1,0],[0,1]]
输出：false
解释：总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0 ；并且学习课程 0 之前，你还应先完成课程 1 。这是不可能的。

提示：
    1 <= numCourses <= 105
    0 <= prerequisites.length <= 5000
    prerequisites[i].length == 2
    0 <= ai, bi < numCourses
    prerequisites[i] 中的所有课程对 互不相同

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/course-schedule
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import (
	"fmt"
	"strings"
)

var g *graph

func canFinish(numCourses int, prerequisites [][]int) bool {
	g = newGraph(numCourses)
	for _, edge := range prerequisites { // 构造图
		g.AddEdge(edge[1], edge[0])
	}
	learned := g.topsort()
	return learned == numCourses
}

func (g *graph) topsort() int {
	learned := 0
	q := &queue{}
	for x := range g.edge {
		if g.inDeg[x] == 0 {
			q.push(x)
		}
	}
	for !q.empty() {
		v := q.pop()
		learned++
		for _, y := range g.GetEdges(v) {
			g.inDeg[y]--
			if g.inDeg[y] == 0 {
				q.push(y)
			}
		}
	}
	return learned
}

type graph struct {
	edge  [][]int
	inDeg []int
}

func newGraph(num int) *graph {
	return &graph{
		edge:  make([][]int, num),
		inDeg: make([]int, num),
	}
}

func (g *graph) GetEdges(x int) []int {
	return g.edge[x]
}

func (g *graph) AddEdge(x, y int) {
	g.edge[x] = append(g.edge[x], y)
	g.inDeg[y]++
}

func (g *graph) String() string {
	sb := strings.Builder{}
	for x, line := range g.edge {
		sb.WriteString(fmt.Sprintf("%d->%v\n", x, line))
	}
	return sb.String()
}

type queue struct {
	data []int
}

func (q *queue) push(val int) {
	q.data = append(q.data, val)
}

func (q *queue) pop() int {
	v := q.data[0]
	q.data = q.data[1:]
	return v
}

func (q *queue) empty() bool {
	return len(q.data) == 0
}
