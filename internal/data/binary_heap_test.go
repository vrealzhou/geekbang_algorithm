package data

import "testing"

func TestBinaryHeap(t *testing.T) {
	heap := NewMaxHeap(5, func(t, o int) int {
		return t - o
	}, 0)
	heap.Add(10)
	if r, _ := heap.Top(); r != 10 {
		t.Errorf("Incorrect top value: %d", r)
		t.FailNow()
	}
	heap.Add(20)
	if r, _ := heap.Top(); r != 20 {
		t.Errorf("Incorrect top value: %d", r)
		t.FailNow()
	}
	heap.Add(4)
	if r, _ := heap.Top(); r != 20 {
		t.Errorf("Incorrect top value: %d", r)
		t.FailNow()
	}
	heap.Add(1)
	heap.Add(5)
	heap.Add(6)
	heap.Add(8)
	if r, _ := heap.Top(); r != 8 {
		t.Errorf("Incorrect top value: %d", r)
		t.FailNow()
	}
	heap.Add(21)
	if r, _ := heap.Top(); r != 21 {
		t.Errorf("Incorrect top value: %d", r)
		t.FailNow()
	}
}

func TestBinaryHeapPop(t *testing.T) {
	heap := NewMaxHeap(5, func(t, o int) int {
		return t - o
	}, 0)
	heap.Add(10)
	heap.Add(1)
	heap.Add(5)
	heap.Add(6)
	heap.Add(8)
	if v, _ := heap.Top(); v != 10 {
		t.Errorf("Incorrect top value: %d", v)
		t.FailNow()
	}
	v, more := heap.Pop()
	if v != 10 || !more {
		t.Errorf("Incorrect top value: %d, %t", v, more)
		t.FailNow()
	}
	if v, _ := heap.Top(); v != 8 {
		t.Errorf("Incorrect top value: %d", v)
		t.FailNow()
	}
	v, more = heap.Pop()
	if v != 8 || !more {
		t.Errorf("Incorrect top value: %d, %t", v, more)
		t.FailNow()
	}
	v, more = heap.Pop()
	if v != 6 || !more {
		t.Errorf("Incorrect top value: %d, %t", v, more)
		t.FailNow()
	}
	v, more = heap.Pop()
	if v != 5 || !more {
		t.Errorf("Incorrect top value: %d, %t", v, more)
		t.FailNow()
	}
	v, more = heap.Pop()
	if v != 1 || more {
		t.Errorf("Incorrect top value: %d, %t", v, more)
		t.FailNow()
	}
}
