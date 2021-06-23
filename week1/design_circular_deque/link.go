package deque

import "fmt"

type node struct {
	Val  int
	Prev *node
	Next *node
}

type ListDeque struct {
	head, tail  *node
	limit, size int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func NewListDeque(k int) *ListDeque {
	return &ListDeque{
		limit: k,
		size:  0,
		head:  nil,
		tail:  nil,
	}
}

func (d *ListDeque) String() string {
	l := make([]int, 0)
	n := d.head
	for n != nil {
		l = append(l, n.Val)
		if n.Next != nil {
			n = n.Next
		} else {
			break
		}
	}
	return fmt.Sprintf("data: %v", l)
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (d *ListDeque) InsertFront(value int) bool {
	if d.IsFull() {
		return false
	}
	n := &node{
		Val:  value,
		Next: d.head,
	}
	if d.head == nil { // from empty to not empty
		d.tail = n
	} else {
		d.head.Prev = n
	}
	d.head = n
	d.size++
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (d *ListDeque) InsertLast(value int) bool {
	if d.IsFull() {
		return false
	}
	n := &node{
		Val:  value,
		Prev: d.tail,
	}
	if d.tail == nil { // from empty to not empty
		d.head = n
	} else {
		d.tail.Next = n
	}
	d.tail = n
	d.size++
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (d *ListDeque) DeleteFront() bool {
	if d.IsEmpty() {
		return false
	}
	if d.head == d.tail { // last one
		d.head = nil
		d.tail = nil
	} else {
		d.head = d.head.Next
	}
	d.size--
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (d *ListDeque) DeleteLast() bool {
	if d.IsEmpty() {
		return false
	}
	if d.tail == d.head { // last one
		d.head = nil
		d.tail = nil
	} else {
		d.tail = d.tail.Prev
	}
	d.size--
	return true
}

/** Get the front item from the deque. */
func (d *ListDeque) GetFront() int {
	if d.head == nil {
		return -1
	}
	return d.head.Val
}

/** Get the last item from the deque. */
func (d *ListDeque) GetRear() int {
	if d.tail == nil {
		return -1
	}
	return d.tail.Val
}

/** Checks whether the circular deque is empty or not. */
func (d *ListDeque) IsEmpty() bool {
	return d.size == 0
}

/** Checks whether the circular deque is full or not. */
func (d *ListDeque) IsFull() bool {
	return d.size == d.limit
}
