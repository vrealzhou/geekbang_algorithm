/*
在选举中，第 i 张票是在时间为 times[i] 时投给 persons[i] 的。

现在，我们想要实现下面的查询函数： TopVotedCandidate.q(int t) 将返回在 t 时刻主导选举的候选人的编号。

在 t 时刻投出的选票也将被计入我们的查询之中。在平局的情况下，最近获得投票的候选人将会获胜。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/online-election
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import "fmt"

type TopVotedCandidate struct {
	resultPerTime []int
	times         []int
}

// 思路：因为投票结果预先已经确定了，所以可以在创建TopVotedCandidate的时候
// 就把每个时间点的赢家保存下来。每次执行Q的时候直接取值就可以了
func Constructor(persons []int, times []int) TopVotedCandidate {
	voteCount := make([]int, len(persons))     // 每个人得票统计
	resultPerTime := make([]int, len(persons)) // 每个时间点的结果
	h := newHeap(len(persons))                 // 用大根堆取得赢家
	for i := 0; i < len(voteCount); i++ {      // 统计每一个时间点的赢家
		voteCount[persons[i]] += 1 // 当前时间点每个人得票统计
		h.Push(&count{persons[i], voteCount[persons[i]]})
		top := h.Top()
		if top.vote == voteCount[persons[i]] && top.person != persons[i] { // 如果票数相同，当前得票的人赢
			resultPerTime[i] = persons[i]
			h.SwapPerson(top.person, persons[i]) // 需要交换堆里元素，为下一个时间点判断作准备
		} else {
			resultPerTime[i] = top.person
		}
		// fmt.Println(i, voteCount, h.data, h.dataMap, resultPerTime)
	}
	return TopVotedCandidate{
		resultPerTime: resultPerTime,
		times:         times,
	}
}

func (this *TopVotedCandidate) Q(t int) int {
	// 二分查找最后一个 >= t 的值的位置
	l := -1
	r := len(this.times) - 1
	for l < r {
		mid := (l + r + 1) / 2
		if this.times[mid] <= t {
			l = mid
		} else {
			r = mid - 1
		}
	}
	// fmt.Println(t, r, this.times[r])
	return this.resultPerTime[r]
}

type count struct {
	person int
	vote   int
}

func (c *count) String() string {
	return fmt.Sprintf("(%d, %d)", c.person, c.vote)
}

type maxHeap struct {
	tail      int
	data      []*count
	dataIndex []int // 保存了person->index信息。可以快速跟据person取得*count
}

func newHeap(size int) *maxHeap {
	dataIndex := make([]int, size)
	for i := 0; i < size; i++ {
		dataIndex[i] = -1
	}
	return &maxHeap{
		tail:      -1,
		data:      make([]*count, size),
		dataIndex: dataIndex,
	}
}

func (h *maxHeap) SwapPerson(person1, person2 int) {
	index1 := h.dataIndex[person1]
	index2 := h.dataIndex[person2]
	h.swap(index1, index2)
}

func (h *maxHeap) swap(index1, index2 int) {
	// 每次交换都需要交换dataMap
	h.dataIndex[h.data[index1].person] = index2
	h.dataIndex[h.data[index2].person] = index1
	h.data[index1], h.data[index2] = h.data[index2], h.data[index1]
}

func (h *maxHeap) Push(v *count) {
	// 先查询person是否已经存在，不存在的话先添加在尾部
	index := h.dataIndex[v.person]
	if index < 0 {
		h.tail++
		index = h.tail
	}
	h.data[index] = v
	h.dataIndex[v.person] = index
	for index > 0 && h.data[(index-1)/2].vote < v.vote {
		newIndex := (index - 1) / 2
		h.swap(index, newIndex)
		index = newIndex
	}
}

func (h *maxHeap) Top() *count {
	if h.tail == -1 {
		return nil
	}
	return h.data[0]
}
