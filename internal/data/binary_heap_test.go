package data

import "testing"

func TestBinaryHeap(t *testing.T) {
	heap := NewHeap(5, func(t, o interface{}) int {
		return t.(int) - o.(int)
	})
	heap.Add(10)
	if heap.Top() != 10 {
		t.Errorf("Incorrect top value: %d", heap.Top())
		t.FailNow()
	}
	heap.Add(20)
	if heap.Top() != 20 {
		t.Errorf("Incorrect top value: %d", heap.Top())
		t.FailNow()
	}
	heap.Add(4)
	if heap.Top() != 20 {
		t.Errorf("Incorrect top value: %d", heap.Top())
		t.FailNow()
	}
	heap.Add(1)
	heap.Add(5)
	heap.Add(6)
	heap.Add(8)
	if heap.Top() != 8 {
		t.Errorf("Incorrect top value: %d", heap.Top())
		t.FailNow()
	}
	heap.Add(21)
	if heap.Top() != 21 {
		t.Errorf("Incorrect top value: %d", heap.Top())
		t.FailNow()
	}
}
