package data

type MaxHeap[T any] struct {
	size    int
	data    []maxHeapItem[T]
	compare func(t, o T) int
	zero    T
}

type maxHeapItem[T any] struct {
	value T
	index int
}

// Variable size heap if size is 0
func NewMaxHeap[T any](size int, compare func(t, o T) int, zero T) *MaxHeap[T] {
	return &MaxHeap[T]{
		size:    size,
		data:    make([]maxHeapItem[T], 0, size),
		compare: compare,
		zero:    zero,
	}
}

func (h *MaxHeap[T]) Add(v T) {
	index := len(h.data)
	h.data = append(h.data, maxHeapItem[T]{
		value: v,
		index: index,
	})
	// Heapify up
	for index > 0 && h.compare(h.data[(index-1)/2].value, v) < 0 {
		newIndex := (index - 1) / 2
		h.swap(index, newIndex)
		index = newIndex
	}
}

func (h *MaxHeap[T]) Top() (T, bool) {
	if len(h.data) == 0 {
		return h.zero, false
	}
	h.removeOutdated()
	return h.data[0].value, true
}

func (h *MaxHeap[T]) removeOutdated() {
	for h.data[0].index < len(h.data)-h.size {
		h.pop()
	}
}

func (h *MaxHeap[T]) pop() (T, bool) {
	if len(h.data) == 0 {
		return h.zero, false
	}
	result := h.data[0].value
	tail := len(h.data) - 1
	h.data[0] = h.data[tail]
	h.data = h.data[:tail]
	// Heapify down
	index := 0
	child := index*2 + 1 // left child
	for child < tail {   // compare and swap with max child to make sure parent larger than all children
		if child+1 < tail && h.compare(h.data[child].value, h.data[child+1].value) < 0 { // choose right if it has a bigger right child
			child++
		}
		if h.compare(h.data[index].value, h.data[child].value) >= 0 { // already match condition
			break
		}
		h.swap(index, child)
		index = child
		child = index*2 + 1 // left child
	}
	return result, tail > 0
}

func (h *MaxHeap[T]) Pop() (T, bool) {
	h.removeOutdated()
	return h.pop()
}

func (h *MaxHeap[T]) swap(a, b int) {
	h.data[a], h.data[b] = h.data[b], h.data[a]
}
