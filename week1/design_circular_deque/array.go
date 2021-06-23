/*
设计实现双端队列。
你的实现需要支持以下操作：

    MyCircularDeque(k)：构造函数,双端队列的大小为k。
    insertFront()：将一个元素添加到双端队列头部。 如果操作成功返回 true。
    insertLast()：将一个元素添加到双端队列尾部。如果操作成功返回 true。
    deleteFront()：从双端队列头部删除一个元素。 如果操作成功返回 true。
    deleteLast()：从双端队列尾部删除一个元素。如果操作成功返回 true。
    getFront()：从双端队列头部获得一个元素。如果双端队列为空，返回 -1。
    getRear()：获得双端队列的最后一个元素。 如果双端队列为空，返回 -1。
    isEmpty()：检查双端队列是否为空。
    isFull()：检查双端队列是否满了。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/design-circular-deque
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package deque

import "fmt"

type ArrayDeque struct {
	data       []int
	head, tail int
	empty      bool
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func NewArrayDeque(k int) *ArrayDeque {
	return &ArrayDeque{
		data:  make([]int, k),
		head:  0,
		tail:  0,
		empty: true,
	}
}

func (d *ArrayDeque) String() string {
	return fmt.Sprintf("data: %v, head: %d, tail: %d", d.data, d.head, d.tail)
}

func (d *ArrayDeque) next(index int) int {
	index++
	if index == len(d.data) {
		index = 0
	}
	return index
}

func (d *ArrayDeque) prev(index int) int {
	index--
	if index < 0 {
		index = len(d.data) - 1
	}
	return index
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (d *ArrayDeque) InsertFront(value int) bool {
	if d.IsFull() {
		return false
	}
	if !d.IsEmpty() {
		d.head = d.prev(d.head)
	}
	d.data[d.head] = value
	d.empty = false
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (d *ArrayDeque) InsertLast(value int) bool {
	if d.IsFull() {
		return false
	}
	if !d.IsEmpty() {
		d.tail = d.next(d.tail)
	}
	d.data[d.tail] = value
	d.empty = false
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (d *ArrayDeque) DeleteFront() bool {
	if d.IsEmpty() {
		return false
	}
	if d.head == d.tail {
		d.empty = true
	} else {
		d.head = d.next(d.head)
	}
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (d *ArrayDeque) DeleteLast() bool {
	if d.IsEmpty() {
		return false
	}
	if d.head == d.tail {
		d.empty = true
	} else {
		d.tail = d.prev(d.tail)
	}
	return true
}

/** Get the front item from the deque. */
func (d *ArrayDeque) GetFront() int {
	if d.IsEmpty() {
		return -1
	}
	return d.data[d.head]
}

/** Get the last item from the deque. */
func (d *ArrayDeque) GetRear() int {
	if d.IsEmpty() {
		return -1
	}
	return d.data[d.tail]
}

/** Checks whether the circular deque is empty or not. */
func (d *ArrayDeque) IsEmpty() bool {
	return d.head == d.tail && d.empty
}

/** Checks whether the circular deque is full or not. */
func (d *ArrayDeque) IsFull() bool {
	if d.head-d.tail == 1 {
		return true
	}
	return d.tail == len(d.data)-1 && d.head == 0
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
