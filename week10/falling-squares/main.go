// https://leetcode-cn.com/problems/falling-squares/
package main

import (
	"fmt"
	"math"
	"sort"
)

func fallingSquares(positions [][]int) []int {
	values := make([]int, len(positions)*2)
	for i, item := range positions {
		values[i*2] = item[0]
		values[i*2+1] = item[0] + item[1] - 1 // 之前边界值处理总是不对，看了别人的题解加了个-1就可以了。但是想不出来为什么这个可以解决
	}
	sort.Ints(values)
	values = removeDup(values)
	t := NewSegmentTree(len(values))
	ans := make([]int, len(positions))
	for i, item := range positions {
		l, r := getIndex(values, item[0]), getIndex(values, item[0]+item[1]-1) // 之前边界值处理总是不对，看了别人的题解加了个-1就可以了。但是想不出来为什么这个可以解决
		val := item[1]
		if i > 0 {
			val += t.Find(l, r)
		}
		t.Add(val, l, r)
		ans[i] = t.Find(0, len(values)-1)
	}
	return ans
}

func removeDup(values []int) []int {
	j := 1
	for i := 1; i < len(values); i++ {
		if values[i] != values[i-1] {
			values[j] = values[i]
			j++
		}
	}
	return values[:j]
}

func getIndex(values []int, v int) int {
	l, r := 0, len(values)
	for l < r {
		mid := (l + r) / 2
		if values[mid] >= v {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}

type Node struct {
	l, r, max int
}

func (n *Node) String() string {
	return fmt.Sprintf("{%d, %d, %d}", n.max, n.l, n.r)
}

type SegmentTree struct {
	data []*Node
	size int
}

func NewSegmentTree(size int) *SegmentTree {
	t := &SegmentTree{
		data: make([]*Node, 4*size),
		size: size,
	}
	return t
}

func (t *SegmentTree) Find(l, r int) int {
	return t.query(0, l, r)
}

func (t *SegmentTree) query(curr, l, r int) int {
	if t.data[curr] == nil {
		return 0
	}
	if l <= t.data[curr].l && r >= t.data[curr].r {
		return t.data[curr].max
	}
	mid := (t.data[curr].l + t.data[curr].r) / 2
	max := int(math.MinInt32)
	if l <= mid {
		tmp := t.query(curr*2+1, l, r)
		if tmp > max {
			max = tmp
		}
	}
	if r > mid {
		tmp := t.query(curr*2+2, l, r)
		if tmp > max {
			max = tmp
		}
	}
	return max
}

func (t *SegmentTree) Add(val, l, r int) {
	t.add(0, val, l, r, 0, t.size-1)
}

func (t *SegmentTree) add(curr, val, nl, nr, l, r int) int {
	// fmt.Println("\t", t.data)
	// fmt.Println("\t", curr, val, nl, nr, l, r)
	if t.data[curr] == nil {
		t.data[curr] = &Node{
			l: l,
			r: r,
		}
	}
	if t.data[curr].l == t.data[curr].r {
		if t.data[curr].max < val {
			t.data[curr].max = val
		}
		return t.data[curr].max
	}
	mid := (l + r) / 2
	if nl <= mid {
		tmp := t.add(curr*2+1, val, nl, int(math.Min(float64(nr), float64(mid))), l, mid)
		if tmp > t.data[curr].max {
			t.data[curr].max = tmp
		}
	}
	if nr > mid {
		tmp := t.add(curr*2+2, val, int(math.Max(float64(nl), float64(mid+1))), nr, mid+1, r)
		if tmp > t.data[curr].max {
			t.data[curr].max = tmp
		}
	}
	return t.data[curr].max
}
