/*
有 n 个网络节点，标记为 1 到 n。

给你一个列表 times，表示信号经过 有向 边的传递时间。 times[i] = (ui, vi, wi)，其中 ui 是源节点，vi 是目标节点， wi 是一个信号从源节点传递到目标节点的时间。

现在，从某个节点 K 发出一个信号。需要多久才能使所有节点都收到信号？如果不能使所有节点收到信号，返回 -1 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/network-delay-time
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import "math"

func networkDelayTime(times [][]int, n int, k int) int {
	return bellmanFord(times, n, k)
}

func bellmanFord(times [][]int, n int, k int) int {
	delays := make([]int, n+1)
	delays[0] = int(math.MaxInt32) / 2
	for i := 1; i < len(delays); i *= 2 {
		copy(delays[i:], delays[:i])
	}
	delays[k] = 0
	for i := 0; i < n; i++ {
		updated := false
		for _, edge := range times {
			if delays[edge[1]] > delays[edge[0]]+edge[2] {
				delays[edge[1]] = delays[edge[0]] + edge[2]
				updated = true
			}
		}
		if !updated {
			break
		}
	}
	max := 0
	for _, delay := range delays[1:] {
		if delay > max {
			max = delay
		}
	}
	if max == delays[0] {
		return -1
	}
	return max
}

func dijkstra(times [][]int, n int, k int) int {
	metas := make([]meta, n+1)
	h := NewHeap(len(metas), func(t, o interface{}) int {
		i, j := t.(int), o.(int)
		return metas[i].delay - metas[j].delay
	})
	metas[k] = meta{
		delay:  0,
		marked: false,
	}
	h.Add(k)
	g := make([][]edge, n+1)
	for _, e := range times {
		g[e[0]] = append(g[e[0]], edge{
			end: e[1],
			t:   e[2],
		})
		for i := 0; i < 2; i++ {
			if e[i] == k || metas[e[i]].delay != 0 {
				continue
			}
			metas[e[i]].delay = int(math.MaxInt32) / 2
		}
	}
	for h.Top() != nil {
		i := h.Pop().(int)
		if metas[i].marked {
			continue
		}
		for _, e := range g[i] {
			if metas[e.end].delay > metas[i].delay+e.t {
				metas[e.end].delay = metas[i].delay + e.t
				h.Add(e.end)
			}
		}
		metas[i].marked = true
	}
	max := 0
	for _, meta := range metas[1:] {
		if !meta.marked {
			max = int(math.MaxInt32) / 2
		}
		if meta.delay > max {
			max = meta.delay
		}
	}
	if max == int(math.MaxInt32)/2 {
		return -1
	}
	return max
}

type edge struct {
	t   int // time
	end int
}

type meta struct {
	delay  int
	marked bool
}

type Heap struct {
	tail    int
	data    []interface{}
	indices map[interface{}]int
	compare func(t, o interface{}) int
}

// Variable size heap if size is 0
func NewHeap(size int, compare func(t, o interface{}) int) *Heap {
	return &Heap{
		tail:    -1,
		data:    make([]interface{}, size),
		indices: make(map[interface{}]int),
		compare: compare,
	}
}

func (h *Heap) Add(v interface{}) {
	h.tail++
	index := h.tail
	if len(h.data) == index {
		h.data = append(h.data, v)
	} else {
		h.data[index] = v
		h.indices[v] = index
	}
	for index > 0 && h.compare(h.data[(index-1)/2], v) < 0 {
		newIndex := (index - 1) / 2
		h.swap(index, newIndex)
		index = newIndex
	}
}

func (h *Heap) Top() interface{} {
	if h.tail == -1 {
		return nil
	}
	return h.data[0]
}

func (h *Heap) Pop() interface{} {
	if h.tail == -1 {
		return nil
	}
	result := h.data[0]
	h.data[0] = h.data[h.tail]
	h.indices[h.data[0]] = 0
	h.tail--
	index := 0
	left := index*2 + 1
	right := index*2 + 2
	for left <= h.tail {
		compareLeft := h.compare(h.data[index], h.data[left])
		if right > h.tail && compareLeft < 0 {
			h.swap(index, left)
			index = left
			break
		}
		if compareLeft > 0 && h.compare(h.data[index], h.data[right]) > 0 {
			break
		}
		if h.compare(h.data[left], h.data[right]) > 0 {
			h.swap(index, left)
			index = left
		} else {
			h.swap(index, right)
			index = right
		}
		left = index*2 + 1
		right = index*2 + 2
	}
	return result
}

func (h *Heap) swap(a, b int) {
	h.data[a], h.data[b] = h.data[b], h.data[a]
	h.indices[h.data[a]], h.indices[h.data[b]] = a, b
}
