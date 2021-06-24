/*
运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制 。

实现 LRUCache 类：

    LRUCache(int capacity) 以正整数作为容量 capacity 初始化 LRU 缓存
    int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
    void put(int key, int value) 如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字-值」。当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。

*/
package lru

import (
	"fmt"
	"strings"
)

type LRUCache struct {
	head, tail *node
	m          map[int]*node
	capacity   int
}

type node struct {
	Key  int
	Val  int
	Prev *node
	Next *node
}

func Constructor(capacity int) LRUCache {
	head := &node{}
	tail := &node{
		Prev: head,
	}
	head.Next = tail
	return LRUCache{
		capacity: capacity,
		m:        make(map[int]*node),
		head:     head,
		tail:     tail,
	}
}

func (this *LRUCache) Get(key int) int {
	n := this.m[key]
	if n == nil {
		return -1
	}
	this.remove(n)
	this.putToHead(n)
	fmt.Printf("get %d:\n%s\n", key, this.String())
	return n.Val
}

func (this *LRUCache) Put(key int, value int) {
	n := &node{
		Key: key,
		Val: value,
	}
	if old, ok := this.m[key]; ok {
		this.remove(old)
	}
	this.putToHead(n)
	this.m[key] = n
	fmt.Printf("put %d m %v:\n%s\n", key, this.m, this.String())
	if len(this.m) > this.capacity {
		last := this.tail.Prev
		delete(this.m, last.Key)
		last.Prev.Next = this.tail
		this.tail.Prev = last.Prev
	}
	fmt.Printf("put delete %d m %v:\n%s\n", key, this.m, this.String())
}

func (this *LRUCache) remove(n *node) {
	n.Prev.Next = n.Next
	n.Next.Prev = n.Prev
}

func (this *LRUCache) putToHead(n *node) {
	n.Prev = this.head
	n.Next = this.head.Next
	n.Next.Prev = n
	this.head.Next = n
}

func (this *LRUCache) String() string {
	buf := strings.Builder{}
	n := this.head.Next
	for n.Next != nil {
		buf.WriteString(fmt.Sprintf("\tkey: %d, val: %d, prev: %d, next: %d\n", n.Key, n.Val, n.Prev.Key, n.Next.Key))
		n = n.Next
	}
	buf.WriteString(fmt.Sprintf("\ttail.prev: %d\n", this.tail.Prev.Key))
	return buf.String()
}
